package game

import (
	"fmt"
	"projects/games/warf2/globals"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.state {
	case MainMenu:
		menuState := g.ui.UpdateMainMenu()
		switch menuState {
		case -1:
			return nil
		case 0:
			go func() {
				// To prevent from mouseclick
				// carrying over to game.
				time.Sleep(time.Millisecond * 100)
				g.state = Gameplay
			}()
		case 1:
			panic("help not implemented")
		case 2:
			panic("this is not a graceful exit, but it sorta works?")
		default:
			panic(fmt.Sprintf("%d is not a valid return", menuState))
		}
	case Gameplay:
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
		g.mouseSystem.Handle(&g.WorldMap, &g.Rooms, &g.JobService.Workers)
		HandleKeyboard(g)
		// UI.
		g.ui.UpdateMainMenu()
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
	}
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
	g.checkForReading()
}
