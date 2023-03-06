package screen

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"time"
	"unsafe"

	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/options"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// SDL display shifting pixels out to a single texture.
type SDL struct {
	*UI

	config *options.Options

	palette    []color.RGBA
	newPalette []color.RGBA // Store new value until next frame

	enabled    bool
	window     *sdl.Window
	renderer   *sdl.Renderer
	texture    *sdl.Texture
	blank      *sdl.Texture
	buffer     []byte
	offset     int
	zoom       int // Zoom factor applied to the 144×160 screen.
	screenRect image.Rectangle

	// Set this to true to save the next frame. Will be reset at VBlank.
	screenshotRequested bool

	// GIF recorder. TODO: record video with sound too.
	gif            *GIF
	startRecording bool
	stopRecording  bool
	recordTime     time.Time
}

// NewSDL returns an SDL2 display with a greyish palette and takes a zoom
// factor to size the window (current default is 2x).
func NewSDL(config *options.Options) *SDL {
	// TODO: subfunctions, this is already too big.
	window, err := sdl.CreateWindow("Goholint",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth*int32(config.ZoomFactor), ScreenHeight*int32(config.ZoomFactor),
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

	if config.VSync {
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

	// Go bindings use byte slices but SDL thinks in terms of uint32
	screenLen := ScreenWidth * ScreenHeight * 4
	buffer := make([]byte, screenLen)

	// Keep computed screen size for screenshots.
	screenRect := image.Rectangle{
		image.Point{0, 0},
		image.Point{ScreenWidth * int(config.ZoomFactor), ScreenHeight * int(config.ZoomFactor)},
	}

	// Create UI with actual screen size and colors from config.
	ui := NewUI(renderer, config)

	s := SDL{
		UI:         ui,
		config:     config,
		palette:    config.Palettes[0],
		renderer:   renderer,
		texture:    texture,
		blank:      blank,
		buffer:     buffer,
		zoom:       int(config.ZoomFactor),
		screenRect: screenRect,
		gif:        NewGIF(config),
	}

	// Initialize blank screen texture from color 0.
	renderer.SetRenderTarget(blank)
	renderer.SetDrawColor(s.palette[0].R, s.palette[0].G, s.palette[0].B, 0xff)
	renderer.Clear()
	renderer.SetRenderTarget(nil)

	// Init texture and trigger stuff usually happening at VBlank.
	s.VBlank() // XXX: is this needed?

	return &s
}

// Close frees all resources created by SDL.
func (s *SDL) Close() {
	s.texture.Destroy()
	s.blank.Destroy()
	s.renderer.Destroy()
	s.window.Destroy()
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
		col := s.palette[colorIndex]
		s.buffer[s.offset+0] = col.R
		s.buffer[s.offset+1] = col.G
		s.buffer[s.offset+2] = col.B
		s.buffer[s.offset+3] = col.A
		s.offset += 4

		if s.gif.IsOpen() {
			s.gif.Write(colorIndex)
		}
	}
}

// HBlank is only there as part of the Display interface and has no use in this
// context (yet?).
func (s *SDL) HBlank() {}

// VBlank is called when the PPU reaches VBlank state. At this point, our SDL
// buffer should be ready to display.
func (s *SDL) VBlank() {
	if s.enabled {
		// SDL bindings used to accept a slice but no longer do as of 0.4.33.
		rawPixels := unsafe.Pointer(&s.buffer[0])
		s.texture.Update(nil, rawPixels, ScreenWidth*4)

		if s.offset != ScreenWidth*ScreenHeight*4 {
			log.Warning("MISSING PIXELS!")
		}
		s.offset = 0
	} else {
		// Draw blank screen texture.
		s.renderer.SetRenderTarget(s.texture)
		s.renderer.SetDrawColor(s.palette[0].R, s.palette[0].G, s.palette[0].B, 0xff)
		s.renderer.Clear()
		s.renderer.SetRenderTarget(nil)

	}
	s.renderer.Copy(s.texture, nil, nil)

	// Update GIF frame if recording. We do this before checking startRecording
	// otherwise the call to SaveFrame will always insert a "disabled" frame in
	// first position (since we haven't yet had time to build a full frame in
	// that specific case).
	if s.gif.IsOpen() {
		d := time.Since(s.recordTime)
		text := fmt.Sprintf("•REC [%02d:%02d]", d/time.Minute, d/time.Second)
		s.UI.Text(text)
		s.gif.SaveFrame()
	}

	// Create GIF here if requested.
	if s.startRecording {
		f, err := options.CreateFileIn("gifs", ".gif")
		if err == nil {
			s.startRecording = false
			s.recordTime = time.Now()
			s.UI.Text("•REC [00:00]")
			s.gif.New(f, s.palette)

			fmt.Printf("Recording GIF to %s\n", f.Name())
		} else {
			log.Warningf("creating gif file failed: %v", err)
		}
	}

	if s.stopRecording {
		s.stopRecording = false
		s.gif.Close()
		s.UI.Text("")
		s.UI.Message(fmt.Sprintf("%d frames saved", len(s.gif.GIF.Image)), 2)
	}

	// UI overlay.
	if s.UI.Enabled {
		//s.UI.texture.SetBlendMode(sdl.BLENDMODE_ADD)
		s.renderer.Copy(s.UI.texture, nil, nil)
	}

	s.renderer.Present()

	if s.screenshotRequested {
		s.screenshotRequested = false

		f, err := options.CreateFileIn("screenshots", ".png")
		if err != nil {
			log.Warningf("creating screenshot file failed: %v", err)
			return
		}
		defer f.Close()

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

		if err := png.Encode(f, img); err != nil {
			log.Warningf("saving screenshot failed: %v", err)
			return
		}

		s.Message("Screenshot saved", 2)
		fmt.Printf("Screenshot saved to %s\n", f.Name())
	}

	if s.newPalette != nil {
		s.palette = s.newPalette
		s.newPalette = nil
	}
}

// Dump writes the current pixel buffer to file for debugging purposes.
func (s *SDL) Dump() {
	ioutil.WriteFile("lcd-buffer-dump.bin", s.buffer, 0644)
}

// Screenshot will make the display dump the next frame to file.
func (s *SDL) Screenshot() {
	s.screenshotRequested = true
}

// StartRecord will create a new GIF file and output frames into it until
// StopRecord is called.
//
// We only just raise a flag here, recording should start and stop in VBlank.
func (s *SDL) StartRecord() {
	if s.gif.IsOpen() {
		log.Warningf("recording to %s already in progress", s.gif.Filename)
		return
	}
	s.startRecording = true
}

// StopRecord will flush recorded frames to the previously created GIF file.
// We only just raise a flag here, recording should start and stop in VBlank.
func (s *SDL) StopRecord() {
	s.stopRecording = true
}

// Palette will set a new palette for the display and GIF.
func (s *SDL) Palette(p []color.RGBA) {
	// Wait until next frame to apply new palette.
	s.newPalette = p
}
