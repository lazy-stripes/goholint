package lcd

import (
	"fmt"
	"hash/crc32"
	"image"
	"image/draw"
	"image/gif"
	"os"
)

// FrameDelay is the time each GIF frame lasts, given that the Gameboy's screen
// is refreshed at 59.7Hz. In 100ths of a second (which is about 1.7 but we
// might add that up before we round it to integer).
const FrameDelay = (1 / 59.7) * 100

// FrameBounds holds fixed bounds for each frame.
var FrameBounds = image.Rectangle{Min: image.Point{0, 0},
	Max: image.Point{X: ScreenWidth, Y: ScreenHeight}}

// GIF display plugging into SDL to generate animated images on the fly.
type GIF struct {
	SDL
	gif.GIF

	File string

	frame *image.Paletted // Current frame
	delay float32         // Current frame's delay

	disabled    *image.Paletted // Disabled screen frame
	disabledCRC uint32          // CRC of disabled frame

	lastCRC uint32 // Previous frame's CRC.
}

// NewGIF returns an SDL2 display with a greyish palette and takes a zoom
// factor to size the window (current default is 2x). This will also
// buffer frames to put in a GIF.
func NewGIF(filename string, zoomFactor uint8) *GIF {
	// TODO: check file access, pre-create it.

	// Pre-instanciate diabled screen frame.
	disabled := image.NewPaletted(FrameBounds, DefaultPalette)
	draw.Draw(disabled, disabled.Bounds(), &image.Uniform{DefaultPalette[0]}, image.ZP, draw.Src)
	middle := disabled.Bounds()
	middle.Min.Y /= 2
	middle.Max.Y = (middle.Max.Y / 2) + 1
	draw.Draw(disabled, middle, &image.Uniform{DefaultPalette[3]}, image.ZP, draw.Src)
	disabledCRC := crc32.ChecksumIEEE(disabled.Pix)

	config := image.Config{ColorModel: disabled.ColorModel(),
		Width: ScreenWidth, Height: ScreenHeight}
	gif := GIF{SDL: *NewSDL(zoomFactor),
		disabled:    disabled,
		disabledCRC: disabledCRC,
		GIF:         gif.GIF{Config: config},
		frame:       image.NewPaletted(FrameBounds, DefaultPalette),
		File:        filename}
	return &gif
}

// Clear draws a disabled GB screen. This is a fixed frame.
func (g *GIF) Clear() {
	g.SDL.Clear()
	if g.lastCRC == g.disabledCRC {
		g.delay += FrameDelay
		g.GIF.Delay[len(g.GIF.Delay)] = int(g.delay)
	} else {
		g.delay = FrameDelay
		g.lastCRC = g.disabledCRC
		g.GIF.Image = append(g.GIF.Image, g.disabled)
		g.GIF.Delay = append(g.GIF.Delay, int(g.delay))
	}
}

// Write adds a new pixel to the current GIF frame.
func (g *GIF) Write(colorIndex uint8) {
	g.SDL.Write(colorIndex)
	if g.SDL.enabled {
		//g.frame.Pix = append(g.frame.Pix, colorIndex)
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
		frameCRC := crc32.ChecksumIEEE(g.frame.Pix)
		if g.lastCRC == frameCRC {
			g.delay += FrameDelay
			g.GIF.Delay[len(g.GIF.Delay)-1] = int(g.delay)
		} else {
			//fmt.Printf("Storing new frame at address %p (pix=%p)\n", g.frame, g.frame.Pix)
			g.delay = 2 // GIF players poorly handle 10ms frames delay
			g.lastCRC = frameCRC
			g.GIF.Image = append(g.GIF.Image, g.frame)
			//fmt.Printf("g.GIF.Image: len=%d, last=%p\n", len(g.GIF.Image), g.GIF.Image[len(g.GIF.Image)-1])
			//if f, err := os.Create(fmt.Sprintf("%02d-%s", len(g.GIF.Image), g.File)); err == nil {
			//		defer func() {
			//			f.Close()
			//		}()
			//		gif.Encode(f, g.frame, &gif.Options{NumColors: 4})
			//		fmt.Println("Frame dumped.")
			//	}
			g.GIF.Delay = append(g.GIF.Delay, int(g.delay))
			g.frame = image.NewPaletted(FrameBounds, DefaultPalette)
			//fmt.Printf("New frame. Delay=%d\n", int(g.delay))
			g.lastCRC = frameCRC
		}
	} else {
		g.Clear()
	}
}

// Close writes the actual GIF file to disk.
func (g *GIF) Close() {
	if f, err := os.Create(g.File); err == nil {
		defer func() {
			f.Close()
		}()
		gif.EncodeAll(f, &g.GIF)
		fmt.Printf("%d frames dumped to %s\n", len(g.GIF.Image), g.File)
	}
}
