package gameboy

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

// Action type for user interactions. This might move to a ui package someday.
// FIXME: I want SDL out of Gameboy code, phase this out, use booleans directly.
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

// Helper strings to format UI messages.
var voiceNames = [4]string{
	"Square 1",
	"Square 2",
	"Wave",
	"Noise",
}

func (g *GameBoy) voiceStatusMsg(voice int) string {
	var sb strings.Builder
	for _, m := range g.APU.Muted {
		if m {
			sb.WriteRune('-')
		} else {
			sb.WriteRune('♪')
		}
	}
	sb.WriteRune('\n')
	sb.WriteString(voiceNames[voice])
	if g.APU.Muted[voice] {
		sb.WriteString(" muted")
	} else {
		sb.WriteString(" enabled")
	}
	return sb.String()
}

// ToggleVoice1 mutes or unmutes the first audio generator (Square 1).
func (g *GameBoy) ToggleVoice1(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.APU.Muted[0] = !g.APU.Muted[0]
	//g.UI.Message(g.voiceStatusMsg(0), 2)
}

// ToggleVoice2 mutes or unmutes the second audio generator (Square 2).
func (g *GameBoy) ToggleVoice2(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.APU.Muted[1] = !g.APU.Muted[1]
	//g.UI.Message(g.voiceStatusMsg(1), 2)
}

// ToggleVoice3 mutes or unmutes the third audio generator (Wave).
func (g *GameBoy) ToggleVoice3(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.APU.Muted[2] = !g.APU.Muted[2]
	//g.UI.Message(g.voiceStatusMsg(2), 2)
}

// ToggleVoice4 mutes or unmutes the fourth audio generator (Noise).
func (g *GameBoy) ToggleVoice4(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.APU.Muted[3] = !g.APU.Muted[3]
	// g.UI.Message(g.voiceStatusMsg(3), 2)
}
