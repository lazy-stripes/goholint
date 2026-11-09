package memory

import (
	"testing"
)

// Validates that it is possible to map bank 0 to the high ROM region (4000-7fff)
// for ROMs that are 256KB or less.
//
// See https://gbdev.io/pandocs/MBC1.html#20003fff--rom-bank-number-write-only
// for details.
func TestMBC1MapHighRomToBank0(t *testing.T) {
	banks := 8
	data := make([]byte, 0x4000*banks) // 128KB (TODO: test range from 32KB to 256KB)
	for bank := 0; bank < banks; bank++ {
		for i := 0; i < 0x4000; i++ {
			data[(0x4000*bank)+i] = byte(bank)
		}
	}
	mbc1 := NewMBC1(NewBankedROM(data), nil)

	// Validate low rom (should always be bank 0)
	for i := uint16(0); i < 0x4000; i++ {
		v := mbc1.Read(i)
		if v != 0x00 {
			t.Fatalf("Invalid value %#02x for bank 0", v)
		}
	}

	// Validate high rom against the selected bank. When requesting a bank higher
	// than the amount available, they should wrap around and include bank 0.
	for bank := uint8(0); bank < 0x20; bank++ {
		// Select bank. Bank 0 will translate to bank 1.
		expected := bank & (uint8(banks) - 1)
		if bank == 0 {
			expected = 1
		}
		mbc1.Write(0x2000, bank)
		//t.Logf("Testing bank %d, expecting value %d", bank, expected)

		// Read high rom contents. Should reflect bank number.
		for i := uint16(0x4000); i < 0x8000; i++ {
			v := mbc1.Read(i)
			if v != expected {
				t.Fatalf("Invalid value %#02x for selected bank %#02x (expected %#02x)",
					v, bank, expected)
			}
		}
	}
}
