package memory

// Memory scanner, à la Game Genie, to inspect memory. Modifying or forcing
// memory addresses should be delegated to a specific GameGenie struct.

// TODO: allow looking up 16-bit values. Endianness is gonna be a bitch.

// Haystack is a mapping of a 16-bit address to an 8-bit value representing all
// the memory locations the scanner still needs to search through.
type Haystack map[uint16]uint8

// Scanner stores a mapping of addresses to values (the haystack) that satisfy
// successive lookup criteria (the needles). By keeping the result of each
// successive lookup, this might help find specific memory locations for some
// variables whose value we can observe and search for.
type Scanner struct {
	memory Addressable // GameBoy RAM.

	Haystack Haystack // Addresses susceptible to hold the value.
}

// NewScanner returns a Scanner instance whose starting haystack is populated
// with all values read from the given addressable object at the memory
// locations that correspond to Gameboy RAM (0xc000-0xe000 and 0xff80-0xfffe).
func NewScanner(memory Addressable) (s *Scanner) {
	s = &Scanner{
		memory:   memory,
		Haystack: make(Haystack, RAMSize+HRAMSize),
	}

	// Initialize the current haystack from the whole current RAM.
	s.Clear()

	return s
}

// initHaystackFrom reads memory from addr to (addr + size) and stores the
// current value in the internal haystack.
func (s *Scanner) initHaystackFrom(addr, size uint16) {
	for offset := uint16(0); offset < size; offset++ {
		s.Haystack[addr+offset] = s.memory.Read(addr + offset)
	}
}

// Reset reinitializes the current haystack with a copy of current values in
// RAM in ranges 0xc000-0xe000 and 0xff80-0xfffe.
func (s *Scanner) Clear() {
	s.initHaystackFrom(RAMOffset, RAMSize)
	s.initHaystackFrom(HRAMOffset, HRAMSize)
}

// condition is a comparison function that returns the boolean result of that
// comparison. Used to conditionally decide whether to keep a lookup candidate
// in a new haystack.
type condition func(old, new uint8) bool

// conditionEq returns true if the value for a given address is the same as it
// was when last looked up.
func conditionEq(old, new uint8) bool {
	return new == old
}

// conditionEq returns true if the value for a given address is larger than it
// was when last looked up.
func conditionGt(old, new uint8) bool {
	return new > old
}

// conditionGte returns true if the value for a given address is larger than or
// equal to what it was when last looked up.
func conditionGte(old, new uint8) bool {
	return new >= old
}

// conditionLt returns true if the value for a given address is smaller than it
// was when last looked up.
func conditionLt(old, new uint8) bool {
	return new < old
}

// conditionLte returns true if the value for a given address is smaller than or
// equal to what it was when last looked up.
func conditionLte(old, new uint8) bool {
	return new <= old
}

// conditionValue returns a new condition function that will return true if
// the value for a given address is equal to the user-provided one.
func conditionValue(value uint8) condition {
	return func(_, new uint8) bool {
		return new == value
	}
}

// lookup goes through the current haystack (which will be the whole RAM the
// first time) and compares all of its values with the ones currently in memory
// using the condition function.
//
// Returns a new haystack containing onnly the addresses and values in current
// memory for which the condition function returned true. That new haystack can
// then be saved and used for the next lookup, or discarded to try a different
// lookup.
//
// See all the Lookup* methods for more details.
func (s *Scanner) lookup(cond condition) (newHaystack Haystack) {
	newHaystack = make(Haystack)

	for addr, oldValue := range s.Haystack {
		newValue := s.memory.Read(addr)
		if cond(oldValue, newValue) {
			newHaystack[addr] = newValue
		}
	}

	return newHaystack
}

// LookUp returns a haystack with the contents of all the memory addresses
// whose current value is equal to the given value for the set of addresses in
// the current haystack. The actual values in the current haystack are ignored.
func (s *Scanner) LookUp(value uint8) (newHaystack Haystack) {
	return s.lookup(conditionValue(value))
}

// LookUpEqual returns a haystack with the contents of all the memory addresses
// whose current value is identical to the value at the same address in the
// current haystack.
func (s *Scanner) LookUpEqual() (newHaystack Haystack) {
	return s.lookup(conditionEq)
}

// LookUpLarger returns a haystack with the contents of all the memory addresses
// whose current value is larger than the value at the same address in the
// current haystack.
func (s *Scanner) LookUpLarger() (newHaystack Haystack) {
	return s.lookup(conditionGt)
}

// LookUpLargerOrEqual returns a haystack with the contents of all the memory
// addresses whose current value is larger than or equal to the value at the
// same address in the current haystack.
func (s *Scanner) LookUpLargerOrEqual() (newHaystack Haystack) {
	return s.lookup(conditionGte)
}

// LookUpSmaller returns a haystack with the contents of all the memory addresses
// whose current value is smaller than the value at the same address in the
// current haystack.
func (s *Scanner) LookUpSmaller() (newHaystack Haystack) {
	return s.lookup(conditionLt)
}

// LookUpSmallerOrEqual returns a haystack with the contents of all the memory
// addresses whose current value is smaller than or equal to the value at the
// same address in the current haystack.
func (s *Scanner) LookUpSmallerOrEqual() (newHaystack Haystack) {
	return s.lookup(conditionLte)
}
