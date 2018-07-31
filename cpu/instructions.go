package cpu

import (
	"fmt"
)

// An Operation executed on the CPU as part of an instruction.
type Operation func(c *CPU)

// An Instruction to be executed by a CPU, pushing a number of 4-cycle micro-operations to the CPU.
type Instruction func(c *CPU) (done bool)

// LR35902InstructionSet is an array of opcodes for the DMG CPU.
var LR35902InstructionSet = []Instruction{
	0x00: nop,
	0x01: ldBcD16,
	0x03: incBc,
	0x04: incB,
	0x05: decB,
	0x06: ldBD8,
	0x0b: decBc,
	0x0c: incC,
	0x0d: decC,
	0x0e: ldCD8,
	0x11: ldDeD16,
	0x13: incDe,
	0x14: incD,
	0x15: decD,
	0x16: ldDD8,
	0x17: rlA,
	0x18: jrR8,
	0x1a: ldAAddrDe,
	0x1c: incE,
	0x1d: decE,
	0x1e: ldED8,
	0x20: jrNzR8,
	0x21: ldHlD16,
	0x22: ldiAddrHlA,
	0x23: incHl,
	0x24: incH,
	0x25: decH,
	0x26: ldHD8,
	0x28: jrZR8,
	0x2a: ldiAAddrHl,
	0x2c: incL,
	0x2d: decL,
	0x2e: ldLD8,
	0x2f: cpl,
	0x31: ldSpD16,
	0x32: lddHlA,
	0x33: incSp,
	0x34: incAddrHl,
	0x35: decAddrHl,
	0x36: ldAddrHlD8,
	0x3c: incA,
	0x3d: decA,
	0x3e: ldAD8,
	0x40: ldBB,
	0x41: ldBC,
	0x42: ldBD,
	0x43: ldBE,
	0x44: ldBH,
	0x45: ldBL,
	0x46: ldBAddrHl,
	0x47: ldBA,
	0x48: ldCB,
	0x49: ldCC,
	0x4a: ldCD,
	0x4b: ldCE,
	0x4c: ldCH,
	0x4d: ldCL,
	0x4e: ldCAddrHl,
	0x4f: ldCA,
	0x57: ldDA,
	0x60: ldHB,
	0x61: ldHC,
	0x62: ldHD,
	0x63: ldHE,
	0x64: ldHH,
	0x65: ldHL,
	0x66: ldHAddrHl,
	0x67: ldHA,
	0x68: ldLB,
	0x69: ldLC,
	0x6a: ldLD,
	0x6b: ldLE,
	0x6c: ldLH,
	0x6d: ldLL,
	0x6e: ldLAddrHl,
	0x6f: ldLA,
	/*
		0x70: ldAddrHlB,
		0x71: ldAddrHlC,
		0x72: ldAddrHlD,
		0x73: ldAddrHlE,
		0x74: ldAddrHlH,
		0x75: ldAddrHlL,
	*/
	0x77: ldAddrHlA,
	0x78: ldAB,
	0x79: ldAC,
	0x7a: ldAD,
	0x7b: ldAE,
	0x7c: ldAH,
	0x7d: ldAL,
	0x7e: ldAAddrHl,
	0x7f: ldAA,
	0x80: addAB,
	0x81: addAC,
	0x82: addAD,
	0x83: addAE,
	0x84: addAH,
	0x85: addAL,
	0x86: addAAddrHl,
	0x87: addAA,
	0x90: subB,
	0x91: subC,
	0x92: subD,
	0x93: subE,
	0x94: subH,
	0x95: subL,
	/*
		0x96: subAddrHl,
		0x97: subA,
			0x98: sbcAB,
			0x99: sbcAC,
			0x9a: sbcAD,
			0x9b: sbcAE,
			0x9c: sbcAH,
			0x9d: sbcAL,
			0x9e: sbcAAddrHl,
			0x9f: sbcAA,
	*/
	0xa0: andB,
	0xa1: andC,
	0xa2: andD,
	0xa3: andE,
	0xa4: andH,
	0xa5: andL,
	0xa6: andAddrHl,
	0xa7: andA,
	0xa8: xorB,
	0xa9: xorC,
	0xaa: xorD,
	0xab: xorE,
	0xac: xorH,
	0xad: xorL,
	0xae: xorAddrHl,
	0xaf: xorA,
	0xb0: orB,
	0xb1: orC,
	0xb2: orD,
	0xb3: orE,
	0xb4: orH,
	0xb5: orL,
	0xb6: orAddrHl,
	0xb7: orA,
	0xbe: cpHl,
	0xc1: popBc,
	0xc3: jpA16,
	0xc5: pushBc,
	0xc9: ret,
	0xcd: callA16,
	/*
		0xd1: popDe,
		0xd5: pushDe,
	*/
	0xe0: ldAddrFfA8A,
	/*
		0xe1: popHl,
	*/
	0xe2: ldAddrFfCA,
	/*
		0xe5: pushHl,
	*/
	0xe6: andD8,
	0xea: ldAddrA16A,
	0xef: rst28h,
	0xf0: ldAAddrFfA8,
	0xf1: popAf,
	0xf3: di,
	0xf5: pushAf,
	0xfa: ldAAddrA16,
	0xfb: ei,
	0xfe: cpD8,
}

// LR35902ExtendedInstructionSet is the array of extension opcodes for the DMG CPU.
var LR35902ExtendedInstructionSet = []Instruction{
	0x11: rlC,
	0x37: swapA,
	0x7c: bit7H,
}

// Operations. The pseudo-atomic things the CPU does as part as an Instruction, which might take many cycles.
// All these helper functions return an Operation instance to be pushed to a CPU.

// Read next 8-bit argument into destination address.
func opReadD8(c *CPU, dest *byte) Operation {
	return func(c *CPU) {
		*dest = c.NextByte()
	}
}

// Read 8-bit value from memory into destination address.
func opReadD8At(c *CPU, addr uint, dest *byte) Operation {
	return func(c *CPU) {
		*dest = c.MMU.Read(addr)
	}
}

// Read 8-bit value from memory address (evaluated at run-time) into destination address.
// XXX: I'm not happy with the naming, nor the uint type chagne, nor the fact I need the two function variants. x_x
func opReadD8AtAddr(c *CPU, addr *uint16, dest *byte) Operation {
	return func(c *CPU) {
		*dest = c.MMU.Read(uint(*addr))
	}
}

// Read least significant byte of 16-bits argument into 16-bit destination. Most significant byte is reset.
func opReadD16Low(c *CPU, dest *uint16) Operation {
	return func(c *CPU) {
		*dest = uint16(c.NextByte())
	}
}

// Read most significant byte of 16-bits argument into 16-bit destination.
func opReadD16High(c *CPU, dest *uint16) Operation {
	return func(c *CPU) {
		*dest |= uint16(c.NextByte()) << 8
	}
}

// Read least significant byte of 16-bits value from memory into 16-bit destination. Most significant byte is reset.
func opReadD16LowAt(c *CPU, addr uint, dest *uint16) Operation {
	return func(c *CPU) {
		*dest = uint16(c.MMU.Read(addr))
	}
}

// Read most significant byte of 16-bits value from memory into 16-bit destination.
func opReadD16HighAt(c *CPU, addr uint, dest *uint16) Operation {
	return func(c *CPU) {
		*dest |= uint16(c.MMU.Read(addr)) << 8
	}
}

// Write 8-bit value (evaluated at calling time) to memory.
func opWriteD8(c *CPU, addr uint, value uint8) Operation {
	return func(c *CPU) {
		c.MMU.Write(addr, value)
	}
}

// Set double register value.
func opSetRr(c *CPU, register, value *uint16) Operation {
	return func(c *CPU) {
		*register = *value
	}
}

// Set double register value (evaluated at calling time).
func opSetRrDirect(c *CPU, register *uint16, value uint16) Operation {
	return func(c *CPU) {
		*register = value
	}
}

// Instructions. Each takes a CPU pointer and will modify its internal state.
// Source: http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html

// Number of cycles reflects the information given in resource linked above.
// Note that reading the instruction byte itself takes 4 cycles (8 for CB xx instructions.)
// Those 4 cycles are included in the count indicated in the helper's comment.
// Each subsequent Operation pushed on the CPU will take an additional 4 cycles.

// Helpers
// LD rr,d16		12 cycles
func ldRrD16(c *CPU, high, low *byte) (done bool) {
	c.ops.Push(opReadD8(c, low))
	c.ops.Push(opReadD8(c, high))
	return false
}

// AND r
func andR(c *CPU, register *byte) {
	c.A &= *register
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
}

// OR r
func orR(c *CPU, register *byte) {
	c.A |= *register
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
}

// XOR r
func xorR(c *CPU, register *byte) {
	c.A ^= *register
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
}

// JR condition,r8	8 cycles if condition is false, 12 if true
func jrXxR8(c *CPU, condition bool) {
	c.ops.Push(Operation(func(c *CPU) {
		offset := int8(c.NextByte())
		// If condition matches, take another tick to update PC
		if condition {
			c.ops.Push(Operation(func(c *CPU) {
				// Need cast to signed for the potential substraction
				c.PC = uint16(int16(c.PC) + int16(offset))
			}))
		}
	}))
}

// INC r			4 cycles
func incR(c *CPU, register *byte) {
	// Flags z 0 h -
	c.F &= ^FlagN
	if *register > 0x0F {
		c.F |= FlagH
	}
	*register++
	if *register == 0 {
		c.F |= FlagZ
	}
}

// DEC r			4 cycles
func decR(c *CPU, register *byte) {
	// Flags z 1 h -
	c.F &= FlagC
	c.F |= FlagN
	if *register > 0x0F {
		c.F |= FlagH
	}
	*register--
	if *register == 0 {
		c.F |= FlagZ
	}
}

// INC rr			8 cycles
func incRr(c *CPU, high, low *uint8) {
	c.ops.Push(Operation(func(c *CPU) {
		if *low == 0xff {
			*high++
		}
		*low++
	}))
}

// DEC rr			8 cycles
func decRr(c *CPU, high, low *uint8) {
	c.ops.Push(Operation(func(c *CPU) {
		if *low == 0x00 {
			*high--
		}
		*low--
	}))
}

// PUSH rr			16 cycles
func pushRr(c *CPU, high, low uint8) {
	c.ops.Push(Operation(func(c *CPU) {
		// Waiting cycle. I guess to align with memory timing?
	}))
	c.ops.Push(Operation(func(c *CPU) {
		c.SP--
		c.MMU.Write(uint(c.SP), high)
	}))
	c.ops.Push(Operation(func(c *CPU) {
		c.SP--
		c.MMU.Write(uint(c.SP), low)
	}))
}

// POP rr			12 cycles
func popRr(c *CPU, high, low *uint8) {
	c.ops.Push(Operation(func(c *CPU) {
		*low = c.MMU.Read(uint(c.SP))
		c.SP++
	}))
	c.ops.Push(Operation(func(c *CPU) {
		*high = c.MMU.Read(uint(c.SP))
		c.SP++
	}))
}

// POP PC			12 cycles
func popPc(c *CPU) {
	c.temp = c.SP + 2
	c.ops.Push(opReadD16LowAt(c, uint(c.SP), &c.PC))
	c.ops.Push(opReadD16HighAt(c, uint(c.SP+1), &c.PC))
	c.ops.Push(opSetRr(c, &c.SP, &c.temp))
}

// RL r -- rotate left through carry
func rlR(c *CPU, register *byte) {
	result := *register << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if *register&(1<<7) > 0 {
		c.F |= FlagC
	}
	*register = result
}

// BIT n,r			8 cycles
func bitNR(c *CPU, bit, register byte) {
	// Flags z 0 1 -
	if register&(1<<bit) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
}

// SWAP r			8 cycles
func swapR(c *CPU, register *byte) {
	// Flags z 0 0 0
	*register = (*register << 4) | (*register >> 4)
	if *register == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
}

// ADD A,r
func addAR(c *CPU, register byte) {
	// Flags: z 0 h c
	c.F = 0
	if c.A&0xf+register&0xf > 0xf {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(register)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
}

// SUB r -- SUB d8 -- CP r -- CP d8
// Only sets flags, return substraction result
func subFlags(c *CPU, value byte) byte {
	// Flags: z 1 h c
	c.F = FlagN
	if value&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if value > c.A {
		c.F |= FlagC
	}
	result := c.A - value
	if result == 0 {
		c.F |= FlagZ
	}
	return result
}

func subD8(c *CPU, value byte) {
	c.A = subFlags(c, value)
}

// Opcodes

// 00: NOP			4 cycles
func nop(c *CPU) (done bool) {
	return true
}

// 01: LD BC,d16	12 cycles
func ldBcD16(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.C))
	c.ops.Push(opReadD8(c, &c.B))
	return false
}

// 03 INC BC		8 cycles
func incBc(c *CPU) (done bool) {
	incRr(c, &c.B, &c.C) // FIXME: use same logic as other operations. Rename to opIncRr
	return false
}

// 04: INC B		4 cycles
func incB(c *CPU) (done bool) {
	incR(c, &c.B)
	return true
}

// 05: DEC B		4 cycles
func decB(c *CPU) (done bool) {
	decR(c, &c.B)
	return true
}

// 06: LD B,d8		8 cycles
func ldBD8(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.B))
	return false
}

// 0B: DEC BC		8 cycles
func decBc(c *CPU) (done bool) {
	decRr(c, &c.B, &c.C)
	return false
}

// 0C: INC C		4 cycles
func incC(c *CPU) (done bool) {
	incR(c, &c.C)
	return true
}

// 0D: DEC C		4 cycles
func decC(c *CPU) (done bool) {
	decR(c, &c.C)
	return true
}

// 0E: LD C,d8		8 cycles
func ldCD8(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.C))
	return false
}

// 11: LD DE,d16	12 cycles
func ldDeD16(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.E))
	c.ops.Push(opReadD8(c, &c.D))
	return false
}

// 13: INC DE		8 cycles
func incDe(c *CPU) (done bool) {
	incRr(c, &c.D, &c.E)
	return false
}

// 14: INC D		4 cycles
func incD(c *CPU) (done bool) {
	incR(c, &c.D)
	return true
}

// 15: DEC D		4 cycles
func decD(c *CPU) (done bool) {
	decR(c, &c.D)
	return true
}

// 16: LD D,d8		8 cycles
func ldDD8(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.D))
	return false
}

// 17: RLA -- RL A	4 cycles
func rlA(c *CPU) (done bool) {
	rlR(c, &c.A)
	return true
}

// 18: JR r8		12 cycles
func jrR8(c *CPU) (done bool) {
	jrXxR8(c, true)
	return false
}

// 1A: LD A,(DE)	8 cycles
func ldAAddrDe(c *CPU) (done bool) {
	c.ops.Push(opReadD8At(c, uint(c.DE()), &c.A))
	return false
}

// 1C: INC E		4 cycles
func incE(c *CPU) (done bool) {
	incR(c, &c.E)
	return true
}

// 1D: DEC E		4 cycles
func decE(c *CPU) (done bool) {
	decR(c, &c.E)
	return true
}

// 1E: LD E,d8		8 cycles
func ldED8(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.E))
	return false
}

// 20: JR NZ,r8		12/8 cycles
func jrNzR8(c *CPU) (done bool) {
	jrXxR8(c, c.F&FlagZ == 0)
	return false
}

// 21: LD HL,d16	12 cycles
func ldHlD16(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.L))
	c.ops.Push(opReadD8(c, &c.H))
	return false
}

// 22: LD (HL+),A	8 cycles
func ldiAddrHlA(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		hl := c.HL()
		c.MMU.Write(uint(hl), c.A)
		c.SetHL(hl + 1)
	}))
	return false
}

// 23: INC HL		8 cycles
func incHl(c *CPU) (done bool) {
	incRr(c, &c.H, &c.L)
	return false
}

// 24: INC H		4 cycles
func incH(c *CPU) (done bool) {
	incR(c, &c.H)
	return true
}

// 25: DEC H		4 cycles
func decH(c *CPU) (done bool) {
	decR(c, &c.H)
	return true
}

// 26: LD H,d8		8 cycles
func ldHD8(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.H))
	return false
}

// 28: JR Z,r8		12/8 cycles
func jrZR8(c *CPU) (done bool) {
	jrXxR8(c, c.F&FlagZ == FlagZ)
	return false
}

// 2A: LD A,(HL+)	8 cycles
func ldiAAddrHl(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		hl := c.HL()
		c.A = c.MMU.Read(uint(hl))
		c.SetHL(hl + 1)
	}))
	return false
}

// 2C: INC L		4 cycles
func incL(c *CPU) (done bool) {
	incR(c, &c.L)
	return true
}

// 2D: DEC L		4 cycles
func decL(c *CPU) (done bool) {
	decR(c, &c.L)
	return true
}

// 2E: LD L,d8		8 cycles
func ldLD8(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.L))
	return false
}

// 2F: CPL			4 cycles
func cpl(c *CPU) (done bool) {
	c.F |= FlagN | FlagH
	c.A ^= 0xff
	return true
}

// 31: LD SP,d16	12 cycles
func ldSpD16(c *CPU) (done bool) {
	c.ops.Push(opReadD16Low(c, &c.SP))
	c.ops.Push(opReadD16High(c, &c.SP))
	return false
}

// 32: LD (HL-),A	8 cycles
func lddHlA(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		hl := c.HL()
		c.MMU.Write(uint(hl), c.A)
		c.SetHL(hl - 1)
	}))
	return false
}

// 33: INC SP		8 cycles
func incSp(c *CPU) (done bool) {
	c.ops.Push(opSetRrDirect(c, &c.SP, c.SP+1))
	return false
}

// 34: INC (HL)		12 cycles
func incAddrHl(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.temp8 = c.MMU.Read(uint(c.HL()))
		incR(c, &c.temp8)
	}))
	c.ops.Push(Operation(func(c *CPU) {
		c.MMU.Write(uint(c.HL()), c.temp8)
	}))
	return false
}

// 35: DEC (HL)		12 cycles
func decAddrHl(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.temp8 = c.MMU.Read(uint(c.HL()))
		decR(c, &c.temp8)
	}))
	c.ops.Push(Operation(func(c *CPU) {
		c.MMU.Write(uint(c.HL()), c.temp8)
	}))
	return false
}

// 36: LD (HL),d8	12 cycles
func ldAddrHlD8(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.temp8))
	c.ops.Push(Operation(func(c *CPU) { // FIXME: used at least twice, do we need opWriteD8/opWriteLastD8 or something?
		c.MMU.Write(uint(c.HL()), c.temp8)
	}))
	return false
}

// 3C: INC A		4 cycles
func incA(c *CPU) (done bool) {
	incR(c, &c.A)
	return true
}

// 3D: DEC A		4 cycles
func decA(c *CPU) (done bool) {
	decR(c, &c.A)
	return true
}

// 3E: LD A,d8		8 cycles
func ldAD8(c *CPU) (done bool) {
	c.ops.Push(opReadD8(c, &c.A))
	return false
}

// 40: LD B,B		4 cycles
func ldBB(c *CPU) (done bool) {
	// nop
	return true
}

// 41: LD B,C		4 cycles
func ldBC(c *CPU) (done bool) {
	c.B = c.C
	return true
}

// 42: LD B,D		4 cycles
func ldBD(c *CPU) (done bool) {
	c.B = c.D
	return true
}

// 43: LD B,E		4 cycles
func ldBE(c *CPU) (done bool) {
	c.B = c.E
	return true
}

// 44: LD B,H		4 cycles
func ldBH(c *CPU) (done bool) {
	c.B = c.H
	return true
}

// 45: LD B,L		4 cycles
func ldBL(c *CPU) (done bool) {
	c.B = c.L
	return true
}

// 46: LD B,(HL)	8 cycles
func ldBAddrHl(c *CPU) (done bool) {
	c.ops.Push(opReadD8At(c, uint(c.HL()), &c.B))
	return false
}

// 47: LD B,A		4 cycles
func ldBA(c *CPU) (done bool) {
	c.B = c.A
	return true
}

// 48: LD C,B		4 cycles
func ldCB(c *CPU) (done bool) {
	c.C = c.B
	return true
}

// 49: LD C,C		4 cycles
func ldCC(c *CPU) (done bool) {
	// nop
	return true
}

// 4A: LD C,D		4 cycles
func ldCD(c *CPU) (done bool) {
	c.C = c.D
	return true
}

// 4B: LD C,E		4 cycles
func ldCE(c *CPU) (done bool) {
	c.C = c.E
	return true
}

// 4C: LD C,H		4 cycles
func ldCH(c *CPU) (done bool) {
	c.C = c.H
	return true
}

// 4D: LD C,L		4 cycles
func ldCL(c *CPU) (done bool) {
	c.C = c.L
	return true
}

// 4E: LD C,(HL)	8 cycles
func ldCAddrHl(c *CPU) (done bool) {
	c.ops.Push(opReadD8At(c, uint(c.HL()), &c.C))
	return false
}

// 4F: LD C,A		4 cycles
func ldCA(c *CPU) (done bool) {
	c.C = c.A
	return true
}

// 57: LD D,A		4 cycles
func ldDA(c *CPU) (done bool) {
	c.D = c.A
	return true
}

// 60: LD H,B		4 cycles
func ldHB(c *CPU) (done bool) {
	c.H = c.B
	return true
}

// 61: LD H,C		4 cycles
func ldHC(c *CPU) (done bool) {
	c.H = c.C
	return true
}

// 62: LD H,D		4 cycles
func ldHD(c *CPU) (done bool) {
	c.H = c.D
	return true
}

// 63: LD H,E		4 cycles
func ldHE(c *CPU) (done bool) {
	c.H = c.E
	return true
}

// 64: LD H,H		4 cycles
func ldHH(c *CPU) (done bool) {
	return true
}

// 65: LD H,L		4 cycles
func ldHL(c *CPU) (done bool) {
	c.H = c.L
	return true
}

// 66: LD H,(HL)	8 cycles
func ldHAddrHl(c *CPU) (done bool) {
	c.ops.Push(opReadD8At(c, uint(c.HL()), &c.H))
	return false
}

// 67: LD H,A		4 cycles
func ldHA(c *CPU) (done bool) {
	c.H = c.A
	return true
}

// 68: LD L,B		4 cycles
func ldLB(c *CPU) (done bool) {
	c.L = c.B
	return true
}

// 69: LD L,C		4 cycles
func ldLC(c *CPU) (done bool) {
	c.L = c.C
	return true
}

// 6A: LD L,D		4 cycles
func ldLD(c *CPU) (done bool) {
	c.L = c.D
	return true
}

// 6B: LD L,E		4 cycles
func ldLE(c *CPU) (done bool) {
	c.L = c.E
	return true
}

// 6C: LD L,H		4 cycles
func ldLH(c *CPU) (done bool) {
	c.L = c.H
	return true
}

// 6D: LD L,L		4 cycles
func ldLL(c *CPU) (done bool) {
	return true
}

// 6E: LD L,(HL)	8 cycles
func ldLAddrHl(c *CPU) (done bool) {
	c.ops.Push(opReadD8At(c, uint(c.HL()), &c.L))
	return false
}

// 6F: LD L,A		4 cycles
func ldLA(c *CPU) (done bool) {
	c.L = c.A
	return true
}

// 70: LD (HL),B	8 cycles
func ldAddrHlB(c *CPU) (done bool) {
	c.ops.Push(opWriteD8(c, uint(c.HL()), c.B))
	return false
}

// 71: LD (HL),C	8 cycles
func ldAddrHlC(c *CPU) (done bool) {
	c.ops.Push(opWriteD8(c, uint(c.HL()), c.C))
	return false
}

// 72: LD (HL),D	8 cycles
func ldAddrHlD(c *CPU) (done bool) {
	c.ops.Push(opWriteD8(c, uint(c.HL()), c.D))
	return false
}

// 73: LD (HL),E	8 cycles
func ldAddrHlE(c *CPU) (done bool) {
	c.ops.Push(opWriteD8(c, uint(c.HL()), c.E))
	return false
}

// 74: LD (HL),H	8 cycles
func ldAddrHlH(c *CPU) (done bool) {
	c.ops.Push(opWriteD8(c, uint(c.HL()), c.H))
	return false
}

// 75: LD (HL),L	8 cycles
func ldAddrHlL(c *CPU) (done bool) {
	c.ops.Push(opWriteD8(c, uint(c.HL()), c.L))
	return false
}

// 77: LD (HL),A	8 cycles
func ldAddrHlA(c *CPU) (done bool) {
	c.ops.Push(opWriteD8(c, uint(c.HL()), c.A))
	return false
}

// 78: LD A,B		4 cycles
func ldAB(c *CPU) (done bool) {
	c.A = c.B
	return true
}

// 79: LD A,C		4 cycles
func ldAC(c *CPU) (done bool) {
	c.A = c.C
	return true
}

// 7A: LD A,D		4 cycles
func ldAD(c *CPU) (done bool) {
	c.A = c.D
	return true
}

// 7B: LD A,E		4 cycles
func ldAE(c *CPU) (done bool) {
	c.A = c.E
	return true
}

// 7C: LD A,H		4 cycles
func ldAH(c *CPU) (done bool) {
	c.A = c.H
	return true
}

// 7D: LD A,L		4 cycles
func ldAL(c *CPU) (done bool) {
	c.A = c.L
	return true
}

// 7E: LD A,(HL)	8 cycles
func ldAAddrHl(c *CPU) (done bool) {
	c.ops.Push(opReadD8At(c, uint(c.HL()), &c.A))
	return false
}

// 7F: LD A,A		4 cycles
func ldAA(c *CPU) (done bool) {
	// nop
	return true
}

// 80: ADD A,B		4 cycles
func addAB(c *CPU) (done bool) {
	addAR(c, c.B)
	return true
}

// 81: ADD A,C		4 cycles
func addAC(c *CPU) (done bool) {
	addAR(c, c.C)
	return true
}

// 82: ADD A,D		4 cycles
func addAD(c *CPU) (done bool) {
	addAR(c, c.D)
	return true
}

// 83: ADD A,E		4 cycles
func addAE(c *CPU) (done bool) {
	addAR(c, c.E)
	return true
}

// 84: ADD A,H		4 cycles
func addAH(c *CPU) (done bool) {
	addAR(c, c.H)
	return true
}

// 85: ADD A,L		4 cycles
func addAL(c *CPU) (done bool) {
	addAR(c, c.L)
	return true
}

// 86: ADD A,(HL)	8 cycles
func addAAddrHl(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		addAR(c, c.MMU.Read(uint(c.HL())))
	}))
	return false
}

// 87: ADD A,A		4 cycles
func addAA(c *CPU) (done bool) {
	addAR(c, c.A)
	return true
}

// 90: SUB B		4 cycles
func subB(c *CPU) (done bool) {
	subD8(c, c.B)
	return true
}

// 91: SUB C		4 cycles
func subC(c *CPU) (done bool) {
	subD8(c, c.C)
	return true
}

// 92: SUB D		4 cycles
func subD(c *CPU) (done bool) {
	subD8(c, c.D)
	return true
}

// 93: SUB E		4 cycles
func subE(c *CPU) (done bool) {
	subD8(c, c.E)
	return true
}

// 94: SUB H		4 cycles
func subH(c *CPU) (done bool) {
	subD8(c, c.H)
	return true
}

// 95: SUB L		4 cycles
func subL(c *CPU) (done bool) {
	subD8(c, c.L)
	return true
}

/*
// 96: SUB (HL)
func subAddrHl(c *CPU) (done bool) {
	subD8(c, c.Read(uint(c.HL())))
}

// 97: SUB A
func subA(c *CPU) (done bool) {
	subD8(c, c.A)
}
*/

// A0: AND B		4 cycles
func andB(c *CPU) (done bool) {
	andR(c, &c.B)
	return true
}

// A1: AND C		4 cycles
func andC(c *CPU) (done bool) {
	andR(c, &c.C)
	return true
}

// A2: AND D		4 cycles
func andD(c *CPU) (done bool) {
	andR(c, &c.D)
	return true
}

// A3: AND E		4 cycles
func andE(c *CPU) (done bool) {
	andR(c, &c.E)
	return true
}

// A4: AND H		4 cycles
func andH(c *CPU) (done bool) {
	andR(c, &c.H)
	return true
}

// A5: AND L		4 cycles
func andL(c *CPU) (done bool) {
	andR(c, &c.L)
	return true
}

// A6: AND (HL)		8 cycles
func andAddrHl(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.temp8 = c.MMU.Read(uint(c.HL()))
		andR(c, &c.temp8)
	}))
	return false
}

// A7: AND A		4 cycles
func andA(c *CPU) (done bool) {
	andR(c, &c.A)
	return true
}

// A8: XOR B		4 cycles
func xorB(c *CPU) (done bool) {
	xorR(c, &c.B)
	return true
}

// A9: XOR C		4 cycles
func xorC(c *CPU) (done bool) {
	xorR(c, &c.C)
	return true
}

// AA: XOR D		4 cycles
func xorD(c *CPU) (done bool) {
	xorR(c, &c.D)
	return true
}

// AB: XOR E		4 cycles
func xorE(c *CPU) (done bool) {
	xorR(c, &c.E)
	return true
}

// AC: XOR H		4 cycles
func xorH(c *CPU) (done bool) {
	xorR(c, &c.H)
	return true
}

// AD: XOR L		4 cycles
func xorL(c *CPU) (done bool) {
	xorR(c, &c.L)
	return true
}

// AE: XOR (HL)		8 cycles
func xorAddrHl(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.temp8 = c.MMU.Read(uint(c.HL()))
		xorR(c, &c.temp8)
	}))
	return false
}

// AF: XOR A		4 cycles
func xorA(c *CPU) (done bool) {
	xorR(c, &c.A)
	return true
}

// B0: OR B			4 cycles
func orB(c *CPU) (done bool) {
	orR(c, &c.B)
	return true
}

// B1: OR C			4 cycles
func orC(c *CPU) (done bool) {
	orR(c, &c.C)
	return true
}

// B2: OR D			4 cycles
func orD(c *CPU) (done bool) {
	orR(c, &c.D)
	return true
}

// B3: OR E			4 cycles
func orE(c *CPU) (done bool) {
	orR(c, &c.E)
	return true
}

// B4: OR H			4 cycles
func orH(c *CPU) (done bool) {
	orR(c, &c.H)
	return true
}

// B5: OR L			4 cycles
func orL(c *CPU) (done bool) {
	orR(c, &c.L)
	return true
}

// B6: OR (HL)		8 cycles
func orAddrHl(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.temp8 = c.MMU.Read(uint(c.HL()))
		orR(c, &c.temp8)
	}))
	return false
}

// B7: OR A			4 cycles
func orA(c *CPU) (done bool) {
	orR(c, &c.A)
	return true
}

// BE: CP (HL)		8 cycles
func cpHl(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		subFlags(c, c.MMU.Read(uint(c.HL())))
	}))
	return false
}

// C1: POP BC		12 cycles
func popBc(c *CPU) (done bool) {
	popRr(c, &c.B, &c.C)
	return false
}

// C3: JP a16		16 cycles
func jpA16(c *CPU) (done bool) {
	c.ops.Push(opReadD16Low(c, &c.temp))
	c.ops.Push(opReadD16High(c, &c.temp))
	c.ops.Push(opSetRr(c, &c.PC, &c.temp))
	return false
}

// C5: PUSH BC		16 cycles
func pushBc(c *CPU) (done bool) {
	pushRr(c, c.B, c.C)
	return false
}

// C9: RET			16 cycles
func ret(c *CPU) (done bool) {
	// Simulate POP PC for consistency
	popPc(c)
	return false
}

// CB 11: RL C		8 cycles
func rlC(c *CPU) (done bool) {
	rlR(c, &c.C)
	return true
}

// CB 37: SWAP A	8 cycles
func swapA(c *CPU) (done bool) {
	swapR(c, &c.A)
	return true
}

// CB 7C: BIT 7,H	8 cycles
func bit7H(c *CPU) (done bool) {
	bitNR(c, 7, c.H)
	return true
}

// CD: CALL a16		24 cycles
func callA16(c *CPU) (done bool) {
	// Advance PC before pushing to stack.
	c.ops.Push(opReadD16Low(c, &c.temp))  // 4 cycles
	c.ops.Push(opReadD16High(c, &c.temp)) // 4 cycles

	// Do just like pushRr but only read PC after previous operations, and tack instantaneous PC update at the end.
	newSP := c.SP - 2                     // This feels meh. Either I implement an opSetRrImmediate or move away from operation FIFO altogether.
	c.ops.Push(opSetRr(c, &c.SP, &newSP)) // 4 cycles
	c.ops.Push(Operation(func(c *CPU) {   // 4 cycles
		c.MMU.Write(uint(c.SP+1), uint8(c.PC>>8))
	}))
	c.ops.Push(Operation(func(c *CPU) { // 4 cycles
		c.MMU.Write(uint(c.SP), uint8(c.PC&0xff))
		c.PC = c.temp
	}))
	return false
}

/*
// D1: POP DE
func popDe(c *CPU) (done bool) {
	popRr(c, &c.D, &c.E)
}

// D5: PUSH DE
func pushDe(c *CPU) (done bool) {
	pushRr(c, c.D, c.E)
}
*/
// E0: LD (FF00+a8),A	12 cycles
func ldAddrFfA8A(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.ops.Push(opWriteD8(c, uint(0xff00+uint16(c.NextByte())), c.A))
	}))
	return false
}

/*
// E1: POP HL
func popHl(c *CPU) (done bool) {
	popRr(c, &c.H, &c.L)
}
*/
// E2: LD (FF00+C),A	8 cycles
func ldAddrFfCA(c *CPU) (done bool) {
	c.ops.Push(opWriteD8(c, uint(0xff00+uint16(c.C)), c.A))
	return false
}

/*
// E5: PUSH HL
func pushHl(c *CPU) (done bool) {
	pushRr(c, c.H, c.L)
}
*/

// E6: AND d8		8 cycles
func andD8(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.temp8 = c.NextByte()
		andR(c, &c.temp8)
	}))
	return false
}

// EA: LD (a16),A	16 cycles
func ldAddrA16A(c *CPU) (done bool) {
	c.ops.Push(opReadD16Low(c, &c.temp))
	c.ops.Push(opReadD16High(c, &c.temp))
	c.ops.Push(Operation(func(c *CPU) {
		c.MMU.Write(uint(c.temp), c.A)
	}))
	return false
}

// EF: RST 28h		16 cycles
func rst28h(c *CPU) (done bool) {
	c.temp = c.SP - 2
	c.ops.Push(opSetRr(c, &c.SP, &c.temp)) // 4 cycles
	c.ops.Push(Operation(func(c *CPU) {    // 4 cycles
		c.MMU.Write(uint(c.SP+1), uint8(c.PC>>8))
	}))
	c.ops.Push(Operation(func(c *CPU) { // 4 cycles
		c.MMU.Write(uint(c.SP), uint8(c.PC&0xff))
		c.PC = 0x0028
	}))
	return false
}

// F0: LD A,(FF00+a8)	12 cycles
func ldAAddrFfA8(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.temp = uint16(c.NextByte()) | 0xff00
	}))
	c.ops.Push(Operation(func(c *CPU) {
		c.A = c.MMU.Read(uint(c.temp))
	}))
	return false
}

// F1: POP AF		12 cycles
func popAf(c *CPU) (done bool) {
	popRr(c, &c.A, &c.F)
	return false
}

// F3: DI			4 cycles
func di(c *CPU) (done bool) {
	fmt.Println("DI: Interruptions not implemented yet")
	return true
}

// F5: PUSH AF		16 cycles
func pushAf(c *CPU) (done bool) {
	pushRr(c, c.A, c.F)
	return false
}

// FA: LD A,(a16)	16 cycles
func ldAAddrA16(c *CPU) (done bool) {
	c.ops.Push(opReadD16Low(c, &c.temp))
	c.ops.Push(opReadD16High(c, &c.temp))
	c.ops.Push(opReadD8AtAddr(c, &c.temp, &c.A))
	return false
}

// F3: EI			4 cycles
func ei(c *CPU) (done bool) {
	fmt.Println("EI: Interruptions not implemented yet")
	return true
}

// FE: CP d8		8 cycles
func cpD8(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		subFlags(c, c.NextByte())
	}))
	return false
}
