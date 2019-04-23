package cpu

import (
	"bytes"
	"fmt"
	"os"
	"runtime/pprof"

	"go.tigris.fr/gameboy/cpu/states"
	"go.tigris.fr/gameboy/interrupts"
	"go.tigris.fr/gameboy/memory"
)

// Flag bitfield enum
const (
	FlagC uint8 = 1 << (iota + 4)
	FlagH
	FlagN
	FlagZ
)

// A CPU implementation of the DMG-01's
type CPU struct {
	MMU    memory.Addressable
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
	ticks       uint
	state       int
	interrupt   uint8  // Currently requested interrupt
	temp8       uint8  // Internal work register storing 8-bit micro-operation results
	temp16      uint16 // Internal work register storing 16-bit micro-operation results

	debug     bool
	startFrom uint16
	oldPC     uint16
}

// New CPU running code in the given address space starting from 0.
func New(code memory.Addressable) *CPU {
	return &CPU{MMU: code, state: states.FetchOpCode, startFrom: 0xFFFF}
}

// Tick advances the CPU state one step.
func (c *CPU) Tick() {
	c.Cycle++
	c.ticks++
	if c.ticks < 4 { // FIXME: c.ClockFactor
		return
	}

	// Reset tick counter and execute next state
	c.ticks = 0

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
			fmt.Printf("PC=%04X (%02X)\n", c.PC, c.MMU.Read(c.PC))
		}
		opcode := c.NextByte()

		if opcode == 0xcb { // Extended instruction set
			c.state = states.FetchExtendedOpcode
		} else {
			defer instructionError(c, false)

			c.instruction = LR35902InstructionSet[opcode]
			if c.instruction.Execute(c) { // Instruction is done within the first 4 cycles.
				c.state = states.FetchOpCode
			} else {
				c.state = states.Execute
			}
		}

	case states.FetchExtendedOpcode:
		opcode := c.NextByte()
		defer instructionError(c, true)

		c.instruction = LR35902ExtendedInstructionSet[opcode]
		if c.instruction.Execute(c) { // Instruction is done within the first 8 cycles.
			c.state = states.FetchOpCode
		} else {
			c.state = states.Execute
		}

	case states.Execute:
		if c.instruction.Tick() {
			c.state = states.FetchOpCode
		}

	case states.InterruptWait0:
		// [TCAGBD:4.9] mentions a 2-cycle idle upon handling interrupt request.
		c.state = states.InterruptWait1

	case states.InterruptWait1:
		requested := c.IF & c.IE

		// Doing this in a switch/case instead of a loop because I played too much EXAPUNKS...
		// Unrolling is good for perfs, right?
		switch {
		case requested&interrupts.VBlank != 0:
			c.interrupt = interrupts.VBlank
		case requested&interrupts.LCDStat != 0:
			c.interrupt = interrupts.LCDStat
		case requested&interrupts.Timer != 0:
			c.interrupt = interrupts.Timer
			// TODO: all other interrupts
		default:
			fmt.Printf(" !!! Unimplemented interrupt requested: %02x\n", requested)
		}

		c.state = states.InterruptPushPCHigh

	case states.InterruptPushPCHigh:
		c.SP--
		c.MMU.Write(c.SP, uint8(c.PC>>8))
		c.state = states.InterruptPushPCLow

	case states.InterruptPushPCLow:
		c.SP--
		c.MMU.Write(c.SP, uint8(c.PC&0xff))
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
	fmt.Fprintf(&b, "Flags:\nZ: %d - N: %d - H: %d - C: %d\n\n", c.F&FlagZ>>7, c.F&FlagN>>6, c.F&FlagH>>5, c.F&FlagC>>4)
	fmt.Fprintf(&b, "Cycle: %d\n", c.Cycle)
	return b.String()
}

// NextByte returns the next byte pointed to by PC.
func (c *CPU) NextByte() uint8 {
	value := c.MMU.Read(c.PC)
	c.PC++
	return value
}

// NextWord returns the next 16bit value in proper byte order you'd expect.
func (c *CPU) NextWord() uint16 {
	return uint16(c.NextByte()) | uint16(c.NextByte())<<8
}

// Context returns a printable context to prepend to log messages.
func (c *CPU) Context() string {
	return fmt.Sprintf("[PC=%04x] ", c.PC)
}

// For missing opcodses debugz.
func instructionError(c *CPU, extended bool) {
	if r := recover(); r != nil {
		if extended {
			fmt.Printf("Execute error at extended instruction %#04x (0xCB %#02x) (%v)\n", c.PC-2, c.MMU.Read(c.PC-1), r)
		} else {
			fmt.Printf("Execute error at instruction %#04x (%#02x) (%v)\n", c.PC-1, c.MMU.Read(c.PC-1), r)
		}
		fmt.Printf("CPU's final state:\n%s\n", c)
		// Dump memory
		if f, err := os.Create("ram-dump.bin"); err == nil {
			defer func() {
				f.Close()
			}()
			buf := make([]byte, 1, 1)
			for addr := uint16(0); addr <= 0xffff; addr++ {
				buf[0] = c.MMU.Read(addr)
				f.Write(buf)
			}
			fmt.Println("RAM dumped to ram-dump.bin")
		}

		// Manually stop profile here, since the Exit below will shortcut the deferred call in main.
		pprof.StopCPUProfile()
		os.Exit(255)
	}
}
