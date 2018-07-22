package lcd

import (
	"fmt"
)

// Pixel is an index into a display-defined palette.
type Pixel int

// Display supporting pixel output and palettes.
type Display interface {
	Enable()
	Disable()
	Close()
	Write(pixel Pixel)
	HBlank()
	VBlank()
}

// Console display shifting pixels out to standard output.
type Console struct {
	Palette [4]rune
	Enabled bool
}

// NewConsole returns a Console display with dark-themed unicode pixels.
func NewConsole() *Console {
	return &Console{Palette: [4]rune{'█', '▒', '░', ' '}}
}

func (c *Console) Enable() {
	c.Enabled = true
}

func (c *Console) Disable() {
	c.Enabled = false
}

func (c *Console) Write(pixel Pixel) {
	if c.Enabled {
		fmt.Printf("%c", c.Palette[pixel])
	}
}

func (c *Console) HBlank() {
	if c.Enabled {
		fmt.Print("\n")
	}
}

func (c *Console) VBlank() {
	if c.Enabled {
		fmt.Print("\n === VBLANK ===\n")
		//fmt.Print("\033[2J")
	}
}
