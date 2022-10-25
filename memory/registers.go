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

// HookRegister represents a register value with optional hooks to call on read
// or write. This is mostly used for the APU's registers.
type HookRegister struct {
	Ptr   *uint8 // Address of actual register variable to read/write.
	Read  func() uint8
	Write func(value uint8)
}

// HookRegisters hold a mapping of addresses to registers with an optional hook
// called at read/write time.
type HookRegisters map[uint16]*HookRegister

// Contains returns true if the address corresponds to a register.
func (r HookRegisters) Contains(addr uint16) (present bool) {
	_, present = r[addr]
	return
}

// Read returns the byte at the given address in VRAM corresponding to a register.
func (r HookRegisters) Read(addr uint16) uint8 {
	if reg, present := r[addr]; present {
		if reg.Read != nil {
			return reg.Read()
		}
		return *reg.Ptr
	}
	panic("Broken MMU")
}

// Write sets the byte at the given address in VRAM to the given value. TODO: checks
func (r HookRegisters) Write(addr uint16, value uint8) {
	// Set register by default, unless hook exists. Used to implement R/O registers.
	if reg, present := r[addr]; present {
		// Write the value no matter what, and then call the hook function that
		// may do something with that newly written value.
		*reg.Ptr = value
		if reg.Write != nil {
			reg.Write(value)
		}
	} else {
		log.Warningf("Writing to unknown register address %#4x", addr)
	}
}
