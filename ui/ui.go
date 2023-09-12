package ui

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unsafe"

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
	Enabled bool

	Controls map[options.KeyStroke]Action

	// Send true to this channel to quit the program.
	QuitChan chan bool

	message string // Temporary text on timer
	text    string // Permanent text

	// TODO: root *Widget

	renderer     *sdl.Renderer
	background   *sdl.Texture // UI texture (background)
	texture      *sdl.Texture // UI texture (foreground)
	screen       *sdl.Texture // GameBoy screen texture
	screenBuffer []byte       // GameBoy screen pixels
	screenRect   *sdl.Rect
	font         *ttf.Font
	zoomFactor   int // From -zoom to compute offsets in various textures

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

	// We set background to full UI texture size for higher-def blurring when we get there.
	background, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_TARGET,
		options.ScreenWidth*int32(config.ZoomFactor),
		options.ScreenHeight*int32(config.ZoomFactor))
	if err != nil {
		panic(err)
	}

	// Background transparency.
	texture.SetBlendMode(sdl.BLENDMODE_BLEND)

	// Keep computed screen size for drawing.
	screenRect := &sdl.Rect{
		X: 0,
		Y: 0,
		W: options.ScreenWidth * int32(config.ZoomFactor),
		H: options.ScreenHeight * int32(config.ZoomFactor),
	}

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
		QuitChan:   make(chan bool),
		background: background,
		texture:    texture, // TODO: rename to foreground?
		renderer:   renderer,
		screenRect: screenRect,
		font:       font,
		zoomFactor: int(config.ZoomFactor),
		fg:         fg,
		bg:         bg,
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

// ScreenBuffer creates a new SDL texture suitable to use for the emulator's
// screen, and a pixel buffer that it returns, which the screen should write
// into. This lets us do funny stuff with the GameBoy display's pixels that
// we couldn't easily do if we only had access to a texture.
func (u *UI) ScreenBuffer() (buffer []byte) {
	texture, err := u.renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STATIC,
		options.ScreenWidth,
		options.ScreenHeight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create screen texture: %s\n", err)
		return nil // TODO: result, err
	}

	// Save texture for repaints.
	u.screen = texture

	// Return buffer for the screen to render into.
	u.screenBuffer = make([]byte, options.ScreenWidth*options.ScreenHeight*4)

	return u.screenBuffer
}

// SetControls validates and sets the given control map for the emulator's UI.
func (u *UI) SetControls(keymap options.Keymap) (err error) {
	// Intermediate mapping between labels and actual actions.
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
	u.renderer.SetRenderTarget(u.texture)
	u.renderer.SetDrawColor(0, 0, 0, 0)
	u.renderer.Clear()
	u.renderer.SetRenderTarget(nil)

	// TODO: if enabled, use frozen/blurred screen as background.
	// Display an overlay when emulation is stopped.
	if u.Enabled {
		// TODO: widgets. Gameboy screen may well be one too!
		//       Widgets are probably gonna need a renderer.
		//u.root.Texture()

		// FIXME: might be easier with a dedicated background texture mimicking the frozen screen.
		u.renderer.SetRenderTarget(u.texture)
		u.renderer.SetDrawColor(0, 0, 0, 0x8f)
		u.renderer.FillRect(u.screenRect)
		u.renderer.SetRenderTarget(nil)
		u.renderer.Copy(u.texture, nil, nil)

		// TODO: render root widget to foreground texture. Those words scare me.

	} else {

		// Gameboy screen in the background.
		// SDL bindings used to accept a slice but no longer do as of 0.4.33.
		rawPixels := unsafe.Pointer(&u.screenBuffer[0])
		u.screen.Update(nil, rawPixels, options.ScreenWidth*4)
		u.renderer.Copy(u.screen, nil, nil)

		// Messages. I'm leaving them in the background for now.
		if u.text != "" || u.message != "" {
			u.repaintText()
			u.renderer.Copy(u.texture, nil, nil)
		}
	}

	u.renderer.Present()
}

// Refresh UI texture with permanent text and current message (if any).
func (u *UI) repaintText() {
	// Reset texture.
	u.renderer.SetRenderTarget(u.texture)
	u.renderer.SetDrawColor(0, 0, 0, 0)

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

	u.renderer.SetRenderTarget(nil)
}

// Refresh UI texture with permanent text and current message (if any).
func (u *UI) renderText(text string, row int) {
	// Instantiate text with an outline effect. There's probably an easier way.
	outlineWidth := int(u.zoomFactor)
	u.font.SetOutline(outlineWidth)
	outline, _ := u.font.RenderUTF8Solid(text, u.bg)
	u.font.SetOutline(0)
	msg, _ := u.font.RenderUTF8Solid(text, u.fg)

	// Position vertically. Bottom row is row number 1.
	_, _, _, h, _ := u.texture.Query()
	y := h - int32((u.font.Height())*row) - Margin // TODO: FontSize config

	// Add margin between successive rows.
	if row > 1 {
		y -= Margin * int32(outlineWidth) * 2
	}

	outlineTexture, _ := u.renderer.CreateTextureFromSurface(outline)
	u.renderer.Copy(outlineTexture,
		nil,
		&sdl.Rect{
			X: Margin,
			Y: y - int32(u.zoomFactor),
			W: outline.W,
			H: outline.H,
		})

	msgTexture, _ := u.renderer.CreateTextureFromSurface(msg)
	u.renderer.Copy(msgTexture,
		nil,
		&sdl.Rect{
			X: Margin + int32(u.zoomFactor),
			Y: y,
			W: msg.W,
			H: msg.H,
		})
}

// Set permanent text (useful for persistent UI). Call with empty string to
// clear.
func (u *UI) Text(text string) {
	u.text = text
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
}
