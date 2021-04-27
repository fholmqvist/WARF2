package ui

import (
	"fmt"
	"image/color"
	w "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
	e "github.com/hajimehoshi/ebiten"
	i "github.com/hajimehoshi/ebiten/inpututil"
)

var (
	width   = 240
	height  = 80
	xOffset = (w.ScreenWidth / 2) - width/2
	yOffset = 100
	buttons = []Button{
		{Element{"Start", xOffset, yOffset, width, height, color.Gray{100}}, func() {}},
		{Element{"Help", xOffset, yOffset * 2, width, height, color.Gray{100}}, func() {}},
		{Element{"Quit", xOffset, yOffset * 3, width, height, color.Gray{100}}, func() {}},
	}
)

type MainMenu struct {
	idx int

	keySensitivity int
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	for _, b := range buttons {
		b.Draw(screen)
	}
}

func (m *MainMenu) Update() {
	m.keySensitivity++
	if m.keySensitivity < 8 {
		return
	}
	if i.IsKeyJustPressed(e.KeyUp) || i.IsKeyJustPressed(e.KeyW) {
		m.idx--
		m.idx %= len(buttons)
		fmt.Println(m.idx)
		m.keySensitivity = 0
		return
	}
	if i.IsKeyJustPressed(e.KeyDown) || i.IsKeyJustPressed(e.KeyS) {
		m.idx++
		m.idx %= len(buttons)
		fmt.Println(m.idx)
		m.keySensitivity = 0
		return
	}
}
