package memory

// Memory banking for larger ROMs or cartridges containing more than 8Kb RAM.

// Initialize sub-logger for banked memory accesses.
func init() {
	log.Add("bank", "Read/writes in banked address spaces (Desperate level only)")
}

const (
	RAMBankSize = 0x2000
	ROMBankSize = 0x4000
)

// RAM as an array of R/W bytes at addresses starting from a given offset.
// If dealing with a battery-backed RAM, this can also be saved to a file.
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
