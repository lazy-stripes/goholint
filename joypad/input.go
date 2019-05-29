package joypad

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

// Input storing needed bits for a button or a direction.
type Input struct {
	Selector uint8 // P14 or P15
	Bit      uint8 // Bit position in JOYP register
	State    bool  // Button state (proper logic here: true is on, false is off)
}

// Joypad inputs whose state is to be updated by our main loop.
var InputRight = Input{P14, P10, false}
var InputLeft = Input{P14, P11, false}
var InputUp = Input{P14, P12, false}
var InputDown = Input{P14, P13, false}

var InputA = Input{P15, P10, false}
var InputB = Input{P15, P11, false}
var InputSelect = Input{P15, P12, false}
var InputStart = Input{P15, P13, false}

// Keymap maps an SDL key code (that can be looked up by name) to an Input.
// To be valid, a Keymap must map exactly eight keycodes to each possible
// Input.
type Keymap map[sdl.Keycode]*Input

// Pre-instanciated keymap errors.
var errEntriesNumber = errors.New("a keymap must have exactly eight keys")
var errUnknownInput = errors.New("a keymap can only to map known Inputs")
var errDuplicateKey = errors.New("a keymap needs keys to be distinct")
var errDuplicateInput = errors.New("a keymap needs inputs to be distinct")

// Validate returns an error if the keymap isn't valid, and nil otherwise.
// A valid keymap needs:
//  - Exactly eight keys
//  - Each key must map to one of the eight possible Inputs
//  - All mapped inputs must be distinct
//  - All mapped keys must be distinct
//
// TODO: Room for improvement here: several keys might map to one input.
func (k Keymap) Validate() error {
	// List of all needed inputs to check.
	allInputs := map[*Input]bool{
		&InputRight:  false,
		&InputLeft:   false,
		&InputUp:     false,
		&InputDown:   false,
		&InputA:      false,
		&InputB:      false,
		&InputSelect: false,
		&InputStart:  false,
	}

	if len(k) != len(allInputs) {
		return errEntriesNumber
	}

	allKeys := make(map[sdl.Keycode]bool)
	for key, input := range k {
		if allKeys[key] {
			return errDuplicateKey
		}
		present, known := allInputs[input]
		if !known {
			return errUnknownInput
		}
		if present {
			return errDuplicateInput
		}
		allKeys[key] = true
		allInputs[input] = true
	}
	return nil
}

// DefaultMapping should work okay out of the box for qwerty/azerty keyboards.
// It uses the arrow keys for directions, and the following button mappings:
// A:      S
// B:      D
// Select: Backspace
// Start:  Return
var DefaultMapping = Keymap{
	sdl.K_RIGHT:     &InputRight,
	sdl.K_LEFT:      &InputLeft,
	sdl.K_UP:        &InputUp,
	sdl.K_DOWN:      &InputDown,
	sdl.K_s:         &InputA,
	sdl.K_d:         &InputB,
	sdl.K_BACKSPACE: &InputSelect,
	sdl.K_RETURN:    &InputStart,
}
