package room

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/item"
	"projects/games/warf2/worldmap"
)

// Library room relaxes
// dwarves and increases
// their knowledge.
type Library struct {
	// bookShelfAmount int
	// chairAmount     int

	// tiles []worldmap.Tile
	// items []worldmap.Tile
}

func NewLibrary(m *worldmap.Map, x1, y1, x2, y2 int) Library {
	l := Library{}
	m.SetFloorTiles(x1, y1, x2, y2)
	l.generateBookShelves(m, x1, y1, x2, y2)
	l.generateFurniture(m, x1, y1, x2, y2)
	return l
}

// Use library.
func (l *Library) Use(dwarf *dwarf.Dwarf) {
	// Every amount of time add
	// more knowledge to Dwarf
	// based on the amount of
	// bookshelves in the library.
	// Decrease stress by same amount.

	// If chair has an adjacent table,
	// stress decreases even faster.
}

func (l *Library) generateBookShelves(m *worldmap.Map, x1, y1, x2, y2 int) {
	for y := y1; y < y2; y += 4 {
		for x := x1; x < x2; x++ {
			idx := worldmap.XYToIdx(x, y)
			item.PlaceRandomIdx(m, idx, item.RandomBookshelf)
		}
	}
}

func (l *Library) generateFurniture(m *worldmap.Map, x1, y1, x2, y2 int) {
	for y := y1 + 2; y < y2; y += 4 {
		for x := x1; x < x2-1; x += 3 {
			item.Place(m, x, y, item.ChairLeft)
			item.Place(m, x+1, y, item.Table)
			if x+2 < x2 {
				item.Place(m, x+2, y, item.ChairRight)
			}
		}
	}
}

// func middle(x1, x2 int) int {
// 	return x1 + ((x2 - x1) / 2)
// }
