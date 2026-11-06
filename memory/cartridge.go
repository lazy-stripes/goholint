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

// TODO: this is getting big, do I dare create a cartridge package?
var OldLicensees = map[byte]string{
	0x00: "",
	0x01: "Nintendo",
	0x08: "Capcom",
	0x09: "HOT-B",
	0x0a: "Jaleco",
	0x0b: "Coconuts Japan",
	0x0c: "Elite Systems",
	0x13: "EA (Electronic Arts)",
	0x18: "Hudson Soft",
	0x19: "ITC Entertainment",
	0x1a: "Yanoman",
	0x1d: "Japan Clary",
	0x1f: "Virgin Games Ltd.",
	0x24: "PCM Complete",
	0x25: "San-X",
	0x28: "Kemco",
	0x29: "SETA Corporation",
	0x30: "Infogrames",
	0x31: "Nintendo",
	0x32: "Bandai",
	0x33: "See new licensee code",
	0x34: "Konami",
	0x35: "HectorSoft",
	0x38: "Capcom",
	0x39: "Banpresto",
	0x3c: "Entertainment Interactive (stub)",
	0x3e: "Gremlin",
	0x41: "Ubi Soft",
	0x42: "Atlus",
	0x44: "Malibu Interactive",
	0x46: "Angel",
	0x47: "Spectrum HoloByte",
	0x49: "Irem",
	0x4a: "Virgin Games Ltd.",
	0x4d: "Malibu Interactive",
	0x4f: "U.S. Gold",
	0x50: "Absolute",
	0x51: "Acclaim Entertainment",
	0x52: "Activision",
	0x53: "Sammy USA Corporation",
	0x54: "GameTek",
	0x55: "Park Place",
	0x56: "LJN",
	0x57: "Matchbox",
	0x59: "Milton Bradley Company",
	0x5a: "Mindscape",
	0x5b: "Romstar",
	0x5c: "Naxat Soft",
	0x5d: "Tradewest",
	0x60: "Titus Interactive",
	0x61: "Virgin Games Ltd.",
	0x67: "Ocean Software",
	0x69: "EA (Electronic Arts)",
	0x6e: "Elite Systems",
	0x6f: "Electro Brain",
	0x70: "Infogrames",
	0x71: "Interplay Entertainment",
	0x72: "Broderbund",
	0x73: "Sculptured Software",
	0x75: "The Sales Curve Limited",
	0x78: "THQ",
	0x79: "Accolade8",
	0x7a: "Triffix Entertainment",
	0x7c: "MicroProse",
	0x7f: "Kemco",
	0x80: "Misawa Entertainment",
	0x83: "LOZC G.",
	0x86: "Tokuma Shoten",
	0x8b: "Bullet-Proof Software",
	0x8c: "Vic Tokai Corp.",
	0x8e: "Ape Inc.",
	0x8f: "I'Max",
	0x91: "Chunsoft Co.",
	0x92: "Video System",
	0x93: "Tsubaraya Productions",
	0x95: "Varie",
	0x96: "Yonezawa/S'Pal",
	0x97: "Kemco",
	0x99: "Arc",
	0x9a: "Nihon Bussan",
	0x9b: "Tecmo",
	0x9c: "Imagineer",
	0x9d: "Banpresto",
	0x9f: "Nova",
	0xa1: "Hori Electric",
	0xa2: "Bandai",
	0xa4: "Konami",
	0xa6: "Kawada",
	0xa7: "Takara",
	0xa9: "Technos Japan",
	0xaa: "Broderbund",
	0xac: "Toei Animation",
	0xad: "Toho",
	0xaf: "Namco",
	0xb0: "Acclaim Entertainment",
	0xb1: "ASCII Corporation or Nexsoft",
	0xb2: "Bandai",
	0xb4: "Square Enix",
	0xb6: "HAL Laboratory",
	0xb7: "SNK",
	0xb9: "Pony Canyon",
	0xba: "Culture Brain",
	0xbb: "Sunsoft",
	0xbd: "Sony Imagesoft",
	0xbf: "Sammy Corporation",
	0xc0: "Taito",
	0xc2: "Kemco",
	0xc3: "Square",
	0xc4: "Tokuma Shoten",
	0xc5: "Data East",
	0xc6: "Tonkin House",
	0xc8: "Koei",
	0xc9: "UFL",
	0xca: "Ultra Games",
	0xcb: "VAP, Inc.",
	0xcc: "Use Corporation",
	0xcd: "Meldac",
	0xce: "Pony Canyon",
	0xcf: "Angel",
	0xd0: "Taito",
	0xd1: "SOFEL (Software Engineering Lab)",
	0xd2: "Quest",
	0xd3: "Sigma Enterprises",
	0xd4: "ASK Kodansha Co.",
	0xd6: "Naxat Soft",
	0xd7: "Copya System",
	0xd9: "Banpresto",
	0xda: "Tomy",
	0xdb: "LJN",
	0xdd: "Nippon Computer Systems",
	0xde: "Human Ent.",
	0xdf: "Altron",
	0xe0: "Jaleco",
	0xe1: "Towa Chiki",
	0xe2: "Yutaka",
	0xe3: "Varie",
	0xe5: "Epoch",
	0xe7: "Athena",
	0xe8: "Asmik Ace Entertainment",
	0xe9: "Natsume",
	0xea: "King Records",
	0xeb: "Atlus",
	0xec: "Epic/Sony Records",
	0xee: "IGS",
	0xf0: "A Wave",
	0xf3: "Extreme Entertainment",
	0xff: "LJN",
}

var NewLicensees = map[string]string{
	"00": "",
	"01": "Nintendo Research & Development ",
	"08": "Capcom",
	"13": "EA (Electronic Arts)",
	"18": "Hudson Soft",
	"19": "B-AI",
	"20": "KSS",
	"22": "Planning Office WADA",
	"24": "PCM Complete",
	"25": "San-X",
	"28": "Kemco",
	"29": "SETA Corporation",
	"30": "Viacom",
	"31": "Nintendo",
	"32": "Bandai",
	"33": "Ocean Software/Acclaim Entertainment",
	"34": "Konami",
	"35": "HectorSoft",
	"37": "Taito",
	"38": "Hudson Soft",
	"39": "Banpresto",
	"41": "Ubi Soft",
	"42": "Atlus",
	"44": "Malibu Interactive",
	"46": "Angel",
	"47": "Bullet-Proof Software",
	"49": "Irem",
	"50": "Absolute",
	"51": "Acclaim Entertainment",
	"52": "Activision",
	"53": "Sammy USA Corporation",
	"54": "Konami",
	"55": "Hi Tech Expressions",
	"56": "LJN",
	"57": "Matchbox",
	"58": "Mattel",
	"59": "Milton Bradley Company",
	"60": "Titus Interactive",
	"61": "Virgin Games Ltd.",
	"64": "Lucasfilm Games",
	"67": "Ocean Software",
	"69": "EA (Electronic Arts)",
	"70": "Infogrames",
	"71": "Interplay Entertainment",
	"72": "Broderbund",
	"73": "Sculptured Software",
	"75": "The Sales Curve Limited",
	"78": "THQ",
	"79": "Accolade",
	"80": "Misawa Entertainment",
	"83": "LOZC G.",
	"86": "Tokuma Shoten",
	"87": "Tsukuda Original",
	"91": "Chunsoft Co.",
	"92": "Video System",
	"93": "Ocean Software/Acclaim Entertainment",
	"95": "Varie",
	"96": "Yonezawa/S'Pal",
	"97": "Kaneko",
	"99": "Pack-In-Video",
	"9H": "Bottom Up",
	"A4": "Konami (Yu-Gi-Oh!)",
	"BL": "MTO",
	"DK": "Kodansha",
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
				return nil, fmt.Errorf("opening ROM inside ZIP failed: %w", err)
			}
			data, err = io.ReadAll(rc)
			if err != nil {
				return nil, fmt.Errorf("reading ROM inside ZIP failed: %w", err)
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
	// XXX: Any reason why I should accept less than 32Kb anyway?
	if len(romData) < 0x150 {
		return nil, fmt.Errorf("truncated ROM: %[1]d (0x%[1]X) bytes", len(romData))
	}

	// Check what kind of chip is in the ROM, select the proper struct. Also
	// store metadata for nice logging and maybe using templates in file names.
	cartType := romData[0x0147]
	romSize := romData[0x0148]
	ramSize := romData[0x0149]

	romBanks := chips.ROMBanks[romSize]
	ramBanks := chips.RAMBanks[ramSize]

	cgbFlag := romData[0x0143]
	oldLicensee := romData[0x014b]
	newLicensee := string(romData[0x0144:0x145])

	// Use old licensee by default.
	licensee := OldLicensees[oldLicensee]

	// Docs say title is between 11 and 16 bytes depending whether we have a
	// "later cartridge" but it's unclear how we can tell. Licensee code and
	// CGB flag will have to do.
	titleLen := 16

	if cgbFlag&0x80 > 0 {
		// Non-ASCII value, non-zero value, so probably a valid CGB flag
		titleLen -= 1
	}

	if oldLicensee == 0x33 {
		// Use new licensee code, probably a "modern" cartridge.
		titleLen -= 4
		licensee = NewLicensees[newLicensee]
	}

	title := string(romData[0x134 : 0x134+titleLen])
	strings.TrimSuffix(title, "\x00")

	log.Infof("Cartridge type 0x%02x (%s)", cartType, chips.Names[cartType])
	log.Infof("ROM size type 0x%02x (%d banks)", romSize, romBanks)
	log.Infof("RAM size type 0x%02x (%d banks)", ramSize, ramBanks)
	log.Infof("Title: %s", title)
	if licensee != "" {
		log.Infof("Licensee: %s", licensee)
	}

	// TODO: validate
	var bus Addressable
	var ram Persistable

	switch cartType {
	case chips.ROMOnly:
		bus = NewROM(romData)
	case chips.MBC1:
		bus = NewMBC1(rom, uint8(romBanks), 0, false, "")
	case chips.MBC1RAM:
		bus = NewMBC1(rom, uint8(romBanks), ramBanks, false, "")
	case chips.MBC1RAMBattery:
		bus = NewMBC1(rom, uint8(romBanks), ramBanks, true, savePath)
	case chips.MBC2:
		bus = NewMBC2(rom, uint8(romBanks), false, "")
	case chips.MBC2Battery:
		bus = NewMBC2(rom, uint8(romBanks), true, savePath)

	default:
		// It's easier for developing if we actually stop here.
		return nil, fmt.Errorf("unknown cartridge type 0x%02x (%s)",
			cartType, chips.Names[cartType])
	}

	return cart
}
