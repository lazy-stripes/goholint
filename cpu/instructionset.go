// Auto-generated on 2018-10-25 13:25:16.624381248 +0200 CEST. See instructions.go
package cpu

// LR35902InstructionSet is an array of instrutions for the DMG CPU.
var LR35902InstructionSet = [...]Instruction{
	0x00: &op00{},
	0x01: &op01{},
	0x02: &op02{},
	0x03: &op03{},
	0x04: &op04{},
	0x05: &op05{},
	0x06: &op06{},
	0x07: &op07{},
	0x08: &op08{},
	0x09: &op09{},
	0x0a: &op0a{},
	0x0b: &op0b{},
	0x0c: &op0c{},
	0x0d: &op0d{},
	0x0e: &op0e{},
	0x11: &op11{},
	0x12: &op12{},
	0x13: &op13{},
	0x14: &op14{},
	0x15: &op15{},
	0x16: &op16{},
	0x17: &op17{},
	0x18: &op18{},
	0x19: &op19{},
	0x1a: &op1a{},
	0x1b: &op1b{},
	0x1c: &op1c{},
	0x1d: &op1d{},
	0x1e: &op1e{},
	0x20: &op20{},
	0x21: &op21{},
	0x22: &op22{},
	0x23: &op23{},
	0x24: &op24{},
	0x25: &op25{},
	0x26: &op26{},
	0x28: &op28{},
	0x29: &op29{},
	0x2a: &op2a{},
	0x2b: &op2b{},
	0x2c: &op2c{},
	0x2d: &op2d{},
	0x2e: &op2e{},
	0x30: &op30{},
	0x31: &op31{},
	0x32: &op32{},
	0x38: &op38{},
	0x39: &op39{},
	0x3a: &op3a{},
	0x3b: &op3b{},
	0x3c: &op3c{},
	0x3d: &op3d{},
	0x3e: &op3e{},
	0x40: &op40{},
	0x41: &op41{},
	0x42: &op42{},
	0x43: &op43{},
	0x44: &op44{},
	0x45: &op45{},
	0x46: &op46{},
	0x47: &op47{},
	0x48: &op48{},
	0x49: &op49{},
	0x4a: &op4a{},
	0x4b: &op4b{},
	0x4c: &op4c{},
	0x4d: &op4d{},
	0x4e: &op4e{},
	0x4f: &op4f{},
	0x50: &op50{},
	0x51: &op51{},
	0x52: &op52{},
	0x53: &op53{},
	0x54: &op54{},
	0x55: &op55{},
	0x56: &op56{},
	0x57: &op57{},
	0x58: &op58{},
	0x59: &op59{},
	0x5a: &op5a{},
	0x5b: &op5b{},
	0x5c: &op5c{},
	0x5d: &op5d{},
	0x5e: &op5e{},
	0x5f: &op5f{},
	0x60: &op60{},
	0x61: &op61{},
	0x62: &op62{},
	0x63: &op63{},
	0x64: &op64{},
	0x65: &op65{},
	0x66: &op66{},
	0x67: &op67{},
	0x68: &op68{},
	0x69: &op69{},
	0x6a: &op6a{},
	0x6b: &op6b{},
	0x6c: &op6c{},
	0x6d: &op6d{},
	0x6e: &op6e{},
	0x6f: &op6f{},
	0x70: &op70{},
	0x71: &op71{},
	0x72: &op72{},
	0x73: &op73{},
	0x74: &op74{},
	0x75: &op75{},
	0x77: &op77{},
	0x78: &op78{},
	0x79: &op79{},
	0x7a: &op7a{},
	0x7b: &op7b{},
	0x7c: &op7c{},
	0x7d: &op7d{},
	0x7e: &op7e{},
	0x7f: &op7f{},
	0x80: &op80{},
	0x81: &op81{},
	0x82: &op82{},
	0x83: &op83{},
	0x84: &op84{},
	0x85: &op85{},
	0x86: &op86{},
	0x87: &op87{},
	0x90: &op90{},
	0x91: &op91{},
	0x92: &op92{},
	0x93: &op93{},
	0x94: &op94{},
	0x95: &op95{},
	0x96: &op96{},
	0x97: &op97{},
	0xa0: &opA0{},
	0xa1: &opA1{},
	0xa2: &opA2{},
	0xa3: &opA3{},
	0xa4: &opA4{},
	0xa5: &opA5{},
	0xa6: &opA6{},
	0xa7: &opA7{},
	0xa8: &opA8{},
	0xa9: &opA9{},
	0xaa: &opAa{},
	0xab: &opAb{},
	0xac: &opAc{},
	0xad: &opAd{},
	0xae: &opAe{},
	0xaf: &opAf{},
	0xb0: &opB0{},
	0xb1: &opB1{},
	0xb2: &opB2{},
	0xb3: &opB3{},
	0xb4: &opB4{},
	0xb5: &opB5{},
	0xb6: &opB6{},
	0xb7: &opB7{},
	0xb8: &opB8{},
	0xb9: &opB9{},
	0xba: &opBa{},
	0xbb: &opBb{},
	0xbc: &opBc{},
	0xbd: &opBd{},
	0xbe: &opBe{},
	0xbf: &opBf{},
	0xc1: &opC1{},
	0xc5: &opC5{},
	0xc9: &opC9{},
	0xcd: &opCd{},
	0xd1: &opD1{},
	0xd5: &opD5{},
	0xd9: &opD9{},
	0xe0: &opE0{},
	0xe1: &opE1{},
	0xe2: &opE2{},
	0xe5: &opE5{},
	0xea: &opEa{},
	0xf0: &opF0{},
	0xf1: &opF1{},
	0xf5: &opF5{},
	0xfe: &opFe{},
}

// LR35902ExtendedInstructionSet is the array of extension opcodes for the DMG CPU.
var LR35902ExtendedInstructionSet = [...]Instruction{
	0x10: &opCb10{},
	0x11: &opCb11{},
	0x12: &opCb12{},
	0x13: &opCb13{},
	0x14: &opCb14{},
	0x15: &opCb15{},
	0x17: &opCb17{},
	0x40: &opCb40{},
	0x41: &opCb41{},
	0x42: &opCb42{},
	0x43: &opCb43{},
	0x44: &opCb44{},
	0x45: &opCb45{},
	0x46: &opCb46{},
	0x47: &opCb47{},
	0x48: &opCb48{},
	0x49: &opCb49{},
	0x4a: &opCb4a{},
	0x4b: &opCb4b{},
	0x4c: &opCb4c{},
	0x4d: &opCb4d{},
	0x4e: &opCb4e{},
	0x4f: &opCb4f{},
	0x50: &opCb50{},
	0x51: &opCb51{},
	0x52: &opCb52{},
	0x53: &opCb53{},
	0x54: &opCb54{},
	0x55: &opCb55{},
	0x56: &opCb56{},
	0x57: &opCb57{},
	0x58: &opCb58{},
	0x59: &opCb59{},
	0x5a: &opCb5a{},
	0x5b: &opCb5b{},
	0x5c: &opCb5c{},
	0x5d: &opCb5d{},
	0x5e: &opCb5e{},
	0x5f: &opCb5f{},
	0x60: &opCb60{},
	0x61: &opCb61{},
	0x62: &opCb62{},
	0x63: &opCb63{},
	0x64: &opCb64{},
	0x65: &opCb65{},
	0x66: &opCb66{},
	0x67: &opCb67{},
	0x68: &opCb68{},
	0x69: &opCb69{},
	0x6a: &opCb6a{},
	0x6b: &opCb6b{},
	0x6c: &opCb6c{},
	0x6d: &opCb6d{},
	0x6e: &opCb6e{},
	0x6f: &opCb6f{},
	0x70: &opCb70{},
	0x71: &opCb71{},
	0x72: &opCb72{},
	0x73: &opCb73{},
	0x74: &opCb74{},
	0x75: &opCb75{},
	0x76: &opCb76{},
	0x77: &opCb77{},
	0x78: &opCb78{},
	0x79: &opCb79{},
	0x7a: &opCb7a{},
	0x7b: &opCb7b{},
	0x7c: &opCb7c{},
	0x7d: &opCb7d{},
	0x7e: &opCb7e{},
	0x7f: &opCb7f{},
}

// 00: NOP				4 cycles
type op00 struct {
	SingleStepOp
}

func (op *op00) Execute(c *CPU) (done bool) {
	return true
}

// 01: LD BC,d16		12 cycles
type op01 struct {
	MultiStepsOp
}

func (op *op01) Tick() (done bool) {
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

// 02: LD (BC),A		8 cycles
type op02 struct {
	MultiStepsOp
}

func (op *op02) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.BC()), op.cpu.A)
	return true
}

// 03: INC BC			8 cycles
type op03 struct {
	MultiStepsOp
}

func (op *op03) Tick() (done bool) {
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

func (op *op04) Execute(c *CPU) (done bool) {
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

func (op *op05) Execute(c *CPU) (done bool) {
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

// 06: LD B,d8			8 cycles
type op06 struct {
	MultiStepsOp
}

func (op *op06) Tick() (done bool) {
	op.cpu.B = op.cpu.NextByte()
	return true
}

// 07: RLCA				4 cycles
type op07 struct {
	SingleStepOp
}

func (op *op07) Execute(c *CPU) (done bool) {
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

// 08: LD (a16),SP		20 cycles
type op08 struct {
	MultiStepsOp
}

func (op *op08) Tick() (done bool) {
	switch op.step {
	case 0:
		// XXX: template snippet for opReadD16Low/high
		op.cpu.temp16 = uint16(op.cpu.NextByte())
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
	case 2:
		op.cpu.MMU.Write(uint(op.cpu.temp16), uint8(op.cpu.SP&0xff))
	case 3:
		op.cpu.MMU.Write(uint(op.cpu.temp16), uint8(op.cpu.SP>>8))
		done = true
	}
	op.step++
	return
}

// 09: ADD HL,BC		8 cycles
type op09 struct {
	MultiStepsOp
}

func (op *op09) Tick() (done bool) {
	// Flags: - 0 h c
	op.cpu.F &= FlagZ
	hl := uint(op.cpu.HL())
    rr := uint(op.cpu.C) | uint(op.cpu.B)<<8
	if hl&0xfff+rr&0xfff > 0xfff {
		op.cpu.F |= FlagH
	}
	result := hl + rr
	if result > 0xffff {
		op.cpu.F |= FlagC
	}
	op.cpu.SetHL(uint16(result & 0xffff))
    return true
}

// 0A: LD (BC),A		8 cycles
type op0a struct {
	MultiStepsOp
}

func (op *op0a) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.BC()), op.cpu.A)
	return true
}

// 0B: DEC BC			8 cycles
type op0b struct {
	MultiStepsOp
}

func (op *op0b) Tick() (done bool) {
	if op.cpu.C == 0x00 {
		op.cpu.B--
	}
	op.cpu.C--
	return true
}

// 0C: INC C			4 cycles
type op0c struct {
	SingleStepOp
}

func (op *op0c) Execute(c *CPU) (done bool) {
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

func (op *op0d) Execute(c *CPU) (done bool) {
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

// 0E: LD C,d8			8 cycles
type op0e struct {
	MultiStepsOp
}

func (op *op0e) Tick() (done bool) {
	op.cpu.C = op.cpu.NextByte()
	return true
}

// 11: LD DE,d16		12 cycles
type op11 struct {
	MultiStepsOp
}

func (op *op11) Tick() (done bool) {
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

// 12: LD (DE),A		8 cycles
type op12 struct {
	MultiStepsOp
}

func (op *op12) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.DE()), op.cpu.A)
	return true
}

// 13: INC DE			8 cycles
type op13 struct {
	MultiStepsOp
}

func (op *op13) Tick() (done bool) {
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

func (op *op14) Execute(c *CPU) (done bool) {
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

func (op *op15) Execute(c *CPU) (done bool) {
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

// 16: LD D,d8			8 cycles
type op16 struct {
	MultiStepsOp
}

func (op *op16) Tick() (done bool) {
	op.cpu.D = op.cpu.NextByte()
	return true
}

// 17: RLA				4 cycles
type op17 struct {
	SingleStepOp
}

func (op *op17) Execute(c *CPU) (done bool) {
	// FIXME: [OPCODES] says flags are 0 0 0 c. Couldn't confirm.
	result := c.A << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.A&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.A = result

	return true
}

// 18: JR r8		12 cycles
type op18 struct {
	MultiStepsOp
	offset int8
}

func (op *op18) Tick() (done bool) {
	switch op.step {
	case 0:
		op.offset = int8(op.cpu.NextByte())
		op.step++
	case 1:
		// Need cast to signed for the potential substraction
		op.cpu.PC = uint16(int16(op.cpu.PC) + int16(op.offset))
		done = true
	}
	return
}

// 19: ADD HL,DE		8 cycles
type op19 struct {
	MultiStepsOp
}

func (op *op19) Tick() (done bool) {
	// Flags: - 0 h c
	op.cpu.F &= FlagZ
	hl := uint(op.cpu.HL())
    rr := uint(op.cpu.E) | uint(op.cpu.D)<<8
	if hl&0xfff+rr&0xfff > 0xfff {
		op.cpu.F |= FlagH
	}
	result := hl + rr
	if result > 0xffff {
		op.cpu.F |= FlagC
	}
	op.cpu.SetHL(uint16(result & 0xffff))
    return true
}

// 1A: LD A,(DE)			8 cycles
type op1a struct {
	MultiStepsOp
}

func (op *op1a) Tick() (done bool) {
	op.cpu.A = op.cpu.MMU.Read(uint(op.cpu.DE()))
	return true
}

// 1B: DEC DE			8 cycles
type op1b struct {
	MultiStepsOp
}

func (op *op1b) Tick() (done bool) {
	if op.cpu.E == 0x00 {
		op.cpu.D--
	}
	op.cpu.E--
	return true
}

// 1C: INC E			4 cycles
type op1c struct {
	SingleStepOp
}

func (op *op1c) Execute(c *CPU) (done bool) {
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

func (op *op1d) Execute(c *CPU) (done bool) {
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

// 1E: LD E,d8			8 cycles
type op1e struct {
	MultiStepsOp
}

func (op *op1e) Tick() (done bool) {
	op.cpu.E = op.cpu.NextByte()
	return true
}

// 20: JR NZ,r8		12/8 cycles
type op20 struct {
	MultiStepsOp
	offset int8
}

func (op *op20) Tick() (done bool) {
	switch op.step {
	case 0:
		op.offset = int8(op.cpu.NextByte())
		if op.cpu.F&FlagZ != FlagZ {
			op.step++
		} else {
			done = true
		}
	case 1:
		// Need cast to signed for the potential substraction
		op.cpu.PC = uint16(int16(op.cpu.PC) + int16(op.offset))
		done = true
	}
	return
}

// 21: LD HL,d16		12 cycles
type op21 struct {
	MultiStepsOp
}

func (op *op21) Tick() (done bool) {
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

// 22: LD (HL+),A			8 cycles
type op22 struct {
	MultiStepsOp
}

func (op *op22) Tick() (done bool) {
	hl := op.cpu.HL()
	op.cpu.MMU.Write(uint(hl), op.cpu.A)
	op.cpu.SetHL(hl+1)
	return true
}

// 23: INC HL			8 cycles
type op23 struct {
	MultiStepsOp
}

func (op *op23) Tick() (done bool) {
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

func (op *op24) Execute(c *CPU) (done bool) {
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

func (op *op25) Execute(c *CPU) (done bool) {
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

// 26: LD H,d8			8 cycles
type op26 struct {
	MultiStepsOp
}

func (op *op26) Tick() (done bool) {
	op.cpu.H = op.cpu.NextByte()
	return true
}

// 28: JR Z,r8		12/8 cycles
type op28 struct {
	MultiStepsOp
	offset int8
}

func (op *op28) Tick() (done bool) {
	switch op.step {
	case 0:
		op.offset = int8(op.cpu.NextByte())
		if op.cpu.F&FlagZ == FlagZ {
			op.step++
		} else {
			done = true
		}
	case 1:
		// Need cast to signed for the potential substraction
		op.cpu.PC = uint16(int16(op.cpu.PC) + int16(op.offset))
		done = true
	}
	return
}

// 29: ADD HL,HL		8 cycles
type op29 struct {
	MultiStepsOp
}

func (op *op29) Tick() (done bool) {
	// Flags: - 0 h c
	op.cpu.F &= FlagZ
	hl := uint(op.cpu.HL())
    rr := uint(op.cpu.L) | uint(op.cpu.H)<<8
	if hl&0xfff+rr&0xfff > 0xfff {
		op.cpu.F |= FlagH
	}
	result := hl + rr
	if result > 0xffff {
		op.cpu.F |= FlagC
	}
	op.cpu.SetHL(uint16(result & 0xffff))
    return true
}

// 2A: LD A,(HL+)			8 cycles
type op2a struct {
	MultiStepsOp
}

func (op *op2a) Tick() (done bool) {
	hl := op.cpu.HL()
	op.cpu.A = op.cpu.MMU.Read(uint(hl))
	op.cpu.SetHL(hl+1)
	return true
}

// 2B: DEC HL			8 cycles
type op2b struct {
	MultiStepsOp
}

func (op *op2b) Tick() (done bool) {
	if op.cpu.L == 0x00 {
		op.cpu.H--
	}
	op.cpu.L--
	return true
}

// 2C: INC L			4 cycles
type op2c struct {
	SingleStepOp
}

func (op *op2c) Execute(c *CPU) (done bool) {
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

func (op *op2d) Execute(c *CPU) (done bool) {
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

// 2E: LD L,d8			8 cycles
type op2e struct {
	MultiStepsOp
}

func (op *op2e) Tick() (done bool) {
	op.cpu.L = op.cpu.NextByte()
	return true
}

// 30: JR NC,r8		12/8 cycles
type op30 struct {
	MultiStepsOp
	offset int8
}

func (op *op30) Tick() (done bool) {
	switch op.step {
	case 0:
		op.offset = int8(op.cpu.NextByte())
		if op.cpu.F&FlagC != FlagC {
			op.step++
		} else {
			done = true
		}
	case 1:
		// Need cast to signed for the potential substraction
		op.cpu.PC = uint16(int16(op.cpu.PC) + int16(op.offset))
		done = true
	}
	return
}

// 31: LD SP,d16		12 cycles
type op31 struct {
	MultiStepsOp
}

func (op *op31) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.SP |= uint16(op.cpu.NextByte()) << 8
		done = true
	}
	return
}

// 32: LD (HL-),A			8 cycles
type op32 struct {
	MultiStepsOp
}

func (op *op32) Tick() (done bool) {
	hl := op.cpu.HL()
	op.cpu.MMU.Write(uint(hl), op.cpu.A)
	op.cpu.SetHL(hl-1)
	return true
}

// 38: JR C,r8		12/8 cycles
type op38 struct {
	MultiStepsOp
	offset int8
}

func (op *op38) Tick() (done bool) {
	switch op.step {
	case 0:
		op.offset = int8(op.cpu.NextByte())
		if op.cpu.F&FlagC == FlagC {
			op.step++
		} else {
			done = true
		}
	case 1:
		// Need cast to signed for the potential substraction
		op.cpu.PC = uint16(int16(op.cpu.PC) + int16(op.offset))
		done = true
	}
	return
}

// 39: ADD HL,SP		8 cycles
type op39 struct {
	MultiStepsOp
}

func (op *op39) Tick() (done bool) {
	// Flags: - 0 h c
	op.cpu.F &= FlagZ
	hl := uint(op.cpu.HL())
	rr := uint(op.cpu.SP)
	if hl&0xfff+rr&0xfff > 0xfff {
		op.cpu.F |= FlagH
	}
	result := hl + rr
	if result > 0xffff {
		op.cpu.F |= FlagC
	}
	op.cpu.SetHL(uint16(result & 0xffff))
    return true
}

// 3A: LD A,(HL-)			8 cycles
type op3a struct {
	MultiStepsOp
}

func (op *op3a) Tick() (done bool) {
	hl := op.cpu.HL()
	op.cpu.A = op.cpu.MMU.Read(uint(hl))
	op.cpu.SetHL(hl-1)
	return true
}

// 3B: DEC SP			8 cycles
type op3b struct {
	MultiStepsOp
}

func (op *op3b) Tick() (done bool) {
	op.cpu.SP--
	return true
}

// 3C: INC A			4 cycles
type op3c struct {
	SingleStepOp
}

func (op *op3c) Execute(c *CPU) (done bool) {
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

func (op *op3d) Execute(c *CPU) (done bool) {
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

// 3E: LD A,d8			8 cycles
type op3e struct {
	MultiStepsOp
}

func (op *op3e) Tick() (done bool) {
	op.cpu.A = op.cpu.NextByte()
	return true
}

// 40: LD B,B			4 cycles
type op40 struct {
	SingleStepOp
}

func (op *op40) Execute(c *CPU) (done bool) {
	c.B = c.B
	return true
}

// 41: LD B,C			4 cycles
type op41 struct {
	SingleStepOp
}

func (op *op41) Execute(c *CPU) (done bool) {
	c.B = c.C
	return true
}

// 42: LD B,D			4 cycles
type op42 struct {
	SingleStepOp
}

func (op *op42) Execute(c *CPU) (done bool) {
	c.B = c.D
	return true
}

// 43: LD B,E			4 cycles
type op43 struct {
	SingleStepOp
}

func (op *op43) Execute(c *CPU) (done bool) {
	c.B = c.E
	return true
}

// 44: LD B,H			4 cycles
type op44 struct {
	SingleStepOp
}

func (op *op44) Execute(c *CPU) (done bool) {
	c.B = c.H
	return true
}

// 45: LD B,L			4 cycles
type op45 struct {
	SingleStepOp
}

func (op *op45) Execute(c *CPU) (done bool) {
	c.B = c.L
	return true
}

// 46: LD B,(HL)			8 cycles
type op46 struct {
	MultiStepsOp
}

func (op *op46) Tick() (done bool) {
	op.cpu.B = op.cpu.MMU.Read(uint(op.cpu.HL()))
	return true
}

// 47: LD B,A			4 cycles
type op47 struct {
	SingleStepOp
}

func (op *op47) Execute(c *CPU) (done bool) {
	c.B = c.A
	return true
}

// 48: LD C,B			4 cycles
type op48 struct {
	SingleStepOp
}

func (op *op48) Execute(c *CPU) (done bool) {
	c.C = c.B
	return true
}

// 49: LD C,C			4 cycles
type op49 struct {
	SingleStepOp
}

func (op *op49) Execute(c *CPU) (done bool) {
	c.C = c.C
	return true
}

// 4A: LD C,D			4 cycles
type op4a struct {
	SingleStepOp
}

func (op *op4a) Execute(c *CPU) (done bool) {
	c.C = c.D
	return true
}

// 4B: LD C,E			4 cycles
type op4b struct {
	SingleStepOp
}

func (op *op4b) Execute(c *CPU) (done bool) {
	c.C = c.E
	return true
}

// 4C: LD C,H			4 cycles
type op4c struct {
	SingleStepOp
}

func (op *op4c) Execute(c *CPU) (done bool) {
	c.C = c.H
	return true
}

// 4D: LD C,L			4 cycles
type op4d struct {
	SingleStepOp
}

func (op *op4d) Execute(c *CPU) (done bool) {
	c.C = c.L
	return true
}

// 4E: LD C,(HL)			8 cycles
type op4e struct {
	MultiStepsOp
}

func (op *op4e) Tick() (done bool) {
	op.cpu.C = op.cpu.MMU.Read(uint(op.cpu.HL()))
	return true
}

// 4F: LD C,A			4 cycles
type op4f struct {
	SingleStepOp
}

func (op *op4f) Execute(c *CPU) (done bool) {
	c.C = c.A
	return true
}

// 50: LD D,B			4 cycles
type op50 struct {
	SingleStepOp
}

func (op *op50) Execute(c *CPU) (done bool) {
	c.D = c.B
	return true
}

// 51: LD D,C			4 cycles
type op51 struct {
	SingleStepOp
}

func (op *op51) Execute(c *CPU) (done bool) {
	c.D = c.C
	return true
}

// 52: LD D,D			4 cycles
type op52 struct {
	SingleStepOp
}

func (op *op52) Execute(c *CPU) (done bool) {
	c.D = c.D
	return true
}

// 53: LD D,E			4 cycles
type op53 struct {
	SingleStepOp
}

func (op *op53) Execute(c *CPU) (done bool) {
	c.D = c.E
	return true
}

// 54: LD D,H			4 cycles
type op54 struct {
	SingleStepOp
}

func (op *op54) Execute(c *CPU) (done bool) {
	c.D = c.H
	return true
}

// 55: LD D,L			4 cycles
type op55 struct {
	SingleStepOp
}

func (op *op55) Execute(c *CPU) (done bool) {
	c.D = c.L
	return true
}

// 56: LD D,(HL)			8 cycles
type op56 struct {
	MultiStepsOp
}

func (op *op56) Tick() (done bool) {
	op.cpu.D = op.cpu.MMU.Read(uint(op.cpu.HL()))
	return true
}

// 57: LD D,A			4 cycles
type op57 struct {
	SingleStepOp
}

func (op *op57) Execute(c *CPU) (done bool) {
	c.D = c.A
	return true
}

// 58: LD E,B			4 cycles
type op58 struct {
	SingleStepOp
}

func (op *op58) Execute(c *CPU) (done bool) {
	c.E = c.B
	return true
}

// 59: LD E,C			4 cycles
type op59 struct {
	SingleStepOp
}

func (op *op59) Execute(c *CPU) (done bool) {
	c.E = c.C
	return true
}

// 5A: LD E,D			4 cycles
type op5a struct {
	SingleStepOp
}

func (op *op5a) Execute(c *CPU) (done bool) {
	c.E = c.D
	return true
}

// 5B: LD E,E			4 cycles
type op5b struct {
	SingleStepOp
}

func (op *op5b) Execute(c *CPU) (done bool) {
	c.E = c.E
	return true
}

// 5C: LD E,H			4 cycles
type op5c struct {
	SingleStepOp
}

func (op *op5c) Execute(c *CPU) (done bool) {
	c.E = c.H
	return true
}

// 5D: LD E,L			4 cycles
type op5d struct {
	SingleStepOp
}

func (op *op5d) Execute(c *CPU) (done bool) {
	c.E = c.L
	return true
}

// 5E: LD E,(HL)			8 cycles
type op5e struct {
	MultiStepsOp
}

func (op *op5e) Tick() (done bool) {
	op.cpu.E = op.cpu.MMU.Read(uint(op.cpu.HL()))
	return true
}

// 5F: LD E,A			4 cycles
type op5f struct {
	SingleStepOp
}

func (op *op5f) Execute(c *CPU) (done bool) {
	c.E = c.A
	return true
}

// 60: LD H,B			4 cycles
type op60 struct {
	SingleStepOp
}

func (op *op60) Execute(c *CPU) (done bool) {
	c.H = c.B
	return true
}

// 61: LD H,C			4 cycles
type op61 struct {
	SingleStepOp
}

func (op *op61) Execute(c *CPU) (done bool) {
	c.H = c.C
	return true
}

// 62: LD H,D			4 cycles
type op62 struct {
	SingleStepOp
}

func (op *op62) Execute(c *CPU) (done bool) {
	c.H = c.D
	return true
}

// 63: LD H,E			4 cycles
type op63 struct {
	SingleStepOp
}

func (op *op63) Execute(c *CPU) (done bool) {
	c.H = c.E
	return true
}

// 64: LD H,H			4 cycles
type op64 struct {
	SingleStepOp
}

func (op *op64) Execute(c *CPU) (done bool) {
	c.H = c.H
	return true
}

// 65: LD H,L			4 cycles
type op65 struct {
	SingleStepOp
}

func (op *op65) Execute(c *CPU) (done bool) {
	c.H = c.L
	return true
}

// 66: LD H,(HL)			8 cycles
type op66 struct {
	MultiStepsOp
}

func (op *op66) Tick() (done bool) {
	op.cpu.H = op.cpu.MMU.Read(uint(op.cpu.HL()))
	return true
}

// 67: LD H,A			4 cycles
type op67 struct {
	SingleStepOp
}

func (op *op67) Execute(c *CPU) (done bool) {
	c.H = c.A
	return true
}

// 68: LD L,B			4 cycles
type op68 struct {
	SingleStepOp
}

func (op *op68) Execute(c *CPU) (done bool) {
	c.L = c.B
	return true
}

// 69: LD L,C			4 cycles
type op69 struct {
	SingleStepOp
}

func (op *op69) Execute(c *CPU) (done bool) {
	c.L = c.C
	return true
}

// 6A: LD L,D			4 cycles
type op6a struct {
	SingleStepOp
}

func (op *op6a) Execute(c *CPU) (done bool) {
	c.L = c.D
	return true
}

// 6B: LD L,E			4 cycles
type op6b struct {
	SingleStepOp
}

func (op *op6b) Execute(c *CPU) (done bool) {
	c.L = c.E
	return true
}

// 6C: LD L,H			4 cycles
type op6c struct {
	SingleStepOp
}

func (op *op6c) Execute(c *CPU) (done bool) {
	c.L = c.H
	return true
}

// 6D: LD L,L			4 cycles
type op6d struct {
	SingleStepOp
}

func (op *op6d) Execute(c *CPU) (done bool) {
	c.L = c.L
	return true
}

// 6E: LD L,(HL)			8 cycles
type op6e struct {
	MultiStepsOp
}

func (op *op6e) Tick() (done bool) {
	op.cpu.L = op.cpu.MMU.Read(uint(op.cpu.HL()))
	return true
}

// 6F: LD L,A			4 cycles
type op6f struct {
	SingleStepOp
}

func (op *op6f) Execute(c *CPU) (done bool) {
	c.L = c.A
	return true
}

// 70: LD (HL),B		8 cycles
type op70 struct {
	MultiStepsOp
}

func (op *op70) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.B)
	return true
}

// 71: LD (HL),C		8 cycles
type op71 struct {
	MultiStepsOp
}

func (op *op71) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.C)
	return true
}

// 72: LD (HL),D		8 cycles
type op72 struct {
	MultiStepsOp
}

func (op *op72) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.D)
	return true
}

// 73: LD (HL),E		8 cycles
type op73 struct {
	MultiStepsOp
}

func (op *op73) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.E)
	return true
}

// 74: LD (HL),H		8 cycles
type op74 struct {
	MultiStepsOp
}

func (op *op74) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.H)
	return true
}

// 75: LD (HL),L		8 cycles
type op75 struct {
	MultiStepsOp
}

func (op *op75) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.L)
	return true
}

// 77: LD (HL),A		8 cycles
type op77 struct {
	MultiStepsOp
}

func (op *op77) Tick() (done bool) {
	op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.A)
	return true
}

// 78: LD A,B			4 cycles
type op78 struct {
	SingleStepOp
}

func (op *op78) Execute(c *CPU) (done bool) {
	c.A = c.B
	return true
}

// 79: LD A,C			4 cycles
type op79 struct {
	SingleStepOp
}

func (op *op79) Execute(c *CPU) (done bool) {
	c.A = c.C
	return true
}

// 7A: LD A,D			4 cycles
type op7a struct {
	SingleStepOp
}

func (op *op7a) Execute(c *CPU) (done bool) {
	c.A = c.D
	return true
}

// 7B: LD A,E			4 cycles
type op7b struct {
	SingleStepOp
}

func (op *op7b) Execute(c *CPU) (done bool) {
	c.A = c.E
	return true
}

// 7C: LD A,H			4 cycles
type op7c struct {
	SingleStepOp
}

func (op *op7c) Execute(c *CPU) (done bool) {
	c.A = c.H
	return true
}

// 7D: LD A,L			4 cycles
type op7d struct {
	SingleStepOp
}

func (op *op7d) Execute(c *CPU) (done bool) {
	c.A = c.L
	return true
}

// 7E: LD A,(HL)			8 cycles
type op7e struct {
	MultiStepsOp
}

func (op *op7e) Tick() (done bool) {
	op.cpu.A = op.cpu.MMU.Read(uint(op.cpu.HL()))
	return true
}

// 7F: LD A,A			4 cycles
type op7f struct {
	SingleStepOp
}

func (op *op7f) Execute(c *CPU) (done bool) {
	c.A = c.A
	return true
}

// 80: ADD A,B		4 cycles
type op80 struct {
	SingleStepOp
}

func (op *op80) Execute(c *CPU) (done bool) {
	// Flags: z 0 h c
	c.F = 0
	if c.A&0xf+c.B&0xf > 0xf {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.B)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 81: ADD A,C		4 cycles
type op81 struct {
	SingleStepOp
}

func (op *op81) Execute(c *CPU) (done bool) {
	// Flags: z 0 h c
	c.F = 0
	if c.A&0xf+c.C&0xf > 0xf {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.C)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 82: ADD A,D		4 cycles
type op82 struct {
	SingleStepOp
}

func (op *op82) Execute(c *CPU) (done bool) {
	// Flags: z 0 h c
	c.F = 0
	if c.A&0xf+c.D&0xf > 0xf {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.D)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 83: ADD A,E		4 cycles
type op83 struct {
	SingleStepOp
}

func (op *op83) Execute(c *CPU) (done bool) {
	// Flags: z 0 h c
	c.F = 0
	if c.A&0xf+c.E&0xf > 0xf {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.E)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 84: ADD A,H		4 cycles
type op84 struct {
	SingleStepOp
}

func (op *op84) Execute(c *CPU) (done bool) {
	// Flags: z 0 h c
	c.F = 0
	if c.A&0xf+c.H&0xf > 0xf {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.H)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 85: ADD A,L		4 cycles
type op85 struct {
	SingleStepOp
}

func (op *op85) Execute(c *CPU) (done bool) {
	// Flags: z 0 h c
	c.F = 0
	if c.A&0xf+c.L&0xf > 0xf {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.L)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 86: ADD A,(HL)		8 cycles
type op86 struct {
	MultiStepsOp
}

func (op *op86) Tick() (done bool) {
	// Flags: z 0 h c
	op.cpu.F = 0
	value := op.cpu.MMU.Read(uint(op.cpu.HL()))
	if op.cpu.A&0xf+value&0xf > 0xf {
		op.cpu.F |= FlagH
	}
	result := uint(op.cpu.A) + uint(value)
	if result > 0xff {
		op.cpu.F |= FlagC
	}
	op.cpu.A = uint8(result & 0xff)
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
    return true
}


// 87: ADD A,A		4 cycles
type op87 struct {
	SingleStepOp
}

func (op *op87) Execute(c *CPU) (done bool) {
	// Flags: z 0 h c
	c.F = 0
	if c.A&0xf+c.A&0xf > 0xf {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.A)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 90: SUB B		4 cycles
type op90 struct {
	SingleStepOp
}

func (op *op90) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.B&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.B > c.A {
		c.F |= FlagC
	}
	result := c.A - c.B
	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result
    return true
}


// 91: SUB C		4 cycles
type op91 struct {
	SingleStepOp
}

func (op *op91) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.C&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.C > c.A {
		c.F |= FlagC
	}
	result := c.A - c.C
	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result
    return true
}


// 92: SUB D		4 cycles
type op92 struct {
	SingleStepOp
}

func (op *op92) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.D&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.D > c.A {
		c.F |= FlagC
	}
	result := c.A - c.D
	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result
    return true
}


// 93: SUB E		4 cycles
type op93 struct {
	SingleStepOp
}

func (op *op93) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.E&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.E > c.A {
		c.F |= FlagC
	}
	result := c.A - c.E
	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result
    return true
}


// 94: SUB H		4 cycles
type op94 struct {
	SingleStepOp
}

func (op *op94) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.H&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.H > c.A {
		c.F |= FlagC
	}
	result := c.A - c.H
	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result
    return true
}


// 95: SUB L		4 cycles
type op95 struct {
	SingleStepOp
}

func (op *op95) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.L&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.L > c.A {
		c.F |= FlagC
	}
	result := c.A - c.L
	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result
    return true
}


// 96: SUB (HL)		8 cycles
type op96 struct {
	MultiStepsOp
}

func (op *op96) Tick() (done bool) {
	// Flags: z 1 h c
	value := op.cpu.MMU.Read(uint(op.cpu.HL()))
	op.cpu.F = FlagN
	if value&0xf > op.cpu.A&0xf {
		op.cpu.F |= FlagH
	}
	if value > op.cpu.A {
		op.cpu.F |= FlagC
	}
	result := op.cpu.A - value
	if result == 0 {
		op.cpu.F |= FlagZ
	}
	op.cpu.A = result
    return true
}


// 97: SUB A		4 cycles
type op97 struct {
	SingleStepOp
}

func (op *op97) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.A&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.A > c.A {
		c.F |= FlagC
	}
	result := c.A - c.A
	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result
    return true
}


// A0: AND B			4 cycles
type opA0 struct {
	SingleStepOp
}

func (op *opA0) Execute(c *CPU) (done bool) {
	c.A &= c.B
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A1: AND C			4 cycles
type opA1 struct {
	SingleStepOp
}

func (op *opA1) Execute(c *CPU) (done bool) {
	c.A &= c.C
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A2: AND D			4 cycles
type opA2 struct {
	SingleStepOp
}

func (op *opA2) Execute(c *CPU) (done bool) {
	c.A &= c.D
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A3: AND E			4 cycles
type opA3 struct {
	SingleStepOp
}

func (op *opA3) Execute(c *CPU) (done bool) {
	c.A &= c.E
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A4: AND H			4 cycles
type opA4 struct {
	SingleStepOp
}

func (op *opA4) Execute(c *CPU) (done bool) {
	c.A &= c.H
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A5: AND L			4 cycles
type opA5 struct {
	SingleStepOp
}

func (op *opA5) Execute(c *CPU) (done bool) {
	c.A &= c.L
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A6: AND (HL)			8 cycles
type opA6 struct {
	MultiStepsOp
}

func (op *opA6) Tick() (done bool) {
	op.cpu.A &= op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 0 0
	if op.cpu.A == 0 {
		op.cpu.F = FlagZ
	} else {
		op.cpu.F = 0
	}
	return true
}

// A7: AND A			4 cycles
type opA7 struct {
	SingleStepOp
}

func (op *opA7) Execute(c *CPU) (done bool) {
	c.A &= c.A
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A8: XOR B			4 cycles
type opA8 struct {
	SingleStepOp
}

func (op *opA8) Execute(c *CPU) (done bool) {
	c.A ^= c.B
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// A9: XOR C			4 cycles
type opA9 struct {
	SingleStepOp
}

func (op *opA9) Execute(c *CPU) (done bool) {
	c.A ^= c.C
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// AA: XOR D			4 cycles
type opAa struct {
	SingleStepOp
}

func (op *opAa) Execute(c *CPU) (done bool) {
	c.A ^= c.D
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// AB: XOR E			4 cycles
type opAb struct {
	SingleStepOp
}

func (op *opAb) Execute(c *CPU) (done bool) {
	c.A ^= c.E
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// AC: XOR H			4 cycles
type opAc struct {
	SingleStepOp
}

func (op *opAc) Execute(c *CPU) (done bool) {
	c.A ^= c.H
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// AD: XOR L			4 cycles
type opAd struct {
	SingleStepOp
}

func (op *opAd) Execute(c *CPU) (done bool) {
	c.A ^= c.L
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// AE: XOR (HL)			8 cycles
type opAe struct {
	MultiStepsOp
}

func (op *opAe) Tick() (done bool) {
	op.cpu.A ^= op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 0 0
	if op.cpu.A == 0 {
		op.cpu.F = FlagZ
	} else {
		op.cpu.F = 0
	}
	return true
}

// AF: XOR A			4 cycles
type opAf struct {
	SingleStepOp
}

func (op *opAf) Execute(c *CPU) (done bool) {
	c.A ^= c.A
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// B0: OR B			4 cycles
type opB0 struct {
	SingleStepOp
}

func (op *opB0) Execute(c *CPU) (done bool) {
	c.A |= c.B
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// B1: OR C			4 cycles
type opB1 struct {
	SingleStepOp
}

func (op *opB1) Execute(c *CPU) (done bool) {
	c.A |= c.C
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// B2: OR D			4 cycles
type opB2 struct {
	SingleStepOp
}

func (op *opB2) Execute(c *CPU) (done bool) {
	c.A |= c.D
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// B3: OR E			4 cycles
type opB3 struct {
	SingleStepOp
}

func (op *opB3) Execute(c *CPU) (done bool) {
	c.A |= c.E
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// B4: OR H			4 cycles
type opB4 struct {
	SingleStepOp
}

func (op *opB4) Execute(c *CPU) (done bool) {
	c.A |= c.H
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// B5: OR L			4 cycles
type opB5 struct {
	SingleStepOp
}

func (op *opB5) Execute(c *CPU) (done bool) {
	c.A |= c.L
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// B6: OR (HL)			8 cycles
type opB6 struct {
	MultiStepsOp
}

func (op *opB6) Tick() (done bool) {
	op.cpu.A |= op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 0 0
	if op.cpu.A == 0 {
		op.cpu.F = FlagZ
	} else {
		op.cpu.F = 0
	}
	return true
}

// B7: OR A			4 cycles
type opB7 struct {
	SingleStepOp
}

func (op *opB7) Execute(c *CPU) (done bool) {
	c.A |= c.A
	// Flags z 0 0 0
	if c.A == 0 {
		c.F = FlagZ
	} else {
		c.F = 0
	}
	return true
}

// B8: CP B		4 cycles
type opB8 struct {
	SingleStepOp
}

func (op *opB8) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.B&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.B > c.A {
		c.F |= FlagC
	}
	result := c.A - c.B
	if result == 0 {
		c.F |= FlagZ
	}
    return true
}


// B9: CP C		4 cycles
type opB9 struct {
	SingleStepOp
}

func (op *opB9) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.C&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.C > c.A {
		c.F |= FlagC
	}
	result := c.A - c.C
	if result == 0 {
		c.F |= FlagZ
	}
    return true
}


// BA: CP D		4 cycles
type opBa struct {
	SingleStepOp
}

func (op *opBa) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.D&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.D > c.A {
		c.F |= FlagC
	}
	result := c.A - c.D
	if result == 0 {
		c.F |= FlagZ
	}
    return true
}


// BB: CP E		4 cycles
type opBb struct {
	SingleStepOp
}

func (op *opBb) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.E&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.E > c.A {
		c.F |= FlagC
	}
	result := c.A - c.E
	if result == 0 {
		c.F |= FlagZ
	}
    return true
}


// BC: CP H		4 cycles
type opBc struct {
	SingleStepOp
}

func (op *opBc) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.H&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.H > c.A {
		c.F |= FlagC
	}
	result := c.A - c.H
	if result == 0 {
		c.F |= FlagZ
	}
    return true
}


// BD: CP L		4 cycles
type opBd struct {
	SingleStepOp
}

func (op *opBd) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.L&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.L > c.A {
		c.F |= FlagC
	}
	result := c.A - c.L
	if result == 0 {
		c.F |= FlagZ
	}
    return true
}


// BE: CP (HL)		8 cycles
type opBe struct {
	MultiStepsOp
}

func (op *opBe) Tick() (done bool) {
	// Flags: z 1 h c
	value := op.cpu.MMU.Read(uint(op.cpu.HL()))
	op.cpu.F = FlagN
	if value&0xf > op.cpu.A&0xf {
		op.cpu.F |= FlagH
	}
	if value > op.cpu.A {
		op.cpu.F |= FlagC
	}
	result := op.cpu.A - value
	if result == 0 {
		op.cpu.F |= FlagZ
	}
    return true
}


// BF: CP A		4 cycles
type opBf struct {
	SingleStepOp
}

func (op *opBf) Execute(c *CPU) (done bool) {
	// Flags: z 1 h c
	c.F = FlagN
	if c.A&0xf > c.A&0xf {
		c.F |= FlagH
	}
	if c.A > c.A {
		c.F |= FlagC
	}
	result := c.A - c.A
	if result == 0 {
		c.F |= FlagZ
	}
    return true
}


// C1: POP BC		12 cycles
type opC1 struct {
	MultiStepsOp
}

func (op *opC1) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.C = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		op.step++
	case 1:
		op.cpu.B = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		done = true
	}
	return
}

// C5: PUSH BC		16 cycles
type opC5 struct {
	MultiStepsOp
}

func (op *opC5) Tick() (done bool) {
	switch op.step {
	case 0:
		// Waiting cycle according to [GEKKIO]. To align with memory timing?
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.B))
		op.step++
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.C))
		done = true
	}
	return
}

// C9: RET		16 cycles
type opC9 struct {
	MultiStepsOp
}

func (op *opC9) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.PC = uint16(op.cpu.MMU.Read(uint(op.cpu.SP)))
		op.cpu.SP++
		op.step++
	case 1:
		op.cpu.PC |= uint16(op.cpu.MMU.Read(uint(op.cpu.SP))) << 8
		op.cpu.SP++
		op.step++
	case 2:
		// [GEKKIO] mentions an internal delay.
		done = true
	}
	return
}

// CD: CALL a16		24 cycles
type opCd struct {
	MultiStepsOp
}

func (op *opCd) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		op.step++
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 3:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0xff))
		op.step++
	case 4:
		op.cpu.PC = op.cpu.temp16
		done = true
	}
	return
}

// D1: POP DE		12 cycles
type opD1 struct {
	MultiStepsOp
}

func (op *opD1) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.E = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		op.step++
	case 1:
		op.cpu.D = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		done = true
	}
	return
}

// D5: PUSH DE		16 cycles
type opD5 struct {
	MultiStepsOp
}

func (op *opD5) Tick() (done bool) {
	switch op.step {
	case 0:
		// Waiting cycle according to [GEKKIO]. To align with memory timing?
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.D))
		op.step++
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.E))
		done = true
	}
	return
}

// D9: RETI		16 cycles
type opD9 struct {
	MultiStepsOp
}

func (op *opD9) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.PC = uint16(op.cpu.MMU.Read(uint(op.cpu.SP)))
		op.cpu.SP++
		op.step++
	case 1:
		op.cpu.PC |= uint16(op.cpu.MMU.Read(uint(op.cpu.SP))) << 8
		op.cpu.SP++
		op.step++
	case 2:
		op.cpu.IME = true
		// [GEKKIO] mentions an internal delay.
		done = true
	}
	return
}

// E0: LD (FF00+a8),A		12 cycles
type opE0 struct {
	MultiStepsOp
}

func (op *opE0) Tick() (done bool) {
	op.cpu.MMU.Write(uint(0xff00+uint16(op.cpu.NextByte())), op.cpu.A)
	return true
}

// E1: POP HL		12 cycles
type opE1 struct {
	MultiStepsOp
}

func (op *opE1) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.L = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		op.step++
	case 1:
		op.cpu.H = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		done = true
	}
	return
}

// E2: LD (FF00+C),A		8 cycles
type opE2 struct {
	MultiStepsOp
}

func (op *opE2) Tick() (done bool) {
	op.cpu.MMU.Write(uint(0xff00+uint16(op.cpu.C)), op.cpu.A)
	return true
}

// E5: PUSH HL		16 cycles
type opE5 struct {
	MultiStepsOp
}

func (op *opE5) Tick() (done bool) {
	switch op.step {
	case 0:
		// Waiting cycle according to [GEKKIO]. To align with memory timing?
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.H))
		op.step++
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.L))
		done = true
	}
	return
}

// EA: LD (a16),A		16 cycles
type opEa struct {
	MultiStepsOp
}

func (op *opEa) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
	case 2:
		op.cpu.MMU.Write(uint(op.cpu.temp16), uint8(op.cpu.A))
		done = true
	}
	op.step++
	return
}

// F0: LD A,(FF00+a8)		12 cycles
type opF0 struct {
	MultiStepsOp
}

func (op *opF0) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.NextByte()
		op.step++
	case 1:
		op.cpu.A = op.cpu.MMU.Read(uint(0xff00+uint16(op.cpu.temp8)))
		done = true
	}
	return
}

// F1: POP AF		12 cycles
type opF1 struct {
	MultiStepsOp
}

func (op *opF1) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.F = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		op.step++
	case 1:
		op.cpu.A = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		done = true
	}
	return
}

// F5: PUSH AF		16 cycles
type opF5 struct {
	MultiStepsOp
}

func (op *opF5) Tick() (done bool) {
	switch op.step {
	case 0:
		// Waiting cycle according to [GEKKIO]. To align with memory timing?
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.A))
		op.step++
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.F))
		done = true
	}
	return
}

// FE: CP d8		8 cycles
type opFe struct {
	MultiStepsOp
}

func (op *opFe) Tick() (done bool) {
	// Flags: z 1 h c
	value := op.cpu.NextByte()
	op.cpu.F = FlagN
	if value&0xf > op.cpu.A&0xf {
		op.cpu.F |= FlagH
	}
	if value > op.cpu.A {
		op.cpu.F |= FlagC
	}
	result := op.cpu.A - value
	if result == 0 {
		op.cpu.F |= FlagZ
	}
    return true
}


// CB 10: RL B				8 cycles
type opCb10 struct {
	SingleStepOp
}

func (op *opCb10) Execute(c *CPU) (done bool) {
	result := c.B << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.B&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.B = result

	return true
}

// CB 11: RL C				8 cycles
type opCb11 struct {
	SingleStepOp
}

func (op *opCb11) Execute(c *CPU) (done bool) {
	result := c.C << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.C&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.C = result

	return true
}

// CB 12: RL D				8 cycles
type opCb12 struct {
	SingleStepOp
}

func (op *opCb12) Execute(c *CPU) (done bool) {
	result := c.D << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.D&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.D = result

	return true
}

// CB 13: RL E				8 cycles
type opCb13 struct {
	SingleStepOp
}

func (op *opCb13) Execute(c *CPU) (done bool) {
	result := c.E << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.E&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.E = result

	return true
}

// CB 14: RL H				8 cycles
type opCb14 struct {
	SingleStepOp
}

func (op *opCb14) Execute(c *CPU) (done bool) {
	result := c.H << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.H&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.H = result

	return true
}

// CB 15: RL L				8 cycles
type opCb15 struct {
	SingleStepOp
}

func (op *opCb15) Execute(c *CPU) (done bool) {
	result := c.L << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.L&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.L = result

	return true
}

// CB 17: RL A				8 cycles
type opCb17 struct {
	SingleStepOp
}

func (op *opCb17) Execute(c *CPU) (done bool) {
	result := c.A << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.A&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.A = result

	return true
}

// CB 40: BIT 0,B			8 cycles
type opCb40 struct {
	SingleStepOp
}

func (op *opCb40) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.B&(1<<0) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 41: BIT 0,C			8 cycles
type opCb41 struct {
	SingleStepOp
}

func (op *opCb41) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.C&(1<<0) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 42: BIT 0,D			8 cycles
type opCb42 struct {
	SingleStepOp
}

func (op *opCb42) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.D&(1<<0) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 43: BIT 0,E			8 cycles
type opCb43 struct {
	SingleStepOp
}

func (op *opCb43) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.E&(1<<0) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 44: BIT 0,H			8 cycles
type opCb44 struct {
	SingleStepOp
}

func (op *opCb44) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.H&(1<<0) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 45: BIT 0,L			8 cycles
type opCb45 struct {
	SingleStepOp
}

func (op *opCb45) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.L&(1<<0) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 46: BIT 0,(HL)			16 cycles
type opCb46 struct {
	MultiStepsOp
}

func (op *opCb46) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 1 -
		if op.cpu.temp8&(1<<0) == 0 {
			op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
		} else {
			op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
		}
		done = true
	}
	return
}

// CB 47: BIT 0,A			8 cycles
type opCb47 struct {
	SingleStepOp
}

func (op *opCb47) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.A&(1<<0) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 48: BIT 1,B			8 cycles
type opCb48 struct {
	SingleStepOp
}

func (op *opCb48) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.B&(1<<1) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 49: BIT 1,C			8 cycles
type opCb49 struct {
	SingleStepOp
}

func (op *opCb49) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.C&(1<<1) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 4A: BIT 1,D			8 cycles
type opCb4a struct {
	SingleStepOp
}

func (op *opCb4a) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.D&(1<<1) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 4B: BIT 1,E			8 cycles
type opCb4b struct {
	SingleStepOp
}

func (op *opCb4b) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.E&(1<<1) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 4C: BIT 1,H			8 cycles
type opCb4c struct {
	SingleStepOp
}

func (op *opCb4c) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.H&(1<<1) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 4D: BIT 1,L			8 cycles
type opCb4d struct {
	SingleStepOp
}

func (op *opCb4d) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.L&(1<<1) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 4E: BIT 1,(HL)			16 cycles
type opCb4e struct {
	MultiStepsOp
}

func (op *opCb4e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 1 -
		if op.cpu.temp8&(1<<1) == 0 {
			op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
		} else {
			op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
		}
		done = true
	}
	return
}

// CB 4F: BIT 1,A			8 cycles
type opCb4f struct {
	SingleStepOp
}

func (op *opCb4f) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.A&(1<<1) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 50: BIT 2,B			8 cycles
type opCb50 struct {
	SingleStepOp
}

func (op *opCb50) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.B&(1<<2) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 51: BIT 2,C			8 cycles
type opCb51 struct {
	SingleStepOp
}

func (op *opCb51) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.C&(1<<2) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 52: BIT 2,D			8 cycles
type opCb52 struct {
	SingleStepOp
}

func (op *opCb52) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.D&(1<<2) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 53: BIT 2,E			8 cycles
type opCb53 struct {
	SingleStepOp
}

func (op *opCb53) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.E&(1<<2) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 54: BIT 2,H			8 cycles
type opCb54 struct {
	SingleStepOp
}

func (op *opCb54) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.H&(1<<2) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 55: BIT 2,L			8 cycles
type opCb55 struct {
	SingleStepOp
}

func (op *opCb55) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.L&(1<<2) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 56: BIT 2,(HL)			16 cycles
type opCb56 struct {
	MultiStepsOp
}

func (op *opCb56) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 1 -
		if op.cpu.temp8&(1<<2) == 0 {
			op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
		} else {
			op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
		}
		done = true
	}
	return
}

// CB 57: BIT 2,A			8 cycles
type opCb57 struct {
	SingleStepOp
}

func (op *opCb57) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.A&(1<<2) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 58: BIT 3,B			8 cycles
type opCb58 struct {
	SingleStepOp
}

func (op *opCb58) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.B&(1<<3) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 59: BIT 3,C			8 cycles
type opCb59 struct {
	SingleStepOp
}

func (op *opCb59) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.C&(1<<3) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 5A: BIT 3,D			8 cycles
type opCb5a struct {
	SingleStepOp
}

func (op *opCb5a) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.D&(1<<3) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 5B: BIT 3,E			8 cycles
type opCb5b struct {
	SingleStepOp
}

func (op *opCb5b) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.E&(1<<3) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 5C: BIT 3,H			8 cycles
type opCb5c struct {
	SingleStepOp
}

func (op *opCb5c) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.H&(1<<3) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 5D: BIT 3,L			8 cycles
type opCb5d struct {
	SingleStepOp
}

func (op *opCb5d) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.L&(1<<3) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 5E: BIT 3,(HL)			16 cycles
type opCb5e struct {
	MultiStepsOp
}

func (op *opCb5e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 1 -
		if op.cpu.temp8&(1<<3) == 0 {
			op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
		} else {
			op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
		}
		done = true
	}
	return
}

// CB 5F: BIT 3,A			8 cycles
type opCb5f struct {
	SingleStepOp
}

func (op *opCb5f) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.A&(1<<3) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 60: BIT 4,B			8 cycles
type opCb60 struct {
	SingleStepOp
}

func (op *opCb60) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.B&(1<<4) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 61: BIT 4,C			8 cycles
type opCb61 struct {
	SingleStepOp
}

func (op *opCb61) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.C&(1<<4) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 62: BIT 4,D			8 cycles
type opCb62 struct {
	SingleStepOp
}

func (op *opCb62) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.D&(1<<4) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 63: BIT 4,E			8 cycles
type opCb63 struct {
	SingleStepOp
}

func (op *opCb63) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.E&(1<<4) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 64: BIT 4,H			8 cycles
type opCb64 struct {
	SingleStepOp
}

func (op *opCb64) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.H&(1<<4) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 65: BIT 4,L			8 cycles
type opCb65 struct {
	SingleStepOp
}

func (op *opCb65) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.L&(1<<4) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 66: BIT 4,(HL)			16 cycles
type opCb66 struct {
	MultiStepsOp
}

func (op *opCb66) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 1 -
		if op.cpu.temp8&(1<<4) == 0 {
			op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
		} else {
			op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
		}
		done = true
	}
	return
}

// CB 67: BIT 4,A			8 cycles
type opCb67 struct {
	SingleStepOp
}

func (op *opCb67) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.A&(1<<4) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 68: BIT 5,B			8 cycles
type opCb68 struct {
	SingleStepOp
}

func (op *opCb68) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.B&(1<<5) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 69: BIT 5,C			8 cycles
type opCb69 struct {
	SingleStepOp
}

func (op *opCb69) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.C&(1<<5) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 6A: BIT 5,D			8 cycles
type opCb6a struct {
	SingleStepOp
}

func (op *opCb6a) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.D&(1<<5) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 6B: BIT 5,E			8 cycles
type opCb6b struct {
	SingleStepOp
}

func (op *opCb6b) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.E&(1<<5) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 6C: BIT 5,H			8 cycles
type opCb6c struct {
	SingleStepOp
}

func (op *opCb6c) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.H&(1<<5) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 6D: BIT 5,L			8 cycles
type opCb6d struct {
	SingleStepOp
}

func (op *opCb6d) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.L&(1<<5) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 6E: BIT 5,(HL)			16 cycles
type opCb6e struct {
	MultiStepsOp
}

func (op *opCb6e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 1 -
		if op.cpu.temp8&(1<<5) == 0 {
			op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
		} else {
			op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
		}
		done = true
	}
	return
}

// CB 6F: BIT 5,A			8 cycles
type opCb6f struct {
	SingleStepOp
}

func (op *opCb6f) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.A&(1<<5) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 70: BIT 6,B			8 cycles
type opCb70 struct {
	SingleStepOp
}

func (op *opCb70) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.B&(1<<6) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 71: BIT 6,C			8 cycles
type opCb71 struct {
	SingleStepOp
}

func (op *opCb71) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.C&(1<<6) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 72: BIT 6,D			8 cycles
type opCb72 struct {
	SingleStepOp
}

func (op *opCb72) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.D&(1<<6) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 73: BIT 6,E			8 cycles
type opCb73 struct {
	SingleStepOp
}

func (op *opCb73) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.E&(1<<6) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 74: BIT 6,H			8 cycles
type opCb74 struct {
	SingleStepOp
}

func (op *opCb74) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.H&(1<<6) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 75: BIT 6,L			8 cycles
type opCb75 struct {
	SingleStepOp
}

func (op *opCb75) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.L&(1<<6) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 76: BIT 6,(HL)			16 cycles
type opCb76 struct {
	MultiStepsOp
}

func (op *opCb76) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 1 -
		if op.cpu.temp8&(1<<6) == 0 {
			op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
		} else {
			op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
		}
		done = true
	}
	return
}

// CB 77: BIT 6,A			8 cycles
type opCb77 struct {
	SingleStepOp
}

func (op *opCb77) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.A&(1<<6) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 78: BIT 7,B			8 cycles
type opCb78 struct {
	SingleStepOp
}

func (op *opCb78) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.B&(1<<7) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 79: BIT 7,C			8 cycles
type opCb79 struct {
	SingleStepOp
}

func (op *opCb79) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.C&(1<<7) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 7A: BIT 7,D			8 cycles
type opCb7a struct {
	SingleStepOp
}

func (op *opCb7a) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.D&(1<<7) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 7B: BIT 7,E			8 cycles
type opCb7b struct {
	SingleStepOp
}

func (op *opCb7b) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.E&(1<<7) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 7C: BIT 7,H			8 cycles
type opCb7c struct {
	SingleStepOp
}

func (op *opCb7c) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.H&(1<<7) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 7D: BIT 7,L			8 cycles
type opCb7d struct {
	SingleStepOp
}

func (op *opCb7d) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.L&(1<<7) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

// CB 7E: BIT 7,(HL)			16 cycles
type opCb7e struct {
	MultiStepsOp
}

func (op *opCb7e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 1 -
		if op.cpu.temp8&(1<<7) == 0 {
			op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
		} else {
			op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
		}
		done = true
	}
	return
}

// CB 7F: BIT 7,A			8 cycles
type opCb7f struct {
	SingleStepOp
}

func (op *opCb7f) Execute(c *CPU) (done bool) {
	// Flags z 0 1 -
	if c.A&(1<<7) == 0 {
		c.F = (c.F & ^FlagN) | FlagZ | FlagH
	} else {
		c.F = (c.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
}

