package gameboy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lazy-stripes/goholint/memory"
	"github.com/veandco/go-sdl2/sdl"
)

// Action type for user interactions. This might move to a ui package someday.
type Action func(eventType uint32)

// JoypadUp updates the Joypad's registers for the Up direction.
func (g *GameBoy) JoypadUp(eventType uint32) {
	g.JPad.Up.State = (eventType == sdl.KEYDOWN)
}

// JoypadDown updates the Joypad's registers for the Down direction.
func (g *GameBoy) JoypadDown(eventType uint32) {
	g.JPad.Down.State = (eventType == sdl.KEYDOWN)
}

// JoypadLeft updates the Joypad's registers for the Left direction.
func (g *GameBoy) JoypadLeft(eventType uint32) {
	g.JPad.Left.State = (eventType == sdl.KEYDOWN)
}

// JoypadRight updates the Joypad's registers for the Right direction.
func (g *GameBoy) JoypadRight(eventType uint32) {
	g.JPad.Right.State = (eventType == sdl.KEYDOWN)
}

// JoypadA updates the Joypad's registers for the A button.
func (g *GameBoy) JoypadA(eventType uint32) {
	g.JPad.A.State = (eventType == sdl.KEYDOWN)
}

// JoypadB updates the Joypad's registers for the B button.
func (g *GameBoy) JoypadB(eventType uint32) {
	g.JPad.B.State = (eventType == sdl.KEYDOWN)
}

// JoypadSelect updates the Joypad's registers for the Select button.
func (g *GameBoy) JoypadSelect(eventType uint32) {
	g.JPad.Select.State = (eventType == sdl.KEYDOWN)
}

// JoypadStart updates the Joypad's registers for the Start button.
func (g *GameBoy) JoypadStart(eventType uint32) {
	g.JPad.Start.State = (eventType == sdl.KEYDOWN)
}

// Screenshot saves the current frame to disk as a PNG file.
// TODO: configurable folder, obviously.
func (g *GameBoy) Screenshot(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	g.Display.Screenshot()
}

// StartStopRecord starts recording video output to GIF and closes the file
// when done. Defined as a single action to toggle between the two and avoid
// opening several GIFs at once.
func (g *GameBoy) StartStopRecord(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	if g.recording { // TODO: query directly from screen?
		g.Display.StopRecord()
		g.recording = false
	} else {
		g.recording = true
		g.Display.StartRecord()
	}
}

// NextPalette switches colors to the next defined palette, wrapping around.
// There should always be at least a default palette in the config object.
func (g *GameBoy) NextPalette(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	g.paletteIndex = (g.paletteIndex + 1) % len(g.config.Palettes)
	g.Display.Palette(g.config.Palettes[g.paletteIndex])
	g.Display.Message(g.config.PaletteNames[g.paletteIndex], 2)
}

// PreviousPalette switches colors to the previous defined palette, wrapping
// around. There should always be at least a default palette in the config
// object.
func (g *GameBoy) PreviousPalette(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}

	g.paletteIndex -= 1
	if g.paletteIndex < 0 {
		// Wrap around (can't use % with negative values).
		g.paletteIndex = len(g.config.Palettes) - 1
	}
	g.Display.Palette(g.config.Palettes[g.paletteIndex])
	g.Display.Message(g.config.PaletteNames[g.paletteIndex], 2)
}

// Helper strings to format UI messages.
var voiceNames = [4]string{
	"Square 1",
	"Square 2",
	"Wave",
	"Noise",
}

func (g *GameBoy) voiceStatusMsg(voice int) string {
	var sb strings.Builder
	for _, m := range g.APU.Muted {
		if m {
			sb.WriteRune('-')
		} else {
			sb.WriteRune('♪')
		}
	}
	sb.WriteRune('\n')
	sb.WriteString(voiceNames[voice])
	if g.APU.Muted[voice] {
		sb.WriteString(" muted")
	} else {
		sb.WriteString(" enabled")
	}
	return sb.String()
}

// ToggleVoice1 mutes or unmutes the first audio generator (Square 1).
func (g *GameBoy) ToggleVoice1(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.APU.Muted[0] = !g.APU.Muted[0]
	g.Display.Message(g.voiceStatusMsg(0), 2)
}

// ToggleVoice2 mutes or unmutes the second audio generator (Square 2).
func (g *GameBoy) ToggleVoice2(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.APU.Muted[1] = !g.APU.Muted[1]
	g.Display.Message(g.voiceStatusMsg(1), 2)
}

// ToggleVoice3 mutes or unmutes the third audio generator (Wave).
func (g *GameBoy) ToggleVoice3(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.APU.Muted[2] = !g.APU.Muted[2]
	g.Display.Message(g.voiceStatusMsg(2), 2)
}

// ToggleVoice4 mutes or unmutes the fourth audio generator (Noise).
func (g *GameBoy) ToggleVoice4(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.APU.Muted[3] = !g.APU.Muted[3]
	g.Display.Message(g.voiceStatusMsg(3), 2)
}

// Quit cleanly quits the program.
func (g *GameBoy) Quit(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.QuitChan <- true
}

// Home stops the emulation to display some kind of home menu. Clearing it
// should be done from the UI itself. Which, in itself, is going to be another
// whole thing.
//
// Doing this should pave the way for high-level stuff the emulator should do on
// VBlank like saving or loading states and suchlike.
func (g *GameBoy) Home(eventType uint32) {
	if eventType != sdl.KEYDOWN {
		return
	}
	g.home = !g.home
	if g.home {
		g.Display.Message("HOME", 2)

		// TODO: stop audio callback, start waiting on SDL Events, use our own set of UI Actions.

		// TEMPORARY
		// Display Scanner menu. This will keep the emulator halted. I'll want
		// to integrate everything into the main loop in the end anyway.
		g.ScannerMenu()
	}
}

// TODO: ui.go or somewhere in ui package.
// Oh boy delegating events is gonna be so fun isn't it.
func (g *GameBoy) ScannerMenu() {
	if g.scanner == nil {
		// Capture initial state.
		g.scanner = memory.NewScanner(g.CPU.Memory)
	}

	// No implementing CLI logic, this is a quick and dirty test, I promise.
	// Also printing out a menu takes me back. So, so far.

	searchResults := g.scanner.Haystack
	reader := bufio.NewReader(os.Stdin)
	choice := ""
	for choice != "s" && choice != "q" {
		fmt.Printf("\nScanner state: %d values\n\n", len(searchResults))

		// Did we find something?
		if len(searchResults) < 5 {
			for addr, val := range searchResults {
				fmt.Printf(" %#4x: %#2x\n", addr, val)
			}
			fmt.Println()
		}

		fmt.Println(" 0-255 - Look up value")

		// Comparison operators need at least one initial haystack.
		if len(g.scanner.Haystack) > 0 {
			fmt.Println(" g     - Larger values")
			fmt.Println(" l     - Smaller values")
			fmt.Println(" e     - Equal (unchanged) values")
		}

		fmt.Println(" c     - Clear state")
		fmt.Println(" s     - Save results and close")
		//fmt.Println(" f     - Force memory value") // TODO: memory.Overlay
		fmt.Println(" q     - Discard results and close")
		fmt.Print("\n > ")

		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		choice = strings.ToLower(choice)

		// Weed out invalid commands.
		if len(g.scanner.Haystack) == 0 && strings.Contains("gle", choice) {
			fmt.Println("Needs at least one result set. Try 's' instead.")
			continue
		}

		switch choice {
		case "g":
			searchResults = g.scanner.LookUpLarger()

		case "l":
			searchResults = g.scanner.LookUpSmaller()

		case "e":
			searchResults = g.scanner.LookUpEqual()

		case "c":
			g.scanner.Clear()

		case "s":
			g.scanner.Haystack = searchResults

		case "q":
			// exit through loop condition
			continue

		default:
			// Try converting choice into a lookup value.
			value, err := strconv.Atoi(choice)
			if err != nil || value > 255 {
				fmt.Println("Please enter a value between 0 and 255")
				continue
			}
			searchResults = g.scanner.LookUp(uint8(value))
		}
	}

	g.home = false
}

// TODO: so many things! Save states, toggle features...
