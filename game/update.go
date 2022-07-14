package game

import (
	"fmt"
	"os"

	gl "github.com/Holmqvist1990/WARF2/globals"

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
		if g.debugFunc != nil && gl.DEBUG {
			f := *g.debugFunc
			f(g)
		}
		g.ui.MouseOverInformation.Text = g.handleInput()
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
		dwarf.Needs.Update(dwarf.Attributes)
	}
}

func (g *Game) handleInput() string {
	HandleKeyboard(g)
	if gl.GAME_PAUSED {
		return ""
	}
	if mode, changed := g.ui.UpdateGameplayMenus(); changed {
		g.SetMouseMode(mode)
		return ""
	}
	overviewText := g.mouseSystem.Handle(g.WorldMap, g.Rooms, &g.JobService.Workers)
	return overviewText
}

func (g *Game) updateMainMenu() {
	menuState := g.ui.UpdateMainMenu()
	switch menuState {
	case -1:
		return
	case 0:
		gl.Delay(func() {
			g.state = Gameplay
		})
	case 1:
		g.state = HelpMenu
		gl.Delay(func() {
			g.ui.HelpMenu.Clickable = true
		})
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
