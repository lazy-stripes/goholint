package cpu

import (
	"testing"

	"tigris.fr/gameboy/memory"
)

func TestLdRrD16(t *testing.T) {
	want := uint16(0xABCD)
	ram := memory.NewRAM(0, 2)
	ram.Write(0, 0xCD)
	ram.Write(1, 0xAB)
	cpu := New(memory.NewMMU([]memory.AddressSpace{ram}))
	ldRrD16(cpu, &cpu.B, &cpu.C)
	got := cpu.BC()
	if got != want {
		t.Errorf("LD rr,d16 failed: wanted %#x, got %#x", want, got)
	}
}
