package gameboy

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"runtime/pprof"

	"github.com/lazy-stripes/goholint/apu"
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
	"github.com/veandco/go-sdl2/sdl"
)

// Package-wide logger.
var log = logger.New("gameboy", "interface-related logs")

// TickResult type to group return values from Tick.
type TickResult struct {
	Left, Right int8
	Play        bool
}

// GameBoy structure grouping all our state machines to tick them together.
type GameBoy struct {
	config *options.Options

	ticks   uint64
	APU     *apu.APU
	CPU     *cpu.CPU
	PPU     *ppu.PPU
	Display screen.Display // Interface, not pointer.
	DMA     *memory.DMA
	Serial  *serial.Serial
	Timer   *timer.Timer
	JPad    *joypad.Joypad

	Controls map[options.KeyStroke]Action

	// Send true to this channel to quit the program.
	QuitChan chan bool

	// Current palette.
	paletteIndex int

	// For GIF record toggle.
	recording bool
}

// SetControls validates and sets the given control map for the emulator.
func (g *GameBoy) SetControls(keymap options.Keymap) (err error) {
	// Intermediate mapping between labels and actual actions. This feels
	// unnecessarily complicated, but should make sense when I start translating
	// these from a config file. I hope.
	actions := map[string]Action{
		"up":              g.JoypadUp,
		"down":            g.JoypadDown,
		"left":            g.JoypadLeft,
		"right":           g.JoypadRight,
		"a":               g.JoypadA,
		"b":               g.JoypadB,
		"select":          g.JoypadSelect,
		"start":           g.JoypadStart,
		"screenshot":      g.Screenshot,
		"recordgif":       g.StartStopRecord,
		"nextpalette":     g.NextPalette,
		"previouspalette": g.PreviousPalette,
		"togglevoice1":    g.ToggleVoice1,
		"togglevoice2":    g.ToggleVoice2,
		"togglevoice3":    g.ToggleVoice3,
		"togglevoice4":    g.ToggleVoice4,
		"quit":            g.Quit,
	}

	g.Controls = make(map[options.KeyStroke]Action)
	for label, keyStroke := range keymap {
		g.Controls[keyStroke] = actions[label]
	}
	return nil
}

// New just instantiates most of the emulator. No biggie.
// TODO: try cleaning this mess up a little.
func New(config *options.Options) *GameBoy {
	g := GameBoy{
		config:   config,
		QuitChan: make(chan bool),
	}

	g.SetControls(config.Keymap)

	// Create CPU and interrupts first so other components can access them too.
	g.CPU = cpu.New(nil)
	ints := interrupts.New(&g.CPU.IF, &g.CPU.IE)

	g.APU = apu.New(config.Mono)

	g.Display = screen.NewSDL(config)
	if config.GIFPath != "" {
		//g.Display.Record(args.GIFPath)
		fmt.Printf("Saving GIF to %s\n", config.GIFPath)
	}

	// TODO: shouldn't we just pass Interrupts to New() functions?
	g.PPU = ppu.New(g.Display)
	g.PPU.Interrupts = ints

	g.Serial = serial.New()
	g.Serial.Interrupts = ints

	g.Timer = timer.New()
	g.Timer.Interrupts = ints

	var boot memory.Addressable
	if config.FastBoot {
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
		boot = memory.NewBoot(config.BootROM)
	}

	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff80, 0x7e)
	g.JPad = joypad.New() // TODO: interrupts
	mmu := memory.NewMMU([]memory.Addressable{
		boot,
		g.APU,
		g.PPU,
		wram,
		ints,
		g.JPad,
		g.Serial,
		g.Timer,
		hram,
	})

	// Memory space for the CPU, taking DMA transfers into account.
	mem := memory.NewDMAMemory(mmu, g.PPU.OAM.RAM)
	g.DMA = mem.DMA
	mmu.Add(g.DMA)
	g.CPU.Memory = mem

	if config.ROMPath != "" {
		// Build save path in case the cartridge uses one. Or use one
		// specified by the user.
		savePath := config.SavePath
		if savePath == "" {
			prefix := filepath.Dir(config.ROMPath)
			suffix := filepath.Base(config.ROMPath)
			savePath = prefix + "/" + suffix + ".sav"
		}
		// TODO: save-related error management.
		mmu.Add(memory.NewCartridge(config.ROMPath, savePath))
	}

	return &g
}

// Tick advances the whole emulator one step at a theoretical 4MHz. Since we're
// using SDL audio for timing this, we also return the current value of audio
// samples for each stereo channel as well as whether they should be played now.
func (g *GameBoy) Tick() (res TickResult) {
	g.ticks++

	// Poll events 1000 times per second.
	if g.ticks%4000 == 0 {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				eventType := event.GetType()
				switch eventType {

				// Button presses and UI keys
				case sdl.KEYDOWN, sdl.KEYUP:
					keyEvent := event.(*sdl.KeyboardEvent)
					keyStroke := options.KeyStroke{
						Code: keyEvent.Keysym.Sym,
						Mod:  sdl.Keymod(keyEvent.Keysym.Mod & options.ModMask),
					}
					if action := g.Controls[keyStroke]; action != nil {
						action(eventType)
					} else {
						if eventType == sdl.KEYDOWN {
							log.Infof("unknown key code: 0x%x", keyStroke.Code)
							log.Infof("        modifier: 0x%x", sdl.GetModState())
						}
					}

				// Window-closing event
				case sdl.QUIT:
					g.QuitChan <- true
				}
			}
		})
	}

	// DMA ticks occur every 4 machine ticks.
	if g.ticks%4 == 0 {
		g.DMA.Tick()
	}

	// CPU ticks occur every 4 machine ticks.
	if g.ticks%4 == 0 {
		g.CPU.Tick()
	}

	// PPU ticks occur every machine tick.
	g.PPU.Tick()

	// Timer tick occur every machine tick.
	g.Timer.Tick()

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
	// Make sure GIF file is written to disk.
	g.Display.Close()

	// If debugging at all, dump debug info.
	if len(g.config.DebugModules) > 0 {
		fmt.Println(g.CPU)
		fmt.Println(g.PPU)

		// Dump memory
		//g.CPU.DumpMemory()
	}
}

// Recover should be called at the end of each Tick. If the program panics, it
// should then display some useful debug info before crashing.
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
