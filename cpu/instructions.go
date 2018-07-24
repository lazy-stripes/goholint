package cpu

// An Operation executed on the CPU as part of an instruction.
type Operation func(c *CPU)

// An Instruction to be executed by a CPU, pushing a number of 4-cycle micro-operations to the CPU.
type Instruction func(c *CPU) (done bool)

// LR35902InstructionSet is an array of opcodes for the DMG CPU.
var LR35902InstructionSet = []Instruction{
	0x00: nop,
	0x01: ldBcD16,
	/*
		0x03: incBc,
		0x04: incB,
	*/
	0x05: decB,
	0x06: ldBD8,
	0x0c: incC,
	0x0d: decC,
	0x0e: ldCD8,
	0x11: ldDeD16,
	0x13: incDe,
	/*
		0x14: incD,
		0x15: decD,
		0x16: ldDD8,
	*/
	0x17: rlA,
	/*
		0x18: jrR8,
	*/
	0x1a: ldAAddrDe,
	/*
		0x1c: incE,
		0x1d: decE,
		0x1e: ldED8,
	*/
	0x20: jrNzR8,
	0x21: ldHlD16,
	0x22: ldiHlA,
	0x23: incHl,
	/*
		0x24: incH,
		0x25: decH,
		0x26: ldHD8,
		0x28: jrZR8,
		0x2c: incL,
		0x2d: decL,
		0x2e: ldLD8,
	*/
	0x31: ldSpD16,
	0x32: lddHlA,
	/*
		0x33: incSp,
		0x34: incAddrHl,
		0x35: decAddrHl,
		0x3c: incA,
		0x3d: decA,
	*/
	0x3e: ldAD8,
	/*
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
	*/
	0x4f: ldCA,
	0x57: ldDA,
	0x60: ldHB,
	0x61: ldHC,
	0x62: ldHD,
	0x63: ldHE,
	0x64: ldHH,
	0x65: ldHL,
	/*
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
	/*
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
	0xaf: xorA,
	0xc1: popBc,
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
		0xea: ldAddrA16A,
		0xf0: ldAAddrFfA8,
		0xf1: popAf,
		0xf5: pushAf,
		0xfa: ldAAddrA16,
	*/
	0xfe: cpD8,
}

// LR35902ExtendedInstructionSet is the array of extension opcodes for the DMG CPU.
var LR35902ExtendedInstructionSet = []Instruction{
	0x11: rlC,
	0x7c: bit7H,
}

// Operations. The pseudo-atomic things the CPU does as part as an Instruction, which might take many cycles.

// Read next 8-bit argument into destination address.
func readD8(c *CPU, dest *byte) Operation {
	return func(c *CPU) {
		*dest = c.NextByte()
	}
}

// Read 8-bit value from memory into destination address.
func readD8At(c *CPU, addr uint, dest *byte) Operation {
	return func(c *CPU) {
		*dest = c.MMU.Read(addr)
	}
}

// Read least significant byte of 16-bits argument into 16-bit destination.
func readD16Low(c *CPU, dest *uint16) Operation {
	return func(c *CPU) {
		*dest = uint16(c.NextByte())
	}
}

// Read most significant byte of 16-bits argument into 16-bit destination.
func readD16High(c *CPU, dest *uint16) Operation {
	return func(c *CPU) {
		*dest |= uint16(c.NextByte()) << 8
	}
}

// Read least significant byte of 16-bits value from memory into 16-bit destination.
func readD16LowAt(c *CPU, addr uint, dest *uint16) Operation {
	return func(c *CPU) {
		*dest = uint16(c.MMU.Read(addr))
	}
}

// Read most significant byte of 16-bits value from memory into 16-bit destination.
func readD16HighAt(c *CPU, addr uint, dest *uint16) Operation {
	return func(c *CPU) {
		*dest |= uint16(c.MMU.Read(addr)) << 8
	}
}

// Write 8-bit value to memory.
func writeD8(c *CPU, addr uint, value uint8) Operation {
	return func(c *CPU) {
		c.MMU.Write(addr, value)
	}
}

// Set double register value.
func setRr(c *CPU, register *uint16, value uint16) Operation {
	return func(c *CPU) {
		*register = value
	}
}

// Instructions. Each takes a CPU pointer and will modify its internal state.
// Source: http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html

// Helpers
// LD rr,d16
func ldRrD16(c *CPU, high, low *byte) bool {
	c.ops.Push(readD8(c, low))
	c.ops.Push(readD8(c, high))
	return false
}

/*
// LD r,(HL)
func ldRAddrHl(c *CPU, register *byte) {
	*register = c.Read(uint(c.HL()))
}
*/

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

// PUSH rr			16 cycles
func pushRr(c *CPU, high, low uint8) {
	c.ops.Push(setRr(c, &c.SP, c.SP-2))
	c.ops.Push(writeD8(c, uint(c.SP-1), high)) // SP hasn't been decremented yet
	c.ops.Push(writeD8(c, uint(c.SP-2), low))  // SP hasn't been decremented yet
}

// POP rr			12 cycles
func popRr(c *CPU, high, low *uint8) {
	c.ops.Push(readD8At(c, uint(c.SP), low))
	c.ops.Push(readD8At(c, uint(c.SP+1), high))
	c.ops.Push(setRr(c, &c.SP, c.SP+2))
}

// POP PC			12 cycles
func popPc(c *CPU) {
	c.ops.Push(readD16LowAt(c, uint(c.SP), &c.PC))
	c.ops.Push(readD16HighAt(c, uint(c.SP+1), &c.PC))
	c.ops.Push(setRr(c, &c.SP, c.SP+2))
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

/*
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
*/

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
	c.ops.Push(readD8(c, &c.C))
	c.ops.Push(readD8(c, &c.B))
	return false
}

/*
// 03 INC BC
func incBc(c *CPU) {
	incRr(c, &c.B, &c.C)
}

// 04: INC B
func incB(c *CPU) {
	incR(c, &c.B)
}
*/

// 05: DEC B		4 cycles.
func decB(c *CPU) (done bool) {
	decR(c, &c.B)
	return true
}

// 06: LD B,d8		8 cycles
func ldBD8(c *CPU) (done bool) {
	c.ops.Push(readD8(c, &c.B))
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
	c.ops.Push(readD8(c, &c.C))
	return false
}

// 11: LD DE,d16
func ldDeD16(c *CPU) (done bool) {
	c.ops.Push(readD8(c, &c.E))
	c.ops.Push(readD8(c, &c.D))
	return false
}

// 13: INC DE		8 cycles
func incDe(c *CPU) (done bool) {
	incRr(c, &c.D, &c.E)
	return false
}

/*
// 14: INC D
func incD(c *CPU) {
	incR(c, &c.D)
}

// 15: DEC D
func decD(c *CPU) {
	decR(c, &c.D)
}

// 16: LD D,d8
func ldDD8(c *CPU) {
	c.D = c.NextByte()
}
*/
// 17: RLA -- RL A	4 cycles
func rlA(c *CPU) (done bool) {
	rlR(c, &c.A)
	return true
}

/*
// 18: JR r8
func jrR8(c *CPU) {
	jrXxR8(c, true)
}
*/
// 1A: LD A,(DE)	8 cycles
func ldAAddrDe(c *CPU) (done bool) {
	c.ops.Push(readD8At(c, uint(c.DE()), &c.A))
	return false
}

/*
// 1C: INC E
func incE(c *CPU) {
	incR(c, &c.E)
}

// 1D: DEC E
func decE(c *CPU) {
	decR(c, &c.E)
}
*/

// 1E: LD E,d8
func ldED8(c *CPU) (done bool) {
	c.ops.Push(readD8(c, &c.E))
	return false
}

// 20: JR NZ,r8
func jrNzR8(c *CPU) (done bool) {
	jrXxR8(c, c.F&FlagZ == 0)
	return false
}

// 21: LD HL,d16
func ldHlD16(c *CPU) (done bool) {
	c.ops.Push(readD8(c, &c.L))
	c.ops.Push(readD8(c, &c.H))
	return false
}

// 22: LD (HL+),A	8 cycles
func ldiHlA(c *CPU) (done bool) {
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

/*
// 24: INC H
func incH(c *CPU) {
	incR(c, &c.H)
}

// 25: DEC H
func decH(c *CPU) {
	decR(c, &c.H)
}

// 26: LD H,d8
func ldHD8(c *CPU) {
	c.H = c.NextByte()
}

// 28: JR Z,r8
func jrZR8(c *CPU) {
	jrXxR8(c, c.F&FlagZ == FlagZ)
}

// 2C: INC L
func incL(c *CPU) {
	incR(c, &c.L)
}

// 2D: DEC L
func decL(c *CPU) {
	decR(c, &c.L)
}
*/

// 2E: LD L,d8		8 cycles
func ldLD8(c *CPU) (done bool) {
	c.ops.Push(readD8(c, &c.L))
	return false
}

// 31: LD SP,d16	12 cycles
func ldSpD16(c *CPU) (done bool) {
	c.ops.Push(readD16Low(c, &c.SP))
	c.ops.Push(readD16High(c, &c.SP))
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

/*
// 33: INC SP
func incSp(c *CPU) {
	c.SP++
}

// 34: INC (HL)
func incAddrHl(c *CPU) {
	value := c.Read(uint(c.HL()))
	incR(c, &value)
	c.Write(uint(c.HL()), value)
}

// 35: DEC (HL)
func decAddrHl(c *CPU) {
	value := c.Read(uint(c.HL()))
	decR(c, &value)
	c.Write(uint(c.HL()), value)
}

// 3C: INC A
func incA(c *CPU) {
	incR(c, &c.A)
}

// 2D: DEC A
func decA(c *CPU) {
	decR(c, &c.A)
}
*/
// 3E: LD A,d8		8 cycles
func ldAD8(c *CPU) (done bool) {
	c.ops.Push(readD8(c, &c.A))
	return false
}

/*
// 40: LD B,B
func ldBB(c *CPU) {
	// nop
}

// 41: LD B,C
func ldBC(c *CPU) {
	c.B = c.C
}

// 42: LD B,D
func ldBD(c *CPU) {
	c.B = c.D
}

// 43: LD B,E
func ldBE(c *CPU) {
	c.B = c.E
}

// 44: LD B,H
func ldBH(c *CPU) {
	c.B = c.H
}

// 45: LD B,L
func ldBL(c *CPU) {
	c.B = c.L
}

// 46: LD B,(HL)
func ldBAddrHl(c *CPU) {
	ldRAddrHl(c, &c.B)
}

// 47: LD B,A
func ldBA(c *CPU) {
	c.B = c.A
}

// 48: LD C,B
func ldCB(c *CPU) {
	c.C = c.B
}

// 49: LD C,C
func ldCC(c *CPU) {
	// nop
}

// 4A: LD C,D
func ldCD(c *CPU) {
	c.C = c.D
}

// 4B: LD C,E
func ldCE(c *CPU) {
	c.C = c.E
}

// 4C: LD C,H
func ldCH(c *CPU) {
	c.C = c.H
}

// 4D: LD C,L
func ldCL(c *CPU) {
	c.C = c.L
}

// 4E: LD C,(HL)
func ldCAddrHl(c *CPU) {
	ldRAddrHl(c, &c.C)
}
*/
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

/*
// 66: LD H,(HL)
func ldHAddrHl(c *CPU) {
	c.H = c.Read(uint(c.HL()))
}

// 67: LD H,A
func ldHA(c *CPU) {
	c.H = c.A
}

// 68: LD L,B
func ldLB(c *CPU) {
	c.L = c.B
}

// 69: LD L,C
func ldLC(c *CPU) {
	c.L = c.C
}

// 6A: LD L,D
func ldLD(c *CPU) {
	c.L = c.D
}

// 6B: LD L,E
func ldLE(c *CPU) {
	c.L = c.E
}

// 6C: LD L,H
func ldLH(c *CPU) {
	c.L = c.H
}

// 6D: LD L,L
func ldLL(c *CPU) {
	c.L = c.L
}

// 6E: LD L,(HL)
func ldLAddrHl(c *CPU) {
	c.L = c.Read(uint(c.HL()))
}

// 6F: LD L,A
func ldLA(c *CPU) {
	c.L = c.A
}

// 70: LD (HL),B
func ldAddrHlB(c *CPU) {
	ldAddrHlR(c, c.A)
}

// 71: LD (HL),C
func ldAddrHlC(c *CPU) {
	ldAddrHlR(c, c.C)
}

// 72: LD (HL),D
func ldAddrHlD(c *CPU) {
	ldAddrHlR(c, c.D)
}

// 73: LD (HL),E
func ldAddrHlE(c *CPU) {
	ldAddrHlR(c, c.E)
}

// 74: LD (HL),H
func ldAddrHlH(c *CPU) {
	ldAddrHlR(c, c.H)
}

// 75: LD (HL),L
func ldAddrHlL(c *CPU) {
	ldAddrHlR(c, c.L)
}
*/
// 77: LD (HL),A	8 cycles
func ldAddrHlA(c *CPU) (done bool) {
	c.ops.Push(writeD8(c, uint(c.HL()), c.A))
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

/*
// 7E: LD A,(HL)
func ldAAddrHl(c *CPU) {
	ldRAddrHl(c, &c.A)
}

// 7F: LD A,A
func ldAA(c *CPU) {
	// nop
}

// 80: ADD A,B
func addAB(c *CPU) {
	addAR(c, c.B)
}

// 81: ADD A,C
func addAC(c *CPU) {
	addAR(c, c.C)
}

// 82: ADD A,D
func addAD(c *CPU) {
	addAR(c, c.D)
}

// 83: ADD A,E
func addAE(c *CPU) {
	addAR(c, c.E)
}

// 84: ADD A,H
func addAH(c *CPU) {
	addAR(c, c.H)
}

// 85: ADD A,L
func addAL(c *CPU) {
	addAR(c, c.L)
}

// 86: ADD A,(HL)
func addAAddrHl(c *CPU) {
	addAR(c, c.Read(uint(c.HL())))
}

// 87: ADD A,A
func addAA(c *CPU) {
	addAR(c, c.A)
}

// 90: SUB B
func subB(c *CPU) {
	subD8(c, c.B)
}

// 91: SUB C
func subC(c *CPU) {
	subD8(c, c.C)
}

// 92: SUB D
func subD(c *CPU) {
	subD8(c, c.D)
}

// 93: SUB E
func subE(c *CPU) {
	subD8(c, c.E)
}

// 94: SUB H
func subH(c *CPU) {
	subD8(c, c.H)
}

// 95: SUB L
func subL(c *CPU) {
	subD8(c, c.L)
}

// 96: SUB (HL)
func subAddrHl(c *CPU) {
	subD8(c, c.Read(uint(c.HL())))
}

// 97: SUB A
func subA(c *CPU) {
	subD8(c, c.A)
}
*/
// AF: XOR A		4 cycles
func xorA(c *CPU) (done bool) {
	xorR(c, &c.A)
	return true
}

// C1: POP BC		12 cycles
func popBc(c *CPU) (done bool) {
	popRr(c, &c.B, &c.C)
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

// CB 7C: BIT 7,H	8 cycles
func bit7H(c *CPU) (done bool) {
	bitNR(c, 7, c.H)
	return true
}

// CD: CALL a16		24 cycles
func callA16(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		// Advance PC before pushing
		//var addr uint16
		// FIXME: wrong number of cycles, reading addr should take 2 ticks, let's add a private addr uint16 to CPU!
		addr := c.NextWord()
		pushRr(c, uint8(c.PC>>8), uint8(c.PC&0xff)) // 12 cycles
		c.PC = addr
	}))
	return false
}

/*
// D1: POP DE
func popDe(c *CPU) {
	popRr(c, &c.D, &c.E)
}

// D5: PUSH DE
func pushDe(c *CPU) {
	pushRr(c, c.D, c.E)
}
*/
// E0: LD (FF00+a8),A	12 cycles
func ldAddrFfA8A(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		c.ops.Push(writeD8(c, uint(0xff00+uint16(c.NextByte())), c.A))
	}))
	return false
}

/*
// E1: POP HL
func popHl(c *CPU) {
	popRr(c, &c.H, &c.L)
}
*/
// E2: LD (FF00+C),A	8 cycles
func ldAddrFfCA(c *CPU) (done bool) {
	c.ops.Push(writeD8(c, uint(0xff00+uint16(c.C)), c.A))
	return false
}

/*
// E5: PUSH HL
func pushHl(c *CPU) {
	pushRr(c, c.H, c.L)
}
*/
// EA: LD (a16),A	16 cycles
func ldAddrA16A(c *CPU) {
	c.Write(uint(c.NextWord()), c.A)
}

/*
// F0: LD A,(FF00+a8)
func ldAAddrFfA8(c *CPU) {
	c.A = c.Read(uint(0xff00 + uint16(c.NextByte())))
}

// F1: POP AF
func popAf(c *CPU) {
	popRr(c, &c.A, &c.F)
}

// F5: PUSH AF
func pushAf(c *CPU) {
	pushRr(c, c.A, c.F)
}

// FA: LD A,(a16)
func ldAAddrA16(c *CPU) {
	c.A = c.Read(uint(c.NextWord()))
}
*/

// FE: CP d8		8 cycles
func cpD8(c *CPU) (done bool) {
	c.ops.Push(Operation(func(c *CPU) {
		subFlags(c, c.NextByte())
	}))
	return false
}
