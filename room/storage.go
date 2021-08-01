package room

import (
	"math"

	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var storageAutoID = 0

///////////////////////////////////
// TODO
// If the storage is filled
// with StorageTiles that
// aren't maxxed, we should
// cleanup so that we have
// fewer yet filled tiles.
//
// This opens more slots and
// prevents storages from being
// artificially full due to
// saturation of one specific tile.
//
// Now, bad:
// [1 Rock, 2 Rock, 1 Rock, 4 Rock]
//
// Later, good:
// [4 Rock, 4 Rock,   None,   None]
//
///////////////////////////////////
type Storage struct {
	ID           int
	Center       int
	StorageTiles []StorageTile
}

func NewStorage(mp *m.Map, x, y int) *Storage {
	s := &Storage{}
	tiles := mp.FloodFillRoom(x, y, m.RandomFloorBrick)
	if len(tiles) == 0 {
		return nil
	}
	for _, idxs := range tiles {
		mp.Tiles[idxs].Room = s
	}
	s.StorageTiles = createStorageTiles(mp, tiles)
	s.Center = determineCenter(mp, tiles)
	s.ID = storageAutoID
	storageAutoID++
	return s
}

func (s *Storage) GetID() int {
	return s.ID
}

func (s *Storage) Update(mp *m.Map) {}

func (s *Storage) GetAvailableTile(r entity.Resource) (idx int, ok bool) {
	for _, t := range s.StorageTiles {
		if t.Unavailable(r) {
			continue
		}
		return t.Idx, true
	}
	return -1, false
}

func (s *Storage) AddItem(idx int, amount uint, r entity.Resource) (int, bool) {
	tIdx, ok := s.getStorageTileIdxFromWorldIdx(idx)
	if !ok {
		return -1, false
	}
	st := s.StorageTiles[tIdx]
	if st.Available(r) {
		// Add to nearest.
		s.StorageTiles[tIdx].Add(r, amount)
		return st.Idx, true
	}
	for idx, st := range s.StorageTiles {
		if st.Unavailable(r) {
			continue
		}
		// Add to first available.
		s.StorageTiles[idx].Add(r, amount)
		return st.Idx, true
	}
	return -1, false
}

func (s *Storage) HasSpace(res entity.Resource) bool {
	for _, t := range s.StorageTiles {
		if t.Available(res) {
			return true
		}
	}
	return false
}

func (s *Storage) String() string {
	return "Storage"
}

func (s *Storage) Tiles() []int {
	idxs := make([]int, len(s.StorageTiles))
	for i, t := range s.StorageTiles {
		idxs[i] = t.Idx
	}
	return idxs
}

func determineCenter(mp *m.Map, tiles []int) int {
	minx, maxx := math.MaxInt16, -1
	miny, maxy := math.MaxInt16, -1
	for _, idx := range tiles {
		if mp.Tiles[idx].X < minx {
			minx = mp.Tiles[idx].X
		}
		if mp.Tiles[idx].X > maxx {
			maxx = mp.Tiles[idx].X
		}
		if mp.Tiles[idx].Y < miny {
			miny = mp.Tiles[idx].Y
		}
		if mp.Tiles[idx].Y > maxy {
			maxy = mp.Tiles[idx].Y
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

func (s *Storage) getStorageTileIdxFromWorldIdx(idx int) (int, bool) {
	for worldIndex, t := range s.StorageTiles {
		if t.Idx == idx {
			return worldIndex, true
		}
	}
	return -1, false
}
