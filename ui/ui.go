package ui

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/lazy-stripes/goholint/assets"
	"github.com/lazy-stripes/goholint/gameboy"
	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ui/widgets"
	"github.com/lazy-stripes/goholint/ui/widgets/align"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	// Margin is the space in pixels between screen border and UI text.
	Margin = 2
)

// Package-wide logger.
var log = logger.New("ui", "UI-related messages")

type dialog struct {
	widgets.Widget

	title string
}

// UI structure to manage user commands and overlay.
// TODO: move Gameboy inside UI, implement UI.Tick(), see if it works.
type UI struct {
	Emulator *gameboy.GameBoy

	Controls map[options.KeyStroke]Action

	SigINTChan chan os.Signal

	// Send true to this channel to quit the program.
	QuitChan chan bool

	config *options.Options

	paused bool

	// Current palette.
	paletteIndex int

	screen  *widgets.Screen
	dialogs *widgets.Stack
	root    *widgets.Group

	zoomFactor int // From -zoom to compute offsets in various textures

	renderer   *sdl.Renderer
	screenRect *sdl.Rect // Screen dimensions (accounting for zoom factor)
	font       *ttf.Font

	fgColor sdl.Color // Text color
	bgColor sdl.Color // Text outline color

	ticks uint
}

// Return a UI instance given a renderer to create the overlay texture.
func New(config *options.Options) *UI {
	window, err := sdl.CreateWindow("Goholint",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		options.ScreenWidth*int32(config.ZoomFactor),
		options.ScreenHeight*int32(config.ZoomFactor),
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	icon, err := img.LoadRW(assets.WindowIconRW(), false)
	if err != nil {
		panic(err)
	}
	window.SetIcon(icon)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	if config.VSync {
		if err = sdl.GLSetSwapInterval(-1); err != nil {
			log.Infof("Can't set adaptive vsync: %s", sdl.GetError())
			// Try 'just' syncing to vblank then.
			if err = sdl.GLSetSwapInterval(1); err != nil {
				log.Warningf("Can't sync to vblank: %s", sdl.GetError())
			}
		}
	}

	if info, err := renderer.GetInfo(); err == nil {
		log.Info("SDL_RENDERER info:")
		w, h, _ := renderer.GetOutputSize()
		log.Infof("RESOLUTION: %d×%d", w, h)
		log.Infof("SOFTWARE: %t", info.Flags&sdl.RENDERER_SOFTWARE != 0)
		log.Infof("ACCELERATED: %t", info.Flags&sdl.RENDERER_ACCELERATED != 0)
		log.Infof("PRESENTVSYNC: %t", info.Flags&sdl.RENDERER_PRESENTVSYNC != 0)
	}

	// TODO: ui.Font(size) -> *ttf.Font (cached by size).
	font, err := ttf.OpenFontRW(assets.UIFontRW(), 1, int(8*config.ZoomFactor))
	if err != nil {
		panic(err)
	}
	titleFont, err := ttf.OpenFontRW(assets.UIFontRW(), 1, int(12*config.ZoomFactor))
	if err != nil {
		panic(err)
	}

	// Keep computed screen size for drawing.
	screenRect := &sdl.Rect{
		X: 0,
		Y: 0,
		W: options.ScreenWidth * int32(config.ZoomFactor),
		H: options.ScreenHeight * int32(config.ZoomFactor),
	}

	// TODO: try and move all of this to the widgets package to clean up the ui one.
	// Colors from config.
	fg := sdl.Color{
		R: config.UIForeground.R,
		G: config.UIForeground.G,
		B: config.UIForeground.B,
		A: config.UIForeground.A,
	}
	bg := sdl.Color{
		R: config.UIBackground.R,
		G: config.UIBackground.G,
		B: config.UIBackground.B,
		A: config.UIBackground.A,
	}

	// By default, use BG color at zero transparency for clearing widgets. It
	// makes the outline of labels blend better.
	background := bg
	background.A = 0

	// Store default widget properties in the widget package namespace. This
	// will be copied to every new widget.
	// FIXME: should that be set directly from options package?
	widgets.DefaultProperties = widgets.Properties{
		Font:            font,
		TitleFont:       titleFont,
		BgColor:         bg,
		FgColor:         fg,
		HorizontalAlign: align.Center,
		VerticalAlign:   align.Middle,
		Border:          1,
		BorderColor:     sdl.Color{0xff, 0x00, 0x00, 0xff},
		//Background:      sdl.Color{0xff, 0xff, 0xff, 0x00},
		Background: background,
		Zoom:       int(config.ZoomFactor),
	}

	widgets.Init(renderer)

	// Screen widget the emulator will write into via the PixelWriter interface.
	gbScreen := widgets.NewScreen(screenRect, config)
	emulator := gameboy.New(gbScreen, config)

	ui := &UI{
		config:     config,
		QuitChan:   make(chan bool),
		SigINTChan: make(chan os.Signal, 1),
		Emulator:   emulator,
		renderer:   renderer,
		screenRect: screenRect,
		font:       font,
		zoomFactor: int(config.ZoomFactor),
		fgColor:    fg,
		bgColor:    bg,
		dialogs:    widgets.NewStack(screenRect, nil),
		root:       widgets.NewGroup(screenRect, nil), // WIP, still not sure how I'll organize this. I still like bg/fg.
		screen:     gbScreen,
	}

	// The UI should primarily show the emulator's screen, with some menu or
	// special widget on top whenever emulation is paused.
	ui.root.Add(gbScreen)
	ui.root.Add(ui.dialogs)
	ui.dialogs.SetVisible(false)

	// Create menu stack with extra widgets. Those will only be shown when the
	// emulator is paused.
	ui.buildMenu()

	ui.SetControls(config.Keymap)

	go ui.handleSIGINT()

	return ui
}

// handleSIGINT prints and dumps debug data on CTRL+C. This should be run as a
// goroutine at startup after creating UI.
func (u *UI) handleSIGINT() {
	// Wait for signal, quit cleanly with potential extra debug info if needed.
	signal.Notify(u.SigINTChan, os.Interrupt)

	<-u.SigINTChan
	fmt.Println("\nTerminated...")

	// TODO: quit-time cleanup in gb, ui, etc.
	//gb.Display.Close()

	// TODO: only dump RAM/VRAM/Other if requested in parameters.
	fmt.Print(u.Emulator.CPU)
	fmt.Print(u.Emulator.PPU)
	u.Emulator.CPU.DumpMemory()

	// Force stopping CPU profiling.
	pprof.StopCPUProfile()

	// FIXME: quit cleanly
	os.Exit(-1)
}

// TODO: FillAudioBuffer
func (u *UI) Tick() (res gameboy.TickResult) {
	u.ticks++
	if u.ticks%1000 == 0 {
		sdl.Do(u.ProcessEvents)
	}

	// FIXME: pause on vblank. Should be using gb.Tick() return here anyway.
	if !u.paused {
		defer u.Emulator.Recover()
		return u.Emulator.Tick()
	}
	return gameboy.TickResult{Play: true} // Always return silence. TODO: play samples of our UI SFXes...es here!
}

func (u *UI) buildMenu() {
	mainMenu := widgets.NewMenu(u.screenRect)

	// Main menu should always be the Stack's bottom widget. It's shown first.
	u.dialogs.Add(mainMenu)

	mainMenu.AddChoice("Resume", u.Hide)

	// FIXME: just do this manually, like addChoice("Input", widgets.NewInput)
	//for i, d := range u.dialogs {
	//	// Each choice outside of the defaults (resume or quit) will just
	//	// show the corresponding widget in the root stack.
	//	mainMenu.AddChoice(d.title, u.dialogs.Show(i+1)) // Main menu is entry 0
	//	u.dialogs.Add(d)
	//}

	mainMenu.AddChoice("Quit", func() { u.QuitChan <- true })
	mainMenu.Select(0) // highlight first entry
}

func (u *UI) Show() {
	u.paused = true
	u.screen.Pause()
	u.dialogs.SetVisible(true)
	u.Repaint()
}

func (u *UI) Hide() {
	u.paused = false
	u.screen.Unpause()
	u.dialogs.SetVisible(false)
	u.Repaint()
}

// SetControls validates and sets the given control map for the emulator's UI.
func (u *UI) SetControls(keymap options.Keymap) (err error) {
	// Intermediate mapping between labels and actual actions.
	// TODO: re-add all the stuff previously done on the gb side (screenshots, gif, etc)
	actions := map[string]Action{
		// High-level actions.
		"home": u.Home,
		"quit": u.Quit,

		// TODO: restrict those to when we're not paused. Maybe I could have
		//       subcontrols for widgets? Might need a bespoke root widget type.
		"nextpalette":     u.NextPalette,
		"previouspalette": u.PreviousPalette,
		"screenshot":      u.Screenshot,
		//		"recordgif":       g.StartStopRecord,
		"togglevoice1": u.ToggleVoice1,
		"togglevoice2": u.ToggleVoice2,
		"togglevoice3": u.ToggleVoice3,
		"togglevoice4": u.ToggleVoice4,

		// Button presses that could either be handled by GB or UI.
		"up":     u.ButtonPressAction(widgets.ButtonUp, u.Emulator.JoypadUp),
		"down":   u.ButtonPressAction(widgets.ButtonDown, u.Emulator.JoypadDown),
		"left":   u.ButtonPressAction(widgets.ButtonLeft, u.Emulator.JoypadLeft),
		"right":  u.ButtonPressAction(widgets.ButtonRight, u.Emulator.JoypadRight),
		"a":      u.ButtonPressAction(widgets.ButtonA, u.Emulator.JoypadA),
		"b":      u.ButtonPressAction(widgets.ButtonB, u.Emulator.JoypadB),
		"select": u.ButtonPressAction(widgets.ButtonSelect, u.Emulator.JoypadSelect),
		"start":  u.ButtonPressAction(widgets.ButtonStart, u.Emulator.JoypadStart),
	}

	u.Controls = make(map[options.KeyStroke]Action)
	for label, keyStroke := range keymap {
		u.Controls[keyStroke] = actions[label]
	}
	return nil
}

func (u *UI) ButtonPressAction(e widgets.Event, gbAction gameboy.Action) Action {
	// Convert keystroke into simpler one-shot widget event. We only care about
	// given event type to tell if a key was pressed.
	return func(eventType uint32) {
		if !u.paused {
			gbAction(eventType == sdl.KEYDOWN)
		} else if u.root != nil && eventType == sdl.KEYDOWN {
			u.root.ProcessEvent(e)
			u.Repaint()
		}
	}
}

func (u *UI) ProcessEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		eventType := event.GetType()

		switch eventType {
		// Button presses and UI keys
		case sdl.KEYDOWN, sdl.KEYUP:
			keyEvent := event.(*sdl.KeyboardEvent)
			keyStroke := options.KeyStroke{
				Code: keyEvent.Keysym.Sym,
				Mod:  sdl.Keymod(keyEvent.Keysym.Mod & options.ModMask),
			}

			if action := u.Controls[keyStroke]; action != nil {
				action(eventType)
			} else {
				if eventType == sdl.KEYDOWN {
					log.Infof("unknown key code: 0x%x", keyStroke.Code)
					log.Infof("        modifier: 0x%x", sdl.GetModState())
				}
			}

		// Window redraw events
		case sdl.WINDOWEVENT:
			u.Repaint()

		// Window-closing event
		case sdl.QUIT:
			u.QuitChan <- true
		}
	}
}

func (u *UI) Repaint() {
	// At this point, we can pretty much just render the root widget.
	texture := u.root.Texture()
	u.renderer.SetRenderTarget(nil)
	u.renderer.Copy(texture, nil, nil)

	// Debug stuff
	var gridSize int32 = 8
	if log.Enabled() && logger.Level >= logger.Debug {
		u.renderer.SetDrawColor(0xff, 0x00, 0x00, 0xff)
		for x := int32(0); x < u.screenRect.W; x += gridSize * int32(u.config.ZoomFactor) {
			u.renderer.DrawLine(x, 0, x, u.screenRect.H)
		}
		for y := int32(0); y < u.screenRect.H; y += gridSize * int32(u.config.ZoomFactor) {
			u.renderer.DrawLine(0, y, u.screenRect.W, y)
		}
	}

	u.renderer.Present()
}
