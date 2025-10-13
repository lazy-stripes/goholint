package apu

// Sources:
//
// [OBSCURE] https://gbdev.io/pandocs/Audio_details.html#obscure-behavior
// [SWEEP] https://gbdev.io/pandocs/Audio_details.html#pulse-channel-with-sweep-ch1

// Sweep structure that will act as a state machine managing the frequency sweep
// for the first Square channel.
type Sweep struct {
	// The properties below can be set by the APU itself.
	Pace     uint8 // NR10 bits 6-4
	Increase bool  // NR10 bit 3
	Step     uint8 // NR10 bits 2-0

	shadow uint // Copy of Square 1's frequency

	enabled bool

	sweepTimer uint8 // Sweep iteration counter.
}

func (s *Sweep) ReloadTimer() {
	s.sweepTimer = s.Pace
	if s.sweepTimer == 0 {
		// The volume envelope and sweep timers treat a period of 0 as 8.
		// [OBSCURE]
		s.sweepTimer = 8
	}
}

// Reset is called whenever the corresponding channel is triggered. If a new
// frequency was computed at this time, `updated` is set to true and the new
// frequency, as well as whether it overflows, are returned for the caller to
// update NRx3 and NRx4.
func (s *Sweep) Reset(freq uint) (updated bool, newFreq uint, overflow bool) {
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
	//   check are performed immediately [the frequency itself is *not* stored
	//   back to NR13/NR14].
	//
	// Source: https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware
	s.shadow = freq
	s.enabled = s.Pace > 0 || s.Step > 0
	s.ReloadTimer()

	if s.Step > 0 {
		_, overflow = s.UpdatedFrequency()
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
	step := s.shadow >> uint(s.Step)
	if s.Increase {
		newFreq = s.shadow + step
	} else {
		newFreq = s.shadow - step
	}
	return newFreq, newFreq > 2047
}

// Tick advances the sweep one step. It will recompute a new frequency every
// 1/128 seconds (i.e. every 4 DIV-APU ticks). Returns whether the signal
// generator should update its frequency, a new frequency, and whether it
// overflows.
//
// Tick won't be called if the generator itself is not enabled.
func (s *Sweep) Tick() (updated bool, newFreq uint, overflow bool) {
	if !s.enabled || s.Pace == 0 {
		return
	}

	s.sweepTimer--
	if s.sweepTimer <= 0 {
		newFreq, overflow = s.UpdatedFrequency()
		if !overflow && s.Step > 0 {
			updated = true
			s.shadow = newFreq

			// ... then frequency calculation and overflow check are run
			// AGAIN immediately using this new value, but this second new
			// frequency is not written back. [SWEEP]
			_, overflow = s.UpdatedFrequency()
		}
		s.ReloadTimer()
	}

	return
}
