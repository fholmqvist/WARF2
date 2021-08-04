package room

import (
	"fmt"
	"math"
	"sync"

	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

// Service for gathering data
// and functionality related to rooms.
type Service struct {
	Rooms []Room
	sync.RWMutex
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Update(mp *m.Map) {
	for _, rm := range s.Rooms {
		rm.Update(mp)
	}
}

func (s *Service) AddRoom(mp *m.Map, currentMousePos int, rm Room) {
	s.Lock()
	defer s.Unlock()
	x, y := globals.IdxToXY(currentMousePos)
	var newRoom Room
	switch rm.(type) {
	case *Storage:
		newRoom = NewStorage(mp, x, y)
	case *SleepHall:
		newRoom = NewSleepHall(mp, x, y)
	case *Farm:
		newRoom = NewFarm(mp, x, y)
	case *Brewery:
		newRoom = NewBrewery(mp, x, y)
	case *Bar:
		newRoom = NewBar(mp, x, y)
	case *Library:
		newRoom = NewLibrary(mp, x, y)
	default:
		panic(fmt.Sprintf("unknown room type: %v", rm))
	}
	s.Rooms = append(s.Rooms, newRoom)
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
		if !ok || storage == nil {
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
