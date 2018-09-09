package main

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"go.tigris.fr/gameboy/cpu"
	"go.tigris.fr/gameboy/interrupts"
	"go.tigris.fr/gameboy/lcd"
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu"
)

func run() int {
	rompath := "bin/DMG_ROM.bin"
	boot := memory.NewBoot(rompath)

	// Pre-instantiate CPU and interrupts so other components can access them too.
	cpu := cpu.New(nil)
	ints := interrupts.New(&cpu.IF, &cpu.IE)

	lcd := lcd.NewSDL()
	ppu := ppu.New(lcd)
	ppu.Interrupts = ints

	//cartridge := memory.NewROM("bin/tetris.gb", 0)
	cartridge := memory.NewROM("bin/cpu_instrs/cpu_instrs.gb", 0)
	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff00, 0x100) // I/O ports, HRAM, IE
	mmu := memory.NewMMU([]memory.Addressable{boot, ppu, wram, ints, hram, cartridge})
	cpu.MMU = mmu

	// Main loop TODO: Gameboy.Run()
	tick := 0
	for {
		cpu.Tick()
		ppu.Tick()
		//fmt.Printf("Tick=%10d, cpu.PC=%02x   \r", tick, cpu.PC)
		tick++
		if tick == 229976-96 {
			//			fmt.Println("STOP")
		}
		if cpu.PC == 0x98 {
			//			fmt.Println("STOP")
		}
	}

	return 0
}

func main() {
	runtime.LockOSThread()
	sdl.Init(sdl.INIT_VIDEO)
	run()
}
