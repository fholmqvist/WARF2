package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Draw loop for Game.
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {

	case MainMenu:
		menuState := g.ui.DrawMainMenu(screen, g.font)
		switch menuState {
		case -1:
			return
		case 0:
			g.state = Gameplay
		case 1:
			panic("help not implemented")
		case 2:
			panic("this is not a graceful exit, but it sorta works?")
		default:
			panic(fmt.Sprintf("%d is not a valid return", menuState))
		}

	case Gameplay:
		drawMap(g, screen)
		drawWorkers(g, screen)
		g.ui.DrawGameplay(screen, g.font, g.Dwarves)
		drawTPS(g, screen)

	default:
		panic(fmt.Sprintf("unknown gamestate: %v", g.state))
	}
}

func drawMap(g *Game, screen *ebiten.Image) {
	for idx, tile := range g.WorldMap.Tiles {
		DrawGraphic(idx, tile.Sprite, screen, g.worldTiles, 1)
	}
	for idx, tile := range g.WorldMap.SelectedTiles {
		DrawGraphic(idx, tile.Sprite, screen, g.worldTiles, 1)
	}
	for idx, tile := range g.WorldMap.Items {
		DrawGraphic(idx, tile.Sprite, screen, g.itemTiles, 1)
	}
}

func drawTPS(g *Game, screen *ebiten.Image) {
	if g.debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	}
}

func drawWorkers(g *Game, screen *ebiten.Image) {
	for _, dwarf := range g.JobService.Workers {
		DrawGraphic(dwarf.Idx, dwarf.Sprite, screen, g.dwarfTiles, 1)
	}
}
