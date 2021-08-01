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
	Rooms []Room
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Update(mp *m.Map) {
	for _, rm := range s.Rooms {
		rm.Update(mp)
	}
}

func (s *Service) AddStorage(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	st := NewStorage(mp, x, y)
	if st == nil {
		return
	}
	s.Rooms = append(s.Rooms, st)
}

func (s *Service) AddSleepHall(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	sh := NewSleepHall(mp, x, y)
	if sh == nil {
		return
	}
	s.Rooms = append(s.Rooms, sh)
}

func (s *Service) AddFarm(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	f := NewFarm(mp, x, y)
	if f == nil {
		return
	}
	s.Rooms = append(s.Rooms, f)
}

func (s *Service) AddBrewery(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	b := NewBrewery(mp, x, y)
	if b == nil {
		return
	}
	s.Rooms = append(s.Rooms, b)
}

func (s *Service) AddLibrary(mp *m.Map, currentMousePos int) {
	x, y := globals.IdxToXY(currentMousePos)
	l := NewLibrary(mp, x, y)
	if l == nil {
		return
	}
	s.Rooms = append(s.Rooms, l)
}

func (s *Service) GetFarm(id int) (*Farm, bool) {
	for _, rm := range s.Rooms {
		if id != rm.GetID() {
			continue
		}
		farm, ok := rm.(*Farm)
		if !ok {
			continue
		}
		return farm, true
	}
	return nil, false
}

func (s *Service) GetStorage(id int) (*Storage, bool) {
	for _, r := range s.Rooms {
		if r.GetID() != id {
			continue
		}
		storage, ok := r.(*Storage)
		if !ok {
			continue
		}
		return storage, true
	}
	return nil, false
}

func (s *Service) FindNearestStorage(mp *m.Map, x, y int, res entity.Resource) (*Storage, int, bool) {
	closest := math.MaxFloat64
	idx := -1
	for i, rm := range s.Rooms {
		storage, ok := rm.(*Storage)
		if !ok {
			continue
		}
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
	return s.Rooms[idx].(*Storage), idx, true
}
