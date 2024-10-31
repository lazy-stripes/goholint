package widgets

import (
	"fmt"

	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ppu"
	"github.com/veandco/go-sdl2/sdl"
)

// TODO: make this a feature of Screen instead. Either that or we need yet
//       another interface for UI display (which would embed Widget,
//       screen.Display and probably ui.Display or some Messager interface).
type DebugScreen struct {
	*Screen

	PPU *ppu.PPU
}

func NewDebugScreen(sizeHint *sdl.Rect, config *options.Options) *DebugScreen {
	return &DebugScreen{Screen: NewScreen(sizeHint, config)}
}

func (s *DebugScreen) Texture() *sdl.Texture {
	// Draw sprite boudaries to texture.
	var spriteRect sdl.Rect
	spriteRect.W = 8 * int32(s.Zoom)
	if s.PPU.LCDC&ppu.LCDCSpriteSize != 0 {
		spriteRect.H = 16 * int32(s.Zoom)
	} else {
		spriteRect.H = 8 * int32(s.Zoom)
	}

	t := s.Screen.Texture()
	renderer.SetDrawColor(
		s.BorderColor.R,
		s.BorderColor.G,
		s.BorderColor.B,
		s.BorderColor.A,
	)
	renderer.SetRenderTarget(t)

	oam := s.PPU.OAM.Bytes
	for i := 0; i < len(oam); i += 4 {
		y := int32(oam[i+0]) - 16
		x := int32(oam[i+1]) - 8
		//tile := oam[i+2]
		//flags := oam[i+3]

		spriteRect.X = x * int32(s.Zoom)
		spriteRect.Y = y * int32(s.Zoom)
		renderer.DrawRect(&spriteRect)
	}
	renderer.SetRenderTarget(nil)

	s.Text(fmt.Sprintf("BGP: %08b", s.PPU.BGP))

	return t
}
