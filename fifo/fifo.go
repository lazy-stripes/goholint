package fifo

import "errors"

type Thing interface{}

// FIFO shifting out pixels to the display. FIXME: using bytes for tests. TODO: Pixel.
type FIFO struct {
	fifo   []interface{}
	out    int
	in     int
	len    int
	minLen int
}

func New(size, minLen int) *FIFO {
	return &FIFO{fifo: make([]interface{}, size, size), minLen: minLen}
}

// Pre-defined errors to only instantiate them once.
var errFIFOOverflow = errors.New("FIFO buffer overflow")
var errFIFOUnderrun = errors.New("FIFO buffer underrun")

// Clear and reset FIFO.
func (f *FIFO) Clear() {
	f.in, f.out, f.len = 0, 0, 0
}

// Push a pixel in the FIFO.
func (f *FIFO) Push(item interface{}) error {
	if f.len == len(f.fifo) {
		return errFIFOOverflow
	}
	f.fifo[f.in] = item
	f.in = (f.in + 1) % len(f.fifo)
	f.len++
	return nil
}

// Pop a pixel out of the FIFO. TODO: f.Display.Write(pixel)
func (f *FIFO) Pop() (item interface{}, err error) {
	// Do nothing if we only have 8 pixels or less available to shift out..
	if f.len <= f.minLen {
		return 0, errFIFOUnderrun
	}
	item = f.fifo[f.out]
	f.out = (f.out + 1) % len(f.fifo)
	f.len--
	return item, nil
}

// PushAll calls Push for each item in the slice it receives.
func (f *FIFO) PushAll(items []interface{}) error {
	for _, item := range items {
		if err := f.Push(item); err != nil {
			return err
		}
	}
	return nil
}

// Size returns the current number of items in the FIFO.
func (f *FIFO) Size() int {
	return f.len
}
