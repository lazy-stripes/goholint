package widgets

import (
	"io/fs"
	"os"
	"path"

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
	dialog
	*List
}

func NewFileDialog(sizeHint *sdl.Rect, dir string, p ...Properties) *FileDialog {
	entries, err := os.ReadDir(dir)
	log.Warningf("error creating FileDialog: %s", err)

	var items []ListItem
	for _, e := range entries {
		item := &FileItem{
			DirEntry: e,
			baseDir:  dir,
		}
		items = append(items, item)
	}

	fd := &FileDialog{
		List: NewList(sizeHint, items, p...),
	}

	return fd
}
