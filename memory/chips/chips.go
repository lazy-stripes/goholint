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
