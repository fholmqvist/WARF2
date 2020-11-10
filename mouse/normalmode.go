package mouse

import (
	m "projects/games/warf2/worldmap"
)

func noneMode(mp *m.Map, currentMousePos int) {
	clickFunctions(mp, currentMousePos,
		func() {
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
			removeOldSelectionTiles(mp, x, y)
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

	// No change
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

func removeOldSelectionTiles(mp *m.Map, x, y int) {
	selectionTile, ok := mp.GetSelectionTile(x, y)
	if !ok {
		return
	}

	selectionTile.Sprite = m.None
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

	tile.NeedsInteraction = true

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
