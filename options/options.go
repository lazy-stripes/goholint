package options

import (
	"flag"
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// ModMask only keeps the keyboard modifiers we allow in keymaps.
const ModMask = sdl.KMOD_CTRL | sdl.KMOD_SHIFT

// KeyStroke describes a single keyboard input, including potential modifiers.
type KeyStroke struct {
	Code sdl.Keycode
	Mod  sdl.Keymod
}

// Keymap associating an action name (joypad input, UI command...) to an input.
type Keymap map[string]KeyStroke

// Options structure grouping command line flags and config file values.
type Options struct {
	BootROM      string         // -boot <path>
	CPUProfile   string         // -cpuprofile <path>
	DebugLevel   string         // -level <debug level>
	DebugModules module         // -debug <module>
	Duration     uint           // -cycles <amount>
	FastBoot     bool           // -fastboot
	GIFPath      string         // -gif <path>
	Keymap       Keymap         // From config.
	Mono         bool           // -mono
	Palettes     [][]color.RGBA // From config.
	PaletteNames []string       // From config, same order.
	ROMPath      string         // -rom <path>
	SavePath     string         // -save <full path>
	UIBackground color.RGBA     // From config.
	UIForeground color.RGBA     // From config.
	VSync        bool           // -vsync
	WaitKey      bool           // -waitkey
	ZoomFactor   uint           // -zoom <factor>
}

// CreateFileIn creates a new file with the requested suffix (which can be only
// an extension, a timestamp + an extension, etc) in the requested subfolder.
// The folder will be created under the configuration path if it doesn't already
// exist.
//
// Returns an open file or an error.
func CreateFileIn(subfolder, suffix string) (*os.File, error) {
	// TODO: could be nice to add metadata to file name, like cartridge name.
	folder := filepath.Join(expandHome(DefaultConfigDir), subfolder)
	filename := fmt.Sprintf("goholint-%s%s", time.Now().Format(DateFormat), suffix)
	path := filepath.Join(folder, filename)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Printf("Creating subfolder %s\n", folder)

		if err := os.MkdirAll(folder, 0755); err != nil {
			fmt.Printf("Can't create subfolder %s: %v\n", folder, err)
			return nil, err
		}
	}
	return os.Create(path)
}

// User-defined type to parse a list of module names for which debug output must be enabled.
type module []string

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (m *module) String() string {
	return fmt.Sprint(*m)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// Flag can be specified multiple times.
func (m *module) Set(value string) error {
	*m = append(*m, value)
	return nil
}

// Supported command-line options.
var bootROM = flag.String("boot", "", "Full path to boot ROM")
var configPath = flag.String("config", "", "Path to custom config file")
var cpuprofile = flag.String("cpuprofile", "", "Write cpu profile to file")
var duration = flag.Uint("cycles", 0, "Stop after executing that many cycles")
var debugModules module
var debugLevel = flag.String("level", "info", "Debug level (-level help for full list)")
var fastBoot = flag.Bool("fastboot", false, "Bypass boot ROM execution")
var gifPath = flag.String("gif", "", "Record gif file")
var monoSound = flag.Bool("mono", false, "Disable stereo panning")
var romPath = flag.String("rom", "", "ROM file to load")
var vSync = flag.Bool("vsync", false, "Force sync to VBlank")
var waitKey = flag.Bool("waitkey", false, "Wait for keypress to start CPU (to help with screen captures)")
var zoomFactor = flag.Uint("zoom", 2, "Zoom factor (default is 2x)")

// Initialize dynamic options.
func init() {
	flag.Var(&debugModules, "debug", "Turn on debug mode for the given module (-debug help for the full list)")
}

// Parse commend-line arguments and return their value in a struct the caller
// can easily pass around.
func Parse() *Options {
	// I like having config files that you can override with command-line
	// parameters.
	flag.Parse()

	// Parse will populate all our variables with either the given or default
	// value, and then we load parameters from the config but avoid overwriting
	// any variable that's been explicitly set by a flag.
	options := Options{
		BootROM:      *bootROM,
		CPUProfile:   *cpuprofile,
		Duration:     *duration,
		DebugModules: debugModules,
		DebugLevel:   *debugLevel,
		FastBoot:     *fastBoot,
		GIFPath:      *gifPath,
		Mono:         *monoSound,
		ROMPath:      *romPath,
		VSync:        *vSync,
		WaitKey:      *waitKey,
		ZoomFactor:   *zoomFactor,
	}

	flagsSet := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		flagsSet[f.Name] = true
	})

	// Other defaults used if there is no config file.

	options.Keymap = DefaultKeymap

	// Always include the default palette as palette 0.
	options.Palettes = append(options.Palettes, DefaultPalette)
	options.PaletteNames = append(options.PaletteNames, "default")

	options.UIBackground = DefaultUIBackground
	options.UIForeground = DefaultUIForeground

	// Use default config if no -config flag was used.
	fullConfigPath := *configPath
	if fullConfigPath == "" {
		createDefaultConfig()
		fullConfigPath = DefaultConfigPath
	}
	fullConfigPath = expandHome(fullConfigPath) // Allow ~ prefix.

	// Load everything else from config, and don't touch values that were set on
	// the command-line.
	options.Update(fullConfigPath, flagsSet)

	return &options
}
