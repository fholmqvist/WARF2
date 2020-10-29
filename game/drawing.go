package game

import (
	"fmt"
	h "projects/games/warf2/helpers"
	m "projects/games/warf2/worldmap"

	"github.com/beefsack/go-astar"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Draw loop for Game.
func (g *Game) Draw(screen *ebiten.Image) {
	drawMap(g, screen)

	g.ui.Draw(screen, g.gameFont, *g.Data)

	t := g.testChar

	drawPathForTestCharacter(g, screen)
	h.DrawGraphic(t.Entity.Idx, t.Entity.Sprite, screen, g.tilesDwarves, 1)

	drawTPS(g, screen)
}

func drawMap(g *Game, screen *ebiten.Image) {
	for idx, tile := range g.WorldMap.Tiles {
		h.DrawGraphic(idx, tile.Sprite, screen, g.tilesWorld, 1)
	}
}

func drawTPS(g *Game, screen *ebiten.Image) {
	if g.debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	}
}

func drawPathForTestCharacter(g *Game, screen *ebiten.Image) {
	from, ok := g.WorldMap.GetTileByIndex(g.testChar.Entity.Idx)
	if !ok {
		return
	}

	to, ok := g.WorldMap.GetTileByIndex(g.mousePos)
	if !ok {
		return
	}

	path, _, ok := astar.Path(from, to)
	if !ok {
		return
	}

	for _, t := range path {
		tile := t.(*m.Tile)
		h.DrawGraphic(tile.Idx, m.WallSelectedSolid, screen, g.tilesWorld, 0.85)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.testChar.Walker.InitiateWalk(path)
	}
}
