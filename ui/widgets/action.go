package widgets

type Action interface {
	Text() string
	Trigger()
}

// action is a base embeddable struct for actial Action types.
type action struct {
	text      string
	onTrigger func()
}

func newAction(text string, onTrigger func()) *action {
	return &action{text: text, onTrigger: onTrigger}
}
