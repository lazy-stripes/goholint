package memory

import (
	"math/rand"
	"testing"
)

func TestRAMContains(t *testing.T) {
	cases := []struct {
		in   uint
		want bool
	}{
		{0x0000, true},
		{0x0042, true},
		{0x007F, true},
		{0x0080, false},
		{0xFFFF, false},
	}

	ram := NewRAM(128)
	for _, c := range cases {
		if got := ram.Contains(c.in); got != c.want {
			t.Errorf("RAM(%d).Contains(%d) == %t, want %t", len(ram), c.in, got, c.want)
		}
	}
}

func TestRAMWrite(t *testing.T) {
	cases := []uint{
		0x0000,
		0x0042,
		0x007F,
	}

	error := uint(0x0080)

	ram := NewRAM(128)
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
	error := uint(0x0080)

	ram := NewRAM(128)
	for addr := uint(0); addr < uint(len(ram)); addr++ {
		in := uint8(rand.Intn(0xFF))
		ram.Write(addr, in)
		if got := ram.Read(addr); got != in {
			t.Errorf("RAM(%d).Read(%d) == %X, want %X", len(ram), addr, got, in)
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
