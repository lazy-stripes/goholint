package options

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"

	"gopkg.in/ini.v1"
)

// Keymap associating an action name (joypad input, UI command...) to an input.
type Keymap map[string]sdl.Keycode

// DefaultKeymap is a reasonable default mapping for QWERTY/AZERTY layouts.
var DefaultKeymap = Keymap{
	"up":         sdl.K_UP,
	"down":       sdl.K_DOWN,
	"left":       sdl.K_LEFT,
	"right":      sdl.K_RIGHT,
	"a":          sdl.K_s,
	"b":          sdl.K_d,
	"select":     sdl.K_BACKSPACE,
	"start":      sdl.K_RETURN,
	"screenshot": sdl.K_F12,
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

// apply a parameter value from the config file to the string variable whose
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

// Samme as apply for unsigned integers.
func applyUint(cfg *ini.File, flags map[string]bool, name string, dst *uint) {
	if key := configKey(cfg, flags, name); key != nil {
		if i, err := key.Uint(); err == nil {
			*dst = i
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

	// Go doesn't natively handle ~ in paths, fair enough.
	if configPath[0] == '~' {
		if u, err := user.Current(); err == nil {
			configPath = filepath.Join(u.HomeDir, configPath[1:])
		}
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
	applyBool(cfg, flags, "nosync", &o.NoSync)
	// TODO: savedir (and just ditch savepath altogether)
	applyBool(cfg, flags, "waitkey", &o.WaitKey)
	applyUint(cfg, flags, "zoom", &o.ZoomFactor)

	// Ignoring options that are not really interesting as a config.
	// Such as -cyles, -gif or -rom...

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
}
