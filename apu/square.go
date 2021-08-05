package apu

// DutyCycles represents available duty patterns. For any given frequency,
// we'll internally split one period of that frequency in 8, and for each
// of those slices, this will specify whether the signal should be on or off.
var DutyCycles = [4][8]bool{
	{false, false, false, false, false, false, false, true}, // 00000001, 12.5%
	{true, false, false, false, false, false, false, true},  // 10000001, 25%
	{true, false, false, false, false, true, true, true},    // 10000111, 50%
	{false, true, true, true, true, true, true, false},      // 01111110, 75%
}

// SquareWave structure implementing sound sample generation for two of the four
// possible sounds the Game Boy can produce at once. A.k.a Sound1 and Sound2.
type SquareWave struct {
	NRx0 uint8 // Sweep pattern (if applicable)
	NRx1 uint8 // Pattern duty and sound length
	NRx2 uint8 // Volume envelope
	NRx3 uint8 // Frequency's lower 8 bits
	NRx4 uint8 // Control and frequency' higher 3 bits

	enabled bool // Only output silence if this is false

	// Duty-related variables.
	dutyStep int  // Sub-index into DutyCycles to set the signal high or low.
	ticks    uint // Clock ticks counter for advancing duty step.

	envelope VolumeEnvelope
}

// SetNRx2 is called whenever the NRx2 register's value was changed, so that it
// can update the volume envelope state machine.
func (s *SquareWave) SetNRx2(value uint8) {
	s.envelope.Initial = value >> 4
	s.envelope.Sweep = value & 7
	if value&NRx2EnvelopeDirection != 0 {
		s.envelope.Direction = 1
	} else {
		s.envelope.Direction = -1
	}
}

// Tick produces a sample of the signal to generate based on the current value
// in the signal generator's registers. We use a named return value, which is
// conveniently set to zero (silence) by default.
func (s *SquareWave) Tick() (sample uint8) {
	// Enable that signal if requested. NR14 being write-only, we can reset it
	// each time it goes to 1 without worrying.
	if s.NRx4&NRx4RestartSound != 0 {
		s.NRx4 &= ^NRx4RestartSound // Reset trigger bit
		log.Debug("NR14 triggered")
		s.enabled = true // It's fine if the signal is already enabled.

		// "Restarting a pulse channel causes its "duty step timer" to reset."
		// Source: https://gbdev.gg8.se/wiki/articles/Sound_Controller#PitFalls
		s.ticks = 0

		s.envelope.Enable()
	}

	if !s.enabled {
		return
	}

	s.envelope.Tick()

	// With `x` the 11-bit value in NR13/NR14, frequency is 131072/(2048-x) Hz.
	rawFreq := ((uint(s.NRx4) & 7) << 8) | uint(s.NRx3)
	freq := 131072 / (2048 - rawFreq)

	// Advance duty step every 1/(8f) where f is the sound's real frequency
	// for as many machine ticks as necessary to generate one sample.
	for i := 0; i < SoundOutRate; i++ {
		if s.ticks++; s.ticks >= GameBoyRate/(freq*8) {
			s.dutyStep = (s.dutyStep + 1) % 8
			s.ticks = 0
		}
	}

	if DutyCycles[s.NRx1>>6][s.dutyStep] {
		sample = s.envelope.Volume()
	}

	return
}
