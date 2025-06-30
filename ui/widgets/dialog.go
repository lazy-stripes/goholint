package widgets

type DialogResult uint8

const (
	DialogOK DialogResult = iota
	DialogCancel
)

type DialogCloser func(DialogResult)

type DialogWidget interface {
	Widget

	OnClose(callback DialogCloser)
	Close(DialogResult)
	Confirm()
	Cancel()
}

type Dialog struct {
	closer DialogCloser
}

func NewDialog(cb DialogCloser) *Dialog {
	return &Dialog{closer: cb}
}

func (d *Dialog) OnClose(cb DialogCloser) {
	d.closer = cb
}

func (d *Dialog) Close(res DialogResult) {
	if d.closer != nil {
		d.closer(res)
	}
}

func (d *Dialog) Confirm() {
	d.Close(DialogOK)
}

func (d *Dialog) Cancel() {
	d.Close(DialogCancel)
}

// ProcessEvent handles button presses and will call either Confirm() on A or
// Start press, or Cancel() on B or Select.
func (d *Dialog) ProcessEvent(e Event) bool {
	switch e {
	case ButtonA:
		d.Confirm()
	case ButtonB:
		d.Cancel()
	case ButtonSelect:
		d.Cancel()
	case ButtonStart:
		d.Confirm()
	default:
		// Let caller handle this event.
		return false
	}

	return true
}
