// Package ui handles all the
// in-game graphical overlays.
package ui

import (
	"image/color"
	"projects/games/warf2/dwarf"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

// UI wraps all the UI elements for Game
type UI struct {
	*MainMenu
	MouseMode Element
}

func (ui *UI) DrawMainMenu(screen *ebiten.Image, font font.Face) int {
	ui.MainMenu.Draw(screen, font)
	return ui.MainMenu.Update()
}

func (ui *UI) DrawGameplay(screen *ebiten.Image, gameFont font.Face, dw []dwarf.Dwarf) {
	mm := ui.MouseMode
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		drawOverview(screen, gameFont, dw, mm)
	}
	text.Draw(screen, mm.Text, gameFont, mm.X, mm.Y, mm.Color)
}

func drawOverview(screen *ebiten.Image, gameFont font.Face, dw []dwarf.Dwarf, mm Element) {
	x, y, xo, yo := 20, 20, 10, 20
	drawBackground(screen, x, y, (len(dw)*y)+y+yo)
	text.Draw(screen, "Dwarves:", gameFont, x+xo, y*2, mm.Color)
	for i, d := range dw {
		text.Draw(screen, d.Characteristics.Name, gameFont, (x*2)+xo, yo+(y*(i+2)), mm.Color)
	}
}

func drawBackground(screen *ebiten.Image, x, y, height int) {
	sw, _ := screen.Size()
	e := Element{
		"",
		x,
		y,
		sw - 40,
		height,
		color.Gray{50},
	}
	DrawSquare(screen, e)
}
