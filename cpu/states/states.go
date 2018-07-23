package states

// Mere collection of constants because I was procrastinating.
const (
	FetchOpCode = iota // 0, default value for cpu.state
	FetchExtendedOpcode
	Execute
)
