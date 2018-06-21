package gameboy

import (
	"tigris.fr/gameboy/memory"
)

// GameBoy (naive) implementation.
type GameBoy struct {
	MMU memory.MMU
	// CPU cpu.CPU
	// LCD gpu.LCD
}

// New GameBoy instance.
func New() *GameBoy {
	return &GameBoy{memory.NewBoot(make([]uint8, 0x100), make([]uint8, 0x200))}
}
