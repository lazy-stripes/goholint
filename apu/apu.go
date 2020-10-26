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
)

// Audio settings for SDL.

// Constant values you can tweak to see their effect on the produced sound.
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

	// NRx4 - Bit 7 - Initial (1=Restart Sound)
	NRx4RestartSound uint8 = 1 << 7

	// NR43 - Bit 3 - Counter Step/Width (0=15 bits, 1=7 bits)
	NR43Width7 uint8 = 1 << 3
)

// APU structure grouping all sound signal generators and keeping track of when
// to actually output a sample for the sound card to play. For now we only use
// two generators for sterao sound, but in time, we'll mix the output of four of
// those and the stereo channel they'll go to will be configurable as well.
type APU struct {
	memory.Registers

	Square1 SquareWave
	Square2 SquareWave
	Wave    WaveTable
	Noise   Noise

	ticks uint // Clock ticks counter for mixing samples
}

// New APU instance. So many registers.
func New() *APU {
	a := APU{Wave: *NewWave()}

	a.Registers = memory.Registers{
		AddrNR10: &a.Square1.NRx0,
		AddrNR11: &a.Square1.NRx1,
		AddrNR12: &a.Square1.NRx2,
		AddrNR13: &a.Square1.NRx3,
		AddrNR14: &a.Square1.NRx4,
		AddrNR21: &a.Square2.NRx1,
		AddrNR22: &a.Square2.NRx2,
		AddrNR23: &a.Square2.NRx3,
		AddrNR24: &a.Square2.NRx4,
		AddrNR31: &a.Wave.NRx1,
		AddrNR32: &a.Wave.NRx2,
		AddrNR33: &a.Wave.NRx3,
		AddrNR34: &a.Wave.NRx4,
		AddrNR41: &a.Noise.NRx1,
		AddrNR42: &a.Noise.NRx2,
		AddrNR43: &a.Noise.NRx3,
		AddrNR44: &a.Noise.NRx4,
	}

	return &a
}

// Overrides to Read/Write methods because of masks and special cases.
func (a *APU) Write(addr uint16, value uint8) {
	// Do write the value anyway. TODO: masks. Ugh.
	a.Registers.Write(addr, value)

	// Special case for some registers.
	switch addr {
	case AddrNR12:
		log.Debugf("NR12 = 0x%02x", value)
		a.Square1.SetNRx2(value)
	case AddrNR22:
		log.Debugf("NR22 = 0x%02x", value)
		a.Square2.SetNRx2(value)
	case AddrNR42:
		log.Debugf("NR42 = 0x%02x", value)
		a.Noise.SetNRx2(value)
	}
}

// Tick advances the state machine of all signal generators to produce a single
// stereo sample for the sound card. This sample is only actually sent to the
// sound card at the chosen sampling rate.
func (a *APU) Tick() (left, right uint8, play bool) {
	// Advance all signal generators a step. Right now we only have two but
	// if we were to implement all four, we'd actually mix all their outputs
	// together here (with various per-generator parameters to account for).

	// TODO: mix signals here according to the relevant registers.
	left = a.Square1.Tick() + a.Square2.Tick() + a.Wave.Tick() + a.Noise.Tick()
	right = left

	// We're ticking as fast as the Game Boy CPU goes, but our sound sample rate
	// is much lower than that so we only need to yield an actual sample every
	// so often.
	// Yes I'm probably missing a very obvious way to optimize this all.
	if a.ticks++; a.ticks >= SoundOutRate {
		// TODO: mix channels here.

		a.ticks = 0
		play = true
	}

	return
}
