package screen

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// UI structure to manage user commands and overlay.
type UI struct {
	Enabled bool

	texture  *sdl.Texture
	renderer *sdl.Renderer

	font *ttf.Font

	fg color.RGBA // TODO: make it configurable
	bg color.RGBA // TODO: make it configurable
}

// Return a UI instance given a texture on which it will write.
func NewUI(renderer *sdl.Renderer, zoom uint) *UI {
	font, err := ttf.OpenFont("assets/ui.ttf", 8) // FIXME: based on screen size (or configurable)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		return nil // TODO: result, err
	}

	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		ScreenWidth*int32(zoom),
		ScreenHeight*int32(zoom))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create UI texture: %s\n", err)
		return nil // TODO: result, err
	}

	// Background transparency.
	texture.SetBlendMode(sdl.BLENDMODE_BLEND)

	ui := UI{
		Enabled:  true,
		texture:  texture,
		renderer: renderer,
		font:     font,
		fg:       ColorBlack,
		bg:       ColorWhite,
	}
	return &ui
}

// Message creates a new UI texture with the given message, enables UI and
// starts a timer that will hide the UI when it's done.
func (u *UI) Message(text string, duration uint) {
	// Reset texture. FIXME: can we do without the background texture altogether?
	u.renderer.SetRenderTarget(u.texture)
	u.renderer.SetDrawColor(0, 0, 0, 0)
	u.renderer.Clear()

	// Instantiate text with an outline effect. There's probably an easier way.
	u.font.SetOutline(1)
	outline, _ := u.font.RenderUTF8Solid(text, sdl.Color(u.bg))
	u.font.SetOutline(0)
	msg, _ := u.font.RenderUTF8Solid(text, sdl.Color(u.fg))

	_, _, _, h, _ := u.texture.Query()
	y := h - 8 - 5 // TODO: FontSize config

	outlineTexture, _ := u.renderer.CreateTextureFromSurface(outline)
	u.renderer.Copy(outlineTexture, nil, &sdl.Rect{X: 4, Y: y - 1, W: outline.W, H: outline.H})

	msgTexture, _ := u.renderer.CreateTextureFromSurface(msg)
	u.renderer.Copy(msgTexture, nil, &sdl.Rect{X: 5, Y: y, W: msg.W, H: msg.H})

	u.renderer.SetRenderTarget(nil)

	// TODO: detect duplicate message, add (×<num>) suffix
	u.Enabled = true
	time.AfterFunc(time.Second*time.Duration(duration), func() { u.Enabled = false })
}
