package cpu

import (
	"testing"

	"tigris.fr/gameboy/memory"
)

func TestCPU(t *testing.T) {
	rompath := "../bin/DMG_ROM.bin"
	rom := memory.NewBootROM(rompath)
	//ram := memory.NewRAM(0x10000)
	mmu := memory.NewMMU([]memory.AddressSpace{rom})
	cpu := New(mmu)
	cpu.Run()
}
