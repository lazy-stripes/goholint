package memory

// Registers map arbitrary 16-bit addresses to an 8-bit value that can be read
// and written. It's very useful to expose a bunch of hardware registers that
// may not all have contiguous addresses (like, say, APU registers...)
type Registers map[uint16]*uint8

// Contains returns true if the given address maps to an existing register.
func (r Registers) Contains(addr uint16) bool {
	return r[addr] != nil
}

// Read returns the byte at the given address.
func (r Registers) Read(addr uint16) uint8 {
	if regPtr := r[addr]; regPtr != nil {
		return *regPtr
	}
	log.Warningf("Reading unknown register address %#4x", addr)
	return 0xff
}

// Write sets the byte at the given address to the given value. TODO: checks
func (r Registers) Write(addr uint16, value uint8) {
	if regPtr := r[addr]; regPtr != nil {
		*regPtr = value
	} else {
		log.Warningf("Writing to unknown register address %#4x", addr)
	}
}
