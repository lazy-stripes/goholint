package memory

import (
	"fmt"
)

// IORegister represents a register value with an associated hook to call on Write.
type IORegister struct {
	Register  *uint8
	WriteHook func(value uint8)
}

// IORegisters represented as an address space mapping to a memory location and a write hook.
type IORegisters map[uint]IORegister

// Contains returns true if the address corresponds to a register.
func (r IORegisters) Contains(addr uint) (present bool) {
	_, present = r[addr]
	return
}

// Read returns the byte at the given address in VRAM corresponding to a register.
func (r IORegisters) Read(addr uint) uint8 {
	if io, present := r[addr]; present {
		return *io.Register
	}
	fmt.Printf("Reading unknown I/O register address %#4x\n", addr)
	return 0xff
}

// Write sets the byte at the given address in VRAM to the given value. TODO: checks
func (r IORegisters) Write(addr uint, value uint8) {
	// Set register by default, unless hook exists. Used to implement R/O registers.
	if io, present := r[addr]; present {
		if io.WriteHook != nil {
			io.WriteHook(value)
		} else {
			*io.Register = value
		}
	} else {
		fmt.Printf("Writing to unknown I/O register address %#4x\n", addr)
	}
}
