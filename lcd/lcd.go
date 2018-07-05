package lcd

import (
	"fmt"
)

// Pixel is an index into a display-defined palette.
type Pixel int

// Display supporting pixel output and palettes.
type Display interface {
	Write(pixel Pixel)
	HBlank()
	VBlank()
}

// Console display shifting piwels out to standard output.
type Console struct {
	Palette [4]rune
}

// NewConsole returns a Console display with dark-themed unicode pixels.
func NewConsole() *Console {
	return &Console{Palette: [4]rune{'█', '▒', '░', ' '}}
}

func (c Console) Write(pixel Pixel) {
	fmt.Print(c.Palette[pixel])
}

func (c Console) HBlank() {
	fmt.Print('\n')
}

func (c *Console) VBlank() {
	fmt.Print("\n === VBLANK ===\n")
}
