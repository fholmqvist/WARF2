package item

import (
	"math/rand"

	"github.com/Holmqvist1990/WARF2/entity"
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
}

func RandomBookshelf() int {
	return entity.BookShelves[rand.Intn(len(entity.BookShelves))]
}

func RandomFurniture() int {
	return entity.Furniture[rand.Intn(len(entity.Furniture))]
}

func RandomCrumbledWall() int {
	return rand.Intn(entity.WallCrumbled4-entity.WallCrumbled1+1) + entity.WallCrumbled1
}

func RandomBed() (int, int) {
	n := rand.Intn(len(entity.Beds)/2) * 2
	return entity.Beds[n], entity.Beds[n+1]
}
