package memory

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"strings"
)

// ROM is a read-only special case of RAM, initialized from a binary file.
type ROM struct {
	RAM
}

// NewROM instantiates a read-only chunk of memory from a binary dump.
// Try to support ZIP files (room for improvement there).
func NewROM(filename string, start uint16) *ROM {
	// Try decompressing ZIP first, if not, read raw bytes.
	// TODO: Try to see if pre-read bytes can be decompressed, if not use them as-is.
	var bytes []byte
	archive, err := zip.OpenReader(filename)
	switch err {
	case zip.ErrFormat:
		// Not a ZIP file, treat as GB ROM directly.
		bytes, err = ioutil.ReadFile(filename)
		if err != nil {
			panic(fmt.Sprintf(" !!! Cannot read ROM file %s (%s)\n", filename, err))
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
			bytes, err = ioutil.ReadAll(rc)
			if err != nil {
				log.Fatal(err.Error())
			}
			rc.Close()
		}
		if bytes == nil {
			log.Fatalf("No GB ROM found in %s", filename)
		}
	default:
		// Improper ZIP file.
		log.Fatal(err.Error())
	}

	return &ROM{RAM{Start: start, Bytes: bytes}}
}

// Write does nothing and displays an error, for reasons I hope are obvious.
func (r *ROM) Write(addr uint16, value uint8) {
	log.Sub("rom").Warningf("Attempt to write %x to read-only address space at %#x",
		value, addr)
}

// Internal read that doesn't conform to the Adressable interface, used for
// ROMs with memory controllers, which can then have a size well over 0xffff.
func (r *ROM) read(addr uint) uint8 {
	offset := addr - uint(r.Start)
	if offset > uint(len(r.Bytes)) {
		log.Sub("rom").Warningf("Read overflow at %#x", addr)
		return 0xff
	}
	return r.Bytes[offset]
}
