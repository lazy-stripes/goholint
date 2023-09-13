package ui

import (
	"github.com/lazy-stripes/goholint/assets"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Widget interface {
	ProcessEvent(*sdl.Event)
	Repaint()
	Texture() *sdl.Texture
}

type RootWidget struct {
	renderer *sdl.Renderer
	texture  *sdl.Texture

	font *ttf.Font // Title font
	opts *UIOptions

	width, height int32
	// TODO: children widgets, layouts, overengineering...
}

// Convenience type for stuff we want to pass to widgets. All of these should
// be optional depending on widget type.
type UIOptions struct {
	font   *ttf.Font
	zoom   int
	bg, fg sdl.Color
}

// TODO: UIOptions struct for all parameters.
func NewRootWidget(renderer *sdl.Renderer, opts *UIOptions) *RootWidget {
	widget := &RootWidget{
		renderer: renderer,
		opts:     opts,
	}

	widget.width, widget.height, _ = renderer.GetOutputSize()
	widget.texture, _ = renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		widget.width,
		widget.height,
	)

	widget.font, _ = ttf.OpenFontRW(assets.UIFont, 1, int(12*opts.zoom))

	// Testing some centered logo.

	return widget
}

func (w *RootWidget) ProcessEvent(e *sdl.Event) {
	// TODO
}

func (w *RootWidget) Repaint() {
	// Seek back to beginning of icon, since it was read when creating window.
	assets.WindowIcon.Seek(0, sdl.RW_SEEK_SET)
	icon, err := img.LoadTextureRW(w.renderer, assets.WindowIcon, false)
	if err != nil {
		panic(err)
	}
	_, _, iconW, iconH, _ := icon.Query()

	title := w.renderText("Goholint")
	_, _, titleW, titleH, _ := title.Query()
	w.renderer.SetRenderTarget(w.texture)
	w.texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	w.renderer.SetDrawColor(0xcc, 0xcc, 0xcc, 0x90)
	w.renderer.Clear()

	title.SetBlendMode(sdl.BLENDMODE_BLEND)
	w.renderer.Copy(title, nil, &sdl.Rect{
		X: (w.width - titleW) / 2,
		Y: (w.height - titleH - iconH) / 2,
		W: titleW,
		H: titleH,
	})

	icon.SetBlendMode(sdl.BLENDMODE_BLEND)
	w.renderer.Copy(icon, nil, &sdl.Rect{
		X: (w.width - iconW) / 2,
		Y: ((w.height - iconH + titleH) / 2),
		W: iconW,
		H: iconH,
	})

	w.renderer.SetRenderTarget(nil)
}

func (w *RootWidget) Texture() *sdl.Texture {
	// TODO: repaint if needed.
	return w.texture
}

// TODO: widgets.Label. For now just render text to a texture.
// TODO: find out where to store font size and outline width
func (w *RootWidget) renderText(text string) *sdl.Texture {
	// Instantiate text with an outline effect. There's probably an easier way.
	// TODO: shouldn't we freeing most of this?
	w.font.SetOutline(w.opts.zoom)
	outline, _ := w.font.RenderUTF8Solid(text, w.opts.bg)
	w.font.SetOutline(0)
	msg, _ := w.font.RenderUTF8Solid(text, w.opts.fg)

	// I can't draw the text directly on the outline as CreateTextureFromSurface
	// creates static textures. Bummer.
	outlineTexture, _ := w.renderer.CreateTextureFromSurface(outline)
	msgTexture, _ := w.renderer.CreateTextureFromSurface(msg)

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
			X: int32(w.opts.zoom),
			Y: int32(w.opts.zoom),
			W: msg.W,
			H: msg.H,
		})
	w.renderer.SetRenderTarget(nil)

	return labelTexture
}

//utlineTexture, _ := u.renderer.CreateTextureFromSurface(outline)
// .renderer.Copy(outlineTexture,
//	nil,
//	&sdl.Rect{
//		X: Margin,
//		Y: y - int32(u.zoomFactor),
//		W: outline.W,
//		H: outline.H,
//	})
//
//sgTexture, _ := u.renderer.CreateTextureFromSurface(msg)
//.renderer.Copy(msgTexture,
//	nil,
//	&sdl.Rect{
//		X: Margin + int32(u.zoomFactor),
//		Y: y,
//		W: msg.W,
//		H: msg.H,
//	})
