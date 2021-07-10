package screen

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// SDL display shifting pixels out to a single texture.
type SDL struct {
	Palette    color.Palette
	enabled    bool
	window     *sdl.Window
	renderer   *sdl.Renderer
	texture    *sdl.Texture
	blank      *sdl.Texture
	buffer     []byte
	offset     int
	zoom       int // Zoom factor applied to the 144×160 screen.
	screenRect image.Rectangle

	// Set this to non-empty to save the next frame. Will be reset at VBlank.
	screenshotPath string
}

var testPalette = [4]color.NRGBA{
	{0xbd, 0xff, 0x9d, 0xff},
	{0xff, 0xaa, 0x00, 0xff},
	{0x00, 0xaa, 0xff, 0xff},
	{0xff, 0x00, 0x00, 0xff},
}

// NewSDL returns an SDL2 display with a greyish palette and takes a zoom
// factor to size the window (current default is 2x).
func NewSDL(zoomFactor uint, vSync bool) *SDL {
	window, err := sdl.CreateWindow("Goholint",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth*int32(zoomFactor), ScreenHeight*int32(zoomFactor),
		sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return nil // TODO: result, err
	}

	// FIXME: embed assets.
	icon, err := img.Load("assets/icon.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load icon: %s\n", err)
	} else {
		window.SetIcon(icon)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return nil // TODO: result, err
	}

	if vSync {
		if err = sdl.GLSetSwapInterval(-1); err != nil {
			log.Infof("Can't set adaptive vsync: %s", sdl.GetError())
			// Try 'just' syncing to vblank then.
			if err = sdl.GLSetSwapInterval(1); err != nil {
				log.Warningf("Can't sync to vblank: %s", sdl.GetError())
			}
		}
	}

	if info, err := renderer.GetInfo(); err == nil {
		log.Info("SDL_RENDERER info:")
		log.Infof("SOFTWARE: %t", info.Flags&sdl.RENDERER_SOFTWARE != 0)
		log.Infof("ACCELERATED: %t", info.Flags&sdl.RENDERER_ACCELERATED != 0)
		log.Infof("PRESENTVSYNC: %t", info.Flags&sdl.RENDERER_PRESENTVSYNC != 0)
	}

	// The way SDL textures handle endianness is unclear, but it seems ABGR
	// format works with our RGBA buffer.
	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STATIC,
		ScreenWidth,
		ScreenHeight)
	if err != nil {
		renderer.Destroy()
		window.Destroy()
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return nil // TODO: result, err
	}

	// Also pre-instantiate disabled screen texture as it'll be used a lot.
	blank, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		ScreenWidth,
		ScreenHeight)
	if err != nil {
		texture.Destroy()
		renderer.Destroy()
		window.Destroy()
		fmt.Fprintf(os.Stderr, "Failed to create blank texture: %s\n", err)
		return nil // TODO: result, err
	}
	renderer.SetRenderTarget(blank)
	renderer.SetDrawColor(ColorWhiteR, ColorWhiteG, ColorWhiteB, sdl.ALPHA_OPAQUE)
	renderer.Clear()
	renderer.SetRenderTarget(nil)

	// Go bindings use byte slices but SDL thinks in terms of uint32
	screenLen := ScreenWidth * ScreenHeight * 4
	buffer := make([]byte, screenLen)

	// Keep computed screen size for screenshots.
	screenRect := image.Rectangle{
		image.Point{0, 0},
		image.Point{ScreenWidth * int(zoomFactor), ScreenHeight * int(zoomFactor)},
	}

	sdl := SDL{Palette: DefaultPalette, renderer: renderer, texture: texture,
		blank: blank, buffer: buffer, zoom: int(zoomFactor), screenRect: screenRect}

	sdl.Clear()

	return &sdl
}

// Close frees all resources created by SDL.
func (s *SDL) Close() {
	s.texture.Destroy()
	s.blank.Destroy()
	s.renderer.Destroy()
	s.window.Destroy()
}

// Clear draws a disabled GB screen (white background).
func (s *SDL) Clear() {
	s.renderer.Copy(s.blank, nil, nil)
	s.renderer.Present()
}

// Enable turns on the display. Pixels will be drawn to our texture and showed at VBlank time.
func (s *SDL) Enable() {
	s.enabled = true
}

// Enabled returns whether the display is enabled or not (as part of the Display interface).
func (s *SDL) Enabled() bool {
	return s.enabled
}

// Disable turns off the display. A disabled GB screen will be drawn at VBlank time.
func (s *SDL) Disable() {
	s.offset = 0
	s.enabled = false
}

// Write adds a new pixel (a mere index into a palette) to the texture buffer.
func (s *SDL) Write(colorIndex uint8) {
	if s.enabled {
		s.buffer[s.offset+0] = s.Palette[colorIndex].(color.NRGBA).R
		s.buffer[s.offset+1] = s.Palette[colorIndex].(color.NRGBA).G
		s.buffer[s.offset+2] = s.Palette[colorIndex].(color.NRGBA).B
		s.buffer[s.offset+3] = s.Palette[colorIndex].(color.NRGBA).A
		s.offset += 4
	}
}

// HBlank is only there as part of the Display interface and has no use in this
// context (yet?).
func (s *SDL) HBlank() {}

// VBlank is called when the PPU reaches VBlank state. At this point, our SDL
// buffer should be ready to display.
func (s *SDL) VBlank() {
	if s.enabled {
		s.texture.Update(nil, s.buffer, ScreenWidth*4)
		s.renderer.Copy(s.texture, nil, nil)
		s.renderer.Present()

		if s.offset != ScreenWidth*ScreenHeight*4 {
			log.Warning("MISSING PIXELS!")
		}
		s.offset = 0
	} else {
		s.Clear() // Phase out Clear()
	}

	// TODO: UI overlay here?
	//if s.ui.Enabled {
	//	s.renderer.Copy(s.ui.Texture, nil, nil)
	//}

	if s.screenshotPath != "" {
		// Reset screenshotPath for next call.
		path := s.screenshotPath
		s.screenshotPath = ""

		// Populate image from buffer, taking zoom into account.
		img := image.NewRGBA(s.screenRect)
		for x := 0; x < img.Rect.Dx(); x++ {
			for y := 0; y < img.Rect.Dy(); y++ {
				srcX := x / s.zoom
				srcY := y / s.zoom
				srcOffset := srcY*ScreenWidth*4 + srcX*4
				dstOffset := y*ScreenWidth*s.zoom*4 + x*4

				// Copy RGBA components.
				img.Pix[dstOffset+0] = s.buffer[srcOffset+0]
				img.Pix[dstOffset+1] = s.buffer[srcOffset+1]
				img.Pix[dstOffset+2] = s.buffer[srcOffset+2]
				img.Pix[dstOffset+3] = s.buffer[srcOffset+3]
			}
		}

		f, err := os.Create(path)

		if err != nil {
			log.Warningf("creating screenshot file failed: %v", err)
			return
		}
		defer f.Close()

		if err := png.Encode(f, img); err != nil {
			log.Warningf("saving screenshot failed: %v", err)
			return
		}

		fmt.Printf("screenshot saved to %s\n", path)
	}
}

// Blank draws an empty GB screen when the display is disabled.
func (s *SDL) Blank() {
	s.Clear()
}

// Dump writes the current pixel buffer to file for debugging purposes.
func (s *SDL) Dump() {
	ioutil.WriteFile("lcd-buffer-dump.bin", s.buffer, 0644)
}

// Screenshot will make the display dump the next frame to file.
func (s *SDL) Screenshot(filename string) {
	s.screenshotPath = filename
}
