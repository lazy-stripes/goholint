package ppu

import (
	"github.com/lazy-stripes/goholint/memory"
	"github.com/lazy-stripes/goholint/ppu/states"
)

// Fetcher reads tile data from VRAM and pushes pixels to PPU FIFO.
type Fetcher struct {
	Enabled         bool
	fifo            *FIFO
	vRAM            memory.Addressable
	oamRAM          memory.Addressable
	ticks           int
	state, oldState states.State
	lcdc            *uint8 // Reference to LCDC for sprites height bit
	mapAddr         uint16 // Start address of BG/Windows map row
	dataAddr        uint16 // Start address of Sprite/BG tile data
	tileOffset      uint8  // X offset in the tile map row (will wrap around)
	tileLine        uint8  // Y offset (in pixels) in the tile
	signedID        bool

	tileID   uint8
	tileData [8]uint8

	sprite       Sprite // Stores X, Y and address in OAM
	spriteID     uint8
	spriteFlags  uint8
	spriteOffset uint8 // X offset for sprite (if not fully on screen)
	spriteLine   uint8 // Y offset (in pixels) in the sprite
	spriteData   [8]uint8
}

// NewFetcher creates a pixel fetcher instance that can read directly from
// video and OAM RAM.
func NewFetcher(ppu *PPU) *Fetcher {
	f := Fetcher{
		fifo:   &ppu.FIFO,
		lcdc:   &ppu.LCDC,
		vRAM:   ppu.VRAM.RAM,
		oamRAM: ppu.OAM.RAM,
	}
	return &f
}

// Start fetching a line of pixels from the given tile in the given tilemap
// address space when Tick() is called.
func (f *Fetcher) Start(mapAddr, dataAddr uint16, tileOffset, tileLine uint8, signedID bool) {
	f.mapAddr, f.dataAddr = mapAddr, dataAddr
	f.tileOffset, f.tileLine = tileOffset, tileLine
	f.signedID = signedID
	f.state = states.ReadTileID
	f.Enabled = true
	f.fifo.Clear()
}

// FetchSprite pauses the current fetching state to read sprite data and mix it
// in the pixel FIFO.
func (f *Fetcher) FetchSprite(sprite Sprite, spriteOffset, spriteLine uint8) {
	f.sprite = sprite
	f.spriteOffset, f.spriteLine = spriteOffset, spriteLine
	f.oldState = f.state
	f.state = states.ReadSpriteID
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
		f.tileID = f.vRAM.Read(f.mapAddr + uint16(f.tileOffset))
		f.state = states.ReadTileData0
		//logger.Printf("fetcher", "%04x: %02x\n", f.mapAddr+uint(f.tileOffset), f.tileID)

	case states.ReadTileData0:
		f.ReadTileLine(0, f.dataAddr, f.tileID, f.signedID, f.tileLine, 0, &f.tileData)
		f.state = states.ReadTileData1

	case states.ReadTileData1:
		f.ReadTileLine(1, f.dataAddr, f.tileID, f.signedID, f.tileLine, 0, &f.tileData)
		f.state = states.PushToFIFO

	case states.PushToFIFO:
		if f.fifo.Size() <= 8 {
			for i := 0; i < 8; i++ {
				f.fifo.Push(Pixel{f.tileData[i], PixelBGP, false})
			}
			f.tileOffset = (f.tileOffset + 1) % 32
			f.state = states.ReadTileID
		}
	case states.ReadSpriteID:
		// Read directly from OAM RAM.
		f.spriteID = f.oamRAM.Read(f.sprite.Address + 2) // We already read X&Y
		f.state = states.ReadSpriteFlags

	case states.ReadSpriteFlags:
		f.spriteFlags = f.oamRAM.Read(f.sprite.Address + 3)
		f.state = states.ReadSpriteData0

		// Account for 8×16 sprites. Quoting PanDocs [4.3 OAM]:
		//
		// In 8×16 mode (LCDC bit 2 = 1), the memory area at $8000-$8FFF is
		// still interpreted as a series of 8×8 tiles, where every 2 tiles form
		// an object. In this mode, this byte specifies the index of the first
		// (top) tile of the object. This is enforced by the hardware: the least
		// significant bit of the tile index is ignored; that is, the top 8×8
		// tile is “NN & $FE”, and the bottom 8×8 tile is “NN | $01”.
		if *f.lcdc&LCDCSpriteSize != 0 {
			if f.spriteLine < 8 {
				if f.spriteFlags&SpriteFlipY != 0 {
					f.spriteID |= 0x01 // Swap tiles
				} else {
					f.spriteID &= 0xfe
				}
			} else {
				if f.spriteFlags&SpriteFlipY != 0 {
					f.spriteID &= 0xfe // Swap tiles
				} else {
					f.spriteID |= 0x01
				}
				f.spriteLine -= 8
			}
		}

	case states.ReadSpriteData0:
		f.ReadTileLine(0, 0x8000, f.spriteID, false, f.spriteLine, f.spriteFlags, &f.spriteData)
		f.state = states.ReadSpriteData1

	case states.ReadSpriteData1:
		f.ReadTileLine(1, 0x8000, f.spriteID, false, f.spriteLine, f.spriteFlags, &f.spriteData)
		f.state = states.MixInFIFO

	case states.MixInFIFO:
		if f.fifo.Size() < 8 {
			break
		}

		// Mix sprite pixels with FIFO, taking into account offset if sprite
		// is only partially displayed (i.e. entering screen from the left).
		var palette uint8
		if f.spriteFlags&0x10 == 0 {
			palette = PixelOBP0
		} else {
			palette = PixelOBP1
		}
		bgOverObj := f.spriteFlags&SpritePriority != 0
		for i := int(f.spriteOffset); i < 8; i++ {
			f.fifo.Mix(i-int(f.spriteOffset), Pixel{f.spriteData[i], palette, bgOverObj})
		}
		f.state = f.oldState
	}
}

// ReadTileLine updates internal pixel buffer with LSB or MSB tile line
// depending on current state.
func (f *Fetcher) ReadTileLine(bitPlane uint8, tileDataAddr uint16, tileID uint8, signedID bool, tileLine uint8, flags uint8, data *[8]uint8) {
	var offset uint16

	if signedID {
		offset = uint16(int16(tileDataAddr) + int16(int8(tileID))*16)
	} else {
		offset = tileDataAddr + (uint16(tileID) * 16)
	}
	if flags&SpriteFlipY != 0 {
		// If flipping, get line at (spriteSize-1-line)
		tileLine = 7 - tileLine
	}
	addr := offset + (uint16(tileLine) * 2)

	pixelData := f.vRAM.Read(addr + uint16(bitPlane))
	for bitPos := 7; bitPos >= 0; bitPos-- {
		var pixelIndex uint
		if flags&SpriteFlipX != 0 {
			pixelIndex = uint(bitPos)
		} else {
			pixelIndex = 7 - uint(bitPos)
		}
		if bitPlane == 0 {
			// Least significant bit, replace previous value.
			data[pixelIndex] = (pixelData >> uint(bitPos)) & 1
		} else {
			// Most significant bit, update previous value.
			data[pixelIndex] |= ((pixelData >> uint(bitPos)) & 1) << 1
		}
	}
}
