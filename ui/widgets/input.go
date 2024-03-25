package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Input widget that could be used for numbers, names, etc. In essence, a
// horizontal menu widget where entries can be changed.

const (
	charsetAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNum   = "0123456789"
	charsetHex   = charsetNum + "ABCDEF"
)

// character is a one-character label selectable from a list of choices.
type character struct {
	*Label

	charset  string // Available characters
	selected int    // Currently selected character in charset
}

func newChar(charset string) *character {
	c := &character{
		Label:   NewLabel(noSizeHint, charset[0:1]),
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

func (c *character) ProcessEvent(e Event) bool {
	// TODO: common code for Prev/Next with Menu?
	switch e {
	case ButtonUp:
		c.Next()
	case ButtonDown:
		c.Prev()
	default:
		// Unknown event, not handled.
		return false
	}

	// Refresh texture if something changed.
	c.repaint()

	return true
}

func (c *character) Next() {
	c.selected = (c.selected + 1) % len(c.charset)
	c.text = c.charset[c.selected : c.selected+1]
}

func (c *character) Prev() {
	c.selected = (c.selected + len(c.charset) - 1) % len(c.charset)
	c.text = c.charset[c.selected : c.selected+1]
}

type Input struct {
	*HorizontalLayout

	chars    []*character
	selected int // Index of selected character

}

// NewInput instantiates an Input widget with the given number of editable
// characters. Each character can be selected within the given charset. By
// default, the first character of the string given as a charset is used.
func NewInput(sizeHint *sdl.Rect, size int, charset string) *Input {
	in := &Input{
		HorizontalLayout: NewHorizontalLayout(sizeHint),
	}

	for i := 0; i < size; i++ {
		c := newChar(charset)
		if i == 0 {
			c.highlight(true)
		}
		in.Add(c)
	}

	return in
}

func (in *Input) ProcessEvent(e Event) bool {
	switch e {
	case ButtonUp:
		in.children[in.selected].ProcessEvent(e)
	case ButtonDown:
		in.children[in.selected].ProcessEvent(e)
	case ButtonLeft:
		in.Prev()
	case ButtonRight:
		in.Next()
	case ButtonA:
		// Next char
	case ButtonB:
		// Previous char
	case ButtonSelect:
		//
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
	in.children[in.selected].(*character).highlight(false)
	in.selected = (in.selected + len(in.children) - 1) % len(in.children)
	in.children[in.selected].(*character).highlight(true)
}

func (in *Input) Next() {
	in.children[in.selected].(*character).highlight(false)
	in.selected = (in.selected + 1) % len(in.children)
	in.children[in.selected].(*character).highlight(true)
}
