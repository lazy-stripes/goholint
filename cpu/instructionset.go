// Generated code. See instructions.go
package cpu

// 00: NOP			4 cycles
type nop struct {
	SingleStepOp
}

func (op nop) Execute(c *CPU) (done bool) {
	return true
}

// A0: AND B
type andB struct {
	SingleStepOp
}

func (op andB) Execute(c *CPU) (done bool) {
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
type andC struct {
	SingleStepOp
}

func (op andC) Execute(c *CPU) (done bool) {
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
type andD struct {
	SingleStepOp
}

func (op andD) Execute(c *CPU) (done bool) {
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
type andE struct {
	SingleStepOp
}

func (op andE) Execute(c *CPU) (done bool) {
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
type andH struct {
	SingleStepOp
}

func (op andH) Execute(c *CPU) (done bool) {
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
type andL struct {
	SingleStepOp
}

func (op andL) Execute(c *CPU) (done bool) {
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
type andAddrHL struct {
	MultiStepsOp
}

func (op andAddrHL) Tick() (done bool) {
	op.cpu.A &= op.cpu.MMU.Read(uint(op.cpu.HL()))
	// Flags z 0 0 0
	if op.cpu.A == 0 {
		op.cpu.F = FlagZ
	} else {
		op.cpu.F = 0
	}
	return true
}
