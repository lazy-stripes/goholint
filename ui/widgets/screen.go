package widgets

import (
	"image"
	"image/color"
	"io/ioutil"
	"time"
	"unsafe"

	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/screen"
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/veandco/go-sdl2/sdl"
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
	msgTimer *time.Timer // Timer for clearing messages
	message  string      // Temporary text on timer
	text     string      // Permanent text

	buffer []byte       // Texture buffer for each frame.
	frozen *sdl.Texture // Blurred grayscale version of GB display.
	screen *sdl.Texture // Gameboy display texture (160×144)

	vblankCallbacks []func() // List of callbacks to invoke at VBlank time.

	newPalette []color.RGBA // Requested new palette, will be set next VBlank.
	palette    []color.RGBA // Current palette.

	blank  []byte // Static texture buffer for "blank screen" frames.
	offset int
	///Rectangle image.Rectangle

	// Set this to true to save the next frame. Will be reset at VBlank.
	screenshotRequested bool

	// GIF recorder. TODO: record video with sound too.
	gif            *screen.GIF
	startRecording bool
	stopRecording  bool
	recordTime     time.Time
}

// New returns an SDL2 display with a greyish palette and takes a zoom
// factor to size the window (current default is 2x).
func NewScreen(sizeHint *sdl.Rect, config *options.Options) *Screen {
	// Go bindings use byte slices but SDL thinks in terms of uint32
	screenLen := options.ScreenWidth * options.ScreenHeight * 4
	blank := make([]byte, screenLen) // TODO: phase out, just redraw blank to internal texture if/when needed.

	// Ignore size hint for main texture, Gameboy's screen is 160×144 pixels.
	screenRect := sdl.Rect{W: options.ScreenWidth, H: options.ScreenHeight}

	// XXX For testing
	//props := DefaultProperties
	//props.Border = 1
	//props.BorderColor = sdl.Color{R: 255, A: 255}

	layoutProps := DefaultProperties
	layoutProps.VerticalAlign = align.Bottom
	layoutProps.HorizontalAlign = align.Left

	s := Screen{
		widget:  new(sizeHint),
		overlay: NewVerticalLayout(sizeHint, nil, layoutProps),
		screen:  texture(&screenRect),
		config:  config,
		palette: config.Palettes[0],
		buffer:  make([]byte, options.ScreenWidth*options.ScreenHeight*4),
		blank:   blank,
		///Rectangle: screenRect,
		gif: screen.NewGIF(config),
	}

	return &s
}

// Set permanent text (useful for persistent UI). Call with empty string to
// clear.
func (s *Screen) Text(text string) {
	s.text = text
}

// Clear temporary message. Texture will be repainted next VBlank.
func (s *Screen) clearMessage() {
	s.message = ""
}

// Message shows a temporary message that will be cleared after the given
// duration (in seconds). The message stacks with permanent text set via Text().
func (s *Screen) Message(text string, seconds time.Duration) {
	// Stop reset timer, a new one will be started.
	// TODO: stack messages (up to, like, 3 or something)
	if s.msgTimer != nil {
		s.msgTimer.Stop()
	}
	s.message = text
	s.msgTimer = time.AfterFunc(time.Second*seconds, s.clearMessage)
}

// makeBlank prepares a static texture buffer to represent the screen when it
// is off.
func (s *Screen) makeBlank() {
	col := s.palette[3] // Background color
	for offset := 0; offset < len(s.blank); offset += 4 {
		s.blank[s.offset+0] = col.R
		s.blank[s.offset+1] = col.G
		s.blank[s.offset+2] = col.B
		s.blank[s.offset+3] = col.A
	}
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

	// Dimensions of UI screen.
	// FIXME: this should all be deduced from the widget texture, but scaling up the gb screen is where the friction happens.
	width := int(options.ScreenWidth * s.config.ZoomFactor)
	height := int(options.ScreenHeight * s.config.ZoomFactor)

	// Intermediate image for easier blurring.
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Map source offset (in 160×144 space) to the current UI pixel.
			srcX := x / int(s.config.ZoomFactor)
			srcY := y / int(s.config.ZoomFactor)
			srcOffset := (srcY * options.ScreenWidth * 4) + (srcX * 4)

			// Extract RGB, compute greyscale, strore in work image.
			r := s.buffer[srcOffset+0]
			g := s.buffer[srcOffset+1]
			b := s.buffer[srcOffset+2]
			a := s.buffer[srcOffset+3]
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			grey := uint8(lum)

			img.SetRGBA(x, y, color.RGBA{grey, grey, grey, a})
		}
	}
	// Blur the background. Apply enough times for sufficient effect.
	// TODO: ... I could make the iterations and overlay configurable I guess?
	img = blur(blur(blur(img)))
	rawPixels := unsafe.Pointer(&img.Pix[0])
	s.frozen.Update(nil, rawPixels, width*4)

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
	s.buffer[s.offset+0] = col.R
	s.buffer[s.offset+1] = col.G
	s.buffer[s.offset+2] = col.B
	s.buffer[s.offset+3] = col.A

	// If all goes well, we'll get VBlank'ed just as we wrap up.
	s.offset = (s.offset + 4) % len(s.buffer)

	if s.gif.IsOpen() {
		s.gif.Write(colorIndex)
	}
}

// clear overrides the VerticalLayout method to draw the gameboy screen to the
// background texture.
func (s *Screen) clear() {
	renderer.SetRenderTarget(s.texture)
	if s.paused {
		renderer.Copy(s.frozen, nil, nil)
	} else {
		rawPixels := unsafe.Pointer(&s.buffer[0])
		s.screen.Update(nil, rawPixels, options.ScreenWidth*4)
		renderer.Copy(s.screen, nil, nil)
	}
	renderer.SetRenderTarget(nil)
}

// Texture will draw the screen and optionally the text overlay on top.
func (s *Screen) Texture() *sdl.Texture {
	// If paused, only show the blurred background instead.
	if !s.paused {
		rawPixels := unsafe.Pointer(&s.buffer[0])
		s.screen.Update(nil, rawPixels, options.ScreenWidth*4)
		// TODO: maybe having a proper RenderTo(texture) that would take care of the target might help.
		renderer.SetRenderTarget(s.texture)
		renderer.Copy(s.screen, nil, nil)

		// TODO: don't draw overlay if not needed.
		overlayTexture := s.overlay.Texture()
		renderer.SetRenderTarget(s.texture)
		renderer.Copy(overlayTexture, nil, nil)
		renderer.SetRenderTarget(nil)
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
	s.vblankCallbacks = append(s.vblankCallbacks, callback)
}

// VBlank is called when the PPU reaches VBlank state. At this point, our SDL
// buffer should be ready to display.
func (s *Screen) VBlank() {
	//if s.enabled {
	//	// Reset offset for drawing the next frame.
	//	s.offset = 0
	//}
	//
	//// Update GIF frame if recording. We do this before checking startRecording
	//// otherwise the call to SaveFrame will always insert a "disabled" frame in
	//// first position (since we haven't yet had time to build a full frame in
	//// that specific case).
	//// FIXME: timer behavior when pausing the emulator. I most likely need to move something to ui package. Or use the GameBoy timer itself.
	//if s.gif.IsOpen() {
	//	d := time.Since(s.recordTime)
	//	text := fmt.Sprintf("•REC [%02d:%02d]", d/time.Minute, d/time.Second)
	//	s.ui.Text(text)
	//	s.gif.SaveFrame()
	//}
	//
	//// Create GIF here if requested.
	//if s.startRecording {
	//	f, err := options.CreateFileIn("gifs", ".gif")
	//	if err == nil {
	//		s.startRecording = false
	//		s.recordTime = time.Now()
	//		s.ui.Text("•REC [00:00]")
	//		s.gif.New(f, s.palette)
	//
	//		fmt.Printf("Recording GIF to %s\n", f.Name())
	//	} else {
	//		log.Warningf("creating gif file failed: %v", err)
	//	}
	//}
	//
	//if s.stopRecording {
	//	s.stopRecording = false
	//	s.gif.Close()
	//	s.ui.Text("")
	//	s.ui.Message(fmt.Sprintf("%d frames saved", len(s.gif.GIF.Image)), 2)
	//}
	//
	//if s.screenshotRequested {
	//	s.screenshotRequested = false
	//
	//	f, err := options.CreateFileIn("screenshots", ".png")
	//	if err != nil {
	//		log.Warningf("creating screenshot file failed: %v", err)
	//		return
	//	}
	//	defer f.Close()
	//
	//	// Populate image from buffer, taking zoom into account.
	//	img := image.NewRGBA(s.Rectangle)
	//	for x := 0; x < img.Rect.Dx(); x++ {
	//		for y := 0; y < img.Rect.Dy(); y++ {
	//			srcX := x / s.zoom
	//			srcY := y / s.zoom
	//			srcOffset := srcY*options.ScreenWidth*4 + srcX*4
	//			dstOffset := y*options.ScreenWidth*s.zoom*4 + x*4
	//
	//			// Copy RGBA components.
	//			img.Pix[dstOffset+0] = s.buffer[srcOffset+0]
	//			img.Pix[dstOffset+1] = s.buffer[srcOffset+1]
	//			img.Pix[dstOffset+2] = s.buffer[srcOffset+2]
	//			img.Pix[dstOffset+3] = s.buffer[srcOffset+3]
	//		}
	//	}
	//
	//	if err := png.Encode(f, img); err != nil {
	//		log.Warningf("saving screenshot failed: %v", err)
	//		return
	//	}
	//
	//	s.ui.Message("Screenshot saved", 2)
	//	fmt.Printf("Screenshot saved to %s\n", f.Name())
	//
	//	// Semi-hack to dump RAM and debug Marioland. In time it should be made
	//	// into a "Game Genie" kind of feature. For now, this will do.
	//	if logger.Level >= logger.Debug {
	//		// TODO: I'd like to be able to call some cpu.DumpRAM() here to make
	//		//       sure I'm getting the exact RAM state for the current frame
	//		//       but scope is being an issue.
	//		//       Maybe through some debug.DumpRAM() where debug would hold
	//		//       the necessary references. Meh.
	//	}
	//}
	//

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
	ioutil.WriteFile("lcd-buffer-dump.bin", s.buffer, 0644)
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
