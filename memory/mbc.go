package memory

import "go.tigris.fr/gameboy/logger"

// Memory Bank Controllers. Source:
// [PANMBC] http://bgb.bircd.org/pandocs.htm#memorybankcontrollers

// MBC1 (max 2MByte ROM and/or 32KByte RAM)
type MBC1 struct {
	*ROM       // Complete ROM (will be addressed according to ROMBank)
	*RAM       // Optional RAM (up to 32KB)
	ROMBank    uint8
	RAMBank    uint8
	RAMEnabled bool

	// True in ROM Banking Mode (up to 8KByte RAM, 2MByte ROM) (default).
	// False in RAM Banking Mode (up to 32KByte RAM, 512KByte ROM).
	ROMBankingMode bool
}

// NewMBC1 creates an address space emulating a cartridge with an MBC1 chip.
// Takes an instance of ROM because to know which kind of chip it uses, we need
// to read it beforehand anyway.
// TODO: load RAM state from external file too.
func NewMBC1(rom *ROM) *MBC1 {
	return &MBC1{ROM: rom, RAM: NewRAM(0, 0x8000), ROMBank: 1}
}

// Contains returns true if the requested address is anywhere in ROM or RAM.
func (m *MBC1) Contains(addr uint16) bool {
	switch {
	case addr >= 0x0000 && addr <= 0x7fff:
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
	case addr >= 0x0000 && addr <= 0x3fff:
		return m.ROM.Read(addr)
	case addr >= 0x4000 && addr <= 0x7fff:
		logger.Printf("mbc/read", "Read ROM at %x.", uint(m.ROMBank)*0x4000+uint(addr-0x4000))
		return m.ROM.read(uint(m.ROMBank)*0x4000 + uint(addr-0x4000))
	case addr >= 0xa000 && addr <= 0xbfff:
		return m.RAM.Read(uint16(m.RAMBank)*0x2000 + uint16(addr-0xa000))
	default:
		return 0xff
	}
}

// Write value to RAM, enable RAM or select ROM/RAM banks.
func (m *MBC1) Write(addr uint16, value uint8) {
	switch {
	case addr >= 0x0000 && addr <= 0x1fff:
		m.RAMEnabled = (value&0x0a == 0x0a)
	case addr >= 0x2000 && addr <= 0x3fff:
		if value&0x1f == 0 { // Bank 0 (or over 0x1f) is not selectable.
			value = 1
		}
		m.ROMBank = m.ROMBank&0x60 | value
		logger.Printf("mbc/write", "ROM Bank 0x%02x selected", m.ROMBank)
	case addr >= 0x4000 && addr <= 0x5fff:
		if m.ROMBankingMode {
			m.ROMBank = m.ROMBank&0x1f | (value&3)<<5
			logger.Printf("mbc/write", "ROM Bank 0x%02x selected", m.ROMBank)
		} else {
			m.RAMBank = value & 3
		}
	case addr >= 0x6000 && addr <= 0x7fff:
		m.ROMBankingMode = (value == 0)
	case addr >= 0xa000 && addr <= 0xbfff:
		if !m.RAMEnabled {
			logger.Printf("mbc", "RAM not enabled, write to 0x%04x ignored.",
				addr)
			return
		}
		m.RAM.Write(uint16(m.RAMBank)*0x2000+uint16(addr-0xa000), value)
	}
}