package widgets

import (
	"github.com/veandco/go-sdl2/sdl"
)

// MenuChoice groups a choice string with the corresponding action.
type MenuChoice struct {
	Text   string
	Action func()
}

// Menu widget displaying a list of potential choices, each of which should map
// to some kind of Action.
type Menu struct {
	*widget

	choices  []MenuChoice
	selected int // Index of selected choice
}

func NewMenu(r *sdl.Renderer, s *sdl.Rect, choices []MenuChoice) *Menu {
	return &Menu{
		widget:   new(r, s),
		choices:  choices,
		selected: 0,
	}
}

func (m *Menu) ProcessEvent(e *sdl.Event) {
	// TODO
}

func (m *Menu) Texture() *sdl.Texture {
	m.repaint()
	return m.texture
}

// repaint refreshes the widget's texture with its current state (menu choice).
func (m *Menu) repaint() {
	if len(m.choices) == 0 {
		return
	}
	margin := int32(properties.TitleFont.Height())

	// Render choices first, grab height, render menu.
	var choiceTextures []*sdl.Texture
	for _, c := range m.choices {
		choiceTextures = append(choiceTextures, m.renderText(c.Text))
		//fmt.Println(c.Text)
	}

	// I really hope renderText creates textures that are all the same height.
	_, _, _, choiceH, _ := choiceTextures[0].Query()
	menuH := int32(len(choiceTextures))*choiceH + int32(len(choiceTextures)-1)*margin

	menuTexture, _ := m.renderer.CreateTexture( // TODO: I should probably make a helper method
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		m.width,
		int32(menuH),
	)

	m.renderer.SetRenderTarget(menuTexture)
	y := int32(0) // Start at the top of the texture
	for i, ct := range choiceTextures {
		if i > 0 {
			y += margin
		}

		if i == m.selected {
			m.renderer.SetDrawColor(
				properties.BgColor.R,
				properties.BgColor.G,
				properties.BgColor.B,
				properties.BgColor.A,
			)
			m.renderer.FillRect(&sdl.Rect{
				X: 0, //margin / 2,
				Y: y,
				W: m.width, // - margin,
				H: choiceH,
			})
		}

		_, _, choiceW, choiceH, _ := ct.Query()
		ct.SetBlendMode(sdl.BLENDMODE_BLEND)
		m.renderer.Copy(ct, nil, &sdl.Rect{
			X: (m.width - choiceW) / 2, // Center text, this should probably be in widgets.Label too.
			Y: y,
			W: choiceW,
			H: choiceH,
		})

		y += choiceH
	}

	// Center menu on widget texture.
	m.renderer.SetRenderTarget(m.texture)
	m.renderer.Copy(menuTexture, nil, &sdl.Rect{
		X: 0,
		Y: (m.height - menuH) / 2,
		W: m.width,
		H: menuH,
	})
}
