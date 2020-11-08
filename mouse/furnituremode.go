package mouse

import (
	"projects/games/warf2/item"
	m "projects/games/warf2/worldmap"
)

func placeFurnitureMode(mp *m.Map, currentMousePos int) {
	clickFunctions(mp, currentMousePos,
		func() {
			tile, ok := mp.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}

			if m.IsAnyWall(tile.Sprite) {
				return
			}

			iTile, ok := mp.GetItemTileByIndex(currentMousePos)
			if !ok {
				return
			}

			iTile.Sprite = item.RandomFurniture()
			if item.IsBlocking(iTile.Sprite) {
				tile.Blocked = true
			}
		},
		[]func(*m.Map, int, int){})
}
