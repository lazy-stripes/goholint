// Package timer implements the DMG Timer as described in:
// [TIMER1] https://gbdev.io/pandocs/Timer_and_Divider_Registers.html
// [TIMER2] https://gbdev.io/pandocs/Timer_Obscure_Behaviour.html
//
// And also in:
// [DIV-APU] https://gbdev.io/pandocs/Audio_details.html
// Because, of course, sound is complicated.
package timer

import (
	"github.com/lazy-stripes/goholint/apu"
	"github.com/lazy-stripes/goholint/interrupts"
	"github.com/lazy-stripes/goholint/logger"
)

// Package-wide logger.
var log = logger.New("timer", "timer registers read/writes (Desperate level only)")

// Register addresses
const (
	AddrDIV  = 0xff04
	AddrTIMA = 0xff05
	AddrTMA  = 0xff06
	AddrTAC  = 0xff07
)

// FrequencyBits map TAC frequency select value to related bits in DIV.
var FrequencyBits = [4]uint8{7, 1, 3, 5}

// Timer address space handling timers and related interrupts.
type Timer struct {
	Interrupts *interrupts.Interrupts
	APU        *apu.APU

	DIV  uint16 // Actually 14-bit wide, only bits 7-14 are visible.
	TIMA uint8
	TMA  uint8
	TAC  uint8

	divApu uint8 // DIV-APU value, will loop from 0 to 7.

	freqBit uint8 // Saved frequency bit updated when TAC is written to.

	lastAudioEdge bool // Falling edge detector of sorts.
	lastTimerEdge bool // Falling edge detector of sorts.

	reloadDelay uint8 // Delay countdown for TIMA interrupt and TMA load.
}

// New Timer instance.
func New(ints *interrupts.Interrupts, apu *apu.APU) *Timer {
	return &Timer{
		Interrupts: ints,
		APU:        apu,
		freqBit:    FrequencyBits[0],
	}
}

// Contains returns true is requested address is a timer register.
func (t *Timer) Contains(addr uint16) bool {
	return addr >= AddrDIV && addr <= AddrTAC
}

// Read a byte from one of the registers, accounting for DIV and TAC.
func (t *Timer) Read(addr uint16) (value uint8) {
	switch addr {
	case AddrDIV:
		value = uint8(t.DIV >> 6)
	case AddrTIMA:
		value = t.TIMA
	case AddrTMA:
		value = t.TMA
	case AddrTAC:
		value = t.TAC & 0xf8
	default:
		panic("Broken MMU")
	}
	log.Desperatef("Timer.Read(0x%04x): 0x%02x", addr, value)
	return value
}

// Write a byte to one of the registers, accounting for DIV.
func (t *Timer) Write(addr uint16, value uint8) {
	log.Desperatef("Timer.Write(0x%04x, 0x%02x)", addr, value)
	switch addr {
	case AddrDIV:
		t.DIV = 0
		t.checkForTimerEvent()
		t.checkForAudioEvent()
	case AddrTIMA:
		t.TIMA = value
	case AddrTMA:
		t.TMA = value
	case AddrTAC:
		t.TAC = value
		t.freqBit = FrequencyBits[value&3]
		t.checkForTimerEvent()
	default:
		panic("Broken MMU")
	}
}

// Check whether falling-edge occurred, and do stuff if it did.
func (t *Timer) checkForTimerEvent() {
	// Detect TAC-related edge falling.
	edge := (t.DIV&(1<<t.freqBit) != 0) && (t.TAC&4 != 0)
	if !edge && t.lastTimerEdge {
		// Timer tick event.
		t.TIMA++
		if t.TIMA == 0 {
			// [TIMER2] TMA loading and interrupt are delayed 1 M-cycle.
			t.reloadDelay = 1
		}
	}
	t.lastTimerEdge = edge
}

// Check whether falling-edge occurred, and do stuff if it did.
func (t *Timer) checkForAudioEvent() {
	// [DIV-APU] Counter is increased every time DIV’s bit 4 goes from 1 to 0.
	edge := t.DIV&(1<<10) != 0
	if !edge && t.lastAudioEdge {
		// Audio event.
		t.divApu = (t.divApu + 1) % 8
		if t.divApu&1 == 0 { // Sound Length clocked every 2 DIV-APU ticks.
			//t.APU.TickLength() TODO
		}

		if t.divApu&3 == 0 { // Frequency Sweep clocked every 4 DIV-APU ticks.
			//t.APU.TickSweep() TODO
		}

		if t.divApu&7 == 0 { // Volume Envelope clocked every 8 DIV-APU ticks.
			//t.APU.TickEnvelope() TODO
		}
	}
	t.lastAudioEdge = edge
}

// Tick advances the timer state one step. This should be called every M-cycle
// and since the exposed part of DIV starts at bit 6, this will look like DIV
// increases at 16KihZ (4MiHz/4/2^6).
func (t *Timer) Tick() {
	// [TIMER2] shows DIV as being 14-bit wide, so we wrap around beyond that.
	t.DIV = (t.DIV + 1) & ((1 << 14) - 1)

	// If a timer event occurred, TMA copy and interrupt will happen, but after
	// one M-cycle delay.
	if t.reloadDelay > 0 {
		t.reloadDelay--
		if t.reloadDelay == 0 {
			t.TIMA = t.TMA
			t.Interrupts.Request(interrupts.Timer)
		}
	}

	// Detect falling edge in (TAC.Freq AND TAC.Enable).
	t.checkForTimerEvent()

	// Detect falling edge in DIV-APU (DIV&(1<<4))
	t.checkForAudioEvent()
}
