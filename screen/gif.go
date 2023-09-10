package screen

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"

	"github.com/lazy-stripes/goholint/options"
)

// TODO: this could probably also be moved to its own package.

// FrameDelay is the time each GIF frame lasts, given that the Gameboy's screen
// is refreshed at 59.7Hz. In 100ths of a second (which is about 1.7 but we
// might add that up before we round it to integer).
// In any event, browsers seem to ignore any value of 0 or 1 (or more depending
// on sources) so delay will be initialized at 2 for new frames.
const FrameDelay = (1 / 59.7) * 100

// FrameBounds holds fixed bounds for each frame.
var FrameBounds = image.Rectangle{
	Min: image.Point{0, 0},
	Max: image.Point{X: options.ScreenWidth, Y: options.ScreenHeight},
}

// GIF recorder generating animated images on the fly.
type GIF struct {
	gif.GIF

	config *options.Options

	palette color.Palette // Pre-instanciated Palette from our RGBA array.

	Filename string
	fd       *os.File

	frame     *image.Paletted // Current frame
	lastFrame *image.Paletted // Previous frame
	delay     float32         // Current frame's delay
	offset    uint            // Current frame's current pixel offset

	disabled *image.Paletted // Disabled screen frame
}

// NewGIF instantiates a GIF recorder that will buffer frames and then output a
// GIF file when required.
func NewGIF(config *options.Options) *GIF {
	// TODO: save path in config
	return &GIF{
		config: config,
	}
}

// Write adds a new pixel to the current GIF frame.
func (g *GIF) Write(colorIndex uint8) {
	g.frame.Pix[g.offset] = colorIndex
	g.offset++
}

// SaveFrame adds the current frame to GIF slice and pre-instantiate next. We
// detect if the display was disabled. If so, save a "disabled screen" frame
// instead.
// TODO: stream saved frames to disk.
func (g *GIF) SaveFrame() {
	// Pixel offset should be at the very end of the frame. If not, screen was
	// off and we save the "disabled" frame instead.
	var currentFrame *image.Paletted
	if g.offset == 0 {
		currentFrame = g.disabled
	} else {
		currentFrame = g.frame
	}

	// If current frame is the same as the previous one, only update delay of
	// the latest frame.
	if g.lastFrame != nil && bytes.Equal(currentFrame.Pix, g.lastFrame.Pix) {
		g.delay += FrameDelay
		g.GIF.Delay[len(g.GIF.Delay)-1] = int(g.delay)
	} else {
		g.delay = FrameDelay
		g.lastFrame = currentFrame
		g.GIF.Image = append(g.GIF.Image, g.frame)
		g.GIF.Delay = append(g.GIF.Delay, 2) // GIF players poorly handle 10ms frames delay
		g.frame = image.NewPaletted(FrameBounds, g.palette)
	}

	g.offset = 0
}

// IsOpen returns true if GIF recording is already in progress (i.e. we have a
// file currently open) or false otherwise.
func (g *GIF) IsOpen() bool {
	return g.fd != nil
}

// New starts recording a new GIF file to the provided descriptor and starts recording screen output. This should
// be called at VBlank time to prevent incomplete frames.
func (g *GIF) New(file *os.File, palette []color.RGBA) {
	if g.IsOpen() {
		log.Sub("gif").Warning("GIF recording already in progress, closing it.")
		g.Close()
	}

	log.Sub("gif").Infof("recording to %s", file.Name())

	// Convert the current palette's RGBA array to Color interface slice.
	g.palette = []color.Color{palette[0], palette[1], palette[2], palette[3]}

	// Pre-instanciate "disabled screen" frame with current palette.
	g.disabled = image.NewPaletted(FrameBounds, g.palette)
	draw.Draw(g.disabled, g.disabled.Bounds(), &image.Uniform{g.palette[0]}, image.Point{}, draw.Src)

	// Dimensions and colors for generated GIF file.
	gifConfig := image.Config{
		ColorModel: g.disabled.ColorModel(),
		Width:      options.ScreenWidth,
		Height:     options.ScreenHeight,
	}

	g.GIF = gif.GIF{Config: gifConfig}
	g.frame = image.NewPaletted(FrameBounds, g.palette)
	g.lastFrame = nil
	g.Filename = file.Name()
	g.fd = file
	g.offset = 0
}

// Close writes the actual GIF file to disk.
func (g *GIF) Close() {
	g.SaveFrame()
	defer func() {
		g.fd.Close()
		g.fd = nil
	}()
	gif.EncodeAll(g.fd, &g.GIF)
	log.Sub("gif").Infof("%d frames dumped to %s", len(g.GIF.Image), g.Filename)
}
