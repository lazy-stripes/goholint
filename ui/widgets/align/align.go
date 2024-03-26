package align

type Align int

const (
	Left Align = iota
	Right
	Center
	Top
	Bottom
	Middle
	Justified
)

type Axis int

const (
	Horizontal Axis = 1 << iota
	Vertical
)
