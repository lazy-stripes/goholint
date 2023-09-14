package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Widget interface {
	ProcessEvent(*sdl.Event)
	Repaint() // This is probably not needed and should be internal.
	Texture() *sdl.Texture
}

// Base widget type.
type widget struct {
	renderer *sdl.Renderer
	texture  *sdl.Texture

	width, height int32
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
