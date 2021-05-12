package ui

import (
	"image/color"
	"log"
	"projects/games/warf2/worldmap"
	w "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
	e "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	inp "github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
)

var (
	width   = 240
	height  = 80
	xOffset = (w.ScreenWidth / 2) - width/2
	yOffset = 100
	buttons = []*Button{
		{Element{"Start", xOffset, yOffset, width, height, color.Gray{100}}},
		{Element{"Help", xOffset, yOffset * 2, width, height, color.Gray{100}}},
		{Element{"Quit", xOffset, yOffset * 3, width, height, color.Gray{100}}},
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
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(-m.logo.Bounds().Dx()/2), 28)
	opt.GeoM.Translate(worldmap.ScreenWidth/2, 0)
	_ = screen.DrawImage(m.logo, opt)
	for _, b := range buttons {
		b.Draw(screen, font)
	}
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

func (m *MainMenu) Update() int {
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
