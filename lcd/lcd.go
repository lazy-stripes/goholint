package lcd

import (
	"fmt"
	"image/color"
)

// Pixel is an index into a display-defined palette.
type Pixel uint8

// Palette containing 4 indexed colors.
type Palette [4]color.NRGBA

// Display supporting pixel output and palettes.
type Display interface {
	Enable()
	Enabled() bool
	Disable()
	Close()
	Write(pixel Pixel)
	HBlank()
	VBlank()
	Blank()
}

// Screen dimensions.
const (
	ScreenWidth  = 160
	ScreenHeight = 144
)

// Default palette colors with separate RGB components for easier use with SDL API.
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
var DefaultPalette = [4]color.NRGBA{
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

func (c *Console) Enable() {
	c.enabled = true
}

func (c *Console) Enabled() bool {
	return c.enabled
}

func (c *Console) Disable() {
	c.enabled = false
}

func (c *Console) Write(pixel Pixel) {
	if c.enabled {
		fmt.Printf("%c", c.Palette[pixel])
	}
}

func (c *Console) HBlank() {
	if c.enabled {
		fmt.Print("\n")
	}
}

func (c *Console) VBlank() {
	if c.enabled {
		fmt.Print("\n === VBLANK ===\n")
		//fmt.Print("\033[2J")
	}
}

func (c *Console) Blank() {
	// Nothing. This method is only for refreshing SDL texture for "disabled" screen.
}
