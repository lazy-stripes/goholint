package apu

// Source: [SOUND1] https://gbdev.gg8.se/wiki/articles/Sound_Controller
//         [SOUND2] https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware

import (
	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/memory"
)

// Package-wide logger.
var log = logger.New("apu", "sound-related operations")

// Register addresses... yeah, sound is complicated.
const (
	AddrNR10 = 0xff10
	AddrNR11 = 0xff11
	AddrNR12 = 0xff12
	AddrNR13 = 0xff13
	AddrNR14 = 0xff14

	AddrNR21 = 0xff16
	AddrNR22 = 0xff17
	AddrNR23 = 0xff18
	AddrNR24 = 0xff19

	AddrNR30 = 0xff1a
	AddrNR31 = 0xff1b
	AddrNR32 = 0xff1c
	AddrNR33 = 0xff1d
	AddrNR34 = 0xff1e

	AddrNR41 = 0xff20
	AddrNR42 = 0xff21
	AddrNR43 = 0xff22
	AddrNR44 = 0xff23

	AddrNR50 = 0xff24
	AddrNR51 = 0xff25
	AddrNR52 = 0xff26

	AddrWavePattern = 0xff30
)

// Audio settings for SDL.

const (
	SamplingRate    = 22050 // How many sample frames to send per second.
	FramesPerBuffer = 1024  // Number of sample frames fitting the audio buffer.
	Volume          = 63    // 25% volume for unsigned 8-bit samples.
)

// GameBoyRate is the main CPU frequence to be used in so many divisions.
const GameBoyRate = 4 * 1024 * 1024 // 4194304Hz or 4MiHz

// SoundOutRate represents CPU cycles to wait before producing one sample frame.
const SoundOutRate = GameBoyRate / SamplingRate

// Audio Control register bits.
const (
	// NRx2 - Bit 3 - Envelope Direction (0=Decrease, 1=Increase)
	NRx2EnvelopeDirection uint8 = 1 << 3

	// NR30 - Bit 7 - Sound Channel 3 Off  (0=Stop, 1=Playback)
	NR30SoundOn uint8 = 1 << 7

	// NRx4 - Bit 7 - Initial (1=Restart Sound)
	NRx4RestartSound uint8 = 1 << 7

	// NRx4 - Bit 6 - Sound Length (1=Stop output when length in NR11 expires)
	NRx4EnableLength uint8 = 1 << 6

	// NR43 - Bit 3 - Counter Step/Width (0=15 bits, 1=7 bits)
	NR43Width7 uint8 = 1 << 3
)

// APU structure grouping all sound signal generators and keeping track of when
// to actually output a sample for the sound card to play. For now we only use
// two generators for sterao sound, but in time, we'll mix the output of four of
// those and the stereo channel they'll go to will be configurable as well.
type APU struct {
	memory.MMU

	Square1 SquareWave
	Square2 SquareWave
	Wave    WaveTable
	Noise   Noise
}

// New APU instance. So many registers.
func New() *APU {
	a := APU{Wave: *NewWave()}

	// Make APU an address space covering its registers and the Wave Pattern
	// memory.
	a.Add(a.Wave.Pattern)
	a.Add(APURegisters{
		AddrNR10: {Ptr: &a.Square1.NRx0, Mask: 0x80, OnWrite: a.Square1.SetNRx0},
		AddrNR11: {Ptr: &a.Square1.NRx1, Mask: 0x3f, OnWrite: a.Square1.SetNRx1},
		AddrNR12: {Ptr: &a.Square1.NRx2, Mask: 0x00, OnWrite: a.Square1.SetNRx2},
		AddrNR13: {Ptr: &a.Square1.NRx3, Mask: 0xff, OnWrite: a.Square1.SetNRx3},
		AddrNR14: {Ptr: &a.Square1.NRx4, Mask: 0xbf, OnWrite: a.Square1.SetNRx4},
		AddrNR21: {Ptr: &a.Square2.NRx1, Mask: 0x3f, OnWrite: a.Square2.SetNRx1},
		AddrNR22: {Ptr: &a.Square2.NRx2, Mask: 0x00, OnWrite: a.Square2.SetNRx2},
		AddrNR23: {Ptr: &a.Square2.NRx3, Mask: 0xff, OnWrite: a.Square2.SetNRx3},
		AddrNR24: {Ptr: &a.Square2.NRx4, Mask: 0xbf, OnWrite: a.Square2.SetNRx4},
		AddrNR30: {Ptr: &a.Wave.NRx0, Mask: 0x7f},
		AddrNR31: {Ptr: &a.Wave.NRx1, Mask: 0xff},
		AddrNR32: {Ptr: &a.Wave.NRx2, Mask: 0x9f},
		AddrNR33: {Ptr: &a.Wave.NRx3, Mask: 0xff, OnWrite: a.Wave.SetNRx3},
		AddrNR34: {Ptr: &a.Wave.NRx4, Mask: 0xbf, OnWrite: a.Wave.SetNRx4},
		AddrNR41: {Ptr: &a.Noise.NRx1, Mask: 0xff, OnWrite: a.Noise.SetNRx1},
		AddrNR42: {Ptr: &a.Noise.NRx2, Mask: 0x00, OnWrite: a.Noise.SetNRx2},
		AddrNR43: {Ptr: &a.Noise.NRx3, Mask: 0x00, OnWrite: a.Noise.SetNRx3},
		AddrNR44: {Ptr: &a.Noise.NRx4, Mask: 0xbf, OnWrite: a.Noise.SetNRx4},
	})

	// Pre-compute default frequencies.
	a.Square1.RecomputeFrequency()
	a.Square2.RecomputeFrequency()
	a.Wave.RecomputeFrequency()
	a.Noise.RecomputeFrequency() // This makes sure divisor is not zero.

	return &a
}

// Tick advances the state machine of all signal generators to produce a single
// stereo sample for the sound card. Note that the number of internal cycles
// happening on each signal generator depends on the output frequency.
func (a *APU) Tick() (left, right uint8) {
	// Advance all signal generators a step. Right now we only have two but
	// if we were to implement all four, we'd actually mix all their outputs
	// together here (with various per-generator parameters to account for).

	// TODO: mix signals here according to the relevant registers.
	// TODO: use signed samples for proper addition.
	// Because we're returning unsigned ints, the silence point is at 128.
	//left = 128 + a.Square1.Tick() - a.Square2.Tick() + a.Wave.Tick()// - a.Noise.Tick()
	left = a.Square1.Tick() + a.Square2.Tick() + a.Wave.Tick() + a.Noise.Tick()
	right = left

	return
}
