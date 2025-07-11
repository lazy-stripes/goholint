package widgets

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"time"
	"unsafe"

	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ppu"
	"github.com/lazy-stripes/goholint/ppu/states"
	"github.com/lazy-stripes/goholint/screen"
	"github.com/lazy-stripes/goholint/ui/widgets/align"
	"github.com/lazy-stripes/goholint/utils"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/draw"
)

// Initialize sub-logger for unmapped MMU accesses.
func init() {
	log.Add("screen", "show sprite boundaries (Debug level only)")
}

// Screen represents the LCD display for a GameBoy. It works by shifting out
// individual pixels to a single dedicated texture.
type Screen struct {
	*widget

	PPU *ppu.PPU // For debugging

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

	statesCallbacks [4][]func() // Lists of callbacks indexed by PPU state.

	newPalette []color.RGBA // Requested new palette, will be set next VBlank.
	palette    []color.RGBA // Current palette.

	///Rectangle image.Rectangle

	// GIF recorder. TODO: record video with sound too.
	gif            *screen.GIF
	startRecording bool
	stopRecording  bool
	recordTime     time.Time     // Record timer.
	recordSkipped  time.Duration // Time to skip after unpausing.
}

// NewScreen returns a widget suitable for use as a Gameboy display (conforming
// to the screen.Display interface) and supporting screenshots, palette changes,
// GIF recording and overlay messages.
func NewScreen(sizeHint *sdl.Rect, config *options.Options) *Screen {
	// Ignore size hint for main texture, Gameboy's screen is 160×144 pixels.
	w, h := options.ScreenWidth, options.ScreenHeight

	layoutProps := DefaultProperties
	layoutProps.VerticalAlign = align.Bottom
	layoutProps.HorizontalAlign = align.Left
	layoutProps.Padding = int32(config.ZoomFactor)

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
		// Request UI to repaint itself since VBlank won't do it anymore.
		sdl.PushEvent(&sdl.RenderEvent{Type: sdl.RENDER_TARGETS_RESET})
	}
}

func (s *Screen) Enabled() bool {
	return s.enabled
}

func (s *Screen) drawDisabled() {
	// Fill the front buffer with background color.
	img := s.frontBuffer
	bg := (color.RGBA)(s.BgColor)
	fg := (color.RGBA)(s.FgColor)
	draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
	bar := img.Bounds() // Draw a bar in the middle of the frame
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
	s.overlay.Insert(s.message)

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

	// Save the duration of our GIF recording so far. We'll use it when we
	// unpause so whatever time was spent paused isn't counted in the recording
	// timer.
	if s.gif.IsOpen() {
		s.recordSkipped = time.Since(s.recordTime)
	}

	// Turn front buffer to greyscale, resize and blur.
	bounds := s.frontBuffer.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			pixel := s.frontBuffer.RGBAAt(x, y)
			r, g, b, a := pixel.R, pixel.G, pixel.B, pixel.A
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			grey := uint8(lum)
			s.frontBuffer.SetRGBA(x, y, color.RGBA{grey, grey, grey, a})
		}
	}

	// Resize to texture size before blurring.
	src := s.frontBuffer
	dstRect := image.Rect(0, 0, int(s.size.W), int(s.size.H))
	dst := image.NewRGBA(dstRect)
	draw.NearestNeighbor.Scale(dst, dstRect, src, src.Bounds(), draw.Over, nil)

	blurred := blur(blur(blur(dst)))
	rawPixels := unsafe.Pointer(&blurred.Pix[0])
	s.texture.Update(nil, rawPixels, blurred.Stride)

	s.paused = true
}

func (s *Screen) Unpause() {
	if s.gif.IsOpen() {
		s.recordTime = time.Now().Add(-s.recordSkipped)
	}

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

func (s *Screen) State(state states.State) {
	switch state {
	case states.HBlank:
	case states.VBlank:
		// Specifically invoke VBlank code in the UI thread if we want to display
		// messages and such.
		sdl.Do(s.vblank)
	case states.OAMSearch:
	case states.PixelTransfer:
	}
	s.invokeCallbacks(state)
}

func (s *Screen) invokeCallbacks(state states.State) {
	// Invoke all stored callbacks and clear slice.
	for _, cb := range s.statesCallbacks[state] {
		sdl.Do(cb)
	}
	s.statesCallbacks[state] = s.statesCallbacks[state][:0]
}

// OnState takes a callback function that will be invoked once when the PPU
// reaches the given state. This is mostly used to ensure some operations are
// only performed at specific times, like VBlank.
//
// The given callback is stored into an internal list. When a new state is
// reached through a call to State(), all callbacks in the list will be invoked
// in the order they were added.
//
// If the screen is currently disabled, no state change can occur so the given
// callback will be invoked immediately.
func (s *Screen) OnState(state states.State, callback func()) {
	if s.enabled {
		s.statesCallbacks[state] = append(s.statesCallbacks[state], callback)
	} else {
		// State won't change so we do the callback now.
		sdl.Do(callback)
	}
}

// vblank writes the contents of the back buffer to the internal widget texture,
// updates GIF recording if needed, and applies a new palette if requested.
//
// This method is called when the PPU reaches VBlank (by invoking State() with
// the value associated to VBlank).
func (s *Screen) vblank() {
	// Swap buffers and update texture (possibly with extra debug stuff).
	s.frontBuffer, s.backBuffer = s.backBuffer, s.frontBuffer

	rawPixels := unsafe.Pointer(&s.frontBuffer.Pix[0])
	s.screen.Update(nil, rawPixels, s.frontBuffer.Stride)
	// TODO: maybe having a proper RenderTo(texture) that would take care of
	// the target might help.
	renderer.SetRenderTarget(s.texture)
	renderer.Copy(s.screen, nil, nil)

	if log.Sub("screen").Enabled() && logger.Level >= logger.Debug {
		s.drawSpriteBorders()
	}

	// Don't draw overlay if not needed.
	if s.text != nil || s.message != nil {
		overlayTexture := s.overlay.Texture()
		renderer.SetRenderTarget(s.texture)
		renderer.Copy(overlayTexture, nil, nil)
	}

	// Without this, blurring no longer works, and I do not understand why. x_x
	renderer.SetRenderTarget(nil)

	// Reset offset for drawing the next frame.
	s.offset = 0

	// Update GIF frame if recording. FIXME: I'd like doing this from the UI, like SaveFrame(screen.Frame).
	if s.gif.IsOpen() {
		d := time.Since(s.recordTime)
		text := fmt.Sprintf("•REC [%02d:%02d]", d/time.Minute, d/time.Second)
		s.Text(text)
		s.gif.SaveFrame()
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
}

// drawSpriteBorders renders sprite boundaries on top of the widget's texture.
func (s *Screen) drawSpriteBorders() {

	var spriteRect sdl.Rect
	spriteRect.W = 8 * int32(s.Zoom)
	if s.PPU.LCDC&ppu.LCDCSpriteSize != 0 {
		spriteRect.H = 16 * int32(s.Zoom)
	} else {
		spriteRect.H = 8 * int32(s.Zoom)
	}

	renderer.SetRenderTarget(s.texture)
	oam := s.PPU.OAM.Bytes
	for i := 0; i < len(oam); i += 4 {
		y := int32(oam[i+0]) - 16
		x := int32(oam[i+1]) - 8
		//tile := oam[i+2]
		flags := oam[i+3]

		spriteRect.X = x * int32(s.Zoom)
		spriteRect.Y = y * int32(s.Zoom)
		renderer.SetDrawColor(0xff, 0x00, 0x00, 0xff)
		renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		renderer.DrawRect(&spriteRect)

		// Add orientation given flip flags.
		x = spriteRect.X + 1
		y = spriteRect.Y + 1
		w := int32(8)

		opX := int32(+1)
		opY := int32(+1)
		if flags&ppu.SpriteFlipX != 0 {
			x = spriteRect.X + spriteRect.W - 2
			opX = -1
		}
		if flags&ppu.SpriteFlipY != 0 {
			y = spriteRect.Y + spriteRect.H - 2
			opY = -1
		}

		for i := int32(0); i <= w; i++ {
			renderer.SetDrawColor(0x80, 0x00, 0x00, 0xff/2)
			renderer.DrawLine(
				x,
				y,
				x,
				y+(opY*(w-i)),
			)
			x += opX
		}

	}
	renderer.SetRenderTarget(nil)
}

// Dump writes the current pixel buffer to file for debugging purposes.
func (s *Screen) Dump() {
	ioutil.WriteFile("lcd-buffer-dump.bin", s.frontBuffer.Pix, 0644)
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
