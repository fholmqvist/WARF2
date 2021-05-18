package room

import (
	"math"
	"projects/games/warf2/globals"
	m "projects/games/warf2/worldmap"
)

// Service for gathering data
// and functionality related to rooms.
type Service struct {
	Storages  []Storage
	Libraries []Library
}

func (s *Service) AddLibrary(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	l := NewLibrary(mp, x, y)
	if l != nil {
		s.Libraries = append(s.Libraries, *l)
	}
}

func (s *Service) AddStorage(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	st := NewStorage(mp, x, y)
	if st != nil {
		s.Storages = append(s.Storages, *st)
	}
}

func (s *Service) FindNearestStorage(mp *m.Map, x, y int) (*Storage, int, bool) {
	if len(s.Storages) == 0 {
		return nil, -1, false
	}
	closest := math.MaxFloat64
	idx := -1
	for i, storage := range s.Storages {
		bx, by := globals.IdxToXY(storage.Center)
		d := globals.Dist(x, y, bx, by)
		if d < closest {
			closest = d
			idx = i
		}
	}
	return &s.Storages[idx], idx, true
}
