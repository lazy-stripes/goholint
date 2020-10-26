package gameboy

import (
	"bufio"
	"fmt"
	"os"

	"github.com/faiface/mainthread"
	"github.com/lazy-stripes/goholint/apu"
	"github.com/lazy-stripes/goholint/cpu"
	"github.com/lazy-stripes/goholint/interrupts"
	"github.com/lazy-stripes/goholint/joypad"
	"github.com/lazy-stripes/goholint/lcd"
	"github.com/lazy-stripes/goholint/memory"
	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ppu"
	"github.com/lazy-stripes/goholint/serial"
	"github.com/lazy-stripes/goholint/timer"
	"github.com/veandco/go-sdl2/sdl"
)

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
	Display lcd.Display // Interface, not pointer.
	DMA     *memory.DMA
	Serial  *serial.Serial
	Timer   *timer.Timer
	JPad    *joypad.Joypad
}

// New just instantiates most of the emulator. No biggie.
func New(args *options.Options) *GameBoy {
	g := GameBoy{args: args}

	// Create CPU and interrupts first so other components can access them too.
	g.CPU = cpu.New(nil)
	ints := interrupts.New(&g.CPU.IF, &g.CPU.IE)

	g.APU = apu.New()

	if args.GIFPath != "" {
		fmt.Printf("Saving GIF to %s\n", args.GIFPath)
		g.Display = lcd.NewGIF(args.GIFPath, args.ZoomFactor, args.NoSync)
	} else {
		g.Display = lcd.NewSDL(args.ZoomFactor, args.NoSync)
	}
	g.PPU = ppu.New(g.Display)
	g.PPU.Interrupts = ints

	g.Serial = serial.New()
	g.Timer = timer.New()
	g.Timer.Interrupts = ints

	var boot memory.Addressable
	if args.FastBoot {
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
	g.JPad = joypad.New(joypad.DefaultMapping) // TODO: interrupts
	g.DMA = &memory.DMA{}
	mmu := memory.NewMMU([]memory.Addressable{boot, g.APU, g.APU.Wave.Pattern, g.PPU, wram, ints, g.JPad, g.Serial, g.Timer, g.DMA, hram})
	g.DMA.MMU = mmu
	g.CPU.MMU = mmu

	if args.ROMPath != "" {
		mmu.Add(memory.NewCartridge(args.ROMPath))
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
				switch event.GetType() {
				case sdl.KEYDOWN:
					keyEvent := event.(*sdl.KeyboardEvent)
					g.JPad.KeyDown(keyEvent.Keysym.Sym)
				case sdl.KEYUP:
					keyEvent := event.(*sdl.KeyboardEvent)
					g.JPad.KeyUp(keyEvent.Keysym.Sym)
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

// Run starts the emulator's main loop. This is used right now but we'll move to
// Tick() soon.
func (g *GameBoy) Run() {
	// Wait for keypress if requested, so obs has time to capture window.
	// Less useful now that we have -gif flag.
	if g.args.WaitKey {
		fmt.Print("Press 'Enter' to start...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	// Main loop.
	tick := 0
	quit := false
	for !quit {
		// FIXME: Ideally, we should plug into Blank/VBlank display methods.
		// Actually try polling for events at a little more than 60Hz.
		if g.PPU.Cycle%((456*153)/2) == 0 {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.GetType() {
				case sdl.KEYDOWN:
					keyEvent := event.(*sdl.KeyboardEvent)
					g.JPad.KeyDown(keyEvent.Keysym.Sym)
				case sdl.KEYUP:
					keyEvent := event.(*sdl.KeyboardEvent)
					g.JPad.KeyUp(keyEvent.Keysym.Sym)
				case sdl.QUIT:
					quit = true
				}
			}
		}

		g.CPU.Tick()
		g.DMA.Tick()
		g.PPU.Tick()
		g.Timer.Tick()

		tick++

		if g.args.Duration > 0 && g.CPU.Cycle >= g.args.Duration {
			break
		}
	}

	g.Display.Close()
}

// Stop should be called before quitting the program and will close all needed
// resources.
func (g *GameBoy) Stop() {
	// Make sure GIF file is written to disk.
	g.Display.Close()
}
