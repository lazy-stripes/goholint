package memory

import (
	"errors"
	"io/ioutil"
	"math/rand"
)

// RAM as an array of R/W bytes at addresses starting from a given offset.
// If dealing with a batter-backed RAM, this can also be saved to a file.
type RAM struct {
	Bytes []uint8
	Start uint16

	saveFile string // For batter-backed RAM chips
}

// NewRAM instantiates a RAM addressable initialized from a save file. The
// file's size must match the RAM's exactly.
func NewRAM(start, size uint16, filename string) *RAM {
	ram := &RAM{make([]uint8, size), start, filename}

	// Don't try loading save file if not needed.
	if filename == "" {
		return ram
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		// The file may not yet exist, it's fine.
		log.Warningf("cannot read saved RAM file %s (%s)", filename, err)
		return ram
	}

	// Token attempt at checking file validity.
	if len(bytes) != int(size) {
		log.Warningf("RAM size (0x%04x) does not match save file's (0x%04x)",
			size, len(bytes))
		return ram
	}

	// Replace empty RAM with file contents.
	ram.Bytes = bytes

	return &RAM{bytes, start, filename}
}

// NewEmptyRAM instantiates a zeroed slice of the given size to represent RAM.
func NewEmptyRAM(start, size uint16) *RAM {
	return &RAM{make([]uint8, size), start, ""}
}

func (r RAM) Read(addr uint16) uint8 {
	return r.Bytes[addr-r.Start]
}

func (r RAM) Write(addr uint16, value uint8) {
	r.Bytes[addr-r.Start] = value
}

// Contains indicates true as long as address fits in the slice. Careful not
// to wrap uint16 here.
func (r RAM) Contains(addr uint16) bool {
	return addr >= r.Start && addr <= r.Start+uint16(len(r.Bytes))-1
}

// NewVRAM instantiates a slice of the given size to represent RAM, initialized
// with random values.
func NewVRAM(start, size uint16) *RAM {
	vram := NewEmptyRAM(start, size)
	for i := range vram.Bytes {
		vram.Bytes[i] = uint8(rand.Intn(0xff))
	}
	return vram
}

// Save dumps the current content of RAM into the associated save file (if any).
func (r RAM) Save() error {
	if r.saveFile == "" {
		return errors.New("trying to Save() RAM with no save file defined")
	}

	return ioutil.WriteFile(r.saveFile, r.Bytes, 0644)
}
