package ppu

import (
	"tigris.fr/gameboy/fifo"
	"tigris.fr/gameboy/memory"
	"tigris.fr/gameboy/ppu/states"
)

// Fetcher reads tile data from VRAM and pushes pixels to PPU FIFO.
type Fetcher struct {
	Enabled    bool
	fifo       *fifo.FIFO
	vRAM       memory.Addressable
	ticks      int
	state      states.State
	mapAddr    uint // Start address of tile map row
	dataAddr   uint
	tileOffset uint8 // X offset in the tile map row (will wrap around)
	tileLine   uint8 // Y offset (in pixels) in the tile

	tileID   uint8
	tileData [8]uint8
}

// Fetch a line of pixels from the given tile in the given tilemap address space when Tick() is called.
func (f *Fetcher) Fetch(mapAddr, dataAddr uint, tileOffset, tileLine uint8) {
	f.mapAddr, f.dataAddr = mapAddr, dataAddr
	f.tileOffset, f.tileLine = tileOffset, tileLine
	f.state = states.ReadTileID
	f.Enabled = true
}

// Tick advances the fetcher's state machine one step.
func (f *Fetcher) Tick() {
	if !f.Enabled {
		return
	}

	f.ticks++
	if f.ticks < ClockFactor {
		return
	}

	// Reset tick counter and execute next state
	f.ticks = 0

	switch f.state {
	case states.ReadTileID:
		f.tileID = f.vRAM.Read(f.mapAddr + uint(f.tileOffset))
		f.state = states.ReadTileData0

	case states.ReadTileData0:

	}
}
