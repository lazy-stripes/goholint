package widgets

import (
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/veandco/go-sdl2/sdl"
)

// menuChoice embedds a label widget and an action function on confirm.
type menuChoice struct {
	*Label
	action func()
}

// Menu widget displaying a list of potential choices, each of which should map
// to some kind of Action.
type Menu struct {
	*VerticalLayout

	selected int // Index of selected choice

	labelSizeHint sdl.Rect   // Cached sizehint for adding choices.
	labelProps    Properties // Cached properties for adding choices.
}

func NewMenu(sizeHint *sdl.Rect) *Menu {
	props := DefaultProperties
	props.HorizontalAlign = align.Center
	props.VerticalAlign = align.Justified

	m := Menu{
		VerticalLayout: NewVerticalLayout(sizeHint, nil, props),
	}

	// Cache labelSizeHint and label props for AddChoice
	props.Margin = int32(DefaultProperties.Zoom * 8)
	props.VerticalAlign = align.Middle
	m.labelProps = props
	m.labelSizeHint = *sizeHint
	m.labelSizeHint.H = int32(DefaultProperties.TitleFont.Height())

	return &m
}

func (m *Menu) AddChoice(title string, action func()) {
	m.Add(&menuChoice{
		Label:  NewLabel(&m.labelSizeHint, title, m.labelProps),
		action: action,
	})
	m.repaint()
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

// current returns the selected character instance from the internal list of
// children.
// FIXME: lots of common code with input, I should write a SelectableLayout or something.
func (m *Menu) current() *menuChoice {
	return m.children[m.selected].(*menuChoice)
}

// XXX: Maybe wrap label in a common type that provides .highlight()
func (m *Menu) highlightCurrent(v bool) {
	label := m.current()
	if v {
		label.Background = label.BgColor
	} else {
		label.Background = DefaultProperties.Background
	}
	label.repaint()
}

func (m *Menu) Select(index uint) {
	if index < uint(len(m.children)) {
		m.highlightCurrent(false)
		m.selected = int(index)
		m.highlightCurrent(true)
		m.repaint()
	}
}

func (m *Menu) Up() {
	m.highlightCurrent(false)
	if m.selected > 0 {
		m.selected -= 1
	}
	// TODO: else, blink? How? widget.Animate(...)?
	m.highlightCurrent(true)
}

func (m *Menu) Down() {
	m.highlightCurrent(false)
	if m.selected < len(m.children)-1 {
		m.selected += 1
	}
	m.highlightCurrent(true)
}

func (m *Menu) Confirm() {
	m.current().action()
}
