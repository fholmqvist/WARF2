package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

const SLEEP_TIME = 600

// Just needs to walk to the destination.
type Sleep struct {
	bedIdx       int
	destinations []int
	dwarf        *dwarf.Dwarf
	sleepTime    int
	arrivedAtIdx int
	remove       bool
}

func NewSleep(bedIdx int, destinations []int) *Sleep {
	return &Sleep{bedIdx, destinations, nil, SLEEP_TIME, -1, false}
}

func (s *Sleep) Remove() bool {
	return s.remove
}

func (s *Sleep) PerformWork(*m.Map, []*dwarf.Dwarf, *room.Service) bool {
	// Put dwarf in bed.
	if s.arrivedAtIdx == -1 {
		s.arrivedAtIdx = s.dwarf.Idx
		s.dwarf.Path = nil
		s.dwarf.Idx = s.bedIdx
	}
	// Finished.
	if s.sleepTime == 0 {
		s.remove = true
		s.dwarf.Idx = s.arrivedAtIdx
		return finished
	}
	// Sleep.
	s.dwarf.Needs.Sleep = 0
	s.sleepTime--
	return unfinished
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
