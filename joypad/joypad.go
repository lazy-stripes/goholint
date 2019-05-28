package joypad

// Source: [JOYPAD] http://gbdev.gg8.se/wiki/articles/Joypad_Input

// Register addresses
const (
	AddrJOYP = 0xff00
)

// Bit values
const (
	P10 = 1 << iota // Bit 0 - Input Right or Button A (0=Pressed) (Read Only)
	P11             // Bit 1 - Input Left  or Button B (0=Pressed) (Read Only)
	P12             // Bit 2 - Input Up    or Select   (0=Pressed) (Read Only)
	P13             // Bit 3 - Input Down  or Start    (0=Pressed) (Read Only)
	P14             // Bit 4 - Select Direction Keys   (0=Select)
	P15             // Bit 5 - Select Button Keys      (0=Select)
)

// Joypad register and event manager for game inputs.
type Joypad struct {
	JOYP   uint8
	Keymap map[uint]Input
}

// New instantiates a Joypad addressable mapping to FF00 that will wait for
// events from the main loop.
func New(keymap map[uint]Input) *Joypad {
	return &Joypad{Keymap: keymap}
}

// Contains returns true if the requested address is the JOYP register.
func (j *Joypad) Contains(addr uint16) bool {
	return addr == AddrJOYP
}

// Read returns the state of selected inputs (inverted logic).
func (j *Joypad) Read(addr uint16) (value uint8) {
	value = j.JOYP & 0x30
	// Set bits for inactive/unselected inputs. Easier to read, I thought.
	for _, input := range j.Keymap {
		if j.JOYP&input.Selector == 0 || !input.State {
			value |= input.Bit
		}
	}
	return value
}

func (j *Joypad) Write(addr uint16, value uint8) {
	j.JOYP = value & 0x30
}
