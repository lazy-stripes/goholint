package apu

import "github.com/lazy-stripes/goholint/memory"

// WaveTable structure implementing sound sample generation for the third
// signal generator (A.k.a Sound3).
type WaveTable struct {
	NRx0 uint8 // Sound on/off - bit 7
	NRx1 uint8 // Sound length
	NRx2 uint8 // Output level - bits 6-5
	NRx3 uint8 // Frequency's lower 8 bits
	NRx4 uint8 // Control and frequency' higher 3 bits

	Pattern *memory.RAM // Wave table pattern (32 4-bit samples)

	enabled bool // Only output silence if this is false

	sample       uint8 // Current sample to play
	sampleOffset int   // Sub-index of the current sample into the wave table
	ticks        uint  // Clock ticks counter for advancing sample index
}

// NewWave returns a WaveTable instance and is also kinda funny as a function
// name. Mostly it allocates 16 bytes of addressable RAM we'll pass along to
// the MMU.
func NewWave() *WaveTable {
	// Create RAM Addressable to store samples.
	w := &WaveTable{Pattern: memory.NewRAM(0xff30, 16)}
	return w
}

// Tick produces a sample of the signal to generate based on the current value
// in the signal generator's registers. We use a named return value, which is
// conveniently set to zero (silence) by default.
func (w *WaveTable) Tick() (sample uint8) {
	// Enable that signal if requested. NR34 being write-only, we can reset it
	// each time it goes to 1 without worrying.
	if w.NRx4&NRx4RestartSound != 0 {
		w.NRx4 &= ^NRx4RestartSound // Reset trigger bit
		log.Debug("NR34 triggered")
		w.enabled = true // It's fine if the signal is already enabled.

		// "Restarting a pulse channel causes its "duty step timer" to reset."
		// Source: https://gbdev.gg8.se/wiki/articles/Sound_Controller#PitFalls
		w.ticks = 0
	}

	if !w.enabled {
		return
	}

	// With `x` the 11-bit value in NR33/NR34, frequency is 65536/(2048-x) Hz.
	rawFreq := ((uint(w.NRx4) & 7) << 8) | uint(w.NRx3)
	freq := 65536 / (2048 - rawFreq)

	// Advance sample index every 1/(32f) where f is the sound's real frequency.
	if w.ticks++; w.ticks >= GameBoyRate/(freq*32) {
		w.sampleOffset = (w.sampleOffset + 1) % 32
		w.ticks = 0

		// Each byte in the wave table contains 2 samples. Read it and only
		// output the proper nibble.
		sampleByte := w.sampleOffset / 2
		sampleShift := 4 - ((w.sampleOffset % 2) * 4) // Upper nibble first
		w.sample = (w.Pattern.Bytes[sampleByte] >> sampleShift) & 0xf
	}

	return w.sample
}
