package widgets

import (
	"io/fs"
	"os"
	"path"
	"regexp"

	"github.com/veandco/go-sdl2/sdl"
)

// TODO: wrap list and filedir entries.

type FileItem struct {
	fs.DirEntry

	baseDir string
}

func (i *FileItem) Text() string {
	return i.Name()
}

func (i *FileItem) Value() any {
	return path.Join(i.baseDir, i.Name())
}

type FileDialog struct {
	*Dialog
	*List
}

func NewFileDialog(sizeHint *sdl.Rect, dir string, p ...Properties) *FileDialog {
	entries, err := os.ReadDir(dir)
	log.Warningf("error creating FileDialog: %s", err)

	var items []ListItem
	for _, e := range entries {
		// Filter by extensions.
		// TODO: make it configurable obv.
		pattern := regexp.MustCompile(`\.(gb|zip)$`)
		if !pattern.MatchString(e.Name()) {
			continue
		}

		item := &FileItem{
			DirEntry: e,
			baseDir:  dir,
		}
		items = append(items, item)
	}

	fd := &FileDialog{
		Dialog: &Dialog{},
		List:   NewList(sizeHint, items, p...),
	}

	return fd
}

func (d *FileDialog) ProcessEvent(e Event) bool {
	switch e {
	case ButtonA:
		d.Confirm()
	case ButtonB:
		d.Cancel()
	case ButtonSelect:
		d.Confirm()
	case ButtonStart:
		d.Cancel()
	default:
		// Navigate list with other button inputs.
		return d.List.ProcessEvent(e)
	}

	return true
}
