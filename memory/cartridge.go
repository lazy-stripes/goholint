package memory

import (
	"github.com/lazy-stripes/goholint/memory/chips"
)

// Initialize sub-logger for cartridge details.
func init() {
	log.Add("cartridge", "cartridge address space details")
}

// NewCartridge instantiates the proper kind of adress space depending on the
// given ROM's header.
// TODO: we only handle ROM-only and MBC1 so far.
func NewCartridge(romPath, savePath string) (cart Addressable) {
	log := log.Sub("cartridge") // Override default logger
	if romPath == "" {
		log.Warning("No cartridge loaded.")
		return nil
	}

	rom := NewROMFromFile(romPath)

	// Check what kind of chip is in the ROM, return the proper struct.
	cartType := rom.Read(0x0147)
	romSize := rom.Read(0x0148)
	ramSize := rom.Read(0x0149)

	romBanks := chips.ROMBanks[romSize]
	ramBanks := chips.RAMBanks[ramSize]

	log.Infof("Cartridge type 0x%02x (%s)", cartType, chips.Names[cartType])
	log.Infof("ROM size type 0x%02x (%d banks)", romSize, romBanks)
	log.Infof("RAM size type 0x%02x (%d banks)", ramSize, ramBanks)

	switch cartType {
	case chips.ROMOnly:
		cart = rom
	case chips.MBC1:
		cart = NewMBC1(rom, uint8(romBanks), 0, false, "")
	case chips.MBC1RAM:
		cart = NewMBC1(rom, uint8(romBanks), ramBanks, false, "")
	case chips.MBC1RAMBattery:
		cart = NewMBC1(rom, uint8(romBanks), ramBanks, true, savePath)
	default:
		log.Warningf("Unknown cartridge type 0x%02x", cartType)
		cart = rom
	}

	return cart
}
