package cpu

// Instruction templates are used to generate state-machine structs for each opcode.
//go:generate go run instructions/make.go -o instructionset.go

// Instructions. Each takes a CPU pointer and will modify its internal state.
// Source: http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html

// Number of cycles reflects the information given in resource linked above.
// Note that reading the instruction byte itself takes 4 cycles (8 for CB xx instructions.)
// Those 4 cycles are included in the count indicated in the helper's comment.
// Each subsequent Operation pushed on the CPU will take an additional 4 cycles.

// An Instruction to be executed by a CPU, implemented as a state machine executing a new step every 4 cycles.
type Instruction interface {
	Execute(c *CPU) (done bool)
	Tick() (done bool)
}

// SingleStepOp is an instruction executed within the 4 (or 8) cycles needed to read the opcode.
// To be embedded in actual instructions that don't need to implement Tick(), but
// need to implement Execute(c *CPU) and make sure it returns false. Kinda meh but ought to work.
type SingleStepOp struct{}

// Tick is a mere placeholder for derived single-instruction types and panics if called.
func (op *SingleStepOp) Tick() (done bool) {
	panic("Tick() called on instruction supposed to complete within Execute()")
}

// MultiStepsOp is an instruction needing more than the fetching cycle(s) to complete.
// It stores the CPU reference passed to Execute() and then expects derived types
// to implement Tick().
type MultiStepsOp struct {
	cpu  *CPU
	step uint // XXX: do we need an enum?
}

// Execute keeps the passed CPU pointer and resets step number used for the state machine.
func (op *MultiStepsOp) Execute(c *CPU) (done bool) {
	op.cpu = c
	op.step = 0
	return false
}

// Tick executes this instruction's next step and returns true as long as there are
// further steps to take. This is a placeholder to be overridden.
func (op *MultiStepsOp) Tick() (done bool) {
	panic("Tick() hasn't been implemented for this Instruction!")
}
