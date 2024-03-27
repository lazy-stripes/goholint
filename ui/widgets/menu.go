package widgets

import (
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/veandco/go-sdl2/sdl"
)

// MenuChoice groups a choice string with the corresponding action.
type MenuChoice struct {
	Text   string
	Action func()
}

// item wraps a Label with a margin and an optional background for selected
// items. TODO: add margin and background options to Label, obviously.
type item struct {
	*widget

	label *Label
}

func newItem(s *sdl.Rect, text string) *item {
	// Item texture should we font-high and window-wide.
	//viewPort := sdl.Rect{H: int32(DefaultProperties.Font.Height()), W: s.W}
	l := NewLabel(noSizeHint, text)

	// Create item texture with margin.
	margin := DefaultProperties.Zoom * 2
	_, _, _, h, _ := l.Texture().Query()

	itemSize := sdl.Rect{
		X: 0,
		Y: 0,
		W: s.W,
		H: h + int32(margin*2),
	}
	item := &item{
		widget: new(&itemSize),
		label:  l,
	}

	return item
}

func (i *item) highlight(v bool) {
	if v {
		i.Background = i.BgColor
		i.label.Background = i.BgColor
	} else {
		i.Background = DefaultProperties.Background
		i.label.Background = DefaultProperties.Background
	}
	i.label.repaint()
}

// Texture renders the label and an optional background if the item is selected.
func (i *item) Texture() *sdl.Texture {
	// Render transparent or filled (selected) background.
	i.clear()

	// Render label on top of it.
	labelTexture := i.label.Texture()
	_, _, w, h, _ := labelTexture.Query()
	labelTexture.SetBlendMode(sdl.BLENDMODE_BLEND)
	renderer.SetRenderTarget(i.texture)
	renderer.Copy(labelTexture, nil, &sdl.Rect{
		X: (i.width - w) / 2,                 // Center text, this should probably be in widgets.Label too.
		Y: int32(DefaultProperties.Zoom * 2), // Margin
		W: w,
		H: h,
	})
	renderer.SetRenderTarget(nil)
	return i.widget.Texture()
}

// Menu widget displaying a list of potential choices, each of which should map
// to some kind of Action.
type Menu struct {
	*VerticalLayout

	items    []*item
	choices  []MenuChoice
	selected int // Index of selected choice
}

func NewMenu(sizeHint *sdl.Rect, choices []MenuChoice) *Menu {
	props := DefaultProperties
	props.VerticalAlign = align.Middle
	layout := NewVerticalLayout(sizeHint, nil, props)
	var items []*item
	for i, c := range choices {
		item := newItem(sizeHint, c.Text)
		items = append(items, item)
		layout.Add(item)
		// TODO: phase out items, use layout.Add(NewLabel(c.Text)), change label background

		// Pre-select first item in list.
		if i == 0 {
			item.Background = item.BgColor
		}
	}

	return &Menu{
		VerticalLayout: layout,
		items:          items,
		choices:        choices,
		selected:       0,
	}
}

func (m *Menu) ProcessEvent(e Event) bool {
	switch e {
	case ButtonUp:
		m.Up()
	case ButtonDown:
		m.Down()
	case ButtonA:
		m.Confirm()
	case ButtonB:
		m.Confirm()
	case ButtonSelect:
		m.Confirm()
	case ButtonStart:
		m.Confirm()
	default:
		// Unknown event, not handled.
		return false
	}

	// Refresh texture if something changed.
	m.repaint()

	return true
}

func (m *Menu) Up() {
	m.items[m.selected].highlight(false)
	if m.selected > 0 {
		m.selected -= 1
	}
	// TODO: else, blink? How? widget.Animate(...)?
	m.items[m.selected].highlight(true)
}

func (m *Menu) Down() {
	m.items[m.selected].highlight(false)
	if m.selected < len(m.choices)-1 {
		m.selected += 1
	}
	m.items[m.selected].highlight(true)
}

func (m *Menu) Confirm() {
	m.choices[m.selected].Action()
}
