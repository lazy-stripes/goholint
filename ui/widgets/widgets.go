package widgets

import (
	"fmt"

	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/veandco/go-sdl2/sdl"
)

// Package-wide logger for widgets.
var log = logger.New("widgets", "widget-level debug")

// noSizeHint is a safe non-nil zero-size rect to use when creating widgets that
// are expected to be able to handle their own size (i.e. labels).
var noSizeHint = &sdl.Rect{}

// Globally available renderer instance. This feels dirty but I'm going for
// convenient right now.
var renderer *sdl.Renderer

func texture(size *sdl.Rect) *sdl.Texture {
	texture, err := renderer.CreateTexture(
		// My understanding of RGBA8888 was too naive, but fortunately I finally
		// found the explanation on the SDL_Color wiki page:
		//
		// "The bits of this structure can be directly reinterpreted as an
		// integer-packed color which uses the SDL_PIXELFORMAT_RGBA32 format
		// (SDL_PIXELFORMAT_ABGR8888 on little-endian systems and
		// SDL_PIXELFORMAT_RGBA8888 on big-endian systems)."
		//
		// So I just needed PIXELFORMAT_RGBA32 all along. For more context, the
		// Wikipedia page about RGBA has a neat little table describing it:
		//
		// |          | Little-endian | Big-endian |
		// |----------|---------------|------------|
		// | RGBA8888 | ABGR32        | RGBA32     |
		// | ARGB32   | BGRA8888      | ARGB8888   |
		// | RGBA32   | ABGR8888      | RGBA8888   |
		//
		// TL;DR: just use RGBA32, you big stripy dumbass.
		uint32(sdl.PIXELFORMAT_RGBA32),
		sdl.TEXTUREACCESS_TARGET,
		size.W,
		size.H,
	)

	if err != nil {
		// I'm already not checking for error anywhere else, but this should at
		// least provide a log before the caller panics on a nil texture.
		log.Warningf("failed to create texture: %v", err)
	}

	texture.SetBlendMode(sdl.BLENDMODE_BLEND)

	return texture
}

func Init(r *sdl.Renderer) {
	// For debugging purposes. Someday it might even be configurable.
	if log.Enabled() && logger.Level >= logger.Debug {
		DefaultProperties.BorderColor = sdl.Color{0xff, 0x00, 0x00, 0xff}
	}

	renderer = r
}

type Widget interface {
	// SetVisible sets the widget's internal visiblity flag. This can be used to
	// temporarily hide a widget within a group or layout without removing the
	// widget from it.
	SetVisible(bool)

	// Visible returns the current value of the internal visibility flag.
	Visible() bool

	// ProcessEvent returns true if the widget caught and handled the event,
	// false if it did not.
	ProcessEvent(Event) bool

	// Texture return the widget's internal texture in its current state. This
	// call might modify the renderer's state if a widget redraws its texture
	// just-in-time.
	Texture() *sdl.Texture

	// Destroy releases all resources dynamically allocated by the widget, like
	// its internal texture.
	Destroy()
}

// Base widget type.
type widget struct {
	Properties

	texture *sdl.Texture

	width, height int32 // Widget's actual size (may not be the same as texture)
	visible       bool  // If false, widget may not show up in groups/layouts
}

// new instantiates a widget, stores the renderer and its drawing size, and
// creates the texture to render the widget to. Optional properties can also
// be provided.
func new(sizeHint *sdl.Rect, props ...Properties) *widget {
	p := DefaultProperties
	if len(props) > 0 {
		p = props[0]
	}

	// Apply margin before creating texture.
	size := *sizeHint
	size.W += p.Margin * 2 // Apply margin to left + right
	size.H += p.Margin * 2 // Apply margin to top + bottom

	widget := &widget{
		Properties: p,
		texture:    texture(&size),
		width:      size.W,
		height:     size.H,
		visible:    true,
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
	renderer.SetRenderTarget(nil) // TODO: remove? We usually write to the texture after a clear anyway.
}

// alignX returns the horizontal offset relative to the widget's left for the
// given width.
func (w *widget) alignX(width int32) (offset int32) {
	switch w.HorizontalAlign {
	case align.Left:
		offset = w.Padding
	case align.Center:
		offset = (w.width - width) / 2
	case align.Right:
		offset = w.width - width - w.Padding
	}
	return
}

// alignY returns the vertical offset relative to the widget's top for the
// given height.
func (w *widget) alignY(height int32) (offset int32) {
	switch w.VerticalAlign {
	case align.Top:
		offset = w.Padding
	case align.Middle:
		offset = (w.height - height) / 2
	case align.Bottom:
		offset = w.height - height - w.Padding
	}
	return
}

// SetVisible takes a boolean that will define whether the widget should be hidden or
// visible. A widget is visible by default at creation time.
func (w *widget) SetVisible(visible bool) {
	w.visible = visible
}

// Visible makes the visibility flag state accessible to the Widget interface.
func (w *widget) Visible() bool {
	return w.visible
}

// ProcessEvent should be overridden in widgets that actually do process events.
// The default implementation always returns false to indicate no event is
// handled.
func (w *widget) ProcessEvent(e Event) bool {
	return false
}

// Texture returns the widget's internal texture after applying properties like
// border color to it.
func (w *widget) Texture() *sdl.Texture {
	// TODO: clear() and return transparent texture if not visible?
	// Draw border on top of internal texture.
	_, _, width, height, _ := w.texture.Query()
	renderer.SetRenderTarget(w.texture)
	renderer.SetDrawColor(
		w.BorderColor.R,
		w.BorderColor.G,
		w.BorderColor.B,
		w.BorderColor.A,
	)
	rect := sdl.Rect{}
	for i := int32(0); i < w.Border; i++ {
		rect.X = i
		rect.Y = i
		rect.W = width - 2*i
		rect.H = height - 2*i
		renderer.DrawRect(&rect)
	}

	// Render padding in a semi-transparent color based on border.
	renderer.SetDrawColor(
		w.BorderColor.R,
		w.BorderColor.G,
		w.BorderColor.B,
		w.BorderColor.A/2,
	)
	for i := w.Border; i < w.Padding; i++ {
		rect.X = i
		rect.Y = i
		rect.W = width - 2*i
		rect.H = height - 2*i
		renderer.DrawRect(&rect)
	}
	renderer.SetRenderTarget(nil)

	return w.texture
}

// Destroy frees the widget's internal texture.
func (w *widget) Destroy() {
	if w.texture != nil {
		if err := w.texture.Destroy(); err != nil {
			fmt.Printf("error while destroying texture: %v\n", err)
		}
	}
}
