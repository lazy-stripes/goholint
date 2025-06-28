package serial

import (
	"fmt"

	"github.com/lazy-stripes/goholint/apu"
	"github.com/lazy-stripes/goholint/interrupts"
	"github.com/lazy-stripes/goholint/logger"
)

// Register addresses
const (
	AddrSB = 0xff01
	AddrSC = 0xff02
)

const SerialRate = 8192 // Serial clock is 8KiHz

const BitOutRate = apu.GameBoyRate / SerialRate

// Serial registers for game link. Used only for debug for now.
type Serial struct {
	SB, SC     uint8
	Interrupts *interrupts.Interrupts

	master  bool // True if using internal clock.
	enabled bool // True if currently transmitting.
	bit     int  // Number of SC bit to transmit next.

	ticks uint
}

// New instantiates a Serial addressable mapping to FF01 and FF02.
func New() *Serial {
	return &Serial{}
}

// Tick should be called at a 8KiHZ rate.
func (s *Serial) Tick() {
	if !s.enabled {
		return
	}

	// TODO: implement external ticks for slave mode.
	if !s.master {
		return
	}

	s.ticks += 1
	if s.ticks < BitOutRate {
		return
	}
	s.ticks = 0

	if logger.Enabled["serial"] {
		// Print characters for now to print GB ROM test results.
		fmt.Printf("SB=%#08[1]b (%#02[1]x)\n", s.SB)
	}

	// Shift our next bit out, read their next bit in. For now we'll always
	// assume there is no connection active.
	s.SB = (s.SB << 1) | 0x01

	s.bit += 1
	if s.bit > 7 {
		s.bit = 0
		s.enabled = false
		s.SC &= 0x7f // Clear SC's high bit.
		s.Interrupts.Request(interrupts.Serial)
	}
}

func (s *Serial) Contains(addr uint16) bool {
	return addr == AddrSB || addr == AddrSC
}

func (s *Serial) Read(addr uint16) uint8 {
	if addr == AddrSB {
		return s.SB
	} else if addr == AddrSC {
		return s.SC
	} else {
		panic(fmt.Sprintf("broken MMU: illegal address %04x requested!", addr))
	}
}

func (s *Serial) Write(addr uint16, value uint8) {
	if addr == AddrSB {
		s.SB = value
	} else if addr == AddrSC {
		// FIXME: What to do if a transfer is already in progress?
		// The Serial register is used by a lot of test ROMs for printing
		// out logs to standard output. We'll still do it if logging is
		// specifically requested.
		// FIXME: There should be some dedicated argument for that.
		if logger.Enabled["serial"] {
			// Print characters for now to print GB ROM test results.
			fmt.Printf("SC=%#02x\n", value)
		}
		if value&(1<<7) != 0 {
			// Start transfer. TODO: master/slave bit.
			s.SC |= 0x80
			s.enabled = true
			s.bit = 0
		} else {
			s.SC &= 0x7f
			s.enabled = false
		}

		if value&1 != 0 {
			s.master = true
			s.SC |= 0x01
		}
	}
}
