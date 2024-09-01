package widgets

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"time"
	"unsafe"

	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/screen"
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/lazy-stripes/goholint/utils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/draw"
)

// Screen represents the LCD display for a GameBoy. It works by shifting out
// individual pixels to a single dedicated texture.
type Screen struct {
	*widget

	// FIXME: I use this virtually everywhere... do I dare a global options.Runtime instance?
	config *options.Options

	paused bool

	// Overlay messages.
	overlay  *VerticalLayout
	msgTimer *time.Timer // Timer for clearing messages.
	message  *Label      // Temporary text on timer.
	text     *Label      // Permanent text.

	enabled     bool
	screen      *sdl.Texture // Gameboy display texture (160×144).
	backBuffer  *image.RGBA  // Work buffer for the frame in progress.
	frontBuffer *image.RGBA  // Buffer for the displayed frame.
	offset      int          // Current pixel offset in frame.

	vblankCallbacks []func() // List of callbacks to invoke at VBlank time.

	newPalette []color.RGBA // Requested new palette, will be set next VBlank.
	palette    []color.RGBA // Current palette.

	///Rectangle image.Rectangle

	// Set this to true to save the next frame. Will be reset at VBlank.
	screenshotRequested bool

	// GIF recorder. TODO: record video with sound too.
	gif            *screen.GIF
	startRecording bool
	stopRecording  bool
	recordTime     time.Time
}

// NewScreen returns a widget suitable for use as a Gameboy display (conforming
// to the screen.Display interface) and supporting screenshots, palette changes,
// GIF recording and overlay messages.
func NewScreen(sizeHint *sdl.Rect, config *options.Options) *Screen {
	// Ignore size hint for main texture, Gameboy's screen is 160×144 pixels.
	w, h := options.ScreenWidth, options.ScreenHeight

	// XXX For testing
	//props := DefaultProperties
	//props.Border = 1
	//props.BorderColor = sdl.Color{R: 255, A: 255}

	layoutProps := DefaultProperties
	layoutProps.VerticalAlign = align.Bottom
	layoutProps.HorizontalAlign = align.Left

	s := Screen{
		widget:      new(sizeHint),
		overlay:     NewVerticalLayout(sizeHint, nil, layoutProps),
		screen:      texture(&sdl.Rect{W: int32(w), H: int32(h)}),
		backBuffer:  image.NewRGBA(image.Rect(0, 0, w, h)),
		frontBuffer: image.NewRGBA(image.Rect(0, 0, w, h)),
		palette:     config.Palettes[0],
		gif:         screen.NewGIF(config),
		config:      config,
	}

	s.drawDisabled()

	return &s
}

func (s *Screen) Enable(enabled bool) {
	s.enabled = enabled
	if !enabled {
		s.drawDisabled()
		sdl.PushEvent(&sdl.RenderEvent{Type: sdl.RENDER_TARGETS_RESET})
	}
}

func (s *Screen) Enabled() bool {
	return s.enabled
}

func (s *Screen) drawDisabled() {
	// Fill the front buffer with background color.
	img := s.frontBuffer
	bg := color.RGBA{ // XXX SDL and Go types don't mingle that well
		s.BgColor.R, // and it sucks that sdl.Color doesn't implement RGBA()
		s.BgColor.G,
		s.BgColor.B,
		s.BgColor.A,
	}
	fg := color.RGBA{
		s.FgColor.R,
		s.FgColor.G,
		s.FgColor.B,
		s.FgColor.A,
	}
	draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
	bar := img.Bounds() // Middle of frame
	bar.Min.Y = bar.Max.Y / 2
	bar.Max.Y = bar.Min.Y + 1
	draw.Draw(img, bar, &image.Uniform{fg}, image.Point{}, draw.Src)
}

// Frame returns the current front buffer. This allows the UI to grab whatever
// is currently being displayed for GIFs of screenshots without having to worry
// about VBlank.
func (s *Screen) Frame() *image.RGBA {
	return s.frontBuffer
}

// Set permanent text (useful for persistent UI). Call with empty string to
// clear.
func (s *Screen) Text(text string) {
	if s.text != nil {
		// Do not recreate label if text hasn't changed.
		if s.text.text == text {
			return
		}

		s.overlay.Remove(s.text)
		s.text.Destroy()
		s.text = nil
	}

	if text != "" {
		s.text = NewLabel(noSizeHint, text)
		s.overlay.Add(s.text)
	}
}

// Clear temporary message. Texture will be repainted next VBlank.
func (s *Screen) clearMessage() {
	// TODO: might need a lock here.
	if s.message != nil {
		s.overlay.Remove(s.message)
		s.message.Destroy()
		s.message = nil
	}
}

// Message shows a temporary message that will be cleared after the given
// duration (in seconds). The message stacks with permanent text set via Text().
func (s *Screen) Message(text string, secs time.Duration) {
	// Stop reset timer, a new one will be started.
	// TODO: stack messages (up to, like, 3 or something)
	if s.msgTimer != nil && s.msgTimer.Stop() {
		s.clearMessage()
	}

	s.message = NewLabel(noSizeHint, text)
	s.overlay.Add(s.message)

	s.msgTimer = time.AfterFunc(secs*time.Second, utils.WrapSDL(s.clearMessage))
}

func averagePixels(pixels []color.RGBA) (avg color.RGBA) {
	var sumR, sumG, sumB int
	for _, pixel := range pixels {
		sumR += int(pixel.R)
		sumG += int(pixel.G)
		sumB += int(pixel.B)
	}

	avg = color.RGBA{
		uint8(sumR / len(pixels)),
		uint8(sumG / len(pixels)),
		uint8(sumB / len(pixels)),
		0xff,
	}

	return avg
}

// blur returns a copy of the image after applying the box blur algorithm to it.
// Image has to be at least 2px×2px, or you will have a bad time.
func blur(img *image.RGBA) (blurred *image.RGBA) {
	blurred = image.NewRGBA(img.Bounds())

	// Apply blur to inner pixels (radius is 1 pixel).
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	for x := 1; x < w-1; x++ {
		for y := 1; y < h-1; y++ {
			neighbors := []color.RGBA{
				img.RGBAAt(x-1, y+1), // Top left
				img.RGBAAt(x+0, y+1), // Top center
				img.RGBAAt(x+1, y+1), // Top right
				img.RGBAAt(x-1, y+0), // Mid left
				img.RGBAAt(x+0, y+0), // Current pixel
				img.RGBAAt(x+1, y+0), // Mid right
				img.RGBAAt(x-1, y-1), // Low left
				img.RGBAAt(x+0, y-1), // Low center
				img.RGBAAt(x+1, y-1), // Low right
			}

			avg := averagePixels(neighbors)
			blurred.SetRGBA(x, y, avg)

			// Duplicate left column of blurred pixels.
			if x == 1 {
				blurred.SetRGBA(0, y, avg)
			}

			// Duplicate right column of blurred pixels.
			if x == w-2 {
				blurred.SetRGBA(w-1, y, avg)
			}

			// Duplicate top row of blurred pixels.
			if y == 1 {
				blurred.SetRGBA(x, 0, avg)
			}

			// Duplicate bottom row of blurred pixels.
			if y == h-2 {
				blurred.SetRGBA(x, h-1, avg)
			}
		}
	}

	// Copy corner pixels.
	blurred.SetRGBA(0, 0, img.RGBAAt(0, 0))
	blurred.SetRGBA(w, 0, img.RGBAAt(w, 0))
	blurred.SetRGBA(0, h, img.RGBAAt(0, h))
	blurred.SetRGBA(w, h, img.RGBAAt(w, h))

	return blurred
}

// Pause is called whenever the emulator is paused. This method takes a copy of
// the current GameBoy screen and turns it to blurred greyscale for use as a
// background in the main UI.
func (s *Screen) Pause() {
	if s.paused {
		return
	}

	// Turn front buffer to greyscale, resize and blur.
	bounds := s.frontBuffer.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			pixel := s.frontBuffer.At(x, y).(color.RGBA)
			r, g, b, a := pixel.R, pixel.G, pixel.B, pixel.A
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			grey := uint8(lum)
			s.frontBuffer.SetRGBA(x, y, color.RGBA{grey, grey, grey, a})
		}
	}

	// Resize to texture size before blurring.
	src := s.frontBuffer
	dstRect := image.Rect(0, 0, int(s.width), int(s.height))
	dst := image.NewRGBA(dstRect)
	draw.NearestNeighbor.Scale(dst, dstRect, src, src.Bounds(), draw.Over, nil)

	blurred := blur(blur(blur(dst)))
	rawPixels := unsafe.Pointer(&blurred.Pix[0])
	s.texture.Update(nil, rawPixels, int(s.width)*4)

	s.paused = true
}

func (s *Screen) Unpause() {
	s.paused = false
}

// Close frees allocated resources.
func (s *Screen) Close() {
	if s.gif.IsOpen() {
		s.gif.Close()
	}
}

// Write adds a new pixel (a mere index into a palette) to the texture buffer.
func (s *Screen) Write(colorIndex uint8) {
	if s.paused {
		return
	}

	col := s.palette[colorIndex]
	s.backBuffer.Pix[s.offset+0] = col.R
	s.backBuffer.Pix[s.offset+1] = col.G
	s.backBuffer.Pix[s.offset+2] = col.B
	s.backBuffer.Pix[s.offset+3] = col.A

	// If all goes well, we'll get VBlank'ed just as we wrap up.
	s.offset = (s.offset + 4) % len(s.backBuffer.Pix)

	if s.gif.IsOpen() {
		s.gif.Write(colorIndex)
	}
}

// Texture will draw the screen and optionally the text overlay on top.
func (s *Screen) Texture() *sdl.Texture {
	// If paused, only show the blurred background instead.
	if !s.paused {
		rawPixels := unsafe.Pointer(&s.frontBuffer.Pix[0])
		s.screen.Update(nil, rawPixels, s.frontBuffer.Stride)
		// TODO: maybe having a proper RenderTo(texture) that would take care of
		// the target might help.
		renderer.SetRenderTarget(s.texture)
		renderer.Copy(s.screen, nil, nil)

		// Don't draw overlay if not needed.
		if s.text != nil || s.message != nil {
			overlayTexture := s.overlay.Texture()
			renderer.SetRenderTarget(s.texture)
			renderer.Copy(overlayTexture, nil, nil)
			renderer.SetRenderTarget(nil)
		}
	}
	return s.widget.Texture()
}

// OnVBlank takes a callback function that will be invoked once when VBlank() is
// called. Use this method to ensure certain operations only happen when a
// screen frame has been fully drawn.
//
// The given callback is stored into an internal list. At the end of VBlank, all
// callbacks in the list will be invoked in the order they were given.
func (s *Screen) OnVBlank(callback func()) {
	if s.enabled {
		s.vblankCallbacks = append(s.vblankCallbacks, callback)
	} else {
		// We may not get vblank at all, do the thing now.
		callback()
	}
}

// VBlank is called when the PPU reaches VBlank state. At this point, our SDL
// buffer should be ready to display.
func (s *Screen) VBlank() {
	// Specifically invoke VBlank code in the UI thread if we want to display
	// messages and such.
	sdl.Do(s.vblank)
}

func (s *Screen) vblank() {
	// Swap buffers.
	s.frontBuffer, s.backBuffer = s.backBuffer, s.frontBuffer

	// Reset offset for drawing the next frame.
	s.offset = 0

	// Update GIF frame if recording.
	// FIXME: timer behavior when pausing the emulator. I most likely need to move something to ui package. Or use the GameBoy timer itself.
	if s.gif.IsOpen() {
		d := time.Since(s.recordTime)
		text := fmt.Sprintf("•REC [%02d:%02d]", d/time.Minute, d/time.Second)
		s.Text(text)
		s.gif.SaveFrame() // TODO: SaveFrame(frontBuffer) instead of using Write
	}

	// Create GIF here if requested.
	if s.startRecording {
		f, err := options.CreateFileIn("gifs", ".gif")
		if err == nil {
			s.startRecording = false
			s.recordTime = time.Now()
			s.Text("•REC [00:00]")
			s.gif.Open(f, s.palette)

			fmt.Printf("Recording GIF to %s\n", f.Name())
		} else {
			log.Warningf("creating gif file failed: %v", err)
		}
	}

	if s.stopRecording {
		s.stopRecording = false
		s.gif.Close()
		s.Text("")
		s.Message(fmt.Sprintf("%d frames saved", len(s.gif.GIF.Image)), 2)
	}

	// Apply new palette if one was requested.
	if s.newPalette != nil {
		s.palette = s.newPalette
		s.newPalette = nil
	}

	// Invoke all stored callbacks and clear slice.
	for _, cb := range s.vblankCallbacks {
		cb()
	}
	s.vblankCallbacks = s.vblankCallbacks[:0]
}

// Dump writes the current pixel buffer to file for debugging purposes.
func (s *Screen) Dump() {
	ioutil.WriteFile("lcd-buffer-dump.bin", s.frontBuffer.Pix, 0644)
}

// Screenshot will make the display dump the next frame to file.
func (s *Screen) Screenshot() {
	s.screenshotRequested = true
}

// StartRecord will create a new GIF file and output frames into it until
// StopRecord is called.
//
// We only just raise a flag here, recording should start and stop in VBlank.
func (s *Screen) StartRecord() {
	if s.gif.IsOpen() {
		log.Warningf("recording to %s already in progress", s.gif.Filename)
		return
	}
	s.startRecording = true
}

// StopRecord will flush recorded frames to the previously created GIF file.
// We only just raise a flag here, recording should start and stop in VBlank.
func (s *Screen) StopRecord() {
	s.stopRecording = true
}

// Palette will set a new palette for the display and GIF. The new palette will
// only go into effect at VBlank time.
func (s *Screen) Palette(p []color.RGBA) {
	s.newPalette = p
}
