package game

import (
	"projects/games/warf2/characters"

	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	g.time.Tick()

	handleMouse(g)
	handleKeyboard(g)

	g.updateCharacters()

	g.JobSystem.Update()

	return nil
}

func (g *Game) updateCharacters() {
	for _, worker := range g.JobSystem.Workers {
		ch := worker.(*characters.Character)
		ch.Walk(&g.WorldMap)
	}
}
