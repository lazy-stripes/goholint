package gameboy

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

// Action type for user interactions. This might move to a ui package someday.
// FIXME: I want SDL out of Gameboy code, phase this out, use booleans directly.
type Action func(state bool)

// JoypadUp updates the Joypad's registers for the Up direction.
func (g *GameBoy) JoypadUp(state bool) {
	g.JPad.Up.State = state
}

// JoypadDown updates the Joypad's registers for the Down direction.
func (g *GameBoy) JoypadDown(state bool) {
	g.JPad.Down.State = state
}

// JoypadLeft updates the Joypad's registers for the Left direction.
func (g *GameBoy) JoypadLeft(state bool) {
	g.JPad.Left.State = state
}

// JoypadRight updates the Joypad's registers for the Right direction.
func (g *GameBoy) JoypadRight(state bool) {
	g.JPad.Right.State = state
}

// JoypadA updates the Joypad's registers for the A button.
func (g *GameBoy) JoypadA(state bool) {
	g.JPad.A.State = state
}

// JoypadB updates the Joypad's registers for the B button.
func (g *GameBoy) JoypadB(state bool) {
	g.JPad.B.State = state
}

// JoypadSelect updates the Joypad's registers for the Select button.
func (g *GameBoy) JoypadSelect(state bool) {
	g.JPad.Select.State = state
}

// JoypadStart updates the Joypad's registers for the Start button.
func (g *GameBoy) JoypadStart(state bool) {
	g.JPad.Start.State = state
}

// Screenshot saves the current frame to disk as a PNG file.
// TODO: configurable folder, obviously.
func (g *GameBoy) Screenshot(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	//g.Display.Screenshot()
}

// StartStopRecord starts recording video output to GIF and closes the file
// when done. Defined as a single action to toggle between the two and avoid
// opening several GIFs at once.
func (g *GameBoy) StartStopRecord(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	if g.recording { // TODO: query directly from screen?
		//g.Display.StopRecord()
		g.recording = false
	} else {
		g.recording = true
		//g.Display.StartRecord()
	}
}

// ToggleVoice1 mutes or unmutes the first audio generator (Square 1).
func (g *GameBoy) ToggleVoice1() {
	g.APU.Muted[0] = !g.APU.Muted[0]
}

// ToggleVoice2 mutes or unmutes the second audio generator (Square 2).
func (g *GameBoy) ToggleVoice2() {
	g.APU.Muted[1] = !g.APU.Muted[1]
}

// ToggleVoice3 mutes or unmutes the third audio generator (Wave).
func (g *GameBoy) ToggleVoice3() {
	g.APU.Muted[2] = !g.APU.Muted[2]
}

// ToggleVoice4 mutes or unmutes the fourth audio generator (Noise).
func (g *GameBoy) ToggleVoice4() {
	g.APU.Muted[3] = !g.APU.Muted[3]
}
