package memory

// Memory Bank Controllers. Source:
// [PANMBC] https://gbdev.io/pandocs/#mbc1

// 6000-7FFF - ROM/RAM Mode Select
const (
	ROMBanking = 0x00
	RAMBanking = 0x01
)

// MBC1 (max 2MByte ROM and/or 32KByte RAM)
type MBC1 struct {
	*ROM            // Complete ROM (will be addressed according to ROMBank)
	*RAM            // Optional RAM (up to 32KB)
	RAMEnabled bool // 0000-1FFF - RAM Enable

	BankLow  uint8 // 2000-3FFF - ROM Bank Number
	BankHigh uint8 // 4000-5FFF - RAM Bank Number / Upper Bits of ROM Bank Number

	// 6000-7FFF - ROM/RAM Mode Select
	// 00h = ROM Banking Mode (up to 8KByte RAM, 2MByte ROM) (default)
	// 01h = RAM Banking Mode (up to 32KByte RAM, 512KByte ROM)
	BankingMode uint8

	// If true, save RAM values. (FIXME: when?)
	battery bool

	// Max number of ROM/RAM banks.
	romBanks uint8 // MBC1 has max 2 MByte ROM (128 banks)
	ramBanks uint8
}

// NewMBC1 creates an address space emulating a cartridge with an MBC1 chip.
// Takes an instance of ROM because to know which kind of chip it uses, we need
// to read it beforehand anyway.
func NewMBC1(rom *ROM, romBanks uint8, ramBanks uint8, battery bool, savePath string) *MBC1 {
	ramSize := uint16(ramBanks) * 0x2000
	ram := NewRAM(0, ramSize) // FIXME: base address and banks

	// If the cartridge has a battery-backed RAM, restore it here.
	if battery {
		if err := ram.Load(savePath); err != nil {
			log.Warning(err.Error())
		}
	}

	return &MBC1{
		ROM:     rom,
		RAM:     ram,
		BankLow: 1,

		romBanks: romBanks,
		ramBanks: ramBanks,
		battery:  battery,
	}
}

// ROMBank returns the currently selected ROM bank according to our internal
// registers.
func (m *MBC1) ROMBank() (bank uint8) {
	bank = m.BankLow
	if m.BankingMode == ROMBanking {
		bank |= m.BankHigh << 5
	}

	// [PANMBC] If the ROM Bank Number is set to a higher value than the number
	// of banks in the cart, the bank number is masked to the required number
	// of bits.

	// MBC1 only supports ROM sizes up to 2MB (128 banks) and all those
	// supported sizes are powers of two we can use as masks.
	bank &= m.romBanks - 1

	return
}

// RAMBank returns the currently selected RAM bank according to our internal
// registers.
func (m *MBC1) RAMBank() (bank uint8) {
	if m.ramBanks > 0 && m.BankingMode == RAMBanking {
		if m.BankHigh < m.ramBanks-1 {
			return m.BankHigh
		}
	}
	return 0
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
		// [GEEKIO] 7.2. says BANK2 is used if RAM banking mode is enabled.
		if m.BankingMode == RAMBanking {
			// TODO: Let ROM type check overflow on its own.
			bank := (m.BankLow << 5) & (m.romBanks - 1)
			return m.ROM.read(uint(bank)*0x4000 + uint(addr))
		}
		return m.ROM.Read(addr)

	case addr >= 0x4000 && addr <= 0x7fff:
		log.Sub("mbc/read").Desperatef("Read ROM at %x.",
			uint(m.ROMBank())*0x4000+uint(addr-0x4000))
		return m.ROM.read(uint(m.ROMBank())*0x4000 + uint(addr-0x4000))

	case m.RAMEnabled && addr >= 0xa000 && addr <= 0xbfff:
		return m.RAM.Read(uint16(m.RAMBank())*0x2000 + uint16(addr-0xa000))

	default:
		return 0xff
	}
}

// Write value to RAM, enable RAM or select ROM/RAM banks.
func (m *MBC1) Write(addr uint16, value uint8) {
	switch {
	// 0000-1FFF - RAM Enable
	case addr >= 0x0000 && addr <= 0x1fff:
		m.RAMEnabled = (value&0x0a == 0x0a)

	// 2000-3FFF - ROM Bank Number
	case addr >= 0x2000 && addr <= 0x3fff:
		if value&0x1f == 0 { // Bank 0 (or over 0x1f) is not selectable.
			value = 1
		}
		m.BankLow = value & 0x1f
		log.Sub("mbc/write").Debugf("BankLow=0x%02x", value&0x1f)

	// 4000-5FFF - RAM Bank Number - or - Upper Bits of ROM Bank Number
	case addr >= 0x4000 && addr <= 0x5fff:
		m.BankHigh = value & 3
		log.Sub("mbc/write").Debugf("BankHigh=0x%02x", value&3)

	// 6000-7FFF - ROM/RAM Mode Select
	case addr >= 0x6000 && addr <= 0x7fff:
		log.Sub("mbc/write").Debugf("Banking Mode 0x%02x selected", value)
		m.BankingMode = value & 1

	// A000-BFFF - RAM Bank 00-03, if any
	case addr >= 0xa000 && addr <= 0xbfff:
		if !m.RAMEnabled {
			log.Sub("mbc/write").Desperatef("RAM not enabled, write to 0x%04x ignored.",
				addr)
			return
		}
		// FIXME: this looks messy, shouldn't banking be handled in RAM itself?
		// Also maybe a mere boundary check here would do?
		m.RAM.Write(uint16(m.RAMBank())*0x2000+uint16(addr-0xa000), value)

		// Write save file on enable. FIXME: Buffer it.
		if m.battery && m.RAMEnabled {
			if err := m.RAM.Save(); err != nil {
				log.Sub("mbc").Warningf("save RAM failed (%s)", err)
			}
		}
	}
}
