package memory

// Memory Bank Controllers. Source:
// [PANMBC] https://gbdev.io/pandocs/MBC1.html
// [PANMBC2] https://gbdev.io/pandocs/MBC2.html

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
	// Avenging Spirit uses MBC1 with no RAM, but still enables it and attempts
	// writing to it, so we still need to check whether there is RAM at all.
	ramSize := uint16(ramBanks) * 0x2000
	ram := NewBankedRAM(AddrExternalRAM, ramSize)

	// If the cartridge has a battery-backed RAM, restore it here.
	if battery && (ramBanks > 0) {
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
		if m.RAMEnabled && (m.ramBanks > 0) {
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
		if m.BankHigh < m.ramBanks {
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
		if !m.RAMEnabled || (m.ramBanks == 0) {
			log.Sub("mbc/write").Desperatef("RAM not enabled, write to 0x%04x ignored.", addr)
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

const MBC2RAMSize = 512

// MBC2 (max 256 KiB ROM and 512×4 bits RAM)
type MBC2 struct {
	ROM *BankedROM // Complete ROM (will be addressed according to ROMBank)
	RAM *BankedRAM // Optional RAM (up to 32KB)

	RAMEnabled bool // 0000-1FFF - RAM Enable

	// If true, save RAM values. (FIXME: when?)
	battery bool

	// Max number of ROM/RAM banks.
	romBanks uint8 // MBC2 has max 256 KByte ROM (16 banks)
}

// NewMBC2 creates an address space emulating a cartridge with an MBC2 chip.
// Takes an instance of ROM because to know which kind of chip it uses, we had
// to read it beforehand anyway.
func NewMBC2(rom *ROM, romBanks uint8, battery bool, savePath string) *MBC2 {
	ram := NewBankedRAM(AddrExternalRAM, MBC2RAMSize)

	// If the cartridge has a battery-backed RAM, restore it here.
	if battery {
		if err := ram.Load(savePath); err != nil {
			log.Warning(err.Error())
		}
	}

	m := &MBC2{
		ROM: NewBankedROM(rom.Bytes),
		RAM: ram,

		romBanks: romBanks,
		battery:  battery,
	}

	// [PANMBC2] The ROM bank is set to 1 by default.
	m.ROM.Bank1 = 1

	return m
}

// Contains returns true if the requested address is anywhere in ROM or RAM.
func (m *MBC2) Contains(addr uint16) bool {
	return addr <= 0x7fff || (addr >= 0xa000 && addr <= 0xbfff)
}

// Read returns the byte at requested address in current ROM bank or RAM. RAM
// values are stored on 4 bits, the high nibble will be zeroed out. RAM is also
// only 512 bytes and wraps around.
func (m *MBC2) Read(addr uint16) uint8 {
	switch {
	case addr <= 0x7fff:
		return m.ROM.Read(addr)

	case addr >= 0xa000 && addr <= 0xbfff:
		if m.RAMEnabled {
			// Wrap around every 512 bytes (only 9 lowest bits of address used).
			wrappedAddr := AddrExternalRAM + (addr & (MBC2RAMSize - 1))
			value := m.RAM.Read(wrappedAddr) & 0x0f
			log.Sub("mbc/read").Desperatef("RAM[%04x]: 0x%02x", addr, value)
			return value
		}
		log.Sub("mbc/read").Desperatef("RAM disabled, read ignored")
		fallthrough // Return "ff" as our undefined value.

	default:
		return 0xff
	}
}

// Write value to RAM, enable RAM or select ROM/RAM banks.
func (m *MBC2) Write(addr uint16, value uint8) {
	switch {
	// 0000-3FFF - RAM Enable, ROM Bank Number
	case addr <= 0x3fff:
		// If bit 8 of address is set, enable/disable RAM. Otherwise, select
		// ROM bank.
		if addr&0x0100 == 0 {
			m.RAMEnabled = (value&0x0f == 0x0a)
			log.Sub("mbc/write").Desperatef("RAMEnabled=%t", m.RAMEnabled)
		} else {
			// [PANMBC2] Specifically, the lower 4 bits of the value written to
			// this address range specify the ROM bank number. If bank 0 is
			// written, the resulting bank will be bank 1 instead.
			m.ROM.Bank1 = uint(value & 0x0f)
			if m.ROM.Bank1 == 0 {
				m.ROM.Bank1 = 1
			}
			log.Sub("mbc/write").Debugf("ROMBank=0x%02x", m.ROM.Bank1)
		}

	// A000-BFFF - RAM
	case addr >= 0xa000 && addr <= 0xbfff:
		if !m.RAMEnabled {
			log.Sub("mbc/write").Desperatef("RAM not enabled, write to 0x%04x ignored.", addr)
			return
		}
		// Wrap around every 512 bytes (only 9 lowest bits of address used).
		wrappedAddr := AddrExternalRAM + (addr & (MBC2RAMSize - 1))
		m.RAM.Write(wrappedAddr, value&0x0f)
		log.Sub("mbc/write").Desperatef("RAM[%04x]=0x%02x", addr, value)

		// Write save file on enable. FIXME: Buffer it. Put it in a common mbc base struct.
		if m.battery && m.RAMEnabled {
			if err := m.RAM.Save(); err != nil {
				log.Sub("mbc").Warningf("save RAM failed (%s)", err)
			}
		}
	}
}
