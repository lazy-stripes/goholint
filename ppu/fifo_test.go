package ppu

import "testing"

func TestFIFO(t *testing.T) {
	f := FIFO{}

	for p := byte(1); p < 12; p++ {
		f.Push(Pixel{p, 0})

		if f.len != int(p) {
			t.Errorf("FIFO length mismatch. Expected %d, got %d", p, f.len)
		}
	}

	for p := byte(1); p < 4; p++ {
		pixel, err := f.Pop()
		if err != nil {
			t.Errorf("Error during Pop(): %s", err)
		}

		if pixel.Color != p {
			t.Errorf("Pop returned wrong value %x instead of 0x01", pixel)
		}
	}
}
