package gameboy

import (
	"go.tigris.fr/gameboy/memory"
)

// GameBoy (naive) implementation. Not used yet.
type GameBoy struct {
	MMU memory.MMU
	// CPU cpu.CPU
	// LCD gpu.LCD
}
