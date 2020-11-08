package room

import (
	"projects/games/warf2/character"
	"projects/games/warf2/item"
	"projects/games/warf2/worldmap"
)

// Library room relaxes
// dwarves and increases
// their knowledge.
type Library struct {
	tileAmount      int
	bookShelfAmount int
	chairAmount     int

	tiles []worldmap.Tile
	items []worldmap.Tile
}

// NewLibrary takes a slice of
// tiles and items and returns
// a new library.
func NewLibrary(tiles []worldmap.Tile, items []worldmap.Tile) *Library {
	l := Library{}

	l.tileAmount = len(tiles)

	for _, i := range items {
		if item.IsBookShelf(i.Sprite) {
			l.bookShelfAmount++
		}

		if item.IsChair(i.Sprite) {
			l.chairAmount++
		}
	}

	return &l
}

// Use library.
func (l *Library) Use(dwarf *character.Dwarf) {
	// Every amount of time add
	// more knowledge to Dwarf
	// based on the amount of
	// bookshelves in the library.
	// Decrease stress by same amount.

	// If chair has an adjacent table,
	// stress decreases even faster.
}
