package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Properties stores useful configuration variables to be passed to all widgets
// at creation time.
type Properties struct {
	Font      *ttf.Font
	TitleFont *ttf.Font

	BgColor sdl.Color
	FgColor sdl.Color

	Zoom int // Zoom factor for the GameBoy display. Used for outlines, margins.
}

var properties *Properties

func SetProperties(p *Properties) {
	properties = p
}
