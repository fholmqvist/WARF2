package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

// Just needs to walk to the destination.
type Sleep struct {
	destinations []int
	dwarf        *dwarf.Dwarf
}

func NewSleep(destination int) *Sleep {
	return &Sleep{[]int{destination}, nil}
}

func (s *Sleep) NeedsToBeRemoved(*m.Map, *room.Service) bool {
	return false
}

func (s *Sleep) PerformWork(*m.Map, []*dwarf.Dwarf, *room.Service) bool {
	return finished
}

func (s *Sleep) Finish(*m.Map, *room.Service) {
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
