package widgets

import "github.com/veandco/go-sdl2/sdl"

// Let's do this.

// VerticalLayout renders its children widgets from top to bottom, vertically
// centered.
type VerticalLayout struct {
	*widget

	viewport *sdl.Rect
	children []Widget // TODO: make that part of widget struct?
}

func NewVerticalLayout(size *sdl.Rect, children []Widget) *VerticalLayout {
	l := &VerticalLayout{
		widget:   new(size),
		viewport: size,
		children: children,
	}

	return l
}

func (l *VerticalLayout) Add(child Widget) {
	l.children = append(l.children, child)
}

func (l *VerticalLayout) Texture() *sdl.Texture {
	l.repaint()
	return l.texture
}

func (l *VerticalLayout) ProcessEvent(Event) bool {
	return false
}

// repaint renders children top-down and spaces them vertically.
func (l *VerticalLayout) repaint() {
	var textures []*sdl.Texture
	totalHeight := int32(0)

	for _, c := range l.children {
		// Render child to texture. Keep track of height.
		// TODO: obviously, at some point, make a HorizontalLayout using common code.
		t := c.Texture()
		_, _, _, h, _ := t.Query()
		totalHeight += h
		textures = append(textures, t)
	}

	// Compute inter-widget space if any.
	margin := (l.height - int32(totalHeight)) / int32(len(l.children)+1)
	if margin < 0 {
		margin = 0
	}

	// Render to our texture, horizontally center each child.
	// TODO: making that horizontal aligment configurable would be neat.
	renderer.SetRenderTarget(l.texture)
	y := int32(0) // Start at the top of the texture
	for _, t := range textures {
		y += margin

		_, _, w, h, _ := t.Query()
		renderer.Copy(t, nil, &sdl.Rect{
			X: 0, // FIXME: horizontal align
			Y: y,
			W: w,
			H: h,
		})

		y += h
	}
}
