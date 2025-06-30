package interrupts

import (
	"fmt"

	"github.com/lazy-stripes/goholint/logger"
)

// Package-wide logger.
var log = logger.New("interrupts", "interrupt manager")

// Namespaced const flags because still procrastinating.
const (
	VBlank  = 1 << iota // Bit 0 (INT 40h)
	LCDStat             // Bit 1 (INT 48h)
	Timer               // Bit 2 (INT 50h)
	Serial              // Bit 3 (INT 58h)
	Joypad              // Bit 4 (INT 60h)
)

// LCDC STAT specific bits for PPU-related interrupts
const (
	STATMode0 = 1 << (iota + 3) // Bit 3 (HBlank)
	STATMode1                   // Bit 4 (VBlank)
	STATMode2                   // Bit 5 (OAM search)
	STATLYCLY                   // Bit 6 (LYC=LY)

)

// Interrupt vectors and register addresses.
const (
	AddrVBlank  = 0x0040
	AddrLCDStat = 0x0048
	AddrTimer   = 0x0050
	AddrSerial  = 0x0058
	AddrJoypad  = 0x0060
	AddrIF      = 0xff0f
	AddrIE      = 0xffff
)

// Addresses is a quick and dirty mapping between an interrupt flag and
// its address.
var Addresses = [...]uint16{
	VBlank:  AddrVBlank,
	LCDStat: AddrLCDStat,
	Timer:   AddrTimer,
	Serial:  AddrSerial,
	Joypad:  AddrJoypad,
}

var Names = [...]string{
	VBlank:  "VBlank",
	LCDStat: "LCDStat",
	Timer:   "Timer",
	Serial:  "Serial",
	Joypad:  "Joypad",
}

// Interrupts represents an address space to access IF (5 LSB bits) and IE
// with added methods to request interrupts.
type Interrupts struct {
	regIF, regIE *uint8
}

// New interrupt registers.
func New(regIF, regIE *uint8) *Interrupts {
	return &Interrupts{regIF, regIE}
}

// Contains returns true if the given address belongs to the address space,
// false otherwise.
func (i *Interrupts) Contains(addr uint16) bool {
	return addr == AddrIF || addr == AddrIE
}

// Read returns the value stored in the requested interrupt register, accounting
// for IF exploitable bits.
func (i *Interrupts) Read(addr uint16) uint8 {
	switch addr {
	case AddrIF:
		return *i.regIF & 0x1f
	case AddrIE:
		return *i.regIE
	}
	panic(fmt.Sprintf("Broken MMU: out-of-range address %#x requested", addr))
}

// Write stores the given value in the given interrupt register, accounting for
// IF exploitable bits.
func (i *Interrupts) Write(addr uint16, value uint8) {
	switch addr {
	case AddrIF:
		*i.regIF = value & 0x1f
		log.Debugf(" !!! IF=%#02x", value)
	case AddrIE:
		*i.regIE = value
		log.Debugf(" !!! IE=%#02x", value)
	}
}

// Request sets the bit corresponding to the requested interrupt type.
func (i *Interrupts) Request(interrupt uint8) {
	*i.regIF |= interrupt
	log.Desperatef("requested interrupt %s (IF=%#02x)", Names[interrupt], *i.regIF)
}
