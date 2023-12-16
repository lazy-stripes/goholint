package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Widget interface {
	ProcessEvent(Event) bool
	Texture() *sdl.Texture
}

// Base widget type.
type widget struct {
	renderer *sdl.Renderer
	texture  *sdl.Texture

	width, height int32

	// I'm only using a linked list of widgets for now.
	next Widget
}

// new instantiates a widget, stores the renderer and its drawing size, and
// creates the texture to render the widget to.
func new(renderer *sdl.Renderer, size *sdl.Rect) *widget {

	// Take texture size from the clipping rectangle set by the parent widget.
	texture, _ := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		size.W,
		size.H,
	)

	widget := &widget{
		renderer: renderer,
		texture:  texture,
		width:    size.W,
		height:   size.H,
	}

	return widget
}

// ProcessEvent should be overridden in widgets that actually do process events.
func (w *widget) ProcessEvent(Event) bool {
	// Default widget behavior is to not catch events.
	return false
}

func (w *widget) SetNext(next Widget) {
	w.next = next
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
	outlineTexture, _ := w.renderer.CreateTextureFromSurface(outline)
	msgTexture, _ := w.renderer.CreateTextureFromSurface(text)

	labelTexture, _ := w.renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		outline.W,
		outline.H,
	)

	w.renderer.SetRenderTarget(labelTexture)
	w.renderer.Copy(outlineTexture,
		nil,
		&sdl.Rect{
			X: 0,
			Y: 0,
			W: outline.W,
			H: outline.H,
		})
	w.renderer.Copy(msgTexture,
		nil,
		&sdl.Rect{
			// Render text on top of outline, offset by outline width.
			X: int32(properties.Zoom),
			Y: int32(properties.Zoom),
			W: text.W,
			H: text.H,
		})
	w.renderer.SetRenderTarget(nil)

	return labelTexture
}
