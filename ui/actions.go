package ui

import (
	"fmt"
	"image"
	"image/png"
	"strings"

	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ppu/states"
	"github.com/lazy-stripes/goholint/ui/widgets"
	"golang.org/x/image/draw"

	"github.com/veandco/go-sdl2/sdl"
)

// Action type for user interactions. TODO: see if we can use a single one for gameboy/ui.
// TODO: see if we can't move that to widgets (see QAction for inspiration)
type Action func(state uint8)

// Quit cleanly quits the program.
func (u *UI) Quit(state uint8) {
	if state != sdl.PRESSED {
		return
	}
	u.QuitChan <- true
}

// Home hides the UI and resumes emulation.
func (u *UI) Home(state uint8) {
	if state != sdl.PRESSED {
		return
	}
	if u.paused {
		u.Hide()
	} else {
		// Wait for full frame before pausing emulator.
		u.screen.OnState(states.VBlank, u.Show)
	}
}

// NextPalette switches colors to the next defined palette, wrapping around.
// There should always be at least a default palette in the config object.
func (u *UI) NextPalette(state uint8) {
	if state != sdl.PRESSED {
		return
	}

	u.paletteIndex = (u.paletteIndex + 1) % len(u.config.Palettes)
	u.screen.Palette(u.config.Palettes[u.paletteIndex])
	u.screen.Message(u.config.PaletteNames[u.paletteIndex], 2)
}

// PreviousPalette switches colors to the previous defined palette, wrapping
// around. There should always be at least a default palette in the config
// object.
func (u *UI) PreviousPalette(state uint8) {
	if state != sdl.PRESSED {
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
func (u *UI) Screenshot(state uint8) {
	if state != sdl.PRESSED {
		return
	}

	f, err := options.CreateFileIn("screenshots", ".png")
	if err != nil {
		log.Warningf("creating screenshot file failed: %v", err)
		return
	}
	defer f.Close()

	// Grab current frame, resize and save.
	src := u.screen.Frame()
	dstRect := image.Rect(
		0,
		0,
		int(u.screenRect.W),
		int(u.screenRect.H),
	)
	dst := image.NewRGBA(dstRect)
	draw.NearestNeighbor.Scale(dst, dstRect, src, src.Bounds(), draw.Over, nil)

	if err := png.Encode(f, dst); err != nil {
		log.Warningf("saving screenshot failed: %v", err)
		return
	}

	u.screen.Message("Screenshot saved", 2)
	fmt.Printf("Screenshot saved to %s\n", f.Name())

}

// StartStopRecord starts recording video output to GIF and closes the file
// when done. Defined as a single action to toggle between the two and avoid
// opening several GIFs at once.
// TODO: move to ui
func (u *UI) StartStopRecord(state uint8) {
	if state != sdl.PRESSED {
		return
	}

	if u.recording { // TODO: query directly from screen?
		u.screen.StopRecord()
		u.recording = false
	} else {
		u.recording = true
		u.screen.StartRecord() // TODO: pass file descriptor there.
	}
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
func (u *UI) ToggleVoice1(state uint8) {
	if state != sdl.PRESSED {
		return
	}
	u.Emulator.ToggleVoice1()
	u.screen.Message(u.voiceStatusMsg(0), 2)
}

// ToggleVoice2 mutes or unmutes the second audio generator (Square 2).
func (u *UI) ToggleVoice2(state uint8) {
	if state != sdl.PRESSED {
		return
	}
	u.Emulator.ToggleVoice2()
	u.screen.Message(u.voiceStatusMsg(1), 2)
}

// ToggleVoice3 mutes or unmutes the third audio generator (Wave).
func (u *UI) ToggleVoice3(state uint8) {
	if state != sdl.PRESSED {
		return
	}
	u.Emulator.ToggleVoice3()
	u.screen.Message(u.voiceStatusMsg(2), 2)
}

// ToggleVoice4 mutes or unmutes the fourth audio generator (Noise).
func (u *UI) ToggleVoice4(state uint8) {
	if state != sdl.PRESSED {
		return
	}
	u.Emulator.ToggleVoice4()
	u.screen.Message(u.voiceStatusMsg(3), 2)
}

// TODO: so many things! Save states, toggle features...

// Actions available in Paused state.
func (u *UI) OpenROM() {
	openROMDialog := widgets.NewFileDialog(u.screenRect, "./bin/roms/")
	u.ShowDialog(openROMDialog, func(res widgets.DialogResult) {
		if res == widgets.DialogOK {
			path := openROMDialog.Selected().Value().(string)
			fmt.Printf("Opening ROM %s\n", path)

			// FIXME(?): GIF
			u.Emulator.Load(path)

		}
	})
}

// TODO: ... another dialogs.go, I guess.
func (u *UI) ShowDialog(dialog widgets.DialogWidget, cb widgets.DialogCloser) {
	dialog.OnClose(func(r widgets.DialogResult) {
		cb(r)
		u.dialogs.Pop() // XXX NavigateBack()?
		u.Hide()        // TODO: rename to Unpause()
	})
	u.dialogs.Push(dialog)

}
