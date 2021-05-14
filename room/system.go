package room

import (
	"fmt"
	"math"
	"projects/games/warf2/globals"
	m "projects/games/warf2/worldmap"
)

// System for gathering data
// and functionality related to rooms.
type System struct {
	Storages  []Storage
	Libraries []Library
}

func (s *System) AddLibrary(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	l := NewLibrary(mp, x, y)
	if l != nil {
		s.Libraries = append(s.Libraries, *l)
	}
}

func (s *System) AddStorage(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	st := NewStorage(mp, x, y)
	if st != nil {
		s.Storages = append(s.Storages, *st)
	}
}

func (s *System) FindNearestStorage(mp *m.Map, x, y int) (*Storage, bool) {
	if len(s.Storages) == 0 {
		return nil, false
	}
	closest := math.MaxFloat64
	idx := -1
	for i, storage := range s.Storages {
		fmt.Println("Current:", storage.Center)
		bx, by := globals.IdxToXY(storage.Center)
		d := globals.Dist(x, y, bx, by)
		if d < closest {
			closest = d
			idx = i
			fmt.Println("Was closer.")
		} else {
			fmt.Println("Wasn't closer.")
		}
	}
	return &s.Storages[idx], true
}
