package apu

// Length structure that will act as a state machine only managing the
// current length counter a Square or Noise signal generator.
type Length struct {
	// The properties below can be set by the APU itself.
	Counter uint8 // NRx1 bits 5-0

	enabled bool

	ticks uint // Clock ticks counter.
}

// Enable is called whenever the corresponding channel is triggered.
func (l *Length) Enable() {
	l.enabled = true
	l.ticks = 0
}

// Disable is called whenever the corresponding channel is triggered without the
// length bit.
func (l *Length) Disable() {
	l.enabled = false
}

// Tick advances the length one step. It will adjust the internal length value
// every 1/256 seconds (i.e. <cpuFreq>/256 ticks). Returns whether the signal
// generator should be disabled.
// Source: https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware
func (l *Length) Tick() (disable bool) {
	if !l.enabled {
		return
	}

	// If the Counter hasn't maxed out yet, advance it.
	// TODO: follow old GB wiki logic and load Counter with max value, then
	//       decrement, so we can use the same struct for Wave too.
	if l.Counter < 64 {
		// Advance length counter every <cpufreq>/256 (256Hz).
		stepRate := uint(GameBoyRate / 256)
		steps := (l.ticks + SoundOutRate) / stepRate
		l.ticks = (l.ticks + SoundOutRate) % stepRate

		l.Counter += uint8(steps)
	}

	// If the Counter has maxed out, tell the channel it should turn off.
	if l.Counter >= 64 {
		disable = true
		l.enabled = false
	}

	return
}
