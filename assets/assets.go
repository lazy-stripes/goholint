package assets

import (
	_ "embed"

	"github.com/veandco/go-sdl2/sdl"
)

// Embedded assets (window icon, UI font...) so we can just distribute a single
// binary file. Thanks to Go >= 1.16 we can now do that pretty much natively
// using the embed package and the //go:embed directive. See the embed package
// documentation for details.

//go:embed boot.rom
var BootROM []byte

//go:embed icon.png
var icon []byte // Raw bytes for the window's icon.

//go:embed ui.ttf
var font []byte // Raw bytes for the UI's TTF font.

// Publicly accessible SDL buffers with our embedded assets.

var iconRW *sdl.RWops // Embedded buffer containing window icon.
var fontRW *sdl.RWops // Embedded buffer containing TTF font for UI.

// Init converts the raw embedded files into RW buffers that SDL can use.
func init() {
	var err error
	if iconRW, err = sdl.RWFromMem(icon); err != nil {
		panic(err)
	}
	if fontRW, err = sdl.RWFromMem(font); err != nil {
		panic(err)
	}
}

// WindowIconRW returns an RW buffer to the window icon that is ready and safe
// to use.
func WindowIconRW() *sdl.RWops {
	// Seek back to start first, in case it's been read before.
	iconRW.Seek(0, sdl.RW_SEEK_SET)
	return iconRW
}

// UIFontRW returns an RW buffer to the UI font that is ready and safe to use.
func UIFontRW() *sdl.RWops {
	// Seek back to start first, in case it's been read before.
	fontRW.Seek(0, sdl.RW_SEEK_SET)
	return fontRW
}
