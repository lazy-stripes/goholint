package memory

import (
	"errors"
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
	data     []uint8
	saveFile string
}

// Load sets the current memory contents from the given file, and stores the
// path to that file for subsequent saves.
func (m *Memory) Load(filename string) error {
	if len(m.saveFile) > 0 && m.saveFile != filename {
		log.Warningf("calling Load(%s) on memory with an existing save file (%s)",
			filename, m.saveFile)
	}
	oldSavefile := m.saveFile
	m.saveFile = filename

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("couldn't load file %s in memory (%w)", filename, err)
	}

	// Token attempt at checking file validity.
	if len(bytes) != len(m.data) {
		// In this specific case, reset saveFile because this may well be an
		// innocent mistake and we don't want that file overwritten later.
		m.saveFile = oldSavefile

		return fmt.Errorf("memory size (0x%04x) does not match file size (0x%04x)",
			len(m.data), len(bytes))
	}

	// Replace current (normally empty) memory with file contents.
	m.data = bytes
	log.Infof("loading memory values from %s", filename)

	return nil
}

// Save dumps the current memory contents into the associated save file (if any).
func (m *Memory) Save() error {
	if m.saveFile == "" {
		return errors.New("trying to Save() memory with no save file defined")
	}

	return os.WriteFile(m.saveFile, m.data, 0644)
}
