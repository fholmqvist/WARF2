package mouse

import m "projects/games/warf2/worldmap"

func floorTileMode(mp *m.Map, currentMousePos int) {
	if !hasClicked {
		// Get tile from real tiles
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
	}

	if startPoint >= 0 {
		floorTileSelection(mp, startPoint, currentMousePos)
	}
}

func floorTileSelection(mp *m.Map, start, end int) {
	x1, y1, x2, y2 := tileRange(start, end)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			tile, ok := mp.GetTile(x, y)
			if !ok {
				continue
			}

			if m.IsAnyWall(tile.Sprite) {
				continue
			}

			if m.IsFloorTile(tile.Sprite) {
				continue
			}

			tile.Sprite = m.RandomFloorBrick()
		}
	}
}
