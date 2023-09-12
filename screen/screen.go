package screen

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"time"

	"github.com/lazy-stripes/goholint/logger"
	"github.com/lazy-stripes/goholint/options"
	"github.com/lazy-stripes/goholint/ui"
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

// Screen represents the LCD display for a GameBoy. It works by shifting out
// individual pixels to a single dedicated texture.
type Screen struct {
	config *options.Options
	ui     *ui.UI

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

// New returns an SDL2 display with a greyish palette and takes a zoom
// factor to size the window (current default is 2x).
func New(ui *ui.UI, config *options.Options) *Screen {
	// Request a buffer from UI, which will be used to draw the emulator's
	// output.
	buffer := ui.ScreenBuffer()

	// Go bindings use byte slices but SDL thinks in terms of uint32
	screenLen := options.ScreenWidth * options.ScreenHeight * 4
	blank := make([]byte, screenLen) // TODO: phase out, use a single buffer, redraw blank if/when needed.

	// Keep computed screen size for screenshots.
	screenRect := image.Rectangle{
		image.Point{0, 0},
		image.Point{
			options.ScreenWidth * int(config.ZoomFactor),
			options.ScreenHeight * int(config.ZoomFactor),
		},
	}

	s := Screen{
		config:    config,
		palette:   config.Palettes[0],
		ui:        ui,
		buffer:    buffer,
		blank:     blank,
		zoom:      int(config.ZoomFactor),
		Rectangle: screenRect,
		gif:       NewGIF(config),
	}

	// Pre-instantiate texture buffer for when the scren is off.
	s.makeBlank()

	// Init texture and trigger stuff usually happening at VBlank.
	s.VBlank() // XXX: is this needed?

	return &s
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

// Close frees allocated resources.
func (s *Screen) Close() {
	if s.gif.IsOpen() {
		s.gif.Close()
	}
}

// Enable turns on the display. Pixels will be drawn to our texture and showed
// at VBlank time.
func (s *Screen) Enable() {
	s.enabled = true
}

// Enabled returns whether the display is enabled or not (as part of the Display
// interface).
func (s *Screen) Enabled() bool {
	return s.enabled
}

// Disable turns off the display. A disabled GB screen will be drawn at VBlank
// time.
func (s *Screen) Disable() {
	s.offset = 0
	s.enabled = false
}

// Write adds a new pixel (a mere index into a palette) to the texture buffer.
func (s *Screen) Write(colorIndex uint8) {
	if s.enabled {
		col := s.palette[colorIndex]
		s.buffer[s.offset+0] = col.R
		s.buffer[s.offset+1] = col.G
		s.buffer[s.offset+2] = col.B
		s.buffer[s.offset+3] = col.A
		s.offset += 4

		if s.gif.IsOpen() {
			s.gif.Write(colorIndex)
		}
	}
}

// VBlank is called when the PPU reaches VBlank state. At this point, our SDL
// buffer should be ready to display.
func (s *Screen) VBlank() {
	// Refresh UI at the end of this function, which will draw the GameBoy
	// screen and whatever text overlays we use here.
	defer s.ui.Repaint()

	if s.enabled {
		// Reset offset for drawing the next frame.
		s.offset = 0
	}

	// Update GIF frame if recording. We do this before checking startRecording
	// otherwise the call to SaveFrame will always insert a "disabled" frame in
	// first position (since we haven't yet had time to build a full frame in
	// that specific case).
	// FIXME: timer behavior when pausing the emulator. I most likely need to move something to ui package.
	if s.gif.IsOpen() {
		d := time.Since(s.recordTime)
		text := fmt.Sprintf("•REC [%02d:%02d]", d/time.Minute, d/time.Second)
		s.ui.Text(text)
		s.gif.SaveFrame()
	}

	// Create GIF here if requested.
	if s.startRecording {
		f, err := options.CreateFileIn("gifs", ".gif")
		if err == nil {
			s.startRecording = false
			s.recordTime = time.Now()
			s.ui.Text("•REC [00:00]")
			s.gif.New(f, s.palette)

			fmt.Printf("Recording GIF to %s\n", f.Name())
		} else {
			log.Warningf("creating gif file failed: %v", err)
		}
	}

	if s.stopRecording {
		s.stopRecording = false
		s.gif.Close()
		s.ui.Text("")
		s.ui.Message(fmt.Sprintf("%d frames saved", len(s.gif.GIF.Image)), 2)
	}

	if s.screenshotRequested {
		s.screenshotRequested = false

		f, err := options.CreateFileIn("screenshots", ".png")
		if err != nil {
			log.Warningf("creating screenshot file failed: %v", err)
			return
		}
		defer f.Close()

		// Populate image from buffer, taking zoom into account.
		img := image.NewRGBA(s.Rectangle)
		for x := 0; x < img.Rect.Dx(); x++ {
			for y := 0; y < img.Rect.Dy(); y++ {
				srcX := x / s.zoom
				srcY := y / s.zoom
				srcOffset := srcY*options.ScreenWidth*4 + srcX*4
				dstOffset := y*options.ScreenWidth*s.zoom*4 + x*4

				// Copy RGBA components.
				img.Pix[dstOffset+0] = s.buffer[srcOffset+0]
				img.Pix[dstOffset+1] = s.buffer[srcOffset+1]
				img.Pix[dstOffset+2] = s.buffer[srcOffset+2]
				img.Pix[dstOffset+3] = s.buffer[srcOffset+3]
			}
		}

		if err := png.Encode(f, img); err != nil {
			log.Warningf("saving screenshot failed: %v", err)
			return
		}

		s.ui.Message("Screenshot saved", 2)
		fmt.Printf("Screenshot saved to %s\n", f.Name())

		// Semi-hack to dump RAM and debug Marioland. In time it should be made
		// into a "Game Genie" kind of feature. For now, this will do.
		if logger.Level >= logger.Debug {
			// TODO: I'd like to be able to call some cpu.DumpRAM() here to make
			//       sure I'm getting the exact RAM state for the current frame
			//       but scope is being an issue.
			//       Maybe through some debug.DumpRAM() where debug would hold
			//       the necessary references. Meh.
		}
	}

	if s.newPalette != nil {
		s.palette = s.newPalette
		s.newPalette = nil
		s.makeBlank() // Recreate blank screen texture buffer with new colors.
	}
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

// Palette will set a new palette for the display and GIF.
func (s *Screen) Palette(p []color.RGBA) {
	// Wait until next frame to apply new palette.
	s.newPalette = p
}
