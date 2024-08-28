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

type PixelWriter interface {
	Write(colorIndex uint8)

}
