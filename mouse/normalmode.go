package mouse

import (
	"fmt"
	"projects/games/warf2/globals"
	m "projects/games/warf2/worldmap"
)

func noneMode(mp *m.Map, currentMousePos int) {
	mp.ClearSelectedTiles()
	clickFunctions(mp, currentMousePos,
		func() {
			printMousePos(currentMousePos)
			// Get tile from real tiles.
			tile, ok := mp.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}
			firstClickedSprite = tile.Sprite
			// Replace that tile with one from SelectedTiles.
			tile, ok = mp.GetSelectionTileByIndex(currentMousePos)
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
		},
		func(mp *m.Map, x int, y int) {
			selectionWalls(mp, x, y)
		})
}

// Attempting to collapse these three similar
// functions into one just made the interface
// that much more complicated. Sometimes, not
// having DRY everywhere ain't that bad.
func mouseUpSetWalls(mp *m.Map, x, y int) {
	selectionTile, ok := mp.GetSelectionTile(x, y)
	if !ok {
		return
	}
	// No change.
	if m.IsNone(selectionTile.Sprite) {
		return
	}
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}
	setWalls(tile)
	selectionTile.Sprite = m.None
}

func selectionWalls(mp *m.Map, x, y int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}
	if !m.IsWallOrSelected(tile.Sprite) {
		return
	}
	selectionTile, ok := mp.GetSelectionTile(x, y)
	if !ok {
		return
	}
	// In order to invert between (un)selected.
	selectionTile.Sprite = tile.Sprite
	setWalls(selectionTile)
}

func setWalls(tile *m.Tile) {
	if !m.IsWallOrSelected(tile.Sprite) {
		return
	}
	if m.IsWall(firstClickedSprite) {
		setToSelected(tile)
		return
	}
	setToNormalInteractFalse(tile)
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
	if tile.Sprite == m.WallSelectedSolid {
		tile.Sprite = m.WallSolid
		return
	}
	tile.Sprite = m.WallExposed
}

func printMousePos(idx int) {
	x, y := globals.IdxToXY(idx)
	fmt.Printf("IDX: %d. XY: {%d, %d}.\n", idx, x, y)
}
