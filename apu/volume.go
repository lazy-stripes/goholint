package apu

// Sources:
//
// [AUDIOHW] https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware

// VolumeEnvelope structure that will act as a state machine only managing the
// current volume envelope for a Square or Noise signal generator.
type VolumeEnvelope struct {
	// The properties below can be set by the APU itself.
	Initial   uint8 // NRx2 bits 7-4
	Direction int8  // NRx2 bit 3
	Pace      uint8 // NRx2 bits 2-0

	enabled bool

	volume     int8  // Current calculated volume.
	ticks      uint  // Clock ticks counter.
	sweepSteps uint8 // Sweep pace counter.

}

// Reset is called whenever the corresponding channel is triggered.
func (v *VolumeEnvelope) Reset() {
	v.volume = int8(v.Initial)
	v.enabled = true
	v.ticks = 0
	v.sweepSteps = 0
}

// The envelope ticks at 64 Hz (i.e. every 8 DIV-APU ticks), and the channel’s
// envelope will be increased or decreased (depending on bit 3) every `Pace` of
// those ticks.
//
// Tick won't be called if the generator itself is not enabled.
func (v *VolumeEnvelope) Tick() {
	if !v.enabled {
		return
	}

	v.ticks++
	if v.ticks < 8 {
		return
	}
	v.ticks = 0

	// When the timer generates a clock and the envelope period is not zero, a
	// new volume is calculated by adding or subtracting (as set by NRx2) one
	// from the current volume. If this new volume within the 0 to 15 range,
	// the volume is updated, otherwise it is left unchanged and no further
	// automatic increments/decrements are made to the volume until the channel
	// is triggered again. [AUDIOHW]
	if (v.volume == 0 && v.Direction < 0) || (v.volume == 15 && v.Direction > 0) {
		v.enabled = false
		return
	}

	// Step up until we reach our Pace value.
	v.sweepSteps += 1
	if v.sweepSteps < v.Pace {
		return
	}

	v.volume += v.Direction
	v.sweepSteps = 0
}

// Volume returns the latest computed volume if the envelope Pace is not zero,
// or the initial volume if it is.
// FIXME: signedness
func (v *VolumeEnvelope) Volume() int8 {
	if v.Pace > 0 {
		return v.volume
	}
	return int8(v.Initial)
}
