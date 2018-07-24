package ppu

import (
	"tigris.fr/gameboy/fifo"
)

// Fetcher reads tile data from VRAM and pushes pixels to PPU FIFO.
type Fetcher struct {
	fifo *fifo.FIFO
}

func (f *Fetcher) Tick() {
	// TODO
}
