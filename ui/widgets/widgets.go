package widgets

import (
	"github.com/lazy-stripes/goholint/assets"
	"github.com/veandco/go-sdl2/img"
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

// Home widget displaying our name and logo. Felt cute. Might delete later.
type Home struct {
	*widget

	// Testing some centered logo.
	// TODO: children widgets, layouts, overengineering...
}

func NewHome(renderer *sdl.Renderer, size *sdl.Rect) *Home {
	return &Home{new(renderer, size)}
}

func (h *Home) ProcessEvent(e *sdl.Event) {
	// TODO
}

func (h *Home) Repaint() {
	icon, err := img.LoadTextureRW(h.renderer, assets.WindowIconRW(), false)
	if err != nil {
		panic(err)
	}
	_, _, iconW, iconH, _ := icon.Query()

	title := h.renderText("Goholint")
	_, _, titleW, titleH, _ := title.Query()
	h.renderer.SetRenderTarget(h.texture)
	h.texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	h.renderer.SetDrawColor(0xcc, 0xcc, 0xcc, 0x90) // TODO: overlay-color config while I'm at it.
	h.renderer.Clear()

	title.SetBlendMode(sdl.BLENDMODE_BLEND)
	h.renderer.Copy(title, nil, &sdl.Rect{
		X: (h.width - titleW - iconW) / 2,
		Y: (h.height - titleH) / 2,
		W: titleW,
		H: titleH,
	})

	icon.SetBlendMode(sdl.BLENDMODE_BLEND)
	h.renderer.Copy(icon, nil, &sdl.Rect{
		X: (h.width - iconW + titleW) / 2,
		Y: (h.height - iconH) / 2,
		W: iconW,
		H: iconH,
	})

	h.renderer.SetRenderTarget(nil)
}

func (h *Home) Texture() *sdl.Texture {
	// TODO: repaint if needed.
	return h.texture
}

// TODO: widgets.Label. For now just render text to a texture.
// TODO: find out where to store font size and outline width
func (h *Home) renderText(text string) *sdl.Texture {
	// Instantiate text with an outline effect. There's probably an easier way.
	// TODO: shouldn't we freeing most of this?
	properties.TitleFont.SetOutline(properties.Zoom)
	outline, _ := properties.TitleFont.RenderUTF8Solid(text, properties.BgColor)
	properties.TitleFont.SetOutline(0)
	msg, _ := properties.TitleFont.RenderUTF8Solid(text, properties.FgColor)

	// I can't draw the text directly on the outline as CreateTextureFromSurface
	// creates static textures. Bummer.
	outlineTexture, _ := h.renderer.CreateTextureFromSurface(outline)
	msgTexture, _ := h.renderer.CreateTextureFromSurface(msg)

	labelTexture, _ := h.renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		outline.W,
		outline.H,
	)

	h.renderer.SetRenderTarget(labelTexture)
	h.renderer.Copy(outlineTexture,
		nil,
		&sdl.Rect{
			X: 0,
			Y: 0,
			W: outline.W,
			H: outline.H,
		})
	h.renderer.Copy(msgTexture,
		nil,
		&sdl.Rect{
			// Render text on top of outline, offset by outline width.
			X: int32(properties.Zoom),
			Y: int32(properties.Zoom),
			W: msg.W,
			H: msg.H,
		})
	h.renderer.SetRenderTarget(nil)

	return labelTexture
}
