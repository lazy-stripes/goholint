package cpu

import (
	"testing"

	"github.com/lazy-stripes/goholint/memory"
)

func TestCPU(t *testing.T) {
	rompath := "../bin/DMG_ROM.bin"
	rom := memory.NewROM(rompath, 0)
	//ram := memory.NewRAM(0x10000)
	mmu := memory.NewMMU([]memory.Addressable{rom})
	cpu := New(mmu)
	cpu.Tick()
}
