package memory

// Registers represented as an address space.
type Registers map[uint16]*uint8

// Contains returns true if the address corresponds to a register.
func (r Registers) Contains(addr uint16) bool {
	return r[addr] != nil
}

// Read returns the byte at the given address in VRAM or from register.
func (r Registers) Read(addr uint16) uint8 {
	if regPtr := r[addr]; regPtr != nil {
		return *regPtr
	}
	log.Warningf("Reading unknown register address %#4x", addr)
	return 0xff
}

// Write sets the byte at the given address in VRAM to the given value. TODO: checks
func (r Registers) Write(addr uint16, value uint8) {
	if regPtr := r[addr]; regPtr != nil {
		*regPtr = value
	} else {
		log.Warningf("Writing to unknown register address %#4x", addr)
	}
}
