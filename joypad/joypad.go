package joypad

import (
	"github.com/lazy-stripes/goholint/logger"
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

// Input storing needed bits for a button or a direction.
type Input struct {
	Selector uint8 // P14 or P15
	Bit      uint8 // Bit position in JOYP register
	State    bool  // Button state (proper logic here: true is on, false is off)
}

// Joypad register and event manager for game inputs.
type Joypad struct {
	JOYP uint8

	Up     Input
	Down   Input
	Left   Input
	Right  Input
	A      Input
	B      Input
	Select Input
	Start  Input

	inputs []*Input // To iterate on all inputs at once
}

// New instantiates a Joypad addressable mapping to FF00 that will wait for
// events from the main loop.
func New() *Joypad {
	j := Joypad{}

	j.Right = Input{P14, P10, false}
	j.Left = Input{P14, P11, false}
	j.Up = Input{P14, P12, false}
	j.Down = Input{P14, P13, false}

	j.A = Input{P15, P10, false}
	j.B = Input{P15, P11, false}
	j.Select = Input{P15, P12, false}
	j.Start = Input{P15, P13, false}

	j.inputs = []*Input{
		&j.Up,
		&j.Down,
		&j.Left,
		&j.Right,
		&j.A,
		&j.B,
		&j.Select,
		&j.Start,
	}

	return &j
}

// Contains returns true if the requested address is the JOYP register.
func (j *Joypad) Contains(addr uint16) bool {
	return addr == AddrJOYP
}

// Read returns the state of selected inputs (inverted logic).
func (j *Joypad) Read(addr uint16) (value uint8) {
	selected := j.JOYP & 0x30
	// Set bits for inactive/unselected inputs, then NOT it all.
	for _, input := range j.inputs {
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

// KeyDown updates button states (if needed) when a key was pressed.
func (j *Joypad) KeyDown(input *Input) {
	input.State = true
}

// KeyUp updates button states (if needed) when a key was released.
func (j *Joypad) KeyUp(input *Input) {
	input.State = false
}
