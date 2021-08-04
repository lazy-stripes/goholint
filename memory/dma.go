package memory

// DMA implementation. Source:
// [VIDEO] http://gbdev.gg8.se/wiki/articles/Video_Display#FF46_-_DMA_-_DMA_Transfer_and_Start_Address_.28R.2FW.29

// AddrDMA is the address of DMA register in DMG address space.
const AddrDMA = 0xff46

// DMA address space taking care of memory transfers between main MMU and
// OAM memory.
type DMA struct {
	DMA uint8
	MMU Addressable

	isActive  bool
	ticks     int
	src, dest uint16
}

// NewDMA returns an instance of DMA managing the actual register and memory
// transfers. Parameter is an Addressable that must span source and destination
// address spaces.
func NewDMA(mmu Addressable) *DMA {
	return &DMA{MMU: mmu}
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
	d.ticks = 0
	d.isActive = true

	log.Sub("dma").Debugf("Start DMA transfer 0x%04x→0xfe00", d.src)
}

// Tick advances DMA transfer one step if it's active. Called every clock tick.
func (d *DMA) Tick() {
	if !d.isActive {
		return
	}
	d.ticks++

	// A DMA transfer takes 160 DMA ticks.
	if d.ticks < 160 {
		// [VIDEO] It takes 160 cycles to complete a 160-byte DMA transfer, right?
		d.MMU.Write(d.dest, d.MMU.Read(d.src))
		d.src++
		d.dest++

		return
	}
	log.Sub("dma").Debug("DMA transfer done")
	d.isActive = false
}
