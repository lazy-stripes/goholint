package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Label struct {
	*widget

	font *ttf.Font
	text string
}

// TODO: why couldn't these be methods of a global UI object abstracting the renderer?
//       I'd just need to move all widgets back up to the ui package.
// TODO: alignment
// TODO: actual font size
func NewLabel(sizeHint *sdl.Rect, text string) *Label {
	return newLabel(sizeHint, text, DefaultProperties.TitleFont)
}

func newLabel(sizeHint *sdl.Rect, text string, font *ttf.Font) *Label {
	// Query font size to create internal texture.
	if sizeHint == noSizeHint {
		w, h, _ := font.SizeUTF8(text)
		sizeHint = &sdl.Rect{
			W: int32(w),
			H: int32(h),
		}
	}
	l := &Label{
		widget: new(sizeHint),
		font:   font,
		text:   text,
	}
	l.repaint()
	return l
}

func (l *Label) repaint() {
	l.clear()

	// Instantiate text with an outline effect. There's probably an easier way.
	l.font.SetOutline(DefaultProperties.Zoom)
	outline, _ := l.font.RenderUTF8Blended(l.text, l.BgColor)
	defer outline.Free()

	l.font.SetOutline(0)
	text, _ := l.font.RenderUTF8Blended(l.text, l.FgColor)
	defer text.Free()

	// Render text on top of outline, offset by outline width.
	text.Blit(nil,
		outline,
		&sdl.Rect{
			X: int32(DefaultProperties.Zoom),
			Y: int32(DefaultProperties.Zoom),
			W: text.W,
			H: text.H,
		})

	outlineTexture, _ := renderer.CreateTextureFromSurface(outline)
	renderer.SetRenderTarget(l.texture)
	renderer.Copy(outlineTexture,
		nil,
		// TODO: align, somehow
		&sdl.Rect{
			X: 0,
			Y: 0,
			W: outline.W,
			H: outline.H,
		})
	renderer.SetRenderTarget(nil)

}

// renderText renders the given string on the given renderer. A a new texture
// will be created.
func renderText(s string) *sdl.Texture {
	// Instantiate text with an outline effect. There's probably an easier way.
	DefaultProperties.TitleFont.SetOutline(DefaultProperties.Zoom)
	outline, _ := DefaultProperties.TitleFont.RenderUTF8Solid(s, DefaultProperties.BgColor)
	defer outline.Free()

	DefaultProperties.TitleFont.SetOutline(0)
	text, _ := DefaultProperties.TitleFont.RenderUTF8Solid(s, DefaultProperties.FgColor)
	defer text.Free()

	text.Blit(nil,
		outline,
		&sdl.Rect{
			// Render text on top of outline, offset by outline width.
			X: int32(DefaultProperties.Zoom),
			Y: int32(DefaultProperties.Zoom),
			W: text.W,
			H: text.H,
		})

	// I can't draw the text directly on the outline as CreateTextureFromSurface
	// creates static textures. Bummer.
	outlineTexture, _ := renderer.CreateTextureFromSurface(outline)

	//msgTexture, _ := renderer.CreateTextureFromSurface(text)

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
	//renderer.Copy(msgTexture,
	//	nil,
	//	&sdl.Rect{
	//		// Render text on top of outline, offset by outline width.
	//		X: int32(DefaultProperties.Zoom),
	//		Y: int32(DefaultProperties.Zoom),
	//		W: text.W,
	//		H: text.H,
	//	})
	renderer.SetRenderTarget(nil)

	return labelTexture
}
