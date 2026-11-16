package memory

import (
	"os"
	"testing"
)

// Use the one ROM file officially shipping with Goholint.
const testROMPath = "../assets/boot.rom"

func validateROMData(romData []byte, rom *ROM, t *testing.T) {
	for addr, expected := range romData {
		value := rom.Read(uint16(addr))
		if value != expected {
			t.Fatalf("Invalid value at ROM address %#04x (expected %#02x, read %#02x)",
				addr, expected, value)
		}
	}
}

// Validate that creating a ROM from an existing byte array simply makes a copy
// of that array.
func TestNewROM(t *testing.T) {
	romData, err := os.ReadFile(testROMPath)
	if err != nil {
		t.Fatal(err.Error())
	}

	rom := NewROM(romData)
	validateROMData(romData, rom, t)
}

// Validate that creating a ROM from a file creates a copy of that file in the
// ROMs internal byte array.
func TestNewROMFromFile(t *testing.T) {
	romData, err := os.ReadFile(testROMPath)
	if err != nil {
		t.Fatal(err.Error())
	}
	rom := NewROMFromFile(testROMPath)
	validateROMData(romData, rom, t)
}

// Validate that ROM data is read-only.
func TestReadOnlyROM(t *testing.T) {
	romData, err := os.ReadFile(testROMPath)
	if err != nil {
		t.Fatal(err.Error())
	}

	rom := NewROM(romData)

	// Try overwriting each byte with its complement.
	for addr := range uint16(len(romData)) {
		rom.Write(addr, ^rom.Read(addr))
	}
	validateROMData(romData, rom, t)
}
