package main

import (
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"tigris.fr/gameboy/cpu"
	"tigris.fr/gameboy/lcd"
	"tigris.fr/gameboy/memory"
	"tigris.fr/gameboy/ppu"
)

func run() int {
	rompath := "bin/DMG_ROM.bin"
	bootRom := memory.NewROM(rompath, 0)
	if bootRom == nil {
		return 1
	}
	lcd := lcd.NewSDL()
	ppu := ppu.New(lcd)
	cartridge := memory.NewROM("bin/tetris.gb", 0)
	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff00, 0x100) // I/O ports, HRAM and IE register
	mmu := memory.NewMMU([]memory.Addressable{bootRom, ppu, wram, hram, cartridge})
	cpu := cpu.New(mmu)

	// Main loop TODO: Gameboy.Run()
	tick := 0
	for {
		cpu.Tick()
		//ppu.Clock <- true // Tick(ppu.Clock, 4)
		//fmt.Printf("Tick=%10d, cpu.PC=%02x   \r", tick, cpu.PC)
		tick++
		if tick == 171704 {
			//fmt.Println("STOP")
		}
	}

	return 0
}

func main() {
	// os.Exit(..) must run AFTER sdl.Main(..) below; so keep track of exit
	// status manually outside the closure passed into sdl.Main(..) below
	var exitcode int
	runtime.LockOSThread()
	sdl.Init(sdl.INIT_VIDEO)
	sdl.Main(func() {
		exitcode = run()
	})
	// os.Exit(..) must run here! If run in sdl.Main(..) above, it will cause
	// premature quitting of sdl.Main(..) function; resource cleaning deferred
	// calls/closing of channels may never run
	os.Exit(exitcode)
}
