package screen

import (
	"image/color"
	"time"

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

// Display supporting pixel output and palettes.
type Display interface {
	Enable()
	Enabled() bool
	Disable()
	Close()
	Write(colorIndex uint8)
	HBlank()
	VBlank()

	Text(text string)
	Message(text string, duration time.Duration)

	Screenshot(filename string)

	Record(filename string)
	StopRecord()
}

// Screen dimensions.
const (
	ScreenWidth  = 160
	ScreenHeight = 144
)
