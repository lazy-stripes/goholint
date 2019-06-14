package memory

// IORegister represents a register value with optional hooks to call on read or
// write as well as an optional mask for registers not exploiting all 8 bits.
type IORegister struct {
	Register  *uint8
	ReadMask  uint8
	WriteMask uint8
	ReadHook  func(io *IORegister) uint8
	WriteHook func(io *IORegister, value uint8)
}

// IORegisters represented as an address space mapping to a memory location and a write hook.
type IORegisters map[uint16]*IORegister

// Contains returns true if the address corresponds to a register.
func (r IORegisters) Contains(addr uint16) (present bool) {
	_, present = r[addr]
	return
}

// Read returns the byte at the given address in VRAM corresponding to a register.
func (r IORegisters) Read(addr uint16) uint8 {
	if io, present := r[addr]; present {
		if io.ReadHook != nil {
			return io.ReadHook(io)
		}
		return *io.Register & io.ReadMask
	}
	panic("Broken MMU")
}

// Write sets the byte at the given address in VRAM to the given value. TODO: checks
func (r IORegisters) Write(addr uint16, value uint8) {
	// Set register by default, unless hook exists. Used to implement R/O registers.
	if io, present := r[addr]; present {
		if io.WriteHook != nil {
			io.WriteHook(io, value)
		} else {
			*io.Register = (*io.Register & ^io.WriteMask) | (value & io.WriteMask)
		}
	}
	panic("Broken MMU")
}

// Helper functions below to create various types of registers (RW, RO,
// restricted write...)

func readOnly(io *IORegister, value uint8) {
	log.Debugf("Read-only register, value 0x%02x not written.", value)
}

// NewRWRegister creates an IORegister pointing to a given byte with full R/W
// access, pretty much identical to a regular register.
// TODO: This is growing messy, I'll want to keep a single register type in the end.
func NewRWRegister(register *uint8) *IORegister {
	return &IORegister{Register: register, ReadMask: 0xff, WriteMask: 0xff}
}

// NewRORegister creates an IORegister pointing to a given byte with read-only
// access.
func NewRORegister(register *uint8) *IORegister {
	return &IORegister{Register: register, ReadMask: 0xff, WriteHook: readOnly}
}
