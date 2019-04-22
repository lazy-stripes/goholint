package memory

import (
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestRAMContains(t *testing.T) {
	cases := []struct {
		in   uint16
		want bool
	}{
		{0x0000, true},
		{0x0042, true},
		{0x007F, true},
		{0x0080, false},
		{0xFFFF, false},
	}

	ram := NewRAM(0, 128)
	for _, c := range cases {
		if got := ram.Contains(c.in); got != c.want {
			t.Errorf("RAM(%d).Contains(%d) == %t, want %t", 128, c.in, got, c.want)
		}
	}
}

func TestRAMWrite(t *testing.T) {
	cases := []uint16{
		0x0000,
		0x0042,
		0x007F,
	}

	error := uint16(0x0080)

	ram := NewRAM(0, 128)
	for _, c := range cases {
		ram.Write(c, 42)
	}

	// Expect a panic with the error case
	defer func() {
		if recover() == nil {
			t.Errorf("Invalid write address %d did not cause a panic", error)
		}
	}()

	ram.Write(error, 42)
}

func TestRAMRead(t *testing.T) {
	error := uint16(0x0080)

	ram := NewRAM(0, 128)
	for addr := uint16(0); addr < uint16(128); addr++ {
		in := uint8(rand.Intn(0xFF))
		ram.Write(addr, in)
		if got := ram.Read(addr); got != in {
			t.Errorf("RAM(%d).Read(%d) == %X, want %X", 128, addr, got, in)
		}
	}

	// Expect a panic with the error case
	defer func() {
		if recover() == nil {
			t.Errorf("Invalid read address %d did not cause a panic", error)
		}
	}()

	ram.Read(error)
}

func TestMMU(t *testing.T) {
	rompath := "../bin/DMG_ROM.bin"
	rom := NewROM(rompath, 0)
	ram := NewRAM(0, 0xffff)
	boot := NewMMU([]Addressable{rom, ram})

	romdump, err := ioutil.ReadFile(rompath)
	if err != nil {
		t.Errorf("Invalid ROM path '%s'", rompath)
	}
	for addr, want := range romdump {
		if got := boot.Read(uint16(addr)); got != want {
			t.Errorf("Byte mismatch at offset %d (expected %x, read %x)", addr, want, got)
		}
	}

	for romaddr := uint16(0); romaddr < 0x100; romaddr++ {
		want := boot.Read(romaddr)
		boot.Write(romaddr, want+1)
		got := boot.Read(romaddr)
		if got != want {
			t.Errorf("ROM write error at address %d (%x before write, %x after write)", romaddr, want, got)
		}
	}

	for addr := uint16(0x100); addr <= 0xffff; addr++ {
		want := boot.Read(addr) + 1
		boot.Write(addr, want)
		got := boot.Read(addr)
		if got != want {
			t.Errorf("RAM write failed at address %d (%x before write, %x after write)", addr, want, got)
		}
	}
}

func TestROMWrite(t *testing.T) {
	rom := NewROM("/dev/null", 0)
	rom.Write(0, 42)
}
