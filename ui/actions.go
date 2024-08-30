package ui

import (
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

// TODO: so many things! Save states, toggle features...
