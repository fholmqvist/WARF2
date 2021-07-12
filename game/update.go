package game

import (
	"fmt"
	"time"

	"github.com/Holmqvist1990/WARF2/globals"

	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.state {
	case MainMenu:
		g.updateMainMenu()
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
		g.handleInput()
		// Only run if game is not paused.
		if !g.time.Tick() {
			return nil
		}
		if !g.time.TimeToMove() {
			return nil
		}
		g.UpdateDwarves()
		g.JobService.Update(&g.Rooms, &g.WorldMap)
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
}

func (g *Game) handleInput() {
	if mode, changed := g.ui.UpdateGameplayMenus(); changed {
		g.SetMouseMode(mode)
		return
	}
	g.mouseSystem.Handle(&g.WorldMap, &g.Rooms, &g.JobService.Workers)
	HandleKeyboard(g)
}

func (g *Game) updateMainMenu() {
	menuState := g.ui.UpdateMainMenu()
	switch menuState {
	case -1:
		return
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
}
