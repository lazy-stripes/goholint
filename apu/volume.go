package apu

// VolumeEnvelope structure that will act as a state machine only managing the
// current volume envelope for a Square or Noise signal generator.
type VolumeEnvelope struct {
	// The properties below can be set by the APU itself.
	Initial   uint8 // NRx2 bits 7-4
	Direction int8  // NRx2 bit 3
	Sweep     uint8 // NRx2 bits 2-0

	enabled bool

	volume int8 // Current calculated volume.
	ticks  uint // Clock ticks counter.

}

// Enable is called whenever the corresponding channel is triggered.
func (v *VolumeEnvelope) Enable() {
	v.volume = int8(v.Initial)
	v.enabled = true
	v.ticks = 0
}

// Disable is called whenever sound for the corresponding channel is disabled.
func (v *VolumeEnvelope) Disable() {
	v.enabled = false
}

// Tick advances the volume envelope one step. It will adjust the volume value
// every <sweep>×(1/64) seconds or <sweep>×(<sample rate>/64) APU ticks.
// Source: https://gbdev.gg8.se/wiki/articles/Sound_Controller about NR12.
func (v *VolumeEnvelope) Tick() {
	if !v.enabled {
		return
	}

	// Volume must always stay in the 0-15 range.
	if (v.volume == 0 && v.Direction < 0) || (v.volume == 15 && v.Direction > 0) {
		v.enabled = false
		return
	}

	// Update volume every <sweep>×(<sample rate>/64) APU ticks.
	v.ticks++
	if v.ticks < uint(v.Sweep)*(SamplingRate/64) {
		return
	}
	v.ticks = 0

	v.volume += v.Direction
}

// Volume returns the latest computed volume if the envelope sweep is not zero,
// or the initial volume if it is.
func (v *VolumeEnvelope) Volume() uint8 {
	if v.Sweep > 0 {
		return uint8(v.volume)
	}
	return v.Initial
}
