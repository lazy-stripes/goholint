package apu

// APURegister wraps a hardware register and associates it with a read mask as
// well as an optional function to call when the register is written (to trigger
// a channel, recompute a frequency, etc).
type APURegister struct {
	Ptr     *uint8            // Pointer to the actual register variable.
	Mask    uint8             // Value to OR with the register's current value on read.
	OnWrite func(value uint8) // Function to call on write, if not nil.
}

// APURegisters works like memory.Registers and treats a mapping of addresses
// to APURegisters as an address space.
type APURegisters map[uint16]APURegister

// Contains returns true if the requested address is one of an APU register.
func (a APURegisters) Contains(addr uint16) (present bool) {
	_, present = a[addr]
	return
}

// Read returns the value of the APU register at the given address, OR'ed with
// this register's specific mask.
// See https://gbdev.gg8.se/wiki/articles/Gameboy_sound_hardware#Register_Reading
func (a APURegisters) Read(addr uint16) uint8 {
	if reg, ok := a[addr]; ok {
		return *reg.Ptr | reg.Mask
	}
	log.Warningf("Reading unknown APU register address %#4x", addr)
	return 0xff
}

// Write stores the given value in the APU register at the given address. If
// there is a write function for this register, it is called after the value is
// written.
func (a APURegisters) Write(addr uint16, value uint8) {
	if reg, ok := a[addr]; ok {
		*reg.Ptr = value
		if reg.OnWrite != nil {
			reg.OnWrite(value)
		}
	} else {
		log.Warningf("Writing to unknown APU register address %#4x", addr)
	}
}
