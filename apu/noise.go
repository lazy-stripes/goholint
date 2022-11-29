// Source: https://gbdev.io/pandocs/Audio_details.html
package apu

// Divisors for the generator's frequency depending on NR43.
var Divisors = map[uint8]uint{
	0: 8,
	1: 16,
	2: 32,
	3: 48,
	4: 64,
	5: 80,
	6: 96,
	7: 112,
}

// LFSR contains the value of a 16-bit Linear Feed Shift Register and a boolean
// defining whether "short mode" (reinjecting bit 7 during shifting) is enabled.
type LFSR struct {
	Value     uint16
	ShortMode bool
}

// Tick advances the LFSR one step and returns its rightmost bit as a boolean.
func (l *LFSR) Tick() bool {
	// [https://gbdev.io/pandocs/Audio_details.html]
	// CH4 revolves around an LFSR. The LFSR is 16-bit internally, but really
	// acts as if it was 15-bit.
	//
	// When CH4 is ticked (at the frequency specified via NR43):
	//  1. The result of ~(LFSR0 ⊕ LFSR1) (1 if bit 0 and bit 1 are identical, 0
	//     otherwise) is written to bit 15.
	//  2. If “short mode” was selected in NR43, then bit 15 is copied to bit 7
	//     as well.
	//  3. Finally, the entire LFSR is shifted right, and bit 0 selects between
	//     0 and the chosen volume.

	xor := (l.Value & 1) ^ ((l.Value & 2) >> 1)
	xor = (^xor) & 1 // Invert XORed value

	l.Value = l.Value & 0x7fff // Reset bit 15, put our XORed value there.
	l.Value |= xor << 15

	if l.ShortMode {
		l.Value = l.Value & 0xff7f // Reset bit 7, put our XORed value there.
		l.Value |= xor << 7
	}

	l.Value >>= 1

	return l.Value&1 == 1 // Convert bit 0 to bool
}

// Noise structure implementing sound sample generation for the fourth signal
// generator (A.k.a Sound4).
type Noise struct {
	NRx1 uint8 // Sound length
	NRx2 uint8 // Volume envelope
	NRx3 uint8 // Clock shift, Width mode of LFSR, Divisor code
	NRx4 uint8 // Counter/consecutive; Inital

	LFSR LFSR // 15-bit shift register

	enabled bool // Only output silence if this is false

	freq uint // Computed from NRx3

	ticks uint // Clock ticks counter for advancing duty step.

	envelope VolumeEnvelope
}

// RecomputeFrequency updates our internal frequency whenever NRx3 changes.
func (n *Noise) RecomputeFrequency() {
	shift := (n.NRx3 & 0xf0) >> 4 // Bit 7-4 - Clock shift
	divider := n.NRx3 & 0x07      // Bit 2-0 - Clock divider code
	n.freq = GameBoyRate / (Divisors[divider] << shift)
}

// SetNRx2 is called whenever the NRx2 register's value was changed, so that it
// can update the volume envelope state machine.
// TODO: make this method common with SquareWave.
func (n *Noise) SetNRx2(value uint8) {
	n.envelope.Initial = value >> 4
	n.envelope.Sweep = value & 7
	if value&NRx2EnvelopeDirection != 0 {
		n.envelope.Direction = 1
	} else {
		n.envelope.Direction = -1
	}
}

// SetNRx3 is called whenever the NRx3 register's value is written, so that it
// can update the internal generator's frequency.
func (n *Noise) SetNRx3(value uint8) {
	n.RecomputeFrequency()
	n.LFSR.ShortMode = value&0x04 != 0 // Bit 3 - LFSR width (15 or 7 bits)
}

// SetNRx4 is called whenever the NRx4 register's value is written, so that it
// can trigger the channel or enable the length counter.
func (n *Noise) SetNRx4(value uint8) {
	// Enable that signal if requested. NR14 being write-only, we can reset it
	// each time it goes to 1 without worrying.
	if value&NRx4RestartSound != 0 {
		n.NRx4 &= ^NRx4RestartSound // Reset trigger bit
		log.Debug("NR44 triggered")
		n.enabled = true // It's fine if the signal is already enabled.

		// The LFSR is set to 0 when (re)triggering the channel.
		n.LFSR.Value = 0

		n.ticks = 0

		n.envelope.Enable()
	}

	// TODO: bit 6 (length)

}

// Tick produces a sample of the signal to generate based on the current value
// in the signal generator's registers. We use a named return value, which is
// conveniently set to zero (silence) by default.
func (n *Noise) Tick() (sample uint8) {
	if !n.enabled {
		return
	}

	n.envelope.Tick()

	// Advance LFSR at the frequency requested by NR43.
	stepRate := GameBoyRate / n.freq
	steps := (n.ticks + SoundOutRate) / stepRate
	n.ticks = (n.ticks + SoundOutRate) % stepRate

	// FIXME: try simplifying this instead of blindly ticking.
	for i := uint(0); i < steps; i++ {
		n.LFSR.Tick()
	}

	if n.LFSR.Value&1 != 0 {
		sample = n.envelope.Volume()
	}

	return
}
