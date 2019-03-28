package memory

// MMU manages an arbitrary number of ordered address spaces, starting with the DMG boot ROM by default.
// It also satisfies the AddressSpace interface.
type MMU struct {
	Spaces []Addressable
}

// NewMMU returns an instance of MMU initialized with optional address spaces.
func NewMMU(spaces []Addressable) *MMU {
	return &MMU{spaces}
}

// NewEmptyMMU returns an instance of MMU with no address space.
func NewEmptyMMU() *MMU {
	return &MMU{[]Addressable{}}
}

// Add an address space at the end of this MMU's list.
func (m *MMU) Add(space Addressable) {
	m.Spaces = append(m.Spaces, space)
}

// Contains returns whether one of the address spaces known to the MMU contains the given address. The first
// address space in the internal list containing a given address will shadow any other.
func (m *MMU) Contains(addr uint) bool {
	for _, space := range m.Spaces {
		if space.Contains(addr) {
			return true
		}
	}
	return false
}

// Returns the first space for which the address is handled.
func (m *MMU) space(addr uint) Addressable {
	for _, space := range m.Spaces {
		if space.Contains(addr) {
			return space
		}
	}
	return nil // TODO: VOID
}

// Read finds the first address space compatible with the given address and returns the value at that address.
func (m *MMU) Read(addr uint) uint8 {
	if space := m.space(addr); space != nil {
		return space.Read(addr)
	}
	return 0xff
}

// Write finds the first address space compatible with the given address and attempts writing the given value to that
// address. TODO: error handling for write only?
func (m *MMU) Write(addr uint, value uint8) {
	if space := m.space(addr); space != nil {
		space.Write(addr, value)
	}
}
