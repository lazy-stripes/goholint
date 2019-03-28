package states

// Mere collection of constants because I was procrastinating. Also a bitfield for easy combination.
const (
	FetchOpCode = 1 << iota
	FetchExtendedOpcode
	Execute
	Halted
	Stopped
	InterruptWait0
	InterruptWait1
	InterruptPushPCHigh
	InterruptPushPCLow
	InterruptCall

	// Useful combinations
	Interruptible     = FetchOpCode | Halted | Stopped
	HandlingInterrupt = InterruptWait0 | InterruptWait1 | InterruptPushPCHigh | InterruptPushPCLow | InterruptCall
)
