package gameboy

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Action type for user interactions. This might move to a ui package someday.
type Action func(eventType uint32)

// JoypadUp updates the Joypad's registers for the Up direction.
func (g *GameBoy) JoypadUp(eventType uint32) {
	g.JPad.Up.State = (eventType == sdl.KEYDOWN)
}

// JoypadDown updates the Joypad's registers for the Down direction.
func (g *GameBoy) JoypadDown(eventType uint32) {
	g.JPad.Down.State = (eventType == sdl.KEYDOWN)
}

// JoypadLeft updates the Joypad's registers for the Left direction.
func (g *GameBoy) JoypadLeft(eventType uint32) {
	g.JPad.Left.State = (eventType == sdl.KEYDOWN)
}

// JoypadRight updates the Joypad's registers for the Right direction.
func (g *GameBoy) JoypadRight(eventType uint32) {
	g.JPad.Right.State = (eventType == sdl.KEYDOWN)
}

// JoypadA updates the Joypad's registers for the A button.
func (g *GameBoy) JoypadA(eventType uint32) {
	g.JPad.A.State = (eventType == sdl.KEYDOWN)
}

// JoypadB updates the Joypad's registers for the B button.
func (g *GameBoy) JoypadB(eventType uint32) {
	g.JPad.B.State = (eventType == sdl.KEYDOWN)
}

// JoypadSelect updates the Joypad's registers for the Select button.
func (g *GameBoy) JoypadSelect(eventType uint32) {
	g.JPad.Select.State = (eventType == sdl.KEYDOWN)
}

// JoypadStart updates the Joypad's registers for the Start button.
func (g *GameBoy) JoypadStart(eventType uint32) {
	g.JPad.Start.State = (eventType == sdl.KEYDOWN)
}

// Screenshot saves the current frame to disk as a PNG file.
// TODO: configurable folder, obviously.
func (g *GameBoy) Screenshot(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	// Build a nice enough filename. TODO: configurable path.
	filename := fmt.Sprintf("goholint-%s-%d.png", time.Now().Format(DateFormat),
		g.CPU.Cycle)

	// Saving the current frame should really be up to the display (so it can
	// wait until VBlank for instance.)
	g.Display.Screenshot(filename)
}

// TODO: so many things! Toggle GIF record, save states, toggle features...
