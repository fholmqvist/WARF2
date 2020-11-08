package mouse

import (
	"projects/games/warf2/item"
	m "projects/games/warf2/worldmap"
)

func placeItemMode(mp *m.Map, currentMousePos int) {
	firstClick(mp, currentMousePos,
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

			iTile.Sprite = item.RandomBookshelf()
			if item.IsBlocking(iTile.Sprite) {
				tile.Blocked = true
			}
		},
		[]func(*m.Map, int, int){})
}
