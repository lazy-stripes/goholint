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
	Palette  Palette
	enabled  bool
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
	buffer   []byte
	offset   int
}

var testPalette = [4]color.NRGBA{
	color.NRGBA{0xbd, 0xff, 0x9d, 0xff},
	color.NRGBA{0xff, 0xaa, 0x00, 0xff},
	color.NRGBA{0x00, 0xaa, 0xff, 0xff},
	color.NRGBA{0xff, 0x00, 0x00, 0xff},
}

// NewSDL returns an SDL2 display with a greyish palette.
func NewSDL() *SDL {
	/*
		if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
			fmt.Fprintf(os.Stderr, "SDL Init failed: %s\n", err)
			return nil // TODO: result, err
		}
	*/
	window, err := sdl.CreateWindow("gb.go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth, ScreenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return nil // TODO: result, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if info, err := renderer.GetInfo(); err == nil {
		fmt.Println("Renderer info:")
		fmt.Printf("SDL_RENDERER_SOFTWARE: %t\n", info.Flags&sdl.RENDERER_SOFTWARE != 0)
		fmt.Printf("SDL_RENDERER_ACCELERATED: %t\n", info.Flags&sdl.RENDERER_ACCELERATED != 0)
		fmt.Printf("SDL_RENDERER_PRESENTVSYNC: %t\n", info.Flags&sdl.RENDERER_PRESENTVSYNC != 0)
	}
	if err != nil {
		window.Destroy()
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return nil // TODO: result, err
	}

	// The way SDL textures handle endianness is unclear, but it seems ABGR format works with our RGBA buffer.
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, ScreenWidth, ScreenHeight)
	if err != nil {
		renderer.Destroy()
		window.Destroy()
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return nil // TODO: result, err
	}

	screenLen := ScreenWidth * ScreenHeight * 4 // Each pixel fits in uint32
	buffer := make([]byte, screenLen, screenLen)
	sdl := SDL{Palette: testPalette, renderer: renderer, texture: texture, buffer: buffer}
	sdl.Clear()
	return &sdl
}

func (s *SDL) Close() {
	s.texture.Destroy()
	s.renderer.Destroy()
	s.window.Destroy()
}

func (s *SDL) Clear() {
	s.renderer.SetDrawColor(ColorWhiteR, ColorWhiteG, ColorWhiteB, sdl.ALPHA_OPAQUE)
	s.renderer.Clear()
	s.renderer.SetDrawColor(ColorBlackR, ColorBlackG, ColorBlackB, sdl.ALPHA_OPAQUE)
	s.renderer.DrawLine(0, ScreenHeight/2, ScreenWidth, ScreenHeight/2)
	s.renderer.Present()
}

func (s *SDL) Enable() {
	s.enabled = true
}

func (s *SDL) Enabled() bool {
	return s.enabled
}

func (s *SDL) Disable() {
	s.enabled = false
}

func (s *SDL) Write(pixel Pixel) {
	if s.enabled {
		s.buffer[s.offset] = s.Palette[pixel].R
		s.buffer[s.offset+1] = s.Palette[pixel].G
		s.buffer[s.offset+2] = s.Palette[pixel].B
		s.buffer[s.offset+3] = s.Palette[pixel].A
		s.offset += 4
	}
}

func (s *SDL) HBlank() {
	// XXX: Do we truly need this function?
}

func (s *SDL) VBlank() {
	if s.enabled {
		s.texture.Update(nil, s.buffer, ScreenWidth*4)
		s.renderer.Clear()
		s.renderer.Copy(s.texture, nil, nil)
		s.renderer.Present()
		//for t := time.Now(); time.Now().Sub(t) < time.Nanosecond*400; {
		//}

		if s.offset != ScreenWidth*ScreenHeight*4 {
			fmt.Println("MISSING PIXELS!")
		}
		s.offset = 0
	} else {
		s.Clear()
	}
}

func (s *SDL) Blank() {
	s.Clear()
}

func (s *SDL) Dump() {
	ioutil.WriteFile("lcd-buffer-dump.bin", s.buffer, 0644)
}
