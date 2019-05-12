package chips

// Cartridge chip types as found in ROM header at address 0x0147.
const (
	ROMOnly        = 0x00
	MBC1           = 0x01
	MBC1RAM        = 0x02
	MBC1RAMBattery = 0x03
	MBC2           = 0x05
	MBC2Battery    = 0x06
	ROMRAM         = 0x08
	ROMRAMBattery  = 0x09
	// TODO: all others that are not strictly used with CGB.
)

// ROMBanks number depending on "ROM Size" cartridge header.
var ROMBanks = map[uint8]uint16{
	0x00: 0,
	0x01: 4,
	0x02: 8,
	0x03: 16,
	0x04: 32,
	0x05: 64,
	0x06: 128,
	0x07: 256,
	0x08: 512,
	0x52: 72,
	0x53: 80,
	0x54: 96,
}

// RAMBanks number depending on "RAM Size" cartridge header.
var RAMBanks = map[uint8]uint8{
	0x00: 0,
	0x01: 1, // Technically 1/4th of a 8KB bank?
	0x02: 1,
	0x03: 4,
	0x04: 16,
	0x05: 8,
}
