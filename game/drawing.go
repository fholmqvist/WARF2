package game

import (
	"fmt"
	m "projects/games/warf2/gmap"
	h "projects/games/warf2/helpers"

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
	for idx, tile := range g.Gmap.Tiles {
		h.DrawGraphic(idx, tile.Sprite, screen, g.tilesWorld, 1)
	}
}

func drawTPS(g *Game, screen *ebiten.Image) {
	if g.debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	}
}

func drawPathForTestCharacter(g *Game, screen *ebiten.Image) {
	idx := g.mousePos
	cIdx := g.testChar.Entity.Idx
	cX, cY := m.IdxToX(cIdx), m.IdxToY(cIdx)
	x, y := m.IdxToX(idx), m.IdxToY(idx)

	path, _, ok := astar.Path(g.Gmap.GetTile(cX, cY), g.Gmap.GetTile(x, y))

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
