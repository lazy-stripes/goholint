package apu

// Sources:
// https://gbdev.io/pandocs/Audio_Registers.html#ff10--nr10-channel-1-sweep
// https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware

// Sweep structure that will act as a state machine managing the frequency sweep
// for the first Square channel.
type Sweep struct {
	// The properties below can be set by the APU itself.
	Sweep    uint8 // NRx0 bits 6-4
	Increase bool  // NRx0 bit 3
	Shift    uint8 // NRx0 bits 2-0

	Shadow uint // Copy of Square 1's frequency

	enabled bool

	ticks      uint  // Clock ticks counter.
	sweepSteps uint8 // Sweep pace counter.
}

// Enable is called whenever the corresponding channel is triggered. If a new
// frequency was computed at this time, `updated` is set to true and the new
// frequency, as well as whether it overflows, are returned for the caller to
// update NRx3 and NRx4.
// XXX: Can't decide whether this is the way to go or if Sweep should have
//      direct access to Square1's registers.
func (s *Sweep) Enable(freq uint) (updated bool, newFreq uint, overflow bool) {
	// During a trigger event, several things occur:
	//
	// * Square 1's frequency is copied to the shadow register.
	//
	// * The sweep timer is reloaded.
	//
	// * The internal enabled flag is set if either the sweep period or shift
	//   are non-zero, cleared otherwise.
	//
	// * If the sweep shift is non-zero, frequency calculation and the overflow
	//   check are performed immediately.
	//
	// Source: https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware
	s.Shadow = freq
	s.enabled = s.Sweep > 0 || s.Shift > 0
	s.ticks = 0
	s.sweepSteps = 0

	if s.Shift > 0 {
		updated = true
		newFreq, overflow = s.UpdatedFrequency()
		return
	}

	// Tell caller the frequency hasn't been updated.
	return false, 0, false
}

// Performs frequency calculation and overflow check based on current shadow
// frequency and sweep parameters. Returns the updated frequency and a boolean
// indicating if it overflows (i.e. is bigger than 2047).
func (s *Sweep) UpdatedFrequency() (newFreq uint, overflow bool) {
	// the new wavelength Lₜ₊₁ is computed from the current one Lₜ as follows:
	// Lₜ₊₁ = Lₜ ± Lₜ/2ⁿ
	step := s.Shadow >> uint(s.Shift)
	if s.Increase {
		newFreq = s.Shadow + step
	} else {
		newFreq = s.Shadow - step
	}
	return newFreq, newFreq > 2047
}

// Tick advances the sweep one step. It will recompute a new frequency every
// every 1/128 seconds (i.e. <cpuFreq>/128 ticks). Returns whether the signal
// generator should update its frequency, a new frequency, and whether it
// overflows.
// Source: https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware
func (s *Sweep) Tick() (updated bool, newFreq uint, overflow bool) {
	if !s.enabled || s.Sweep == 0 {
		return
	}

	// Update frequency sweep step every <cpufreq>/128 (128Hz).
	stepRate := uint(GameBoyRate / 128)
	steps := (s.ticks + SoundOutRate) / stepRate
	s.ticks = (s.ticks + SoundOutRate) % stepRate

	// FIXME: there's got to be a more elegant way to do this.
	for ; steps > 0; steps-- {
		s.sweepSteps += 1
		if s.sweepSteps >= s.Sweep {
			newFreq, overflow = s.UpdatedFrequency()
			if !overflow && s.Shift > 0 {
				updated = true
				s.Shadow = newFreq

				// ... then frequency calculation and overflow check are run
				// AGAIN immediately using this new value, but this second new
				// frequency is not written back.
				// Source: https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware#Frequency_Sweep
				_, overflow = s.UpdatedFrequency()
			}
			s.sweepSteps = 0
		}
	}

	return
}
