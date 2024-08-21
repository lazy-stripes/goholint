package ui

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"
	"time"
	"unsafe"

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
	emulator *gameboy.GameBoy

	paused bool

	Controls map[options.KeyStroke]Action

	// Send true to this channel to quit the program.
	QuitChan chan bool

	msgTimer *time.Timer // Timer for clearing messages
	message  string      // Temporary text on timer
	text     string      // Permanent text

	dialogs []dialog       // Subwidgets for various systems.
	root    *widgets.Stack // WIP

	zoomFactor int // From -zoom to compute offsets in various textures

	renderer   *sdl.Renderer
	background *sdl.Texture // UI texture (background)
	texture    *sdl.Texture // UI texture (foreground)
	screenRect *sdl.Rect    // Screen dimensions (accounting for zoom factor)
	font       *ttf.Font

	gbScreen       *sdl.Texture // GameBoy screen texture
	gbScreenBuffer []byte       // GameBoy screen pixels

	fgColor sdl.Color // Text color
	bgColor sdl.Color // Text outline color

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

	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		options.ScreenWidth*int32(config.ZoomFactor),
		options.ScreenHeight*int32(config.ZoomFactor))
	if err != nil {
		panic(err)
	}

	// We set background to full UI texture size for higher-def blurring.
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

	// Store default widget properties in the widget package namespace. This
	// will be copied to every new widget.
	widgets.DefaultProperties = widgets.Properties{
		Font:            font,
		TitleFont:       titleFont,
		BgColor:         bg,
		FgColor:         fg,
		HorizontalAlign: align.Center,
		VerticalAlign:   align.Middle,
		//Border:     1,
		Background: sdl.Color{0x20, 0x7f, 0x20, 0x7f},
		Zoom:       int(config.ZoomFactor),
	}

	widgets.Init(renderer)

	// Screen widget the emulator will write into via the PixelWriter interface.
	gbScreen := widgets.NewScreen(screenRect, config)
	emulator := gameboy.New(config, gbScreen)

	ui := &UI{
		QuitChan:   make(chan bool),
		emulator:   emulator,
		background: background,
		texture:    texture, // TODO: rename to foreground?
		renderer:   renderer,
		screenRect: screenRect,
		font:       font,
		zoomFactor: int(config.ZoomFactor),
		fgColor:    fg,
		bgColor:    bg,
		//root:       widgets.NewStack(screenRect, nil),
		root: widgets.NewStack(screenRect, []widgets.Widget{gbScreen}),
	}

	// FIXME: more helper functions.
	//ui.addDialog(dialog{
	//	widgets.NewInput(screenRect, 5, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", ui.Hide),
	//	"Input",
	//})

	//ui.buildMenu()

	ui.SetControls(config.Keymap)

	return ui
}

func (u *UI) Tick() (res gameboy.TickResult) {
	// WIP Only tick emulator for now. We'll do the "pause to menu" again from scratch. Weeee.
	defer u.emulator.Recover()
	return u.emulator.Tick()
	//return TickResult{Play: true} // Always return silence. TODO: play samples of our UI SFXes...es here!
}

func (u *UI) addDialog(d dialog) {
	u.dialogs = append(u.dialogs, d)
}

func (u *UI) showDialog(idx int) func() {
	return func() { u.root.Show(uint(idx)) }
}

func (u *UI) buildMenu() {
	mainMenu := widgets.NewMenu(u.screenRect)

	// Main menu should always be the Stack's bottom widget. It's shown first.
	u.root.Add(mainMenu)
	u.root.Show(0)

	mainMenu.AddChoice("Resume", u.Hide)

	for i, d := range u.dialogs {
		// Each choice outside of the defaults (resume or quit) will just
		// show the corresponding widget in the root stack.
		mainMenu.AddChoice(d.title, u.showDialog(i+1)) // Main menu is entry 0
		u.root.Add(d)
	}

	mainMenu.AddChoice("Quit", func() { u.QuitChan <- true })
	mainMenu.Select(0) // highlight first entry
}

func (u *UI) Show() {
	u.paused = true
	u.freezeBackground()
	u.Repaint()
}

func (u *UI) Hide() {
	u.paused = false
}

func averagePixels(pixels []color.RGBA) (avg color.RGBA) {
	var sumR, sumG, sumB int
	for _, pixel := range pixels {
		sumR += int(pixel.R)
		sumG += int(pixel.G)
		sumB += int(pixel.B)
	}

	avg = color.RGBA{
		uint8(sumR / len(pixels)),
		uint8(sumG / len(pixels)),
		uint8(sumB / len(pixels)),
		0xff,
	}

	return avg
}

// blur returns a copy of the image after applying the box blur algorithm to it.
// Image has to be at least 2px×2px, or you will have a bad time.
func blur(img *image.RGBA) (blurred *image.RGBA) {
	blurred = image.NewRGBA(img.Bounds())

	// Apply blur to inner pixels (radius is 1 pixel).
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	for x := 1; x < w-1; x++ {
		for y := 1; y < h-1; y++ {
			neighbors := []color.RGBA{
				img.RGBAAt(x-1, y+1), // Top left
				img.RGBAAt(x+0, y+1), // Top center
				img.RGBAAt(x+1, y+1), // Top right
				img.RGBAAt(x-1, y+0), // Mid left
				img.RGBAAt(x+0, y+0), // Current pixel
				img.RGBAAt(x+1, y+0), // Mid right
				img.RGBAAt(x-1, y-1), // Low left
				img.RGBAAt(x+0, y-1), // Low center
				img.RGBAAt(x+1, y-1), // Low right
			}

			avg := averagePixels(neighbors)
			blurred.SetRGBA(x, y, avg)

			// Duplicate left column of blurred pixels.
			if x == 1 {
				blurred.SetRGBA(0, y, avg)
			}

			// Duplicate right column of blurred pixels.
			if x == w-2 {
				blurred.SetRGBA(w-1, y, avg)
			}

			// Duplicate top row of blurred pixels.
			if y == 1 {
				blurred.SetRGBA(x, 0, avg)
			}

			// Duplicate bottom row of blurred pixels.
			if y == h-2 {
				blurred.SetRGBA(x, h-1, avg)
			}
		}
	}

	// Copy corner pixels.
	blurred.SetRGBA(0, 0, img.RGBAAt(0, 0))
	blurred.SetRGBA(w, 0, img.RGBAAt(w, 0))
	blurred.SetRGBA(0, h, img.RGBAAt(0, h))
	blurred.SetRGBA(w, h, img.RGBAAt(w, h))

	return blurred
}

// freezeBackground takes a copy of the current GameBoy screen and turns it to
// blurred greyscale for use as a background in the main UI.
func (u *UI) freezeBackground() {
	// We need the screen buffer here. This tightly couples UI and screen.
	// I'm sure it's fine.

	// Dimensions of UI screen.
	_, _, w, h, _ := u.background.Query()
	width := int(w)
	height := int(h)

	// Intermediate image for easier blurring.
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Map source offset (in 160×144 space) to the current UI pixel.
			srcX := x / u.zoomFactor
			srcY := y / u.zoomFactor
			srcOffset := (srcY * options.ScreenWidth * 4) + (srcX * 4)

			// Extract RGB, compute greyscale, strore in work image.
			r := u.gbScreenBuffer[srcOffset+0]
			g := u.gbScreenBuffer[srcOffset+1]
			b := u.gbScreenBuffer[srcOffset+2]
			a := u.gbScreenBuffer[srcOffset+3]
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			grey := uint8(lum)

			img.SetRGBA(x, y, color.RGBA{grey, grey, grey, a})
		}
	}
	// Blur the background. Apply enough times for sufficient effect.
	// TODO: ... I could make the iterations and overlay configurable I guess?
	img = blur(blur(blur(img)))
	rawPixels := unsafe.Pointer(&img.Pix[0])
	u.background.Update(nil, rawPixels, width*4)
}

// ScreenBuffer creates a new SDL texture suitable to use for the emulator's
// screen, and a pixel buffer that it returns, which the PPU should write into.
// This lets us do funny stuff with the GameBoy display's pixels in the UI that
// we couldn't easily do if we only had access to a texture.
func (u *UI) ScreenBuffer() (buffer []byte) {
	texture, err := u.renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STATIC,
		options.ScreenWidth,
		options.ScreenHeight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create screen texture: %s\n", err)
		return nil // XXX Honestly, at this point we might as well panic()
	}

	// Save texture for repaints.
	u.gbScreen = texture

	// Return buffer for the screen to render into.
	u.gbScreenBuffer = make([]byte, options.ScreenWidth*options.ScreenHeight*4)

	return u.gbScreenBuffer
}

// SetControls validates and sets the given control map for the emulator's UI.
func (u *UI) SetControls(keymap options.Keymap) (err error) {
	// Intermediate mapping between labels and actual actions.
	actions := map[string]Action{
		"quit":   u.Quit,
		"home":   u.Home,
		"up":     u.propagate(widgets.ButtonUp),
		"down":   u.propagate(widgets.ButtonDown),
		"left":   u.propagate(widgets.ButtonLeft),
		"right":  u.propagate(widgets.ButtonRight),
		"a":      u.propagate(widgets.ButtonA),
		"b":      u.propagate(widgets.ButtonB),
		"select": u.propagate(widgets.ButtonSelect),
		"start":  u.propagate(widgets.ButtonStart),
	}

	u.Controls = make(map[options.KeyStroke]Action)
	for label, keyStroke := range keymap {
		u.Controls[keyStroke] = actions[label]
	}
	return nil
}

func (u *UI) propagate(e widgets.Event) Action {
	// Convert keystroke into simpler one-shot widget event. We only care about
	// given event type to tell if a key was pressed.
	return func(eventType uint32) {
		if u.root != nil && eventType == sdl.KEYDOWN {
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
	// Reset background texture.
	u.renderer.SetRenderTarget(u.texture)
	u.renderer.SetDrawColor(0, 0, 0, 0)
	u.renderer.Clear()
	u.renderer.SetRenderTarget(nil)

	// If UI is enabled, use frozen/blurred screen as background. Otherwise,
	// render GameBoy screen and overlay messages.
	if true || u.paused {
		// TODO: widgets. Gameboy screen may well be one too!
		//       Widgets are probably gonna need a renderer.

		// Render blurred greyscale background to window.
		u.renderer.SetRenderTarget(nil)
		u.renderer.Copy(u.background, nil, nil)

		// Render overlay on foreground texture and copy it to window.
		u.renderer.SetRenderTarget(u.texture)
		u.renderer.SetDrawColor(0xcc, 0xcc, 0xcc, 0x90)
		u.renderer.FillRect(u.screenRect)

		u.renderer.SetRenderTarget(nil)
		u.texture.SetBlendMode(sdl.BLENDMODE_BLEND)
		u.renderer.Copy(u.texture, nil, nil)

		// Retrieve texture for root widget and copy it to window.
		root := u.root.Texture()

		u.renderer.SetRenderTarget(nil)
		root.SetBlendMode(sdl.BLENDMODE_BLEND)
		u.renderer.Copy(root, nil, nil)

		// Debug stuff
		if log.Enabled() && logger.Level >= logger.Debug {
			u.renderer.SetDrawColor(0xff, 0x00, 0x00, 0xff)
			for x := int32(0); x < u.screenRect.W; x += 8 {
				u.renderer.DrawLine(x, 0, x, u.screenRect.H)
			}
			for y := int32(0); y < u.screenRect.H; y += 8 {
				u.renderer.DrawLine(0, y, u.screenRect.W, y)
			}
		}

	} else {
		// GameBoy screen and text/message.

		// SDL bindings used to accept a slice but no longer do as of 0.4.33.
		rawPixels := unsafe.Pointer(&u.gbScreenBuffer[0])
		u.gbScreen.Update(nil, rawPixels, options.ScreenWidth*4)

		u.renderer.SetRenderTarget(nil)
		u.renderer.Copy(u.gbScreen, nil, nil)

		// Messages. I'm leaving them in the background for now.
		if u.text != "" || u.message != "" {
			u.repaintText() // FIXME: make it clearer that it's rendering to u.texture.
			// I *really* need some more generic function to render text anyway.
			u.renderer.SetRenderTarget(nil)
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
}

// Refresh UI texture with permanent text and current message (if any).
// TODO: widgets.Label
func (u *UI) renderText(s string, row int) {
	// Instantiate text with an outline effect. There's probably an easier way.
	outlineWidth := int(u.zoomFactor)
	u.font.SetOutline(outlineWidth)

	outline, _ := u.font.RenderUTF8Solid(s, u.bgColor)
	defer outline.Free()

	u.font.SetOutline(0)
	text, _ := u.font.RenderUTF8Solid(s, u.fgColor)
	defer text.Free()

	// Position vertically. Bottom row is row number 1.
	_, _, _, h, _ := u.texture.Query()
	y := h - int32((u.font.Height())*row) - Margin // TODO: FontSize config

	// Add margin between successive rows.
	if row > 1 {
		y -= Margin * int32(outlineWidth) * 2
	}

	outlineTexture, _ := u.renderer.CreateTextureFromSurface(outline)
	defer outlineTexture.Destroy()
	u.renderer.Copy(outlineTexture,
		nil,
		&sdl.Rect{
			X: Margin,
			Y: y - int32(u.zoomFactor),
			W: outline.W,
			H: outline.H,
		})

	msgTexture, _ := u.renderer.CreateTextureFromSurface(text)
	defer msgTexture.Destroy()
	u.renderer.Copy(msgTexture,
		nil,
		&sdl.Rect{
			X: Margin + int32(u.zoomFactor),
			Y: y,
			W: text.W,
			H: text.H,
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

// Message shows a temporary message that will be cleared after the given
// duration (in seconds). This message stacks with permanent text set with Text().
func (u *UI) Message(text string, seconds time.Duration) {
	// Stop reset timer, a new one will be started.
	// TODO: stack messages (up to, like, 3 or something)
	if u.msgTimer != nil {
		u.msgTimer.Stop()
	}
	u.message = text
	u.msgTimer = time.AfterFunc(time.Second*seconds, u.clearMessage)
}
