package widgets

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

type Label struct {
	*widget

	text string
}

// TODO: why couldn't these be methods of a global UI object abstracting the renderer?
//       I'd just need to move all widgets back up to the ui package.
// TODO: actual font size
func NewLabel(sizeHint *sdl.Rect, text string, p ...Properties) *Label {
	props := DefaultProperties
	if len(p) > 0 {
		props = p[0]
	}

	if sizeHint == noSizeHint {
		// Query font size to create our own size hint for the internal texture.
		sizeHint = &sdl.Rect{}

		// Allow Labels to have several lines. We'll have to compute the max
		// width and total added height for the whole thing.
		lines := strings.Split(text, "\n")
		props.Font.SetOutline(props.Zoom)
		for _, line := range lines {
			w, h, _ := props.Font.SizeUTF8(line)
			if int32(w) > sizeHint.W {
				sizeHint.W = int32(w)
			}
			sizeHint.H += int32(h)
		}
		props.Font.SetOutline(0)

	}
	l := &Label{
		widget: new(sizeHint, props),
		text:   text,
	}
	l.repaint()
	return l
}

func (l *Label) repaint() {
	l.clear()

	// Instantiate text with an outline effect. There's probably an easier way.
	l.Font.SetOutline(DefaultProperties.Zoom)
	// Render*Wrapped nicely handles newlines and will only wrap on them if
	// called with wrap length 0.
	outline, _ := l.Font.RenderUTF8BlendedWrapped(l.text, l.BgColor, 0)
	defer outline.Free()

	l.Font.SetOutline(0)
	text, _ := l.Font.RenderUTF8BlendedWrapped(l.text, l.FgColor, 0)
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
