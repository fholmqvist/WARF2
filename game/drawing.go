package game

import (
	"fmt"
	"projects/games/warf2/characters"
	h "projects/games/warf2/helpers"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Draw loop for Game.
func (g *Game) Draw(screen *ebiten.Image) {
	drawMap(g, screen)

	g.ui.Draw(screen, g.gameFont, *g.Data)

	drawWorkers(g, screen)

	drawTPS(g, screen)
}

func drawMap(g *Game, screen *ebiten.Image) {
	for idx, tile := range g.WorldMap.Tiles {
		h.DrawGraphic(idx, tile.Sprite, screen, g.worldTiles, 1)
	}
}

func drawTPS(g *Game, screen *ebiten.Image) {
	if g.debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	}
}

func drawWorkers(g *Game, screen *ebiten.Image) {
	for _, worker := range g.JobSystem.Workers {
		ch := worker.(*characters.Character)
		h.DrawGraphic(ch.Entity.Idx, ch.Entity.Sprite, screen, g.dwarfTiles, 1)
	}
}
