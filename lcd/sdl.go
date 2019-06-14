package lcd

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// SDL display shifting pixels out to a single texture.
type SDL struct {
	Palette  color.Palette
	enabled  bool
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
	blank    *sdl.Texture
	buffer   []byte
	offset   int
}

var testPalette = [4]color.NRGBA{
	color.NRGBA{0xbd, 0xff, 0x9d, 0xff},
	color.NRGBA{0xff, 0xaa, 0x00, 0xff},
	color.NRGBA{0x00, 0xaa, 0xff, 0xff},
	color.NRGBA{0xff, 0x00, 0x00, 0xff},
}

// NewSDL returns an SDL2 display with a greyish palette and takes a zoom
// factor to size the window (current default is 2x).
func NewSDL(zoomFactor uint8) *SDL {
	/*
		if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
			fmt.Fprintf(os.Stderr, "SDL Init failed: %s\n", err)
			return nil // TODO: result, err
		}
	*/
	window, err := sdl.CreateWindow("gb.go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth*int32(zoomFactor), ScreenHeight*int32(zoomFactor),
		sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return nil // TODO: result, err
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return nil // TODO: result, err
	}

	if err = sdl.GLSetSwapInterval(-1); err != nil {
		log.Infof("Can't set adaptive vsync: %s", sdl.GetError())
		// Try 'just' syncing to vblank then.
		if err = sdl.GLSetSwapInterval(1); err != nil {
			log.Warningf("Can't sync to vblank: %s", sdl.GetError())
		}
	}
	if info, err := renderer.GetInfo(); err == nil {
		log.Info("Renderer info:")
		log.Infof("SDL_RENDERER_SOFTWARE: %t", info.Flags&sdl.RENDERER_SOFTWARE != 0)
		log.Infof("SDL_RENDERER_ACCELERATED: %t", info.Flags&sdl.RENDERER_ACCELERATED != 0)
		log.Infof("SDL_RENDERER_PRESENTVSYNC: %t", info.Flags&sdl.RENDERER_PRESENTVSYNC != 0)
	}

	// The way SDL textures handle endianness is unclear, but it seems ABGR format works with our RGBA buffer.
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, ScreenWidth, ScreenHeight)
	if err != nil {
		renderer.Destroy()
		window.Destroy()
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return nil // TODO: result, err
	}

	// Also pre-instantiate disabled screen texture as it'll be used a lot.
	blank, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, ScreenWidth, ScreenHeight)
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
	renderer.SetDrawColor(ColorBlackR, ColorBlackG, ColorBlackB, sdl.ALPHA_OPAQUE)
	renderer.DrawLine(0, ScreenHeight/2, ScreenWidth, ScreenHeight/2)
	renderer.SetRenderTarget(nil)

	screenLen := ScreenWidth * ScreenHeight * 4 // Go bindings use byte slices but SDL thinks in terms of uint32
	buffer := make([]byte, screenLen)
	sdl := SDL{Palette: DefaultPalette, renderer: renderer, texture: texture, blank: blank, buffer: buffer}
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

// Clear draws a disabled GB screen (white background with a black line through the middle).
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

// Write adds a new pixel (a mere index into our screen palette) to the texture buffer.
func (s *SDL) Write(colorIndex uint8) {
	if s.enabled {
		s.buffer[s.offset] = s.Palette[colorIndex].(color.NRGBA).R
		s.buffer[s.offset+1] = s.Palette[colorIndex].(color.NRGBA).G
		s.buffer[s.offset+2] = s.Palette[colorIndex].(color.NRGBA).B
		s.buffer[s.offset+3] = s.Palette[colorIndex].(color.NRGBA).A
		s.offset += 4
	}
}

// HBlank is only there as part of the Display interface and has no use in this context (yet?).
func (s *SDL) HBlank() {
}

// VBlank is called when the PPU reaches VBlank state. At this point, our SDL buffer should be ready to display.
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
		s.Clear()
	}
}

// Blank is called on each PPU step when the display is disabled, drawing the disabled GB screen.
func (s *SDL) Blank() {
	s.Clear()
}

// Dump writes the current pixel buffer to file for debugging purposes.
func (s *SDL) Dump() {
	ioutil.WriteFile("lcd-buffer-dump.bin", s.buffer, 0644)
}
