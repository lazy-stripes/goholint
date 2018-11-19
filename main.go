package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/veandco/go-sdl2/sdl"
	"go.tigris.fr/gameboy/cpu"
	"go.tigris.fr/gameboy/interrupts"
	"go.tigris.fr/gameboy/lcd"
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu"
	"go.tigris.fr/gameboy/serial"
)

func run(fastBoot bool) int {

	// Pre-instantiate CPU and interrupts so other components can access them too.
	cpu := cpu.New(nil)
	ints := interrupts.New(&cpu.IF, &cpu.IE)

	lcd := lcd.NewSDL()
	ppu := ppu.New(lcd)
	ppu.Interrupts = ints

	serial := serial.New()

	var boot memory.Addressable
	if fastBoot {
		// TODO: set vram and other registers
		boot = memory.NewRAM(memory.BootAddr, 1)
		cpu.PC = 0x0100
	} else {
		rompath := "bin/DMG_ROM.bin"
		boot = memory.NewBoot(rompath)
	}

	//cartridge := memory.NewROM("bin/tetris.gb", 0)
	//cartridge := memory.NewROM("bin/sml.gb", 0)
	//cartridge := memory.NewROM("bin/cpu_instrs/individual/03-op sp,hl.gb", 0)
	cartridge := memory.NewROM("bin/cpu_instrs/individual/04-op r,imm.gb", 0)
	//cartridge := memory.NewROM("bin/cpu_instrs/individual/05-op rp.gb", 0)
	//cartridge := memory.NewROM("bin/cpu_instrs/individual/06-ld r,r.gb", 0)
	//cartridge := memory.NewROM("bin/cpu_instrs/individual/07-jr,jp,call,ret,rst.gb", 0)
	//cartridge := memory.NewROM("bin/cpu_instrs/individual/08-misc instrs.gb", 0)
	//cartridge := memory.NewROM("bin/cpu_instrs/individual/09-op r,r.gb", 0)
	//cartridge := memory.NewROM("bin/cpu_instrs/individual/10-bit ops.gb", 0)
	//cartridge := memory.NewROM("bin/cpu_instrs/individual/11-op a,(hl).gb", 0)
	//cartridge := memory.NewRAM(0, 0)
	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff00, 0x100) // I/O ports, HRAM, IE FIXME: remove overlaps
	mmu := memory.NewMMU([]memory.Addressable{boot, ppu, wram, ints, serial, hram, cartridge})
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
	}

	return 0
}

func main() {
	runtime.LockOSThread()

	var fastBoot = flag.Bool("fastboot", false, "bypass boot ROM execution")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		if err = pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()

		log.Println("CPU profiling written to: ", *cpuprofile)
	}
	sdl.Init(sdl.INIT_VIDEO)
	run(*fastBoot)
}
