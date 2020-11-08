package mouse

import (
	"projects/games/warf2/item"
	m "projects/games/warf2/worldmap"
)

func placeItem(mp *m.Map, currentMousePos int) {
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
}

func placeItemMode(mp *m.Map, currentMousePos int) {
	clickFunctions(mp, currentMousePos,
		func() {
			placeItem(mp, currentMousePos)
		},
		// This is a bit unfortunate, but it works.
		[]func(*m.Map, int, int){
			func(*m.Map, int, int) { placeItem(mp, currentMousePos) },
		})
}
