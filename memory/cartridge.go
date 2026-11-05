package memory

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lazy-stripes/goholint/memory/chips"
)

// Initialize sub-logger for cartridge details.
func init() {
	log.Add("cartridge", "cartridge address space details")
}

type Persistable interface {
	Load(path string) error
	Save() error
}

type Cartridge struct {
	bus Addressable
	ram Persistable

	cartType uint8
	romSize  uint8
	ramSize  uint8
	battery  bool
	// TODO: more details from header (game name, etc)
}

func (c *Cartridge) Bus() Addressable {
	return c.bus
}

func (c *Cartridge) HasBattery() bool {
	return c.battery
}

func (c *Cartridge) LoadRAM(path string) error {
	if !c.battery {
		return errors.New("cannot load RAM contents for a cartridge with no battery")
	}
	return c.ram.Load(path)
}

func (c *Cartridge) SaveRAM() error {
	if !c.battery {
		return errors.New("cannot save RAM contents for a cartridge with no battery")
	}
	return c.ram.Save()
}

// OpenROM returns the contents of the Game Boy ROM at the given path. The file
// can be a ZIP file, in which case the first compressed file in the archive
// with the .gb extension will be loaded.
func OpenROM(filename string) (data []byte, err error) {
	// Try decompressing ZIP first, if not, read raw data.
	// TODO: Try to see if pre-read data can be decompressed, if not use them as-is.
	archive, err := zip.OpenReader(filename)
	switch err {
	case zip.ErrFormat:
		// Not a ZIP file, treat as GB ROM directly.
		data, err = os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("reading ROM failed: %w", err)
		}
	case nil:
		// Proper ZIP file, try finding a GB ROM in there.
		for _, f := range archive.File {
			if !strings.HasSuffix(f.Name, ".gb") {
				continue
			}
			log.Debugf("Extracting %s from %s", f.Name, filename)
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err.Error())
			}
			data, err = io.ReadAll(rc)
			if err != nil {
				log.Fatal(err.Error())
			}
			rc.Close()
			break
		}
		if data == nil {
			return nil, fmt.Errorf("no GB ROM found in %s", filename)
		}
	default:
		// Improper ZIP file.
		return nil, err
	}

	return data, nil
}

// NewCartridge instantiates the proper kind of adress space depending on the
// given ROM's header.
// TODO: implement missing mappers.
func NewCartridge(romPath, savePath string) (cart *Cartridge, err error) {
	log := log.Sub("cartridge") // Override default logger
	if romPath == "" {
		return nil, errors.New("empty path given, cannot load ROM")
	}

	romData, err := OpenROM(romPath)
	if err != nil {
		return nil, fmt.Errorf("reading ROM failed: %w", err)
	}

	// Make sure we have at least enough bytes to parse the header.
	if len(romData) < 0x150 {
		return nil, fmt.Errorf("truncated ROM failed: %w", err)
	}

	// Check what kind of chip is in the ROM, select the proper struct.
	cartType := romData[0x0147]
	romSize := romData[0x0148]
	ramSize := romData[0x0149]

	romBanks := chips.ROMBanks[romSize]
	ramBanks := chips.RAMBanks[ramSize]

	log.Infof("Cartridge type 0x%02x (%s)", cartType, chips.Names[cartType])
	log.Infof("ROM size type 0x%02x (%d banks)", romSize, romBanks)
	log.Infof("RAM size type 0x%02x (%d banks)", ramSize, ramBanks)

	// TODO: validate

	switch cartType {
	case chips.ROMOnly:
		cart = rom
	case chips.MBC1:
		cart = NewMBC1(rom, uint8(romBanks), 0, false, "")
	case chips.MBC1RAM:
		cart = NewMBC1(rom, uint8(romBanks), ramBanks, false, "")
	case chips.MBC1RAMBattery:
		cart = NewMBC1(rom, uint8(romBanks), ramBanks, true, savePath)
	case chips.MBC2:
		cart = NewMBC2(rom, uint8(romBanks), false, "")
	case chips.MBC2Battery:
		cart = NewMBC2(rom, uint8(romBanks), true, savePath)

	default:
		// It's easier for developing if we actually stop here.
		log.Fatalf("Unknown cartridge type 0x%02x (%s)", cartType, chips.Names[cartType])
	}

	return cart
}
