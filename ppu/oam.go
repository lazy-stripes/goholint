package ppu

import (
	"github.com/lazy-stripes/goholint/memory"
	"github.com/lazy-stripes/goholint/ppu/states"
)

// AddrOAM represents the base address of OAM RAM.
const AddrOAM = 0xfe00

// OAM computes a list of sprites to display for the current scanline.
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
	o.sprite.Address = AddrOAM + uint16(o.index*4)
	switch o.state {
	case states.ReadSpriteY:
		o.sprite.Y = o.ram.Read(o.sprite.Address)
		o.state = states.ReadSpriteX
	case states.ReadSpriteX:
		// Sprite table must have room left, and current sprite tile must
		// contain current scanline.
		if len(o.Sprites) == 10 {
			return true
		}

		height := uint8(8)
		if *o.lcdc&LCDCSpriteSize != 0 {
			height = 16
		}

		// [TUGBT] 00:46:00
		//              0
		//              .
		//              |
		// [Sprite.Y]---+-------
		//              |
		// [LY]- - - - -+- - - -
		//              |
		//              |
		// [Sprite.Y+H]-+-------
		//              |
		//              .
		//             144
		o.sprite.X = o.ram.Read(o.sprite.Address + 1)
		if o.sprite.X != 0 {
			y := *o.ly + 16
			if o.sprite.Y <= y && o.sprite.Y+height > y {
				o.Sprites = append(o.Sprites, o.sprite)
			}
		}

		o.state = states.ReadSpriteY
		o.index++
	}

	return o.index >= 40
}
