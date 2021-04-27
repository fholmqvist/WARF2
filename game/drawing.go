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
		g.ui.DrawMainMenu(screen)

	case Gameplay:
		drawMap(g, screen)
		drawWorkers(g, screen)
		g.ui.DrawGameplay(screen, g.gameFont, g.Dwarves)
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
