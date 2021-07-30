package room

import (
	"math"

	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

// Service for gathering data
// and functionality related to rooms.
type Service struct {
	Storages   []Storage
	SleepHalls []SleepHall
	Farms      []Farm
	Libraries  []Library
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Update(mp *m.Map) {
	for _, f := range s.Farms {
		f.Update(mp)
	}
}

func (s *Service) AddSleepHall(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	sh := NewSleepHall(mp, x, y)
	if sh == nil {
		return
	}
	s.SleepHalls = append(s.SleepHalls, *sh)
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

func (s *Service) GetStorage(storageIdx int) (*Storage, bool) {
	/////////////////////////
	// TODO
	// Should be ID, not IDX.
	/////////////////////////
	if storageIdx > len(s.Storages) {
		return nil, false
	}
	return &s.Storages[storageIdx], true
}

func (s *Service) FindNearestStorage(mp *m.Map, x, y int, res entity.Resource) (*Storage, int, bool) {
	if len(s.Storages) == 0 {
		return nil, -1, false
	}
	if len(s.Storages) == 1 {
		st := &s.Storages[0]
		if st.HasSpace(res) {
			return st, 0, true
		} else {
			return nil, -1, false
		}
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
	if idx == -1 {
		// No available space
		// in any storage.
		return nil, -1, false
	}
	return &s.Storages[idx], idx, true
}
