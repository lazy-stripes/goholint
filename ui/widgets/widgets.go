package widgets

import (
	"fmt"

	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/veandco/go-sdl2/sdl"
)

// Package-wide logger.
var log = logger.New("widgets", "widget-level debug")

// noSizeHint is a safe non-nil zero-size rect to use when creating widgets that
// are expected to be able to handle their own size (i.e. labels).
var noSizeHint = &sdl.Rect{}

// This feels dirty but I'm going for convenient right now.

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
		// | -------- |-------------- | ---------- |
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
	renderer = r
}

type Widget interface {
	// Hide sets the widget's internal visiblity flag. This can be used to
	// temporarily hide a widget within a group or layout without removing the
	// widget from it.
	Hide(bool)

	// IsVisible just returns the current value of the internal visibility flag.
	IsVisible() bool

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
	hidden        bool  // If true, widget may not show up in groups and layouts
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
		offset = 0
	case align.Center:
		offset = (w.width - width) / 2
	case align.Right:
		offset = w.width - width
	}
	return
}

// alignX returns the vertical offset relative to the widget's top for the
// given height.
func (w *widget) alignY(height int32) (offset int32) {
	switch w.VerticalAlign {
	case align.Top:
		offset = 0
	case align.Middle:
		offset = (w.height - height) / 2
	case align.Bottom:
		offset = w.height - height
	}
	return
}

// Hide takes a boolean that will define whether the widget should be hidden or
// visible. A widget is visible by default at creation time.
func (w *widget) Hide(hidden bool) {
	w.hidden = hidden
}

// IsVIsible makes the visibility flag accessible to the interface.
func (w *widget) IsVisible() bool {
	return !w.hidden
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
	// TODO: clear() and return transparent texture on hidden?
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
	renderer.SetRenderTarget(nil)

	return w.texture
}

// Destroy frees the widget's internal texture.
func (w *widget) Destroy() {
	if w.texture != nil {
		if err := w.texture.Destroy(); err != nil {
			fmt.Printf("error while destroying texture: %v", err)
		}
	}
}
