package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Group of widgets where only one is visible. Inspired from QStackedWidget.

type Stack struct {
	*Group

	current Widget // Currently displayed child widget
}

func NewStack(sizeHint *sdl.Rect, children []Widget, props ...Properties) *Stack {
	s := Stack{Group: NewGroup(sizeHint, children, props...)}
	if len(children) > 0 {
		s.Show(0)
	}
	return &s
}

// ProcessEvent calls ProcessEvent for the currently displayed widget if any.
func (s *Stack) ProcessEvent(e Event) bool {
	if s.current != nil {
		return s.current.ProcessEvent(e)
	}
	return false
}

// Add appends the given widget to the stack's internal children, and shows it
// if it's the very first child to be added.
func (s *Stack) Add(w Widget) {
	s.Group.Add(w)
	if len(s.children) == 1 {
		s.Show(0)
	}
}

// Texture updates the internal texture with the currently shown child (if any)
// then calls the base class.
func (s *Stack) Texture() *sdl.Texture {
	if s.current != nil {
		s.texture = s.current.Texture()
	}
	return s.widget.Texture()
}

// Show sets the current index of the widget to be drawn and repaints the stack.
func (s *Stack) Show(index uint) {
	if index >= uint(len(s.children)) {
		log.Warningf("Stack.Show(%d) out of bounds (%d)", index, len(s.children))
	}
	s.current = s.children[index]
}

// ShowWidget sets the index of the given widget to be drawn and repaints the
// stack, if the given widget is a child of the Stack.
func (s *Stack) ShowWidget(w Widget) {
	for _, child := range s.children {
		if child == w {
			s.current = child
			return
		}
	}
}
