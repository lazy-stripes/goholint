package ppu

import "errors"

// FIFO shifting out pixels to the display. FIXME: using bytes for tests. TODO: Pixel.
type FIFO struct {
	fifo [16]byte
	out  int
	in   int
	len  int
}

// Push a pixel in the FIFO.
func (f *FIFO) Push(pixel byte) error {
	if f.len == len(f.fifo) {
		return errors.New("ppu: FIFO buffer overrun")
	}
	f.fifo[f.in] = pixel
	f.in = (f.in + 1) % len(f.fifo)
	f.len++
	return nil
}

// Pop a pixel out of the FIFO. TODO: f.Display.Write(pixel)
func (f *FIFO) Pop() (pixel byte, err error) {
	// Do nothing if less than 8 pixels available to shift out..
	if f.len < 8 {
		return 0, errors.New("ppu: FIFO buffer underrun")
	}
	pixel = f.fifo[f.out]
	f.out = (f.out - 1) % len(f.fifo)
	f.len--
	return pixel, nil
}
