package memory

import "fmt"

// Registers represented as an address space.
type Registers map[uint]*uint8

// Contains returns true if the address corresponds to a register.
func (r Registers) Contains(addr uint) bool {
	return r[addr] != nil
}

// Read returns the byte at the given address in VRAM or from register.
func (r Registers) Read(addr uint) uint8 {
	if regPtr := r[addr]; regPtr != nil {
		return *regPtr
	}
	fmt.Printf("Reading unknown register address %#4x\n", addr)
	return 0xff
}

// Write sets the byte at the given address in VRAM to the given value. TODO: checks
func (r Registers) Write(addr uint, value uint8) {
	// FIXME: check for R/O registers.
	if regPtr := r[addr]; regPtr != nil {
		*regPtr = value
	} else {
		fmt.Printf("Writing to unknown register address %#4x\n", addr)
	}
}
