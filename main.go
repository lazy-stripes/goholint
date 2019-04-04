package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/veandco/go-sdl2/sdl"

	"go.tigris.fr/gameboy/cpu"
	"go.tigris.fr/gameboy/debug"
	"go.tigris.fr/gameboy/interrupts"
	"go.tigris.fr/gameboy/lcd"
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu"
	"go.tigris.fr/gameboy/serial"
	"go.tigris.fr/gameboy/timer"
)

func run(romPath string, fastBoot bool) int {

	// Pre-instantiate CPU and interrupts so other components can access them too.
	cpu := cpu.New(nil)
	ints := interrupts.New(&cpu.IF, &cpu.IE)

	lcd := lcd.NewSDL()
	ppu := ppu.New(lcd)
	ppu.Interrupts = ints

	serial := serial.New()
	timer := timer.New()
	timer.Interrupts = ints

	var boot memory.Addressable
	if fastBoot {
		// TODO: set vram and other registers
		boot = memory.NewRAM(memory.BootAddr, 1)
		cpu.PC = 0x0100
	} else {
		rompath := "bin/DMG_ROM.bin"
		boot = memory.NewBoot(rompath)
	}

	var cartridge memory.Addressable
	if romPath == "" {
		cartridge = memory.NewRAM(0, 0)
	} else {
		cartridge = memory.NewROM(romPath, 0)
	}
	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff00, 0x100) // I/O ports, HRAM, IE FIXME: remove overlaps
	mmu := memory.NewMMU([]memory.Addressable{boot, ppu, wram, ints, serial, timer, hram, cartridge})
	cpu.MMU = mmu

	// Main loop TODO: Gameboy.Run()
	tick := 0
	for {
		timer.Tick()
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

// User-defined type to parse a list of module names for which debug output muist be enabled.
type module []string

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (m *module) String() string {
	return fmt.Sprint(*m)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag. Flag can be specified multiple times.
func (m *module) Set(value string) error {
	*m = append(*m, value)
	return nil
}

func main() {
	runtime.LockOSThread()

	var fastBoot = flag.Bool("fastboot", false, "bypass boot ROM execution")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var romPath = flag.String("rom", "", "ROM file to load")
	var debugModules module
	flag.Var(&debugModules, "debug", "turn on debug mode for the given module")
	flag.Parse()

	for _, m := range debugModules {
		debug.Enabled[m] = true
	}

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
	run(*romPath, *fastBoot)
}
