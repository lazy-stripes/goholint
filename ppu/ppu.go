package ppu

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"

	"tigris.fr/gameboy/cpu"
	"tigris.fr/gameboy/memory"
)

// PPU address space handling video RAM and display.
type PPU struct {
	*memory.MMU
	vram       *memory.RAM
	lcd        *LCD
	registers  map[uint]*uint8
	Clock      cpu.Clock
	LCDC       uint8
	STAT       uint8
	SCY, SCX   uint8
	LY         uint8
	LYC        uint8
	WY, WX     uint8
	BGP        uint8
	OBP0, OBP1 uint8
	// TODO: DMA, address space to OAM, put in CPU
}

// New PPU instance, randomized because why not?
func New() *PPU {
	p := PPU{vram: memory.NewVRAM(0x8000, 0x2000)}
	// Map addresses to their corresponding register.
	var registers = map[uint]*uint8{
		0xff40: &p.LCDC,
		0xff41: &p.STAT,
		0xff42: &p.SCY,
		0xff43: &p.SCX,
		0xff44: &p.LY,
		0xff45: &p.LYC,
		0xff47: &p.BGP,
		0xff48: &p.OBP0,
		0xff49: &p.OBP1,
		0xff4a: &p.WY,
		0xff4b: &p.WX,
	}
	p.registers = registers
	return &p
}

// Contains returns true if requested address is in VRAM or a register.
func (p *PPU) Contains(addr uint) bool {
	if p.registers[addr] != nil {
		return true
	}
	return p.vram.Contains(addr)
}

// Read returns the byte at the given address in VRAM or from register.
func (p *PPU) Read(addr uint) uint8 {
	if regPtr := p.registers[addr]; regPtr != nil {
		return *regPtr
	}
	return p.vram.Read(addr)
}

// Write sets the byte at the given address in VRAM to the given value. TODO: checks
func (p *PPU) Write(addr uint, value uint8) {
	// FIXME: check for R/O registers.
	if regPtr := p.registers[addr]; regPtr != nil {
		*regPtr = value

		// TODO: stuff on some registers?

		// XXX: dump tiles when LCD is enabled
		if addr == 0xff40 {
			p.DumpTiles(0x8190, 8)
			fmt.Println("Tiles dumped.")
			os.Exit(0)
		}
	} else {
		p.vram.Write(addr, value)
	}
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
	ioutil.WriteFile("vram-dump.bin", p.vram.Bytes, 0666)
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
