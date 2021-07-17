package item

import (
	"math/rand"

	gl "github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
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
	if gl.IsBlocking(item.Sprite) {
		tile.Blocked = true
	}
}

func PlaceRandom(mp *m.Map, x, y int, f func() int) {
	idx := gl.XYToIdx(x, y)
	PlaceRandomIdx(mp, idx, f)
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
	if gl.IsBlocking(item.Sprite) {
		tile.Blocked = true
	}
}

func RandomBookshelf() int {
	return gl.BookShelves[rand.Intn(len(gl.BookShelves))]
}

func RandomFurniture() int {
	return gl.Furniture[rand.Intn(len(gl.Furniture))]
}

func RandomCrumbledWall() int {
	return rand.Intn(gl.WallCrumbled4-gl.WallCrumbled1+1) + gl.WallCrumbled1
}
