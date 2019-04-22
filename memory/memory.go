package memory

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
