package timer

import (
	"fmt"

	"go.tigris.fr/gameboy/interrupts"
	"go.tigris.fr/gameboy/log"
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

	ticks int // Only counted to measure overflow delay
}

// New Timer instance.
func New() *Timer {
	return &Timer{MMU: memory.NewEmptyMMU()}
}

// Contains returns true is requested address is a timer register.
func (t *Timer) Contains(addr uint) bool {
	return addr >= AddrDIV && addr <= AddrTAC
}

// Read a byte from one of the registers, accounting for DIV and TAC.
func (t *Timer) Read(addr uint) (value uint8) {
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
	if log.Enabled["timer"] {
		fmt.Printf("Timer.Read(0x%04x): 0x%02x\n", addr, value)
	}
	return
}

// Write a byte to one of the registers, accounting for DIV
func (t *Timer) Write(addr uint, value uint8) {
	if log.Enabled["timer"] {
		fmt.Printf("Timer.Write(0x%04x, 0x%02x)\n", addr, value)
	}
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
	// TODO: pretty much everything
	t.DIV++
}
