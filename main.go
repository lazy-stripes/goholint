package main

import (
	"tigris.fr/gameboy/cpu"
	"tigris.fr/gameboy/memory"
	"tigris.fr/gameboy/ppu"
)

func main() {
	rompath := "bin/DMG_ROM.bin"
	bootRom := memory.NewROM(rompath, 0)
	if bootRom == nil {
		return
	}
	ppu := ppu.New()
	ram := memory.NewRAM(0, 0x10000)
	mmu := memory.NewMMU([]memory.Addressable{bootRom, ppu, ram})
	cpu := cpu.New(mmu)
	cpu.Run()
}
