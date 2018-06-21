package memory

import (
	"fmt"
)

// MMU interface provides get/set byte functions in a given address space.
type MMU interface {
	Read(addr uint) uint8
	Write(addr uint, value uint8)
}

// Boot address space translating memory access to Boot ROM for the lowest 256 bytes.
type Boot struct {
	BootROM MMU
	RAM     MMU
}

// NewBoot allocates new bootstrap memory from BootROM and Card ROM.
func NewBoot(rom, cart AddressSpace) *Boot {
	return &Boot{rom, cart}
}

func (m *Boot) Read(addr uint) uint8 {
	if addr < 0x100 {
		return m.BootROM.Read(addr)
	}
	return m.RAM.Read(addr)
}

func (m *Boot) Write(addr uint, value uint8) {
	if addr < 0x100 {
		fmt.Printf(" !!! Attempt to write to BootROM at %x\n", addr)
		return
	}
	m.RAM.Write(addr, value)
}

// AddressSpace as an arbitrary long list of bytes and some way to get them or write them.
type AddressSpace []uint8

func (r AddressSpace) Read(addr uint) uint8 {
	return r[addr]
}

func (r AddressSpace) Write(addr uint, value uint8) {
	r[addr] = value
}
