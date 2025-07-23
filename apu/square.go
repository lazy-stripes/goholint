package apu

// [AUDIODETAILS] https://gbdev.io/pandocs/Audio_details.html

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

	Enabled bool // Only output silence if this is false

	freq uint // Real signal frequency computed from NRx3 and NRx4

	// Duty-related variables.
	dutyStep uint // Sub-index into DutyCycles to set the signal high or low.
	ticks    uint // Clock ticks counter for advancing duty step.

	lengthEnabled bool
	sweepEnabled  bool

	length   Length
	envelope VolumeEnvelope
	sweep    Sweep
}

// RawFrequency returns the base frequency value stored in NRx3 and NRx4, an
// 11-bit value between 0 and 2047.
func (s *SquareWave) RawFrequency() (rawFreq uint) {
	return ((uint(s.NRx4) & 0x07) << 8) | uint(s.NRx3)
}

// SetRawFrequency sets the base frequency in NRx3 and NRx4 from the given value
// between 0 and 2047. If the given value is larger, nothing happens.
func (s *SquareWave) SetRawFrequency(rawFreq uint) {
	if rawFreq > 2047 {
		log.Warningf("invalid base frequency %d for square channel", rawFreq)
		return
	}

	s.NRx3 = uint8(rawFreq)
	s.NRx4 &= 0xf8
	s.NRx4 |= uint8(rawFreq>>8) & 0x07

	s.RecomputeFrequency()
}

// RecomputeFrequency updates our internal raw frequency value whenever NRx3 or
// NRx4 change.
func (s *SquareWave) RecomputeFrequency() {
	// With `x` the 11-bit value in NR13/NR14, frequency is 131072/(2048-x) Hz.
	rawFreq := ((uint(s.NRx4) & 7) << 8) | uint(s.NRx3)
	s.freq = 131072 / (2048 - rawFreq)
}

// SetNRx0 is called whenever the NRx0 register's value was changed, so that it
// can update the sweep parameters.
func (s *SquareWave) SetNRx0(value uint8) {
	s.sweep.Pace = (value & 0x70) >> 4
	s.sweep.Increase = value&0x08 == 0
	s.sweep.Step = value & 0x07

	s.sweepEnabled = (s.sweep.Pace != 0)
}

// SetNRx1 is called whenever the NRx1 register's value was changed, so that it
// can update the length timer.
func (s *SquareWave) SetNRx1(value uint8) {
	s.length.Initial = value & 0x3f
}

// SetNRx2 is called whenever the NRx2 register's value was changed, so that it
// can update the volume envelope state machine.
func (s *SquareWave) SetNRx2(value uint8) {
	s.envelope.Initial = value >> 4
	s.envelope.Pace = value & 7
	if value&NRx2EnvelopeDirection != 0 {
		s.envelope.Direction = 1
	} else {
		s.envelope.Direction = -1
	}
}

// SetNRx3 is called whenever the NRx3 register's value is written, so that it
// can update the internal generator's frequency.
func (s *SquareWave) SetNRx3(value uint8) {
	s.RecomputeFrequency()
}

// SetNRx4 is called whenever the NRx4 register's value is written, so that it
// can trigger the channel or update the internal generator's frequency.
func (s *SquareWave) SetNRx4(value uint8) {
	// Recompute frequency in case bits 0-3 changed.
	s.RecomputeFrequency()

	s.lengthEnabled = (value&NRx4EnableLength != 0)

	// Enable that signal if requested. NR14 being write-only, we can reset it
	// each time it goes to 1 without worrying.
	if value&NRx4RestartSound != 0 {
		s.NRx4 &= ^NRx4RestartSound // Reset trigger bit

		s.Enabled = true // It's fine if the signal is already enabled.

		// "Restarting a pulse channel causes its "duty step timer" to reset."
		// Source: https://gbdev.gg8.se/wiki/articles/Sound_Controller#PitFalls
		s.ticks = 0

		s.length.Reset(64)

		s.envelope.Reset()

		// Enable sweep, see if a frequency change was already computed.
		updated, newFreq, overflow := s.sweep.Reset(s.RawFrequency())
		if overflow {
			s.Enabled = false
		} else {
			if updated {
				s.SetRawFrequency(newFreq)
			}
		}
	}

}

// Tick is called whenever DIV-APU increases. It will tick the signal generator's
// length, sweep and envelope.
func (s *SquareWave) Tick() {
	if !s.Enabled {
		return
	}

	// Tick length. It will be updated at 256Hz (or every 2 DIV-APU ticks).
	if s.lengthEnabled {
		disabled := s.length.Tick()
		if disabled {
			s.Enabled = false
			return
		}
	}

	// Tick sweep. It will be updated at 128Hz (or every 4 DIV-APU ticks).
	if s.sweepEnabled {
		updated, newFreq, overflow := s.sweep.Tick()
		if updated {
			s.SetRawFrequency(newFreq)
		}

		if overflow {
			s.Enabled = false
			return
		}
	}

	// Tick envelope. It will be updated at 64Hz (or every 8 DIV-APU ticks).
	s.envelope.Tick()
}

// Sample produces a sample of the signal to generate based on the current value
// in the signal generator's registers. We use a named return value, which is
// conveniently set to zero (silence) by default.
func (s *SquareWave) Sample() (sample int8) {
	if !s.Enabled {
		return
	}

	// Advance duty step every 1/(8f) where f is the sound's real frequency.
	stepRate := GameBoyRate / (s.freq * 8)
	steps := (s.ticks + SoundOutRate) / stepRate
	s.ticks = (s.ticks + SoundOutRate) % stepRate

	s.dutyStep = (s.dutyStep + steps) % 8

	// FIXME: The digital value produced by the generator, which ranges between
	// $0 and $F (0 and 15), is linearly translated by the DAC into an analog
	// value between -1 and 1 (the unit is arbitrary). [AUDIODETAILS]
	if DutyCycles[s.NRx1>>6][s.dutyStep] {
		sample = s.envelope.Volume()
	} else {
		sample = -s.envelope.Volume()
	}

	return
}
