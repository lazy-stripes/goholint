package widgets

import (
	"github.com/lazy-stripes/goholint/ui/widgets/align"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Properties stores useful configuration variables to be passed to all widgets
// at creation time. Internally uses SDL types for convenience.
// FIXME: separate actual widget properties (margin, background...) and configuration (Fg/BgColor, fonts...)
type Properties struct {
	Font      *ttf.Font
	TitleFont *ttf.Font // FIXME: only one Font property, to be set to small or large.

	BgColor sdl.Color // FIXME: OutlineColor
	FgColor sdl.Color

	Background sdl.Color // Background color (default is transparent)

	Margin  int32 // TODO: individual horizontal/vertical margin
	Padding int32 // TODO: individual horizontal/vertical padding

	HorizontalAlign align.Align // Widget contents alignment (horizontal)
	VerticalAlign   align.Align // Widget contents alignment (vertical)

	// For debugging if nothing else.
	Border      int32     // Border width in pixels
	BorderColor sdl.Color // Border color

	Zoom int // Zoom factor for the GameBoy display. Used for outlines, margins.
}

var DefaultProperties Properties
