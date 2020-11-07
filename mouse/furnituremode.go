package mouse

import m "projects/games/warf2/worldmap"

func placeFurnitureMode(mp *m.Map, currentMousePos int) {
	firstClick(mp, currentMousePos,
		func() {
			tile, ok := mp.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}

			if m.IsAnyWall(tile.Sprite) {
				return
			}

			item, ok := mp.GetItemTileByIndex(currentMousePos)
			if !ok {
				return
			}

			item.Sprite = m.RandomFurniture()
		},
		[]func(*m.Map, int, int){})
}
