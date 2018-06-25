package memory

import (
	"fmt"
	"io/ioutil"
)

// AddressSpace interface provides functions to read/write bytes in a given address space.
type AddressSpace interface {
	// Contains returns true if the given address belongs to the address space, false otherwise.
	Contains(addr uint) bool
	// Read returns the value stored at the given address.
	Read(addr uint) uint8
	// Write attempts to store the given value at the given address. Not all address spaces are writable.
	Write(addr uint, value uint8)
}

// MMU manages an arbitrary number of ordered address spaces, starting with the DMG boot ROM by default.
// It also satisfies the AddressSpace interface.
type MMU struct {
	Spaces []AddressSpace
}

// NewMMU returns an instance of MMU initialized with optional address spaces.
func NewMMU(spaces []AddressSpace) *MMU {
	return &MMU{spaces}
}

// Contains returns whether one of the address spaces known to the MMU contains the given address. The first
// address space in the internal list containing a given address will shadow any other.
func (m *MMU) Contains(addr uint) bool {
	for _, space := range m.Spaces {
		if space.Contains(addr) {
			return true
		}
	}
	return false
}

// Returns the first space for which the address is handled.
func (m *MMU) space(addr uint) AddressSpace {
	for _, space := range m.Spaces {
		if space.Contains(addr) {
			return space
		}
	}
	return nil // TODO: VOID
}

// Read finds the first address space compatible with the given address and returns the value at that address.
func (m *MMU) Read(addr uint) uint8 {
	if space := m.space(addr); space != nil {
		return space.Read(addr)
	}
	return 0xFF
}

// Write finds the first address space compatible with the given address and attempts writing the given value to that
// address. TODO: error handling for write only?
func (m *MMU) Write(addr uint, value uint8) {
	if space := m.space(addr); space != nil {
		space.Write(addr, value)
	}
}

// Boot address space translating memory access to Boot ROM for the lowest 256 bytes.
type Boot struct {
	BootROM AddressSpace
	RAM     AddressSpace
}

// BootROM is the image of the DMG boot ROM, initialized from dum file.
type BootROM []byte

// NewBootROM instantiates a copy of a DMG ROM dump.
func NewBootROM(filename string) BootROM {
	rom, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf(" !!! Invalid ROM path %s (%s)\n", filename, err)
		return nil
	}
	return rom
}

func (b BootROM) Read(addr uint) uint8 {
	return b[addr]
}

func (b BootROM) Write(addr uint, value uint8) {
	fmt.Printf(" !!! Attempt to write %x to BootROM at %x\n", value, addr)
}

// Contains indicates true as long as address fits in the slice.
func (b BootROM) Contains(addr uint) bool {
	return addr >= 0 && addr < uint(len(b))
}

// RAM as an arbitrary long list of R/W bytes.
type RAM []uint8

// NewRAM instantiates a zeroed slice of the given size to represent RAM.
func NewRAM(size uint) RAM {
	return make(RAM, size)
}

func (r RAM) Read(addr uint) uint8 {
	return r[addr]
}

func (r RAM) Write(addr uint, value uint8) {
	r[addr] = value
}

// Contains indicates true as long as address fits in the slice.
func (r RAM) Contains(addr uint) bool {
	return addr >= 0 && addr < uint(len(r))
}
