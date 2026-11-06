package memory

import (
	"errors"
	"io/fs"
	"os"
)

// RAM as an array of R/W bytes at addresses starting from a given offset.
// If dealing with a battery-backed RAM, this can also be saved to a file.
type RAM struct {
	*Memory

	saveFile string // For battery-backed RAM chips
}

func NewRAM(start uint16, size int) *RAM {
	return &RAM{Memory: New(start, size)}
}

// Load sets the current content of RAM from the given file, and stores the
// path to that file for subsequent saves.
func (r *RAM) Load(filename string) error {
	if len(r.saveFile) > 0 && r.saveFile != filename {
		log.Warningf("calling Load(%s) on RAM with an existing save file (%s)",
			filename, r.saveFile)
	}
	oldSavefile := r.saveFile
	r.saveFile = filename

	err := r.Memory.Load(filename) // The file may not yet exist, it's fine.
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		// Reset saveFile because this may well be an innocent mistake and we
		// don't want that file overwritten later.
		r.saveFile = oldSavefile
		return err
	}

	return nil
}

// Save dumps the current content of RAM into the associated save file (if any).
func (r *RAM) Save() error {
	if r.saveFile == "" {
		return errors.New("trying to Save() RAM with no save file defined")
	}

	return os.WriteFile(r.saveFile, r.Bytes, 0644)
}
