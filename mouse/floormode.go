package mouse

import m "projects/games/warf2/worldmap"

func floorTileMode(mp *m.Map, currentMousePos int) {
	firstClick(mp, currentMousePos,
		func() {
			tile, ok := mp.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}

			setHasClicked(currentMousePos)

			if m.IsAnyWall(tile.Sprite) {
				return
			}

			if m.IsFloorTile(tile.Sprite) {
				return
			}

			tile.Sprite = m.RandomFloorBrick()
		},
		[]func(*m.Map, int, int){
			floorTileSelection,
		})
}

func floorTileSelection(mp *m.Map, x, y int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}

	if m.IsAnyWall(tile.Sprite) {
		return
	}

	if m.IsFloorTile(tile.Sprite) {
		return
	}

	tile.Sprite = m.RandomFloorBrick()
}

func resetFloorMode(mp *m.Map, currentMousePos int) {
	firstClick(mp, currentMousePos,
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
		[]func(*m.Map, int, int){
			resetFloorSelection,
		})
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
