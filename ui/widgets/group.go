package widgets

import "github.com/veandco/go-sdl2/sdl"

// Group adds a list of children to widgets embedding it, as well as an Add()
// method to append children to the list, and event dispatching.
type Group struct {
	*widget

	children []Widget // List of sub-widgets
}

func NewGroup(rect *sdl.Rect, children []Widget, props ...Properties) *Group {
	g := Group{
		widget:   new(rect, props...),
		children: children,
	}
	return &g
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

// Texture draws children texture in order (bottom to top).
func (g *Group) Texture() *sdl.Texture {
	g.clear()
	renderer.SetRenderTarget(g.texture)
	for _, c := range g.children {
		t := c.Texture()
		t.SetBlendMode(sdl.BLENDMODE_BLEND)
		renderer.Copy(t, nil, nil)
	}
	renderer.SetRenderTarget(nil)
	return g.texture
}

// Add appends a child widget to the internal list of children.
func (g *Group) Add(child Widget) {
	g.children = append(g.children, child)
}

// Remove looks for the given widget in the internal list of children and
// removes it if found.
func (g *Group) Remove(child Widget) {
	// See if child is in our list to begin with.
	index := -1
	for i, c := range g.children {
		if c == child {
			index = i
			break
		}
	}
	if index > -1 {
		// Make a new slice without the offending item.
		// FIXME: This is not optimized at all. Make it work first!
		// TODO: Actually what I really want is Widget.SetVisible(bool).
		newChildren := make([]Widget, 0)
		newChildren = append(newChildren, g.children[:index]...)
		newChildren = append(newChildren, g.children[index+1:]...)
		g.children = newChildren
	}
}

// Clear removes all children from the group and clears their textures.
func (g *Group) Clear() {
	for _, c := range g.children {
		c.Destroy()
	}
	g.children = nil
}
