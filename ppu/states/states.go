package states

// State of a PPU at a given step.
type State int

// Mere collection of constants because I was procrastinating when I did that for package cpu.
const (
	OAMSearch = iota // 0, default value for ppu.state
	PixelTransfer
	HBlank
	VBlank
)
