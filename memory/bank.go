package memory

import "fmt"

// Memory banking for larger ROMs or cartridges containing more than 8Kb RAM.

// Initialize sub-logger for banked memory accesses.
func init() {
	log.Add("bank", "Read/writes in banked address spaces (Desperate level only)")
}

const (
	RAMBankSize = 0x2000
	ROMBankSize = 0x4000
)

// Bank represents an addressable memory space that can be backed by a larger
// byte array.
type Bank struct {
	memory []uint8
	offset int

	// For debugging only.
	bank      int    // Current bank selected
	addrRange string // Address range for this bank.

	Start uint16
	Size  uint16
}

// NewBank returns a Bank instance representing a memory area starting at address
// `start` and of length `size`. This area is backed by the `memory` array, which
// must be a multiple of `size`. The exposed area will be determined by the given
// `bank` number.
func NewBank(start, size uint16, bank int, memory []uint8) *Bank {
	b := Bank{
		memory:    memory,
		addrRange: fmt.Sprintf("[%04x-%04x]", start, start+size-1),
		Start:     start,
		Size:      size,
	}
	b.Set(bank)
	return &b
}

// Contains returns whether the requested address is managed by this bank.
func (b *Bank) Contains(addr uint16) bool {
	return addr >= b.Start && addr < b.Start+b.Size
}

// Read returns the byte at the given address, taking the current bank number
// into account.
func (b *Bank) Read(addr uint16) (value uint8) {
	real := int(addr-b.Start) + b.offset
	value = b.memory[real]

	log.Sub("bank").Desperatef("%s Read at 0x%04x in bank %d (real 0x%x): %d (%#x)",
		b.addrRange, addr, b.bank, real, value, value)

	return value
}

// Write modifies the byte at the given address with the given value, taking the
// current bank number into account.
func (b *Bank) Write(addr uint16, value uint8) {
	real := int(addr-b.Start) + b.offset

	log.Sub("bank").Desperatef("%s Write at 0x%04x in bank %d (real 0x%x): %d (%#x)",
		b.addrRange, addr, b.bank, real, value, value)

	b.memory[real] = value
}

// Set modifies the current bank number and recomputes the offset that will be
// used to read or write values from the underlying memory array.
func (b *Bank) Set(bank int) {
	b.bank = bank
	b.offset = (bank * int(b.Size)) % len(b.memory)
}

// BankedRAM wraps an array of R/W bytes at addresses starting from a given
// offset. The array can be larger than 8kb and accessed through banking.
type BankedRAM struct {
	*RAM

	Bank uint // 8Kb bank to map to 0xa000-0xbfff.
}

// NewBankedRAM instantiates a zeroed slice of the given size to represent RAM
// along with a bank index to map an 8Kb bank to the 0xa000-0xbfff range.
func NewBankedRAM(start, size uint16) *BankedRAM {
	return &BankedRAM{RAM: NewRAM(start, size)}
}

// offset returns the offset in our internal buffer for the requested address.
func (r *BankedRAM) offset(addr uint16) uint {
	offset := uint(addr-r.Start) + (r.Bank * RAMBankSize)
	log.Sub("bank").Desperatef("RAM [%04x-%04x] offset at 0x%04x (bank %d): %d (%#x)",
		r.Start, r.Start+RAMBankSize, addr, r.Bank, offset, offset)
	return offset
}

// Read returns the value stored at the given address in RAM, handling offsets
// and bank switching.
func (r *BankedRAM) Read(addr uint16) uint8 {
	return r.Bytes[r.offset(addr)]
}

// Write sets stores the given value at the given address in RAM, handling
// offsets and bank switching.
func (r *BankedRAM) Write(addr uint16, value uint8) {
	r.Bytes[r.offset(addr)] = value
}

// ROM as an array of R/W bytes at addresses starting from a given offset.
// If dealing with a battery-backed ROM, this can also be saved to a file.
type BankedROM struct {
	*ROM

	Bank0 uint // 16Kb bank to map to 0x0000-0x3fff.
	Bank1 uint // 16Kb bank to map to 0x4000-0x7fff
}

// NewBankedROM instantiates a zeroed slice of the given size to represent ROM
// along with bank indexes to map 16Kb bank to the 0x0000-0x3fff and
// 0x4000-0x7fff ranges.
func NewBankedROM(data []byte) *BankedROM {
	return &BankedROM{ROM: NewROM(data), Bank1: 1}
}

// offset returns the offset in our internal buffer for the requested address.
func (r *BankedROM) offset(addr uint16) uint {
	var offset uint
	switch {
	case addr <= 0x3fff:
		offset = uint(addr) + (r.Bank0 * ROMBankSize)
		log.Sub("bank").Desperatef("ROM [0000-3fff] offset at 0x%04x (bank %d): %d (%#x)",
			addr, r.Bank0, offset, offset)

	case addr >= 0x4000 && addr <= 0x7fff:
		offset = uint(addr-0x4000) + (r.Bank1 * ROMBankSize)
		log.Sub("bank").Desperatef("ROM [4000-7fff] offset at 0x%04x (bank %d): %d (%#x)",
			addr, r.Bank1, offset, offset)

	default:
		panic("out of bounds address")
	}
	return offset
}

// Read returns the value stored at the given address in ROM, handling offsets
// and bank switching.
func (r *BankedROM) Read(addr uint16) uint8 {
	return r.Bytes[r.offset(addr)]
}

// Write sets stores the given value at the given address in ROM, handling
// offsets and bank switching.
func (r *BankedROM) Write(addr uint16, value uint8) {
	r.Bytes[r.offset(addr)] = value
}
