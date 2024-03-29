package memory

// DMA implementation. Source:
// [VIDEO] http://gbdev.gg8.se/wiki/articles/Video_Display#FF46_-_DMA_-_DMA_Transfer_and_Start_Address_.28R.2FW.29
// TODO: restrict CPU access to HRAM while a transfer is active.

// Initialize sub-logger for DMA transfers.
func init() {
	log.Add("dma", "DMA register and transfers")
}

// AddrDMA is the address of DMA register in DMG address space.
const AddrDMA = 0xff46

// DMA address space taking care of memory transfers between main MMU and
// OAM memory.
type DMA struct {
	DMA uint8
	MMU Addressable
	OAM Addressable

	pending   bool // For initial delay
	isActive  bool
	written   uint // Total bytes written
	src, dest uint16
}

// NewDMA returns an instance of DMA managing the actual register and memory
// transfers. Parameter is an Addressable that must span source and destination
// address spaces.
func NewDMA(mmu, oam Addressable) *DMA {
	return &DMA{MMU: mmu, OAM: oam}
}

// Contains return true if the requested address is the DMA register.
func (d *DMA) Contains(addr uint16) bool {
	return addr == AddrDMA
}

// Read returns the content of the DMA register.
func (d *DMA) Read(addr uint16) uint8 {
	return d.DMA
}

// Write sets the start address for memory transfer in the DMA register and
// initiates said transfer.
func (d *DMA) Write(addr uint16, value uint8) {
	d.DMA = value
	d.src = uint16(value) * 0x100
	d.dest = 0xfe00 // OAM RAM
	d.written = 0
	d.pending = true

	log.Sub("dma").Debugf("Request DMA transfer 0x%04x→0xfe00", d.src)
}

// Tick advances DMA transfer one step if it's active. Called every 4 CPU ticks.
func (d *DMA) Tick() {
	// Start DMA transfer with a 1-cycle delay according to [GEKKIO].
	if d.pending {
		d.isActive = true
		d.pending = false
		return
	}

	if !d.isActive {
		return
	}

	// Feeling like tracking down DMA timings. Might delete later.
	if d.written == 0 {
		log.Sub("dma").Debugf("Start DMA transfer")
	}

	// A DMA transfer takes 160 CPU cycles (160×4 ticks).
	if d.written < 160 {
		d.OAM.Write(d.dest, d.MMU.Read(d.src))
		d.src++
		d.dest++
		d.written++

		return
	}

	log.Sub("dma").Debug("DMA transfer done")
	d.isActive = false
}

// DMAMemory wraps the whole address space to forbid memory access to the CPU
// while a DMA transfer is taking place.
type DMAMemory struct {
	Addressable
	DMA *DMA
}

// NewDMA returns an instance of DMA managing the actual register and memory
// transfers. Parameter is an Addressable that must span source and destination
// address spaces.
func NewDMAMemory(mmu, oam Addressable) *DMAMemory {
	return &DMAMemory{Addressable: mmu, DMA: NewDMA(mmu, oam)}
}

// Read overrides the embedded Addressable method to only allow reading from
// high RAM if a DMA transfer is currently taking place.
func (d *DMAMemory) Read(addr uint16) uint8 {
	if d.DMA.isActive && addr < 0xfe00 {
		log.Sub("dma").Desperatef("read to %x during DMA transfer", addr)
		return 0xff
	} else {
		return d.Addressable.Read(addr)
	}
}

// Write overrides the embedded Addressable method to only allow writing to high
// RAM if a DMA transfer is currently taking place.
func (d *DMAMemory) Write(addr uint16, value uint8) {
	if d.DMA.isActive && addr < 0xfe00 {
		log.Sub("dma").Desperatef("write to %x during DMA transfer", addr)
		return
	} else {
		d.Addressable.Write(addr, value)
	}
}
