package mouse

import (
	"projects/games/warf2/item"
	m "projects/games/warf2/worldmap"
)

func placeItemMode(mp *m.Map, currentMousePos int) {
	clickFunctions(mp, currentMousePos,
		func() {
			item.PlaceRandomIdx(mp, currentMousePos, item.RandomBookshelf)
		},
		func(*m.Map, int, int) {
			item.PlaceRandomIdx(mp, currentMousePos, item.RandomBookshelf)
		},
	)
}
