package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"

	"github.com/veandco/go-sdl2/sdl"

	"go.tigris.fr/gameboy/cpu"
	"go.tigris.fr/gameboy/interrupts"
	"go.tigris.fr/gameboy/lcd"
	"go.tigris.fr/gameboy/logger"
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu"
	"go.tigris.fr/gameboy/serial"
	"go.tigris.fr/gameboy/timer"
)

func run(romPath string, fastBoot bool, waitKey bool) int {

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
		boot = memory.NewBoot("bin/boot/dmg_rom.bin")
	}

	cartridge := memory.NewCartridge(romPath)
	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff00, 0x100) // I/O ports, HRAM, IE FIXME: remove overlaps
	dma := &memory.DMA{}
	mmu := memory.NewMMU([]memory.Addressable{boot, ppu, wram, ints, serial, timer, dma, hram, cartridge})
	dma.MMU = mmu
	cpu.MMU = mmu

	// Add CPU-specific context to debug output.
	logger.Context = cpu.Context

	// Handle interrupt, store pointer to CPU for debug info.
	c := make(chan os.Signal, 1)
	go handleInterrupt(c, cpu)
	signal.Notify(c, os.Interrupt)

	// Wait for keypress if requested, so obs has time to capture window.
	if waitKey {
		fmt.Print("Press 'Enter' to start...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	// Main loop TODO: Gameboy.Run()
	tick := 0
	for {
		//t := time.Now()
		timer.Tick()
		cpu.Tick()
		dma.Tick()
		ppu.Tick()
		//fmt.Printf("Tick=%10d, cpu.PC=%02x   \r", tick, cpu.PC)
		tick++
		if tick == 229976-96 {
			//			fmt.Println("STOP")
		}
		//for time.Now().Sub(t) < time.Nanosecond*100 {
		//}
	}

	return 0
}

// User-defined type to parse a list of module names for which debug output must be enabled.
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

func handleInterrupt(c chan os.Signal, cpu *cpu.CPU) {
	// Wait for signal, quit cleanly with potential extra debug info if needed.
	<-c
	fmt.Println("\nTerminated...")

	// TODO: only dump RAM/VRAM/Other if requested in parameters.
	fmt.Print(cpu)
	cpu.DumpRAM()

	// Force stopping CPU profiling.
	pprof.StopCPUProfile()

	os.Exit(-1)
}

func main() {
	runtime.LockOSThread()

	var fastBoot = flag.Bool("fastboot", false, "bypass boot ROM execution")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var romPath = flag.String("rom", "", "ROM file to load")
	var waitKey = flag.Bool("waitkey", false, "Wait for keypress to start CPU (to help with screen captures)")
	var debugModules module
	flag.Var(&debugModules, "debug", "turn on debug mode for the given module")
	flag.Parse()

	for _, m := range debugModules {
		logger.Enabled[m] = true
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
	run(*romPath, *fastBoot, *waitKey)
}
