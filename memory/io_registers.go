package memory

import (
	"fmt"
)

type IORegister struct {
	register  *uint8
	writeHook func(value uint8)
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
		return *io.register
	}
	fmt.Printf("Reading unknown I/O register address %#4x\n", addr)
	return 0xff
}

// Write sets the byte at the given address in VRAM to the given value. TODO: checks
func (r IORegisters) Write(addr uint, value uint8) {
	// FIXME: check for R/O registers.
	if io, present := r[addr]; present {
		*io.register = value
		if io.writeHook != nil {
			io.writeHook(value)
		}
	} else {
		fmt.Printf("Writing to unknown I/O register address %#4x\n", addr)
	}
}
