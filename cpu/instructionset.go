// Auto-generated on 2019-04-02 22:24:14.639312645 +0200 CEST m=+0.279342208. See instructions.go
package cpu

import "go.tigris.fr/gameboy/cpu/states"

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
	0x0f: &op0f{},
	0x10: &op10{},
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
	0x1f: &op1f{},
	0x20: &op20{},
	0x21: &op21{},
	0x22: &op22{},
	0x23: &op23{},
	0x24: &op24{},
	0x25: &op25{},
	0x26: &op26{},
	0x27: &op27{},
	0x28: &op28{},
	0x29: &op29{},
	0x2a: &op2a{},
	0x2b: &op2b{},
	0x2c: &op2c{},
	0x2d: &op2d{},
	0x2e: &op2e{},
	0x2f: &op2f{},
	0x30: &op30{},
	0x31: &op31{},
	0x32: &op32{},
	0x33: &op33{},
	0x34: &op34{},
	0x35: &op35{},
	0x36: &op36{},
	0x37: &op37{},
	0x38: &op38{},
	0x39: &op39{},
	0x3a: &op3a{},
	0x3b: &op3b{},
	0x3c: &op3c{},
	0x3d: &op3d{},
	0x3e: &op3e{},
	0x3f: &op3f{},
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
	0x76: &op76{},
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
	0x88: &op88{},
	0x89: &op89{},
	0x8a: &op8a{},
	0x8b: &op8b{},
	0x8c: &op8c{},
	0x8d: &op8d{},
	0x8e: &op8e{},
	0x8f: &op8f{},
	0x90: &op90{},
	0x91: &op91{},
	0x92: &op92{},
	0x93: &op93{},
	0x94: &op94{},
	0x95: &op95{},
	0x96: &op96{},
	0x97: &op97{},
	0x98: &op98{},
	0x99: &op99{},
	0x9a: &op9a{},
	0x9b: &op9b{},
	0x9c: &op9c{},
	0x9d: &op9d{},
	0x9e: &op9e{},
	0x9f: &op9f{},
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
	0xc0: &opC0{},
	0xc1: &opC1{},
	0xc2: &opC2{},
	0xc3: &opC3{},
	0xc4: &opC4{},
	0xc5: &opC5{},
	0xc6: &opC6{},
	0xc7: &opC7{},
	0xc8: &opC8{},
	0xc9: &opC9{},
	0xca: &opCa{},
	0xcc: &opCc{},
	0xcd: &opCd{},
	0xce: &opCe{},
	0xcf: &opCf{},
	0xd0: &opD0{},
	0xd1: &opD1{},
	0xd2: &opD2{},
	0xd4: &opD4{},
	0xd5: &opD5{},
	0xd6: &opD6{},
	0xd7: &opD7{},
	0xd8: &opD8{},
	0xd9: &opD9{},
	0xda: &opDa{},
	0xdc: &opDc{},
	0xde: &opDe{},
	0xdf: &opDf{},
	0xe0: &opE0{},
	0xe1: &opE1{},
	0xe2: &opE2{},
	0xe5: &opE5{},
	0xe6: &opE6{},
	0xe7: &opE7{},
	0xe8: &opE8{},
	0xe9: &opE9{},
	0xea: &opEa{},
	0xee: &opEe{},
	0xef: &opEf{},
	0xf0: &opF0{},
	0xf1: &opF1{},
	0xf2: &opF2{},
	0xf3: &opF3{},
	0xf5: &opF5{},
	0xf6: &opF6{},
	0xf7: &opF7{},
	0xf8: &opF8{},
	0xf9: &opF9{},
	0xfa: &opFa{},
	0xfb: &opFb{},
	0xfe: &opFe{},
	0xff: &opFf{},
}

// LR35902ExtendedInstructionSet is the array of extension opcodes for the DMG CPU.
var LR35902ExtendedInstructionSet = [...]Instruction{
	0x00: &opCb00{},
	0x01: &opCb01{},
	0x02: &opCb02{},
	0x03: &opCb03{},
	0x04: &opCb04{},
	0x05: &opCb05{},
	0x06: &opCb06{},
	0x07: &opCb07{},
	0x08: &opCb08{},
	0x09: &opCb09{},
	0x0a: &opCb0a{},
	0x0b: &opCb0b{},
	0x0c: &opCb0c{},
	0x0d: &opCb0d{},
	0x0e: &opCb0e{},
	0x0f: &opCb0f{},
	0x10: &opCb10{},
	0x11: &opCb11{},
	0x12: &opCb12{},
	0x13: &opCb13{},
	0x14: &opCb14{},
	0x15: &opCb15{},
	0x16: &opCb16{},
	0x17: &opCb17{},
	0x18: &opCb18{},
	0x19: &opCb19{},
	0x1a: &opCb1a{},
	0x1b: &opCb1b{},
	0x1c: &opCb1c{},
	0x1d: &opCb1d{},
	0x1e: &opCb1e{},
	0x1f: &opCb1f{},
	0x20: &opCb20{},
	0x21: &opCb21{},
	0x22: &opCb22{},
	0x23: &opCb23{},
	0x24: &opCb24{},
	0x25: &opCb25{},
	0x26: &opCb26{},
	0x27: &opCb27{},
	0x28: &opCb28{},
	0x29: &opCb29{},
	0x2a: &opCb2a{},
	0x2b: &opCb2b{},
	0x2c: &opCb2c{},
	0x2d: &opCb2d{},
	0x2e: &opCb2e{},
	0x2f: &opCb2f{},
	0x30: &opCb30{},
	0x31: &opCb31{},
	0x32: &opCb32{},
	0x33: &opCb33{},
	0x34: &opCb34{},
	0x35: &opCb35{},
	0x36: &opCb36{},
	0x37: &opCb37{},
	0x38: &opCb38{},
	0x39: &opCb39{},
	0x3a: &opCb3a{},
	0x3b: &opCb3b{},
	0x3c: &opCb3c{},
	0x3d: &opCb3d{},
	0x3e: &opCb3e{},
	0x3f: &opCb3f{},
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
	0x80: &opCb80{},
	0x81: &opCb81{},
	0x82: &opCb82{},
	0x83: &opCb83{},
	0x84: &opCb84{},
	0x85: &opCb85{},
	0x86: &opCb86{},
	0x87: &opCb87{},
	0x88: &opCb88{},
	0x89: &opCb89{},
	0x8a: &opCb8a{},
	0x8b: &opCb8b{},
	0x8c: &opCb8c{},
	0x8d: &opCb8d{},
	0x8e: &opCb8e{},
	0x8f: &opCb8f{},
	0x90: &opCb90{},
	0x91: &opCb91{},
	0x92: &opCb92{},
	0x93: &opCb93{},
	0x94: &opCb94{},
	0x95: &opCb95{},
	0x96: &opCb96{},
	0x97: &opCb97{},
	0x98: &opCb98{},
	0x99: &opCb99{},
	0x9a: &opCb9a{},
	0x9b: &opCb9b{},
	0x9c: &opCb9c{},
	0x9d: &opCb9d{},
	0x9e: &opCb9e{},
	0x9f: &opCb9f{},
	0xa0: &opCbA0{},
	0xa1: &opCbA1{},
	0xa2: &opCbA2{},
	0xa3: &opCbA3{},
	0xa4: &opCbA4{},
	0xa5: &opCbA5{},
	0xa6: &opCbA6{},
	0xa7: &opCbA7{},
	0xa8: &opCbA8{},
	0xa9: &opCbA9{},
	0xaa: &opCbAa{},
	0xab: &opCbAb{},
	0xac: &opCbAc{},
	0xad: &opCbAd{},
	0xae: &opCbAe{},
	0xaf: &opCbAf{},
	0xb0: &opCbB0{},
	0xb1: &opCbB1{},
	0xb2: &opCbB2{},
	0xb3: &opCbB3{},
	0xb4: &opCbB4{},
	0xb5: &opCbB5{},
	0xb6: &opCbB6{},
	0xb7: &opCbB7{},
	0xb8: &opCbB8{},
	0xb9: &opCbB9{},
	0xba: &opCbBa{},
	0xbb: &opCbBb{},
	0xbc: &opCbBc{},
	0xbd: &opCbBd{},
	0xbe: &opCbBe{},
	0xbf: &opCbBf{},
	0xc0: &opCbC0{},
	0xc1: &opCbC1{},
	0xc2: &opCbC2{},
	0xc3: &opCbC3{},
	0xc4: &opCbC4{},
	0xc5: &opCbC5{},
	0xc6: &opCbC6{},
	0xc7: &opCbC7{},
	0xc8: &opCbC8{},
	0xc9: &opCbC9{},
	0xca: &opCbCa{},
	0xcb: &opCbCb{},
	0xcc: &opCbCc{},
	0xcd: &opCbCd{},
	0xce: &opCbCe{},
	0xcf: &opCbCf{},
	0xd0: &opCbD0{},
	0xd1: &opCbD1{},
	0xd2: &opCbD2{},
	0xd3: &opCbD3{},
	0xd4: &opCbD4{},
	0xd5: &opCbD5{},
	0xd6: &opCbD6{},
	0xd7: &opCbD7{},
	0xd8: &opCbD8{},
	0xd9: &opCbD9{},
	0xda: &opCbDa{},
	0xdb: &opCbDb{},
	0xdc: &opCbDc{},
	0xdd: &opCbDd{},
	0xde: &opCbDe{},
	0xdf: &opCbDf{},
	0xe0: &opCbE0{},
	0xe1: &opCbE1{},
	0xe2: &opCbE2{},
	0xe3: &opCbE3{},
	0xe4: &opCbE4{},
	0xe5: &opCbE5{},
	0xe6: &opCbE6{},
	0xe7: &opCbE7{},
	0xe8: &opCbE8{},
	0xe9: &opCbE9{},
	0xea: &opCbEa{},
	0xeb: &opCbEb{},
	0xec: &opCbEc{},
	0xed: &opCbEd{},
	0xee: &opCbEe{},
	0xef: &opCbEf{},
	0xf0: &opCbF0{},
	0xf1: &opCbF1{},
	0xf2: &opCbF2{},
	0xf3: &opCbF3{},
	0xf4: &opCbF4{},
	0xf5: &opCbF5{},
	0xf6: &opCbF6{},
	0xf7: &opCbF7{},
	0xf8: &opCbF8{},
	0xf9: &opCbF9{},
	0xfa: &opCbFa{},
	0xfb: &opCbFb{},
	0xfc: &opCbFc{},
	0xfd: &opCbFd{},
	0xfe: &opCbFe{},
	0xff: &opCbFf{},
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
	c.F &= FlagC
	if c.B&0x0f == 0x0f {
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
	if c.B&0x0f == 0 {
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
	// Flags: 0 0 0 c
	// Flags z 0 0 c
	c.F = 0x00
	result := c.A << 1 & 0xff
	if c.A&0x80 != 0 {
		result |= 1
		c.F |= FlagC
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
		op.cpu.MMU.Write(uint(op.cpu.temp16+1), uint8(op.cpu.SP>>8))
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

// 0A: LD A,(BC)			8 cycles
type op0a struct {
	MultiStepsOp
}

func (op *op0a) Tick() (done bool) {
	op.cpu.A = op.cpu.MMU.Read(uint(op.cpu.BC()))
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
	c.F &= FlagC
	if c.C&0x0f == 0x0f {
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
	if c.C&0x0f == 0 {
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

// 0F: RRCA				4 cycles
type op0f struct {
	SingleStepOp
}

func (op *op0f) Execute(c *CPU) (done bool) {
	// Flags 0 0 0 c
	result := c.A >> 1
	// Flags z 0 0 c
	c.F = 0x00
	if c.A&1 > 0 {
		result |= (1 << 7)
		c.F |= FlagC
	}
	c.A = result

	return true
}

// 10: STOP 0				4 cycles
type op10 struct {
	SingleStepOp
}

func (op *op10) Execute(c *CPU) (done bool) {
	// Source indicates a 2-byte, 4-cycle instruction but this is unclear.
	c.PC++	// Ignore following zero
	c.state = states.Stopped
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
	c.F &= FlagC
	if c.D&0x0f == 0x0f {
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
	if c.D&0x0f == 0 {
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
	// Flags 0 0 0 c
	result := c.A << 1 & 0xff
	if c.F&FlagC > 0 {
		result |= 1
	}
	// Flags z 0 0 c
	c.F = 0x00
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
	c.F &= FlagC
	if c.E&0x0f == 0x0f {
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
	if c.E&0x0f == 0 {
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

// 1F: RRA				4 cycles
type op1f struct {
	SingleStepOp
}

func (op *op1f) Execute(c *CPU) (done bool) {
	// Flags 0 0 0 c
	result := c.A >> 1
	if c.F&FlagC > 0 {
		result |= (1<<7)
	}
	// Flags z 0 0 c
	c.F = 0x00
	if c.A&1 > 0 {
		c.F |= FlagC
	}
	c.A = result

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
	c.F &= FlagC
	if c.H&0x0f == 0x0f {
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
	if c.H&0x0f == 0 {
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

// 27: DAA				4 cycles
type op27 struct {
	SingleStepOp
}

func (op *op27) Execute(c *CPU) (done bool) {
	// Huge thanks to www.z80.info/z80syntx.htm and coffee-gb for that one.
	result := int(c.A)
	if c.F&FlagN != 0 {
		if c.F&FlagH != 0 {
			result = (result - 0x06) & 0xff
		}
		if c.F&FlagC != 0 {
			result = (result - 0x60) & 0xff
		}
	} else {
		if (c.F&FlagH != 0) || (result&0x0f > 9) {
			result += 0x06
		}
		if (c.F&FlagC != 0) || (result > 0x9f) {
			result += 0x60
		}
	}
	// Flags z - 0 c
	c.F &= FlagN | FlagC
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
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
	c.F &= FlagC
	if c.L&0x0f == 0x0f {
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
	if c.L&0x0f == 0 {
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

// 2F: CPL		4 cycles
type op2f struct {
	SingleStepOp
}

func (op *op2f) Execute(c *CPU) (done bool) {
	// Flags: z 1 1 c
	c.F |= FlagN|FlagH
	c.A ^= 0xff
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

// 33: INC SP			8 cycles
type op33 struct {
	MultiStepsOp
}

func (op *op33) Tick() (done bool) {
	op.cpu.SP++
	return true
}

// 34: INC (HL)			12 cycles
type op34 struct {
	MultiStepsOp
}

func (op *op34) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 h -
		op.cpu.F &= FlagC
		if op.cpu.temp8&0x0f == 0x0f {
			op.cpu.F |= FlagH
		}
		op.cpu.temp8++
		if op.cpu.temp8 == 0 {
			op.cpu.F |= FlagZ
		}
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// 35: DEC (HL)			12 cycles
type op35 struct {
	MultiStepsOp
}

func (op *op35) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 1 h -
		op.cpu.F &= FlagC
		op.cpu.F |= FlagN
		if op.cpu.temp8&0x0f == 0 {
			op.cpu.F |= FlagH
		}
		op.cpu.temp8--
		if op.cpu.temp8 == 0 {
			op.cpu.F |= FlagZ
		}
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// 36: LD (HL),d8		12 cycles
type op36 struct {
	MultiStepsOp
}

func (op *op36) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.NextByte()
		op.step++
	case 1:
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// 37: SCF			4 cycles
type op37 struct {
	SingleStepOp
}

func (op *op37) Execute(c *CPU) (done bool) {
	z := c.F & FlagZ
	// Flags: - 0 0 1
	c.F = FlagC | z
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
	c.F &= FlagC
	if c.A&0x0f == 0x0f {
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
	if c.A&0x0f == 0 {
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

// 3F: CCF			4 cycles
type op3f struct {
	SingleStepOp
}

func (op *op3f) Execute(c *CPU) (done bool) {
	z := c.F & FlagZ
	// Flags: - 0 0 c
	c.F = (c.F ^ FlagC) & FlagC
	c.F |= z
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

// 76: HALT				4 cycles
type op76 struct {
	SingleStepOp
}

func (op *op76) Execute(c *CPU) (done bool) {
	// TODO: implement HALT bug [TCAGBD 4.10]
	c.state = states.Halted
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

// 88: ADC A,B		4 cycles
type op88 struct {
	SingleStepOp
}

func (op *op88) Execute(c *CPU) (done bool) {
	carry := (c.F & FlagC) >> 4
	// Flags: z 0 h c
	c.F = 0
	if (c.A & 0x0f) + (c.B & 0x0f) + carry > 0x0f {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.B) + uint(carry)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 89: ADC A,C		4 cycles
type op89 struct {
	SingleStepOp
}

func (op *op89) Execute(c *CPU) (done bool) {
	carry := (c.F & FlagC) >> 4
	// Flags: z 0 h c
	c.F = 0
	if (c.A & 0x0f) + (c.C & 0x0f) + carry > 0x0f {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.C) + uint(carry)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 8A: ADC A,D		4 cycles
type op8a struct {
	SingleStepOp
}

func (op *op8a) Execute(c *CPU) (done bool) {
	carry := (c.F & FlagC) >> 4
	// Flags: z 0 h c
	c.F = 0
	if (c.A & 0x0f) + (c.D & 0x0f) + carry > 0x0f {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.D) + uint(carry)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 8B: ADC A,E		4 cycles
type op8b struct {
	SingleStepOp
}

func (op *op8b) Execute(c *CPU) (done bool) {
	carry := (c.F & FlagC) >> 4
	// Flags: z 0 h c
	c.F = 0
	if (c.A & 0x0f) + (c.E & 0x0f) + carry > 0x0f {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.E) + uint(carry)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 8C: ADC A,H		4 cycles
type op8c struct {
	SingleStepOp
}

func (op *op8c) Execute(c *CPU) (done bool) {
	carry := (c.F & FlagC) >> 4
	// Flags: z 0 h c
	c.F = 0
	if (c.A & 0x0f) + (c.H & 0x0f) + carry > 0x0f {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.H) + uint(carry)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 8D: ADC A,L		4 cycles
type op8d struct {
	SingleStepOp
}

func (op *op8d) Execute(c *CPU) (done bool) {
	carry := (c.F & FlagC) >> 4
	// Flags: z 0 h c
	c.F = 0
	if (c.A & 0x0f) + (c.L & 0x0f) + carry > 0x0f {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.L) + uint(carry)
	if result > 0xff {
		c.F |= FlagC
	}
	c.A = uint8(result & 0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}

// 8E: ADC A,(HL)		8 cycles
type op8e struct {
	MultiStepsOp
}

func (op *op8e) Tick() (done bool) {
	carry := (op.cpu.F & FlagC) >> 4
	// Flags: z 0 h c
	op.cpu.F = 0
	value := op.cpu.MMU.Read(uint(op.cpu.HL()))
	if (op.cpu.A & 0x0f) + (value & 0x0f) + carry > 0x0f {
		op.cpu.F |= FlagH
	}
	result := uint(op.cpu.A) + uint(value) + uint(carry)
	if result > 0xff {
		op.cpu.F |= FlagC
	}
	op.cpu.A = uint8(result & 0xff)
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
    return true
}


// 8F: ADC A,A		4 cycles
type op8f struct {
	SingleStepOp
}

func (op *op8f) Execute(c *CPU) (done bool) {
	carry := (c.F & FlagC) >> 4
	// Flags: z 0 h c
	c.F = 0
	if (c.A & 0x0f) + (c.A & 0x0f) + carry > 0x0f {
		c.F |= FlagH
	}
	result := uint(c.A) + uint(c.A) + uint(carry)
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


// 98: SBC A,B		4 cycles
type op98 struct {
	SingleStepOp
}

func (op *op98) Execute(c *CPU) (done bool) {
	carry := int((c.F & FlagC) >> 4)
	// Flags: z 1 h c
	c.F = FlagN
	result := int(c.A) - int(c.B) - carry
	// Trusting the Internet on this one.
	if (c.A ^ c.B ^ uint8(result&0xff)) & (1 << 4) != 0 {
		c.F |= FlagH
	}
	if result < 0 {
		c.F |= FlagC
	}
	c.A = uint8(result&0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}


// 99: SBC A,C		4 cycles
type op99 struct {
	SingleStepOp
}

func (op *op99) Execute(c *CPU) (done bool) {
	carry := int((c.F & FlagC) >> 4)
	// Flags: z 1 h c
	c.F = FlagN
	result := int(c.A) - int(c.C) - carry
	// Trusting the Internet on this one.
	if (c.A ^ c.C ^ uint8(result&0xff)) & (1 << 4) != 0 {
		c.F |= FlagH
	}
	if result < 0 {
		c.F |= FlagC
	}
	c.A = uint8(result&0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}


// 9A: SBC A,D		4 cycles
type op9a struct {
	SingleStepOp
}

func (op *op9a) Execute(c *CPU) (done bool) {
	carry := int((c.F & FlagC) >> 4)
	// Flags: z 1 h c
	c.F = FlagN
	result := int(c.A) - int(c.D) - carry
	// Trusting the Internet on this one.
	if (c.A ^ c.D ^ uint8(result&0xff)) & (1 << 4) != 0 {
		c.F |= FlagH
	}
	if result < 0 {
		c.F |= FlagC
	}
	c.A = uint8(result&0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}


// 9B: SBC A,E		4 cycles
type op9b struct {
	SingleStepOp
}

func (op *op9b) Execute(c *CPU) (done bool) {
	carry := int((c.F & FlagC) >> 4)
	// Flags: z 1 h c
	c.F = FlagN
	result := int(c.A) - int(c.E) - carry
	// Trusting the Internet on this one.
	if (c.A ^ c.E ^ uint8(result&0xff)) & (1 << 4) != 0 {
		c.F |= FlagH
	}
	if result < 0 {
		c.F |= FlagC
	}
	c.A = uint8(result&0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}


// 9C: SBC A,H		4 cycles
type op9c struct {
	SingleStepOp
}

func (op *op9c) Execute(c *CPU) (done bool) {
	carry := int((c.F & FlagC) >> 4)
	// Flags: z 1 h c
	c.F = FlagN
	result := int(c.A) - int(c.H) - carry
	// Trusting the Internet on this one.
	if (c.A ^ c.H ^ uint8(result&0xff)) & (1 << 4) != 0 {
		c.F |= FlagH
	}
	if result < 0 {
		c.F |= FlagC
	}
	c.A = uint8(result&0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}


// 9D: SBC A,L		4 cycles
type op9d struct {
	SingleStepOp
}

func (op *op9d) Execute(c *CPU) (done bool) {
	carry := int((c.F & FlagC) >> 4)
	// Flags: z 1 h c
	c.F = FlagN
	result := int(c.A) - int(c.L) - carry
	// Trusting the Internet on this one.
	if (c.A ^ c.L ^ uint8(result&0xff)) & (1 << 4) != 0 {
		c.F |= FlagH
	}
	if result < 0 {
		c.F |= FlagC
	}
	c.A = uint8(result&0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}


// 9E: SBC A,(HL)		8 cycles
type op9e struct {
	MultiStepsOp
}

func (op *op9e) Tick() (done bool) {
	carry := int((op.cpu.F & FlagC) >> 4)
	// Flags: z 1 h c
	op.cpu.F = FlagN
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	result := int(op.cpu.A) - int(op.cpu.temp8) - carry
	// Trusting the Internet on this one.
	if (op.cpu.A ^ op.cpu.temp8 ^ uint8(result&0xff)) & (1 << 4) != 0 {
		op.cpu.F |= FlagH
	}
	if result < 0 {
		op.cpu.F |= FlagC
	}
	op.cpu.A = uint8(result&0xff)
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
    return true
}


// 9F: SBC A,A		4 cycles
type op9f struct {
	SingleStepOp
}

func (op *op9f) Execute(c *CPU) (done bool) {
	carry := int((c.F & FlagC) >> 4)
	// Flags: z 1 h c
	c.F = FlagN
	result := int(c.A) - int(c.A) - carry
	// Trusting the Internet on this one.
	if (c.A ^ c.A ^ uint8(result&0xff)) & (1 << 4) != 0 {
		c.F |= FlagH
	}
	if result < 0 {
		c.F |= FlagC
	}
	c.A = uint8(result&0xff)
	if c.A == 0 {
		c.F |= FlagZ
	}
    return true
}


// A0: AND B			4 cycles
type opA0 struct {
	SingleStepOp
}

func (op *opA0) Execute(c *CPU) (done bool) {
	c.A &= c.B
	// Flags z 0 1 0
	c.F = FlagH
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// A1: AND C			4 cycles
type opA1 struct {
	SingleStepOp
}

func (op *opA1) Execute(c *CPU) (done bool) {
	c.A &= c.C
	// Flags z 0 1 0
	c.F = FlagH
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// A2: AND D			4 cycles
type opA2 struct {
	SingleStepOp
}

func (op *opA2) Execute(c *CPU) (done bool) {
	c.A &= c.D
	// Flags z 0 1 0
	c.F = FlagH
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// A3: AND E			4 cycles
type opA3 struct {
	SingleStepOp
}

func (op *opA3) Execute(c *CPU) (done bool) {
	c.A &= c.E
	// Flags z 0 1 0
	c.F = FlagH
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// A4: AND H			4 cycles
type opA4 struct {
	SingleStepOp
}

func (op *opA4) Execute(c *CPU) (done bool) {
	c.A &= c.H
	// Flags z 0 1 0
	c.F = FlagH
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// A5: AND L			4 cycles
type opA5 struct {
	SingleStepOp
}

func (op *opA5) Execute(c *CPU) (done bool) {
	c.A &= c.L
	// Flags z 0 1 0
	c.F = FlagH
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// A6: AND (HL)			8 cycles
type opA6 struct {
	MultiStepsOp
}

func (op *opA6) Tick() (done bool) {
	op.cpu.A &= op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 0
	op.cpu.F = FlagH
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
	return true
}

// A7: AND A			4 cycles
type opA7 struct {
	SingleStepOp
}

func (op *opA7) Execute(c *CPU) (done bool) {
	c.A &= c.A
	// Flags z 0 1 0
	c.F = FlagH
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	op.cpu.F = 0
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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
	op.cpu.F = 0
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
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
	c.F = 0
	if c.A == 0 {
		c.F |= FlagZ
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


// C0: RET NZ		20/8 cycles
type opC0 struct {
	MultiStepsOp
}

func (op *opC0) Tick() (done bool) {
	switch op.step {
	case 0:
		if op.cpu.F&FlagZ != FlagZ {
			op.step++
		} else {
			done = true
		}
	case 1:
		op.cpu.PC = uint16(op.cpu.MMU.Read(uint(op.cpu.SP)))
		op.cpu.SP++
		op.step++
	case 2:
		op.cpu.PC |= uint16(op.cpu.MMU.Read(uint(op.cpu.SP))) << 8
		op.cpu.SP++
		op.step++
	case 3:
		// [GEKKIO] mentions an internal delay.
		done = true
	}
	return
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

// C2: JP NZ,a16		16/12 cycles
type opC2 struct {
	MultiStepsOp
}

func (op *opC2) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		if op.cpu.F&FlagZ != FlagZ {
			op.step++
		} else {
			done = true
		}
	case 2:
		op.cpu.PC = op.cpu.temp16
		done = true
	}
	return
}

// C3: JP a16		16 cycles
type opC3 struct {
	MultiStepsOp
}

func (op *opC3) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		op.step++
	case 2:
		op.cpu.PC = op.cpu.temp16
		done = true
	}
	return
}

// C4: CALL NZ,a16		24/12 cycles
type opC4 struct {
	MultiStepsOp
}

func (op *opC4) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		if op.cpu.F&FlagZ != FlagZ {
			op.step++
		} else {
			done = true
		}
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 3:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 4:
		op.cpu.PC = op.cpu.temp16
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

// C6: ADD A,d8		8 cycles
type opC6 struct {
	MultiStepsOp
}

func (op *opC6) Tick() (done bool) {
	// Flags: z 0 h c
	op.cpu.F = 0
	value := op.cpu.NextByte()
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


// C7: RST 00H		16 cycles
type opC7 struct {
	MultiStepsOp
}

func (op *opC7) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 2:
		op.cpu.PC = 0x00
		done = true
	}
	return
}

// C8: RET Z		20/8 cycles
type opC8 struct {
	MultiStepsOp
}

func (op *opC8) Tick() (done bool) {
	switch op.step {
	case 0:
		if op.cpu.F&FlagZ == FlagZ {
			op.step++
		} else {
			done = true
		}
	case 1:
		op.cpu.PC = uint16(op.cpu.MMU.Read(uint(op.cpu.SP)))
		op.cpu.SP++
		op.step++
	case 2:
		op.cpu.PC |= uint16(op.cpu.MMU.Read(uint(op.cpu.SP))) << 8
		op.cpu.SP++
		op.step++
	case 3:
		// [GEKKIO] mentions an internal delay.
		done = true
	}
	return
}

// C9: RET 		16 cycles
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

// CA: JP Z,a16		16/12 cycles
type opCa struct {
	MultiStepsOp
}

func (op *opCa) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		if op.cpu.F&FlagZ == FlagZ {
			op.step++
		} else {
			done = true
		}
	case 2:
		op.cpu.PC = op.cpu.temp16
		done = true
	}
	return
}

// CC: CALL Z,a16		24/12 cycles
type opCc struct {
	MultiStepsOp
}

func (op *opCc) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		if op.cpu.F&FlagZ == FlagZ {
			op.step++
		} else {
			done = true
		}
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 3:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 4:
		op.cpu.PC = op.cpu.temp16
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
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 4:
		op.cpu.PC = op.cpu.temp16
		done = true
	}
	return
}

// CE: ADC A,d8		8 cycles
type opCe struct {
	MultiStepsOp
}

func (op *opCe) Tick() (done bool) {
	carry := (op.cpu.F & FlagC) >> 4
	// Flags: z 0 h c
	op.cpu.F = 0
	value := op.cpu.NextByte()
	if (op.cpu.A & 0x0f) + (value & 0x0f) + carry > 0x0f {
		op.cpu.F |= FlagH
	}
	result := uint(op.cpu.A) + uint(value) + uint(carry)
	if result > 0xff {
		op.cpu.F |= FlagC
	}
	op.cpu.A = uint8(result & 0xff)
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
    return true
}


// CF: RST 08H		16 cycles
type opCf struct {
	MultiStepsOp
}

func (op *opCf) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 2:
		op.cpu.PC = 0x08
		done = true
	}
	return
}

// D0: RET NC		20/8 cycles
type opD0 struct {
	MultiStepsOp
}

func (op *opD0) Tick() (done bool) {
	switch op.step {
	case 0:
		if op.cpu.F&FlagC != FlagC {
			op.step++
		} else {
			done = true
		}
	case 1:
		op.cpu.PC = uint16(op.cpu.MMU.Read(uint(op.cpu.SP)))
		op.cpu.SP++
		op.step++
	case 2:
		op.cpu.PC |= uint16(op.cpu.MMU.Read(uint(op.cpu.SP))) << 8
		op.cpu.SP++
		op.step++
	case 3:
		// [GEKKIO] mentions an internal delay.
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

// D2: JP NC,a16		16/12 cycles
type opD2 struct {
	MultiStepsOp
}

func (op *opD2) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		if op.cpu.F&FlagC != FlagC {
			op.step++
		} else {
			done = true
		}
	case 2:
		op.cpu.PC = op.cpu.temp16
		done = true
	}
	return
}

// D4: CALL NC,a16		24/12 cycles
type opD4 struct {
	MultiStepsOp
}

func (op *opD4) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		if op.cpu.F&FlagC != FlagC {
			op.step++
		} else {
			done = true
		}
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 3:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 4:
		op.cpu.PC = op.cpu.temp16
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

// D6: SUB d8		8 cycles
type opD6 struct {
	MultiStepsOp
}

func (op *opD6) Tick() (done bool) {
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
	op.cpu.A = result
    return true
}


// D7: RST 10H		16 cycles
type opD7 struct {
	MultiStepsOp
}

func (op *opD7) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 2:
		op.cpu.PC = 0x10
		done = true
	}
	return
}

// D8: RET C		20/8 cycles
type opD8 struct {
	MultiStepsOp
}

func (op *opD8) Tick() (done bool) {
	switch op.step {
	case 0:
		if op.cpu.F&FlagC == FlagC {
			op.step++
		} else {
			done = true
		}
	case 1:
		op.cpu.PC = uint16(op.cpu.MMU.Read(uint(op.cpu.SP)))
		op.cpu.SP++
		op.step++
	case 2:
		op.cpu.PC |= uint16(op.cpu.MMU.Read(uint(op.cpu.SP))) << 8
		op.cpu.SP++
		op.step++
	case 3:
		// [GEKKIO] mentions an internal delay.
		done = true
	}
	return
}

// D9: RETI 		16 cycles
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

// DA: JP C,a16		16/12 cycles
type opDa struct {
	MultiStepsOp
}

func (op *opDa) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		if op.cpu.F&FlagC == FlagC {
			op.step++
		} else {
			done = true
		}
	case 2:
		op.cpu.PC = op.cpu.temp16
		done = true
	}
	return
}

// DC: CALL C,a16		24/12 cycles
type opDc struct {
	MultiStepsOp
}

func (op *opDc) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
		op.step++
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
		if op.cpu.F&FlagC == FlagC {
			op.step++
		} else {
			done = true
		}
	case 2:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 3:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 4:
		op.cpu.PC = op.cpu.temp16
		done = true
	}
	return
}

// DE: SBC A,d8		8 cycles
type opDe struct {
	MultiStepsOp
}

func (op *opDe) Tick() (done bool) {
	carry := int((op.cpu.F & FlagC) >> 4)
	// Flags: z 1 h c
	op.cpu.F = FlagN
	op.cpu.temp8 = op.cpu.NextByte()
	result := int(op.cpu.A) - int(op.cpu.temp8) - carry
	// Trusting the Internet on this one.
	if (op.cpu.A ^ op.cpu.temp8 ^ uint8(result&0xff)) & (1 << 4) != 0 {
		op.cpu.F |= FlagH
	}
	if result < 0 {
		op.cpu.F |= FlagC
	}
	op.cpu.A = uint8(result&0xff)
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
    return true
}


// DF: RST 18H		16 cycles
type opDf struct {
	MultiStepsOp
}

func (op *opDf) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 2:
		op.cpu.PC = 0x18
		done = true
	}
	return
}

// E0: LD (FF00+a8),A		12 cycles
type opE0 struct {
	MultiStepsOp
}

func (op *opE0) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.NextByte()
		op.step++
	case 1:
		op.cpu.MMU.Write(uint(0xff00+uint16(op.cpu.temp8)), op.cpu.A)
		done = true
	}
	return
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

// E6: AND d8			8 cycles
type opE6 struct {
	MultiStepsOp
}

func (op *opE6) Tick() (done bool) {
	op.cpu.A &= op.cpu.NextByte()
	// Flags z 0 1 0
	op.cpu.F = FlagH
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
	return true
}

// E7: RST 20H		16 cycles
type opE7 struct {
	MultiStepsOp
}

func (op *opE7) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 2:
		op.cpu.PC = 0x20
		done = true
	}
	return
}

// E8: ADD SP,r8		16 cycles
type opE8 struct {
	MultiStepsOp
	offset int8
}

func (op *opE8) Tick() (done bool) {
	switch op.step {
	case 0:
		op.offset = int8(op.cpu.NextByte())
		op.step++
	case 1:
		// [REF NEEDED] extra cycle
		op.step++
	case 2:
		// Flags: 0 0 h c
		op.cpu.F = 0

		// Need cast to signed for the potential substraction
		if (int16(op.cpu.SP)&0x0f+int16(op.offset)&0x0f) & 0x10 != 0 {
			op.cpu.F |= FlagH
		}
		if (int16(op.cpu.SP)&0xff+int16(op.offset)&0xff) & 0x100 != 0 {
			op.cpu.F |= FlagC
		}
		result := int16(op.cpu.SP) + int16(op.offset)
		op.cpu.SP = uint16(result)
		done = true
	}
	return
}

// E9: JP HL			4 cycles
type opE9 struct {
	SingleStepOp
}

func (op *opE9) Execute(c *CPU) (done bool) {
	c.PC = c.HL()
	return true
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

// EE: XOR d8			8 cycles
type opEe struct {
	MultiStepsOp
}

func (op *opEe) Tick() (done bool) {
	op.cpu.A ^= op.cpu.NextByte()
	// Flags z 0 0 0
	op.cpu.F = 0
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
	return true
}

// EF: RST 28H		16 cycles
type opEf struct {
	MultiStepsOp
}

func (op *opEf) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 2:
		op.cpu.PC = 0x28
		done = true
	}
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
		op.cpu.F = op.cpu.MMU.Read(uint(op.cpu.SP)) & 0xf0
		op.cpu.SP++
		op.step++
	case 1:
		op.cpu.A = op.cpu.MMU.Read(uint(op.cpu.SP))
		op.cpu.SP++
		done = true
	}
	return
}

// F2: LD A,(FF00+C)		8 cycles
type opF2 struct {
	MultiStepsOp
}

func (op *opF2) Tick() (done bool) {
	op.cpu.A = op.cpu.MMU.Read(uint(0xff00+uint16(op.cpu.C)))
	return true
}

// F3: DI		4 cycles
type opF3 struct {
	SingleStepOp
}

func (op *opF3) Execute(c *CPU) (done bool) {
	c.IME = false
    return true
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

// F6: OR d8			8 cycles
type opF6 struct {
	MultiStepsOp
}

func (op *opF6) Tick() (done bool) {
	op.cpu.A |= op.cpu.NextByte()
	// Flags z 0 0 0
	op.cpu.F = 0
	if op.cpu.A == 0 {
		op.cpu.F |= FlagZ
	}
	return true
}

// F7: RST 30H		16 cycles
type opF7 struct {
	MultiStepsOp
}

func (op *opF7) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 2:
		op.cpu.PC = 0x30
		done = true
	}
	return
}

// F8: LD HL,SP+r8		12 cycles
type opF8 struct {
	MultiStepsOp
	offset int8
}

func (op *opF8) Tick() (done bool) {
	switch op.step {
	case 0:
		op.offset = int8(op.cpu.NextByte())
		op.step++
	case 1:
		// Flags: 0 0 h c
		op.cpu.F = 0

		// Need cast to signed for the potential substraction
		if (int16(op.cpu.SP)&0x0f+int16(op.offset)&0x0f) & 0x10 != 0 {
			op.cpu.F |= FlagH
		}
		if (int16(op.cpu.SP)&0xff+int16(op.offset)&0xff) & 0x100 != 0 {
			op.cpu.F |= FlagC
		}
		result := int16(op.cpu.SP) + int16(op.offset)
		op.cpu.SetHL(uint16(result))
		done = true
	}
	return
}

// F9: LD SP,HL			8 cycles
type opF9 struct {
	MultiStepsOp
}

func (op *opF9) Tick() (done bool) {
	op.cpu.SP = op.cpu.HL()
	return true
}

// FA: LD A,(a16)		16 cycles
type opFa struct {
	MultiStepsOp
}

func (op *opFa) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp16 = uint16(op.cpu.NextByte())
	case 1:
		op.cpu.temp16 |= uint16(op.cpu.NextByte()) << 8
	case 2:
		op.cpu.A = op.cpu.MMU.Read(uint(op.cpu.temp16))
		done = true
	}
	op.step++
	return
}

// FB: EI		4 cycles
type opFb struct {
	SingleStepOp
}

func (op *opFb) Execute(c *CPU) (done bool) {
	c.IME = true
    return true
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


// FF: RST 38H		16 cycles
type opFf struct {
	MultiStepsOp
}

func (op *opFf) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC>>8))
		op.step++
	case 1:
		op.cpu.SP--
		op.cpu.MMU.Write(uint(op.cpu.SP), uint8(op.cpu.PC&0x00ff))
		op.step++
	case 2:
		op.cpu.PC = 0x38
		done = true
	}
	return
}

// CB 00: RLC B				8 cycles
type opCb00 struct {
	SingleStepOp
}

func (op *opCb00) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	result := c.B << 1 & 0xff
	if c.B&0x80 != 0 {
		result |= 1
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.B = result
	return true
}

// CB 01: RLC C				8 cycles
type opCb01 struct {
	SingleStepOp
}

func (op *opCb01) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	result := c.C << 1 & 0xff
	if c.C&0x80 != 0 {
		result |= 1
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.C = result
	return true
}

// CB 02: RLC D				8 cycles
type opCb02 struct {
	SingleStepOp
}

func (op *opCb02) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	result := c.D << 1 & 0xff
	if c.D&0x80 != 0 {
		result |= 1
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.D = result
	return true
}

// CB 03: RLC E				8 cycles
type opCb03 struct {
	SingleStepOp
}

func (op *opCb03) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	result := c.E << 1 & 0xff
	if c.E&0x80 != 0 {
		result |= 1
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.E = result
	return true
}

// CB 04: RLC H				8 cycles
type opCb04 struct {
	SingleStepOp
}

func (op *opCb04) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	result := c.H << 1 & 0xff
	if c.H&0x80 != 0 {
		result |= 1
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.H = result
	return true
}

// CB 05: RLC L				8 cycles
type opCb05 struct {
	SingleStepOp
}

func (op *opCb05) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	result := c.L << 1 & 0xff
	if c.L&0x80 != 0 {
		result |= 1
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.L = result
	return true
}

// CB 06: RLC (HL)				16 cycles
type opCb06 struct {
	MultiStepsOp
}

func (op *opCb06) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 0 c
		op.cpu.F = 0x00
		result := op.cpu.temp8 << 1 & 0xff
		if op.cpu.temp8&0x80 != 0 {
			result |= 1
			op.cpu.F |= FlagC
		}

		if result == 0 {
			op.cpu.F |= FlagZ
		}
		op.cpu.MMU.Write(uint(op.cpu.HL()), result)
		done = true
	}
	return
}

// CB 07: RLC A				8 cycles
type opCb07 struct {
	SingleStepOp
}

func (op *opCb07) Execute(c *CPU) (done bool) {
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

// CB 08: RRC B				8 cycles
type opCb08 struct {
	SingleStepOp
}

func (op *opCb08) Execute(c *CPU) (done bool) {
	result := c.B >> 1
	// Flags z 0 0 c
	c.F = 0x00
	if c.B&1 > 0 {
		result |= (1 << 7)
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.B = result

	return true
}

// CB 09: RRC C				8 cycles
type opCb09 struct {
	SingleStepOp
}

func (op *opCb09) Execute(c *CPU) (done bool) {
	result := c.C >> 1
	// Flags z 0 0 c
	c.F = 0x00
	if c.C&1 > 0 {
		result |= (1 << 7)
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.C = result

	return true
}

// CB 0A: RRC D				8 cycles
type opCb0a struct {
	SingleStepOp
}

func (op *opCb0a) Execute(c *CPU) (done bool) {
	result := c.D >> 1
	// Flags z 0 0 c
	c.F = 0x00
	if c.D&1 > 0 {
		result |= (1 << 7)
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.D = result

	return true
}

// CB 0B: RRC E				8 cycles
type opCb0b struct {
	SingleStepOp
}

func (op *opCb0b) Execute(c *CPU) (done bool) {
	result := c.E >> 1
	// Flags z 0 0 c
	c.F = 0x00
	if c.E&1 > 0 {
		result |= (1 << 7)
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.E = result

	return true
}

// CB 0C: RRC H				8 cycles
type opCb0c struct {
	SingleStepOp
}

func (op *opCb0c) Execute(c *CPU) (done bool) {
	result := c.H >> 1
	// Flags z 0 0 c
	c.F = 0x00
	if c.H&1 > 0 {
		result |= (1 << 7)
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.H = result

	return true
}

// CB 0D: RRC L				8 cycles
type opCb0d struct {
	SingleStepOp
}

func (op *opCb0d) Execute(c *CPU) (done bool) {
	result := c.L >> 1
	// Flags z 0 0 c
	c.F = 0x00
	if c.L&1 > 0 {
		result |= (1 << 7)
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.L = result

	return true
}

// CB 0E: RRC (HL)				16 cycles
type opCb0e struct {
	MultiStepsOp
}

func (op *opCb0e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 0 c
		op.cpu.F = 0x00
		result := op.cpu.temp8 >> 1
		if op.cpu.temp8&1 != 0 {
			result |= (1 << 7)
			op.cpu.F |= FlagC
		}
		if result == 0 {
			op.cpu.F |= FlagZ
		}

		op.cpu.MMU.Write(uint(op.cpu.HL()), result)
		done = true
	}
	return
}

// CB 0F: RRC A				8 cycles
type opCb0f struct {
	SingleStepOp
}

func (op *opCb0f) Execute(c *CPU) (done bool) {
	result := c.A >> 1
	// Flags z 0 0 c
	c.F = 0x00
	if c.A&1 > 0 {
		result |= (1 << 7)
		c.F |= FlagC
	}
	if result == 0 {
		c.F |= FlagZ
	}
	c.A = result

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

// CB 16: RL (HL)				16 cycles
type opCb16 struct {
	MultiStepsOp
}

func (op *opCb16) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 0 c
		result := (op.cpu.temp8 << 1) & 0xff
		if op.cpu.F&FlagC != 0 {
			result |= 1
		}
		op.cpu.F = 0x00
		if op.cpu.temp8&(1<<7) != 0 {
			op.cpu.F |= FlagC
		}
		if result == 0 {
			op.cpu.F |= FlagZ
		}
		op.cpu.MMU.Write(uint(op.cpu.HL()), result)
		done = true
	}
	return
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

// CB 18: RR B				8 cycles
type opCb18 struct {
	SingleStepOp
}

func (op *opCb18) Execute(c *CPU) (done bool) {
	result := c.B >> 1
	if c.F&FlagC > 0 {
		result |= (1<<7)
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.B&1 > 0 {
		c.F |= FlagC
	}
	c.B = result

	return true
}

// CB 19: RR C				8 cycles
type opCb19 struct {
	SingleStepOp
}

func (op *opCb19) Execute(c *CPU) (done bool) {
	result := c.C >> 1
	if c.F&FlagC > 0 {
		result |= (1<<7)
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.C&1 > 0 {
		c.F |= FlagC
	}
	c.C = result

	return true
}

// CB 1A: RR D				8 cycles
type opCb1a struct {
	SingleStepOp
}

func (op *opCb1a) Execute(c *CPU) (done bool) {
	result := c.D >> 1
	if c.F&FlagC > 0 {
		result |= (1<<7)
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.D&1 > 0 {
		c.F |= FlagC
	}
	c.D = result

	return true
}

// CB 1B: RR E				8 cycles
type opCb1b struct {
	SingleStepOp
}

func (op *opCb1b) Execute(c *CPU) (done bool) {
	result := c.E >> 1
	if c.F&FlagC > 0 {
		result |= (1<<7)
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.E&1 > 0 {
		c.F |= FlagC
	}
	c.E = result

	return true
}

// CB 1C: RR H				8 cycles
type opCb1c struct {
	SingleStepOp
}

func (op *opCb1c) Execute(c *CPU) (done bool) {
	result := c.H >> 1
	if c.F&FlagC > 0 {
		result |= (1<<7)
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.H&1 > 0 {
		c.F |= FlagC
	}
	c.H = result

	return true
}

// CB 1D: RR L				8 cycles
type opCb1d struct {
	SingleStepOp
}

func (op *opCb1d) Execute(c *CPU) (done bool) {
	result := c.L >> 1
	if c.F&FlagC > 0 {
		result |= (1<<7)
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.L&1 > 0 {
		c.F |= FlagC
	}
	c.L = result

	return true
}

// CB 1E: RR (HL)				16 cycles
type opCb1e struct {
	MultiStepsOp
}

func (op *opCb1e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 0 c
		result := op.cpu.temp8 >> 1
		if op.cpu.F&FlagC != 0 {
			result |= (1<<7)
		}
		op.cpu.F = 0x00
		if result == 0 {
			op.cpu.F |= FlagZ
		}
		if op.cpu.temp8&1 != 0 {
			op.cpu.F |= FlagC
		}
		op.cpu.MMU.Write(uint(op.cpu.HL()), result)
		done = true
	}
	return
}

// CB 1F: RR A				8 cycles
type opCb1f struct {
	SingleStepOp
}

func (op *opCb1f) Execute(c *CPU) (done bool) {
	result := c.A >> 1
	if c.F&FlagC > 0 {
		result |= (1<<7)
	}
	// Flags z 0 0 c
	c.F = 0x00
	if result == 0 {
		c.F |= FlagZ
	}
	if c.A&1 > 0 {
		c.F |= FlagC
	}
	c.A = result

	return true
}

// CB 20: SLA	B			8 cycles
type opCb20 struct {
	SingleStepOp
}

func (op *opCb20) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	if c.B&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.B <<= 1
	if c.B == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 21: SLA	C			8 cycles
type opCb21 struct {
	SingleStepOp
}

func (op *opCb21) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	if c.C&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.C <<= 1
	if c.C == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 22: SLA	D			8 cycles
type opCb22 struct {
	SingleStepOp
}

func (op *opCb22) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	if c.D&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.D <<= 1
	if c.D == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 23: SLA	E			8 cycles
type opCb23 struct {
	SingleStepOp
}

func (op *opCb23) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	if c.E&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.E <<= 1
	if c.E == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 24: SLA	H			8 cycles
type opCb24 struct {
	SingleStepOp
}

func (op *opCb24) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	if c.H&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.H <<= 1
	if c.H == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 25: SLA	L			8 cycles
type opCb25 struct {
	SingleStepOp
}

func (op *opCb25) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	if c.L&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.L <<= 1
	if c.L == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 26: SLA HL			16 cycles
type opCb26 struct {
	MultiStepsOp
}

func (op *opCb26) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 0 c
		op.cpu.F = 0x00
		if op.cpu.temp8&(1<<7) > 0 {
			op.cpu.F |= FlagC
		}
		op.cpu.temp8 <<= 1
		if op.cpu.temp8 == 0 {
			op.cpu.F |= FlagZ
		}

		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB 27: SLA	A			8 cycles
type opCb27 struct {
	SingleStepOp
}

func (op *opCb27) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c
	c.F = 0x00
	if c.A&(1<<7) > 0 {
		c.F |= FlagC
	}
	c.A <<= 1
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 28: SRA B			8 cycles
type opCb28 struct {
	SingleStepOp
}

func (op *opCb28) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.B&1 != 0 {
		c.F |= FlagC
	}
	c.B = (c.B >> 1) | (c.B & (1 << 7))
	if c.B == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 29: SRA C			8 cycles
type opCb29 struct {
	SingleStepOp
}

func (op *opCb29) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.C&1 != 0 {
		c.F |= FlagC
	}
	c.C = (c.C >> 1) | (c.C & (1 << 7))
	if c.C == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 2A: SRA D			8 cycles
type opCb2a struct {
	SingleStepOp
}

func (op *opCb2a) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.D&1 != 0 {
		c.F |= FlagC
	}
	c.D = (c.D >> 1) | (c.D & (1 << 7))
	if c.D == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 2B: SRA E			8 cycles
type opCb2b struct {
	SingleStepOp
}

func (op *opCb2b) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.E&1 != 0 {
		c.F |= FlagC
	}
	c.E = (c.E >> 1) | (c.E & (1 << 7))
	if c.E == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 2C: SRA H			8 cycles
type opCb2c struct {
	SingleStepOp
}

func (op *opCb2c) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.H&1 != 0 {
		c.F |= FlagC
	}
	c.H = (c.H >> 1) | (c.H & (1 << 7))
	if c.H == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 2D: SRA L			8 cycles
type opCb2d struct {
	SingleStepOp
}

func (op *opCb2d) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.L&1 != 0 {
		c.F |= FlagC
	}
	c.L = (c.L >> 1) | (c.L & (1 << 7))
	if c.L == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 2E: SRA HL			16 cycles
type opCb2e struct {
	MultiStepsOp
}

func (op *opCb2e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
		op.cpu.F = 0x00
		if op.cpu.temp8&1 != 0 {
			op.cpu.F |= FlagC
		}
		op.cpu.temp8 = (op.cpu.temp8 >> 1) | (op.cpu.temp8 & (1 << 7))
		if op.cpu.temp8 == 0 {
			op.cpu.F |= FlagZ
		}

		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB 2F: SRA A			8 cycles
type opCb2f struct {
	SingleStepOp
}

func (op *opCb2f) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.A&1 != 0 {
		c.F |= FlagC
	}
	c.A = (c.A >> 1) | (c.A & (1 << 7))
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 30: SWAP B			8 cycles
type opCb30 struct {
	SingleStepOp
}

func (op *opCb30) Execute(c *CPU) (done bool) {
	// Flags z 0 0 0
	c.F = 0x00
	c.B = ((c.B & 0x0f) << 4) | ((c.B & 0xf0) >> 4)
	if c.B == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 31: SWAP C			8 cycles
type opCb31 struct {
	SingleStepOp
}

func (op *opCb31) Execute(c *CPU) (done bool) {
	// Flags z 0 0 0
	c.F = 0x00
	c.C = ((c.C & 0x0f) << 4) | ((c.C & 0xf0) >> 4)
	if c.C == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 32: SWAP D			8 cycles
type opCb32 struct {
	SingleStepOp
}

func (op *opCb32) Execute(c *CPU) (done bool) {
	// Flags z 0 0 0
	c.F = 0x00
	c.D = ((c.D & 0x0f) << 4) | ((c.D & 0xf0) >> 4)
	if c.D == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 33: SWAP E			8 cycles
type opCb33 struct {
	SingleStepOp
}

func (op *opCb33) Execute(c *CPU) (done bool) {
	// Flags z 0 0 0
	c.F = 0x00
	c.E = ((c.E & 0x0f) << 4) | ((c.E & 0xf0) >> 4)
	if c.E == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 34: SWAP H			8 cycles
type opCb34 struct {
	SingleStepOp
}

func (op *opCb34) Execute(c *CPU) (done bool) {
	// Flags z 0 0 0
	c.F = 0x00
	c.H = ((c.H & 0x0f) << 4) | ((c.H & 0xf0) >> 4)
	if c.H == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 35: SWAP L			8 cycles
type opCb35 struct {
	SingleStepOp
}

func (op *opCb35) Execute(c *CPU) (done bool) {
	// Flags z 0 0 0
	c.F = 0x00
	c.L = ((c.L & 0x0f) << 4) | ((c.L & 0xf0) >> 4)
	if c.L == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 36:  HL			16 cycles
type opCb36 struct {
	MultiStepsOp
}

func (op *opCb36) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 0 0
		op.cpu.F = 0x00
		op.cpu.temp8 = ((op.cpu.temp8 & 0x0f) << 4) | ((op.cpu.temp8 & 0xf0) >> 4)
		if op.cpu.temp8 == 0 {
			op.cpu.F |= FlagZ
		}

		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB 37: SWAP A			8 cycles
type opCb37 struct {
	SingleStepOp
}

func (op *opCb37) Execute(c *CPU) (done bool) {
	// Flags z 0 0 0
	c.F = 0x00
	c.A = ((c.A & 0x0f) << 4) | ((c.A & 0xf0) >> 4)
	if c.A == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 38: SRL B			8 cycles
type opCb38 struct {
	SingleStepOp
}

func (op *opCb38) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.B&1 != 0 {
		c.F |= FlagC
	}
	c.B = (c.B >> 1)
	if c.B == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 39: SRL C			8 cycles
type opCb39 struct {
	SingleStepOp
}

func (op *opCb39) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.C&1 != 0 {
		c.F |= FlagC
	}
	c.C = (c.C >> 1)
	if c.C == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 3A: SRL D			8 cycles
type opCb3a struct {
	SingleStepOp
}

func (op *opCb3a) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.D&1 != 0 {
		c.F |= FlagC
	}
	c.D = (c.D >> 1)
	if c.D == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 3B: SRL E			8 cycles
type opCb3b struct {
	SingleStepOp
}

func (op *opCb3b) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.E&1 != 0 {
		c.F |= FlagC
	}
	c.E = (c.E >> 1)
	if c.E == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 3C: SRL H			8 cycles
type opCb3c struct {
	SingleStepOp
}

func (op *opCb3c) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.H&1 != 0 {
		c.F |= FlagC
	}
	c.H = (c.H >> 1)
	if c.H == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 3D: SRL L			8 cycles
type opCb3d struct {
	SingleStepOp
}

func (op *opCb3d) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.L&1 != 0 {
		c.F |= FlagC
	}
	c.L = (c.L >> 1)
	if c.L == 0 {
		c.F |= FlagZ
	}
	return true
}

// CB 3E: SRL HL			16 cycles
type opCb3e struct {
	MultiStepsOp
}

func (op *opCb3e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
		op.cpu.F = 0x00
		if op.cpu.temp8&1 != 0 {
			op.cpu.F |= FlagC
		}
		op.cpu.temp8 = (op.cpu.temp8 >> 1)
		if op.cpu.temp8 == 0 {
			op.cpu.F |= FlagZ
		}

		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB 3F: SRL A			8 cycles
type opCb3f struct {
	SingleStepOp
}

func (op *opCb3f) Execute(c *CPU) (done bool) {
	// Flags z 0 0 c (though [OPCODES] says z 0 0 0 for SRA)
	c.F = 0x00
	if c.A&1 != 0 {
		c.F |= FlagC
	}
	c.A = (c.A >> 1)
	if c.A == 0 {
		c.F |= FlagZ
	}
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

// CB 46: BIT 0,(HL)			12 cycles
type opCb46 struct {
	MultiStepsOp
}

func (op *opCb46) Tick() (done bool) {
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 -
	if op.cpu.temp8&(1<<0) == 0 {
		op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
	} else {
		op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
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

// CB 4E: BIT 1,(HL)			12 cycles
type opCb4e struct {
	MultiStepsOp
}

func (op *opCb4e) Tick() (done bool) {
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 -
	if op.cpu.temp8&(1<<1) == 0 {
		op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
	} else {
		op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
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

// CB 56: BIT 2,(HL)			12 cycles
type opCb56 struct {
	MultiStepsOp
}

func (op *opCb56) Tick() (done bool) {
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 -
	if op.cpu.temp8&(1<<2) == 0 {
		op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
	} else {
		op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
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

// CB 5E: BIT 3,(HL)			12 cycles
type opCb5e struct {
	MultiStepsOp
}

func (op *opCb5e) Tick() (done bool) {
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 -
	if op.cpu.temp8&(1<<3) == 0 {
		op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
	} else {
		op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
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

// CB 66: BIT 4,(HL)			12 cycles
type opCb66 struct {
	MultiStepsOp
}

func (op *opCb66) Tick() (done bool) {
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 -
	if op.cpu.temp8&(1<<4) == 0 {
		op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
	} else {
		op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
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

// CB 6E: BIT 5,(HL)			12 cycles
type opCb6e struct {
	MultiStepsOp
}

func (op *opCb6e) Tick() (done bool) {
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 -
	if op.cpu.temp8&(1<<5) == 0 {
		op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
	} else {
		op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
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

// CB 76: BIT 6,(HL)			12 cycles
type opCb76 struct {
	MultiStepsOp
}

func (op *opCb76) Tick() (done bool) {
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 -
	if op.cpu.temp8&(1<<6) == 0 {
		op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
	} else {
		op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
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

// CB 7E: BIT 7,(HL)			12 cycles
type opCb7e struct {
	MultiStepsOp
}

func (op *opCb7e) Tick() (done bool) {
	op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 1 -
	if op.cpu.temp8&(1<<7) == 0 {
		op.cpu.F = (op.cpu.F & ^FlagN) | FlagZ | FlagH
	} else {
		op.cpu.F = (op.cpu.F & ^(FlagN | FlagZ)) | FlagH
	}
	return true
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

// CB 80: RES 0,B			8 cycles
type opCb80 struct {
	SingleStepOp
}

func (op *opCb80) Execute(c *CPU) (done bool) {
	c.B &= ^uint8(1<<0)
	return true
}

// CB 81: RES 0,C			8 cycles
type opCb81 struct {
	SingleStepOp
}

func (op *opCb81) Execute(c *CPU) (done bool) {
	c.C &= ^uint8(1<<0)
	return true
}

// CB 82: RES 0,D			8 cycles
type opCb82 struct {
	SingleStepOp
}

func (op *opCb82) Execute(c *CPU) (done bool) {
	c.D &= ^uint8(1<<0)
	return true
}

// CB 83: RES 0,E			8 cycles
type opCb83 struct {
	SingleStepOp
}

func (op *opCb83) Execute(c *CPU) (done bool) {
	c.E &= ^uint8(1<<0)
	return true
}

// CB 84: RES 0,H			8 cycles
type opCb84 struct {
	SingleStepOp
}

func (op *opCb84) Execute(c *CPU) (done bool) {
	c.H &= ^uint8(1<<0)
	return true
}

// CB 85: RES 0,L			8 cycles
type opCb85 struct {
	SingleStepOp
}

func (op *opCb85) Execute(c *CPU) (done bool) {
	c.L &= ^uint8(1<<0)
	return true
}

// CB 86: RES 0,(HL)			16 cycles
type opCb86 struct {
	MultiStepsOp
}

func (op *opCb86) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 &= ^uint8(1<<0)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB 87: RES 0,A			8 cycles
type opCb87 struct {
	SingleStepOp
}

func (op *opCb87) Execute(c *CPU) (done bool) {
	c.A &= ^uint8(1<<0)
	return true
}

// CB 88: RES 1,B			8 cycles
type opCb88 struct {
	SingleStepOp
}

func (op *opCb88) Execute(c *CPU) (done bool) {
	c.B &= ^uint8(1<<1)
	return true
}

// CB 89: RES 1,C			8 cycles
type opCb89 struct {
	SingleStepOp
}

func (op *opCb89) Execute(c *CPU) (done bool) {
	c.C &= ^uint8(1<<1)
	return true
}

// CB 8A: RES 1,D			8 cycles
type opCb8a struct {
	SingleStepOp
}

func (op *opCb8a) Execute(c *CPU) (done bool) {
	c.D &= ^uint8(1<<1)
	return true
}

// CB 8B: RES 1,E			8 cycles
type opCb8b struct {
	SingleStepOp
}

func (op *opCb8b) Execute(c *CPU) (done bool) {
	c.E &= ^uint8(1<<1)
	return true
}

// CB 8C: RES 1,H			8 cycles
type opCb8c struct {
	SingleStepOp
}

func (op *opCb8c) Execute(c *CPU) (done bool) {
	c.H &= ^uint8(1<<1)
	return true
}

// CB 8D: RES 1,L			8 cycles
type opCb8d struct {
	SingleStepOp
}

func (op *opCb8d) Execute(c *CPU) (done bool) {
	c.L &= ^uint8(1<<1)
	return true
}

// CB 8E: RES 1,(HL)			16 cycles
type opCb8e struct {
	MultiStepsOp
}

func (op *opCb8e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 &= ^uint8(1<<1)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB 8F: RES 1,A			8 cycles
type opCb8f struct {
	SingleStepOp
}

func (op *opCb8f) Execute(c *CPU) (done bool) {
	c.A &= ^uint8(1<<1)
	return true
}

// CB 90: RES 2,B			8 cycles
type opCb90 struct {
	SingleStepOp
}

func (op *opCb90) Execute(c *CPU) (done bool) {
	c.B &= ^uint8(1<<2)
	return true
}

// CB 91: RES 2,C			8 cycles
type opCb91 struct {
	SingleStepOp
}

func (op *opCb91) Execute(c *CPU) (done bool) {
	c.C &= ^uint8(1<<2)
	return true
}

// CB 92: RES 2,D			8 cycles
type opCb92 struct {
	SingleStepOp
}

func (op *opCb92) Execute(c *CPU) (done bool) {
	c.D &= ^uint8(1<<2)
	return true
}

// CB 93: RES 2,E			8 cycles
type opCb93 struct {
	SingleStepOp
}

func (op *opCb93) Execute(c *CPU) (done bool) {
	c.E &= ^uint8(1<<2)
	return true
}

// CB 94: RES 2,H			8 cycles
type opCb94 struct {
	SingleStepOp
}

func (op *opCb94) Execute(c *CPU) (done bool) {
	c.H &= ^uint8(1<<2)
	return true
}

// CB 95: RES 2,L			8 cycles
type opCb95 struct {
	SingleStepOp
}

func (op *opCb95) Execute(c *CPU) (done bool) {
	c.L &= ^uint8(1<<2)
	return true
}

// CB 96: RES 2,(HL)			16 cycles
type opCb96 struct {
	MultiStepsOp
}

func (op *opCb96) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 &= ^uint8(1<<2)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB 97: RES 2,A			8 cycles
type opCb97 struct {
	SingleStepOp
}

func (op *opCb97) Execute(c *CPU) (done bool) {
	c.A &= ^uint8(1<<2)
	return true
}

// CB 98: RES 3,B			8 cycles
type opCb98 struct {
	SingleStepOp
}

func (op *opCb98) Execute(c *CPU) (done bool) {
	c.B &= ^uint8(1<<3)
	return true
}

// CB 99: RES 3,C			8 cycles
type opCb99 struct {
	SingleStepOp
}

func (op *opCb99) Execute(c *CPU) (done bool) {
	c.C &= ^uint8(1<<3)
	return true
}

// CB 9A: RES 3,D			8 cycles
type opCb9a struct {
	SingleStepOp
}

func (op *opCb9a) Execute(c *CPU) (done bool) {
	c.D &= ^uint8(1<<3)
	return true
}

// CB 9B: RES 3,E			8 cycles
type opCb9b struct {
	SingleStepOp
}

func (op *opCb9b) Execute(c *CPU) (done bool) {
	c.E &= ^uint8(1<<3)
	return true
}

// CB 9C: RES 3,H			8 cycles
type opCb9c struct {
	SingleStepOp
}

func (op *opCb9c) Execute(c *CPU) (done bool) {
	c.H &= ^uint8(1<<3)
	return true
}

// CB 9D: RES 3,L			8 cycles
type opCb9d struct {
	SingleStepOp
}

func (op *opCb9d) Execute(c *CPU) (done bool) {
	c.L &= ^uint8(1<<3)
	return true
}

// CB 9E: RES 3,(HL)			16 cycles
type opCb9e struct {
	MultiStepsOp
}

func (op *opCb9e) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 &= ^uint8(1<<3)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB 9F: RES 3,A			8 cycles
type opCb9f struct {
	SingleStepOp
}

func (op *opCb9f) Execute(c *CPU) (done bool) {
	c.A &= ^uint8(1<<3)
	return true
}

// CB A0: RES 4,B			8 cycles
type opCbA0 struct {
	SingleStepOp
}

func (op *opCbA0) Execute(c *CPU) (done bool) {
	c.B &= ^uint8(1<<4)
	return true
}

// CB A1: RES 4,C			8 cycles
type opCbA1 struct {
	SingleStepOp
}

func (op *opCbA1) Execute(c *CPU) (done bool) {
	c.C &= ^uint8(1<<4)
	return true
}

// CB A2: RES 4,D			8 cycles
type opCbA2 struct {
	SingleStepOp
}

func (op *opCbA2) Execute(c *CPU) (done bool) {
	c.D &= ^uint8(1<<4)
	return true
}

// CB A3: RES 4,E			8 cycles
type opCbA3 struct {
	SingleStepOp
}

func (op *opCbA3) Execute(c *CPU) (done bool) {
	c.E &= ^uint8(1<<4)
	return true
}

// CB A4: RES 4,H			8 cycles
type opCbA4 struct {
	SingleStepOp
}

func (op *opCbA4) Execute(c *CPU) (done bool) {
	c.H &= ^uint8(1<<4)
	return true
}

// CB A5: RES 4,L			8 cycles
type opCbA5 struct {
	SingleStepOp
}

func (op *opCbA5) Execute(c *CPU) (done bool) {
	c.L &= ^uint8(1<<4)
	return true
}

// CB A6: RES 4,(HL)			16 cycles
type opCbA6 struct {
	MultiStepsOp
}

func (op *opCbA6) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 &= ^uint8(1<<4)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB A7: RES 4,A			8 cycles
type opCbA7 struct {
	SingleStepOp
}

func (op *opCbA7) Execute(c *CPU) (done bool) {
	c.A &= ^uint8(1<<4)
	return true
}

// CB A8: RES 5,B			8 cycles
type opCbA8 struct {
	SingleStepOp
}

func (op *opCbA8) Execute(c *CPU) (done bool) {
	c.B &= ^uint8(1<<5)
	return true
}

// CB A9: RES 5,C			8 cycles
type opCbA9 struct {
	SingleStepOp
}

func (op *opCbA9) Execute(c *CPU) (done bool) {
	c.C &= ^uint8(1<<5)
	return true
}

// CB AA: RES 5,D			8 cycles
type opCbAa struct {
	SingleStepOp
}

func (op *opCbAa) Execute(c *CPU) (done bool) {
	c.D &= ^uint8(1<<5)
	return true
}

// CB AB: RES 5,E			8 cycles
type opCbAb struct {
	SingleStepOp
}

func (op *opCbAb) Execute(c *CPU) (done bool) {
	c.E &= ^uint8(1<<5)
	return true
}

// CB AC: RES 5,H			8 cycles
type opCbAc struct {
	SingleStepOp
}

func (op *opCbAc) Execute(c *CPU) (done bool) {
	c.H &= ^uint8(1<<5)
	return true
}

// CB AD: RES 5,L			8 cycles
type opCbAd struct {
	SingleStepOp
}

func (op *opCbAd) Execute(c *CPU) (done bool) {
	c.L &= ^uint8(1<<5)
	return true
}

// CB AE: RES 5,(HL)			16 cycles
type opCbAe struct {
	MultiStepsOp
}

func (op *opCbAe) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 &= ^uint8(1<<5)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB AF: RES 5,A			8 cycles
type opCbAf struct {
	SingleStepOp
}

func (op *opCbAf) Execute(c *CPU) (done bool) {
	c.A &= ^uint8(1<<5)
	return true
}

// CB B0: RES 6,B			8 cycles
type opCbB0 struct {
	SingleStepOp
}

func (op *opCbB0) Execute(c *CPU) (done bool) {
	c.B &= ^uint8(1<<6)
	return true
}

// CB B1: RES 6,C			8 cycles
type opCbB1 struct {
	SingleStepOp
}

func (op *opCbB1) Execute(c *CPU) (done bool) {
	c.C &= ^uint8(1<<6)
	return true
}

// CB B2: RES 6,D			8 cycles
type opCbB2 struct {
	SingleStepOp
}

func (op *opCbB2) Execute(c *CPU) (done bool) {
	c.D &= ^uint8(1<<6)
	return true
}

// CB B3: RES 6,E			8 cycles
type opCbB3 struct {
	SingleStepOp
}

func (op *opCbB3) Execute(c *CPU) (done bool) {
	c.E &= ^uint8(1<<6)
	return true
}

// CB B4: RES 6,H			8 cycles
type opCbB4 struct {
	SingleStepOp
}

func (op *opCbB4) Execute(c *CPU) (done bool) {
	c.H &= ^uint8(1<<6)
	return true
}

// CB B5: RES 6,L			8 cycles
type opCbB5 struct {
	SingleStepOp
}

func (op *opCbB5) Execute(c *CPU) (done bool) {
	c.L &= ^uint8(1<<6)
	return true
}

// CB B6: RES 6,(HL)			16 cycles
type opCbB6 struct {
	MultiStepsOp
}

func (op *opCbB6) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 &= ^uint8(1<<6)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB B7: RES 6,A			8 cycles
type opCbB7 struct {
	SingleStepOp
}

func (op *opCbB7) Execute(c *CPU) (done bool) {
	c.A &= ^uint8(1<<6)
	return true
}

// CB B8: RES 7,B			8 cycles
type opCbB8 struct {
	SingleStepOp
}

func (op *opCbB8) Execute(c *CPU) (done bool) {
	c.B &= ^uint8(1<<7)
	return true
}

// CB B9: RES 7,C			8 cycles
type opCbB9 struct {
	SingleStepOp
}

func (op *opCbB9) Execute(c *CPU) (done bool) {
	c.C &= ^uint8(1<<7)
	return true
}

// CB BA: RES 7,D			8 cycles
type opCbBa struct {
	SingleStepOp
}

func (op *opCbBa) Execute(c *CPU) (done bool) {
	c.D &= ^uint8(1<<7)
	return true
}

// CB BB: RES 7,E			8 cycles
type opCbBb struct {
	SingleStepOp
}

func (op *opCbBb) Execute(c *CPU) (done bool) {
	c.E &= ^uint8(1<<7)
	return true
}

// CB BC: RES 7,H			8 cycles
type opCbBc struct {
	SingleStepOp
}

func (op *opCbBc) Execute(c *CPU) (done bool) {
	c.H &= ^uint8(1<<7)
	return true
}

// CB BD: RES 7,L			8 cycles
type opCbBd struct {
	SingleStepOp
}

func (op *opCbBd) Execute(c *CPU) (done bool) {
	c.L &= ^uint8(1<<7)
	return true
}

// CB BE: RES 7,(HL)			16 cycles
type opCbBe struct {
	MultiStepsOp
}

func (op *opCbBe) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 &= ^uint8(1<<7)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB BF: RES 7,A			8 cycles
type opCbBf struct {
	SingleStepOp
}

func (op *opCbBf) Execute(c *CPU) (done bool) {
	c.A &= ^uint8(1<<7)
	return true
}

// CB C0: SET 0,B			8 cycles
type opCbC0 struct {
	SingleStepOp
}

func (op *opCbC0) Execute(c *CPU) (done bool) {
	c.B |= uint8(1<<0)
	return true
}

// CB C1: SET 0,C			8 cycles
type opCbC1 struct {
	SingleStepOp
}

func (op *opCbC1) Execute(c *CPU) (done bool) {
	c.C |= uint8(1<<0)
	return true
}

// CB C2: SET 0,D			8 cycles
type opCbC2 struct {
	SingleStepOp
}

func (op *opCbC2) Execute(c *CPU) (done bool) {
	c.D |= uint8(1<<0)
	return true
}

// CB C3: SET 0,E			8 cycles
type opCbC3 struct {
	SingleStepOp
}

func (op *opCbC3) Execute(c *CPU) (done bool) {
	c.E |= uint8(1<<0)
	return true
}

// CB C4: SET 0,H			8 cycles
type opCbC4 struct {
	SingleStepOp
}

func (op *opCbC4) Execute(c *CPU) (done bool) {
	c.H |= uint8(1<<0)
	return true
}

// CB C5: SET 0,L			8 cycles
type opCbC5 struct {
	SingleStepOp
}

func (op *opCbC5) Execute(c *CPU) (done bool) {
	c.L |= uint8(1<<0)
	return true
}

// CB C6: SET 0,(HL)			16 cycles
type opCbC6 struct {
	MultiStepsOp
}

func (op *opCbC6) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 |= uint8(1<<0)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB C7: SET 0,A			8 cycles
type opCbC7 struct {
	SingleStepOp
}

func (op *opCbC7) Execute(c *CPU) (done bool) {
	c.A |= uint8(1<<0)
	return true
}

// CB C8: SET 1,B			8 cycles
type opCbC8 struct {
	SingleStepOp
}

func (op *opCbC8) Execute(c *CPU) (done bool) {
	c.B |= uint8(1<<1)
	return true
}

// CB C9: SET 1,C			8 cycles
type opCbC9 struct {
	SingleStepOp
}

func (op *opCbC9) Execute(c *CPU) (done bool) {
	c.C |= uint8(1<<1)
	return true
}

// CB CA: SET 1,D			8 cycles
type opCbCa struct {
	SingleStepOp
}

func (op *opCbCa) Execute(c *CPU) (done bool) {
	c.D |= uint8(1<<1)
	return true
}

// CB CB: SET 1,E			8 cycles
type opCbCb struct {
	SingleStepOp
}

func (op *opCbCb) Execute(c *CPU) (done bool) {
	c.E |= uint8(1<<1)
	return true
}

// CB CC: SET 1,H			8 cycles
type opCbCc struct {
	SingleStepOp
}

func (op *opCbCc) Execute(c *CPU) (done bool) {
	c.H |= uint8(1<<1)
	return true
}

// CB CD: SET 1,L			8 cycles
type opCbCd struct {
	SingleStepOp
}

func (op *opCbCd) Execute(c *CPU) (done bool) {
	c.L |= uint8(1<<1)
	return true
}

// CB CE: SET 1,(HL)			16 cycles
type opCbCe struct {
	MultiStepsOp
}

func (op *opCbCe) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 |= uint8(1<<1)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB CF: SET 1,A			8 cycles
type opCbCf struct {
	SingleStepOp
}

func (op *opCbCf) Execute(c *CPU) (done bool) {
	c.A |= uint8(1<<1)
	return true
}

// CB D0: SET 2,B			8 cycles
type opCbD0 struct {
	SingleStepOp
}

func (op *opCbD0) Execute(c *CPU) (done bool) {
	c.B |= uint8(1<<2)
	return true
}

// CB D1: SET 2,C			8 cycles
type opCbD1 struct {
	SingleStepOp
}

func (op *opCbD1) Execute(c *CPU) (done bool) {
	c.C |= uint8(1<<2)
	return true
}

// CB D2: SET 2,D			8 cycles
type opCbD2 struct {
	SingleStepOp
}

func (op *opCbD2) Execute(c *CPU) (done bool) {
	c.D |= uint8(1<<2)
	return true
}

// CB D3: SET 2,E			8 cycles
type opCbD3 struct {
	SingleStepOp
}

func (op *opCbD3) Execute(c *CPU) (done bool) {
	c.E |= uint8(1<<2)
	return true
}

// CB D4: SET 2,H			8 cycles
type opCbD4 struct {
	SingleStepOp
}

func (op *opCbD4) Execute(c *CPU) (done bool) {
	c.H |= uint8(1<<2)
	return true
}

// CB D5: SET 2,L			8 cycles
type opCbD5 struct {
	SingleStepOp
}

func (op *opCbD5) Execute(c *CPU) (done bool) {
	c.L |= uint8(1<<2)
	return true
}

// CB D6: SET 2,(HL)			16 cycles
type opCbD6 struct {
	MultiStepsOp
}

func (op *opCbD6) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 |= uint8(1<<2)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB D7: SET 2,A			8 cycles
type opCbD7 struct {
	SingleStepOp
}

func (op *opCbD7) Execute(c *CPU) (done bool) {
	c.A |= uint8(1<<2)
	return true
}

// CB D8: SET 3,B			8 cycles
type opCbD8 struct {
	SingleStepOp
}

func (op *opCbD8) Execute(c *CPU) (done bool) {
	c.B |= uint8(1<<3)
	return true
}

// CB D9: SET 3,C			8 cycles
type opCbD9 struct {
	SingleStepOp
}

func (op *opCbD9) Execute(c *CPU) (done bool) {
	c.C |= uint8(1<<3)
	return true
}

// CB DA: SET 3,D			8 cycles
type opCbDa struct {
	SingleStepOp
}

func (op *opCbDa) Execute(c *CPU) (done bool) {
	c.D |= uint8(1<<3)
	return true
}

// CB DB: SET 3,E			8 cycles
type opCbDb struct {
	SingleStepOp
}

func (op *opCbDb) Execute(c *CPU) (done bool) {
	c.E |= uint8(1<<3)
	return true
}

// CB DC: SET 3,H			8 cycles
type opCbDc struct {
	SingleStepOp
}

func (op *opCbDc) Execute(c *CPU) (done bool) {
	c.H |= uint8(1<<3)
	return true
}

// CB DD: SET 3,L			8 cycles
type opCbDd struct {
	SingleStepOp
}

func (op *opCbDd) Execute(c *CPU) (done bool) {
	c.L |= uint8(1<<3)
	return true
}

// CB DE: SET 3,(HL)			16 cycles
type opCbDe struct {
	MultiStepsOp
}

func (op *opCbDe) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 |= uint8(1<<3)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB DF: SET 3,A			8 cycles
type opCbDf struct {
	SingleStepOp
}

func (op *opCbDf) Execute(c *CPU) (done bool) {
	c.A |= uint8(1<<3)
	return true
}

// CB E0: SET 4,B			8 cycles
type opCbE0 struct {
	SingleStepOp
}

func (op *opCbE0) Execute(c *CPU) (done bool) {
	c.B |= uint8(1<<4)
	return true
}

// CB E1: SET 4,C			8 cycles
type opCbE1 struct {
	SingleStepOp
}

func (op *opCbE1) Execute(c *CPU) (done bool) {
	c.C |= uint8(1<<4)
	return true
}

// CB E2: SET 4,D			8 cycles
type opCbE2 struct {
	SingleStepOp
}

func (op *opCbE2) Execute(c *CPU) (done bool) {
	c.D |= uint8(1<<4)
	return true
}

// CB E3: SET 4,E			8 cycles
type opCbE3 struct {
	SingleStepOp
}

func (op *opCbE3) Execute(c *CPU) (done bool) {
	c.E |= uint8(1<<4)
	return true
}

// CB E4: SET 4,H			8 cycles
type opCbE4 struct {
	SingleStepOp
}

func (op *opCbE4) Execute(c *CPU) (done bool) {
	c.H |= uint8(1<<4)
	return true
}

// CB E5: SET 4,L			8 cycles
type opCbE5 struct {
	SingleStepOp
}

func (op *opCbE5) Execute(c *CPU) (done bool) {
	c.L |= uint8(1<<4)
	return true
}

// CB E6: SET 4,(HL)			16 cycles
type opCbE6 struct {
	MultiStepsOp
}

func (op *opCbE6) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 |= uint8(1<<4)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB E7: SET 4,A			8 cycles
type opCbE7 struct {
	SingleStepOp
}

func (op *opCbE7) Execute(c *CPU) (done bool) {
	c.A |= uint8(1<<4)
	return true
}

// CB E8: SET 5,B			8 cycles
type opCbE8 struct {
	SingleStepOp
}

func (op *opCbE8) Execute(c *CPU) (done bool) {
	c.B |= uint8(1<<5)
	return true
}

// CB E9: SET 5,C			8 cycles
type opCbE9 struct {
	SingleStepOp
}

func (op *opCbE9) Execute(c *CPU) (done bool) {
	c.C |= uint8(1<<5)
	return true
}

// CB EA: SET 5,D			8 cycles
type opCbEa struct {
	SingleStepOp
}

func (op *opCbEa) Execute(c *CPU) (done bool) {
	c.D |= uint8(1<<5)
	return true
}

// CB EB: SET 5,E			8 cycles
type opCbEb struct {
	SingleStepOp
}

func (op *opCbEb) Execute(c *CPU) (done bool) {
	c.E |= uint8(1<<5)
	return true
}

// CB EC: SET 5,H			8 cycles
type opCbEc struct {
	SingleStepOp
}

func (op *opCbEc) Execute(c *CPU) (done bool) {
	c.H |= uint8(1<<5)
	return true
}

// CB ED: SET 5,L			8 cycles
type opCbEd struct {
	SingleStepOp
}

func (op *opCbEd) Execute(c *CPU) (done bool) {
	c.L |= uint8(1<<5)
	return true
}

// CB EE: SET 5,(HL)			16 cycles
type opCbEe struct {
	MultiStepsOp
}

func (op *opCbEe) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 |= uint8(1<<5)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB EF: SET 5,A			8 cycles
type opCbEf struct {
	SingleStepOp
}

func (op *opCbEf) Execute(c *CPU) (done bool) {
	c.A |= uint8(1<<5)
	return true
}

// CB F0: SET 6,B			8 cycles
type opCbF0 struct {
	SingleStepOp
}

func (op *opCbF0) Execute(c *CPU) (done bool) {
	c.B |= uint8(1<<6)
	return true
}

// CB F1: SET 6,C			8 cycles
type opCbF1 struct {
	SingleStepOp
}

func (op *opCbF1) Execute(c *CPU) (done bool) {
	c.C |= uint8(1<<6)
	return true
}

// CB F2: SET 6,D			8 cycles
type opCbF2 struct {
	SingleStepOp
}

func (op *opCbF2) Execute(c *CPU) (done bool) {
	c.D |= uint8(1<<6)
	return true
}

// CB F3: SET 6,E			8 cycles
type opCbF3 struct {
	SingleStepOp
}

func (op *opCbF3) Execute(c *CPU) (done bool) {
	c.E |= uint8(1<<6)
	return true
}

// CB F4: SET 6,H			8 cycles
type opCbF4 struct {
	SingleStepOp
}

func (op *opCbF4) Execute(c *CPU) (done bool) {
	c.H |= uint8(1<<6)
	return true
}

// CB F5: SET 6,L			8 cycles
type opCbF5 struct {
	SingleStepOp
}

func (op *opCbF5) Execute(c *CPU) (done bool) {
	c.L |= uint8(1<<6)
	return true
}

// CB F6: SET 6,(HL)			16 cycles
type opCbF6 struct {
	MultiStepsOp
}

func (op *opCbF6) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 |= uint8(1<<6)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB F7: SET 6,A			8 cycles
type opCbF7 struct {
	SingleStepOp
}

func (op *opCbF7) Execute(c *CPU) (done bool) {
	c.A |= uint8(1<<6)
	return true
}

// CB F8: SET 7,B			8 cycles
type opCbF8 struct {
	SingleStepOp
}

func (op *opCbF8) Execute(c *CPU) (done bool) {
	c.B |= uint8(1<<7)
	return true
}

// CB F9: SET 7,C			8 cycles
type opCbF9 struct {
	SingleStepOp
}

func (op *opCbF9) Execute(c *CPU) (done bool) {
	c.C |= uint8(1<<7)
	return true
}

// CB FA: SET 7,D			8 cycles
type opCbFa struct {
	SingleStepOp
}

func (op *opCbFa) Execute(c *CPU) (done bool) {
	c.D |= uint8(1<<7)
	return true
}

// CB FB: SET 7,E			8 cycles
type opCbFb struct {
	SingleStepOp
}

func (op *opCbFb) Execute(c *CPU) (done bool) {
	c.E |= uint8(1<<7)
	return true
}

// CB FC: SET 7,H			8 cycles
type opCbFc struct {
	SingleStepOp
}

func (op *opCbFc) Execute(c *CPU) (done bool) {
	c.H |= uint8(1<<7)
	return true
}

// CB FD: SET 7,L			8 cycles
type opCbFd struct {
	SingleStepOp
}

func (op *opCbFd) Execute(c *CPU) (done bool) {
	c.L |= uint8(1<<7)
	return true
}

// CB FE: SET 7,(HL)			16 cycles
type opCbFe struct {
	MultiStepsOp
}

func (op *opCbFe) Tick() (done bool) {
	switch op.step {
	case 0:
		op.cpu.temp8 = op.cpu.MMU.Read(uint(op.cpu.HL()))
		op.step++
	case 1:
		op.cpu.temp8 |= uint8(1<<7)
		op.cpu.MMU.Write(uint(op.cpu.HL()), op.cpu.temp8)
		done = true
	}
	return
}

// CB FF: SET 7,A			8 cycles
type opCbFf struct {
	SingleStepOp
}

func (op *opCbFf) Execute(c *CPU) (done bool) {
	c.A |= uint8(1<<7)
	return true
}

