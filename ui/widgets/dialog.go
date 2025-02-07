package widgets

type DialogResult uint8

const (
	DialogOK DialogResult = iota
	DialogCancel
)

type DialogWidget interface {
	Widget

	OnClose(callback DialogCloser)
	Confirm()
	Cancel()
}

type DialogCloser func(DialogResult)

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
