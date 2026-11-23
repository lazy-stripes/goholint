package memory

// DMA implementation. Source:
// [VIDEO] http://gbdev.gg8.se/wiki/articles/Video_Display#FF46_-_DMA_-_DMA_Transfer_and_Start_Address_.28R.2FW.29

// FIXME: gekkio confirms that the "only HRAM is accessible during DMA" is a myth. It's more complicated than that.
// [GEKKIODMA] https://www.reddit.com/r/EmuDev/comments/5hahss/gb_readwrite_memory_during_an_oam_dma/daz9cbi/
// [BUSES] https://iceboy.a-singer.de/doc/mem_patterns.html
// TODO: start downloading and storing those sources offline, because we never know.
// TL;DR:
// external bus: $0000-$7FFF, $A000-$FDFF
// video ram bus: $8000-$9FFF

// Initialize sub-logger for DMA transfers.
func init() {
	log.Add("dma", "DMA register and transfers")
}

// AddrDMA is the address of DMA register in DMG address space.
const AddrDMA = 0xff46

// Bus identifiers returned by TargetBus. See TargetBus documentation for
// details.
type Bus int

const (
	ExternalBus Bus = iota // 0x0000-0x7fff, 0xa000-0xfdff
	VideoBus               // 0x8000-0x9fff, 0xfe00-0xfeff
	InternalBus            // 0xff00-0xffff
)

// TargetBus returns the identifier of the bus that contains the given address.
// We define the following buses based on available documentation (not 100%
// accurate but hopefully sufficient for our needs):
//
//   - External bus: 0x0000-0x7fff (ROM) and 0xa000-0xfdff (RAM + "echo" RAM)
//   - Video bus: 0x8000-0x9fff (VRAM) and 0xfe00-0xfeff (OAM)
//   - Internal bus: all other addresses (0xff00-0xffff, always reachable)
func TargetBus(addr uint16) Bus {
	switch {
	case addr < 0x8000:
		return ExternalBus
	case addr >= 0x8000 && addr < 0xa000:
		return VideoBus
	case addr >= 0xa000 && addr < 0xfe00:
		return ExternalBus
	case addr >= 0xfe00 && addr < 0xfea0:
		return VideoBus
	default:
		return InternalBus
	}
}

// DMA address space taking care of memory transfers between main MMU and
// OAM memory.
type DMA struct {
	DMA uint8
	MMU Addressable
	OAM Addressable

	pending   bool // For initial delay
	running   bool
	written   uint // Total bytes written
	src, dest uint16

	srcBus Bus // For blocking reads/writes
}

// NewDMA returns an instance of DMA managing the actual register and memory
// transfers. Parameter is an Addressable that must span source (0x0000-0xdfff)
// and destination (0xfe00-0xfe9f) address spaces.
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
	d.src = uint16(value) * 0x100 // FIXME: if value is > 0xdf, what do?
	d.srcBus = TargetBus(d.src)
	d.dest = 0xfe00 // OAM RAM
	d.written = 0
	d.pending = true

	log.Sub("dma").Debugf("Request DMA transfer 0x%04x→0xfe00", d.src)
}

// Tick advances DMA transfer one step if it's active. Called every 4 CPU ticks.
func (d *DMA) Tick() {
	// Start DMA transfer with a 1-cycle delay according to [GEKKIO].
	if d.pending {
		d.running = true
		d.pending = false
		return
	}

	if !d.running {
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
	d.running = false
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

// Read overrides the embedded Addressable method to block reads from buses that
// a DMA transfer is currently using. A DMA transfer always blocks the Video bus
// and will also block the External bus if transferring data from it.
func (d *DMAMemory) Read(addr uint16) uint8 {
	if d.DMA.running {
		bus := TargetBus(addr)
		if bus == VideoBus || (bus == d.DMA.srcBus) {
			log.Sub("dma").Desperatef("blocked read to %x during DMA transfer",
				addr)
			return 0xff
		}
	}
	return d.Addressable.Read(addr)
}

// Write overrides the embedded Addressable method to block writes from buses that
// a DMA transfer is currently using. A DMA transfer always blocks the Video bus
// and will also block the External bus if transferring data from it.
func (d *DMAMemory) Write(addr uint16, value uint8) {
	if d.DMA.running {
		bus := TargetBus(addr)
		if bus == VideoBus || (bus == d.DMA.srcBus) {
			log.Sub("dma").Desperatef("blocked write to %x during DMA transfer",
				addr)
			return
		}
	}
	d.Addressable.Write(addr, value)
}
