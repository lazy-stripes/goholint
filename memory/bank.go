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

// Banked represents an addressable memory space that can be backed by a larger
// byte array.
type Banked struct {
	*Memory

	offset int

	// For debugging only.
	bank      int    // Current bank selected
	addrRange string // Address range for this bank.

	Start    uint16
	BankSize uint16
}

// NewBanked returns a Banked instance representing a memory area with a start
// address and a size, which is the size of the bank, not the underlying memory.
// The backing memory array must be a multiple of the given size, and its
// accessible range will be determined by the given bank number.
func NewBanked(start, bankSize uint16, bank int, memory *Memory) *Banked {
	b := Banked{
		Memory:    memory,
		Start:     start,
		BankSize:  bankSize,
		addrRange: fmt.Sprintf("[%04x-%04x]", start, start+bankSize-1),
	}
	b.Set(bank)
	return &b
}

// Contains returns whether the requested address is managed by this bank.
func (b *Banked) Contains(addr uint16) bool {
	return addr >= b.Start && addr < b.Start+b.BankSize
}

// Read returns the byte at the given address, taking the current bank number
// into account.
func (b *Banked) Read(addr uint16) (value uint8) {
	real := int(addr-b.Start) + b.offset
	value = b.Bytes[real]

	log.Sub("bank").Desperatef("%s Read at 0x%04x in bank %d (real 0x%x): %d (%#x)",
		b.addrRange, addr, b.bank, real, value, value)

	return value
}

// Write modifies the byte at the given address with the given value, taking the
// current bank number into account.
func (b *Banked) Write(addr uint16, value uint8) {
	real := int(addr-b.Start) + b.offset

	log.Sub("bank").Desperatef("%s Write at 0x%04x in bank %d (real 0x%x): %d (%#x)",
		b.addrRange, addr, b.bank, real, value, value)

	b.Bytes[real] = value
}

// Set modifies the current bank number and recomputes the offset that will be
// used to read or write values from the underlying memory array.
func (b *Banked) Set(bank int) {
	b.bank = bank
	b.offset = (bank * int(b.BankSize)) % len(b.Bytes) // TODO: subtract Start here directly?
}

// BankedRAM wraps an array of R/W bytes at addresses starting from a given
// offset. The array can be larger than 8kb and accessed through banking.
type BankedRAM struct {
	*Banked
	*RAM
}

// NewBankedRAM instantiates a zeroed slice of the given size to represent RAM
// along with a bank index to map an 8Kb bank to the 0xa000-0xbfff range.
func NewBankedRAM(start uint16, size int) *BankedRAM {
	ram := NewRAM(start, size)
	banked := NewBanked(start, RAMBankSize, 0, ram.Memory)
	return &BankedRAM{RAM: ram, Banked: banked}
}

// ROM as an array of R/W bytes at addresses starting from a given offset.
// If dealing with a battery-backed ROM, this can also be saved to a file.
type BankedROM struct {
	*ROM

	Bank0 *Banked // 16Kb bank to map to 0x0000-0x3fff.
	Bank1 *Banked // 16Kb bank to map to 0x4000-0x7fff
}

// NewBankedROM instantiates a memory array with the given data to represent a
// ROM that can be accessed through two banks in the 0x0000-0x3fff and
// 0x4000-0x7fff ranges.
func NewBankedROM(data []byte) *BankedROM {
	rom := NewROM(data)
	bank0 := NewBanked(0x0000, ROMBankSize, 0, rom.Memory)
	bank1 := NewBanked(0x4000, ROMBankSize, 1, rom.Memory)
	return &BankedROM{ROM: rom, Bank0: bank0, Bank1: bank1}
}

func (r *BankedROM) Contains(addr uint16) bool {
	return addr < 0x8000
}

// Read returns the value stored at the given address in ROM, taking banks into
// account.
func (r *BankedROM) Read(addr uint16) uint8 {
	if addr < 0x4000 {
		return r.Bank0.Read(addr)
	}
	return r.Bank1.Read(addr)
}
