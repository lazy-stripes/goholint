package ppu

import (
	"image"
	"image/color"

	"tigris.fr/gameboy/memory"
)

// PPU address space handling video RAM and display.
type PPU struct {
	vram *memory.RAM
}

// New PPU instance, randomized because why not?
func New() *PPU {
	return &PPU{memory.NewVRAM(0x8000, 8*1024)}
}

// Contains returns true if requested address is in VRAM space.
func (p *PPU) Contains(addr uint) bool {
	return p.vram.Contains(addr)
}

// Read returns the byte at the given address in VRAM. TODO: checks
func (p *PPU) Read(addr uint) uint8 {
	return p.vram.Read(addr)
}

// Write sets the byte at the given address in VRAM to the given value. TODO: checks
func (p *PPU) Write(addr uint, value uint8) {
	p.vram.Write(addr, value)
}

// DumpVRAM saves the current state of the video RAM as a PNG file.
func (p *PPU) DumpVRAM(addr uint) {

	// FIXME: handle native palettes
	palette := color.Palette{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0x40, 0x40, 0x40, 0xff},
		color.RGBA{0x80, 0x80, 0x80, 0xff},
		color.RGBA{0xC0, 0xC0, 0xC0, 0xff},
	}

	start := 0x8000
	dump := image.NewPaletted(image.Rect(0, 0, 128, 192), palette)
	offset := 0
	for tile := 0; tile < 8*1024; tile++ {
		pixels := p.Decode(uint(start + offset))
		for _, pixel := range pixels {
			dump.Pix[pixel] = pixel
			offset++
		}
	}
}

// Decode reads 8 pixels from VRAM and returns them as an array of colors (aka palette indexes). TODO: Fetcher.
func (p *PPU) Decode(addr uint) (line []uint8) {
	lineHi := p.Read(addr)
	lineLow := p.Read(addr)
	// TODO: push directly to FIFO
	line = make([]uint8, 0, 8)
	for bit := uint(7); bit >= 0; bit-- {
		pixel := (lineHi>>bit&1)<<1 | (lineLow >> bit & 1)
		line = append(line, pixel)
	}
	return line
}
