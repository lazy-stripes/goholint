package ppu

// Pixel palettes used at display time. Will map to ppu.palettes.
const (
	PixelBGP  = 0
	PixelOBP0 = 1
	PixelOBP1 = 2
)

// Pixel holding its color index and palette to be used in our FIFO.
type Pixel struct {
	Color   uint8
	Palette uint8
}
