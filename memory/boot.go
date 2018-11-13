package memory

import (
	"fmt"
)

// Address of BOOT register in I/O RAM.
const bootAddr = 0xff50

// Boot address space holding Boot ROM and BOOT register to disable it.
type Boot struct {
	Register uint8
	ROM      ROM
	disabled bool
}

// NewBoot returns a new Boot address space containing the given boot ROM.
func NewBoot(filename string) *Boot {
	return &Boot{ROM: *NewROM(filename, 0)}
}

// Contains returns true if the given address belongs to the ROM or BOOT register, false otherwise.
func (b *Boot) Contains(addr uint) bool {
	return addr == bootAddr || (!b.disabled && b.ROM.Contains(addr))
}

// Read returns the value stored at the given address in ROM or in BOOT register.
func (b *Boot) Read(addr uint) uint8 {
	if addr == bootAddr {
		return b.Register
	}
	return b.ROM.Read(addr)
}

// Write is only supported for BOOT register and disables the boot ROM.
func (b *Boot) Write(addr uint, value uint8) {
	if addr == bootAddr {
		b.Register = value
		b.disabled = true
		fmt.Println(" !!! BootROM disabled.")
	} else {
		b.ROM.Write(addr, value) // Shouldn't happen but will log it if it does
	}
}