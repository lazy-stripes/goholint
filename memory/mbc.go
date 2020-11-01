package memory

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

	// If true, save RAM values. (TODO: when?)
	battery bool

	// Max number of ROM/RAM banks.
	romBanks uint16
	ramBanks uint8
}

// NewMBC1 creates an address space emulating a cartridge with an MBC1 chip.
// Takes an instance of ROM because to know which kind of chip it uses, we need
// to read it beforehand anyway.
// TODO: load RAM state from external file too.
func NewMBC1(rom *ROM, romBanks uint16, ramBanks uint8, battery bool, savePath string) *MBC1 {
	// If the cartridge has a battery-backed RAM, restore it here.
	ramSize := uint16(ramBanks) * 0x2000
	var ram *RAM
	if battery {
		ram = NewRAM(0, ramSize, savePath)
	} else {
		ram = NewEmptyRAM(0, ramSize)
	}

	return &MBC1{
		ROM:      rom,
		RAM:      ram,
		ROMBank:  1,
		romBanks: romBanks,
		ramBanks: ramBanks,
		battery:  battery,
	}
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
		// [GEEKIO] 7.2. says BANK2 is used if RAM banking mode is 1.
		if m.ROMBankingMode {
			return m.ROM.Read(addr)
		}
		return m.ROM.read(uint(m.ROMBank&0x60)*0x4000 + uint(addr))
	case addr >= 0x4000 && addr <= 0x7fff:
		log.Sub("mbc/read").Desperatef("Read ROM at %x.",
			uint(m.ROMBank)*0x4000+uint(addr-0x4000))
		return m.ROM.read(uint(m.ROMBank)*0x4000 + uint(addr-0x4000))
	case m.RAMEnabled && addr >= 0xa000 && addr <= 0xbfff:
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

		// Write save file on enable. FIXME: it does that continuously. Buffer it.
		if m.RAMEnabled {
			if err := m.RAM.Save(); err != nil {
				log.Sub("mbc").Warningf("save RAM failed (%s)", err)
			}
		}
	case addr >= 0x2000 && addr <= 0x3fff:
		if value&0x1f == 0 { // Bank 0 (or over 0x1f) is not selectable.
			value = 1
		}
		m.ROMBank = m.ROMBank&0x60 | value&0x1f
		log.Sub("mbc/write").Debugf("ROM Bank 0x%02x selected", m.ROMBank)
	case addr >= 0x4000 && addr <= 0x5fff:
		if m.ROMBankingMode {
			m.ROMBank = m.ROMBank&0x1f | (value&3)<<5
			log.Sub("mbc/write").Debugf("ROM Bank 0x%02x selected", m.ROMBank)
		} else {
			m.RAMBank = value & 3
		}
	case addr >= 0x6000 && addr <= 0x7fff:
		m.ROMBankingMode = (value == 0)
	case addr >= 0xa000 && addr <= 0xbfff:
		if !m.RAMEnabled {
			log.Sub("mbc/write").Desperatef("RAM not enabled, write to 0x%04x ignored.",
				addr)
			return
		}
		// FIXME: this looks messy, shouldn't banking be handled in RAM itself?
		m.RAM.Write(uint16(m.RAMBank)*0x2000+uint16(addr-0xa000), value)
	}
}
