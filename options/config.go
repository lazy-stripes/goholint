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
	applyBool(cfg, flags, "fastboot", &o.FastBoot)
	apply(cfg, flags, "level", &o.DebugLevel)
	applyBool(cfg, flags, "mono", &o.Mono)
	applyBool(cfg, flags, "vsync", &o.VSync)
	applyBool(cfg, flags, "waitkey", &o.WaitKey)
	applyUint(cfg, flags, "zoom", &o.ZoomFactor)

	// Ignoring flags that are not really interesting as a config, such as
	// -cyles, -gif or -rom...

	// Set keymap here. Build on top of default. TODO: validate.
	keySection := cfg.Section("keymap")
	for key := range o.Keymap {
		// Key() will return the empty string if it doesn't exist, it's fine.
		keyName := keySection.Key(key).String()
		if keyName == "" {
			continue
		}

		keySym := sdl.GetKeyFromName(keyName)
		if keySym != sdl.K_UNKNOWN {
			o.Keymap[key] = keySym
		} else {
			fmt.Printf("Unknown key name '%s' for action '%s'.\n", keyName, key)
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
