package game

import (
	"projects/games/warf2/globals"

	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	// Only update state if
	// player is actively playing.
	if g.state != Gameplay {
		return nil
	}
	if g.debugFunc != nil && globals.DEBUG {
		f := *g.debugFunc
		f(g)
	}
	// Handle input.
	g.mouseSystem.Handle(&g.WorldMap, &g.Rooms)
	HandleKeyboard(g)
	// Only run if game is not paused.
	if !g.time.Tick() {
		return nil
	}
	if !g.time.TimeToMove() {
		return nil
	}
	g.UpdateDwarves()
	g.JobService.Update(&g.Rooms)
	g.RailService.Update(&g.WorldMap)
	return nil
}

func (g *Game) UpdateDwarves() {
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
	// Should be redesigned so that we get
	// the highest desired need for each
	// available dwarf, and assign a job to
	// satisfy that need.
	/////////////////////////////////////////////////
	g.checkForLibraryReading()
}
