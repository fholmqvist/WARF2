package mouse

import (
	m "projects/games/warf2/worldmap"
)

func noneMode(mp *m.Map, currentMousePos int) {
	if !hasClicked {
		// Get tile from real tiles
		tile, ok := mp.GetTileByIndex(currentMousePos)
		if !ok {
			return
		}

		firstClickedSprite = tile.Sprite

		// Replace that tile with one from SelectedTiles
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

		setHasClicked(currentMousePos)
	}

	if startPoint >= 0 {
		removeOldSelectionTiles(mp)
		selectionWalls(mp, startPoint, currentMousePos)
	}
}

// Attempting to collapse these three similar
// functions into one just made the interface
// that much more complicated. Sometimes, not
// having DRY everywhere ain't that bad.
func mouseUpSetWalls(mp *m.Map, start, end int) {
	x1, y1, x2, y2 := tileRange(start, end)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			selectionTile, ok := mp.GetSelectionTile(x, y)
			if !ok {
				continue
			}

			// No change
			if m.IsNone(selectionTile.Sprite) {
				continue
			}

			tile, ok := mp.GetTile(x, y)
			if !ok {
				continue
			}
			setWalls(tile)
			selectionTile.Sprite = m.None
		}
	}
}

func selectionWalls(mp *m.Map, start, end int) {
	x1, y1, x2, y2 := tileRange(start, end)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			tile, ok := mp.GetTile(x, y)
			if !ok {
				continue
			}

			if !m.IsWallOrSelected(tile.Sprite) {
				continue
			}

			selectionTile, ok := mp.GetSelectionTile(x, y)
			if !ok {
				continue
			}

			// In order to invert between (un)selected
			selectionTile.Sprite = tile.Sprite

			setWalls(selectionTile)
		}
	}

	previousStartPoint = start
	previousEndPoint = end
}

func removeOldSelectionTiles(mp *m.Map) {
	x1, y1, x2, y2 := tileRange(previousStartPoint, previousEndPoint)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			selectionTile, ok := mp.GetSelectionTile(x, y)
			if !ok {
				continue
			}

			selectionTile.Sprite = m.None
		}
	}
}

func tileRange(start, end int) (int, int, int, int) {
	x1, y1 := m.IdxToXY(start)
	x2, y2 := m.IdxToXY(end)

	if x1 > x2 {
		x1, x2 = x2, x1
	}

	if y1 > y2 {
		y1, y2 = y2, y1
	}

	return x1, y1, x2, y2
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
