package ppu

// Sprite type holds (x,y) coordinates of the current pixel of a sprite in the
// current scanline as well as its address in OAM RAM.
type Sprite struct {
	X, Y    uint8
	Address uint
}
