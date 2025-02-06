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

	selected int // Currently selected item

	// Paging/scrolling.
	itemsPerPage int
	page         int
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
	itemsPerPage := int(sizeHint.H) / props.Font.Height()
	if itemsPerPage > len(items) {
		itemsPerPage = len(items)
	}

	labels := make([]Widget, 0, itemsPerPage)
	hint := *sizeHint
	hint.H = int32(props.Font.Height())

	// TODO: reserve width to the right for scrollbar if needed.

	for i := 0; i < itemsPerPage; i++ {
		labels = append(labels, NewLabel(&hint, items[i].Text(), props))
	}

	l := List{
		VerticalLayout: NewVerticalLayout(sizeHint, labels),
		items:          items,
		itemsPerPage:   itemsPerPage,
	}
	l.repaint()

	return &l
}

func (l *List) repaint() {
	// TODO: only render currently visible items.
	// .. circular buffer of labels?
	for idx, item := range l.items[:l.itemsPerPage] {
		label := l.children[idx].(*Label)
		label.text = item.Text() // XXX I should probably have an actual .SetText() by now.

		if idx == l.selected {
			label.Background = label.BgColor // TODO: label.Highlight()
		}

		label.repaint()
	}

	l.VerticalLayout.repaint()
}

func (l *List) Selected() ListItem {
	if len(l.items) > 0 {
		return l.items[l.selected]
	}
	return nil
}
