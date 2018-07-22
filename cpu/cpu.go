package cpu

import (
	"bytes"
	"fmt"
	"os"

	"tigris.fr/gameboy/memory"
	"tigris.fr/gameboy/ppu"
	"tigris.fr/gameboy/timer"
)

// A CPU implementation of the DMG-01's
type CPU struct {
	timer.Clock
	MMU                    *memory.MMU
	Cycle                  uint
	IME                    bool // Interrupt Master Enable flag
	A, F, B, C, D, E, H, L uint8
	SP                     uint16
	PC                     uint16
}

// New CPU running DMG code in the given address space starting from 0.
func New(mmu *memory.MMU) *CPU {
	return &CPU{Clock: make(timer.Clock), MMU: mmu}
}
// Flag bitfield enum
const (
	FlagC uint8 = 1 << (iota + 4)
	FlagH
	FlagN
	FlagZ
)

// Helper methods to read/write 16-bit registers
func readRR(high, low byte) uint16 {
	return uint16(high)<<8 | uint16(low)
}

func writeRR(value uint16, high, low *byte) {
	*high = byte(value >> 8)
	*low = byte(value & 0x00ff)
}

// AF returns the 16-bit value stored in registers A and F.
func (c *CPU) AF() uint16 {
	return readRR(c.A, c.F)
}

// SetAF writes the bytes of the given 16-bit value to A and F.
func (c *CPU) SetAF(word uint16) {
	writeRR(word, &c.A, &c.F)
}

// BC returns the 16-bit value stored in registers B and C.
func (c *CPU) BC() uint16 {
	return readRR(c.B, c.C)
}

// SetBC writes the bytes of the given 16-bit value to B and C.
func (c *CPU) SetBC(word uint16) {
	writeRR(word, &c.B, &c.C)
}

// DE returns the 16-bit value stored in registers D and E.
func (c *CPU) DE() uint16 {
	return readRR(c.D, c.E)
}

// SetDE writes the bytes of the given 16-bit value to D and E.
func (c *CPU) SetDE(word uint16) {
	writeRR(word, &c.D, &c.E)
}

// HL returns the 16-bit value stored in registers H and L.
func (c *CPU) HL() uint16 {
	return readRR(c.H, c.L)
}

// SetHL writes the bytes of the given 16-bit value to H and L.
func (c *CPU) SetHL(word uint16) {
	writeRR(word, &c.H, &c.L)
}

// String returns a human-readable representation of the CPU's current state.
func (c *CPU) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "A: %#02x - F: %#02x - AF: %#04x\n", c.A, c.F, c.AF())
	fmt.Fprintf(&b, "B: %#02x - C: %#02x - BC: %#04x\n", c.B, c.C, c.BC())
	fmt.Fprintf(&b, "D: %#02x - E: %#02x - DE: %#04x\n", c.D, c.E, c.DE())
	fmt.Fprintf(&b, "H: %#02x - L: %#02x - HL: %#04x\n", c.H, c.L, c.HL())
	fmt.Fprintf(&b, "                    SP: %#04x\n", c.SP)
	fmt.Fprintf(&b, "                    PC: %#04x\n", c.PC)
	fmt.Fprintf(&b, "Flags:\nZ: %d - N: %d - H: %d - C: %d\n", c.F&FlagZ>>7, c.F&FlagN>>6, c.F&FlagH>>5, c.F&FlagC>>4)
	fmt.Fprintf(&b, "Cycle: %d\n", c.Cycle)
	return b.String()
}

// Read a byte from MMU in the proper number of cycles.
func (c *CPU) Read(addr uint) uint8 {
	c.Ticks(2)
	return c.MMU.Read(addr)
}

// Write a byte to MMU in the proper number of cycles.
func (c *CPU) Write(addr uint, value byte) {
	c.Ticks(2)
	c.MMU.Write(addr, value)
}

// NextByte returns the next byte pointed to by PC.
func (c *CPU) NextByte() uint8 {
	value := c.Read(uint(c.PC))
	c.PC++
	return value
}

// NextWord returns the next 16bit value in proper byte order you'd expect.
func (c *CPU) NextWord() uint16 {
	return uint16(c.NextByte()) | uint16(c.NextByte())<<8
}

// For missing opcodses debugz.
func instructionError(c *CPU, extended bool) {
	if r := recover(); r != nil {
		if extended {
			fmt.Printf("Execute error at extended instruction %#04x (0xCB %#02x) (%v)\n", c.PC-2, c.MMU.Read(uint(c.PC-1)), r)
		} else {
			fmt.Printf("Execute error at instruction %#04x (%#02x) (%v)\n", c.PC-1, c.MMU.Read(uint(c.PC-1)), r)
		}
		fmt.Printf("CPU's final state:\n%s\n", c)
		// Dump memory
		/*
			if f, err := os.Create("ram-dump.bin"); err == nil {
				defer func() {
					f.Close()
				}()
				buf := make([]byte, 1, 1)
				for addr := uint(0); addr < 0x10000; addr++ {
					buf[0] = c.MMU.Read(addr)
					f.Write(buf)
				}
				fmt.Println("RAM dumped.")
			}*/

		os.Exit(255)
	}
}

// Execute the next instruction (and handles extensions to base instruction set.)
func (c *CPU) Execute() {
	opcode := c.NextByte()
	if opcode == 0xcb { // Extended instruction set
		defer instructionError(c, true)
		LR35902ExtendedInstructionSet[c.NextByte()](c)
	} else {
		defer instructionError(c, false)
		LR35902InstructionSet[opcode](c)
	}
}

// Run CPU on the current address space.
func (c *CPU) Run() {
	debugFrom := uint(28816)
	for {
		if c.PC == 0x8c && c.D == 63 {
			c.MMU.Spaces[1].(*ppu.PPU).LCD.Enable()
		}
		c.Execute()
		c.Cycle++
		if c.Cycle >= debugFrom {
			fmt.Printf("========= Cycle: %#4x ========\n", c.Cycle)
			fmt.Print(c)
			fmt.Printf("==============================\n")
		}
		if c.PC > 0x100 {
			fmt.Print("Jumped out of BootROM!")
			os.Exit(1)
		}
	}
}
