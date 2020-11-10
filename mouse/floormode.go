package mouse

import m "projects/games/warf2/worldmap"

func floorTileMode(mp *m.Map, currentMousePos int) {
	clickFunctions(mp, currentMousePos,
		func() {
			tile, ok := mp.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}

			setHasClicked(currentMousePos)

			if m.IsAnyWall(tile.Sprite) {
				return
			}

			if m.IsFloorBrick(tile.Sprite) {
				return
			}

			tile.Sprite = m.RandomFloorBrick()
		},
		func(mp *m.Map, x int, y int) {
			removeOldSelectionTiles(mp, x, y)
			floorTileSelection(mp, x, y)
		},
	)
}

func floorTileSelection(mp *m.Map, x, y int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}

	if m.IsAnyWall(tile.Sprite) {
		return
	}

	selectionTile, ok := mp.GetSelectionTile(x, y)
	if !ok {
		return
	}

	if m.IsFloorBrick(selectionTile.Sprite) {
		return
	}

	selectionTile.Sprite = m.FloorBricksOne
}

func floorTiles(mp *m.Map, x, y int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}

	if m.IsAnyWall(tile.Sprite) {
		return
	}

	if m.IsFloorBrick(tile.Sprite) {
		return
	}

	tile.Sprite = m.RandomFloorBrick()
}

func resetFloorMode(mp *m.Map, currentMousePos int) {
	clickFunctions(mp, currentMousePos,
		func() {
			if !hasClicked {
				tile, ok := mp.GetTileByIndex(currentMousePos)
				if !ok {
					return
				}

				setHasClicked(currentMousePos)

				if m.IsAnyWall(tile.Sprite) {
					return
				}

				tile.Sprite = m.Ground
			}
		},
		resetFloorSelection,
	)
}

func resetFloorSelection(mp *m.Map, x int, y int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}

	if m.IsAnyWall(tile.Sprite) {
		return
	}

	tile.Sprite = m.Ground
}
