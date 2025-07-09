package apu

// Length structure that will act as a state machine only managing the
// current length counter a Square or Noise signal generator.
type Length struct {
	// The properties below can be set by the APU itself.
	Initial uint8 // NRx1 bits 5-0 (or 8-0 for Wave generator)

	timer uint16 // Current internal timer value.
	ticks uint   // DIV-APU ticks counter.
}

// Reset is called whenever the corresponding channel is triggered. It takes the
// maximum timer value (256 for Wave generator, 64 for the others) as parameter
// and will set the internal timer to (max-Initial) if it's currently zero.
func (l *Length) Reset(max uint16) {
	l.ticks = 0
	if l.timer == 0 {
		l.timer = max - uint16(l.Initial)
	}
}

// Tick advances the length one step. It will adjust the internal length value
// every 1/256 seconds (i.e. every 2 DIV-APU ticks). Returns whether the signal
// generator should be disabled.
//
// Tick won't be called if the generator itself is not enabled.
func (l *Length) Tick() (disable bool) {
	l.ticks++
	if l.ticks < 2 {
		return
	}
	l.ticks = 0

	// Decrement internal timer until zero.
	if l.timer > 0 {
		l.timer--

		// If the timer has expired, generator must be turned off.
		if l.timer == 0 {
			disable = true
		}
	}

	return
}
