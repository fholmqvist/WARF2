package mouse

import m "projects/games/warf2/worldmap"

func removeItemMode(mp *m.Map, currentMousePos int) {
	firstClick(mp, currentMousePos,
		func() {
			item, ok := mp.GetItemTileByIndex(currentMousePos)
			if !ok {
				return
			}

			item.Sprite = m.NoItem
		},
		[]func(*m.Map, int, int){})
}
