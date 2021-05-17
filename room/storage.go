package room

import (
	"math"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/globals"
	"projects/games/warf2/resource"
	m "projects/games/warf2/worldmap"
)

type Storage struct {
	Center       int
	Tiles        m.Tiles
	StorageTiles []StorageTile
}

type StorageTile struct {
	Idx    int
	Tpe    resource.Resource
	Amount uint
}

func NewStorage(mp *m.Map, x, y int) *Storage {
	s := &Storage{}
	tiles := mp.FloodFillRoom(x, y, m.RandomFloorBrick)
	if len(tiles) == 0 {
		return nil
	}
	s.Tiles = tiles
	s.StorageTiles = createStorageTiles(tiles)
	s.Center = determineCenter(mp, tiles)
	return s
}

// Use storage.
func (s *Storage) Use(dwarf *dwarf.Dwarf) {
	// Nothing yet.
}

func determineCenter(mp *m.Map, tiles m.Tiles) int {
	minx, maxx := math.MaxInt16, -1
	miny, maxy := math.MaxInt16, -1
	for _, t := range tiles {
		if t.X < minx {
			minx = t.X
		}
		if t.X > maxx {
			maxx = t.X
		}
		if t.Y < miny {
			miny = t.Y
		}
		if t.Y > maxy {
			maxy = t.Y
		}
	}
	midx := minx + ((maxx - minx) / 2)
	midy := miny + ((maxy - miny) / 2)
	center := globals.XYToIdx(midx, midy)
	for m.IsAnyWall(mp.Tiles[center].Sprite) {
		center++
	}
	return center
}

func createStorageTiles(tt m.Tiles) []StorageTile {
	var st []StorageTile
	for _, t := range tt {
		st = append(st, StorageTile{
			Idx:    t.Idx,
			Tpe:    resource.None,
			Amount: 0,
		})
	}
	return st
}
