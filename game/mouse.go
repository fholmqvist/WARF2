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

// This cluster of variables
// help with (de)selecting walls.
var startPoint = -1
var endPoint = -1
var hasClicked = false
var firstClickedSprite = -1

// Remembering last frame
// in order to reset selected
// tiles without having to
// redraw the entire screen.
var previousStartPoint = -1
var previousEndPoint = -1

func handleMouse(g *Game) {
	mouseHover(g)

	idx := mousePos()

	if idx < 0 || idx > m.TilesT {
		return
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseClick(g, idx)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		endPoint = idx
		mouseUp(g)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.mouseMode = None
	}
}

func mouseClick(g *Game, currentMousePos int) {
	switch g.mouseMode {

	case None:
		// Identity
		if !hasClicked {
			// Get tile from real tiles
			tile, ok := g.WorldMap.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}

			firstClickedSprite = tile.Sprite

			// Replace that tile with one from SelectedTiles
			tile, ok = g.WorldMap.GetSelectionTileByIndex(currentMousePos)
			if !ok {
				return
			}

			// Selecting a non-wall defaults to
			// wall in order to enable wall selection
			// without having first clicked on a wall.
			if !m.IsSelectedWall(firstClickedSprite) {
				firstClickedSprite = m.WallSolid
			}

			if m.IsWallOrSelected(tile.Sprite) {
				tile.Sprite = invertSelected(firstClickedSprite)
			}

			startPoint = currentMousePos
			hasClicked = true

		}

		if startPoint >= 0 {
			removeOldSelectionTiles(g)
			selectionWalls(g, startPoint, currentMousePos)
		}

	default:
		fmt.Println("mouseClick: unknown MouseMode:", g.mouseMode)
	}
}

func mouseUp(g *Game) {
	if startPoint >= 0 {
		mouseUpSetWalls(g, startPoint, endPoint)
	}

	startPoint = -1
	hasClicked = false
}

func mouseHover(g *Game) {
	switch g.mouseMode {
	default:
	}
}

func mousePos() int {
	mx, my := ebiten.CursorPosition()
	mx, my = mx/m.TileSize, my/m.TileSize

	return mx + (my * m.TilesW)
}
