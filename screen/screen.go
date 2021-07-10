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

// Display supporting pixel output and palettes.
type Display interface {
	Enable()
	Enabled() bool
	Disable()
	Close()
	Write(colorIndex uint8)
	HBlank()
	VBlank()
	Blank()

	Message(text string, duration uint)
	Screenshot(filename string)
}

// Screen dimensions.
const (
	ScreenWidth  = 160
	ScreenHeight = 144
)

// Default palette colors with separate RGB components for easier use with SDL
// API. Kinda greenish.
const (
	ColorWhiteRGB     = 0xe0f0e7
	ColorWhiteR       = (ColorWhiteRGB >> 16) & 0xff
	ColorWhiteG       = (ColorWhiteRGB >> 8) & 0xff
	ColorWhiteB       = ColorWhiteRGB & 0xff
	ColorLightGrayRGB = 0x8ba394
	ColorLightGrayR   = (ColorLightGrayRGB >> 16) & 0xff
	ColorLightGrayG   = (ColorLightGrayRGB >> 16) & 0xff
	ColorLightGrayB   = ColorLightGrayRGB & 0xff
	ColorDarkGrayRGB  = 0x55645a
	ColorDarkGrayR    = (ColorDarkGrayRGB >> 16) & 0xff
	ColorDarkGrayG    = (ColorDarkGrayRGB >> 16) & 0xff
	ColorDarkGrayB    = ColorDarkGrayRGB & 0xff
	ColorBlackRGB     = 0x343d37
	ColorBlackR       = (ColorBlackRGB >> 16) & 0xff
	ColorBlackG       = (ColorBlackRGB >> 16) & 0xff
	ColorBlackB       = ColorBlackRGB & 0xff
)

var ColorWhite = color.RGBA{ColorWhiteR, ColorWhiteG, ColorWhiteB, 0xff}
var ColorLightGray = color.RGBA{ColorLightGrayR, ColorLightGrayG, ColorLightGrayB, 0xff}
var ColorDarkGray = color.RGBA{ColorDarkGrayR, ColorDarkGrayG, ColorDarkGrayB, 0xff}
var ColorBlack = color.RGBA{ColorBlackR, ColorBlackG, ColorBlackB, 0xff}

// DefaultPalette represents the selectable colors in the DMG.
var DefaultPalette = []color.Color{
	ColorWhite,
	ColorLightGray,
	ColorDarkGray,
	ColorBlack,
}
