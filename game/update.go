package game

import (
	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	g.time.Tick()

	g.mouseSystem.Handle(&g.WorldMap, &g.Rooms)
	handleKeyboard(g)

	g.updateDwarves()

	g.JobSystem.Update()

	return nil
}

func (g *Game) updateDwarves() {
	for _, dwarf := range g.JobSystem.Workers {
		dwarf.Walk(&g.WorldMap)
	}
	if !g.time.NewCycle() {
		return
	}
	for _, dwarf := range g.JobSystem.Workers {
		dwarf.Needs.Update(dwarf.Characteristics)
	}
	/////////////////////////////////////////////////
	// TODO
	//
	// Should be redesigned so that we get
	// the highest desired need for each
	// available dwarf, and assign a job to
	// satisfy that need.
	/////////////////////////////////////////////////
	g.checkForLibraryReading()
}
