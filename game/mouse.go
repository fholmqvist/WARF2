package game

import (
	"fmt"

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
		mouseUp()
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

			if !m.IsWallOrSelected(tile.Sprite) {
				return
			}

			firstClickedSprite = tile.Sprite
			tile.Sprite = invertSelected(firstClickedSprite)

			startPoint = idx
			hasClicked = true

		} else if startPoint >= 0 {
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
					if !m.IsWallOrSelected(tile.Sprite) {
						continue
					}
					if m.IsWall(firstClickedSprite) {
						tile.Sprite = setToSelected(tile.Sprite)
					} else {
						tile.Sprite = setToNormal(tile.Sprite)
					}
				}
			}
		}

	default:
		fmt.Println("mouseClick: unknown MouseMode:", g.mouseMode)
	}
}

func mouseUp() {
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

func setToSelected(sprite int) int {
	if m.IsSelectedWall(sprite) {
		return sprite
	}

	if sprite == m.WallSolid {
		return m.WallSelectedSolid
	}
	return m.WallSelectedExposed
}

func setToNormal(sprite int) int {
	if m.IsWall(sprite) {
		return sprite
	}

	if sprite == m.WallSelectedSolid {
		return m.WallSolid
	}
	return m.WallExposed
}
