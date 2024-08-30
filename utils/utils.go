package utils

import "github.com/veandco/go-sdl2/sdl"

// Didn't know where else to put that stuff.

// WrapSDL returns the given function wrapped into an sdl.Do call, to be safely
// used as a callback function. This helps keeping SDL entirely out of gameboy
// code.
func WrapSDL(f func()) func() {
	return func() { sdl.Do(f) }
}
