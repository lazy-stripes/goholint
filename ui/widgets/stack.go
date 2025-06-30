package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Stack of widgets only showing and forwarding events to the top one.
type Stack struct {
	*Group
}

func NewStack(sizeHint *sdl.Rect, children []Widget, props ...Properties) *Stack {
	return &Stack{Group: NewGroup(sizeHint, children, props...)}
}

// ProcessEvent calls ProcessEvent for the currently displayed widget if any.
func (s *Stack) ProcessEvent(e Event) bool {
	if len(s.children) > 0 {
		return s.children[len(s.children)-1].ProcessEvent(e)
	}
	return false
}

// Push appends the given widget to the internal list, effectively pushing the
// widget to the top of the stack.
func (s *Stack) Push(w Widget) {
	s.Group.Add(w)
}

// Pop returns the last widget in the internal list and removes it, effectively
// popping it off the top of the stack.
func (s *Stack) Pop() (w Widget) {
	if len(s.children) > 0 {
		last := len(s.children) - 1
		w = s.children[last]
		s.children = s.children[:last]
	}
	return w
}

// Texture updates the internal texture with the currently shown child (if any)
// then calls the base class.
func (s *Stack) Texture() *sdl.Texture {
	if len(s.children) > 0 {
		last := len(s.children) - 1
		s.texture = s.children[last].Texture()
	}
	return s.widget.Texture()
}
