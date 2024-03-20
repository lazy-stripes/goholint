package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Properties stores useful configuration variables to be passed to all widgets
// at creation time. Internally uses SDL types for convenience.
type Properties struct {
	Font      *ttf.Font
	TitleFont *ttf.Font // FIXME: only one Font property, to be set to small or large.

	BgColor sdl.Color
	FgColor sdl.Color

	Margin uint // TODO: all sides

	// For debugging if nothing else.
	Border      int32
	BorderColor sdl.Color

	Zoom int // Zoom factor for the GameBoy display. Used for outlines, margins.
}

var DefaultProperties Properties
