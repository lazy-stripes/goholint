package memory_test

import (
	"math/rand"
	"testing"

	"github.com/lazy-stripes/goholint/memory"
)

// testRegisters stores addresses and register values in a bespoke struct that
// embeds Registers, the way it's supposed to be used. The used addressed and
// values are available to basically validate that Registers acts like a map.
type testRegisters struct {
	memory.Registers

	addresses []uint16
	values    []byte
}

// newRandomTestRegisters returns a testRegisters instance of the given size
// filled with random values at random addresses. Both values and addresses are
// stored in the order they were set for further validation.
func newRandomTestRegisters(size int) *testRegisters {
	tr := &testRegisters{
		Registers: make(memory.Registers),
		addresses: make([]uint16, size),
		values:    make([]byte, size),
	}

	for i := range size {
		// Make sure to not generate duplicate addresses.
		addr := uint16(rand.Intn(0x10000))
		for tr.Registers[addr] != nil {
			addr = uint16(rand.Intn(0x10000))
		}
		value := uint8(rand.Intn(0xfe) + 1) // Avoid zero value or 0xff.
		tr.addresses[i] = addr
		tr.values[i] = value
		tr.Registers[addr] = &tr.values[i]
	}

	return tr
}

// Validate that Contains() returns the proper values for a Registers instance
// initialized with random registers.
func TestRegistersContains(t *testing.T) {
	// Test with an arbitrary large number of registers.
	tr := newRandomTestRegisters(1024)

	// Validate that all generated addresses are seen as contained in the
	// Register address space. Store them so we can check that all other
	// addresses are not contained in the space.
	contained := make(map[uint16]bool)
	for _, addr := range tr.addresses {
		if !tr.Contains(addr) {
			t.Fatalf("Address %#04x missing from Registers address space", addr)
		}
		contained[addr] = true
	}

	// Validate that all other available addresses are not present in Registers.
	for addr := range 0x10000 {
		if contained[uint16(addr)] {
			continue
		}
		if tr.Contains(uint16(addr)) {
			t.Fatalf("Address %#04x should not be in Registers address space", addr)
		}
	}
}

// Validate that we can read values stored in registers and that incorrect
// addresses produce a default value.
func TestRegistersRead(t *testing.T) {
	// Test with an arbitrary large number of registers.
	tr := newRandomTestRegisters(1024)

	// Validate that all generated addresses map to a non-default value (neither
	// 0x00 nor 0xff). Store those addresses to validate that every other read
	// should return a default value.
	contained := make(map[uint16]bool)
	for i := range tr.addresses {
		addr := tr.addresses[i]
		expected := tr.values[i]
		value := tr.Read(addr)

		if value != expected {
			t.Fatalf("Expected %#02x at register address %#04x, got %#02x",
				expected, value, addr)
		}
		contained[addr] = true
	}

	// Validate that all other available addresses return a default value.
	for addr := range 0x10000 {
		if contained[uint16(addr)] {
			continue
		}
		value := tr.Read(uint16(addr))
		if value != 0xff {
			t.Fatalf("Unmapped register address %#04x has non-default value %#02x",
				addr, value)
		}
	}
}

// Validate that we can write values stored in registers and that incorrect
// addresses cannot be written to.
func TestRegistersWrite(t *testing.T) {
	// Test with an arbitrary large number of registers.
	tr := newRandomTestRegisters(1024)

	// Validate that all generated addresses map to a value that we can modify
	// and that others will always ever return 0xff.
	for addr := range 0x10000 {
		// Read value, increment and write it back, then read it again.
		value := tr.Read(uint16(addr))
		expected := value + 1
		tr.Write(uint16(addr), expected)
		modified := tr.Read(uint16(addr))

		// If the address is an actual register, the value must have changed.
		if tr.Contains(uint16(addr)) {
			if modified != expected {
				t.Fatalf("Write to %#04x failed (expected %#02x, got %#02x)",
					addr, expected, modified)
			}
		} else {
			// If the address is not an actual register, value must be the same
			// default one.
			if value != 0xff {
				t.Fatalf("Unmapped register address %#04x has non-default value %#02x",
					addr, value)
			}

			if modified != value {
				t.Fatalf("Unmapped register address %#04x has modified value %#02x instead of %#02x",
					addr, modified, value)
			}
		}
	}
}
