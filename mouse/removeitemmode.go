package mouse

import (
	"projects/games/warf2/item"
	m "projects/games/warf2/worldmap"
)

func removeItemMode(mp *m.Map, currentMousePos int) {
	clickFunctions(mp, currentMousePos,
		func() {
			iTile, ok := mp.GetItemTileByIndex(currentMousePos)
			if !ok {
				return
			}

			iTile.Sprite = item.NoItem

			tile, ok := mp.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}

			tile.Blocked = false
		},
		[]func(*m.Map, int, int){})
}
