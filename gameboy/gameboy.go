package gameboy

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"runtime/pprof"

	"github.com/lazy-stripes/goholint/apu"
	"github.com/lazy-stripes/goholint/assets"
	"github.com/lazy-stripes/goholint/cpu"
	"github.com/lazy-stripes/goholint/interrupts"
	"github.com/lazy-stripes/goholint/joypad"
	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/memory"
	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ppu"
	"github.com/lazy-stripes/goholint/screen"
	"github.com/lazy-stripes/goholint/serial"
	"github.com/lazy-stripes/goholint/timer"
)

// Package-wide logger.
var log = logger.New("gameboy", "interface-related logs")

// VBlankRate represents ticks to wait before refreshing the Home screen.
const VBlankRate = apu.GameBoyRate / 60

// TickResult type to group return values from Tick.
type TickResult struct {
	Left, Right int8
	Play        bool // If true, use stereo sample in Left and Right.
	VBlank      bool // If true, repaint UI screen.
}

// GameBoy structure grouping all our state machines to tick them together.
type GameBoy struct {
	config *options.Options

	display screen.Display

	ticks  uint64
	APU    *apu.APU
	CPU    *cpu.CPU
	PPU    *ppu.PPU
	DMA    *memory.DMA
	Serial *serial.Serial
	Timer  *timer.Timer
	JPad   *joypad.Joypad

	// For GIF record toggle.
	recording bool
}

// New just instantiates most of the emulator. No biggie.
func New(display screen.Display, config *options.Options) *GameBoy {
	g := GameBoy{
		config:  config,
		display: display,
	}

	g.Reset()

	return &g
}

// Reset bootstraps the emulator given its currently stored config. Should be
// called when the emulator is paused (TODO or I must implement display.Reset()).
// TODO: try cleaning this mess up a little harder.
func (g *GameBoy) Reset() {
	// Create CPU and interrupts first so other components can access them too.
	g.CPU = cpu.New(nil)
	ints := interrupts.New(&g.CPU.IF, &g.CPU.IE)

	g.APU = apu.New(g.config.Mono)

	// TODO: move GIF handling to UI.
	if g.config.GIFPath != "" {
		//g.Display.Record(args.GIFPath)
		fmt.Printf("Saving GIF to %s\n", g.config.GIFPath)
	}

	// TODO: shouldn't we just pass Interrupts to New() functions?
	g.PPU = ppu.New(g.display)
	g.PPU.Interrupts = ints

	g.Serial = serial.New()
	g.Serial.Interrupts = ints

	g.Timer = timer.New()
	g.Timer.Interrupts = ints

	var boot memory.Addressable
	if g.config.FastBoot {
		// TODO: just implement save states, at this point.

		// XXX: What the BootROM does RAM-wise:
		// - Zero out/write logo tiles to 0x8000->0x9fff
		// - Write to audio registers
		// - Write to PPU registers
		// - Write to stack
		boot = memory.NewRAM(memory.BootAddr, 1)
		boot.Write(memory.BootAddr, 0x01)

		// Values below are what the CPU contains after booting the DMG ROM.
		g.CPU.A = 0x01
		g.CPU.F = 0xb0
		g.CPU.B = 0x00
		g.CPU.C = 0x13
		g.CPU.D = 0x00
		g.CPU.E = 0xd8
		g.CPU.H = 0x01
		g.CPU.L = 0x4d
		g.CPU.PC = 0x0100
		g.CPU.SP = 0xfffe

		// FIXME: properly pre-initialize PPU.
		//g.ppu.LCDC = 0x91
		//g.ppu.LY = 0x96
		//g.ppu.BGP = 0xfc

		for addr := 0x8000; addr <= 0x9fff; addr++ {
			// TODO: set RAM/VRAM
		}
	} else {
		if g.config.BootROM != "" {
			boot = memory.NewBootFromFile(g.config.BootROM)
		} else {
			boot = memory.NewBoot(assets.BootROM)
		}
	}

	hram := memory.NewRAM(0xff80, 0x7f)
	wram := memory.NewRAM(0xc000, 0x2000)
	//prohibited := memory.NewRAM(0xfea0, 0x60)
	g.JPad = joypad.New() // TODO: interrupts
	mmu := memory.NewMMU([]memory.Addressable{
		g.APU,
		hram,
		boot,
		g.PPU,
		//prohibited,
		wram,
		ints,
		g.JPad,
		g.Serial,
		g.Timer,
	})

	// Memory space for the CPU, taking DMA transfers into account.
	mem := memory.NewDMAMemory(mmu, g.PPU.OAM.RAM)
	g.DMA = mem.DMA
	mmu.Add(g.DMA)
	g.CPU.Memory = mem

	if g.config.ROMPath != "" {
		// Build save path in case the cartridge uses one. Or use one
		// specified by the user.
		savePath := g.config.SavePath
		if savePath == "" {
			prefix := filepath.Dir(g.config.ROMPath)
			suffix := filepath.Base(g.config.ROMPath)
			savePath = prefix + "/" + suffix + ".sav"
		}
		// TODO: save-related error management.
		mmu.Add(memory.NewCartridge(g.config.ROMPath, savePath))
	}
}

func (g *GameBoy) Load(path string) {
	// TODO: error handling, will need more work on UI side.
	g.config.ROMPath = path
	g.Reset()
}

// Tick advances the whole emulator one step at a theoretical 4MHz. Since we're
// using SDL audio for timing this, we also return the current value of audio
// samples for each stereo channel as well as whether they should be played now.
func (g *GameBoy) Tick() (res TickResult) {
	g.ticks++

	// PPU ticks occur every machine tick.
	res.VBlank = g.PPU.Tick()

	// Timer tick occur every machine tick.
	g.Timer.Tick()

	// DMA ticks occur every 4 machine ticks.
	if g.ticks%4 == 0 {
		g.DMA.Tick()
	}

	// CPU ticks occur every 4 machine ticks.
	if g.ticks%4 == 0 {
		g.CPU.Tick()
	}

	// Serial ticks happen at 8KHz.
	if g.ticks%serial.BitOutRate == 0 {
		g.Serial.Tick()
	}

	// APU ticks occur only when we need to generate the next sample.
	// Note that the Gameboy machine frequency is not an exact multiple of the
	// sound output frequency, so this is in fact an approximation. So long as
	// no one can hear the difference, let's call it good enough.
	if g.ticks%apu.SoundOutRate == 0 {
		res.Left, res.Right = g.APU.Tick()
		res.Play = true
	}

	return
}

// Stop should be called before quitting the program and will close all needed
// resources.
func (g *GameBoy) Stop() {
	// FIXME: Make sure GIF file is written to disk. That sounds like a UI-problem.
	// g.Display.Close()

	// If debugging at all, dump debug info.
	if g.config.DebugModules != nil {
		fmt.Println(g.CPU)
		fmt.Println(g.PPU)

		// Dump memory TODO: make that configurable, -dumpmem or something.
		g.CPU.DumpMemory()
	}
}

// Recover should be called at the end of each Tick. If the program panics, it
// should then display some useful debug info before crashing.
// TODO: there is some overlap with handleSIGINT, maybe part or all of this could be done from UI too.
func (g *GameBoy) Recover() {
	if r := recover(); r != nil {
		fmt.Printf("Goholint seems to have crashed (%v). I'm sorry.\n\n", r)

		fmt.Println(g.CPU)
		fmt.Println(g.PPU)

		// Dump memory
		g.CPU.DumpMemory()

		// Force-close CPU-profile if any.
		pprof.StopCPUProfile()

		// Still display stack trace for debugging.
		fmt.Print("\nThe information below may help fix the problem:\n\n")
		debug.PrintStack()

		// Bail. Bummer.
		os.Exit(1)
	}
}
