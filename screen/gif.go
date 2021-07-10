package screen

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"
)

// FrameDelay is the time each GIF frame lasts, given that the Gameboy's screen
// is refreshed at 59.7Hz. In 100ths of a second (which is about 1.7 but we
// might add that up before we round it to integer).
// In any event, browsers seem to ignore any value of 0 or 1 (or more depending
// on sources) so delay will be initialized at 2 for new frames.
const FrameDelay = (1 / 59.7) * 100

// FrameBounds holds fixed bounds for each frame.
var FrameBounds = image.Rectangle{Min: image.Point{0, 0},
	Max: image.Point{X: ScreenWidth, Y: ScreenHeight}}

// GIF display plugging into SDL to generate animated images on the fly.
type GIF struct {
	SDL
	gif.GIF

	File string

	frame     *image.Paletted // Current frame
	lastFrame *image.Paletted // Previous frame
	delay     float32         // Current frame's delay

	disabled *image.Paletted // Disabled screen frame
}

// NewGIF returns an SDL2 display with a greyish palette and takes a zoom
// factor to size the window (current default is 2x). This will also
// buffer frames to put in a GIF.
func NewGIF(filename string, zoomFactor uint, noSync bool) *GIF {
	// TODO: check file access, (pre-create it?)

	// Pre-instanciate disabled screen frame.
	disabled := image.NewPaletted(FrameBounds, DefaultPalette)
	draw.Draw(disabled, disabled.Bounds(), &image.Uniform{DefaultPalette[0]}, image.Point{}, draw.Src)
	middle := disabled.Bounds()
	middle.Min.Y /= 2
	middle.Max.Y = (middle.Max.Y / 2) + 1
	draw.Draw(disabled, middle, &image.Uniform{DefaultPalette[3]}, image.Point{}, draw.Src)

	config := image.Config{
		ColorModel: disabled.ColorModel(),
		Width:      ScreenWidth,
		Height:     ScreenHeight,
	}

	return &GIF{
		SDL:       *NewSDL(zoomFactor, noSync),
		disabled:  disabled,
		GIF:       gif.GIF{Config: config},
		frame:     image.NewPaletted(FrameBounds, DefaultPalette),
		lastFrame: disabled, // Acceptable zero value to avoid a nil check later
		File:      filename,
	}
}

// Clear draws a disabled GB screen. This is a fixed frame.
func (g *GIF) Clear() {
	g.SDL.Clear()
	if g.lastFrame == g.disabled {
		g.delay += FrameDelay
		g.GIF.Delay[len(g.GIF.Delay)] = int(g.delay)
	} else {
		g.delay = FrameDelay
		g.lastFrame = g.disabled
		g.GIF.Image = append(g.GIF.Image, g.disabled)
		g.GIF.Delay = append(g.GIF.Delay, 2)
	}
}

// Write adds a new pixel to the current GIF frame.
func (g *GIF) Write(colorIndex uint8) {
	g.SDL.Write(colorIndex)
	if g.SDL.enabled {
		// SDL.Write already advanced its internal offset by 4.
		g.frame.Pix[(g.SDL.offset/4)-1] = colorIndex
	}
}

// HBlank not used yet. TODO: duplicate last pixel line up to ZoomFactor.
func (g *GIF) HBlank() {
}

// VBlank adds the current frame to GIF slice and pre-instantiate next.
func (g *GIF) VBlank() {
	g.SDL.VBlank()
	if g.SDL.enabled {
		// If current frame is the same as the previous one, only update delay.
		if bytes.Equal(g.frame.Pix, g.lastFrame.Pix) {
			g.delay += FrameDelay
			g.GIF.Delay[len(g.GIF.Delay)-1] = int(g.delay)
		} else {
			g.delay = FrameDelay
			g.lastFrame = g.frame
			g.GIF.Image = append(g.GIF.Image, g.frame)
			g.GIF.Delay = append(g.GIF.Delay, 2) // GIF players poorly handle 10ms frames delay
			g.frame = image.NewPaletted(FrameBounds, DefaultPalette)
		}
	} else {
		g.Clear()
	}
}

// Close writes the actual GIF file to disk.
func (g *GIF) Close() {
	g.VBlank()
	f, err := os.Create(g.File)
	if err == nil {
		defer func() {
			f.Close()
		}()
		gif.EncodeAll(f, &g.GIF)
		log.Sub("gif").Infof("%d frames dumped to %s", len(g.GIF.Image), g.File)
	} else {
		fmt.Println(err)
	}
}
