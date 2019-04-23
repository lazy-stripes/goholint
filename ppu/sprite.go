package ppu

// Sprite flags
const (
	SpritePalette = 1 << (iota + 4)
	SpriteFlipX
	SpriteFlipY
	SpritePriority
)

// Sprite type holds (x,y) coordinates of the current pixel of a sprite in the
// current scanline as well as its address in OAM RAM.
type Sprite struct {
	X, Y    uint8
	Address uint16
	Fetched bool // Set to true after this sprite was treated for a given line.
}
