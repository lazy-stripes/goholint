package widgets

import (
	"github.com/lazy-stripes/goholint/assets"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Home widget displaying our name and logo. Felt cute. Might delete later.
type Home struct {
	*VerticalLayout

	// Testing some centered logo.
	viewport *sdl.Rect

	// TODO: children widgets, layouts, overengineering...
}

func NewHome(size *sdl.Rect) *Home {
	h := &Home{
		VerticalLayout: NewVerticalLayout(size, nil),
	}

	// Compute viewport size. It'd be easier to render once and get it from
	// title texture size.
	headerHeight := h.drawHeader()
	h.viewport = &sdl.Rect{
		W: h.width,
		H: h.height - headerHeight,
		X: 0,
		Y: headerHeight,
	}

	// TODO: set next from UI.
	choices := []MenuChoice{
		{"Resume", nil},
		{"Quit", nil},
	}
	h.Add(NewMenu(h.viewport, choices))

	h.repaint()

	return h
}

func (h *Home) repaint() {
	// TODO: widget.renderNext(&viewport) should be feasible.
	if len(h.children) > 0 {
		t := h.children[0].Texture()
		t.SetBlendMode(sdl.BLENDMODE_BLEND)
		renderer.SetRenderTarget(h.texture)
		renderer.Copy(t, nil, h.viewport)

		//renderer.SetDrawColor(0xff, 0, 0x80, 128)
		//renderer.DrawRect(h.viewport)

		renderer.SetRenderTarget(nil)
	}
}

func (h *Home) drawHeader() (height int32) {
	icon, err := img.LoadTextureRW(renderer, assets.WindowIconRW(), false)
	if err != nil {
		panic(err)
	}
	defer icon.Destroy()
	_, _, iconW, iconH, _ := icon.Query()

	title := h.renderText("Goholint")
	defer title.Destroy()

	_, _, titleW, titleH, _ := title.Query()
	renderer.SetRenderTarget(h.texture)
	//renderer.SetRenderTarget(nil)
	h.texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	renderer.SetDrawColor(0xcc, 0xcc, 0xcc, 0x90) // TODO: overlay-color config while I'm at it.
	renderer.Clear()

	// Show name and logo as a header.
	margin := 8 * int32(DefaultProperties.Zoom)
	title.SetBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Copy(title, nil, &sdl.Rect{
		X: (h.width - titleW - iconW) / 2,
		Y: margin,
		W: titleW,
		H: titleH,
	})

	icon.SetBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Copy(icon, nil, &sdl.Rect{
		X: (h.width - iconW + titleW) / 2,
		Y: (titleH-iconH)/2 + margin, // Aligned with title center
		W: iconW,
		H: iconH,
	})

	renderer.SetRenderTarget(nil)

	// Return height to compute viewport size. Height is max(titleH, iconH).
	if titleH > iconH {
		return titleH + margin
	}
	return iconH + margin
}

func (h *Home) Texture() *sdl.Texture {
	// No need to repaint, this is a static widget.
	//h.repaint()
	return h.texture
}
