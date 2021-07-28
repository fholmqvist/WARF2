package jobservice

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/item"
	"github.com/Holmqvist1990/WARF2/job"
	"github.com/Holmqvist1990/WARF2/room"
	"github.com/Holmqvist1990/WARF2/worldmap"
)

const (
	TIME_FACTOR         = 20
	LIBRARY_READ_CUTOFF = 80
)

func (s *Service) checkForNeeds(mp *worldmap.Map, rs *room.Service) {
	for _, dwf := range s.AvailableWorkers {
		if !dwf.Available() {
			continue
		}
		if checkForSleep(dwf, s, mp, rs) {
			continue
		}
		if checkForReading(dwf, s, mp) {
			continue
		}
	}
}

func checkForSleep(dwf *dwarf.Dwarf, s *Service, mp *worldmap.Map, rs *room.Service) (added bool) {
	if dwf.Needs.Sleep < dwarf.MAX_NEED {
		return false
	}
	if sleepAlreadyExists(s, dwf) {
		return false
	}
	//////////////////////////////////////
	// TODO
	// FindNearestBeds (plural).
	// Don't rest in occupied bed.
	//////////////////////////////////////
	bedIndex, ok := item.FindNearestBed(mp, dwf.Idx)
	if !ok {
		return false
	}
	jb := job.NewSleep(
		bedIndex,
		worldmap.TileDirsToIdxs(worldmap.SurroundingTilesFour(bedIndex)),
	)
	SetWorkerAndMove(jb, dwf, mp)
	s.Jobs = append(s.Jobs, jb)
	dwf.Needs.Sleep = 0
	return true
}

func checkForReading(dwf *dwarf.Dwarf, s *Service, mp *worldmap.Map) (added bool) {
		if dwf.Needs.ToRead < LIBRARY_READ_CUTOFF {
			return false
		}
		if readingAlreadyExists(s, dwf) {
			return false
		}
		destination, ok := getBookshelfDestination(mp, *dwf)
		if !ok {
			return false
		}
		jb := job.NewLibraryRead(
			[]int{destination},
			int(dwf.Characteristics.DesireToRead*TIME_FACTOR),
		)
		SetWorkerAndMove(jb, dwf, mp)
		s.Jobs = append(s.Jobs, jb)
		dwf.Needs.ToRead = 0
		return true
}

func getBookshelfDestination(mp *worldmap.Map, dwf dwarf.Dwarf) (int, bool) {
	bookshelf, ok := item.FindNearestBookshelf(mp, dwf.Idx)
	if !ok {
		return -1, false
	}
	destination := mp.OneTileDown(bookshelf)
	if !worldmap.IsExposed(destination.Sprite) {
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
