package ppu

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"

	"go.tigris.fr/gameboy/debug"
	"go.tigris.fr/gameboy/fifo"
	"go.tigris.fr/gameboy/interrupts"
	"go.tigris.fr/gameboy/lcd"
	"go.tigris.fr/gameboy/memory"
	"go.tigris.fr/gameboy/ppu/states"
)

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
	*fifo.FIFO
	Fetcher
	Interrupts *interrupts.Interrupts
	Cycle      int
	LCD        lcd.Display
	LCDC       uint8
	STAT       uint8
	SCY, SCX   uint8
	LY         uint8
	LYC        uint8
	WY, WX     uint8
	BGP        uint8
	OBP0, OBP1 uint8
	// TODO: address space to OAM, put in CPU

	ticks int
	state states.State

	oamIndex int

	x uint
}

// New PPU instance.
func New(display lcd.Display) *PPU {
	fifo := fifo.New(16, 8)
	p := PPU{MMU: memory.NewEmptyMMU(), FIFO: fifo, LCD: display}
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
	p.Add(memory.NewVRAM(0x8000, 0x2000)) // VRAM
	p.Add(memory.NewVRAM(0xfe00, 0xa0))   // OAM RAM (TODO: mapped OBJ struct)
	p.Fetcher = Fetcher{fifo: fifo, vRAM: p.MMU}
	p.state = states.OAMSearch
	return &p
}

// Write override that handles read-only registers and bits.
func (p *PPU) Write(addr uint, value uint8) {
	switch addr {
	case AddrSTAT:
		debug.Printf("ppu", "PPU.Write(0x%04x[STAT], 0x%02x)", addr, value)
		p.STAT = value & 0xf8
	case AddrLY:
		// [PANDOCS] says writing to it "resets counter"?
		debug.Printf("ppu", "PPU.Write(0x%04x[LY], 0x%02x)", addr, value)
		debug.Printf("ppu", "Write to LY. What do?")
	default:
		p.MMU.Write(addr, value)
	}
}

// Read override that handles exact STAT lower bits' values at any given time.
func (p *PPU) Read(addr uint) uint8 {
	if addr == AddrSTAT {
		// We never write to STAT bits 0-2 so we (safely?) assume they're 0.
		stat := p.STAT | uint8(p.state) // Mode
		if p.LY == p.LYC {
			stat |= 4
		}
		return stat
	}
	return p.MMU.Read(addr)
}

// Tick advances the CPU state one step.
func (p *PPU) Tick() {
	p.Cycle++
	p.ticks++

	if !p.LCD.Enabled() {
		if p.LCDC&LCDCDisplayEnable == 0 {
			// Refresh window with "disabled screen" texture at about the same
			// rate we'd display the current texture upon VBlank.
			if p.ticks%(456*153) == 0 {
				p.LCD.Blank()
			}
		} else {
			p.state = states.OAMSearch
			p.LCD.Enable()
		}
	} else {
		if p.LCDC&LCDCDisplayEnable == 0 {
			// Disable LCD. Clean up internal state.
			p.LY = 0
			p.x = 0
			// STAT mode flag is zero when LCD is disabled. Apparently.
			p.state = 0
			p.LCD.Disable()
		}
	}

	if !p.LCD.Enabled() {
		return
	}

	switch p.state {
	case states.OAMSearch:
		// TODO
		p.oamIndex++
		if p.oamIndex >= 40 {
			// Initialize fetcher for background.
			y := p.SCY + p.LY
			tileLine := y % 8
			tileOffset := p.SCX / 8
			tileMapRowAddr := p.BGMap() + (uint(y/8) * 32)
			tileDataAddr, signedID := p.TileData()
			p.Fetcher.Start(tileMapRowAddr, tileDataAddr, tileOffset, tileLine, signedID)

			p.x = 0
			p.state = states.PixelTransfer
		}

	case states.PixelTransfer:
		p.Fetcher.Tick()
		// TODO: handle display mode
		// TODO: drop pixels according to SCX
		// TODO: sprites display
		if p.FIFO.Size() <= 8 {
			return
		}

		p.x += p.Pop()
		if p.x == 160 {
			p.LCD.HBlank()
			p.state = states.HBlank
		}

	case states.HBlank:
		// Simply wait the proper number of clock cycles.
		if p.ticks >= 456 {
			// Done, either move to new line, or VBlank.
			p.ticks = 0
			p.LY++
			if p.LY == 144 {
				p.LCD.VBlank()
				p.Interrupts.Request(interrupts.VBlank)

				p.state = states.VBlank
			} else {
				// Prepare to go back to OAM search state.
				p.oamIndex = 0
				p.state = states.OAMSearch
			}
		}

	case states.VBlank:
		// Simply wait the proper number of clock cycles. Special case for last line.
		if p.ticks == 4 && p.LY == 153 {
			p.LY = 0
			// Request interrupt. Maybe add a hook to LY setter?
		}

		if p.ticks >= 456 {
			p.ticks = 0
			if p.LY == 0 { // We wrapped back to 0 about 452 ticks ago. Start rendering from top of screen again.
				p.oamIndex = 0
				p.state = states.OAMSearch
				// TODO: interrupts
			} else {
				p.LY++
			}
			// TODO: LYC=LY interrupt
		}
	}
}

// BGMap returns the base address of the background map in VRAM.
func (p *PPU) BGMap() uint {
	if (p.LCDC & LCDCBGTileMapDisplayeSelect) > 0 {
		return 0x9c00
	}
	return 0x9800
}

// TileData returns the base address of the background or window tile data in VRAM.
func (p *PPU) TileData() (addr uint, signedID bool) {
	if (p.LCDC & LCDCBGWindowTileDataSelect) > 0 {
		return 0x8000, false
	}
	return 0x9000, true
}

// Pop tries shifting a pixel out of the FIFO and returns the number of shifted pixels (0 or 1).
func (p *PPU) Pop() uint {
	if pixel, err := p.FIFO.Pop(); err == nil {
		// This was shamefully taken from coffee-gb.
		color := (p.BGP >> (pixel.(uint8) << 1)) & 3
		p.LCD.Write(lcd.Pixel(color))
		return 1
	}
	return 0
}

// DumpTiles writes tiles from VRAM into a PNG file to test the decoder.
func (p *PPU) DumpTiles(addr, len uint) {

	// FIXME: handle native palettes
	palette := color.Palette{
		color.RGBA{0xff, 0xff, 0xff, 0xff},
		color.RGBA{0xaa, 0xaa, 0xaa, 0xff},
		color.RGBA{0x55, 0x55, 0x55, 0xff},
		color.RGBA{0x00, 0x00, 0x00, 0xff},
	}

	start := addr
	// Don't bother re-aligning tile lines yet, use an 8-pixels wide image.
	dump := image.NewPaletted(image.Rect(0, 0, 8, int(8*len)), palette)
	offset := 0
	for tile := 0; tile < int(len); tile++ {
		for line := 0; line < 8; line++ {
			pixels := p.Decode(start)
			for _, pixel := range pixels {
				dump.Pix[offset] = pixel
				offset++
			}
			start += 2 // 2 bytes per tile line
		}
	}

	f, err := os.Create("tiles-dump.png")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(f)
	png.Encode(w, dump)
	w.Flush()

	// Dump VRAM for checks
	//ioutil.WriteFile("vram-dump.bin", p.vram.Bytes, 0666)
}

// Decode reads 8 pixels from VRAM and returns them as an array of colors (aka palette indexes). TODO: Fetcher.
func (p *PPU) Decode(addr uint) (line []uint8) {
	lineLo := p.Read(addr)
	lineHi := p.Read(addr + 1)
	// TODO: push directly to FIFO
	line = make([]uint8, 0, 8)
	for bit := 7; bit >= 0; bit-- {
		pixel := (lineHi>>uint(bit)&1)<<1 | (lineLo >> uint(bit) & 1)
		line = append(line, pixel)
	}
	return line
}
