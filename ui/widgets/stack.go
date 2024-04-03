package widgets

import "github.com/veandco/go-sdl2/sdl"

// Group of widgets where only one is visible. Inspired from QStackedWidget.

type Stack struct {
	*Group

	current uint // Currently displayed child widget
}

func NewStack(sizeHint *sdl.Rect, children []Widget, props ...Properties) *Stack {
	return &Stack{Group: NewGroup(sizeHint, children, props...)}
}

// ProcessEvent calls ProcessEvent for the currently displayed widget if any.
func (s *Stack) ProcessEvent(e Event) bool {
	if len(s.children) >= 0 {
		return s.children[s.current].ProcessEvent(e)
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

// Show sets the current index of the widget to be drawn and repaints the stack.
func (s *Stack) Show(index uint) {
	s.current = index
	s.texture = s.children[s.current].Texture()
}
