package game

import (
	"fmt"

	"projects/games/warf2/worldmap"
	m "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// MouseMode enum for managing mouse action state.
type MouseMode int

// MouseMode enum.
const (
	None MouseMode = iota
)

// This cluster of variables
// help with (de)selecting walls.
var startPoint = -1
var hasClicked = false
var firstClickedSprite = -1

func handleMouse(g *Game) {
	mouseHover(g)

	idx := mousePos()

	if idx < 0 || idx > m.TilesT {
		return
	}

	g.mousePos = idx

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if g.debug {
			fmt.Println("tile:", idx, m.GraphicName(g.WorldMap.Tiles[idx].Idx))
		}

		mouseClick(g, idx)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mouseUp(g)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.mouseMode = None
	}
}

func mouseClick(g *Game, idx int) {

	switch g.mouseMode {

	case None:
		// Identity
		if !hasClicked {
			tile, ok := g.WorldMap.GetTileByIndex(idx)
			if !ok {
				return
			}

			firstClickedSprite = tile.Sprite

			// Selecting a non-wall defaults to
			// wall in order to enable wall selection
			// without having first clicked on a wall.
			if !m.IsSelectedWall(firstClickedSprite) {
				firstClickedSprite = m.WallSolid
			}

			if m.IsWallOrSelected(tile.Sprite) {
				tile.Sprite = invertSelected(firstClickedSprite)
			}

			startPoint = idx
			hasClicked = true

		}
		if startPoint >= 0 {
			runJobOverRangeOfTiles(g, idx, setWalls)
		}

	default:
		fmt.Println("mouseClick: unknown MouseMode:", g.mouseMode)
	}
}

func mouseUp(g *Game) {
	if startPoint >= 0 {
		runJobOverRangeOfTiles(g, mousePos(), wallNeedsInteraction)
	}

	startPoint = -1
	hasClicked = false
}

func mouseHover(g *Game) {
	// idx := mousePos()

	switch g.mouseMode {

	default:
	}
}

func mousePos() int {
	mx, my := ebiten.CursorPosition()
	mx, my = mx/m.TileSize, my/m.TileSize

	return mx + (my * m.TilesW)
}

func runJobOverRangeOfTiles(g *Game, idx int, f func(*worldmap.Tile)) {
	x1, y1 := m.IdxToXY(startPoint)
	x2, y2 := m.IdxToXY(idx)

	if x1 > x2 {
		x1, x2 = x2, x1
	}

	if y1 > y2 {
		y1, y2 = y2, y1
	}

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			tile, ok := g.WorldMap.GetTile(x, y)
			if !ok {
				continue
			}
			f(tile)
		}
	}
}

func setWalls(tile *worldmap.Tile) {
	if !m.IsWallOrSelected(tile.Sprite) {
		return
	}
	if m.IsWall(firstClickedSprite) {
		setToSelected(tile)
	} else {
		setToNormalInteractFalse(tile)
	}
}

func wallNeedsInteraction(tile *worldmap.Tile) {
	if !m.IsSelectedWall(tile.Sprite) {
		return
	}
	tile.NeedsInteraction = true
}

func invertSelected(sprite int) int {
	if m.IsWall(sprite) {
		if sprite == m.WallSolid {
			return m.WallSelectedSolid
		}
		return m.WallSelectedExposed
	}

	if sprite == m.WallSelectedSolid {
		return m.WallSolid
	}
	return m.WallExposed
}

func setToSelected(tile *m.Tile) {
	if m.IsSelectedWall(tile.Sprite) {
		return
	}

	if tile.Sprite == m.WallSolid {
		tile.Sprite = m.WallSelectedSolid
		return
	}
	tile.Sprite = m.WallSelectedExposed
}

func setToNormalInteractFalse(tile *m.Tile) {
	if m.IsWall(tile.Sprite) {
		return
	}

	// Resetting here as this state is
	// no longer valid. With that said,
	// this premature assumption will
	// probably bite me in the ass
	// sometime in the future when I
	// rediscover this after hours
	// of (unnecessary?) debugging.
	tile.NeedsInteraction = false

	if tile.Sprite == m.WallSelectedSolid {
		tile.Sprite = m.WallSolid
		return
	}
	tile.Sprite = m.WallExposed
}
