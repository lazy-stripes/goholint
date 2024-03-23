package widgets

import "github.com/veandco/go-sdl2/sdl"

// Group adds a list of children to widgets embedding it, as well as an Add()
// method to append children to the list, and event dispatching.
type Group struct {
	*widget

	children []Widget // List of sub-widgets
}

// ProcessEvent calls ProcessEvent for each child widget until one of them
// returns true. If none return true, false is returned.
func (g *Group) ProcessEvent(e Event) bool {
	// Propagate event processing until a sub-widget catches it.
	for _, c := range g.children {
		if c.ProcessEvent(e) {
			return true
		}
	}
	return false
}

// Add appends a child widget to the internal list of children.
func (g *Group) Add(child Widget) {
	g.children = append(g.children, child)
}

func NewGroup(rect *sdl.Rect, children ...Widget) *Group {
	g := Group{
		widget:   new(rect),
		children: children,
	}
	return &g
}
