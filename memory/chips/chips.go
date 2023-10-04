package chips

// Cartridge chip types as found in ROM header at address 0x0147.
const (
	ROMOnly        uint8 = 0x00
	MBC1           uint8 = 0x01
	MBC1RAM        uint8 = 0x02
	MBC1RAMBattery uint8 = 0x03
	MBC2           uint8 = 0x05
	MBC2Battery    uint8 = 0x06
	ROMRAM         uint8 = 0x08
	ROMRAMBattery  uint8 = 0x09
	// TODO: all others that are not strictly used with CGB.
)

var Names = map[uint8]string{
	ROMOnly:        "ROM only",
	MBC1:           "MBC1",
	MBC1RAM:        "MBC1+RAM",
	MBC1RAMBattery: "MBC1+RAM+Battery",
	MBC2:           "MBC2",
	MBC2Battery:    "MBC2+Battery",
	ROMRAM:         "ROM+RAM",
	ROMRAMBattery:  "ROM+RAM+Battery",
}

// ROMBanks number depending on "ROM Size" cartridge header.
var ROMBanks = map[uint8]uint16{
	0x00: 2,
	0x01: 4,
	0x02: 8,
	0x03: 16,
	0x04: 32,
	0x05: 64,
	0x06: 128,
	0x07: 256,
	0x08: 512,
	0x52: 72, // Unconfirmed.
	0x53: 80, // Unconfirmed.
	0x54: 96, // Unconfirmed.
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
