package cpu

import (
	"bytes"
	"fmt"

	"github.com/lazy-stripes/goholint/cpu/states"
	"github.com/lazy-stripes/goholint/interrupts"
	"github.com/lazy-stripes/goholint/memory"
	"github.com/lazy-stripes/goholint/options"
)

// [GEKKIO] https://gekkio.fi/files/gb-docs/gbctr.pdf

// Flag bitfield enum
const (
	FlagC uint8 = 1 << (iota + 4)
	FlagH
	FlagN
	FlagZ
)

// CPU emulates a DMG-01 CPU.
type CPU struct {
	Memory memory.Addressable
	Cycle  uint
	IME    bool // Interrupt Master Enable flag
	IF, IE uint8
	A, F   uint8
	B, C   uint8
	D, E   uint8
	H, L   uint8
	SP     uint16
	PC     uint16

	instruction Instruction
	state       int

	// [GEKKIO] says EI takes one instruction to actually set IME.
	IMEScheduled bool
	IMEPending   bool

	interrupt uint8  // Currently requested interrupt
	temp8     uint8  // Internal work register storing 8-bit micro-operation results
	temp16    uint16 // Internal work register storing 16-bit micro-operation results

	debug     bool
	startFrom uint16
	oldPC     uint16
}

// New CPU running code in the given address space starting from 0.
func New(mem memory.Addressable) *CPU {
	return &CPU{Memory: mem, state: states.FetchOpCode, startFrom: 0xFFFF}
}

// Tick advances the CPU state one step.
func (c *CPU) Tick() {
	c.Cycle++

	// Handle interrupts
	if (c.state&states.Interruptible != 0) && c.IME && (c.IF&c.IE != 0) {
		// TODO: re-enable LCD if interrupted after STOP.
		c.state = states.InterruptWait0
	}

	// Exit HALT even if IME is not set
	if (c.state == states.Halted) && (c.IF&c.IE != 0) {
		c.state = states.FetchOpCode
	}

	switch c.state {
	case states.Halted:
		return
	case states.Stopped:
		return
	case states.FetchOpCode:
		if !c.debug && c.PC == c.startFrom {
			c.debug = true
		}
		if c.debug && c.PC != c.oldPC {
			//fmt.Printf("PC=%04X (%02X)\n", c.PC, c.MMU.Read(c.PC))
		}
		opcode := c.NextByte()

		if opcode == 0xcb { // Extended instruction set
			c.state = states.FetchExtendedOpcode
		} else {
			c.instruction = LR35902InstructionSet[opcode]
			if c.instruction.Execute(c) { // Instruction is done within the first 4 cycles.
				c.state = states.FetchOpCode
			} else {
				c.state = states.Execute
			}
		}

	case states.FetchExtendedOpcode:
		opcode := c.NextByte()

		c.instruction = LR35902ExtendedInstructionSet[opcode]
		if c.instruction.Execute(c) { // Instruction is done within the first 8 cycles.
			c.state = states.FetchOpCode
		} else {
			c.state = states.Execute
		}

	case states.Execute:
		if c.instruction.Tick() {
			// Handle one-instruction delay when enabling IME [GEKKIO]
			if c.IMEScheduled {
				if c.IMEPending {
					c.IMEPending = false
				} else {
					c.IME = true
					c.IMEScheduled = false
				}
			}
			c.state = states.FetchOpCode
		}

	case states.InterruptWait0:
		// [TCAGBD:4.9] mentions a 2-cycle idle upon handling interrupt request.
		c.state = states.InterruptWait1

	case states.InterruptWait1:
		requested := c.IF & c.IE

		// Doing this in a switch/case instead of a loop because I played too
		// much EXAPUNKS... Unrolling is good for perfs, right?
		switch {
		case requested&interrupts.VBlank != 0:
			c.interrupt = interrupts.VBlank
		case requested&interrupts.LCDStat != 0:
			c.interrupt = interrupts.LCDStat
		case requested&interrupts.Timer != 0:
			c.interrupt = interrupts.Timer
		case requested&interrupts.Serial != 0:
			c.interrupt = interrupts.Serial
		case requested&interrupts.Joypad != 0:
			c.interrupt = interrupts.Joypad
		default:
			fmt.Printf(" !!! Unknown interrupt requested: 0x%02x\n", requested)
		}

		c.state = states.InterruptPushPCHigh

	case states.InterruptPushPCHigh:
		c.SP--
		c.Memory.Write(c.SP, uint8(c.PC>>8))
		c.state = states.InterruptPushPCLow

	case states.InterruptPushPCLow:
		c.SP--
		c.Memory.Write(c.SP, uint8(c.PC&0xff))
		c.state = states.InterruptCall

	case states.InterruptCall:
		c.PC = interrupts.InterruptAddress[c.interrupt]
		c.IME = false
		c.IF &= ^c.interrupt
		c.state = states.FetchOpCode

	default:
		panic(fmt.Sprintf("Unknown CPU state %d", c.state))
	}
}

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
	fmt.Fprintf(&b, "Flags:\nZ: %d - N: %d - H: %d - C: %d\n\n", c.F&FlagZ>>7,
		c.F&FlagN>>6, c.F&FlagH>>5, c.F&FlagC>>4)
	fmt.Fprintf(&b, "Cycle: %d\n", c.Cycle)
	return b.String()
}

// NextByte returns the next byte pointed to by PC.
func (c *CPU) NextByte() uint8 {
	value := c.Memory.Read(c.PC)
	c.PC++
	return value
}

// NextWord returns the next 16bit value in proper byte order you'd expect.
func (c *CPU) NextWord() uint16 {
	return uint16(c.NextByte()) | uint16(c.NextByte())<<8
}

// Context returns a printable context to prepend to log messages. Currently,
// it only shows the current value of PC.
func (c *CPU) Context() string {
	return fmt.Sprintf("[PC=%04x, Cycle=%08x] ", c.PC, c.Cycle)
}

// DumpMemory writes current RAM values to a file. TODO: make filename configurable.
func (c *CPU) DumpMemory() {
	suffix := fmt.Sprintf("-%d.memory", c.Cycle)
	if f, err := options.CreateFileIn("debug", suffix); err == nil {
		defer func() {
			f.Close()
		}()
		buf := make([]byte, 0xffff)
		for addr := uint16(0); addr <= 0xffff; addr++ {
			buf[addr] = c.Memory.Read(addr)
		}
		f.Write(buf)
		fmt.Printf("Memory dumped to %s\n", f.Name())
	}
}
