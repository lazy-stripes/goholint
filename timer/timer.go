// Package timer implements the DMG Timer as described in:
// [TIMER1] http://gbdev.gg8.se/wiki/articles/Timer_and_Divider_Registers
// [TIMER2] http://gbdev.gg8.se/wiki/articles/Timer_Obscure_Behaviour
package timer

import (
	"github.com/lazy-stripes/goholint/interrupts"
	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/memory"
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
var FrequencyBits = [4]uint{9, 3, 5, 7}

// Timer address space handling timers and related interrupts.
type Timer struct {
	*memory.MMU

	Interrupts *interrupts.Interrupts
	DIV        uint16
	TIMA       uint8
	TMA        uint8
	TAC        uint8

	prevEdge bool // Falling edge detector of sorts
	ticks    int  // Only counted to measure overflow delay

	reloadDelay uint8 // Delay countdown for TIMA interrupt and TMA load
}

// New Timer instance.
func New() *Timer {
	return &Timer{MMU: memory.NewEmptyMMU()}
}

// Contains returns true is requested address is a timer register.
func (t *Timer) Contains(addr uint16) bool {
	return addr >= AddrDIV && addr <= AddrTAC
}

// Read a byte from one of the registers, accounting for DIV and TAC.
func (t *Timer) Read(addr uint16) (value uint8) {
	switch addr {
	case AddrDIV:
		value = uint8(t.DIV >> 8)
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
	case AddrTIMA:
		t.TIMA = value
	case AddrTMA:
		t.TMA = value
	case AddrTAC:
		t.TAC = value
	default:
		panic("Broken MMU")
	}
}

// Tick advances the timer state one step.
func (t *Timer) Tick() {
	t.DIV++

	// [TIMER2] Detect falling edge in (TAC.Freq AND TAC.Enable)
	bit := FrequencyBits[t.TAC&3]
	edge := (t.DIV&(1<<bit) != 0) && (t.TAC&4 != 0)
	if !edge && t.prevEdge {
		t.TIMA++
		if t.TIMA == 0 {
			// [TIMER2] TMA loading and interrupt are delayed 4 clocks.
			t.reloadDelay = 4
		}
	}
	t.prevEdge = edge

	if t.reloadDelay > 0 {
		t.reloadDelay--
		if t.reloadDelay == 0 {
			t.TIMA = t.TMA
			t.Interrupts.Request(interrupts.Timer)
		}
	}
}
