package assets

import (
	_ "embed"

	"github.com/veandco/go-sdl2/sdl"
)

// Embedded assets (window icon, UI font...) so we can just distribute a single
// binary file. Thanks to Go >= 1.16 we can now do that pretty much natively
// using the embed package and the //go:embed directive. See the embed package
// documentation for details.

//go:embed icon.png
var icon []byte // Raw bytes for the window's icon.

//go:embed ui.ttf
var font []byte // Raw bytes for the UI's TTF font.

// Publicly accessible SDL buffers with our embedded assets.

var WindowIcon *sdl.RWops // Embedded buffer containing window icon.
var UIFont *sdl.RWops     // Embedded buffer containing TTF font for UI.

// Init converts the raw embedded files into RW buffers that SDL can use.
func init() {
	var err error
	if WindowIcon, err = sdl.RWFromMem(icon); err != nil {
		panic(err)
	}
	if UIFont, err = sdl.RWFromMem(font); err != nil {
		panic(err)
	}
}
