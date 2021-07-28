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
	width       = 240
	height      = 80
	xOffset     = (globals.ScreenWidth / 2) - width/2
	yOffset     = 48
	ySeparation = 100
	buttons     = []*Button{
		{Element{"Start", xOffset, yOffset + ySeparation, width, height, textColor, color.Gray{100}}},
		{Element{"Help", xOffset, yOffset + ySeparation*2, width, height, textColor, color.Gray{100}}},
		{Element{"Quit", xOffset, yOffset + ySeparation*3, width, height, textColor, color.Gray{100}}},
	}
)

func init() {
	buttons[0].Select()
	buttons[1].Deselect()
	buttons[2].Deselect()
}

type MainMenu struct {
	idx            int
	keySensitivity int
	logo           *e.Image
}

func NewMainMenu() *MainMenu {
	logo, _, err := ebitenutil.NewImageFromFile("art/logo.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	return &MainMenu{
		logo: logo,
	}
}

func (m *MainMenu) Draw(screen *ebiten.Image, font font.Face) {
	for _, b := range buttons {
		b.Draw(screen, font)
	}
	drawLogo(m, screen)
}

func (m *MainMenu) Select() {
	m.idx %= len(buttons)
	m.keySensitivity = 0
	for i, b := range buttons {
		if i == m.idx {
			b.Select()
			continue
		}
		b.Deselect()
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
			m.idx = len(buttons) - 1
		}
		m.Select()
		return -1
	}
	if inp.IsKeyJustPressed(e.KeyDown) || inp.IsKeyJustPressed(e.KeyS) {
		m.idx++
		m.Select()
		return -1
	}
	return -1
}

func (m *MainMenu) mouseAndSelect() (int, bool) {
	for i, b := range buttons {
		x, y := ebiten.CursorPosition()
		if b.MouseIsOver(x, y) {
			m.idx = i
			m.Select()
			if inp.IsMouseButtonJustPressed(e.MouseButtonLeft) {
				return i, true
			}
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
