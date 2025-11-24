package main

// The comments below are used by Golang's C pseudo-package, which is used to
// interface with external C code. As we're doing low-level data transfers
// between Go pointers and C, we have to use this special syntax. This is way
// out of scope, but if you're curious, see: https://golang.org/cmd/cgo/
//
// The point is that the C-like comments below will make the Uint8 SDL type
// and our callback function usable as if they were part of a "C" package.

// typedef signed char Uint8;
// void mainLoopCallback(void *userdata, Uint8 *stream, int len);
import "C"

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/lazy-stripes/goholint/apu"
	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ui"
)

// Using module variable instead of pinned pointer passed as userdata to the
// audio callback on account I like it better that way. Apologies to my beloved
// university CS teachers.
var mainUI *ui.UI
var midPoint uint16

// Audio callback function that SDL will call at a regular interval that
// should be roughly <sampling rate> / (<audio buffer size> / <channels>).
//
//export mainLoopCallback
func mainLoopCallback(_ unsafe.Pointer, ptr *C.Uint8, len C.int) {
	// Newer Go lets us cast a C-array pointer to a slice a bit more gracefully.
	n := int(len)
	buffer := unsafe.Slice(ptr, n)

	// FIXME: move loop below to mainUI.FillAudioBuffer(buffer)
	// Tick the emulator as many times as needed to fill the audio buffer.
	for i := 0; i < n; {
		res := mainUI.Tick()

		if res.Play {
			// XXX Tinkering with signedness for now.
			l := uint8(int16(midPoint) + int16(res.Left))
			r := uint8(int16(midPoint) + int16(res.Right))

			buffer[i+0] = C.Uint8(l)
			buffer[i+1] = C.Uint8(r)
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
		sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO | sdl.INIT_EVENTS | sdl.INIT_GAMECONTROLLER | sdl.INIT_JOYSTICK)
		ttf.Init()

		// List currently connected controllers. TODO: detect new ones.
		for i := range sdl.NumJoysticks() {
			sdl.GameControllerOpen(i)
		}

		// Instantiate main UI. Someday I might add extra windows for debugging.
		mainUI = ui.New()

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
		// FIXME: with the latest Ubuntu update with pipewire and stuff, AUDIO_S8
		//        causes horrible, yet very consistent cracking that I believe are
		//        due to some quirk inside SDL when it converts samples, since
		//        OpenAudio returns _U8 as the preferred format. Reverting to U8
		//        seems to not have the crackling but my samples are no longer
		//        centered around zero.
		spec := sdl.AudioSpec{
			Freq:     apu.SamplingRate,
			Format:   sdl.AUDIO_U8,
			Channels: 2,
			Samples:  apu.FramesPerBuffer,
			Callback: sdl.AudioCallback(C.mainLoopCallback),
		}

		var obtained sdl.AudioSpec

		// We're asking SDL to honor our parameters exactly, or fail.
		if err := sdl.OpenAudio(&spec, &obtained); err != nil {
			panic(err)
		}


		midPoint = uint16(obtained.Silence)

		// Start playing sound. Not sure why we un-pause it instead of starting it.
		sdl.PauseAudio(false)
	})

	defer mainUI.Emulator.Stop()

	<-mainUI.QuitChan // Wait for the callback or an action to signal us.

	// Stop invoking our callback.
	sdl.CloseAudio()
}

func main() {
	// Run main function in a separate goroutine so sdl can reserve the UI thread.
	sdl.Main(run)
}
