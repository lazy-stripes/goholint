package lcd

// TODO: Rename this into 'screen' package. Picking this up after a year or so and I have trouble finding my way.

import (
	"fmt"
	"image/color"

	"github.com/lazy-stripes/goholint/logger"
)

// Package-wide logger.
var log = logger.New("lcd", "actual pixel display operations")

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
	ColorWhite      = 0xe0f0e7
	ColorWhiteR     = (ColorWhite >> 16) & 0xff
	ColorWhiteG     = (ColorWhite >> 8) & 0xff
	ColorWhiteB     = ColorWhite & 0xff
	ColorLightGray  = 0x8ba394
	ColorLightGrayR = (ColorLightGray >> 16) & 0xff
	ColorLightGrayG = (ColorLightGray >> 16) & 0xff
	ColorLightGrayB = ColorLightGray & 0xff
	ColorDarkGray   = 0x55645a
	ColorDarkGrayR  = (ColorDarkGray >> 16) & 0xff
	ColorDarkGrayG  = (ColorDarkGray >> 16) & 0xff
	ColorDarkGrayB  = ColorDarkGray & 0xff
	ColorBlack      = 0x343d37
	ColorBlackR     = (ColorBlack >> 16) & 0xff
	ColorBlackG     = (ColorBlack >> 16) & 0xff
	ColorBlackB     = ColorBlack & 0xff
)

// DefaultPalette represents the selectable colors in the DMG.
var DefaultPalette = []color.Color{
	color.NRGBA{ColorWhiteR, ColorWhiteG, ColorWhiteB, 0xff},
	color.NRGBA{ColorLightGrayR, ColorLightGrayG, ColorLightGrayB, 0xff},
	color.NRGBA{ColorDarkGrayR, ColorDarkGrayG, ColorDarkGrayB, 0xff},
	color.NRGBA{ColorBlackR, ColorBlackG, ColorBlackB, 0xff},
}

// Console display shifting pixels out to standard output.
type Console struct {
	Palette [4]rune
	enabled bool
}

// NewConsole returns a Console display with dark-themed unicode pixels.
func NewConsole() *Console {
	return &Console{Palette: [4]rune{'█', '▒', '░', ' '}}
}

// Enable turns on the display. "Pixels" will be printed out the moment they're written to the display.
func (c *Console) Enable() {
	c.enabled = true
}

// Enabled returns whether the display is enabled or not (as part of the Display interface).
func (c *Console) Enabled() bool {
	return c.enabled
}

// Disable turns off the display. No output will occur.
func (c *Console) Disable() {
	c.enabled = false
}

// Write prints out a pixel from our rune palette if display is enabled.
func (c *Console) Write(colorIndex uint8) {
	if c.enabled {
		fmt.Printf("%c", c.Palette[colorIndex])
	}
}

// HBlank prints a newline at HBlank time to set up the console for the next line.
func (c *Console) HBlank() {
	if c.enabled {
		fmt.Print("\n")
	}
}

// VBlank prints a separation between each console screen frame.
func (c *Console) VBlank() {
	if c.enabled {
		fmt.Print("\n === VBLANK ===\n")
		//fmt.Print("\033[2J")
	}
}

// Blank does nothing for a Console display.
func (c *Console) Blank() {
	// Nothing. This method is only for refreshing SDL texture for "disabled" screen.
}
