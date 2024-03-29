package game

import (
	"fmt"

	gl "github.com/Holmqvist1990/WARF2/globals"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Draw loop for Game.
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case MainMenu:
		g.ui.MainMenu.Draw(screen, g.uiTiles, g.font)
	case HelpMenu:
		g.ui.HelpMenu.Draw(screen, g.uiTiles, g.font)
	case Gameplay:
		drawMap(g, screen)
		drawMovables(g, screen)
		drawWorkers(g, screen)
		g.ui.GameplayUI.Draw(screen, g.uiTiles, g.font, &g.ui, g.JobService.Workers)
		drawTPS(g, screen)
	default:
		panic(fmt.Sprintf("unknown gamestate: %v", g.state))
	}
}

func DrawGraphic(idx, sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha float64) {
	gl.DrawTile(sprite, screen, tileset, alpha, gl.DrawOptions(idx, alpha, 0))
}

func DrawRailGraphic(idx, sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha, rotation float64) {
	if sprite == 0 {
		return
	}
	gl.DrawTile(sprite, screen, tileset, alpha, gl.DrawOptions(idx, alpha, rotation))
}

func drawMap(g *Game, screen *ebiten.Image) {
	for idx, tile := range g.WorldMap.Tiles {
		DrawGraphic(idx, tile.Sprite, screen, g.worldTiles, 1)
	}
	for idx, tile := range g.WorldMap.SelectedTiles {
		DrawGraphic(idx, tile.Sprite, screen, g.worldTiles, 1)
	}
	for idx, tile := range g.WorldMap.Rails {
		DrawRailGraphic(idx, tile.Sprite, screen, g.railTiles, 1, tile.Rotation)
	}
	for idx, tile := range g.WorldMap.Items {
		DrawGraphic(idx, tile.Sprite, screen, g.itemTiles, 1)
	}
}

func drawTPS(g *Game, screen *ebiten.Image) {
	if gl.DEBUG {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	}
}

func drawWorkers(g *Game, screen *ebiten.Image) {
	for _, dwarf := range g.JobService.Workers {
		DrawGraphic(dwarf.Idx, dwarf.Sprite, screen, g.dwarfTiles, 1)
	}
}

func drawMovables(g *Game, screen *ebiten.Image) {
	for _, cart := range g.RailService.Carts {
		DrawGraphic(cart.Idx, cart.Sprite, screen, g.railTiles, 1)
	}
}
