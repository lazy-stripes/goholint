package ppu

import (
	"github.com/lazy-stripes/goholint/memory"
	"github.com/lazy-stripes/goholint/ppu/states"
)

// AddrOAM represents the base address of OAM RAM.
const AddrOAM = 0xfe00

// OAM computes a list of sprites to display for the current scanline.
type OAM struct {
	*memory.RAM

	Sprites []Sprite

	state    states.State
	ly, lcdc *uint8        // References to PPU's registers
	mode     *states.State // Display mode (equivalent to ppu.STAT&3)
	index    uint8
	sprite   Sprite // Current sprite
}

// NewOAM creates an OAM address space. Takes a reference to a PPU instance so
// it can access the LY and LCDC registers.
func NewOAM(ppu *PPU) *OAM {
	o := OAM{
		RAM:     memory.NewRAM(AddrOAM, 0xa0),
		Sprites: make([]Sprite, 0, 10),
		ly:      &ppu.LY,
		lcdc:    &ppu.LCDC,
		mode:    &ppu.state,
	}

	return &o
}

// Read overrides RAM method to restrict access to OAM to PPU modes 0 and 1.
// [https://gbdev.io/pandocs/Accessing_VRAM_and_OAM.html]
func (o *OAM) Read(addr uint16) uint8 {
	if *o.mode == states.HBlank || *o.mode == states.VBlank {
		return o.RAM.Read(addr)
	}
	return 0xff
}

// Write overrides RAM method to restrict access to OAM to PPU modes 0 and 1.
// DMA should bypass this method and directly call the underlying RAM object's
// Write method instead.
// [https://gbdev.io/pandocs/Accessing_VRAM_and_OAM.html]
func (o *OAM) Write(addr uint16, value uint8) {
	if *o.mode == states.HBlank || *o.mode == states.VBlank {
		o.RAM.Write(addr, value)
	}
	// Ignore write otherwise.
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
		o.sprite.Y = o.RAM.Read(o.sprite.Address)
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
		o.sprite.X = o.RAM.Read(o.sprite.Address + 1)
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
