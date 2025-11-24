// Package apu implements the Game Boy's sound generators as described in:
//
//   - [AUDIO1] https://gbdev.io/pandocs/Audio.html
//   - [AUDIO2] https://gbdev.io/pandocs/Audio_Registers.html
//   - [AUDIO3] https://gbdev.io/pandocs/Audio_details.html
//   - [AUDIOHW] https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware
package apu

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
	SamplingRate    = 44100 // How many sample frames to send per second.
	FramesPerBuffer = 256   // Number of sample frames fitting the audio buffer.
)

// Values to multiply with the final volume if not maximum (7) or zero.
var VolumeFactors = [7]float32{
	0.0,
	0.14285714285714285, // Volume 1
	0.2857142857142857,  // Volume 2
	0.42857142857142855, // Volume 3
	0.5714285714285714,  // Volume 4
	0.7142857142857143,  // Volume 5
	0.8571428571428571,  // Volume 6
}

// GameBoyRate is the main CPU frequence to be used in so many divisions.
// FIXME: that should probably live somewhere else, other modules could use it.
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

	// NR51 - Bit 7 - Output sound 4 to SO2 terminal
	NR51Output4Right uint8 = 1 << 7

	// NR51 - Bit 6 - Output sound 3 to SO2 terminal
	NR51Output3Right uint8 = 1 << 6

	// NR51 - Bit 5 - Output sound 2 to SO2 terminal
	NR51Output2Right uint8 = 1 << 5

	// NR51 - Bit 4 - Output sound 1 to SO2 terminal
	NR51Output1Right uint8 = 1 << 4

	// NR51 - Bit 3 - Output sound 4 to SO1 terminal
	NR51Output4Left uint8 = 1 << 3

	// NR51 - Bit 2 - Output sound 3 to SO1 terminal
	NR51Output3Left uint8 = 1 << 2

	// NR51 - Bit 1 - Output sound 2 to SO1 terminal
	NR51Output2Left uint8 = 1 << 1

	// NR51 - Bit 0 - Output sound 1 to SO1 terminal
	NR51Output1Left uint8 = 1
)

// APU structure grouping all sound signal generators and keeping track of when
// to actually output a sample for the sound card to play. It also takes care of
// volume control, channel mixing and stereo panning.
type APU struct {
	memory.MMU

	enabled bool

	ticks uint // APU-DIV ticks

	// Using full objects here instead of pointers in the vague hope that it may
	// somehow improve performance by having those contiguous in memory. Actually
	// benchmarking that might even make for a nice blog article.
	//
	// TODO: replace those with the usual pointers. See that it changes nothing.

	Square1 SquareWave
	Square2 SquareWave
	Wave    WaveTable
	Noise   Noise

	Mono  bool
	Muted [4]bool // Channels muted manually by the user

	NR50 uint8 // FF24 - Channel control / ON-OFF / Volume (R/W)
	NR51 uint8 // FF25 - Selection of Sound output terminal (R/W)
	NR52 uint8 // FF26 - Sound on/off
}

// New APU instance. So many registers.
func New(mono bool) *APU {
	a := APU{Wave: *NewWave()}

	a.Mono = mono

	// Make APU an address space covering its registers and the Wave Pattern
	// memory.
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
		AddrNR31: {Ptr: &a.Wave.NRx1, Mask: 0xff, OnWrite: a.Wave.SetNRx1},
		AddrNR32: {Ptr: &a.Wave.NRx2, Mask: 0x9f, OnWrite: a.Wave.SetNRx2},
		AddrNR33: {Ptr: &a.Wave.NRx3, Mask: 0xff, OnWrite: a.Wave.SetNRx3},
		AddrNR34: {Ptr: &a.Wave.NRx4, Mask: 0xbf, OnWrite: a.Wave.SetNRx4},
		AddrNR41: {Ptr: &a.Noise.NRx1, Mask: 0xff, OnWrite: a.Noise.SetNRx1},
		AddrNR42: {Ptr: &a.Noise.NRx2, Mask: 0x00, OnWrite: a.Noise.SetNRx2},
		AddrNR43: {Ptr: &a.Noise.NRx3, Mask: 0x00, OnWrite: a.Noise.SetNRx3},
		AddrNR44: {Ptr: &a.Noise.NRx4, Mask: 0xbf, OnWrite: a.Noise.SetNRx4},
		AddrNR50: {Ptr: &a.NR50},
		AddrNR51: {Ptr: &a.NR51},
		AddrNR52: {Ptr: &a.NR52, OnWrite: a.SetNR52},
	})
	a.Add(a.Wave.Pattern)

	// Set length timer max threshold (256 for wave, 64 for the others).
	a.Square1.length.Max = 64
	a.Square2.length.Max = 64
	a.Wave.length.Max = 256
	a.Noise.length.Max = 64

	// Pre-compute default frequencies.
	a.Square1.RecomputeFrequency()
	a.Square2.RecomputeFrequency()
	a.Wave.RecomputeFrequency()
	a.Noise.RecomputeFrequency() // This ensures divisor is not zero.

	return &a
}

func (a *APU) SetNR52(value uint8) {
	a.enabled = value&0x80 != 0
}

// Read overrides our internal Addressables to catch reads from NR52.
func (a *APU) Read(addr uint16) (value uint8) {
	if addr == AddrNR52 {
		if a.enabled {
			value |= 0x80
		}
		if a.Square1.Enabled {
			value |= 0x01
		}
		if a.Square2.Enabled {
			value |= 0x02
		}
		if a.Wave.Enabled {
			value |= 0x04
		}
		if a.Noise.Enabled {
			value |= 0x08
		}
		return value
	}
	return a.MMU.Read(addr)
}

// Tick updates the state machine of all signal generators whenever the DIV-APU
// timer increases.
func (a *APU) Tick() {
	a.ticks = (a.ticks + 1) % 8
	if a.ticks&1 == 0 { // Sound Length clocked every 2 DIV-APU ticks (256Hz).
		a.Square1.TickLength()
		a.Square2.TickLength()
		a.Wave.TickLength()
		a.Noise.TickLength()
	}
	if a.ticks&3 == 0 { // Frequency Sweep clocked every 4 DIV-APU ticks (128Hz).
		a.Square1.TickSweep()
	}
	if a.ticks&7 == 0 { // Volume Envelope clocked every 8 DIV-APU ticks (64Hz).
		a.Square1.envelope.Tick()
		a.Square2.envelope.Tick()
		a.Noise.envelope.Tick()
	}
}

// Sample produces a sample of the signal to generate based on the current state
// of all the APU's generators. This method should be called at the configured
// audio sample rate (i.e. CPU Frequency / Sample Rate).
func (a *APU) Sample() (left, right int8) {
	if !a.enabled {
		return
	}

	square1 := a.Square1.Sample()
	square2 := a.Square2.Sample()
	wave := a.Wave.Sample()
	noise := a.Noise.Sample()

	// Suppress output for channels the user manually muted.
	if a.Muted[0] {
		square1 = 0
	}

	if a.Muted[1] {
		square2 = 0
	}

	if a.Muted[2] {
		wave = 0
	}

	if a.Muted[3] {
		noise = 0
	}

	// Each channel can return a sample from -15 to +15. Even at those maxima,
	// adding the four values together should not overflow an int8.
	if a.NR51&NR51Output1Left != 0 || a.Mono {
		left += square1
	}

	if a.NR51&NR51Output2Left != 0 || a.Mono {
		left += square2
	}

	if a.NR51&NR51Output3Left != 0 || a.Mono {
		left += wave
	}

	if a.NR51&NR51Output4Left != 0 || a.Mono {
		left += noise
	}

	if a.NR51&NR51Output1Right != 0 || a.Mono {
		right += square1
	}

	if a.NR51&NR51Output2Right != 0 || a.Mono {
		right += square2
	}

	if a.NR51&NR51Output3Right != 0 || a.Mono {
		right += wave
	}

	if a.NR51&NR51Output4Right != 0 || a.Mono {
		right += noise
	}

	volumeLeft := a.NR50 & 0x07
	volumeRight := (a.NR50 & 0x70) >> 4

	// Adjust global volume for each channel.
	// XXX: should we just use floats for samples? Also what if Mono?
	if volumeLeft == 0 {
		left = 0
	} else if volumeLeft < 7 {
		left = int8(float32(left) * VolumeFactors[volumeLeft])
	}

	if volumeRight == 0 {
		right = 0
	} else if volumeRight < 7 {
		right = int8(float32(right) * VolumeFactors[volumeRight])
	}

	return
}
