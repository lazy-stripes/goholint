package apu

// Divisors for the generator's frequency depending on NR43.
var Divisors = map[uint8]int{
	0: 8,
	1: 16,
	2: 32,
	3: 48,
	4: 64,
	5: 80,
	6: 96,
	7: 112,
}

// Noise structure implementing sound sample generation for the fourth signal
// generator (A.k.a Sound4).
type Noise struct {
	NRx1 uint8 // Sound length
	NRx2 uint8 // Volume envelope
	NRx3 uint8 // Polynomial counter and frequency
	NRx4 uint8 // Counter/consecutive; Inital

	enabled bool // Only output silence if this is false

	register uint16 // 15-bit shift register

	output uint8 // Current output

	ticks uint // Clock ticks counter for advancing duty step.

	envelope VolumeEnvelope
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

// Tick produces a sample of the signal to generate based on the current value
// in the signal generator's registers. We use a named return value, which is
// conveniently set to zero (silence) by default.
func (n *Noise) Tick() (sample uint8) {
	// Enable that signal if requested. NR34 being write-only, we can reset it
	// each time it goes to 1 without worrying.
	if n.NRx4&NRx4RestartSound != 0 {
		n.NRx4 &= ^NRx4RestartSound // Reset trigger bit
		log.Debug("NR44 triggered")
		n.enabled = true

		// [AUDIO2] Trigger Event details:
		// Frequency timer is reloaded with period. (TODO: do it that way instead of from 0)
		// Noise channel's LFSR bits are all set to 1.
		n.ticks = 0
		n.register = 0x7fff // Keep 16th bit zero in prevision for shifting

		log.Debugf("NR43=0x%02x", n.NRx3)
		s := n.NRx3 >> 4
		r := Divisors[n.NRx3&7]
		divisor := r << s
		log.Debugf("s=%d, r=%d, div=%d", r, s, divisor)
		log.Debugf("1048576 / (r + 1) / (1 << (s + 1))=%d", 1048576/(r+1)/(1<<(s+1)))
		log.Debugf("(1048576 / (r + 1)) / (1 << (s + 1))=%d", (1048576/(r+1))/(1<<(s+1)))

	}

	if !n.enabled {
		return
	}

	n.envelope.Tick()

	// [AUDIO1] Frequency = 524288 Hz / r / 2^(s+1) ;For r=0 assume r=0.5 instead
	// [AUDIO2] More details about divisor code.
	// [... reddit] Frequency = 1048576 Hz / (ratio * 2) / (2 ^ (shiftclockfreq + 1))
	// GBSOUND.txt PRNG Frequency = (1048576 Hz / (ratio + 1)) / (2 ^ (shiftclockfreq + 1))
	// TODO: pre-compute frequencies (in squarewave too). This is suboptimal. Also precomputing will be easier to debug.
	s := n.NRx3 >> 4
	r := Divisors[n.NRx3&7]
	//r := uint(n.NRx3 & 7)
	if r == 0 {
		r = 1
	}
	rawFreq := 1048576 / (r + 1) / (1 << (s + 1))
	if rawFreq > SamplingRate {
		rawFreq = SamplingRate
	}
	freq := uint(rawFreq / 16)

	// Update register at the required frequency. FIXME: why is this not working?
	if n.ticks++; n.ticks >= freq {
		// [AUDIO2] When clocked by the frequency timer, the low two bits (0 and
		// 1) are XORed, all bits are shifted right by one, and the result of
		// the XOR is put into the now-empty 15th bit. If width mode is 1
		// (NR43), the XOR result is ALSO put into bit 6 AFTER the shift,
		// resulting in a 7-bit LFSR. The waveform output is bit 0 of the LFSR,
		// INVERTED.
		xor := (n.register & 1) ^ ((n.register & 2) >> 1)
		n.register >>= 1
		n.register &= ^uint16(1<<14) & 0x7fff // Reset bit 14 in case xor is zero
		n.register |= xor << 14
		if n.NRx3&NR43Width7 != 0 {
			n.register &= ^uint16(1<<6) & 0x7fff
			n.register |= xor << 6
		}
		n.ticks = 0
		n.output = uint8((^n.register) & 1)

		log.Desperatef("LFSR=%16b", n.register)

	}

	return n.output * n.envelope.Volume()
}
