package memory

// MMU manages an arbitrary number of ordered address spaces. It also satisfies
// the AddressSpace interface.
type MMU struct {
	Spaces []Addressable
}

// NewMMU returns an instance of MMU initialized with existing address spaces.
func NewMMU(spaces []Addressable) *MMU {
	return &MMU{spaces}
}

// NewEmptyMMU returns an instance of MMU with no address space.
func NewEmptyMMU() *MMU {
	var empty []Addressable
	return &MMU{empty}
}

// Add an address space at the end of this MMU's list.
func (m *MMU) Add(space Addressable) {
	m.Spaces = append(m.Spaces, space)
}

// Contains returns whether one of the address spaces known to the MMU contains
// the given address. The first address space in the internal list containing a
// given address will shadow any other that may contain it.
func (m *MMU) Contains(addr uint16) bool {
	for _, space := range m.Spaces {
		if space.Contains(addr) {
			return true
		}
	}
	return false
}

// Returns the first space for which the address is handled.
func (m *MMU) space(addr uint16) Addressable {
	for _, space := range m.Spaces {
		if space.Contains(addr) {
			return space
		}
	}
	return nil
}

// Read finds the first address space compatible with the given address and
// returns the value at that address. If no space contains the requested
// address, it returns 0xff (emulates black bar on boot).
func (m *MMU) Read(addr uint16) uint8 {
	if space := m.space(addr); space != nil {
		return space.Read(addr)
	}
	log.Sub("mmu/read").Debugf("MMU.Read: Unmapped address 0x%04x", addr)
	return 0xff
}

// Write finds the first address space compatible with the given address and
// attempts writing the given value to that address.
func (m *MMU) Write(addr uint16, value uint8) {
	if space := m.space(addr); space != nil {
		space.Write(addr, value)
	} else {
		log.Sub("mmu/write").Debugf("MMU.Write: Unmapped address 0x%04x", addr)
	}
}
