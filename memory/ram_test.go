package memory_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/lazy-stripes/goholint/memory"
)

// Validate that a RAM instance's Contains() method only returns true for all
// the actual addresses contained in that address space.
func TestRAMContains(t *testing.T) {
	size := 0x2000
	start := 0x4000

	ram := memory.NewRAM(uint16(start), size)
	for addr := range 0x10000 {
		if addr >= start && addr < start+size {
			if !ram.Contains(uint16(addr)) {
				t.Fatalf("Address %#04x should be contained in RAM", addr)
			}
		} else {
			if ram.Contains(uint16(addr)) {
				t.Fatalf("Address %#04x should not be contained in RAM", addr)
			}
		}
	}
}

// Validate that RAM can be read from and written to.
func TestRAMReadWrite(t *testing.T) {
	size := 0x2000
	start := 0x4000

	ram := memory.NewRAM(uint16(start), size)
	for addr := start; addr < start+size; addr++ {
		def := ram.Read(uint16(addr))
		if def != 0x00 {
			t.Fatalf("Invalid default value at address %#04x (expected 0x00, read %#02x)",
				addr, def)
		}

		expected := uint8(rand.Intn(0xfe)) + 1 // Non-zero, non-0xff random value
		ram.Write(uint16(addr), expected)
		value := ram.Read(uint16(addr))
		if value != expected {
			t.Fatalf("Write error at address %#04x (expected %#02x, read %#02x)",
				addr, expected, value)
		}
	}
}

// createTempRAMFile creates a temporary file of the given size, fills it with
// random data, and returns said data, the file path and a non-nil error if
// creating the file failed. If error is nil, it is the caller's responsibility
// to remove this file after testing.
func createTempRAMFile(size int) (path string, data []byte, err error) {
	// Create a RAM-sized file with random values.
	ramValues := make([]byte, size)
	rand.Read(ramValues)
	f, err := os.CreateTemp("", "goholint-test-ram-*")
	if err != nil {
		return "", nil, err
	}

	_, err = f.Write(ramValues)
	if err != nil {
		f.Close()
		os.Remove(f.Name())
		return "", nil, err
	}
	if err := f.Close(); err != nil {
		os.Remove(f.Name())
		return "", nil, err
	}

	return f.Name(), ramValues, nil
}

// Validate that Load() overwrites RAM with values in the given file.
func TestRAMLoad(t *testing.T) {
	size := 0x2000
	start := uint16(0x4000)

	tmpfile, ramValues, err := createTempRAMFile(size)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile) // Clean up when test is done

	ram := memory.NewRAM(start, size)
	err = ram.Load(tmpfile)
	if err != nil {
		t.Fatalf("Loading %s to RAM failed: %v", tmpfile, err)
	}

	for i, expected := range ramValues {
		addr := start + uint16(i)
		value := ram.Read(addr)
		if value != expected {
			t.Fatalf("Invalid RAM value loaded at address %#04x (expected %#02x, read %#02x)",
				addr, expected, value)
		}
	}
}

// validateCleanRAM checks that the given RAM instance is all zeroed-out.
func validateCleanRAM(ram *memory.RAM, t *testing.T) {
	for i, v := range ram.Bytes {
		if v != 0x00 {
			t.Fatalf("Invalid RAM value at address %#04x (expected 0x00, read %#02x)",
				ram.Start+uint16(i), v)
		}
	}
}

// Validate that Load() errors do not overwrite RAM and leave it as it was.
func TestRAMLoadErrors(t *testing.T) {
	size := 0x2000
	start := uint16(0x4000)

	ram := memory.NewRAM(start, size)
	validateCleanRAM(ram, t)

	// Try loading an invalid file path. Note that files that don't exist will
	// not trigger an error here.
	err := ram.Load("/") // Can't load a folder
	if err == nil {
		t.Fatalf("No error when loading invalid file path")
	}
	validateCleanRAM(ram, t)

	// Try loading a file of the wrong size.
	tmpfile, _, err := createTempRAMFile(size * 2)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile) // Clean up when test is done

	err = ram.Load(tmpfile) // Can't load a file not exactly the size of RAM
	if err == nil {
		t.Fatalf("No error when loading file of wrong size")
	}
	validateCleanRAM(ram, t)

	// Try loading a file that does not exist. This should NOT return an error.
	err = ram.Load("/this/should/not/exist/anywhere")
	if err != nil {
		t.Fatalf("Unexpected error when loading non-existant file (%v)", err)
	}
	validateCleanRAM(ram, t)
}

// TODO: test file save.
