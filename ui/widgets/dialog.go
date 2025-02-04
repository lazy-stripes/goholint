package widgets

type DialogResult uint8

const (
	DialogOK DialogResult = iota
	DialogCancel
)

type DialogCloser func(DialogResult)

type Dialog struct {
	*widget

	closer DialogCloser // Callback when closing dialog.
}

func (d *Dialog) Show(closer DialogCloser) {
	d.closer = closer
	d.SetVisible(true)
}

func (d *Dialog) Close(result DialogResult) {
	d.SetVisible(false)
	if d.closer != nil {
		d.closer(result)
		d.closer = nil
	}
}
