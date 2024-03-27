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
// TODO: actual font size
func NewLabel(sizeHint *sdl.Rect, text string, props ...Properties) *Label {
	return newLabel(sizeHint, text, DefaultProperties.TitleFont, props...)
}

func newLabel(sizeHint *sdl.Rect, text string, font *ttf.Font, props ...Properties) *Label {
	if sizeHint == noSizeHint {
		// Query font size to create internal texture.
		w, h, _ := font.SizeUTF8(text)
		sizeHint = &sdl.Rect{
			W: int32(w),
			H: int32(h),
		}
	}
	l := &Label{
		widget: new(sizeHint, props...),
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
		&sdl.Rect{
			X: l.alignX(outline.W),
			Y: l.alignY(outline.H),
			W: outline.W,
			H: outline.H,
		})
	renderer.SetRenderTarget(nil)
}
