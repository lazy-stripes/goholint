package ppu

import "errors"

// FIFO storing generic items and supporting a minimum size under which it can't Pop.
// TODO: try implementing a specific PixelFIFO and see if it helps with performances.
type FIFO struct {
	fifo   []interface{}
	out    int
	in     int
	len    int
	minLen int
}

// NewFifo returns an empty FIFO of the given size with the given minimum length.
func NewFifo(size, minLen int) *FIFO {
	return &FIFO{fifo: make([]interface{}, size), minLen: minLen}
}

// Pre-defined errors to only instantiate them once.
var errFIFOOverflow = errors.New("FIFO buffer overflow")
var errFIFOUnderrun = errors.New("FIFO buffer underrun")

// Clear and reset FIFO.
func (f *FIFO) Clear() {
	f.in, f.out, f.len = 0, 0, 0
}

// Push an item in the FIFO.
func (f *FIFO) Push(item interface{}) error {
	if f.len == len(f.fifo) {
		return errFIFOOverflow
	}
	f.fifo[f.in] = item
	f.in = (f.in + 1) % len(f.fifo)
	f.len++
	return nil
}

// Pop an item out of the FIFO.
func (f *FIFO) Pop() (item interface{}, err error) {
	// Do nothing if we only have the minimum length or less available to shift out.
	if f.len <= f.minLen {
		return 0, errFIFOUnderrun
	}
	item = f.fifo[f.out]
	f.out = (f.out + 1) % len(f.fifo)
	f.len--
	return item, nil
}

// Size returns the current number of items in the FIFO.
func (f *FIFO) Size() int {
	return f.len
}
