package options

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

// DateFormat layout for generated file names.
const DateFormat = "2006-01-02-15-04-05.000"

const (
	// DefaultConfigDir is the base folder where to store the emulator's config
	// file and subfolders for saves/screenshots.
	DefaultConfigDir = "~/.goholint/"

	// DefaultConfigPath is the path to our config file in the user's home.
	DefaultConfigPath = DefaultConfigDir + "config.ini"

	// DefaultConfig contains a reasonable default config.ini that's used
	// automatically if no config exists at run time
	//
	// TODO: template using our real default values?
	DefaultConfig = `# Most of the flags (except, obviously, -config) can be overridden here with
# the exact same name. See -help for details.

#boot = path/to/dmg_rom.bin
#cpuprofile = path/to/cpuprofile.pprof
#fastboot = 1
#level = debug
#mono = 1
#vsync = 1
#waitkey = 1
#zoom = 1

# Customize the default GB palette and UI colors here. Use hex RGB format.
[colors]
gb-0  = e0f0e7 # Lightest
gb-1  = 8ba394 # Light
gb-2  = 55645a # Dark
gb-3  = 343d37 # Darkest
ui-bg = ffffff # UI Background (outline)
ui-fg = 000000 # UI foreground (text)

# Define custom palettes here. Use short names and hex RGB format for colors.
# Format is: <name> = <lightest> <light> <dark> <darkest>
# The following are courtesy of lospec.com. They have tons more.
[palettes]
awakening = ffffb5 7bc67b 6b8c42 5a3921 # https://lospec.com/palette-list/links-awakening-sgb
icarus    = cef7f7 f78e50 9e0000 1e0000 # https://lospec.com/palette-list/kid-icarus-sgb
kirby     = f7bef7 e78686 7733e7 2c2c96 # https://lospec.com/palette-list/kirby-sgb
mario2    = eff7b6 dfa677 11c600 000000 # https://lospec.com/palette-list/super-mario-land-2-sgb
megaman   = cecece 6f9edf 42678e 102533 # https://lospec.com/palette-list/megaman-v-sgb
metroid   = aedf1e b62558 047e60 2c1700 # https://lospec.com/palette-list/metroid-ii-sgb
pokemon   = ffefff f7b58c 84739c 181010 # https://lospec.com/palette-list/pokemon-sgb
sgb       = f7e7c6 d68e49 a63725 331e50 # https://lospec.com/palette-list/nintendo-super-gameboy

# Define your keymap below with <action>=<key>. Key codes are taken from the
# SDL2 documentation (https://wiki.libsdl.org/SDL_Keycode) without the SDLK_
# prefix, and all supported actions are listed hereafter.
[keymap]
up     = UP        # Joypad Up
down   = DOWN      # Joypad Down
left   = LEFT      # Joypad Left
right  = RIGHT     # Joypad Right
a      = s         # A Button
b      = d         # B Button
select = BACKSPACE # Select Button
start  = RETURN    # Start Button

screenshot = F12   # Save a screenshot

recordgif = g      # Start/stop recording video output to GIF

# Cycle through custom palettes.
nexpalette      = PAGEDOWN
previouspalette = PAGEUP

# Enable or disable audio voices.
togglevoice1 = 1
togglevoice2 = 2
togglevoice3 = 3
togglevoice4 = 4

# Cleanly exit the program.
quit = CTRL+q

# TODO: reset, snapshot...
`
)

// DefaultKeymap is a reasonable default mapping for QWERTY/AZERTY layouts.
var DefaultKeymap = Keymap{
	"up":              KeyStroke{Code: sdl.K_UP},
	"down":            KeyStroke{Code: sdl.K_DOWN},
	"left":            KeyStroke{Code: sdl.K_LEFT},
	"right":           KeyStroke{Code: sdl.K_RIGHT},
	"a":               KeyStroke{Code: sdl.K_s},
	"b":               KeyStroke{Code: sdl.K_d},
	"select":          KeyStroke{Code: sdl.K_BACKSPACE},
	"start":           KeyStroke{Code: sdl.K_RETURN},
	"screenshot":      KeyStroke{Code: sdl.K_F12},
	"recordgif":       KeyStroke{Code: sdl.K_g},
	"nextpalette":     KeyStroke{Code: sdl.K_PAGEDOWN},
	"previouspalette": KeyStroke{Code: sdl.K_PAGEUP},
	"togglevoice1":    KeyStroke{Code: sdl.K_1},
	"togglevoice2":    KeyStroke{Code: sdl.K_2},
	"togglevoice3":    KeyStroke{Code: sdl.K_3},
	"togglevoice4":    KeyStroke{Code: sdl.K_4},
	"quit":            KeyStroke{Code: sdl.K_q, Mod: sdl.KMOD_LCTRL},
}

// Default palette colors with separate RGB components for easier use with SDL
// API.
const (
	// Arbitrary default colors that looked good on my screen. Kinda greenish.
	ColorWhiteRGB     = 0xe0f0e7
	ColorLightGrayRGB = 0x8ba394
	ColorDarkGrayRGB  = 0x55645a
	ColorBlackRGB     = 0x343d37

	ColorWhiteR     = (ColorWhiteRGB >> 16) & 0xff
	ColorWhiteG     = (ColorWhiteRGB >> 8) & 0xff
	ColorWhiteB     = ColorWhiteRGB & 0xff
	ColorLightGrayR = (ColorLightGrayRGB >> 16) & 0xff
	ColorLightGrayG = (ColorLightGrayRGB >> 16) & 0xff
	ColorLightGrayB = ColorLightGrayRGB & 0xff
	ColorDarkGrayR  = (ColorDarkGrayRGB >> 16) & 0xff
	ColorDarkGrayG  = (ColorDarkGrayRGB >> 16) & 0xff
	ColorDarkGrayB  = ColorDarkGrayRGB & 0xff
	ColorBlackR     = (ColorBlackRGB >> 16) & 0xff
	ColorBlackG     = (ColorBlackRGB >> 16) & 0xff
	ColorBlackB     = ColorBlackRGB & 0xff
)

var ColorWhite = color.RGBA{ColorWhiteR, ColorWhiteG, ColorWhiteB, 0xff}
var ColorLightGray = color.RGBA{ColorLightGrayR, ColorLightGrayG, ColorLightGrayB, 0xff}
var ColorDarkGray = color.RGBA{ColorDarkGrayR, ColorDarkGrayG, ColorDarkGrayB, 0xff}
var ColorBlack = color.RGBA{ColorBlackR, ColorBlackG, ColorBlackB, 0xff}

// DefaultPalette represents the selectable colors in the DMG.
var DefaultPalette = []color.RGBA{
	ColorWhite,
	ColorLightGray,
	ColorDarkGray,
	ColorBlack,
}

// Default UI colors. Black text, white outline.
var DefaultUIBackground = color.RGBA{0x00, 0x00, 0x00, 0xff}
var DefaultUIForeground = color.RGBA{0xff, 0xff, 0xff, 0xff}
