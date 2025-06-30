package widgets

import (
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/veandco/go-sdl2/sdl"
)

type Selectable interface {
	Widget

	Highlight(bool)

	// TODO: Select(bool)
	// TODO: Selected() bool
}

type ListItem interface {
	Text() string
	Value() any
}

// Scrollable vertical list with selectable text items.
// TODO: build menus out of this at some point.
type List struct {
	*VerticalLayout

	items []ListItem // TODO: richer type

	// Paging/scrolling.
	pages        int
	itemsPerPage int

	page     int // Current page
	selected int // Currently selected item
}

// TODO: items should have methods to get label/value from them. Like, you know, actions.
func NewList(sizeHint *sdl.Rect, items []ListItem, p ...Properties) *List {
	props := DefaultProperties // TODO: getProps() helper
	if len(p) > 0 {
		props = p[0]
	}

	props.HorizontalAlign = align.Left

	// Figure out how many labels at most we can display. Pre-instantiate those
	// labels, we'll repaint them as we go.
	// TODO: helper functions to (pre)compute all that.
	availableHeight := int(sizeHint.H - (props.Margin * 2))
	labelHeight := props.Font.Height() + (props.Zoom * 2) // Take outline into account (TODO: font.LabelHeight())

	// TODO: recompute on Add/Remove
	//       but this works better when called with a pre-existing list anyway.
	itemsPerPage := availableHeight / labelHeight
	pages := len(items) / itemsPerPage

	// Partial last page.
	if pages%itemsPerPage != 0 {
		pages += 1
	}
	if itemsPerPage > len(items) {
		itemsPerPage = len(items)
	}

	labels := make([]Widget, 0, itemsPerPage)
	hint := *sizeHint
	hint.H = int32(labelHeight)

	// TODO: reserve width to the right for scrollbar if needed.

	for i := 0; i < itemsPerPage; i++ {
		labels = append(labels, NewLabel(&hint, items[i].Text(), props))
	}

	l := List{
		VerticalLayout: NewVerticalLayout(sizeHint, labels),
		items:          items,
		pages:          pages,
		itemsPerPage:   itemsPerPage,
	}
	l.repaint()

	return &l
}

// repaint only renders the currently visible labels depending on pagination.
func (l *List) repaint() {
	startIdx := l.page * l.itemsPerPage
	endIdx := startIdx + l.itemsPerPage
	if endIdx > len(l.items) {
		endIdx = len(l.items)
	}
	for idx, item := range l.items[startIdx:endIdx] {
		label := l.children[idx].(*Label)
		label.text = item.Text() // XXX I should probably have an actual .SetText() by now.

		if idx == l.selected {
			label.Background = label.BgColor // TODO: label.Highlight()
		}

		label.repaint()
	}

	// Clear leftover labels on the last page.
	for idx := endIdx - startIdx; idx < l.itemsPerPage; idx++ {
		label := l.children[idx].(*Label)
		label.clear()
	}

	l.VerticalLayout.repaint()
}

func (l *List) Texture() *sdl.Texture {
	t := l.VerticalLayout.Texture()

	// Draw scrollbar.
	if l.Pages() < 2 {
		// No scrollbar for contents with a single page.
		return t
	}

	// Draw scrollbar on top of child widget.
	props := l.Props()
	w := (4 + props.Border*2) * int32(props.Zoom)
	h := l.Size().H

	renderer.SetRenderTarget(t)
	renderer.SetDrawColor(
		props.BgColor.R,
		props.BgColor.G,
		props.BgColor.B,
		props.BgColor.A,
	)

	// Draw actual bar. I currently think it looks better without a border.
	rect := sdl.Rect{
		X: l.Size().W - w,
		Y: (h / int32(l.Pages())) * int32(l.Page()),
		W: w,
		H: h / int32(l.Pages()),
	}

	renderer.SetRenderTarget(t)
	renderer.FillRect(&rect)
	renderer.SetRenderTarget(nil)

	return t
}

// Page returns the current page in the list. Used for scrollbars.
func (l *List) Page() int {
	return l.page
}

// Pages returns the total number of pages in the list. Used for scrollbars.
func (l *List) Pages() int {
	return l.pages
}

func (l *List) Selected() ListItem {
	if len(l.items) > 0 {
		idx := (l.page * l.itemsPerPage) + l.selected
		return l.items[idx]
	}
	return nil
}

func (m *List) ProcessEvent(e Event) bool {
	switch e {
	case ButtonUp:
		m.Up()
	case ButtonDown:
		m.Down()
	case ButtonLeft:
		m.PreviousPage()
	case ButtonRight:
		m.NextPage()
	default:
		// Unknown event, not handled.
		return false
	}

	// Refresh texture if something changed.
	m.repaint()

	return true
}

// current returns the selected label instance from the internal list of
// children.
func (l *List) current() *Label {
	return l.children[l.selected].(*Label)
}

// XXX: Maybe wrap label in a common type that provides .Select() and .Selected() (Selectable?)
func (l *List) highlightCurrent(v bool) {
	label := l.current()
	if v {
		label.Background = label.BgColor
	} else {
		label.Background = DefaultProperties.Background
	}
	label.repaint()
}

func (l *List) Select(index uint) {
	if index < uint(len(l.children)) {
		l.highlightCurrent(false)
		l.selected = int(index)
		l.highlightCurrent(true)
		l.repaint()
	}
}

func (l *List) Up() {
	l.highlightCurrent(false)
	if l.selected > 0 {
		l.selected -= 1
	} else {
		if l.page > 0 {
			l.page -= 1
			l.selected = l.itemsPerPage - 1
		}
	}
	// TODO: else, blink? How? widget.Animate(...)?
	l.highlightCurrent(true)
}

func (l *List) Down() {
	l.highlightCurrent(false)
	// Last page may not be full.
	maxIndex := l.itemsPerPage - 1
	if l.page == l.pages-1 {
		maxIndex = (len(l.items) % l.itemsPerPage) - 1
	}
	if l.selected < maxIndex {
		l.selected += 1
	} else {
		if l.page < l.pages-1 {
			l.page += 1
			l.selected = 0
		}
	}
	l.highlightCurrent(true)
}

func (l *List) PreviousPage() {
	l.highlightCurrent(false)
	if l.page > 0 {
		l.page -= 1
	} else {
		l.selected = 0
	}
	l.highlightCurrent(true)
}

func (l *List) NextPage() {
	l.highlightCurrent(false)
	lastIdx := (len(l.items) % l.itemsPerPage) - 1
	if l.page == l.pages-1 {
		l.selected = lastIdx
	} else {
		l.page += 1
		if (l.page == l.pages-1) && (l.selected > lastIdx) {
			l.selected = lastIdx
		}
	}
	l.highlightCurrent(true)
}
