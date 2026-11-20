package memory_test

import (
	"math/rand"
	"testing"

	"github.com/lazy-stripes/goholint/memory"
)

// Validate that MMU.Contains() properly returns true for all address spaces
// it covers.
func TestMMUContains(t *testing.T) {
	strNot := map[bool]string{true: "", false: " not"}

	// Generate memory zones at semi-random start addresses. Seven zones that
	// are at most 8KB should cover enough of the 64KB memory space to test.
	want := make([]bool, 0x10000)
	memoryZones := make([]*memory.Memory, 7)
	lastAddr := 0
	for i := range len(memoryZones) {
		start := lastAddr + rand.Intn(0x2000)
		size := rand.Intn(0x2000)
		lastAddr = start + size
		memoryZones[i] = memory.New(uint16(start), size)
	}

	// TODO: should I just add Start() and Size() to Addressable and directly
	// store stuff in an []Addressable slice?
	//mmu := memory.NewMMU(memoryZones)
	mmu := memory.NewEmptyMMU()

	// All addresses covered by our memory spaces should return true.
	for _, m := range memoryZones {
		mmu.Add(m)
		for i := int(m.Start); i < int(m.Start)+len(m.Bytes); i++ {
			want[i] = true
		}
	}

	for addr := range 0x10000 {
		if mmu.Contains(uint16(addr)) != want[addr] {
			t.Fatalf("Address %#04x should%s be contained in MMU", addr,
				strNot[want[addr]])
		}
	}
}

// Validate that reading and writing from an MMU instance hits the proper
// address space.
func TestMMUReadWrite(t *testing.T) {
	// Use two arrays to validate test: one for the first read, one for the
	// second read right after a write. By default they contain the value for
	// open bus read (0xff)
	wantRead1 := make([]byte, 0x10000)
	wantRead2 := make([]byte, 0x10000)
	for i := range 0x10000 {
		wantRead1[i] = 0xff
		wantRead2[i] = 0xff
	}

	// Use some read-only and some r/w spaces. Also BootROM while we are at it,
	// so we can see MMU dynamically reacts to changes in their contents'
	// Contains() method.

	// Pseudo boot ROM using a sentinel value. Update array of wanted values
	// accordingly (do not forget the register).
	bootData := make([]byte, 0x100)
	for i := range bootData {
		bootData[i] = 0xbb
		wantRead1[i] = bootData[i]
		// No wantRead2, we will have disabled the boot ROM at that point.
	}
	boot := memory.NewBoot(bootData)
	wantRead1[memory.BootAddr] = boot.Register
	wantRead2[memory.BootAddr] = boot.Register + 1

	// Pseudo ROM using a sentinel value. It should initially be shadowed by
	// the boot ROM.
	romData := make([]byte, 0x8000)
	for i := range romData {
		romData[i] = 0xdd
		if i > 0xff { // First 255 bytes are shadowed by boot ROM at first.
			wantRead1[i] = romData[i]
		}
		wantRead2[i] = romData[i]
	}
	rom := memory.NewROM(romData)

	// RAM block using a sentinel value. Those should change after write.
	ramStart := 0x8000 + rand.Intn(0x4000)
	ram := memory.NewRAM(uint16(ramStart), 0x4000)
	for i := range 0x4000 {
		// Adjust for RAM start address.
		ram.Write(uint16(ramStart+i), 0xaa)
		wantRead1[ramStart+i] = 0xaa
		wantRead2[ramStart+i] = 0xab
	}

	mmu := memory.NewMMU([]memory.Addressable{boot, rom, ram})

	// First read should return the contents of boot, rom and ram where applicable.
	for addr := range 0x10000 {
		expected := wantRead1[addr]
		value := mmu.Read(uint16(addr))
		if value != expected {
			t.Fatalf("Invalid value read from MMU at address %#04x (expected %#02x, got %#02x)",
				addr, expected, value)
		}
	}

	// Disable boot ROM now. Then do a second pass writing and reading again.
	// Still use zero so that we can see it increment (TODO: validate that you
	// can write to FF50 several times).
	boot.Write(memory.BootAddr, 0x00)

	// Try incrementing all values in memory, see which ones were properly
	// updated.
	for addr := range 0x10000 {
		mmu.Write(uint16(addr), mmu.Read(uint16(addr))+1)
		expected := wantRead2[addr]
		value := mmu.Read(uint16(addr))
		if value != expected {
			t.Fatalf("Invalid value reread from MMU at address %#04x (expected %#02x, got %#02x)",
				addr, expected, value)
		}
	}
}
