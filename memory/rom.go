package memory

import (
	"fmt"
	"io/ioutil"

	"go.tigris.fr/gameboy/logger"
)

// ROM is a read-only special case of RAM, initialized from a binary file.
type ROM struct {
	RAM
}

// NewROM instantiates a read-only chunk of memory from a binary dump.
func NewROM(filename string, start uint16) *ROM {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf(" !!! Cannot read ROM file %s (%s)\n", filename, err))
	}
	return &ROM{RAM{Start: start, Bytes: bytes}}
}

// Write does nothing and displays an error, for reasons I hope are obvious.
func (r *ROM) Write(addr uint16, value uint8) {
	logger.Printf("rom", "Attempt to write %x to read-only address space at %#x",
		value, addr)
}

// Internal read that doesn't conform to the Adressable interface, used for
// ROMs with memory controllers, which can then have a size well over 0xffff.
func (r *ROM) read(addr uint) uint8 {
	offset := addr - uint(r.Start)
	if offset > uint(len(r.Bytes)) {
		logger.Printf("rom", "Read overflow at %#x", addr)
		return 0xff
	}
	return r.Bytes[offset]
}
