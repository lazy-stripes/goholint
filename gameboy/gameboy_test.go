package gameboy

import "testing"

func TestEmpty(t *testing.T) {
	gb := New()
	gb.MMU.Write(42, 42)
}
