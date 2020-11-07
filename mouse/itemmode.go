package mouse

import (
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

			item, ok := mp.GetItemTileByIndex(currentMousePos)
			if !ok {
				return
			}

			item.Sprite = m.RandomBookshelf()
			tile.Blocked = true
		},
		[]func(*m.Map, int, int){})
}

func removeItemMode(mp *m.Map, currentMousePos int) {
	firstClick(mp, currentMousePos,
		func() {
			item, ok := mp.GetItemTileByIndex(currentMousePos)
			if !ok {
				return
			}

			item.Sprite = m.NoItem

			tile, ok := mp.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}

			tile.Blocked = false
		},
		[]func(*m.Map, int, int){})
}
