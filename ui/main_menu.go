package ui

import (
	"image/color"
	"log"

	"github.com/Holmqvist1990/WARF2/globals"

	"github.com/hajimehoshi/ebiten"
	e "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	inp "github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
)

var (
	width       = 20
	height      = 1
	xOffset     = (globals.ScreenWidth/globals.TileSize)/2 - width/2
	yOffset     = 8
	ySeparation = 4
)

type MainMenu struct {
	idx            int
	keySensitivity int
	logo           *e.Image
	buttons        []*Button
}

func NewMainMenu() *MainMenu {
	logo, _, err := ebitenutil.NewImageFromFile("art/logo.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	return &MainMenu{
		logo: logo,
		buttons: []*Button{
			{Element{"Start", xOffset, yOffset + ySeparation, width, height, textColor, color.Gray{100}}, true},
			{Element{"Help", xOffset, yOffset + ySeparation*2, width, height, textColor, color.Gray{100}}, false},
			{Element{"Quit", xOffset, yOffset + ySeparation*3, width, height, textColor, color.Gray{100}}, false},
		},
	}
}

func (m *MainMenu) Draw(screen *ebiten.Image, uiTiles *ebiten.Image, font font.Face) {
	for _, b := range m.buttons {
		b.Draw(screen, uiTiles, font)
	}
	drawLogo(m, screen)
}

func (m *MainMenu) SelectHoveredButton() {
	m.idx %= len(m.buttons)
	m.keySensitivity = 0
	for i, b := range m.buttons {
		if i == m.idx {
			b.hovering = true
			continue
		}
		b.hovering = false
	}
}

func (m *MainMenu) UpdateMainMenu() int {
	if v, didSelect := m.mouseAndSelect(); didSelect {
		return v
	}
	m.keySensitivity++
	if m.keySensitivity < 4 {
		return -1
	}
	if inp.IsKeyJustPressed(e.KeyUp) || inp.IsKeyJustPressed(e.KeyW) {
		m.idx--
		if m.idx < 0 {
			m.idx = len(m.buttons) - 1
		}
		m.SelectHoveredButton()
		return -1
	}
	if inp.IsKeyJustPressed(e.KeyDown) || inp.IsKeyJustPressed(e.KeyS) {
		m.idx++
		m.SelectHoveredButton()
		return -1
	}
	return -1
}

func (m *MainMenu) mouseAndSelect() (int, bool) {
	for i, b := range m.buttons {
		x, y := ebiten.CursorPosition()
		x /= globals.TileSize
		y /= globals.TileSize
		if !b.MouseIsOver(x, y) {
			continue
		}
		m.idx = i
		m.SelectHoveredButton()
		if inp.IsMouseButtonJustPressed(e.MouseButtonLeft) {
			return i, true
		}
	}
	if inp.IsKeyJustPressed(e.KeyEnter) {
		return m.idx, true
	}
	return -1, false
}

func drawLogo(m *MainMenu, screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(2, 2)
	opt.GeoM.Translate(float64(-m.logo.Bounds().Dx()), 28)
	opt.GeoM.Translate(globals.ScreenWidth/2, 0)
	_ = screen.DrawImage(m.logo, opt)
}
