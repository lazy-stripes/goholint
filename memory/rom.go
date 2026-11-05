package memory

// Initialize sub-logger for ROM accesses.
func init() {
	log.Add("rom", "Unexpected writes and overflows in ROM address spaces")
}

// ROM is a read-only special case of RAM, initialized from a binary file.
type ROM struct {
	RAM
}

func NewROM(data []byte) *ROM {
	return &ROM{RAM{Start: 0, Bytes: data}}
}

// Write does nothing and displays an error, for reasons I hope are obvious.
func (r *ROM) Write(addr uint16, value uint8) {
	log.Sub("rom").Warningf("Attempt to write %x to read-only address space at %#x",
		value, addr)
}
