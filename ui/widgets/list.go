package widgets

import "github.com/veandco/go-sdl2/sdl"

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
func NewList(sizeHint *sdl.Rect, items []ListItem, props ...Properties) *List {
	// Figure out how many labels at most we can display. Pre-instantiate those
	// labels, we'll repaint them as we go.
	itemsPerPage := int(sizeHint.H) / DefaultProperties.TitleFont.Height()
	labels := make([]Widget, itemsPerPage)
	hint := *sizeHint
	hint.H = int32(DefaultProperties.TitleFont.Height())

	// TODO: reserve width to the right for scrollbar if needed.

	for i := 0; i < itemsPerPage; i++ {
		labels = append(labels, NewLabel(&hint, ""))
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
