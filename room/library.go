package room

import (
	"math/rand"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/item"
	"projects/games/warf2/worldmap"
	m "projects/games/warf2/worldmap"
)

// Library room relaxes
// dwarves and increases
// their knowledge.
type Library struct {
	// bookShelfAmount int
	// chairAmount     int

	tiles []worldmap.Tile
	// items []worldmap.Tile
}

func NewLibrary(mp *m.Map, x, y int) Library {
	l := Library{}
	tiles := mp.FloodFillRoom(x, y, m.RandomFloorBrick)
	bookShelfRows := 0
	for _, t := range tiles {
		bookShelfRows += l.generateBookshelves(mp, t, 4)
		l.generateFurniture(mp, t, 4)
	}
	l.tiles = tiles
	for i := 0; i < bookShelfRows; i++ {
		l.cleanupBookshelves(mp, y+(i*4))
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

func (l *Library) generateBookshelves(mp *m.Map, t m.Tile, every int) int {
	earlyExists := []bool{
		t.Y%every != 0,
		m.IsDoorOpening(mp, m.OneTileUp(t.Idx)),
		m.IsDoorOpening(mp, m.OneTileDown(t.Idx)),
		m.IsDoorOpening(mp, m.OneTileLeft(t.Idx)),
		m.IsDoorOpening(mp, m.OneTileRight(t.Idx)),
		m.IsAnyWall(mp.Tiles[m.OneTileDown(t.Idx)].Sprite),
	}
	for _, ee := range earlyExists {
		if ee {
			return 0
		}
	}
	item.PlaceRandom(mp, t.X, t.Y, item.RandomBookshelf)
	return 1
}

// In case where bookshelves run
// through an entire room unbroken.
func (l *Library) cleanupBookshelves(mp *m.Map, y int) {
	items := []m.Tile{}
	for _, t := range l.tiles {
		if t.Y == y {
			items = append(items, mp.Items[t.Idx])
		}
	}
	if len(items) == 0 {
		return
	}
	for _, i := range items {
		if i.Sprite == item.NoItem {
			return
		}
	}
	mp.Items[items[rand.Intn(len(items))].Idx].Sprite = item.NoItem
}

func (l *Library) generateFurniture(mp *m.Map, t m.Tile, every int) {
	newY := t.Y + 2
	if newY%every != 0 {
		return
	}
	if m.IsAnyWall(mp.Tiles[m.OneTileLeft(t.Idx)].Sprite) {
		return
	}
	if mp.Items[m.OneTileLeft(t.Idx)].Sprite != item.NoItem {
		return
	}
	for x := t.X; x < t.X+3; x++ {
		spr := mp.Tiles[m.XYToIdx(x, t.Y)].Sprite
		if m.IsAnyWall(spr) {
			return
		}
		if !m.IsFloorBrick(spr) {
			return
		}
	}
	item.Place(mp, t.X, t.Y, item.ChairLeft)
	item.Place(mp, t.X+1, t.Y, item.Table)
	item.Place(mp, t.X+2, t.Y, item.ChairRight)
}
