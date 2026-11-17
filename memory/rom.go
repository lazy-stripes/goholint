package memory

// Initialize sub-logger for ROM accesses.
func init() {
	log.Add("rom", "Unexpected writes and overflows in ROM address spaces")
}

// ROM is a read-only memory array that always starts from address zero.
type ROM struct {
	*Memory
}

// NewROM returns a ROM instance with its contents set to the given data.
func NewROM(data []byte) *ROM {
	return &ROM{Memory: &Memory{Start: 0, Bytes: data}}
}

// NewROMFromFile returns a ROM instance with its contents loaded from the
// given path.
//
// FIXME: error handling.
func NewROMFromFile(filename string) *ROM {
	r := &ROM{Memory: &Memory{Start: 0}}
	r.Load(filename)
	return r
}

// Write does nothing and displays an error, for reasons I hope are obvious.
func (r *ROM) Write(addr uint16, value uint8) {
	log.Sub("rom").Warningf("Ignoring Write(%#04x, %#02x) to read-only memory",
		addr, value)
}
