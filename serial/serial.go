package serial

import (
	"fmt"

	"go.tigris.fr/gameboy/memory"
)

// Register addresses
const (
	AddrSB = 0xff01
	AddrSC = 0xff02
)

// Serial registers for game link. Used only for debug for now.
type Serial struct {
	memory.Registers

	SB, SC uint8
}

// New instantiates a Serial addressable mapping to FF01 and FF02.
func New() *Serial {
	s := Serial{Registers: memory.Registers{}}
	s.Registers[0xff01] = &s.SB
	s.Registers[0xff02] = &s.SC
	return &s
}

// Override write for Serial Transfer Control
func (s *Serial) Write(addr uint, value uint8) {
	s.Registers.Write(addr, value)
	if addr == AddrSC {
		if value&(1<<7) > 0 {
			fmt.Printf("%c", s.SB)
		}
	}
}
