package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Action type for user interactions. TODO: see if we can use a single one for gameboy/ui.
type Action func(eventType uint32)

// sdlWrap returns the given function wrapped into an sdl.Do call, to be safely
// used as a callback function. This helps keeping SDL entirely out of gameboy
// code.
func sdlWrap(f func()) func() {
	return func() { sdl.Do(f) }
}

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
		u.screen.OnVBlank(sdlWrap(u.Show))
	}
}

// TODO: so many things! Save states, toggle features...
