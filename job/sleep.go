package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

// Just needs to walk to the destination.
type Sleep struct {
	bedIdx       int
	destinations []int
	dwarf        *dwarf.Dwarf
	sleepTime    int
	arrivedAtIdx int
}

func NewSleep(bedIdx int, destinations []int) *Sleep {
	sleep := 600
	return &Sleep{bedIdx, destinations, nil, sleep, -1}
}

func (s *Sleep) NeedsToBeRemoved(*m.Map, *room.Service) bool {
	return s.dwarf != nil && s.sleepTime == 0
}

func (s *Sleep) PerformWork(*m.Map, []*dwarf.Dwarf, *room.Service) bool {
	if s.dwarf.State == dwarf.Moving {
		return unfinished
	}
	if s.sleepTime == 0 {
		return finished
	}
	if s.arrivedAtIdx == -1 {
		s.arrivedAtIdx = s.dwarf.Idx
		s.dwarf.Idx = s.bedIdx
	}
	s.sleepTime--
	return unfinished
}

func (s *Sleep) Finish(*m.Map, *room.Service) {
	if s.dwarf == nil {
		return
	}
	s.dwarf.Idx = s.arrivedAtIdx
	s.dwarf.SetToAvailable()
	s.dwarf = nil
}

func (s *Sleep) GetWorker() *dwarf.Dwarf {
	return s.dwarf
}

func (s *Sleep) SetWorker(d *dwarf.Dwarf) {
	s.dwarf = d
}

func (s *Sleep) GetDestinations() []int {
	return s.destinations
}

func (s *Sleep) HasInternalMove() bool {
	return false
}

func (s *Sleep) String() string {
	return "Sleep"
}
