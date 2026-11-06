package memory

import (
	"fmt"
	"os"

	"github.com/lazy-stripes/goholint/logger"
)

// Package-wide logger initialized below. Sub-loggers will be created in the
// relevant source files.
var log = logger.New("memory", "memory-related operations")

// Addressable interface provides functions to read/write bytes in a given
// 16-bit address space.
type Addressable interface {
	// Contains returns true if the given address belongs to the address space.
	Contains(addr uint16) bool
	// Read returns the value stored at the given address.
	Read(addr uint16) uint8
	// Write attempts to store the given value at the given address if writable.
	Write(addr uint16, value uint8)
}

type Memory struct {
	Bytes []uint8
	Start uint16
}

// New instantiates a zeroed slice of the given size to represent memory.
func New(start uint16, size int) *Memory {
	return &Memory{
		Bytes: make([]uint8, size),
		Start: start,
	}
}

// Read returns the value stored at the given address in RAM, handling offsets.
func (m *Memory) Read(addr uint16) uint8 {
	return m.Bytes[addr-m.Start]
}

// Write sets stores the given value at the given address in RAM, handling offsets.
func (m *Memory) Write(addr uint16, value uint8) {
	m.Bytes[addr-m.Start] = value
}

// Contains indicates true as long as address fits in the slice. Careful not
// to wrap uint16 here.
func (m *Memory) Contains(addr uint16) bool {
	return addr >= m.Start && addr < m.Start+uint16(len(m.Bytes))
}

// Load sets the current memory contents from the given file, and stores the
// path to that file for subsequent saves.
func (m *Memory) Load(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("couldn't load file %s in memory (%w)", filename, err)
	}

	// Token attempt at checking file validity. Allow replacing an empty data
	// array (e.g. for one-time loading of ROMs).
	if len(m.Bytes) > 0 && len(bytes) != len(m.Bytes) {
		return fmt.Errorf("memory size (0x%04x) does not match file size (0x%04x)",
			len(m.Bytes), len(bytes))
	}

	// Replace current (normally empty) memory with file contents.
	m.Bytes = bytes
	log.Infof("loaded memory values from %s", filename)

	return nil
}
