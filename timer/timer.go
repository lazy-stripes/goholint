package timer

import (
	"go.tigris.fr/gameboy/interrupts"
	"go.tigris.fr/gameboy/memory"
)

// Register addresses
const (
	AddrDIV  = 0xff04
	AddrTIMA = 0xff05
	AddrTMA  = 0xff06
	AddrTAC  = 0xff07
)

// Timer address space handling timers and related interrupts.
type Timer struct {
	*memory.MMU

	Interrupts *interrupts.Interrupts
	DIV        uint16
	TIMA       uint8
	TMA        uint8
	TAC        uint8

	ticks int
}

// New Timer instance.
func New() *Timer {
	t := Timer{memory.NewEmptyMMU()}
	t.Add(memory.Registers{
		AddrTIMA: &t.TIMA,
		AddrTMA:  &t.TMA,
		AddrTAC:  &t.TAC,
	})
	return &t
}

// Read a byte from one of the registers, accounting for DIV and TAC.
func (t *Timer) Read(addr uint) uint8 {
	switch addr {
	case AddrDIV:
		return uint8(t.DIV >> 8)
	case AddrTIMA:
		return t.TIMA
	case AddrTMA:
		return t.TMA
	case AddrTAC:
		return t.TAC & 0xf8
	default:
		return 0xff
	}
}

// Write a byte to one of the registers, accounting for DIV
func (t *Timer) Write(addr uint, value uint8) {
	if addr == AddrDIV {
		t.DIV = 0
	} else {
		t.MMU.Write(addr, value)
	}
}

// Tick advances the timer state one step.
func (t *Timer) Tick() {

}
