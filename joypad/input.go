package joypad

// Input storing needed bits for a button or a direction.
type Input struct {
	Selector uint8 // P14 or P15
	Bit      uint8 // Bit position in JOYP register
	State    bool  // Button state (proper logic here: true is on, false is off)
}

var InputRight = Input{P14, P10, false}
var InputLeft = Input{P14, P11, false}
var InputUp = Input{P14, P12, false}
var InputDown = Input{P14, P13, false}

var InputA = Input{P15, P10, false}
var InputB = Input{P15, P11, false}
var InputSelect = Input{P15, P12, false}
var InputStart = Input{P15, P13, false}
