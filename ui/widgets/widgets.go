package widgets

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// sizeHint is a safe non-nil zero-size rect to use when creating widgets that
// are expected to be able to handle their own size (i.e. labels).
var noSizeHint = sdl.Rect{}

// This feels dirty but I'm going for convenient right now.

var renderer *sdl.Renderer

func texture(size *sdl.Rect) *sdl.Texture {
	texture, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		size.W,
		size.H,
	)

	if err != nil {
		// I'm already not checking for error anywhere else, but this should at
		// least provide a log before the caller panics on a nil texture.
		log.Printf("failed to create texture: %v", err)
	}

	return texture
}

func Init(r *sdl.Renderer) {
	renderer = r
}

type Widget interface {
	// ProcessEvent returns true if the widget caught and handled the event,
	// false if it did not.
	ProcessEvent(Event) bool

	// Texture return the widget's internal texture in its current state. This
	// call might modify the renderer's state if a widget redraws its texture
	// just-in-time.
	Texture() *sdl.Texture
}

// Base widget type.
type widget struct {
	Properties

	texture *sdl.Texture

	width, height int32 // XXX could this be derived from texture?
}

// new instantiates a widget, stores the renderer and its drawing size, and
// creates the texture to render the widget to.
func new(size *sdl.Rect) *widget {
	widget := &widget{
		Properties: DefaultProperties,
		texture:    texture(size),
		width:      size.W,
		height:     size.H,
	}
	widget.clear()
	return widget
}

// clear repaints the widget's internal texture with the current background
// color. Automatically called at creation time. Should be called before a
// repaint.
func (w *widget) clear() {
	renderer.SetDrawColor(
		w.Background.R,
		w.Background.G,
		w.Background.B,
		w.Background.A)
	renderer.SetRenderTarget(w.texture)
	renderer.Clear()
	renderer.SetRenderTarget(nil)
}

// ProcessEvent should be overridden in widgets that actually do process events.
// The default implementation always returns false to indicate no event is
// handled.
func (w *widget) ProcessEvent(e Event) bool {
	return false
}

// Texture should be called by subclasses to apply unused properties like border
// to the widget's internal texture.
func (w *widget) Texture() *sdl.Texture {
	// TODO: call w.repaint() and remove .Texture() from all subclasses that don't need to override it?
	// Draw border on top of internal texture.
	_, _, width, height, _ := w.texture.Query()
	renderer.SetRenderTarget(w.texture)
	renderer.SetDrawColor(w.FgColor.R, w.FgColor.G, w.FgColor.B, w.FgColor.A)
	rect := sdl.Rect{}
	for i := int32(0); i < w.Border; i++ {
		rect.X = i
		rect.Y = i
		rect.W = width - i
		rect.H = height - i
		renderer.DrawRect(&rect)
	}
	renderer.SetRenderTarget(nil)

	return w.texture
}

// renderText lets the widget render outlined text to a new texture using its
// internal renderer.
func (w *widget) renderText(s string) *sdl.Texture {
	// Instantiate text with an outline effect. There's probably an easier way.
	DefaultProperties.TitleFont.SetOutline(DefaultProperties.Zoom)
	outline, _ := DefaultProperties.TitleFont.RenderUTF8Solid(s, DefaultProperties.BgColor)
	defer outline.Free()

	DefaultProperties.TitleFont.SetOutline(0)
	text, _ := DefaultProperties.TitleFont.RenderUTF8Solid(s, DefaultProperties.FgColor)
	defer text.Free()

	// I can't draw the text directly on the outline as CreateTextureFromSurface
	// creates static textures. Bummer.
	outlineTexture, _ := renderer.CreateTextureFromSurface(outline)
	msgTexture, _ := renderer.CreateTextureFromSurface(text)

	labelTexture, _ := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		outline.W,
		outline.H,
	)

	renderer.SetRenderTarget(labelTexture)
	renderer.Copy(outlineTexture,
		nil,
		&sdl.Rect{
			X: 0,
			Y: 0,
			W: outline.W,
			H: outline.H,
		})
	renderer.Copy(msgTexture,
		nil,
		&sdl.Rect{
			// Render text on top of outline, offset by outline width.
			X: int32(DefaultProperties.Zoom),
			Y: int32(DefaultProperties.Zoom),
			W: text.W,
			H: text.H,
		})
	renderer.SetRenderTarget(nil)

	return labelTexture
}
