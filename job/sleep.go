package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

const SLEEP_TIME = 600

type Sleep struct {
	JobBase
	bedIdx       int
	sleepTime    int
	arrivedAtIdx int
}

func NewSleep(bedIdx int, destinations []int) *Sleep {
	return &Sleep{
		NewJobBase(destinations),
		bedIdx,
		SLEEP_TIME,
		-1,
	}
}

func (s *Sleep) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	// Just arrived.
	if s.arrivedAtIdx == -1 {
		// Try again elsehwere.
		if s.bedOccupied(dwarves) {
			s.dwarf.Needs.Sleep = dwarf.MAX
			s.remove = true
			return finished
		}
		// Put dwarf in bed.
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

func (s *Sleep) HasInternalMove() bool {
	return false
}

func (s *Sleep) String() string {
	return "Sleep"
}

func (s *Sleep) bedOccupied(dwarves []*dwarf.Dwarf) bool {
	for _, dwf := range dwarves {
		if dwf.Idx == s.bedIdx {
			return true
		}
	}
	return false
}
