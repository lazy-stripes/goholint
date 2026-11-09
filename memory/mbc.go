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
	ROM *BankedROM // ROM (up to 2MB or 128 banks)
	RAM *BankedRAM // Optional RAM (up to 32KB or 4 banks)

	RAMEnabled bool // 0000-1FFF - RAM Enable

	BankLow  int // 2000-3FFF - ROM Bank Number (bits 0-4)
	BankHigh int // 4000-5FFF - RAM Bank Number / ROM Bank Number (bits 5-6)

	// 6000-7FFF - ROM/RAM Mode Select
	// 00h = ROM Banking Mode (up to 8KByte RAM, 2MByte ROM) (default)
	// 01h = RAM Banking Mode (up to 32KByte RAM, 512KByte ROM)
	BankingMode uint8

	// Pre-computed mask for ROM bank number. Set to (ROM.Banks() - 1).
	// E.g. 32 banks need 5 bits for addressing, so mask is 31 (0x11111).
	romBanksMask int
}

// NewMBC1 creates an address space emulating a cartridge with an MBC1 mapper.
// Takes instances of BankedROM (and BankedRAM is available) because file
// handling (and RAM loading/saving) is handled at the Cartridge level.
func NewMBC1(rom *BankedROM, ram *BankedRAM) *MBC1 {
	return &MBC1{
		ROM:     rom,
		RAM:     ram,
		BankLow: 1,

		romBanksMask: rom.Bank1.Banks() - 1,
	}
}

// Contains returns true if the requested address is anywhere in ROM or RAM.
func (m *MBC1) Contains(addr uint16) bool {
	return addr < 0x8000 || (addr >= 0xa000 && addr < 0xc000)
}

// Read returns the byte at requested address in current ROM or RAM bank.
func (m *MBC1) Read(addr uint16) uint8 {
	switch {
	case addr < 0x8000:
		value := m.ROM.Read(addr)
		log.Sub("mbc/read").Desperatef("ROM[%04x]: 0x%02x", addr, value)
		return value

	case addr >= 0xa000 && addr < 0xc000:
		if m.RAMEnabled && (m.RAM != nil) {
			value := m.RAM.Read(addr)
			log.Sub("mbc/read").Desperatef("RAM[%04x]: 0x%02x", addr, value)
			return value
		}
		log.Sub("mbc/read").Desperatef("RAM disabled, read ignored")
		return 0xff

	default:
		panic("out of bounds MBC1 read")
	}
}

// recomputeBanks updates ROM and RAM bank numbers after writes to the MBC1.
func (m *MBC1) recomputeBanks() {
	if (m.BankingMode == RAMBanking) && (m.RAM != nil) {
		if m.BankHigh < m.RAM.Banks() {
			// FIXME: confirm behavior. We could just do SetBank() and let it wrap from within BankedRAM.
			m.RAM.Set(m.BankHigh)
		}

		m.ROM.Bank0.Set((m.BankHigh << 5) & m.romBanksMask)
	} else {
		if m.RAM != nil {
			m.RAM.Set(0)
		}
		m.ROM.Bank0.Set(0)
	}

	// 4000-7fff range uses all 7 bank number bits regardless of mode.
	m.ROM.Bank1.Set(((m.BankHigh << 5) + m.BankLow) & m.romBanksMask)
}

// Write value to RAM, enable RAM or select ROM/RAM banks.
func (m *MBC1) Write(addr uint16, value uint8) {
	switch {
	// 0000-1FFF - RAM Enable
	case addr < 0x2000:
		m.RAMEnabled = (value&0x0f == 0x0a)
		log.Sub("mbc/write").Desperatef("RAMEnabled=%t", m.RAMEnabled)

	// 2000-3FFF - ROM Bank Number
	case addr >= 0x2000 && addr < 0x4000:

		m.BankLow = int(value & 0x1f)
		if m.BankLow == 0 {
			m.BankLow = 1
		}
		m.recomputeBanks()
		log.Sub("mbc/write").Debugf("BankLow=0x%02x", value&0x1f)

	// 4000-5FFF - RAM Bank Number - or - Upper Bits of ROM Bank Number
	case addr >= 0x4000 && addr < 0x6000:
		m.BankHigh = int(value & 3)
		m.recomputeBanks()
		log.Sub("mbc/write").Debugf("BankHigh=0x%02x", value&3)

	// 6000-7FFF - ROM/RAM Mode Select
	case addr >= 0x6000 && addr < 0x8000:
		log.Sub("mbc/write").Debugf("Banking Mode 0x%02x selected", value)
		m.BankingMode = value & 1
		m.recomputeBanks()

	// A000-BFFF - RAM Bank 00-03, if any
	case addr >= 0xa000 && addr < 0xc000:
		// Avenging Spirit uses MBC1 with no RAM, but still enables it and attempts
		// writing to it, so we still need to check whether there is RAM at all.
		if !m.RAMEnabled || (m.RAM == nil) {
			log.Sub("mbc/write").Desperatef("RAM not enabled, write to 0x%04x ignored.", addr)
			return
		}
		// Proper banks should have been set.
		m.RAM.Write(addr, value)
		log.Sub("mbc/write").Desperatef("RAM[%04x]=0x%02x", addr, value)

	default:
		panic("out of bounds MBC1 write")
	}
}

const MBC2RAMSize = 512

// MBC2 (max 256 KiB ROM and 512×4 bits RAM)
type MBC2 struct {
	ROM *BankedROM // Complete ROM (will be addressed according to ROMBank)
	RAM *RAM       // Built-in RAM (512 half-bytes)

	RAMEnabled bool // 0000-1FFF - RAM Enable
}

// NewMBC2 creates an address space emulating a cartridge with an MBC2 chip.
// Takes an instance of ROM because to know which kind of chip it uses, we had
// to read it beforehand anyway.
func NewMBC2(rom *BankedROM) *MBC2 {
	ram := NewRAM(AddrExternalRAM, MBC2RAMSize)

	m := &MBC2{
		ROM: rom,
		RAM: ram,
	}

	// [PANMBC2] The ROM bank is set to 1 by default.
	m.ROM.Bank1.Set(1)

	return m
}

// Contains returns true if the requested address is anywhere in ROM or RAM.
func (m *MBC2) Contains(addr uint16) bool {
	return addr < 0x8000 || (addr >= 0xa000 && addr < 0xc000)
}

// Read returns the byte at requested address in current ROM bank or RAM. RAM
// values are stored on 4 bits, the high nibble will be zeroed out. RAM is also
// only 512 bytes and wraps around.
func (m *MBC2) Read(addr uint16) uint8 {
	switch {
	case addr < 0x8000:
		return m.ROM.Read(addr)

	case addr >= 0xa000 && addr < 0xc000:
		if m.RAMEnabled {
			// Wrap around every 512 bytes (only 9 lowest bits of address used).
			wrappedAddr := AddrExternalRAM + (addr & (MBC2RAMSize - 1))
			value := m.RAM.Read(wrappedAddr) & 0x0f
			log.Sub("mbc/read").Desperatef("RAM[%04x]: 0x%02x", addr, value)
			return value
		}
		log.Sub("mbc/read").Desperatef("RAM disabled, read ignored")
		return 0xff

	default:
		panic("out of bounds MBC2 read")
	}
}

// Write value to RAM, enable RAM or select ROM/RAM banks.
func (m *MBC2) Write(addr uint16, value uint8) {
	switch {
	// 0000-3FFF - RAM Enable, ROM Bank Number
	case addr < 0x4000:
		// If bit 8 of address is clear, enable/disable RAM. Otherwise, select
		// ROM bank.
		if addr&0x0100 == 0 {
			m.RAMEnabled = (value&0x0f == 0x0a)
			log.Sub("mbc/write").Desperatef("RAMEnabled=%t", m.RAMEnabled)
		} else {
			// [PANMBC2] Specifically, the lower 4 bits of the value written to
			// this address range specify the ROM bank number. If bank 0 is
			// written, the resulting bank will be bank 1 instead.
			bank := int(value & 0x0f)
			if bank == 0 {
				bank = 1
			}
			m.ROM.Bank1.Set(bank)
			log.Sub("mbc/write").Debugf("ROMBank=0x%02x (bank %d)", value, bank)
		}

	// A000-BFFF - RAM
	case addr >= 0xa000 && addr < 0xc000:
		if !m.RAMEnabled {
			log.Sub("mbc/write").Desperatef("RAM not enabled, write to 0x%04x ignored.", addr)
			return
		}
		// Wrap around every 512 bytes (only 9 lowest bits of address used).
		wrappedAddr := AddrExternalRAM + (addr & (MBC2RAMSize - 1))
		m.RAM.Write(wrappedAddr, value&0x0f)
		log.Sub("mbc/write").Desperatef("RAM[%04x]=0x%02x", addr, value)

	default:
		panic("out of bounds MBC2 write")
	}
}
