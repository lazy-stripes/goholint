package memory

import (
	"go.tigris.fr/gameboy/logger"
	"go.tigris.fr/gameboy/memory/chips"
)

// NewCartridge instantiates the proper kind of adress space depending on the
// given ROM's header.
// TODO: we only handle ROM-only and MBC1 so far.
func NewCartridge(romPath string) (cart Addressable) {
	if romPath == "" {
		logger.Printf("cartridge", "No cartridge loaded.")
		return NewRAM(0, 0)
	}

	rom := NewROM(romPath, 0) // XXX: do we actually ever need to specify start > 0?

	// Check what kind of chip is in the ROM, return the proper struct.
	logger.Printf("cartridge", "Cartridge type 0x%02x", rom.Read(0x0147))
	switch chip := rom.Read(0x0147); chip {
	case chips.ROMOnly:
		cart = rom
	case chips.MBC1:
		cart = NewMBC1(rom)
	case chips.MBC1RAM:
		cart = NewMBC1(rom)
	case chips.MBC1RAMBattery:
		cart = NewMBC1(rom)
	default:
		logger.Printf("cartridge", "Unknown cartridge type 0x%02x", chip)
		cart = rom
	}

	return cart
}
