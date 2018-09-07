// Generated code. See instructions.go
package cpu

// LR35902InstructionSet is an array of instrutions for the DMG CPU.
var LR35902InstructionSet = [...]Instruction{
	0x00: op00{},
	0x01: op01{},
	0x02: op02{},
	0x03: op03{},
	0x04: op04{},
	0x05: op05{},
	0x06: op06{},
	0x07: op07{},
	0x0c: op0c{},
	0x0d: op0d{},
	0x0e: op0e{},
	0x11: op11{},
	0x13: op13{},
	0x14: op14{},
	0x15: op15{},
	0x16: op16{},
	0x1c: op1c{},
	0x1d: op1d{},
	0x1e: op1e{},
	0x21: op21{},
	0x23: op23{},
	0x24: op24{},
	0x25: op25{},
	0x26: op26{},
	0x2c: op2c{},
	0x2d: op2d{},
	0x2e: op2e{},
	0x3c: op3c{},
	0x3d: op3d{},
	0x3e: op3e{},
	0xa0: opA0{},
	0xa1: opA1{},
	0xa2: opA2{},
	0xa3: opA3{},
	0xa4: opA4{},
	0xa5: opA5{},
	0xa6: opA6{},
}

// LR35902ExtendedInstructionSet is the array of extension opcodes for the DMG CPU.
var LR35902ExtendedInstructionSet = []Instruction{}

// 00: NOP			4 cycles
type op00 struct {
	SingleStepOp
}

func (op op00) Execute(c *CPU) (done bool) {
	return true
}

// 01: LD BC,d16		(12 cycles)
type op01 struct {
	MultiStepsOp
}

func (op op01) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.C = op.cpu.NextByte()
		op.step++
	case 1:
		op.cpu.B = op.cpu.NextByte()
		done = true
	}
	return
}

// 02: LD (BC),A		(8 cycles)
type op02 struct {
	MultiStepsOp
}

func (op op02) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.BC()), op.cpu.A)
	return true
}

// 03: INC BC			8 cycles
type op03 struct {
	MultiStepsOp
}

func (op op03) Tick() (done bool) {
	if op.cpu.C == 0xff {
		op.cpu.B++
	}
	op.cpu.C++
	return true
}

// 04: INC B			4 cycles
type op04 struct {
	SingleStepOp
}

func (op op04) Execute(c *CPU) (done bool) {
	// Flags z 0 h -
	c.F &= ^FlagN
	if c.B > 0x0F {
		c.F |= FlagH
	}
	c.B++
	if c.B == 0 {
		c.F |= FlagZ
	}
	return true
}

// 05: DEC B			4 cycles
type op05 struct {
	SingleStepOp
}

func (op op05) Execute(c *CPU) (done bool) {
	// Flags z 1 h -
	c.F &= FlagC
	c.F |= FlagN
	if c.B > 0x0F {
		c.F |= FlagH
	}
	c.B--
	if c.B == 0 {
		c.F |= FlagZ
	}
	return true
}

// 06: LD B,d8		(8 cycles)
type op06 struct {
	MultiStepsOp
}

func (op op06) Tick() (done bool) {
	op.cpu.B = op.cpu.NextByte()
	return true
}

// 07: RLCA			4 cycles
type op07 struct {
	SingleStepOp
}

func (op op07) Execute(c *CPU) (done bool) {
	// FIXME: [OPCODES] says flags are 0 0 0 c. Couldn't confirm.
	// Flags z 0 0 c
	c.F = 0x00
	result := c.A << 1 & 0xff
	if c.A&0x80 != 0 {
		result |= 1
		c.F |= FlagC
	}

	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result
	return true
}

// 0C: INC C			4 cycles
type op0c struct {
	SingleStepOp
}

func (op op0c) Execute(c *CPU) (done bool) {
	// Flags z 0 h -
	c.F &= ^FlagN
	if c.C > 0x0F {
		c.F |= FlagH
	}
	c.C++
	if c.C == 0 {
		c.F |= FlagZ
	}
	return true
}

// 0D: DEC C			4 cycles
type op0d struct {
	SingleStepOp
}

func (op op0d) Execute(c *CPU) (done bool) {
	// Flags z 1 h -
	c.F &= FlagC
	c.F |= FlagN
	if c.C > 0x0F {
		c.F |= FlagH
	}
	c.C--
	if c.C == 0 {
		c.F |= FlagZ
	}
	return true
}

// 0E: LD C,d8		(8 cycles)
type op0e struct {
	MultiStepsOp
}

func (op op0e) Tick() (done bool) {
	op.cpu.C = op.cpu.NextByte()
	return true
}

// 11: LD DE,d16		(12 cycles)
type op11 struct {
	MultiStepsOp
}

func (op op11) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.E = op.cpu.NextByte()
		op.step++
	case 1:
		op.cpu.D = op.cpu.NextByte()
		done = true
	}
	return
}

// 13: INC DE			8 cycles
type op13 struct {
	MultiStepsOp
}

func (op op13) Tick() (done bool) {
	if op.cpu.E == 0xff {
		op.cpu.D++
	}
	op.cpu.E++
	return true
}

// 14: INC D			4 cycles
type op14 struct {
	SingleStepOp
}

func (op op14) Execute(c *CPU) (done bool) {
	// Flags z 0 h -
	c.F &= ^FlagN
	if c.D > 0x0F {
		c.F |= FlagH
	}
	c.D++
	if c.D == 0 {
		c.F |= FlagZ
	}
	return true
}

// 15: DEC D			4 cycles
type op15 struct {
	SingleStepOp
}

func (op op15) Execute(c *CPU) (done bool) {
	// Flags z 1 h -
	c.F &= FlagC
	c.F |= FlagN
	if c.D > 0x0F {
		c.F |= FlagH
	}
	c.D--
	if c.D == 0 {
		c.F |= FlagZ
	}
	return true
}

// 16: LD D,d8		(8 cycles)
type op16 struct {
	MultiStepsOp
}

func (op op16) Tick() (done bool) {
	op.cpu.D = op.cpu.NextByte()
	return true
}

// 1C: INC E			4 cycles
type op1c struct {
	SingleStepOp
}

func (op op1c) Execute(c *CPU) (done bool) {
	// Flags z 0 h -
	c.F &= ^FlagN
	if c.E > 0x0F {
		c.F |= FlagH
	}
	c.E++
	if c.E == 0 {
		c.F |= FlagZ
	}
	return true
}

// 1D: DEC E			4 cycles
type op1d struct {
	SingleStepOp
}

func (op op1d) Execute(c *CPU) (done bool) {
	// Flags z 1 h -
	c.F &= FlagC
	c.F |= FlagN
	if c.E > 0x0F {
		c.F |= FlagH
	}
	c.E--
	if c.E == 0 {
		c.F |= FlagZ
	}
	return true
}

// 1E: LD E,d8		(8 cycles)
type op1e struct {
	MultiStepsOp
}

func (op op1e) Tick() (done bool) {
	op.cpu.E = op.cpu.NextByte()
	return true
}

// 21: LD HL,d16		(12 cycles)
type op21 struct {
	MultiStepsOp
}

func (op op21) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.L = op.cpu.NextByte()
		op.step++
	case 1:
		op.cpu.H = op.cpu.NextByte()
		done = true
	}
	return
}

// 23: INC HL			8 cycles
type op23 struct {
	MultiStepsOp
}

func (op op23) Tick() (done bool) {
	if op.cpu.L == 0xff {
		op.cpu.H++
	}
	op.cpu.L++
	return true
}

// 24: INC H			4 cycles
type op24 struct {
	SingleStepOp
}

func (op op24) Execute(c *CPU) (done bool) {
	// Flags z 0 h -
	c.F &= ^FlagN
	if c.H > 0x0F {
		c.F |= FlagH
	}
	c.H++
	if c.H == 0 {
		c.F |= FlagZ
	}
	return true
}

// 25: DEC H			4 cycles
type op25 struct {
	SingleStepOp
}

func (op op25) Execute(c *CPU) (done bool) {
	// Flags z 1 h -
	c.F &= FlagC
	c.F |= FlagN
	if c.H > 0x0F {
		c.F |= FlagH
	}
	c.H--
	if c.H == 0 {
		c.F |= FlagZ
	}
	return true
}

// 26: LD H,d8		(8 cycles)
type op26 struct {
	MultiStepsOp
}

func (op op26) Tick() (done bool) {
	op.cpu.H = op.cpu.NextByte()
	return true
}

// 2C: INC L			4 cycles
type op2c struct {
	SingleStepOp
}

func (op op2c) Execute(c *CPU) (done bool) {
	// Flags z 0 h -
	c.F &= ^FlagN
	if c.L > 0x0F {
		c.F |= FlagH
	}
	c.L++
	if c.L == 0 {
		c.F |= FlagZ
	}
	return true
}

// 2D: DEC L			4 cycles
type op2d struct {
	SingleStepOp
}

func (op op2d) Execute(c *CPU) (done bool) {
	// Flags z 1 h -
	c.F &= FlagC
	c.F |= FlagN
	if c.L > 0x0F {
		c.F |= FlagH
	}
	c.L--
	if c.L == 0 {
		c.F |= FlagZ
	}
	return true
}

// 2E: LD L,d8		(8 cycles)
type op2e struct {
	MultiStepsOp
}

func (op op2e) Tick() (done bool) {
	op.cpu.L = op.cpu.NextByte()
	return true
}

// 3C: INC A			4 cycles
type op3c struct {
	SingleStepOp
}

func (op op3c) Execute(c *CPU) (done bool) {
	// Flags z 0 h -
	c.F &= ^FlagN
	if c.A > 0x0F {
		c.F |= FlagH
	}
	c.A++
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// 3D: DEC A			4 cycles
type op3d struct {
	SingleStepOp
}

func (op op3d) Execute(c *CPU) (done bool) {
	// Flags z 1 h -
	c.F &= FlagC
	c.F |= FlagN
	if c.A > 0x0F {
		c.F |= FlagH
	}
	c.A--
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// 3E: LD A,d8		(8 cycles)
type op3e struct {
	MultiStepsOp
}

func (op op3e) Tick() (done bool) {
	op.cpu.A = op.cpu.NextByte()
	return true
}

// A0: AND B
type opA0 struct {
	SingleStepOp
}

func (op opA0) Execute(c *CPU) (done bool) {
	c.A &= c.B
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A1: AND C
type opA1 struct {
	SingleStepOp
}

func (op opA1) Execute(c *CPU) (done bool) {
	c.A &= c.C
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A2: AND D
type opA2 struct {
	SingleStepOp
}

func (op opA2) Execute(c *CPU) (done bool) {
	c.A &= c.D
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A3: AND E
type opA3 struct {
	SingleStepOp
}

func (op opA3) Execute(c *CPU) (done bool) {
	c.A &= c.E
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A4: AND H
type opA4 struct {
	SingleStepOp
}

func (op opA4) Execute(c *CPU) (done bool) {
	c.A &= c.H
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A5: AND L
type opA5 struct {
	SingleStepOp
}

func (op opA5) Execute(c *CPU) (done bool) {
	c.A &= c.L
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A6: AND (HL)
type opA6 struct {
	MultiStepsOp
}

func (op opA6) Tick() (done bool) {
	op.cpu.A &= op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 0 0
	if op.cpu.A == 0 {
		op.cpu.F = FlagZ
	} else {
		op.cpu.F = 0
	}
	return true
}
