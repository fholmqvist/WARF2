package game

import (
	"projects/games/warf2/character"

	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	g.time.Tick()

	g.mouseSystem.Handle(&g.WorldMap)
	handleKeyboard(g)

	g.updateCharacters()

	g.JobSystem.Update()

	return nil
}

func (g *Game) updateCharacters() {
	for _, worker := range g.JobSystem.Workers {
		dwarf := worker.(*character.Dwarf)
		dwarf.Walk(&g.WorldMap)
	}
}
