package game

import (
	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	g.time.Tick()

	handleMouse(g)
	handleKeyboard(g)

	g.testChar.Walk(&g.Gmap)

	return nil
}
