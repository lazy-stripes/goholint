package memory

import "github.com/lazy-stripes/goholint/logger"

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

// Package-wide logger initialized below.
var log = logger.New("memory", "memory-related operations")

// Package initialization function setting up logger submodules.
func init() {
	log.Add("boot", "boot ROM disable register")
	log.Add("cartridge", "cartridge address space details")
	log.Add("dma", "DMA register and transfers")
	log.Add("mbc", "MBC chip details")
	log.Add("mbc/read", "MBC ROM chip reads (Desperate level only)")
	log.Add("mbc/write", "MBC ROM chip writes (Desperate level only)")
	log.Add("mmu/read", "unmapped memory reads (Debug level only)")
	log.Add("mmu/write", "unmapped memory writes (Debug level only)")
	log.Add("rom", "Unexpected writes and overflows in ROM address spaces")
}
