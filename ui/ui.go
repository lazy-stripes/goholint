package ui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lazy-stripes/goholint/assets"
	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/options"

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

// UI structure to manage user commands and overlay.
type UI struct {
	Enabled  bool
	Renderer *sdl.Renderer

	Screen *sdl.Texture // Gameboy screen texture

	// Send true to this channel to quit the program.
	QuitChan chan bool

	message string // Temporary text on timer
	text    string // Permanent text

	texture *sdl.Texture // UI texture

	font     *ttf.Font
	fontZoom uint

	fg sdl.Color // TODO: make it configurable
	bg sdl.Color // TODO: make it configurable

	msgTimer *time.Timer
}

// Return a UI instance given a renderer to create the overlay texture.
func New(config *options.Options) *UI {
	window, err := sdl.CreateWindow("Goholint",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		options.ScreenWidth*int32(config.ZoomFactor),
		options.ScreenHeight*int32(config.ZoomFactor),
		sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return nil // TODO: result, err
	}

	icon, err := img.LoadRW(assets.WindowIcon, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load icon: %s\n", err)
	} else {
		window.SetIcon(icon)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return nil // TODO: result, err
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
		log.Infof("SOFTWARE: %t", info.Flags&sdl.RENDERER_SOFTWARE != 0)
		log.Infof("ACCELERATED: %t", info.Flags&sdl.RENDERER_ACCELERATED != 0)
		log.Infof("PRESENTVSYNC: %t", info.Flags&sdl.RENDERER_PRESENTVSYNC != 0)
	}

	font, err := ttf.OpenFontRW(assets.UIFont, 1, int(8*config.ZoomFactor))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		return nil // TODO: result, err
	}

	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		options.ScreenWidth*int32(config.ZoomFactor),
		options.ScreenHeight*int32(config.ZoomFactor))
	if err != nil {
		font.Close()
		fmt.Fprintf(os.Stderr, "Failed to create UI texture: %s\n", err)
		return nil // TODO: result, err
	}

	// Scale font up with screen size.
	fontZoom := config.ZoomFactor // TODO: smaller fontZoom for higher zoom.

	// Background transparency.
	texture.SetBlendMode(sdl.BLENDMODE_BLEND)

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

	ui := &UI{
		texture:  texture,
		QuitChan:   make(chan bool),
		Renderer: renderer,
		font:     font,
		fontZoom: fontZoom,
		fg:       fg,
		bg:       bg,
	}

	// TODO: allow several subsystems with .AddUI(scanner). We'll need a complex
	// interface. I can't wait.

	ui.SetControls(config.Keymap)

	return ui
}

func (u *UI) Show() {
	// TODO: background blur, top menu
	u.Enabled = true
	u.Repaint()
}

func (u *UI) Hide() {
	u.Enabled = false
}

// ScreenTexture returns a new SDL texture suitable to use for the emulator's
// screen, and stores it internally to use it during repaints.
func (u *UI) ScreenTexture() (texture *sdl.Texture) {
	texture, err := u.Renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STATIC,
		options.ScreenWidth,
		options.ScreenHeight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create screen texture: %s\n", err)
		return nil // TODO: result, err
	}

	// Save texture for repaints.
	u.Screen = texture

	return texture
}

// SetControls validates and sets the given control map for the emulator's UI.
func (u *UI) SetControls(keymap options.Keymap) (err error) {
	// Intermediate mapping between labels and actual actions. This feels
	// unnecessarily complicated, but should make sense when I start translating
	// these from a config file. I hope.
	actions := map[string]Action{
		"quit": u.Quit,
		"home": u.Home,
	}

	u.Controls = make(map[options.KeyStroke]Action)
	for label, keyStroke := range keymap {
		u.Controls[keyStroke] = actions[label]
	}
	return nil
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
			// TODO: home menu actions
			// TODO: now starts the propagation fun. How to send events down the widget tree?
			if action := u.Controls[keyStroke]; action != nil {
				action(eventType)
			} else {
				if eventType == sdl.KEYDOWN {
					log.Infof("unknown key code: 0x%x", keyStroke.Code)
					log.Infof("        modifier: 0x%x", sdl.GetModState())
				}
			}

		// Window-closing event
		case sdl.QUIT:
			u.QuitChan <- true
		}
	}
}

func (u *UI) Repaint() {
	// Gameboy screen.
	if u.Screen != nil {
		u.Renderer.Copy(u.Screen, nil, nil)
	}

	// UI overlay.
	if u.text != "" || u.message != "" {
		//u.Texture.SetBlendMode(sdl.BLENDMODE_ADD)
		u.Renderer.Copy(u.texture, nil, nil)
	}
	u.Renderer.Present()
}

// Refresh UI texture with permanent text and current message (if any).
func (u *UI) repaintText() {
	// Reset texture.
	u.Renderer.SetRenderTarget(u.texture)
	u.Renderer.SetDrawColor(0, 0, 0, 0)
	u.Renderer.Clear()

	row := 1
	if u.text != "" {
		u.renderText(u.text, row)
		row++
	}

	// TODO: stack messages
	if u.message != "" {
		// Allow messages to have several lines. However, we need to iterate in
		// reverse as we render text from the bottom up.
		lines := strings.Split(u.message, "\n")
		for i := range lines {
			line := lines[len(lines)-1-i]
			u.renderText(line, row)
			row++
		}
	}

	u.Renderer.SetRenderTarget(nil)
}

// Refresh UI texture with permanent text and current message (if any).
func (u *UI) renderText(text string, row int) {
	// Instantiate text with an outline effect. There's probably an easier way.
	u.font.SetOutline(int(u.fontZoom))
	outline, _ := u.font.RenderUTF8Solid(text, u.bg)
	u.font.SetOutline(0)
	msg, _ := u.font.RenderUTF8Solid(text, u.fg)

	// Position vertically. Bottom row is row number 1.
	_, _, _, h, _ := u.texture.Query()
	y := h - int32(u.font.Height()*row) - Margin // TODO: FontSize config

	outlineTexture, _ := u.Renderer.CreateTextureFromSurface(outline)
	u.Renderer.Copy(outlineTexture, nil, &sdl.Rect{X: Margin, Y: y - int32(u.fontZoom), W: outline.W, H: outline.H})

	msgTexture, _ := u.Renderer.CreateTextureFromSurface(msg)
	u.Renderer.Copy(msgTexture, nil, &sdl.Rect{X: Margin + int32(u.fontZoom), Y: y, W: msg.W, H: msg.H})
}

// Set permanent text (useful for persistent UI). Call with empty string to
// clear.
func (u *UI) Text(text string) {
	u.text = text
	u.repaintText()
}

// Clear temporary message and repaint texture.
func (u *UI) clearMessage() {
	// Make sure to execute in the UI thread in case we were called from a
	// timer thread.
	u.message = ""
	sdl.Do(u.Repaint)
}

// Message creates a new UI texture with the given message, enables UI and
// starts a timer that will hide the UI when it's done. Takes a text string and
// a duration (in seconds).
func (u *UI) Message(text string, duration time.Duration) {
	// Stop reset timer, a new one will be started.
	// TODO: stack messages
	if u.msgTimer != nil {
		u.msgTimer.Stop()
	}
	u.message = text
	u.msgTimer = time.AfterFunc(time.Second*duration, u.clearMessage)
	u.repaintText()
}
