package memory

// Memory Bank Controllers. Source:
// [PANMBC] https://gbdev.io/pandocs/MBC1.html

// Initialize sub-logger for MBC details.
func init() {
	log.Add("mbc", "MBC chip details")
	log.Add("mbc/read", "MBC ROM chip reads (Desperate level only)")
	log.Add("mbc/write", "MBC ROM chip writes (Desperate level only)")
}

// 6000-7FFF - ROM/RAM Mode Select
const (
	ROMBanking = 0x00
	RAMBanking = 0x01
)

const AddrExternalRAM = 0xa000

// MBC1 (max 2MByte ROM and/or 32KByte RAM)
type MBC1 struct {
	ROM *BankedROM // Complete ROM (will be addressed according to ROMBank)
	RAM *BankedRAM // Optional RAM (up to 32KB)

	RAMEnabled bool // 0000-1FFF - RAM Enable

	BankLow  uint8 // 2000-3FFF - ROM Bank Number (bits 0-4)
	BankHigh uint8 // 4000-5FFF - RAM Bank Number / ROM Bank Number (bits 5-6)

	// 6000-7FFF - ROM/RAM Mode Select
	// 00h = ROM Banking Mode (up to 8KByte RAM, 2MByte ROM) (default)
	// 01h = RAM Banking Mode (up to 32KByte RAM, 512KByte ROM)
	BankingMode uint8

	// If true, save RAM values. (FIXME: when?)
	battery bool

	// Max number of ROM/RAM banks.
	romBanks uint8 // MBC1 has max 2 MByte ROM (128 banks)
	ramBanks uint8

	// Pre-computed mask for ROM bank number. Set to (romBanks - 1).
	// E.g. 32 banks need 5 bits for addressing, so mask is 31 (0x11111).
	romBanksMask uint
}

// NewMBC1 creates an address space emulating a cartridge with an MBC1 chip.
// Takes an instance of ROM because to know which kind of chip it uses, we had
// to read it beforehand anyway.
func NewMBC1(rom *ROM, romBanks uint8, ramBanks uint8, battery bool, savePath string) *MBC1 {
	ramSize := uint16(ramBanks) * 0x2000
	ram := NewBankedRAM(AddrExternalRAM, ramSize)

	// If the cartridge has a battery-backed RAM, restore it here.
	if battery {
		if err := ram.Load(savePath); err != nil {
			log.Warning(err.Error())
		}
	}

	return &MBC1{
		ROM:     NewBankedROM(rom.Bytes),
		RAM:     ram,
		BankLow: 1,

		romBanks:     romBanks,
		ramBanks:     ramBanks,
		romBanksMask: uint(romBanks - 1),
		battery:      battery,
	}
}

// Contains returns true if the requested address is anywhere in ROM or RAM.
func (m *MBC1) Contains(addr uint16) bool {
	switch {
	case /* addr >= 0x0000 && */ addr <= 0x7fff:
		return true
	case addr >= 0xa000 && addr <= 0xbfff:
		return true
	default:
		return false
	}
}

// Read returns the byte at requested address in current ROM or RAM bank.
func (m *MBC1) Read(addr uint16) uint8 {
	switch {
	case /*  addr >= 0x0000 && */ addr <= 0x7fff:
		return m.ROM.Read(addr)

	case addr >= 0xa000 && addr <= 0xbfff:
		if m.RAMEnabled {
			value := m.RAM.Read(addr)
			log.Sub("mbc/read").Desperatef("RAM[%04x]: 0x%02x", addr, value)
			return value
		}
		log.Sub("mbc/read").Desperatef("RAM disabled, read ignored")
		fallthrough // Return "ff" as our undefined value.

	default:
		return 0xff
	}
}

// recomputeBanks recomputes ROM and RAM bank numbers after writes to the MBC1.
func (m *MBC1) recomputeBanks() {
	if m.BankingMode == RAMBanking {
		if m.ramBanks > 0 && m.BankHigh < m.ramBanks {
			m.RAM.Bank = uint(m.BankHigh)
		}

		m.ROM.Bank0 = uint(m.BankHigh<<5) & m.romBanksMask
	} else {
		m.RAM.Bank = 0
		m.ROM.Bank0 = 0
	}

	// 4000-7fff range uses all 7 bank number bits regardless of mode.
	m.ROM.Bank1 = uint((m.BankHigh<<5)+m.BankLow) & m.romBanksMask
}

// Write value to RAM, enable RAM or select ROM/RAM banks.
func (m *MBC1) Write(addr uint16, value uint8) {
	switch {
	// 0000-1FFF - RAM Enable
	case addr <= 0x1fff:
		m.RAMEnabled = (value&0x0f == 0x0a)
		log.Sub("mbc/write").Desperatef("RAMEnabled=%t", m.RAMEnabled)

	// 2000-3FFF - ROM Bank Number
	case addr >= 0x2000 && addr <= 0x3fff:

		m.BankLow = value & 0x1f
		if m.BankLow == 0 {
			m.BankLow = 1
		}
		m.recomputeBanks()
		log.Sub("mbc/write").Debugf("BankLow=0x%02x", value&0x1f)

	// 4000-5FFF - RAM Bank Number - or - Upper Bits of ROM Bank Number
	case addr >= 0x4000 && addr <= 0x5fff:
		m.BankHigh = value & 3
		m.recomputeBanks()
		log.Sub("mbc/write").Debugf("BankHigh=0x%02x", value&3)

	// 6000-7FFF - ROM/RAM Mode Select
	case addr >= 0x6000 && addr <= 0x7fff:
		log.Sub("mbc/write").Debugf("Banking Mode 0x%02x selected", value)
		m.BankingMode = value & 1
		m.recomputeBanks()

	// A000-BFFF - RAM Bank 00-03, if any
	case addr >= 0xa000 && addr <= 0xbfff:
		if !m.RAMEnabled {
			log.Sub("mbc/write").Desperatef("RAM not enabled, write to 0x%04x ignored.",
				addr)
			return
		}
		// Proper banks should have been set.
		m.RAM.Write(addr, value)
		log.Sub("mbc/write").Desperatef("RAM[%04x]=0x%02x", addr, value)

		// Write save file on enable. FIXME: Buffer it.
		if m.battery && m.RAMEnabled {
			if err := m.RAM.Save(); err != nil {
				log.Sub("mbc").Warningf("save RAM failed (%s)", err)
			}
		}
	}
}
