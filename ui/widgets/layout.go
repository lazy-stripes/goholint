package widgets

import (
	"github.com/lazy-stripes/goholint/ui/widgets/align"

	"github.com/veandco/go-sdl2/sdl"
)

// VerticalLayout renders its children widgets from top to bottom, vertically
// centered.
type VerticalLayout struct {
	*Group
}

func NewVerticalLayout(size *sdl.Rect, children []Widget, props ...Properties) *VerticalLayout {
	l := &VerticalLayout{
		Group: NewGroup(size, children, props...),
	}
	l.repaint()
	return l
}

//func (l *VerticalLayout) Add(child Widget) {
//	l.Group.Add(child)
//	l.repaint()
//}

// Texture bypasses the base Group method to just render aligned children as-is.
func (l *VerticalLayout) Texture() *sdl.Texture {
	l.repaint()
	return l.widget.Texture()
}

// repaint renders children top-down and spaces them vertically.
func (l *VerticalLayout) repaint() {
	l.clear()

	var textures []*sdl.Texture
	totalHeight := int32(0)

	for _, c := range l.children {
		// Render child to texture. Keep track of height.
		// TODO: common layout code.
		t := c.Texture()
		_, _, _, h, _ := t.Query()
		totalHeight += h
		textures = append(textures, t)
	}

	// Compute inter-widget space if any.
	spacing := int32(0)
	if l.VerticalAlign == align.Justified && l.height > totalHeight {
		spacing = (l.height - int32(totalHeight)) / int32(len(l.children)+1)
	}

	// Render to our texture, horizontally align each child.
	renderer.SetRenderTarget(l.texture)

	// Offset contents Y value depending on alignment.
	y := l.alignY(totalHeight)

	for _, t := range textures {
		y += spacing

		_, _, w, h, _ := t.Query()
		renderer.Copy(t, nil, &sdl.Rect{
			X: l.alignX(w),
			Y: y,
			W: w,
			H: h,
		})

		y += h
	}
	renderer.SetRenderTarget(nil)
}

// HorizontalLayout renders its children widgets from left to right, horizontally
// centered.
type HorizontalLayout struct {
	*Group
}

func NewHorizontalLayout(size *sdl.Rect, children []Widget, props ...Properties) *HorizontalLayout {
	l := &HorizontalLayout{
		Group: NewGroup(size, children, props...),
	}
	l.repaint()
	return l
}

//func (l *HorizontalLayout) Add(child Widget) {
//	l.Group.Add(child)
//	l.repaint()
//}

// Texture overrides the base Group method to render aligned children as-is.
func (l *HorizontalLayout) Texture() *sdl.Texture {
	l.repaint()
	return l.widget.Texture()
}

//func (l *HorizontalLayout) ProcessEvent(e Event) bool {
//	// Repaint if one of our children handled the event.
//	caught := l.Group.ProcessEvent(e)
//	if caught {
//		l.repaint()
//	}
//	return caught
//}

// repaint renders children left-right and spaces them horizontally.
func (l *HorizontalLayout) repaint() {
	l.clear()

	var textures []*sdl.Texture
	totalWidth := int32(0)

	for _, c := range l.children {
		// Render child to texture. Keep track of width.
		// TODO: common layout code.
		t := c.Texture()
		_, _, w, _, _ := t.Query()
		totalWidth += w
		textures = append(textures, t)
	}

	// Compute inter-widget space if any.
	spacing := int32(0)
	if l.HorizontalAlign == align.Justified && l.width > totalWidth {
		spacing = (l.width - int32(totalWidth)) / int32(len(l.children)+1)
	}

	// Render to our texture, vertically center each child.
	renderer.SetRenderTarget(l.texture)

	// Offset contents Y value depending on alignment.
	x := l.alignX(totalWidth)

	for _, t := range textures {
		x += spacing

		_, _, w, h, _ := t.Query()
		renderer.Copy(t, nil, &sdl.Rect{
			X: x,
			Y: l.alignY(h),
			W: w,
			H: h,
		})

		x += w
	}
}
