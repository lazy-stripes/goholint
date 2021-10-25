package ppu

import (
	"github.com/lazy-stripes/goholint/memory"
	"github.com/lazy-stripes/goholint/ppu/states"
)

// AddrOAM represents the base address of OAM RAM.
const AddrVRAM = 0x8000

// VRAM address space restricting access to Video RAM to modes 0, 1 and 2.
type VRAM struct {
	*memory.RAM
	mode *states.State // Display mode (equivalent to ppu.STAT&3)
}

// NewOAM creates an OAM address space. Takes a reference to a PPU instance so
// it can access the LY and LCDC registers.
func NewVRAM(ppu *PPU) *VRAM {
	v := VRAM{
		RAM:  memory.NewRAM(AddrVRAM, 0x2000),
		mode: &ppu.state,
	}

	return &v
}

// Read overrides RAM method to restrict access to OAM to PPU modes 0, 1 and 2.
// [https://gbdev.io/pandocs/Accessing_VRAM_and_OAM.html]
func (v *VRAM) Read(addr uint16) uint8 {
	if *v.mode != states.PixelTransfer {
		return v.RAM.Read(addr)
	}
	return 0xff
}

// Write overrides RAM method to restrict access to OAM to PPU modes 0, 1 and 2.
// [https://gbdev.io/pandocs/Accessing_VRAM_and_OAM.html]
func (v *VRAM) Write(addr uint16, value uint8) {
	if *v.mode != states.PixelTransfer {
		v.RAM.Write(addr, value)
	}
	// Ignore write otherwise.
}
