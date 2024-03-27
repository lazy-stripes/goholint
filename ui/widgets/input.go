package widgets

import (
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/veandco/go-sdl2/sdl"
)

// Input widget that could be used for numbers, names, etc. In essence, a
// horizontal menu widget where entries can be changed.

const (
	charsetAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	charsetNum   = "0123456789"
	charsetHex   = charsetNum + "ABCDEF"
)

// character is a one-character label selectable from a larget charset.
type character struct {
	*Label

	charset  string // Available characters
	selected int    // Currently selected character in charset
}

func newChar(charset string) *character {
	props := DefaultProperties
	props.HorizontalAlign = align.Center
	props.VerticalAlign = align.Middle
	props.Margin = 10

	c := &character{
		// Auto-size label texture to its contents.
		Label:   NewLabel(noSizeHint, charset[0:1], props),
		charset: charset,
	}
	return c
}

func (c *character) highlight(v bool) {
	if v {
		c.Background = c.BgColor
	} else {
		c.Background = DefaultProperties.Background
	}
	c.repaint()
}

func (c *character) repaint() {
	// Update label before redrawing.
	c.text = c.charset[c.selected : c.selected+1]
	c.Label.repaint()
}

func (c *character) Next() {
	c.selected = (c.selected + 1) % len(c.charset)
	c.repaint()
}

func (c *character) Prev() {
	c.selected = (c.selected + len(c.charset) - 1) % len(c.charset)
	c.repaint()
}

type Input struct {
	*HorizontalLayout

	selected int // Index of selected character
}

// NewInput instantiates an Input widget with the given number of editable
// characters. Each character can be selected within the given charset. By
// default, the first character of the string given as a charset is used.
func NewInput(sizeHint *sdl.Rect, size int, charset string) *Input {
	in := &Input{
		HorizontalLayout: NewHorizontalLayout(sizeHint, nil),
	}
	in.HorizontalAlign = align.Center
	in.VerticalAlign = align.Middle

	for i := 0; i < size; i++ {
		c := newChar(charset)
		if i == 0 {
			c.highlight(true)
		}
		in.Add(c)
	}

	return in
}

// current returns the selected character instance from the internal list of
// children.
func (in *Input) current() *character {
	return in.children[in.selected].(*character)
}

// highlight sets/unsets the background for the currently selected character.
func (in *Input) highlight(v bool) {
	in.current().highlight(v)
}

func (in *Input) ProcessEvent(e Event) bool {
	switch e {
	case ButtonUp:
		in.current().Next()
	case ButtonDown:
		in.current().Prev()
	case ButtonLeft:
		in.Prev()
	case ButtonRight:
		in.Next()
	case ButtonA:
		// TODO Next char, keep currently selected character
	case ButtonB:
		// TODO Previous char, keep currently selected character
	case ButtonSelect:
		// ?
	case ButtonStart:
		//in.Confirm()
	default:
		// Unknown event, not handled.
		return false
	}

	// Refresh texture if something changed.
	in.repaint()

	return true
}

func (in *Input) Prev() {
	in.highlight(false)
	in.selected = (in.selected + len(in.children) - 1) % len(in.children)
	in.highlight(true)
}

func (in *Input) Next() {
	in.highlight(false)
	in.selected = (in.selected + 1) % len(in.children)
	in.highlight(true)
}
