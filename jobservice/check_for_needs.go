package jobservice

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/item"
	"github.com/Holmqvist1990/WARF2/job"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

const (
	TIME_FACTOR         = 20
	LIBRARY_READ_CUTOFF = 80
)

func (s *Service) checkForNeeds(mp *m.Map, rs *room.Service) {
	for _, dwf := range s.AvailableWorkers {
		if !dwf.Available() {
			continue
		}
		if checkForSleep(dwf, s.Workers, s, mp, rs) {
			continue
		}
		if checkForDrink(dwf) {
			continue
		}
		if checkForReading(dwf, s, mp) {
			continue
		}
	}
}

func checkForSleep(dwf *dwarf.Dwarf, dwarves []*dwarf.Dwarf, s *Service, mp *m.Map, rs *room.Service) (added bool) {
	if dwf.Needs.Sleep < dwarf.MAX {
		return false
	}
	if sleepAlreadyExists(s, dwf) {
		return false
	}
	bedIdxs, ok := item.FindNearestBeds(mp, dwf.Idx)
	if !ok {
		return false
	}
	// Don't rest in occupied bed.
	target := -1
	for _, dst := range bedIdxs {
		for _, dwarf := range dwarves {
			if dwarf.Idx == dst {
				continue
			}
		}
		target = dst
		break
	}
	// No available beds.
	if target == -1 {
		return false
	}
	jb := job.NewSleep(
		target,
		m.NeighTileFour(target),
	)
	SetWorkerAndMove(jb, dwf, mp)
	s.Jobs = append(s.Jobs, jb)
	dwf.Needs.Sleep = 0
	return true
}

func checkForDrink(dwf *dwarf.Dwarf) bool {
	if dwf.Needs.Drink < dwarf.MAX {
		return false
	}
	fmt.Println("THIRSTY!")
	return false
}

func checkForReading(dwf *dwarf.Dwarf, s *Service, mp *m.Map) (added bool) {
	if dwf.Needs.Read < LIBRARY_READ_CUTOFF {
		return false
	}
	if readingAlreadyExists(s, dwf) {
		return false
	}
	destination, ok := getBookshelfDestination(mp, *dwf)
	if !ok {
		return false
	}
	jb := job.NewRead(
		[]int{destination},
		int(dwf.Characteristics.DesireToRead*TIME_FACTOR),
	)
	SetWorkerAndMove(jb, dwf, mp)
	s.Jobs = append(s.Jobs, jb)
	dwf.Needs.Read = 0
	return true
}

func getBookshelfDestination(mp *m.Map, dwf dwarf.Dwarf) (int, bool) {
	bookshelf, ok := item.FindNearestBookshelf(mp, dwf.Idx)
	if !ok {
		return -1, false
	}
	destination := mp.OneTileDown(bookshelf)
	if !m.IsExposed(destination.Sprite) {
		return -1, false
	}
	return destination.Idx, true
}

func sleepAlreadyExists(s *Service, dwf *dwarf.Dwarf) bool {
	for _, jb := range s.Jobs {
		rd, ok := jb.(*job.Sleep)
		if !ok {
			continue
		}
		if rd.GetWorker() == dwf {
			return true
		}
	}
	return false
}

func readingAlreadyExists(g *Service, dwf *dwarf.Dwarf) bool {
	for _, jb := range g.Jobs {
		rd, ok := jb.(*job.Read)
		if !ok {
			continue
		}
		if rd.GetWorker() == dwf {
			return true
		}
	}
	return false
}
