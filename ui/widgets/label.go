package widgets

import "github.com/veandco/go-sdl2/sdl"

type Label struct {
	*widget

	Text string
}

// TODO: why couldn't these be methods of a global UI object abstracting the renderer?
//       I'd just need to move all widgets back up to the ui package.
func NewLabel(r *sdl.Renderer, text string) *Label {
	texture := renderText(r, text)
	_, _, w, h, _ := texture.Query()
	l := &Label{
		widget: &widget{
			renderer: r,
			width: w,
			height: h,
			texture: texture,
		},
		Text: text,
	}

	return l
}

func (l *Label) Texture() *sdl.Texture {
	return l.texture
}

func (l *Label) ProcessEvent(Event) bool {
	return false
}

// renderText renders the given string on the given renderer. A a new texture
// will be created.
func renderText(r *sdl.Renderer, s string) *sdl.Texture {
	// Instantiate text with an outline effect. There's probably an easier way.
	properties.TitleFont.SetOutline(properties.Zoom)
	outline, _ := properties.TitleFont.RenderUTF8Solid(s, properties.BgColor)
	defer outline.Free()

	properties.TitleFont.SetOutline(0)
	text, _ := properties.TitleFont.RenderUTF8Solid(s, properties.FgColor)
	defer text.Free()

	// I can't draw the text directly on the outline as CreateTextureFromSurface
	// creates static textures. Bummer.
	outlineTexture, _ := r.CreateTextureFromSurface(outline)
	msgTexture, _ := r.CreateTextureFromSurface(text)

	labelTexture, _ := r.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		outline.W,
		outline.H,
	)

	r.SetRenderTarget(labelTexture)
	r.Copy(outlineTexture,
		nil,
		&sdl.Rect{
			X: 0,
			Y: 0,
			W: outline.W,
			H: outline.H,
		})
	r.Copy(msgTexture,
		nil,
		&sdl.Rect{
			// Render text on top of outline, offset by outline width.
			X: int32(properties.Zoom),
			Y: int32(properties.Zoom),
			W: text.W,
			H: text.H,
		})
	r.SetRenderTarget(nil)

	return labelTexture
}