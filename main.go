package main

// The comments below are used by Golang's C pseudo-package, which is used to
// interface with external C code. As we're doing low-level data transfers
// between Go pointers and C, we have to use this special syntax. This is way
// out of scope, but if you're curious, see: https://golang.org/cmd/cgo/
//
// The point is that the C-like comments below will make the Uint8 SDL type
// and our callback function usable as if they were part of a "C" package.

// typedef signed char Int8;
// void mainLoopCallback(void *userdata, Int8 *stream, int len);
import "C"

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime/pprof"
	"strings"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/lazy-stripes/goholint/apu"
	"github.com/lazy-stripes/goholint/gameboy"
	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ui"
)

// TODO: minimal (like, REALLY minimal) GUI. And clean all of this up.

var quit chan bool // Used by the callback to tell the main function to quit.

func init() {
	quit = make(chan bool)
}

// I have given up on trying to pass this as userdata to SDL.
// TODO: try runtime.Pinner whenever we upgrade to go 1.21

// TODO: maybe move all that main code to a goholint package?

var goholint struct {
	gb    *gameboy.GameBoy
	ui    *ui.UI
	ticks uint64
}

var mainUI *ui.UI

// Audio callback function that SDL will call at a regular interval that
// should be roughly <sampling rate> / (<audio buffer size> / <channels>).
//
//export mainLoopCallback
func mainLoopCallback(data unsafe.Pointer, buf *C.Int8, len C.int) {
	// We've reached the limits of the Go bindings. I might try and see if it's
	// any cleaner when using a push approach rather than a callback.
	n := int(len)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(buf)), Len: n, Cap: n}
	buffer := *(*[]C.Int8)(unsafe.Pointer(&hdr))

	// FIXME: move loop below to mainUI.FillAudioBuffer(buffer)
	// Tick the emulator as many times as needed to fill the audio buffer.
	for i := 0; i < n; {
		res := mainUI.Tick()

		if res.Play {
			buffer[i] = C.Int8(res.Left)
			buffer[i+1] = C.Int8(res.Right)
			i += 2
		}

		if res.VBlank {
			sdl.Do(mainUI.Repaint)
		}
	}
}

// Separate function to forcefully run in the main thread.
func run() {
	args := options.Parse()

	if args.DebugLevel == "help" {
		logger.HelpLevels()
		os.Exit(0)
	}

	level, ok := logger.Levels[strings.ToLower(args.DebugLevel)]
	if ok {
		logger.Level = level
	} else {
		log.Fatal("unknown log level ", args.DebugLevel)
	}

	for _, m := range args.DebugModules {
		// List available modules if requested.
		if m == "help" {
			logger.Help()
			os.Exit(0)
		}

		// TODO: error if module OR submodule is not registered.
		logger.Enabled[m] = true
	}

	if args.CPUProfile != "" {
		f, err := os.Create(args.CPUProfile)
		if err != nil {
			log.Fatal(err)
		}
		if err = pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()

		log.Println("CPU profiling written to: ", args.CPUProfile)
	}

	// Execute all SDL operations in the main thread.
	sdl.Do(func() {
		sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO | sdl.INIT_EVENTS)
		ttf.Init()

		// Instantiate main UI. Someday I might add extra windows for debugging.
		mainUI = ui.New(args)

		// Wait for keypress if requested, so obs has time to capture window.
		// Less useful now that we have -gif flag.
		if args.WaitKey {
			fmt.Print("Press 'Enter' to start...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}

		// Add CPU-specific context to debug output.
		logger.Context = mainUI.Emulator.CPU.Context
		//logger.Context = func() string { return fmt.Sprintf("%s\n%s\n> ", gb.CPU, gb.PPU) } // TEMPORARY

		// An AudioSpec structure containing our parameters. After calling
		// OpenAudio, it will also contain some values initialized by SDL itself,
		// such as the audio buffer size.
		spec := sdl.AudioSpec{
			Freq:     apu.SamplingRate,
			Format:   sdl.AUDIO_S8,
			Channels: 2,
			Samples:  apu.FramesPerBuffer,
			Callback: sdl.AudioCallback(C.mainLoopCallback),
		}

		// We're asking SDL to honor our parameters exactly, or fail.
		if err := sdl.OpenAudio(&spec, nil); err != nil {
			panic(err)
		}

		// Start playing sound. Not sure why we un-pause it instead of starting it.
		sdl.PauseAudio(false)
	})

	defer mainUI.Emulator.Stop()

	<-mainUI.QuitChan // Wait for the callback or an action to signal us.

	sdl.CloseAudio()
}

func main() {
	// Run main function in a separate goroutine so sdl can reserve the UI thread.
	sdl.Main(run)
}
