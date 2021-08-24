package options

import (
	"fmt"
	"image/color"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/veandco/go-sdl2/sdl"

	"gopkg.in/ini.v1"
)

// Keymap associating an action name (joypad input, UI command...) to an input.
type Keymap map[string]sdl.Keycode

const (
	// DefaultConfigPath is the path to our config file in the user's home.
	DefaultConfigPath = "~/.goholint/config.ini"

	// DefaultConfig contains a reasonable default config.ini that's used
	// automatically if no config exists at run time. TODO: embed from file?
	DefaultConfig = `# Most of the flags (except, obviously, -config) can be overridden here with
# the exact same name. See -help for details.

#boot = path/to/dmg_rom.bin
#cpuprofile = path/to/cpuprofile.pprof
#level = debug
#fastboot = 1
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
icarus = cef7f7 f78e50 9e0000 1e0000    # https://lospec.com/palette-list/kid-icarus-sgb
kirby = f7bef7 e78686 7733e7 2c2c96     # https://lospec.com/palette-list/kirby-sgb
mario2 = eff7b6 dfa677 11c600 000000    # https://lospec.com/palette-list/super-mario-land-2-sgb
megaman = cecece 6f9edf 42678e 102533   # https://lospec.com/palette-list/megaman-v-sgb
metroid = aedf1e b62558 047e60 2c1700   # https://lospec.com/palette-list/metroid-ii-sgb
pokemon = ffefff f7b58c 84739c 181010   # https://lospec.com/palette-list/pokemon-sgb
sgb = f7e7c6 d68e49 a63725 331e50       # https://lospec.com/palette-list/nintendo-super-gameboy

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

screenshot = F12   # Save a screenshot in the current directory

recordgif = g      # Start/stop recording video output to GIF

# Cycle through custom palettes.
nexpalette      = PAGEDOWN
previouspalette = PAGEUP

# TODO: quit, reset, snapshot...
`
)

// DefaultKeymap is a reasonable default mapping for QWERTY/AZERTY layouts.
var DefaultKeymap = Keymap{
	"up":              sdl.K_UP,
	"down":            sdl.K_DOWN,
	"left":            sdl.K_LEFT,
	"right":           sdl.K_RIGHT,
	"a":               sdl.K_s,
	"b":               sdl.K_d,
	"select":          sdl.K_BACKSPACE,
	"start":           sdl.K_RETURN,
	"screenshot":      sdl.K_F12,
	"recordgif":       sdl.K_g,
	"nextpalette":     sdl.K_PAGEDOWN,
	"previouspalette": sdl.K_PAGEUP,
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

// Go doesn't natively handle ~ in paths, fair enough.
func expandHome(path string) string {
	if path[0] == '~' {
		if u, err := user.Current(); err == nil {
			path = filepath.Join(u.HomeDir, path[1:])
		}
	}
	return path
}

// configKey returns a config key by the given name if it's present in the file
// and not already set by command-line arguments.
func configKey(cfg *ini.File, flags map[string]bool, name string) *ini.Key {
	// FIXME: handle section but so far I only use one for controls.
	if !flags[name] && cfg.Section("").HasKey(name) {
		return cfg.Section("").Key(name)
	}
	return nil
}

// Apply a parameter value from the config file to the string variable whose
// address is given, if that parameter was present in the file and not already
// set on the command-line.
func apply(cfg *ini.File, flags map[string]bool, name string, dst *string) {
	if key := configKey(cfg, flags, name); key != nil {
		*dst = key.String()
	}
}

// Same as apply for booleans.
func applyBool(cfg *ini.File, flags map[string]bool, name string, dst *bool) {
	if key := configKey(cfg, flags, name); key != nil {
		if b, err := key.Bool(); err == nil {
			*dst = b
		}
	}
}

// Same as apply for unsigned integers.
func applyUint(cfg *ini.File, flags map[string]bool, name string, dst *uint) {
	if key := configKey(cfg, flags, name); key != nil {
		if i, err := key.Uint(); err == nil {
			*dst = i
		}
	}
}

// Apply a parameter value from the config file to the color variable whose
// address is given, if that parameter was present in the file. If the color
// value can't be parsed, it's silently ignored.
func applyColor(s *ini.Section, name string, dst *color.RGBA) {
	if !s.HasKey(name) {
		return
	}

	if key := s.Key(name); key != nil {
		// Colors should be in hexadecimal (without 0x prefix).
		if rgb, err := strconv.ParseUint(key.String(), 16, 32); err == nil {
			dst.R = uint8((rgb >> 16) & 0xff)
			dst.G = uint8((rgb >> 8) & 0xff)
			dst.B = uint8(rgb & 0xff)
			dst.A = 0xff
		} else {
			fmt.Printf("Invalid value for color '%s': %v\n", name, err)
		}
	}
}

// Add a custom-defined palette to Options.
func (o *Options) addPalette(name, value string) {
	hexColors := strings.Fields(value)
	if len(hexColors) != 4 {
		return
	}

	palette := make([]color.RGBA, 4)
	for i := range palette {
		// Colors should be in hexadecimal (without 0x prefix).
		if rgb, err := strconv.ParseUint(hexColors[i], 16, 32); err == nil {
			palette[i].R = uint8((rgb >> 16) & 0xff)
			palette[i].G = uint8((rgb >> 8) & 0xff)
			palette[i].B = uint8(rgb & 0xff)
			palette[i].A = 0xff
		} else {
			fmt.Printf("Invalid value for color %d in palette %s: %v\n", i, name, err)
			return // Ignore palettes with invalid colors
		}
	}
	o.Palettes = append(o.Palettes, palette)
	o.PaletteNames = append(o.PaletteNames, name)
}

// Attempt to create home config folder and put our default config there, if
// it doesn't already exist.
func createDefaultConfig() {
	configPath := expandHome(DefaultConfigPath)
	folder := filepath.Dir(configPath)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Println("Creating default config folder.")

		if err := os.MkdirAll(folder, 0755); err != nil {
			fmt.Printf("Can't create config folder %s: %v\n", folder, err)
			return
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Creating default config.")

		f, err := os.Create(configPath)
		if err != nil {
			fmt.Printf("Creating %s failed: %v\n", configPath, err)
			return
		}
		defer f.Close()

		if _, err := f.WriteString(DefaultConfig); err != nil {
			fmt.Printf("Writing default config failed: %v\n", err)
			return
		}
	}
}

// Update reads all parameters from a given configuration file and updates the
// Options instance with those values, skipping all options that may already
// have been set on the command-line.
func (o *Options) Update(configPath string, flags map[string]bool) {
	if configPath == "" {
		return
	}

	cfg, err := ini.Load(configPath)
	if err != nil {
		// No real error handling, this method should be forgiving.
		fmt.Printf("Can't load config file %s (%s)\n", configPath, err)
		return
	}

	// Using quick and dirty helpers because mixed types and lazy.
	apply(cfg, flags, "boot", &o.BootROM)
	apply(cfg, flags, "cpuprofile", &o.CPUProfile)
	// TODO: debug special format.
	apply(cfg, flags, "level", &o.DebugLevel)
	applyBool(cfg, flags, "fastboot", &o.FastBoot)
	applyBool(cfg, flags, "nosync", &o.VSync)
	// TODO: savedir (and just ditch savepath altogether)
	applyBool(cfg, flags, "waitkey", &o.WaitKey)
	applyUint(cfg, flags, "zoom", &o.ZoomFactor)

	// Ignoring flags that are not really interesting as a config, such as
	// -cyles, -gif or -rom...

	// Set keymap here. Build on top of default. TODO: validate.
	keySection := cfg.Section("keymap")
	for key := range o.Keymap {
		// Key() will return the empty string if it doesn't exist, it's fine.
		keyName := keySection.Key(key).String()
		keySym := sdl.GetKeyFromName(keyName)
		if keySym != sdl.K_UNKNOWN {
			o.Keymap[key] = keySym
		}
	}

	// Set colors here. Build on top of default as well.
	colorSection := cfg.Section("colors")

	// Default palette is palette 0.
	applyColor(colorSection, "gb-0", &o.Palettes[0][0])
	applyColor(colorSection, "gb-1", &o.Palettes[0][1])
	applyColor(colorSection, "gb-2", &o.Palettes[0][2])
	applyColor(colorSection, "gb-3", &o.Palettes[0][3])

	applyColor(colorSection, "ui-bg", &o.UIBackground)
	applyColor(colorSection, "ui-fg", &o.UIForeground)

	// Add custom palettes.
	palettesSection := cfg.Section("palettes")
	for _, palName := range palettesSection.KeyStrings() {
		palValue := palettesSection.Key(palName).String()
		o.addPalette(palName, palValue)
	}
}
