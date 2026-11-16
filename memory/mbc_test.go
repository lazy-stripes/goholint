package memory

import (
	"slices"
	"testing"
)

// ROM sizes to test for mapping bank 0 to high ROM (<=256KB).
var mbc1MapHighROMTestBanks = []int{2, 4, 8, 16, 32}

// Validates that it is possible to map bank 0 to the high ROM region (4000-7fff)
// for a ROM up to 256KB (i.e. with 2, 4, 8 or 16 banks).
func testMBC1MapHighRomToBank0ForBanks(t *testing.T, banks int) {
	if !slices.Contains(mbc1MapHighROMTestBanks, banks) {
		t.Fatalf("Invalid number of banks (%d), should be 2, 4, 8 or 16", banks)
	}

	t.Logf("Testing MBC1 with %dKB ROM (%d banks)", 16*banks, banks)

	data := make([]byte, ROMBankSize*banks) // 32KB to 256KB ROM
	for bank := range banks {
		for i := range ROMBankSize {
			data[(ROMBankSize*bank)+i] = byte(bank)
		}
	}
	mbc1 := NewMBC1(NewBankedROM(data), nil)

	// Validate low rom (should always be bank 0)
	for i := range uint16(ROMBankSize) {
		v := mbc1.Read(i)
		if v != 0x00 {
			t.Fatalf("Invalid value %#02x for bank 0", v)
		}
	}

	// Validate high rom against the selected bank. When requesting a bank higher
	// than the amount available, they should wrap around and include bank 0.
	for bank := range uint8(0x20) {
		// Select bank. Bank 0 will translate to bank 1.
		expected := bank & (uint8(banks) - 1)
		if bank == 0 {
			expected = 1
		}
		mbc1.Write(0x2000, bank)
		//t.Logf("Testing bank %d, expecting value %d", bank, expected)

		// Read high rom contents. Should reflect bank number.
		for i := uint16(ROMBankSize); i < 0x8000; i++ {
			v := mbc1.Read(i)
			if v != expected {
				t.Fatalf("Invalid value %#02x for selected bank %#02x (expected %#02x)",
					v, bank, expected)
			}
		}
	}
}

// Validates that it is possible to map bank 0 to the high ROM region (4000-7fff)
// for ROMs that are 256KB or less.
//
// See https://gbdev.io/pandocs/MBC1.html#20003fff--rom-bank-number-write-only
// for details.
func TestMBC1MapHighRomToBank0(t *testing.T) {
	// Test behavior for all ROM sizes between 32KB and 256KB.
	for _, banks := range mbc1MapHighROMTestBanks {
		testMBC1MapHighRomToBank0ForBanks(t, banks)
	}
}

// ROM sizes to test for mapping banks 0x20/0x40/0x60 to high ROM (>=1MB).
var mbc1MapLowerZeroBankToHighROMBanks = []int{64, 128}

// Validates that it is not possible to map a bank to the high ROM region
// (4000-7fff) if the lower five bits of the bank number are all zero (for ROMs
// 1MB or larger).
func testMBC1MapLowerZeroBankToHighROMForBanks(t *testing.T, banks int) {
	if !slices.Contains(mbc1MapLowerZeroBankToHighROMBanks, banks) {
		t.Fatalf("Invalid number of banks (%d), should be 64 or 128", banks)
	}

	t.Logf("Testing MBC1 with %dMB ROM (%d banks)", banks/64, banks)

	data := make([]byte, ROMBankSize*banks) // 1MB or 2MB ROM

	for bank := range banks {
		for i := range ROMBankSize {
			data[(ROMBankSize*bank)+i] = byte(bank)
		}
	}
	mbc1 := NewMBC1(NewBankedROM(data), nil)

	// Validate low rom (should always be bank 0)
	for i := range uint16(ROMBankSize) {
		v := mbc1.Read(i)
		if v != 0x00 {
			t.Fatalf("Invalid value %#02x for bank 0", v)
		}
	}

	// Validate high rom against the selected bank. When requesting a bank whose
	// lower 5 bits are zero (0x20, 0x40, 0x60), we should get the bank directly
	// after (0x21, 0x41, 0x61).
	for bank := range uint8(0x80) {
		// Select bank. Bank 0 will translate to bank 1.
		expected := bank & (uint8(banks) - 1)
		if bank&0x1f == 0 {
			expected += 1
		}
		mbc1.Write(0x2000, bank&0x1f)
		mbc1.Write(0x4000, bank>>5)

		//t.Logf("Testing bank %d, expecting value %d", bank, expected)

		// Read high rom contents. Should reflect bank number except for our
		// special cases.
		for i := uint16(ROMBankSize); i < 0x8000; i++ {
			v := mbc1.Read(i)
			if v != expected {
				t.Fatalf("Invalid value %#02x for selected bank %#02x (expected %#02x)",
					v, bank, expected)
			}
		}
	}
}

func TestMBC1MapLowerZeroBankToHighROM(t *testing.T) {
	// Test behavior for all ROM sizes between 32KB and 256KB.
	for _, banks := range mbc1MapLowerZeroBankToHighROMBanks {
		testMBC1MapLowerZeroBankToHighROMForBanks(t, banks)
	}
}
