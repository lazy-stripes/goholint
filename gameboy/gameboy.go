package gameboy

import (
	"fmt"
	"path/filepath"

	"github.com/faiface/mainthread"
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

// DateFormat layout for generated file names.
const DateFormat = "2006-01-02-15-04-05"

// TickResult type to group return values from Tick.
type TickResult struct {
	Left, Right uint8
	Play, Quit  bool
}

// GameBoy structure grouping all our state machines to tick them together.
type GameBoy struct {
	args *options.Options

	APU     *apu.APU
	CPU     *cpu.CPU
	PPU     *ppu.PPU
	Display screen.Display // Interface, not pointer.
	DMA     *memory.DMA
	Serial  *serial.Serial
	Timer   *timer.Timer
	JPad    *joypad.Joypad

	Controls map[sdl.Keycode]Action
}

// SetControls validates and sets the given control map for the emulator.
func (g *GameBoy) SetControls(keymap options.Keymap) (err error) {
	// Intermediate mapping between labels and actual actions. This feels
	// unnecessarily complicated, but should make sense when I start translating
	// these from a config file. I hope.
	actions := map[string]Action{
		"up":         g.JoypadUp,
		"down":       g.JoypadDown,
		"left":       g.JoypadLeft,
		"right":      g.JoypadRight,
		"a":          g.JoypadA,
		"b":          g.JoypadB,
		"select":     g.JoypadSelect,
		"start":      g.JoypadStart,
		"screenshot": g.Screenshot,
	}

	g.Controls = make(map[sdl.Keycode]Action)
	for label, keyCode := range keymap {
		g.Controls[keyCode] = actions[label]
	}
	return nil
}

// New just instantiates most of the emulator. No biggie.
func New(args *options.Options) *GameBoy {
	g := GameBoy{args: args}

	g.SetControls(args.Keymap)

	// Create CPU and interrupts first so other components can access them too.
	g.CPU = cpu.New(nil)
	ints := interrupts.New(&g.CPU.IF, &g.CPU.IE)

	g.APU = apu.New()

	if args.GIFPath != "" {
		fmt.Printf("Saving GIF to %s\n", args.GIFPath)
		g.Display = screen.NewGIF(args.GIFPath, args.ZoomFactor, args.VSync)
	} else {
		g.Display = screen.NewSDL(args.ZoomFactor, args.VSync)
	}
	g.PPU = ppu.New(g.Display)
	g.PPU.Interrupts = ints

	g.Serial = serial.New()
	g.Timer = timer.New()
	g.Timer.Interrupts = ints

	var boot memory.Addressable
	if args.FastBoot {
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
		boot = memory.NewBoot(args.BootROM)
	}

	wram := memory.NewRAM(0xc000, 0x2000)
	hram := memory.NewRAM(0xff80, 0x7e)
	g.JPad = joypad.New() // TODO: interrupts
	g.DMA = &memory.DMA{}
	mmu := memory.NewMMU([]memory.Addressable{
		boot,
		g.APU,
		g.APU.Wave.Pattern,
		g.PPU,
		wram,
		ints,
		g.JPad,
		g.Serial,
		g.Timer,
		g.DMA,
		hram,
	})
	g.DMA.MMU = mmu
	g.CPU.MMU = mmu

	if args.ROMPath != "" {
		// Build save path in case the cartridge uses one. Or use one
		// specified by the user.
		savePath := args.SavePath
		if savePath == "" {
			// The user could also just specify a path to a save folder.
			prefix := args.SaveDir
			if prefix == "" {
				prefix = filepath.Dir(args.ROMPath)
			}
			suffix := filepath.Base(args.ROMPath)
			savePath = prefix + "/" + suffix + ".sav"
		}
		// TODO: save-related error management.
		mmu.Add(memory.NewCartridge(args.ROMPath, savePath))
	}

	return &g
}

// Tick advances the whole emulator one step at a theoretical 4MHz. Since we're
// using SDL audio for timing this, we also return the current value of audio
// samples for each stereo channel as well as whether they should be played now.
func (g *GameBoy) Tick() (res TickResult) {
	// Check for external events (button presses, quit, etc.) first. We do that
	// based on VSync cycles, until I think of something better.
	if g.PPU.Cycle%(456*153) == 0 {
		mainthread.Call(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				eventType := event.GetType()
				switch eventType {

				// Button presses and UI keys
				case sdl.KEYDOWN, sdl.KEYUP:
					keyEvent := event.(*sdl.KeyboardEvent)
					keyCode := keyEvent.Keysym.Sym

					if action := g.Controls[keyCode]; action != nil {
						action(eventType)
					} else {
						log.Infof("unknown key code %v", keyCode)
					}

				// Window-closing event
				case sdl.QUIT:
					res.Quit = true
				}
			}
		})
	}

	g.CPU.Tick()
	g.DMA.Tick()
	g.PPU.Tick()
	g.Timer.Tick()
	res.Left, res.Right, res.Play = g.APU.Tick()
	return
}

// Stop should be called before quitting the program and will close all needed
// resources.
func (g *GameBoy) Stop() {
	// Make sure GIF file is written to disk.
	g.Display.Close()

	// If debugging at all, dump debug info.
	if len(g.args.DebugModules) > 0 {
		fmt.Println(g.CPU)
		fmt.Println(g.PPU)

		// Dump memory
		//g.CPU.DumpRAM()
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
		g.CPU.DumpRAM()
	}
}
