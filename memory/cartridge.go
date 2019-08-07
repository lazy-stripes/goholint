package memory

import (
	"go.tigris.fr/gameboy/memory/chips"
)

// NewCartridge instantiates the proper kind of adress space depending on the
// given ROM's header.
// TODO: we only handle ROM-only and MBC1 so far.
func NewCartridge(romPath string) (cart Addressable) {
	log := log.Sub("cartridge") // Override default logger
	if romPath == "" {
		log.Warning("No cartridge loaded.")
		return nil
	}

	rom := NewROM(romPath, 0) // XXX: do we actually ever need to specify start > 0?

	// Check what kind of chip is in the ROM, return the proper struct.
	log.Infof("Cartridge type 0x%02x", rom.Read(0x0147))
	log.Infof("ROM size type 0x%02x", rom.Read(0x0148))
	log.Infof("RAM size type 0x%02x", rom.Read(0x0149))
	romBanks := chips.ROMBanks[rom.Read(0x0148)]
	ramBanks := chips.RAMBanks[rom.Read(0x0149)]
	switch chip := rom.Read(0x0147); chip {
	case chips.ROMOnly:
		cart = rom
	case chips.MBC1:
		cart = NewMBC1(rom, romBanks, 0, false)
	case chips.MBC1RAM:
		cart = NewMBC1(rom, romBanks, ramBanks, false)
	case chips.MBC1RAMBattery:
		cart = NewMBC1(rom, romBanks, ramBanks, true)
	default:
		log.Warningf("Unknown cartridge type 0x%02x", chip)
		cart = rom
	}

	return cart
}
