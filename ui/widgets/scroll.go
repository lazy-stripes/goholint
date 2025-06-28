package widgets

import "github.com/veandco/go-sdl2/sdl"

// Widget including a scrollbar for paginated contents. The scrollbar isn't
// meant to be interacted with but just shows where you are in a list.

// Scrollable is an interface providing methods to check how a scrollbar should
// be drawn. It uses a current page number and a total number of pages to
// compute the size of the scrollbar and its position.
//
// If the scrollable widget only shows one page, the scrollbar won't be drawn.
type Scrollable interface {
	Widget

	// Page returns the current page in the child contents. This should be an
	// integer in the [0, Pages[ range.
	Page() int

	// Pages returns the total number of pages used by the child widget. If
	// this value is 1, no scrollbar will be drawn.
	Pages() int

	// DrawScrollbar renders the current scrollbar on top of the given widget.
	DrawScrollbar(dst Scrollable)
}

// TODO: horizontal scroll too.
// TODO: better idea maybe: ScrollArea? Still can't figure out how to deal with sizing.

type Scroll struct {
	Scrollable
}

func NewScroll(child Scrollable) *Scroll {
	return &Scroll{Scrollable: child}
}

type Scrollbar struct{}

func (s *Scrollbar) DrawScrollbar(dst Scrollable) {
	t := dst.Texture()

	if dst.Pages() < 2 {
		// No scrollbar for contents with a single page.
		return
	}

	// Draw scrollbar on top of child widget.
	props := dst.Props()
	w := (4 + props.Border*2) * int32(props.Zoom)
	h := dst.Size().H

	// Draw bar border.
	renderer.SetRenderTarget(t)
	renderer.SetDrawColor(
		props.BorderColor.R,
		props.BorderColor.G,
		props.BorderColor.B,
		props.BorderColor.A,
	)
	rect := sdl.Rect{
		X: dst.Size().W - w,
		Y: 0,
		W: w,
		H: h,
	}
	renderer.DrawRect(&rect)

	// Draw actual bar.
	rect.H = h / int32(dst.Pages())
	rect.Y = h * int32(dst.Page())
	renderer.DrawRect(&rect)
	renderer.SetRenderTarget(nil)

}
