package room

import (
	"math"

	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/resource"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

// Service for gathering data
// and functionality related to rooms.
type Service struct {
	Storages  []Storage
	Farms     []Farm
	Libraries []Library
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Update(mp *m.Map) {
	for _, f := range s.Farms {
		f.Update(mp)
	}
}

func (s *Service) AddLibrary(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	l := NewLibrary(mp, x, y)
	if l == nil {
		return
	}
	s.Libraries = append(s.Libraries, *l)
}

func (s *Service) AddFarm(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	f := NewFarm(mp, x, y)
	if f == nil {
		return
	}
	s.Farms = append(s.Farms, *f)
}

func (s *Service) AddStorage(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	st := NewStorage(mp, x, y)
	if st == nil {
		return
	}
	s.Storages = append(s.Storages, *st)
}

func (s *Service) GetFarm(farmID int) (*Farm, bool) {
	for _, f := range s.Farms {
		if farmID == f.ID {
			return &f, true
		}
	}
	return nil, false
}

func (s *Service) FindNearestStorage(mp *m.Map, x, y int, res resource.Resource) (*Storage, int, bool) {
	if len(s.Storages) == 0 {
		return nil, -1, false
	}
	closest := math.MaxFloat64
	idx := -1
	for i, storage := range s.Storages {
		if !storage.HasSpace(res) {
			continue
		}
		bx, by := globals.IdxToXY(storage.Center)
		d := globals.Dist(x, y, bx, by)
		if d < closest {
			closest = d
			idx = i
		}
	}
	return &s.Storages[idx], idx, true
}
