package ppu

import (
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu/states"
)

// AddrOAM represents the base address of OAM RAM.
const AddrOAM = 0xfe00

// OAM represents an address space holding a list of sprites to display for the
// current scanline.
type OAM struct {
	Sprites []Sprite

	ram      memory.Addressable
	state    states.State
	ly, lcdc *uint8 // References to PPU's registers
	index    uint8
	sprite   Sprite // Current sprite
}

// Start OAM search.
func (o *OAM) Start() {
	o.Sprites = o.Sprites[:0] // Reset size only
	o.index = 0
	o.state = states.ReadSpriteY
}

// Tick advances OAM search one step and returns true when the search is over.
func (o *OAM) Tick() (done bool) {
	o.sprite.Address = AddrOAM + uint(o.index*4)
	switch o.state {
	case states.ReadSpriteY:
		o.sprite.Y = o.ram.Read(o.sprite.Address)
		o.state = states.ReadSpriteX
	case states.ReadSpriteX:
		// Sprite table must have room left, and current sprite tile must
		// contain current scanline.
		if len(o.Sprites) == 10 {
			break
		}

		o.index++
	}
	if o.index >= 40 {
		return true
	}
	return false
}
