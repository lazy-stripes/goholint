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
		fmt.Printf(" !!! Cannot read ROM file %s (%s)\n", filename, err)
		return nil
	}
	return &ROM{RAM{Start: start, Bytes: bytes}}
}

// Write does nothing and displays an error, for reasons I hope are obvious.
func (r *ROM) Write(addr uint16, value uint8) {
	logger.Printf("rom", "Attempt to write %x to read-only address space at %#x",
		value, addr)
}
