package room

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/item"
	m "projects/games/warf2/worldmap"
	"sort"
)

// Library room relaxes
// dwarves and increases
// their knowledge.
type Library struct {
	tiles m.Tiles
}

func NewLibrary(mp *m.Map, x, y int) *Library {
	l := Library{}
	tiles := mp.FloodFillRoom(x, y, m.RandomFloorBrick)
	if len(tiles) == 0 {
		return nil
	}
	sort.Sort(tiles)
	l.tiles = tiles
	var (
		firstRow     = tiles[0].Y
		lastShelfRow = -1
	)
	for _, t := range l.tiles {
		lastShelfRow = l.placeItems(mp, t, firstRow, lastShelfRow)
	}
	return &l
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

func (l *Library) placeItems(mp *m.Map, t m.Tile, firstRow int, lastShelfRow int) int {
	if t.Y == firstRow {
		if !m.IsAnyWall(mp.OneTileUp(t.Idx).Sprite) {
			return lastShelfRow
		}
		item.PlaceRandom(mp, t.X, t.Y, item.RandomBookshelf)
		return lastShelfRow
	}
	if (firstRow-t.Y)%4 == 0 {
		l.generateBookshelves(mp, t)
		return t.Y
	}
	if (firstRow-t.Y)%2 == 0 {
		l.generateFurniture(mp, t)
	}
	if t.Y == lastShelfRow+1 {
		l.breakupBookshelves(mp, lastShelfRow)
		return -1
	}
	return lastShelfRow
}

func (l *Library) generateBookshelves(mp *m.Map, t m.Tile) {
	earlyExists := []bool{
		m.IsDoorOpening(mp, m.OneTileUp(t.Idx)),
		m.IsDoorOpening(mp, m.OneTileDown(t.Idx)),
		m.IsDoorOpening(mp, m.OneTileLeft(t.Idx)),
		m.IsDoorOpening(mp, m.OneTileRight(t.Idx)),
		m.IsAnyWall(mp.OneTileDown(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileDownLeft(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileDownRight(t.Idx).Sprite),
	}
	for _, ee := range earlyExists {
		if ee {
			return
		}
	}
	item.PlaceRandom(mp, t.X, t.Y, item.RandomBookshelf)
}

// In case where bookshelves run
// through an entire room unbroken.
func (l *Library) breakupBookshelves(mp *m.Map, y int) {
	items := []m.Tile{}
	for _, t := range l.tiles {
		if t.Y == y {
			items = append(items, mp.Items[t.Idx])
		}
	}
	if len(items) == 0 {
		return
	}
	shelves := 0
	for i, it := range items {
		if i == 0 || i == len(items)-1 {
			continue
		}
		if item.IsBookShelf(it.Sprite) {
			shelves++
		}
	}
	spaceEvery := 5
	if shelves < spaceEvery {
		return
	}
	for i := 0; i < shelves/spaceEvery; i++ {
		mp.Items[items[spaceEvery*i].Idx].Sprite = item.NoItem
	}
}

func (l *Library) generateFurniture(mp *m.Map, t m.Tile) {
	earlyExists := []bool{
		m.IsAnyWall(mp.Tiles[m.OneTileLeft(t.Idx)].Sprite),
		mp.Items[m.OneTileLeft(t.Idx)].Sprite != item.NoItem,
		m.IsDoorOpening(mp, m.OneTileDown(t.Idx)),
	}
	for _, ee := range earlyExists {
		if ee {
			return
		}
	}
	for x := t.X; x < t.X+3; x++ {
		spr := mp.Tiles[m.XYToIdx(x, t.Y)].Sprite
		if m.IsAnyWall(spr) {
			return
		}
		if !m.IsFloorBrick(spr) {
			return
		}
		if m.IsDoorOpening(mp, m.OneTileDown(m.XYToIdx(x, t.Y))) {
			return
		}
	}
	item.Place(mp, t.X, t.Y, item.ChairLeft)
	item.Place(mp, t.X+1, t.Y, item.Table)
	item.Place(mp, t.X+2, t.Y, item.ChairRight)
}
