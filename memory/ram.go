package memory

import "math/rand"

// RAM as an arbitrary long list of R/W bytes at addresses starting from a given offset.
type RAM struct {
	Bytes []uint8
	Start uint
}

// NewRAM instantiates a zeroed slice of the given size to represent RAM.
func NewRAM(start, size uint) *RAM {
	return &RAM{make([]uint8, size), start}
}

func (r RAM) Read(addr uint) uint8 {
	return r.Bytes[addr-r.Start]
}

func (r RAM) Write(addr uint, value uint8) {
	r.Bytes[addr-r.Start] = value
}

// Contains indicates true as long as address fits in the slice.
func (r RAM) Contains(addr uint) bool {
	return addr >= r.Start && addr < r.Start+uint(len(r.Bytes))
}

// NewVRAM instantiates a slice of the given size to represent RAM, initialized
// with random values.
func NewVRAM(start, size uint) *RAM {
	vram := NewRAM(start, size)
	for i := range vram.Bytes {
		vram.Bytes[i] = uint8(rand.Intn(0xff))
	}
	return vram
}
