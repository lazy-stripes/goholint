package widgets

type DialogResult uint8

const (
	DialogOK DialogResult = iota
	DialogCancel
)

type DialogCloser func(DialogResult)

type Dialog interface {
	Widget

	OnClose(cb DialogCloser)
}

type dialog struct {
	closer DialogCloser
}

func (d *dialog) OnClose(cb DialogCloser) {
	d.closer = cb
}
