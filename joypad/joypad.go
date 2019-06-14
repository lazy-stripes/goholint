package joypad

import (
	"github.com/veandco/go-sdl2/sdl"
	"go.tigris.fr/gameboy/logger"
)

// Source: [JOYPAD] http://gbdev.gg8.se/wiki/articles/Joypad_Input

// Package-wide logger.
var log = logger.New("joypad", "user inputs manager")

// Package initialization function setting up logger.
func init() {
	log.Add("read", "register reads (Desperate level only)")
	log.Add("input", "input changes (Debug level or lower)")
}

// Register address.
const (
	AddrJOYP = 0xff00
)

// Bit values to select/update button states.
const (
	P10 = 1 << iota // Bit 0 - Input Right or Button A (0=Pressed) (Read Only)
	P11             // Bit 1 - Input Left  or Button B (0=Pressed) (Read Only)
	P12             // Bit 2 - Input Up    or Select   (0=Pressed) (Read Only)
	P13             // Bit 3 - Input Down  or Start    (0=Pressed) (Read Only)
	P14             // Bit 4 - Select Direction Keys   (0=Select)
	P15             // Bit 5 - Select Button Keys      (0=Select)
)

// Joypad register and event manager for game inputs.
type Joypad struct {
	JOYP   uint8
	Keymap Keymap
}

// New instantiates a Joypad addressable mapping to FF00 that will wait for
// events from the main loop.
func New(keymap Keymap) *Joypad {
	if err := keymap.Validate(); err != nil {
		log.Warningf("Invalid keymap %v. Using default instead.", keymap)
		keymap = DefaultMapping
	}
	return &Joypad{Keymap: keymap}
}

// Contains returns true if the requested address is the JOYP register.
func (j *Joypad) Contains(addr uint16) bool {
	return addr == AddrJOYP
}

// Read returns the state of selected inputs (inverted logic).
func (j *Joypad) Read(addr uint16) (value uint8) {
	selected := j.JOYP & 0x30
	// Set bits for inactive/unselected inputs, then NOT it all.
	for _, input := range j.Keymap {
		if selected&input.Selector == 0 && input.State {
			value |= input.Bit
		}
	}
	value = (^value)&0x0f | selected
	log.Sub("read").Desperatef("JOYP=0x%02x", value)
	return value
}

// Write updates the writeable bits of the JOYP register.
func (j *Joypad) Write(addr uint16, value uint8) {
	j.JOYP = value & 0x30
}

// Helper method setting or resetting an input's state.
func (j *Joypad) setInput(code sdl.Keycode, state bool) {
	if input := j.Keymap[code]; input != nil {
		log.Sub("input").Debugf("%v=%t", input, state)
		input.State = state
	}
}

// KeyDown updates button states (if needed) when a key was pressed.
func (j *Joypad) KeyDown(code sdl.Keycode) {
	j.setInput(code, true)
}

// KeyUp updates button states (if needed) when a key was released.
func (j *Joypad) KeyUp(code sdl.Keycode) {
	j.setInput(code, false)
}
