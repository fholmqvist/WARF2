package room

import (
	"math"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/globals"
	m "projects/games/warf2/worldmap"
)

type Storage struct {
	Center int
	tiles  m.Tiles
}

func NewStorage(mp *m.Map, x, y int) *Storage {
	s := &Storage{}
	tiles := mp.FloodFillRoom(x, y, m.RandomFloorBrick)
	if len(tiles) == 0 {
		return nil
	}
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
