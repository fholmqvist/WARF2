package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/item"
	"github.com/Holmqvist1990/WARF2/room"
	"github.com/Holmqvist1990/WARF2/worldmap"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Read struct {
	JobBase
	readingTime int
	remove      bool
}

func NewRead(destinations []int, readingTime int) *Read {
	return &Read{NewJobBase(destinations), readingTime, false}
}

func (l *Read) PerformWork(m *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	if needsChair(m, l) {
		return getChair(m, l, dwarves)
	}
	// Reading.
	if l.readingTime > 1 {
		l.readingTime--
		return unfinished
	}
	// Finished.
	l.readingTime = 0
	l.remove = true
	return finished
}

func (l *Read) HasInternalMove() bool {
	return false
}

func (l *Read) String() string {
	return "Library"
}

func needsChair(m *worldmap.Map, l *Read) bool {
	return !entity.IsChair(m.Items[l.dwarf.Idx].Sprite) &&
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
