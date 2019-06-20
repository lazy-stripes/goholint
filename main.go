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
	"go.tigris.fr/gameboy/joypad"
	"go.tigris.fr/gameboy/lcd"
	"go.tigris.fr/gameboy/logger"
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu"
	"go.tigris.fr/gameboy/serial"
	"go.tigris.fr/gameboy/timer"
)

// TODO: minimal (like, REALLY minimal) GUI. And clean all of this up.
func run(options *Options) int {

	// Pre-instantiate CPU and interrupts so other components can access them too.
	cpu := cpu.New(nil)
	ints := interrupts.New(&cpu.IF, &cpu.IE)

	var display lcd.Display
	if options.GIFPath != "" {
		display = lcd.NewGIF(options.GIFPath, options.ZoomFactor)
	} else {
		display = lcd.NewSDL(options.ZoomFactor)
	}
	ppu := ppu.New(display)
	ppu.Interrupts = ints

	serial := serial.New()
	timer := timer.New()
	timer.Interrupts = ints

	var boot memory.Addressable
	if options.FastBoot {
		// TODO: set RAM
		boot = memory.NewRAM(memory.BootAddr, 1)
		boot.Write(memory.BootAddr, 0x01)

		// Values below are what the CPU contains when we boot the DMG ROM.
		cpu.A = 0x01
		cpu.F = 0xb0
		cpu.B = 0x00
		cpu.C = 0x13
		cpu.D = 0x00
		cpu.E = 0xd8
		cpu.H = 0x01
		cpu.L = 0x4d
		cpu.PC = 0x0100
	} else {
		boot = memory.NewBoot("bin/boot/dmg_rom.bin")
	}

	cartridge := memory.NewCartridge(options.ROMPath)
	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff80, 0x7e)
	jpad := joypad.New(joypad.DefaultMapping) // TODO: interrupts
	dma := &memory.DMA{}
	mmu := memory.NewMMU([]memory.Addressable{boot, ppu, wram, ints, jpad, serial, timer, dma, hram, cartridge})
	dma.MMU = mmu
	cpu.MMU = mmu

	// Add CPU-specific context to debug output.
	logger.Context = cpu.Context

	// Handle interrupt, store pointer to CPU for debug info.
	c := make(chan os.Signal, 1)
	go handleInterrupt(c, cpu, display)
	signal.Notify(c, os.Interrupt)

	// Wait for keypress if requested, so obs has time to capture window.
	// Less useful now we have -gif flag.
	if options.WaitKey {
		fmt.Print("Press 'Enter' to start...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	// Main loop TODO: Gameboy.Run()
	tick := 0
	quit := false
	for !quit {
		//t := time.Now()
		// FIXME: Ideally, we should plug into Blank/VBlank display methods.
		if ppu.Cycle%(456*153) == 0 {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.GetType() {
				case sdl.KEYDOWN:
					keyEvent := event.(*sdl.KeyboardEvent)
					jpad.KeyDown(keyEvent.Keysym.Sym)
				case sdl.KEYUP:
					keyEvent := event.(*sdl.KeyboardEvent)
					jpad.KeyUp(keyEvent.Keysym.Sym)
				case sdl.QUIT:
					quit = true
				}
			}
		}

		cpu.Tick()
		dma.Tick()
		ppu.Tick()
		timer.Tick()
		//fmt.Printf("Tick=%10d, cpu.PC=%02x   \r", tick, cpu.PC)
		tick++
		//if tick == 229976-96 {
		//			fmt.Println("STOP")
		//}

		if options.Duration > 0 && cpu.Cycle >= options.Duration {
			break
		}
	}

	display.Close()
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

func handleInterrupt(c chan os.Signal, cpu *cpu.CPU, lcd lcd.Display) {
	// Wait for signal, quit cleanly with potential extra debug info if needed.
	<-c
	fmt.Println("\nTerminated...")

	lcd.Close()

	// TODO: only dump RAM/VRAM/Other if requested in parameters.
	fmt.Print(cpu)
	cpu.DumpRAM()

	// Force stopping CPU profiling.
	pprof.StopCPUProfile()

	os.Exit(-1)
}

// Options structure grouping command line flags values.
type Options struct {
	FastBoot   bool
	ROMPath    string
	GIFPath    string
	WaitKey    bool
	ZoomFactor uint8
	Duration   uint
}

func main() {
	runtime.LockOSThread()

	var fastBoot = flag.Bool("fastboot", false, "Bypass boot ROM execution")
	var cpuprofile = flag.String("cpuprofile", "", "Write cpu profile to file")
	var duration = flag.Uint("cycles", 0, "Stop after executing that many cycles")
	var debugModules module
	flag.Var(&debugModules, "debug", "Turn on debug mode for the given module (-debug help for the full list)")
	var gifPath = flag.String("gif", "", "Record gif file")
	var romPath = flag.String("rom", "", "ROM file to load")
	var waitKey = flag.Bool("waitkey", false, "Wait for keypress to start CPU (to help with screen captures)")
	var zoomFactor = flag.Int("zoom", 2, "Zoom factor (default is 2x)")
	flag.Parse()

	for _, m := range debugModules {
		// List available modules if requested.
		if m == "help" {
			logger.Help()
			os.Exit(0)
		}
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

	opt := Options{FastBoot: *fastBoot,
		ROMPath:    *romPath,
		GIFPath:    *gifPath,
		WaitKey:    *waitKey,
		ZoomFactor: uint8(*zoomFactor),
		Duration:   *duration,
	}

	run(&opt)
}
