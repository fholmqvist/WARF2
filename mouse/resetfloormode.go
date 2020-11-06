package mouse

import m "projects/games/warf2/worldmap"

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
