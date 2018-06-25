package main

import (
	"tigris.fr/gameboy/cpu"
	"tigris.fr/gameboy/memory"
)

func main() {
	rompath := "bin/DMG_ROM.bin"
	rom := memory.NewBootROM(rompath)
	if rom == nil {
		return
	}
	ram := memory.NewRAM(0x10000)
	mmu := memory.NewMMU([]memory.AddressSpace{rom, ram})
	cpu := cpu.New(mmu)
	cpu.Run()
}
