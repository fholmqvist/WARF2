package game

import (
	"fmt"
	"os"
	"time"

	"github.com/Holmqvist1990/WARF2/globals"

	"github.com/hajimehoshi/ebiten"
)

// Update loop for Game.
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.state {
	case MainMenu:
		g.updateMainMenu()
	case HelpMenu:
		g.updateHelpMenu()
	case Gameplay:
		if g.debugFunc != nil && globals.DEBUG {
			f := *g.debugFunc
			f(g)
		}
		g.ui.OverviewTab.Text = g.handleInput()
		if !g.time.Tick() {
			return nil
		}
		if !g.time.TimeToMove() {
			return nil
		}
		g.UpdateDwarves()
		g.JobService.Update(g.Rooms, g.WorldMap)
		g.RailService.Update(g.WorldMap)
		if !g.time.QuarterCycle() {
			return nil
		}
		g.Rooms.Update(g.WorldMap)
	}
	return nil
}

func (g *Game) UpdateDwarves() {
	for _, dwarf := range g.JobService.Workers {
		dwarf.Walk(g.WorldMap)
	}
	if !g.time.NewCycle() {
		return
	}
	for _, dwarf := range g.JobService.Workers {
		dwarf.Needs.Update(dwarf.Characteristics)
	}
}

func (g *Game) handleInput() string {
	if mode, changed := g.ui.UpdateGameplayMenus(); changed {
		g.SetMouseMode(mode)
		return ""
	}
	overviewText := g.mouseSystem.Handle(g.WorldMap, g.Rooms, &g.JobService.Workers)
	HandleKeyboard(g)
	return overviewText
}

func (g *Game) updateMainMenu() {
	menuState := g.ui.UpdateMainMenu()
	switch menuState {
	case -1:
		return
	case 0:
		go func() {
			// Discard further input, yield.
			time.Sleep(time.Millisecond * 100)
			g.state = Gameplay
		}()
	case 1:
		g.state = HelpMenu
		go func() {
			// Discard further input, yield.
			time.Sleep(time.Millisecond * 100)
			g.ui.HelpMenu.Clickable = true
		}()
	case 2:
		os.Exit(3)
	default:
		panic(fmt.Sprintf("%d is not a valid return", menuState))
	}
}

func (g *Game) updateHelpMenu() {
	back := g.ui.HelpMenu.Update()
	if back {
		g.state = MainMenu
	}
}
