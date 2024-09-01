package screen

import (
	"image/color"

	"github.com/lazy-stripes/goholint/logger"
)

// Package-wide logger.
var log = logger.New("screen", "actual pixel display operations")

func init() {
	log.Add("gif", "GIF generator operations")
}

// ColorIndex into a display-defined 4-color palette.
type ColorIndex uint8

// Palette containing 4 indexed colors.
type Palette [4]color.NRGBA

type Display interface {
	// Enable or disable the display depending on LCDC.1 bit's value.
	Enable(bool)

	// Enabled returns the current status of the display.
	Enabled() bool

	// Write shifts out a pixel (a color index from 0 to 3 into the current
	// palette) to the display.
	Write(colorIndex uint8)

	// VBlank should be called when vertical blanking starts.
	VBlank()

	// OnVBlank takes a function that will get called once when VBlank starts.
	OnVBlank(func())
}
