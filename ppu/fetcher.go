package ppu

import (
	"go.tigris.fr/gameboy/fifo"
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu/states"
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
	signedID   bool

	tileID   uint8
	tileData [8]uint8
}

// Start fetching a line of pixels from the given tile in the given tilemap address space when Tick() is called.
func (f *Fetcher) Start(mapAddr, dataAddr uint, tileOffset, tileLine uint8, signedID bool) {
	f.mapAddr, f.dataAddr = mapAddr, dataAddr
	f.tileOffset, f.tileLine = tileOffset, tileLine
	f.signedID = signedID
	f.state = states.ReadTileID
	f.Enabled = true
	f.fifo.Clear()
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
		//log.Printf("%04x: %02x\n", f.mapAddr+uint(f.tileOffset), f.tileID)

	case states.ReadTileData0:
		f.ReadTileLine(0)
		f.state = states.ReadTileData1

	case states.ReadTileData1:
		f.ReadTileLine(1)
		f.state = states.PushToFIFO

	case states.PushToFIFO:
		if f.fifo.Size() <= 8 {
			for i := 0; i < 8; i++ { // TODO: PixelFIFO directly handling [8]uint8
				f.fifo.Push(f.tileData[i])
			}
			f.tileOffset = (f.tileOffset + 1) % 32
			f.state = states.ReadTileID
		}
	}
}

// ReadTileLine updates internal pixel buffer with LSB or MSB tile line depending on parameter.
func (f *Fetcher) ReadTileLine(byteOffset uint8) {
	// TODO: attributes, 16-pixel height, reverse (well, sprites really)
	var offset uint
	if f.signedID {
		offset = uint(int(f.dataAddr) + (int(f.tileID) * 16))
	} else {
		offset = f.dataAddr + (uint(f.tileID) * 16)
	}
	addr := offset + (uint(f.tileLine) * 2)

	pixelData := f.vRAM.Read(addr + uint(byteOffset))
	for bitPos := 7; bitPos >= 0; bitPos-- {
		if byteOffset == 0 {
			// Least significant bit, replace previous value.
			f.tileData[7-bitPos] = (pixelData >> uint(bitPos)) & 1
		} else {
			// Most significant bit, update previous value.
			f.tileData[7-bitPos] |= ((pixelData >> uint(bitPos)) & 1) << 1
		}
	}
}
