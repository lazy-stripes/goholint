package gameboy

import (
	"go.tigris.fr/gameboy/memory"
)

// GameBoy (naive) implementation.
type GameBoy struct {
	MMU memory.MMU
	// CPU cpu.CPU
	// LCD gpu.LCD
}
