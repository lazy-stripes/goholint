package widgets

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

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
	ProcessEvent(Event) bool
	Texture() *sdl.Texture
}

// Base widget type.
type widget struct {
	texture *sdl.Texture

	width, height int32 // XXX could this be derived from texture?

	children []Widget // List of sub-widgets

	background sdl.Color // Background color. Default value is transparent.

	// TODO: margins (or better: properties)
}

// new instantiates a widget, stores the renderer and its drawing size, and
// creates the texture to render the widget to.
func new(size *sdl.Rect) *widget {
	widget := &widget{
		texture: texture(size),
		width:   size.W,
		height:  size.H,
	}
	return widget
}

// ProcessEvent should be overridden in widgets that actually do process events.
func (w *widget) ProcessEvent(e Event) bool {
	// Propagate event processing until a sub-widget catches it.
	for _, c := range w.children {
		if c.ProcessEvent(e) {
			return true
		}
	}
	return false
}

// Add appends a sub-widget to the internal list of children.
func (w *widget) Add(child Widget) {
	w.children = append(w.children, child)
}

// renderText lets the widget render outlined text to a new texture using its
// internal renderer.
func (w *widget) renderText(s string) *sdl.Texture {
	// Instantiate text with an outline effect. There's probably an easier way.
	properties.TitleFont.SetOutline(properties.Zoom)
	outline, _ := properties.TitleFont.RenderUTF8Solid(s, properties.BgColor)
	defer outline.Free()

	properties.TitleFont.SetOutline(0)
	text, _ := properties.TitleFont.RenderUTF8Solid(s, properties.FgColor)
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
			X: int32(properties.Zoom),
			Y: int32(properties.Zoom),
			W: text.W,
			H: text.H,
		})
	renderer.SetRenderTarget(nil)

	return labelTexture
}
