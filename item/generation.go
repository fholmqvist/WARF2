package item

import (
	"math/rand"
	m "projects/games/warf2/worldmap"
)

var (
	bookShelves = []int{
		BookShelfOne, BookShelfTwo, BookShelfThree,
		BookShelfFour, BookShelfFive, BookShelfSix,
		BookShelfSeven, BookShelfEight, BookShelfNine,
		BookShelfTen,
	}

	furniture = []int{
		ChairLeft, Table, ChairRight,
	}
)

func Place(mp *m.Map, x, y, sprite int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}
	if m.IsAnyWall(tile.Sprite) {
		return
	}
	item, ok := mp.GetItemTile(x, y)
	if !ok {
		return
	}
	item.Sprite = sprite
	if IsBlocking(item.Sprite) {
		tile.Blocked = true
	}
}

func PlaceRandomIdx(mp *m.Map, idx int, f func() int) {
	tile, ok := mp.GetTileByIndex(idx)
	if !ok {
		return
	}
	if m.IsAnyWall(tile.Sprite) {
		return
	}
	item, ok := mp.GetItemTileByIndex(idx)
	if !ok {
		return
	}
	item.Sprite = f()
	if IsBlocking(item.Sprite) {
		tile.Blocked = true
	}
}

// RandomBookshelf returns
// the sprite for a random BookShelf.
func RandomBookshelf() int {
	return bookShelves[rand.Intn(len(bookShelves))]
}

// RandomFurniture returns
// the sprite for a random Furniture.
func RandomFurniture() int {
	return furniture[rand.Intn(len(furniture))]
}
