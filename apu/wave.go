package apu

import (
	"github.com/lazy-stripes/goholint/memory"
)

// OutputShift maps the output level code in NR32 with the amount of right
// shifts to apply to the generated sample.
var OutputShift = [4]int{
	4, // 0: Mute (no sound)
	0, // 1: 100% Volume (Produce Wave Pattern RAM Data as it is)
	1, // 2: 50% Volume (Produce Wave Pattern RAM data shifted once to the right)
	2, // 3: 25% Volume (Produce Wave Pattern RAM data shifted twice to the right)
}

// WaveTable structure implementing sound sample generation for the third
// signal generator (A.k.a Sound3).
type WaveTable struct {
	NRx0 uint8 // Sound on/off - bit 7
	NRx1 uint8 // Sound length
	NRx2 uint8 // Output level - bits 6-5
	NRx3 uint8 // Frequency's lower 8 bits
	NRx4 uint8 // Control and frequency' higher 3 bits

	Pattern *memory.RAM // Wave table pattern (32 4-bit samples)

	length Length

	freq uint // Computed from NRx3 and NRx4

	Enabled bool // Only output silence if this is false

	sampleOffset uint // Sub-index of the current sample into the wave table
	ticks        uint // Clock ticks counter for advancing sample index
}

// NewWave returns a WaveTable instance and is also kinda funny as a function
// name. Mostly it allocates 16 bytes of addressable RAM we'll pass along to
// the MMU.
func NewWave() *WaveTable {
	// Create RAM Addressable to store samples.
	// TODO: "Wave RAM should only be accessed while CH3 is disabled (NR30 bit
	// 7 reset), otherwise accesses will behave weirdly.
	//
	// On almost all models, the byte will be written at the offset CH3 is
	// currently reading. On GBA, the write will simply be ignored."
	// Source: https://gbdev.gg8.se/wiki/articles/Sound_Controller#FF30-FF3F_-_Wave_Pattern_RAM
	w := &WaveTable{Pattern: memory.NewRAM(AddrWavePattern, 16)}
	return w
}

// RecomputeFrequency updates our internal raw frequency value whenever NRx3 or
// NRx4 change.
func (w *WaveTable) RecomputeFrequency() {
	// With `x` the 11-bit value in NR33/NR34, frequency is 65536/(2048-x) Hz.
	rawFreq := ((uint(w.NRx4) & 7) << 8) | uint(w.NRx3)
	w.freq = 65536 / (2048 - rawFreq)
}

// SetNRx3 is called whenever the NRx3 register's value is written, so that it
// can update the internal generator's frequency.
func (w *WaveTable) SetNRx3(value uint8) {
	w.RecomputeFrequency()
}

// SetNRx4 is called whenever the NRx4 register's value is written, so that it
// can trigger the channel or update the internal generator's frequency.
func (w *WaveTable) SetNRx4(value uint8) {
	// Enable that signal if requested. NR34 being write-only, we can reset it
	// each time it goes to 1 without worrying.
	if w.NRx4&NRx4RestartSound != 0 {
		w.NRx4 &= ^NRx4RestartSound // Reset trigger bit
		log.Debug("NR34 triggered")
		w.Enabled = true // It's fine if the signal is already enabled.

		// TODO:
		// "When restarting CH3, it resumes playing the last 4-bit sample it
		// read from wave RAM, or 0 if no sample has been read since APU reset.
		// (Sample latching is independent of output level control in NR32.)
		// After the latched sample completes, it starts with the second sample
		// in wave RAM (low 4 bits of $FF30). The first sample (high 4 bits of
		// $FF30) is played last."
		// Source: https://gbdev.gg8.se/wiki/articles/Sound_Controller#PitFalls
		w.ticks = 0
	}

	// TODO: bit 6 (length)

	w.RecomputeFrequency()
}

// Tick produces a sample of the signal to generate based on the current value
// in the signal generator's registers. We use a named return value, which is
// conveniently set to zero (silence) by default.
func (w *WaveTable) Tick() (sample int8) {
	if !w.Enabled {
		return
	}

	if w.NRx0&NR30SoundOn == 0 {
		return
	}

	w.length.Tick()

	// Advance sample index every 1/(32f) where f is the sound's real frequency.
	stepRate := GameBoyRate / (w.freq * 32)
	steps := (w.ticks + SoundOutRate) / stepRate
	w.ticks = (w.ticks + SoundOutRate) % stepRate

	w.sampleOffset = (w.sampleOffset + steps) % 32

	// Each byte in the wave table contains 2 samples. Read it and only
	// output the proper nibble.
	sampleByte := w.sampleOffset / 2
	sampleShift := 4 - ((w.sampleOffset % 2) * 4) // Upper nibble first
	sample = int8((w.Pattern.Bytes[sampleByte] >> sampleShift) & 0xf)

	// Adjust for volume.
	sample >>= OutputShift[(w.NRx2&0x60)>>5]

	return sample
}
