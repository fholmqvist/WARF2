// Package ui handles all the
// in-game graphical overlays.
package ui

import (
	"fmt"
	"image/color"
	"projects/games/warf2/entity"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

// UI wraps all the UI elements for Game
type UI struct {
	MouseMode Element
}

// Element wraps data for UI elements
type Element struct {
	Text  string
	X     int
	Y     int
	Color color.Color
}

// Draw function or UI overlays
func (ui *UI) Draw(screen *ebiten.Image, gameFont font.Face, data entity.Data) {
	mm := ui.MouseMode
	x, y, xo := 20, 20, 10

	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		drawBackground(screen, x, y)

		// Crime
		cs := fmt.Sprintf("Example: %d", data.Example)

		text.Draw(screen, cs, gameFont, x+xo, y*2, mm.Color)
	}

	text.Draw(screen, mm.Text, gameFont, mm.X, mm.Y, mm.Color)
}

func drawBackground(screen *ebiten.Image, x, y int) {
	sw, _ := screen.Size()

	uibg, _ := ebiten.NewImage(sw-40, 32, ebiten.FilterDefault)
	_ = uibg.Fill(color.Gray{50})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	_ = screen.DrawImage(uibg, op)
}
