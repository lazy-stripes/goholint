package widgets

// I want to try and simplify event handling for widgets. Not sure this counts.
type Event uint8

const (
	ButtonUp Event = iota
	ButtonDown
	ButtonLeft
	ButtonRight
	ButtonA
	ButtonB
	ButtonSelect
	ButtonStart

	// TODO: anything more complex for (limited) text/numbers input.
	// Or maybe just wrap SDL events?
)
