package widgets

import (
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

	label    *Label
	selected bool
}

func newItem(s *sdl.Rect, text string) *item {
	l := NewLabel(text)

	// Create item texture with margin.
	margin := properties.Zoom * 2
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

// Texture renders the label and an optional background if the item is selected.
func (i *item) Texture() *sdl.Texture {
	// Render transparent or filled (selected) background.
	renderer.SetRenderTarget(i.texture)
	if i.selected {
		renderer.SetDrawColor(
			properties.BgColor.R,
			properties.BgColor.G,
			properties.BgColor.B,
			properties.BgColor.A,
		)
	} else {
		renderer.SetDrawColor(0, 0, 0, 0) // Transparent
	}
	renderer.Clear()

	// Render label on top of it.
	labelTexture := i.label.Texture()
	_, _, w, h, _ := labelTexture.Query()
	labelTexture.SetBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Copy(labelTexture, nil, &sdl.Rect{
		X: (i.width - w) / 2,          // Center text, this should probably be in widgets.Label too.
		Y: int32(properties.Zoom * 2), // Margin
		W: w,
		H: h,
	})
	renderer.SetRenderTarget(nil)
	return i.texture
}

// Menu widget displaying a list of potential choices, each of which should map
// to some kind of Action.
type Menu struct {
	*VerticalLayout

	items    []*item
	choices  []MenuChoice
	selected int // Index of selected choice
}

func NewMenu(s *sdl.Rect, choices []MenuChoice) *Menu {
	layout := NewVerticalLayout(s, nil)
	var items []*item
	for i, c := range choices {
		item := newItem(s, c.Text)
		items = append(items, item)
		layout.Add(item)
		// TODO: phase out items, use layout.Add(NewLabel(c.Text)), change label background

		// Pre-select first item in list.
		if i == 0 {
			item.selected = true
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
	return true
}

func (m *Menu) Up() {
	m.items[m.selected].selected = false

	if m.selected > 0 {
		m.selected -= 1
	}
	// TODO: else, blink? How?

	m.items[m.selected].selected = true
}

func (m *Menu) Down() {
	m.items[m.selected].selected = false

	if m.selected < len(m.choices)-1 {
		m.selected += 1
	}
	// TODO: blink?

	m.items[m.selected].selected = true
}

func (m *Menu) Confirm() {
	m.choices[m.selected].Action()
}
