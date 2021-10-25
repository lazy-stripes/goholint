package memory

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// RAM as an array of R/W bytes at addresses starting from a given offset.
// If dealing with a batter-backed RAM, this can also be saved to a file.
type RAM struct {
	Bytes []uint8
	Start uint16

	saveFile string // For batter-backed RAM chips
}

// NewRAM instantiates a zeroed slice of the given size to represent RAM.
func NewRAM(start, size uint16) *RAM {
	return &RAM{make([]uint8, size), start, ""}
}

// Read returns the value stored at the given address in RAM, handling offsets.
func (r *RAM) Read(addr uint16) uint8 {
	return r.Bytes[addr-r.Start]
}

// Read sets stores the given value at the given address in RAM, handling offsets.
func (r *RAM) Write(addr uint16, value uint8) {
	r.Bytes[addr-r.Start] = value
}

// Contains indicates true as long as address fits in the slice. Careful not
// to wrap uint16 here.
func (r *RAM) Contains(addr uint16) bool {
	return addr >= r.Start && addr <= r.Start+uint16(len(r.Bytes))-1
}

// Load sets the current content of RAM from the given file, and stores the
// path to that file for subsequent saves.
func (r *RAM) Load(filename string) error {
	if r.saveFile != filename {
		log.Warningf("calling Load(%s) on RAM with an existing save file (%s)",
			filename, r.saveFile)
	}
	oldSavefile := r.saveFile
	r.saveFile = filename

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		// The file may not yet exist, it's fine.
		return fmt.Errorf("cannot load RAM file %s (%s)", filename, err)
	}

	// Token attempt at checking file validity.
	if len(bytes) != len(r.Bytes) {
		// In this specific case, reset saveFile because this may well be an
		// innocent mistake and we don't want that file overwritten later.
		r.saveFile = oldSavefile

		return fmt.Errorf("RAM size (0x%04x) does not match save file's (0x%04x)",
			len(r.Bytes), len(bytes))
	}

	// Replace current (normally empty) RAM with file contents.
	r.Bytes = bytes
	log.Infof("loading RAM values from %s", filename)

	return nil
}

// Save dumps the current content of RAM into the associated save file (if any).
func (r *RAM) Save() error {
	if r.saveFile == "" {
		return errors.New("trying to Save() RAM with no save file defined")
	}

	return ioutil.WriteFile(r.saveFile, r.Bytes, 0644)
}
