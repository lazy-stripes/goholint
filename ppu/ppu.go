package ppu

// Source: [TUGBT] https://www.youtube.com/watch?v=HyzD8pNlpwI&t=2747s

import (
	"bytes"
	"fmt"

	"github.com/lazy-stripes/goholint/interrupts"
	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/memory"
	"github.com/lazy-stripes/goholint/ppu/states"
	"github.com/lazy-stripes/goholint/screen"
	"github.com/veandco/go-sdl2/sdl"
)

// Package-wide logger.
var log = logger.New("ppu", "pixel processing unit operations")

// Package initialization function setting up logger submodules.
func init() {
	log.Add("ticks", "ticks taken per PPU phase (Desperate only)")
}

// ClockFactor representing the number of ticks taken by each step (base is 4).
// Used in Fetcher's Tick() method.
var ClockFactor = 2

// Register addresses.
const (
	AddrLCDC = 0xff40
	AddrSTAT = 0xff41
	AddrSCY  = 0xff42
	AddrSCX  = 0xff43
	AddrLY   = 0xff44
	AddrLYC  = 0xff45
	AddrBGP  = 0xff47
	AddrOBP0 = 0xff48
	AddrOBP1 = 0xff49
	AddrWY   = 0xff4a
	AddrWX   = 0xff4b
)

// LCDC flags. XXX: Move to subpackage lcdc for nicer namespacing?
const (
	// Bit 0 - BG/Window Display/Priority     (0=Off, 1=On)
	LCDCBGDisplay uint8 = 1 << iota
	// Bit 1 - OBJ (Sprite) Display Enable    (0=Off, 1=On)
	LCDCSpriteDisplayEnable
	// Bit 2 - OBJ (Sprite) Size              (0=8x8, 1=8x16)
	LCDCSpriteSize
	// Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
	LCDCBGTileMapDisplayeSelect
	// Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
	LCDCBGWindowTileDataSelect
	// Bit 5 - Window Display Enable          (0=Off, 1=On)
	LCDCWindowDisplayEnable
	// Bit 6 - Window Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
	LCDCWindowTileMapDisplayeSelect
	// Bit 7 - LCD Display Enable             (0=Off, 1=On)
	LCDCDisplayEnable
)

// PPU address space handling video RAM and display.
type PPU struct {
	*memory.MMU
	*Fetcher
	FIFO

	VRAM       *VRAM
	OAM        *OAM
	Interrupts *interrupts.Interrupts
	Cycle      int
	LCD        screen.Display
	LCDC       uint8
	STAT       uint8
	SCY, SCX   uint8
	LY         uint8
	LYC        uint8
	WY, WX     uint8
	BGP        uint8
	OBP0, OBP1 uint8

	ticks int
	state states.State

	toDrop uint8 // Pixels to drop for SCX
	x      uint8
	window bool // True if window fetch in progress

	// Quick and dirty mapping of PixelPalette index to palette register
	// for quick access when pushing pixels to LCD.
	palettes [3]*uint8

	frames uint // DEBUG for counting
}

// New PPU instance.
func New(display screen.Display) *PPU {
	p := PPU{MMU: memory.NewEmptyMMU(), LCD: display}
	p.Add(memory.Registers{
		AddrLCDC: &p.LCDC,
		AddrSTAT: &p.STAT,
		AddrSCY:  &p.SCY,
		AddrSCX:  &p.SCX,
		AddrLY:   &p.LY,
		AddrLYC:  &p.LYC,
		AddrBGP:  &p.BGP,
		AddrOBP0: &p.OBP0,
		AddrOBP1: &p.OBP1,
		AddrWY:   &p.WY,
		AddrWX:   &p.WX,
	})

	// FIXME: mode-dependent addressing for those.
	p.VRAM = NewVRAM(&p)
	p.Add(p.VRAM)

	p.OAM = NewOAM(&p)
	p.Add(p.OAM)

	//p.Fetcher = Fetcher{fifo: &p.FIFO, vRAM: p.MMU, lcdc: &p.LCDC}
	p.Fetcher = NewFetcher(&p)

	p.palettes = [3]*uint8{&p.BGP, &p.OBP0, &p.OBP1}
	p.state = states.OAMSearch

	return &p
}

// String returns a human-readable representation of the PPU's current state.
func (p *PPU) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "LCDC: %#02x\n", p.LCDC)
	fmt.Fprintf(&b, "STAT: %#02x\n", p.STAT)
	fmt.Fprintf(&b, "SCY:  %#02x\n", p.SCY)
	fmt.Fprintf(&b, "SCX:  %#02x\n", p.SCX)
	fmt.Fprintf(&b, "LY:   %#02x\n", p.LY)
	fmt.Fprintf(&b, "LYC:  %#02x\n", p.LYC)
	fmt.Fprintf(&b, "BGP:  %#02x\n", p.BGP)
	fmt.Fprintf(&b, "OBP0: %#02x\n", p.OBP0)
	fmt.Fprintf(&b, "OBP1: %#02x\n", p.OBP1)
	fmt.Fprintf(&b, "WY:   %#02x\n", p.WY)
	fmt.Fprintf(&b, "WX:   %#02x\n", p.WX)
	fmt.Fprintf(&b, "\nFrames: %d (%#04x)\n", p.frames, p.frames)
	return b.String()
}

// Requests the LY=LYC interrupt as needed when changing LY.
func (p *PPU) setLY(value uint8) {
	p.LY = value
	if p.LY == p.LYC {
		p.RequestLCDInterrupt(interrupts.STATLYCLY)
	}
}

// Write override that handles read-only registers and bits.
func (p *PPU) Write(addr uint16, value uint8) {
	switch addr {
	case AddrSTAT:
		log.Debugf("PPU.Write(0x%04x[STAT], 0x%02x)", addr, value)
		p.STAT = value & 0xf8
	case AddrLY:
		// [PANDOCS] says writing to it "resets counter"?
		log.Debugf("PPU.Write(0x%04x[LY], 0x%02x)", addr, value)
		log.Warning("Write to LY. What do?")
	default:
		p.MMU.Write(addr, value)
	}
}

// Read override that handles exact STAT lower bits' values at any given time.
func (p *PPU) Read(addr uint16) uint8 {
	if addr == AddrSTAT {
		// We never write to STAT bits 0-2 so we (safely?) assume they're 0.
		//logger.Printf("ppu", "PPU.Read(0x%04x[STAT]) - p.state=0x%02x, p.STAT=0x%02x", addr, p.state, p.STAT)
		stat := p.STAT | uint8(p.state) // Mode
		if p.LY == p.LYC {
			stat |= 4
		}
		log.Debugf("PPU.Read(0x%04x[STAT]) = 0x%02x", addr, stat)
		return stat
	}
	return p.MMU.Read(addr)
}

// Tick advances the CPU state one step. Return whether we reached VBlank so
// that event polling can happen then.
func (p *PPU) Tick() {
	p.Cycle++
	p.ticks++

	if !p.LCD.Enabled() {
		if p.LCDC&LCDCDisplayEnable == 0 {
			// Refresh window with "disabled screen" texture at about the same
			// rate we'd display the current texture upon VBlank.
			if p.ticks%(456*153) == 0 {
				log.Sub("ticks").Desperatef("Disabled: %d ticks", 456*153)
				p.LCD.VBlank()
			}
		} else {
			p.OAM.Start()
			p.LCD.Enable()
			p.state = states.OAMSearch
			p.RequestLCDInterrupt(interrupts.STATMode2)
		}
	} else {
		if p.LCDC&LCDCDisplayEnable == 0 {
			// Disable LCD. Clean up internal state.
			p.LY = 0
			p.x = 0
			// [TCAFBD] STAT mode flag is zero when LCD is off.
			p.state = 0
			p.LCD.Disable()
		}
	}

	if !p.LCD.Enabled() {
		return
	}

	switch p.state {
	case states.OAMSearch:
		// Tick will return true when all OAM space has been searched.
		if p.OAM.Tick() {
			// Initialize fetcher for background.
			y := p.SCY + p.LY
			tileLine := y % 8
			tileOffset := p.SCX / 8
			tileMapRowAddr := p.BGMap() + (uint16(y/8) * 32)
			tileDataAddr, signedID := p.TileData()
			p.Fetcher.Start(tileMapRowAddr, tileDataAddr, tileOffset, tileLine, signedID)

			p.x = 0
			p.toDrop = p.SCX % 8
			p.state = states.PixelTransfer

			log.Sub("ticks").Desperatef("OAM Search: %d ticks", p.ticks)
		}

	case states.PixelTransfer:
		p.Fetcher.Tick()

		// TODO: handle display mode
		if p.FIFO.Size() <= 8 {
			return
		}

		if p.Fetcher.state&states.FetchingSprite > 0 {
			return
		}

		// Check whether we should start fetching window tiles.
		if !p.window && p.LCDC&LCDCWindowDisplayEnable > 0 &&
			p.LY >= p.WY && p.x+7 >= p.WX {
			p.window = true
			p.toDrop = 0 // Window doesn't scroll

			// Reinitialize fetcher for window.
			y := p.LY - p.WY
			tileLine := y % 8
			tileOffset := (p.x - p.WX + 7) / 8
			tileMapRowAddr := p.WindowMap() + (uint16(y/8) * 32)
			tileDataAddr, signedID := p.TileData()
			p.Fetcher.Start(tileMapRowAddr, tileDataAddr, tileOffset, tileLine, signedID)
			return
		}

		// Drop pixels according to SCX
		if p.toDrop > 0 {
			p.toDrop -= p.Drop()
			return
		}

		// Find out if a sprite (that hasn't yet been fetched) should be
		// displayed at the current X position.
		if p.LCDC&LCDCSpriteDisplayEnable != 0 {
			for i, sprite := range p.OAM.Sprites {
				if sprite.Fetched {
					continue
				}

				// Fetch sprite if its X position matches the current pixel.
				// OAM search guarantees that all 10 potential sprites' X
				// position is non-zero.
				fetch := false
				offset := uint8(0)
				switch {
				case sprite.X < 8 && p.x == 0:
					// Special case for sprites scrolling in from the left.
					fetch = true
					offset = 8 - sprite.X
				case sprite.X-8 == p.x:
					fetch = true
				}

				if fetch {
					p.Fetcher.FetchSprite(sprite, offset, p.LY+16-sprite.Y)
					p.OAM.Sprites[i].Fetched = true
					return
				}
			}
		}

		p.x += p.Pop()
		if p.x == 160 {
			p.window = false
			p.LCD.HBlank()
			p.state = states.HBlank
			p.RequestLCDInterrupt(interrupts.STATMode0)

			log.Sub("ticks").Desperatef("Pixel Transfer: %d ticks", p.ticks)
		}

	case states.HBlank:
		// Simply wait the proper number of clock cycles.
		if p.ticks >= 456 {
			log.Sub("ticks").Desperatef("HBlank: %d ticks", p.ticks)

			// Done, either move to new line, or VBlank.
			p.ticks = 0
			p.setLY(p.LY + 1)
			if p.LY == 144 {
				p.frames++
				sdl.Do(p.LCD.VBlank) // Keep GPU stuff in OS thread.
				p.state = states.VBlank
				p.RequestLCDInterrupt(interrupts.STATMode1)

				p.Interrupts.Request(interrupts.VBlank)
			} else {
				// Prepare to go back to OAM search state.
				p.OAM.Start()
				p.state = states.OAMSearch
				p.RequestLCDInterrupt(interrupts.STATMode2)
			}
		}

	case states.VBlank:
		// Simply wait the proper number of clock cycles. Special case for last line.
		if p.ticks == 4 && p.LY == 153 {
			p.setLY(0)
		}

		if p.ticks >= 456 {
			log.Sub("ticks").Desperatef("VBlank: %d ticks (LY=%d)", p.ticks, p.LY)

			p.ticks = 0
			if p.LY == 0 { // We wrapped back to 0 about 452 ticks ago. Start rendering from top of screen again.
				p.OAM.Start()
				p.state = states.OAMSearch
				p.RequestLCDInterrupt(interrupts.STATMode2)
			} else {
				p.setLY(p.LY + 1)
			}
		}
	}
}

// mapAddress returns the base address for BG or Window map according to LCDC.
func (p *PPU) mapAddress(bit uint8) uint16 {
	if p.LCDC&bit != 0 {
		return 0x9c00
	}
	return 0x9800
}

// BGMap returns the base address of the background map in VRAM.
func (p *PPU) BGMap() uint16 {
	return p.mapAddress(LCDCBGTileMapDisplayeSelect)
}

// WindowMap returns the base address of the window map in VRAM.
func (p *PPU) WindowMap() uint16 {
	return p.mapAddress(LCDCWindowTileMapDisplayeSelect)
}

// TileData returns the base address of the background or window tile data in VRAM.
func (p *PPU) TileData() (addr uint16, signedID bool) {
	if (p.LCDC & LCDCBGWindowTileDataSelect) != 0 {
		return 0x8000, false
	}
	return 0x9000, true
}

// Pop tries shifting a pixel out of the FIFO to the LCD and returns the
// number of shifted pixels (0 or 1).
func (p *PPU) Pop() uint8 {
	return p.pop(false)
}

// Drop tries taking a pixel out of the FIFO and discarding it to account for
// SCX. It returns the number of dropped pixels (0 or 1).
func (p *PPU) Drop() uint8 {
	return p.pop(true)
}

func (p *PPU) pop(drop bool) uint8 {
	if pixel, err := p.FIFO.Pop(); err == nil {
		if !drop {
			palette := *p.palettes[pixel.Palette]
			// This was shamefully taken from coffee-gb.
			color := (palette >> (pixel.Color << 1)) & 3
			p.LCD.Write(color)
		}
		return 1
	}
	return 0
}

// RequestLCDInterrupt checks STAT bits when an interrupt condition occurs and
// requests an actual interrupt if the corresponding bit is set.
func (p *PPU) RequestLCDInterrupt(interrupt uint8) {
	if p.STAT&interrupt != 0 {
		p.Interrupts.Request(interrupts.LCDStat)
	}
}
