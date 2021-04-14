package room

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/worldmap"
)

// Library room relaxes
// dwarves and increases
// their knowledge.
type Library struct {
	bookShelfAmount int
	chairAmount     int

	tiles []worldmap.Tile
	items []worldmap.Tile
}

func NewLibrary(m *worldmap.Map, x1, y1, x2, y2 int) Library {
	l := Library{}
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			m.SetFloorTile(x, y)
		}
	}

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
