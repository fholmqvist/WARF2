package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/item"
	"github.com/Holmqvist1990/WARF2/room"
	"github.com/Holmqvist1990/WARF2/worldmap"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Read struct {
	dwarf        *dwarf.Dwarf
	destinations []int
	readingTime  int
}

func NewLibraryRead(destinations []int, readingTime int) *Read {
	return &Read{nil, destinations, readingTime}
}

func (l *Read) NeedsToBeRemoved(*m.Map, *room.Service) bool {
	return l.readingTime <= 0 || l.dwarf == nil
}

func (l *Read) Finish(*m.Map, *room.Service) {
	l.dwarf = nil
}

func (l *Read) PerformWork(m *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	if shouldGetChair(m, l) {
		return getChair(m, l, dwarves)
	}
	// Still reading.
	if l.readingTime > 1 {
		l.readingTime--
		return unfinished
	}
	// Done!
	l.readingTime = 0
	return finished
}

func (l *Read) GetWorker() *dwarf.Dwarf {
	return l.dwarf
}

func (l *Read) SetWorker(dw *dwarf.Dwarf) {
	l.dwarf = dw
}

func (l *Read) GetDestinations() []int {
	return l.destinations
}

func (l *Read) HasInternalMove() bool {
	return false
}

func (l *Read) String() string {
	return "Library"
}

func shouldGetChair(m *worldmap.Map, l *Read) bool {
	return !gl.IsChair(m.Items[l.dwarf.Idx].Sprite) &&
		l.dwarf.State != dwarf.Moving
}

func getChair(m *worldmap.Map, l *Read, dwarves []*dwarf.Dwarf) bool {
	dsts, ok := item.FindNearestChairs(m, l.destinations[0])
	if !ok {
		return unfinished
	}
	// Don't sit in occupied chair.
	target := -1
	for _, dst := range dsts {
		for _, dwarf := range dwarves {
			if dwarf.Idx == dst {
				continue
			}
		}
		target = dst
		break
	}
	// No available chairs.
	if target == -1 {
		return unfinished
	}
	l.destinations[0] = target
	l.dwarf.MoveTo(target, m)
	return unfinished
}
