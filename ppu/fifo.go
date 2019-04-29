package ppu

import "errors"

// FIFO holds the PPU's pixel FIFO shifting out pixels to the display with the
// guarantee that there will always be 8 pixels available to mix sprites in.
type FIFO struct {
	fifo [16]Pixel
	out  int
	in   int
	len  int
}

// Pre-defined errors to only instantiate them once.
var errFIFOOverflow = errors.New("FIFO buffer overflow")
var errFIFOUnderrun = errors.New("FIFO buffer underrun")

// Clear and reset FIFO.
func (f *FIFO) Clear() {
	f.in, f.out, f.len = 0, 0, 0
}

// Push an item in the FIFO.
func (f *FIFO) Push(pixel Pixel) error {
	if f.len == len(f.fifo) {
		return errFIFOOverflow
	}
	f.fifo[f.in] = pixel
	f.in = (f.in + 1) % len(f.fifo)
	f.len++
	return nil
}

// Pop an item out of the FIFO.
func (f *FIFO) Pop() (pixel Pixel, err error) {
	// Do nothing if we only have the minimum length or less to shift out.
	if f.len <= 8 {
		return pixel, errFIFOUnderrun
	}
	pixel = f.fifo[f.out]
	f.out = (f.out + 1) % len(f.fifo)
	f.len--
	return pixel, nil
}

// Size returns the current number of items in the FIFO.
func (f *FIFO) Size() int {
	return f.len
}

// Mix sprite pixel data in the lower half of the FIFO. Priorities are hard to
// understand in most docs I've seen, so this is empirical at best.
func (f *FIFO) Mix(offset int, pixel Pixel) {
	index := (f.out + offset) % len(f.fifo)
	current := f.fifo[index]

	// Discard pixel if it's transparent.
	if pixel.Color == 0 {
		return
	}

	// Mix pixel in if the current one is from the background.
	// TODO: OBJ-BG priority attribute bit.
	// TODO: OAM index, possibly needed in Pixel struct.
	if current.Palette == PixelBGP {
		f.fifo[index] = pixel
	}
}
