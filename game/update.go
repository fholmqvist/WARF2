package game

import (
	"fmt"
	"projects/games/warf2/dwarf"

	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	// Only update state if
	// player is actively playing.
	if g.state != Gameplay {
		return nil
	}
	g.time.Tick()
	g.mouseSystem.Handle(&g.WorldMap, &g.Rooms)
	handleKeyboard(g)
	g.updateDwarves()
	g.JobService.Update()

	for _, d := range g.Dwarves {
		if d.State == dwarf.WorkerIdle {
			continue
		}
		fmt.Println(d)
	}

	return nil
}

func (g *Game) updateDwarves() {
	for _, dwarf := range g.JobService.Workers {
		dwarf.Walk(&g.WorldMap)
	}
	if !g.time.NewCycle() {
		return
	}
	for _, dwarf := range g.JobService.Workers {
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
