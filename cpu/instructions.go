package cpu

// An Instruction to be executed by a CPU.
type Instruction func(c *CPU)

// LR35902InstructionSet is an array of opcodes for the DMG CPU.
var LR35902InstructionSet = [0x100]Instruction{
	0x00: nop,
	0x01: ldBcD16,
	0x03: incBc,
	0x05: decB,
	0x06: ldBD8,
	0x0c: incC,
	0x0e: ldCD8,
	0x11: ldDeD16,
	0x13: incDe,
	0x15: decD,
	0x16: ldDD8,
	0x17: rlA,
	0x1a: ldAAddrDe,
	0x1c: incE,
	0x1e: ldED8,
	0x20: jrNzR8,
	0x21: ldHlD16,
	0x22: ldiHlA,
	0x23: incHl,
	0x25: decH,
	0x26: ldHD8,
	0x2c: incL,
	0x2e: ldLD8,
	0x31: ldSpD16,
	0x32: lddHlA,
	0x33: incSp,
	0x35: decAddrHl,
	0x3c: incA,
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
	0x70: ldAddrHlB,
	0x71: ldAddrHlC,
	0x72: ldAddrHlD,
	0x73: ldAddrHlE,
	0x74: ldAddrHlH,
	0x75: ldAddrHlL,
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
	0xaf: xorA,
	0xc1: popBc,
	0xc5: pushBc,
	0xc9: ret,
	0xcd: callA16,
	0xd1: popDe,
	0xd5: pushDe,
	0xe0: ldAddrFfA8A,
	0xe1: popHl,
	0xe2: ldAddrFfCA,
	0xe5: pushHl,
	0xf1: popAf,
	0xf5: pushAf,
	0xfe: cpD8,
}

// LR35902ExtendedInstructionSet is the array of extension opcodes for the DMG CPU.
var LR35902ExtendedInstructionSet = [0x100]Instruction{
	0x11: rlC,
	0x7c: bit7H,
}

// Instructions. Each takes a CPU pointer and will modify its internal state.
// Opcode is hexadecimal. Source: http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html

// Helpers
// LD rr,d16
func ldRrD16(c *CPU, high, low *byte) {
	*low = c.Read()
	*high = c.Read()
}

// LD r,(HL)
func ldRAddrHl(c *CPU, register *byte) {
	*register = c.MMU.Read(uint(c.HL()))
}

// LD (HL),r
func ldAddrHlR(c *CPU, register byte) {
	c.MMU.Write(uint(c.HL()), register)
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

// JR condition,r8
func jrXxR8(c *CPU, condition bool) {
	offset := int8(c.Read())
	if condition {
		// Need cast to unsigned for the potential substraction
		c.PC = uint16(int16(c.PC) + int16(offset))
	}
}

// INC r
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

// DEC r
func decR(c *CPU, register *byte) {
	// Flags z 1 h -
	c.F |= FlagN
	if *register > 0x0F {
		c.F |= FlagH
	}
	*register--
	if *register == 0 {
		c.F |= FlagZ
	}
}

// INC rr
func incRr(c *CPU, high, low *uint8) {
	if *low == 0xff {
		*high++
	}
	*low++
}

// PUSH rr
func pushRr(c *CPU, high, low uint8) {
	c.SP -= 2
	c.MMU.Write(uint(c.SP), low)
	c.MMU.Write(uint(c.SP+1), high)
}

// POP rr
func popRr(c *CPU, high, low *uint8) {
	*low = c.MMU.Read(uint(c.SP))
	*high = c.MMU.Read(uint(c.SP + 1))
	c.SP += 2
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

// BIT n,r
func bitNR(c *CPU, bit, register byte) {
	// Flags z 0 1 -
	if register&(1<<bit) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
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

// SUB r -- sub d8 -- CP r -- CP d8
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

// Opcodes

// 00: NOP
func nop(c *CPU) {}

// 01: LD BC,d16
func ldBcD16(c *CPU) {
	ldRrD16(c, &c.B, &c.C)
}

// 03 INC BC
func incBc(c *CPU) {
	incRr(c, &c.B, &c.C)
}

// 05: DEC B
func decB(c *CPU) {
	decR(c, &c.B)
}

// 06: LD B,d8
func ldBD8(c *CPU) {
	c.B = c.Read()
}

// 0C: INC C
func incC(c *CPU) {
	incR(c, &c.C)
}

// 0E: LD C,d8
func ldCD8(c *CPU) {
	c.C = c.Read()
}

// 11: LD DE,d16
func ldDeD16(c *CPU) {
	ldRrD16(c, &c.D, &c.E)
}

// 13: INC DE
func incDe(c *CPU) {
	incRr(c, &c.D, &c.E)
}

// 15: DEC D
func decD(c *CPU) {
	decR(c, &c.D)
}

// 16: LD D,d8
func ldDD8(c *CPU) {
	c.D = c.Read()
}

// 17: RLA -- RL A
func rlA(c *CPU) {
	rlR(c, &c.A)
}

// 1A: LD A,(DE)
func ldAAddrDe(c *CPU) {
	c.A = c.MMU.Read(uint(c.DE()))
}

// 1C: INC E
func incE(c *CPU) {
	incR(c, &c.E)
}

// 1E: LD E,d8
func ldED8(c *CPU) {
	c.E = c.Read()
}

// 20: JR NZ,r8
func jrNzR8(c *CPU) {
	jrXxR8(c, c.F&FlagZ == 0)
}

// 21: LD HL,d16
func ldHlD16(c *CPU) {
	ldRrD16(c, &c.H, &c.L)
}

// 22: LD (HL+),A
func ldiHlA(c *CPU) {
	hl := c.HL()
	c.MMU.Write(uint(hl), c.A)
	c.SetHL(hl + 1)
}

// 23: INC HL
func incHl(c *CPU) {
	incRr(c, &c.H, &c.L)
}

// 25: DEC H
func decH(c *CPU) {
	decR(c, &c.H)
}

// 26: LD H,d8
func ldHD8(c *CPU) {
	c.H = c.Read()
}

// 2C: INC L
func incL(c *CPU) {
	incR(c, &c.L)
}

// 2E: LD L,d8
func ldLD8(c *CPU) {
	c.L = c.Read()
}

// 31: LD SP,d16
func ldSpD16(c *CPU) {
	c.SP = c.ReadWord()
}

// 32: LD (HL-),A
func lddHlA(c *CPU) {
	hl := c.HL()
	c.MMU.Write(uint(hl), c.A)
	c.SetHL(hl - 1)
}

// 33: INC SP
func incSp(c *CPU) {
	c.SP++
}

// 35: DEC (HL)
func decAddrHl(c *CPU) {
	value := c.MMU.Read(uint(c.HL()))
	decR(c, &value)
	c.MMU.Write(uint(c.HL()), value)
}

// 3C: INC A
func incA(c *CPU) {
	incR(c, &c.A)
}

// 3E: LD A,d8
func ldAD8(c *CPU) {
	c.A = c.Read()
}

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

// 4F: LD C,A
func ldCA(c *CPU) {
	c.C = c.A
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

// 77: LD (HL),A
func ldAddrHlA(c *CPU) {
	ldAddrHlR(c, c.A)
}

// 78: LD A,B
func ldAB(c *CPU) {
	c.A = c.B
}

// 79: LD A,C
func ldAC(c *CPU) {
	c.A = c.C
}

// 7A: LD A,D
func ldAD(c *CPU) {
	c.A = c.D
}

// 7B: LD A,E
func ldAE(c *CPU) {
	c.A = c.E
}

// 7C: LD A,H
func ldAH(c *CPU) {
	c.A = c.H
}

// 7D: LD A,L
func ldAL(c *CPU) {
	c.A = c.L
}

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
	addAR(c, c.MMU.Read(uint(c.HL())))
}

// 87: ADD A,A
func addAA(c *CPU) {
	addAR(c, c.A)
}

// AF: XOR A
func xorA(c *CPU) {
	xorR(c, &c.A)
}

// C1: POP BC
func popBc(c *CPU) {
	popRr(c, &c.B, &c.C)
}

// C5: PUSH BC
func pushBc(c *CPU) {
	pushRr(c, c.B, c.C)
}

// C9: RET
func ret(c *CPU) {
	// Simulate POP PC for consistency
	var P, C uint8
	popRr(c, &P, &C)
	c.PC = uint16(P)<<8 | uint16(C)
}

// CB 11: RL C
func rlC(c *CPU) {
	rlR(c, &c.C)
}

// CB 7C: BIT 7,H
func bit7H(c *CPU) {
	bitNR(c, 7, c.H)
}

// CD: CALL a16
func callA16(c *CPU) {
	// Advance PC before pushing
	addr := c.ReadWord()
	pushRr(c, uint8(c.PC>>8), uint8(c.PC&0xff))
	c.PC = addr
}

// C1: POP DE
func popDe(c *CPU) {
	popRr(c, &c.D, &c.E)
}

// D5: PUSH DE
func pushDe(c *CPU) {
	pushRr(c, c.D, c.E)
}

// E0: LD (FF00+a8),A
func ldAddrFfA8A(c *CPU) {
	c.MMU.Write(uint(0xff00+uint16(c.Read())), c.A)
}

// E1: POP HL
func popHl(c *CPU) {
	popRr(c, &c.H, &c.L)
}

// E2: LD (FF00+C),A
func ldAddrFfCA(c *CPU) {
	c.MMU.Write(uint(0xff00+uint16(c.C)), c.A)
}

// E5: PUSH HL
func pushHl(c *CPU) {
	pushRr(c, c.H, c.L)
}

// F1: POP AF
func popAf(c *CPU) {
	popRr(c, &c.A, &c.F)
}

// F5: PUSH AF
func pushAf(c *CPU) {
	pushRr(c, c.A, c.F)
}

// FE: CP d8
func cpD8(c *CPU) {
	subFlags(c, c.Read())
}
