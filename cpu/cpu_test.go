package cpu

import (
	"testing"

	"go.tigris.fr/gameboy/memory"
)

func TestCPU(t *testing.T) {
	rompath := "../bin/DMG_ROM.bin"
	rom := memory.NewROM(rompath, 0)
	//ram := memory.NewRAM(0x10000)
	mmu := memory.NewMMU([]memory.Addressable{rom})
	cpu := New(mmu)
	cpu.Run()
}
