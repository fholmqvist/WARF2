package ui

import (
	"image/color"
	w "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
	e "github.com/hajimehoshi/ebiten"
	i "github.com/hajimehoshi/ebiten/inpututil"
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
}

func (m *MainMenu) Draw(screen *ebiten.Image, font font.Face) {
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
	m.keySensitivity++
	if m.keySensitivity < 4 {
		return -1
	}
	if i.IsKeyJustPressed(e.KeyUp) || i.IsKeyJustPressed(e.KeyW) {
		m.idx--
		if m.idx < 0 {
			m.idx = len(buttons) - 1
		}
		m.Select()
		return -1
	}
	if i.IsKeyJustPressed(e.KeyDown) || i.IsKeyJustPressed(e.KeyS) {
		m.idx++
		m.Select()
		return -1
	}
	if i.IsKeyJustPressed(e.KeyEnter) {
		return m.idx
	}
	return -1
}
