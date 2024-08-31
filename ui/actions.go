package ui

import (
	"strings"

	"github.com/lazy-stripes/goholint/utils"

	"github.com/veandco/go-sdl2/sdl"
)

// Action type for user interactions. TODO: see if we can use a single one for gameboy/ui.
type Action func(eventType uint32)

// Quit cleanly quits the program.
func (u *UI) Quit(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	u.QuitChan <- true
}

// Home hides the UI and resumes emulation.
func (u *UI) Home(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	if u.paused {
		u.Hide()
	} else {
		// Wait for full frame before pausing emulator.
		u.screen.OnVBlank(utils.WrapSDL(u.Show))
	}
}

// NextPalette switches colors to the next defined palette, wrapping around.
// There should always be at least a default palette in the config object.
func (u *UI) NextPalette(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	u.paletteIndex = (u.paletteIndex + 1) % len(u.config.Palettes)
	u.screen.Palette(u.config.Palettes[u.paletteIndex])
	u.screen.Message(u.config.PaletteNames[u.paletteIndex], 2)
}

// PreviousPalette switches colors to the previous defined palette, wrapping
// around. There should always be at least a default palette in the config
// object.
func (u *UI) PreviousPalette(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	u.paletteIndex -= 1
	if u.paletteIndex < 0 {
		// Wrap around (can't use % with negative values).
		u.paletteIndex = len(u.config.Palettes) - 1
	}
	u.screen.Palette(u.config.Palettes[u.paletteIndex])
	u.screen.Message(u.config.PaletteNames[u.paletteIndex], 2)
}

// Screenshot saves the current frame to disk as a PNG file.
// TODO: configurable folder, obviously.
// TODO: move to ui
func (u *UI) Screenshot(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	//g.Display.Screenshot()
}

// StartStopRecord starts recording video output to GIF and closes the file
// when done. Defined as a single action to toggle between the two and avoid
// opening several GIFs at once.
// TODO: move to ui
func (u *UI) StartStopRecord(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	//if u.recording { // TODO: query directly from screen?
	//	//g.Display.StopRecord()
	//	u.recording = false
	//} else {
	//	u.recording = true
	//	//g.Display.StartRecord()
	//}
}

// Helper strings to format UI messages.
var voiceNames = [4]string{
	"Square 1",
	"Square 2",
	"Wave",
	"Noise",
}

// TODO: have gb.ToggleVoice return current state, keep track of voice states here.
func (u *UI) voiceStatusMsg(voice int) string {
	var sb strings.Builder
	for _, m := range u.Emulator.APU.Muted {
		if m {
			sb.WriteRune('-')
		} else {
			sb.WriteRune('♪')
		}
	}
	sb.WriteRune('\n')
	sb.WriteString(voiceNames[voice])
	if u.Emulator.APU.Muted[voice] {
		sb.WriteString(" muted")
	} else {
		sb.WriteString(" enabled")
	}
	return sb.String()
}

// ToggleVoice1 mutes or unmutes the first audio generator (Square 1).
func (u *UI) ToggleVoice1(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	u.Emulator.ToggleVoice1()
	u.screen.Message(u.voiceStatusMsg(0), 2)
}

// ToggleVoice2 mutes or unmutes the second audio generator (Square 2).
func (u *UI) ToggleVoice2(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	u.Emulator.APU.Muted[1] = !u.Emulator.APU.Muted[1]
	u.screen.Message(u.voiceStatusMsg(1), 2)
}

// ToggleVoice3 mutes or unmutes the third audio generator (Wave).
func (u *UI) ToggleVoice3(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	u.Emulator.APU.Muted[2] = !u.Emulator.APU.Muted[2]
	u.screen.Message(u.voiceStatusMsg(2), 2)
}

// ToggleVoice4 mutes or unmutes the fourth audio generator (Noise).
func (u *UI) ToggleVoice4(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	u.Emulator.APU.Muted[3] = !u.Emulator.APU.Muted[3]
	u.screen.Message(u.voiceStatusMsg(3), 2)
}

// TODO: so many things! Save states, toggle features...
