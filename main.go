package main

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"go.tigris.fr/gameboy/cpu"
	"go.tigris.fr/gameboy/lcd"
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu"
)

func run() int {
	rompath := "bin/DMG_ROM.bin"
	bootRom := memory.NewROM(rompath, 0)
	if bootRom == nil {
		return 1
	}
	boot := memory.NewBoot(rompath)

	lcd := lcd.NewSDL()
	ppu := ppu.New(lcd)
	cartridge := memory.NewROM("bin/tetris.gb", 0)
	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff00, 0x100) // I/O ports, HRAM and IE register
	mmu := memory.NewMMU([]memory.Addressable{boot, ppu, wram, hram, cartridge})
	cpu := cpu.New(mmu)

	// Main loop TODO: Gameboy.Run()
	tick := 0
	for {
		cpu.Tick()
		ppu.Tick()
		//fmt.Printf("Tick=%10d, cpu.PC=%02x   \r", tick, cpu.PC)
		tick++
		if tick == 171704 {
			//fmt.Println("STOP")
		}
	}

	return 0
}

func main() {
	runtime.LockOSThread()
	sdl.Init(sdl.INIT_VIDEO)
	run()
}
