package widgets

// Simple enumeration of high-level events that can be used from the emulator's
// and the ui's package independently.

// I want to try and simplify event handling for widgets. Not sure this counts.
type Event uint8

const (
	// Button presses.
	ButtonUp Event = iota
	ButtonDown
	ButtonLeft
	ButtonRight
	ButtonA
	ButtonB
	ButtonSelect
	ButtonStart

	// Emulator state change.
	Paused
	Unpaused

	// TODO: anything more complex for (limited) text/numbers input.
	// Or maybe just wrap SDL events?
)
