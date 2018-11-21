package serial

import (
	"fmt"
)

// Register addresses
const (
	AddrSB = 0xff01
	AddrSC = 0xff02
)

// Serial registers for game link. Used only for debug for now.
type Serial struct {
	SB, SC uint8
}

// New instantiates a Serial addressable mapping to FF01 and FF02.
func New() *Serial {
	return &Serial{}
}

func (s *Serial) Contains(addr uint) bool {
	return addr == AddrSB || addr == AddrSC
}

func (s *Serial) Read(addr uint) uint8 {
	if addr == AddrSB {
		return s.SB
	} else if addr == AddrSC {
		return s.SC | 0x7e
	} else {
		panic(fmt.Sprintf("broken MMU: illegal address %04x requested!", addr))
	}
}

func (s *Serial) Write(addr uint, value uint8) {
	if addr == AddrSB {
		s.SB = value
	} else if addr == AddrSC {
		s.SC = value
		if value&(1<<7) != 0 {
			fmt.Printf("%c", s.SB)
		}
	}
}
