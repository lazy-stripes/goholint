package screen

import (
	"image"
	"image/color"
	"time"

	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/options"
)

// Package-wide logger.
var log = logger.New("screen", "actual pixel display operations")

func init() {
	log.Add("gif", "GIF generator operations")
}

// ColorIndex into a display-defined 4-color palette.
type ColorIndex uint8

// Palette containing 4 indexed colors.
type Palette [4]color.NRGBA

type PixelWriter interface {
	Write(colorIndex uint8)
}

// Screen represents the LCD display for a GameBoy. It works by shifting out
// individual pixels to a single dedicated texture.
// TODO: before anything else, find in which package this thing should live.
// TODO: Maybe just a PixelWriter interface here with a Write method?
type Screen struct {
	config *options.Options
	//ui     *ui.UI

	palette    []color.RGBA
	newPalette []color.RGBA // Store new value until next frame

	enabled   bool
	buffer    []byte // Texture buffer for each frame
	blank     []byte // Static texture buffer for "blank screen" frames
	offset    int
	zoom      int // Zoom factor applied to the 144×160 screen.
	Rectangle image.Rectangle

	// Set this to true to save the next frame. Will be reset at VBlank.
	screenshotRequested bool

	// GIF recorder. TODO: record video with sound too.
	gif            *GIF
	startRecording bool
	stopRecording  bool
	recordTime     time.Time
}
