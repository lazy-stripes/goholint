package memory

// AddressSpace interface provides functions to read/write bytes in a given address space.
type AddressSpace interface {
	// Contains returns true if the given address belongs to the address space, false otherwise.
	Contains(addr uint) bool
	// Read returns the value stored at the given address.
	Read(addr uint) uint8
	// Write attempts to store the given value at the given address. Not all address spaces are writable.
	Write(addr uint, value uint8)
}
