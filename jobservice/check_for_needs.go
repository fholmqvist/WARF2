package jobservice

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/item"
	"github.com/Holmqvist1990/WARF2/job"
	"github.com/Holmqvist1990/WARF2/worldmap"
)

const (
	TIME_FACTOR         = 20
	LIBRARY_READ_CUTOFF = 80
)

func (j *Service) checkForReading(mp *worldmap.Map) {
	for _, dwf := range j.AvailableWorkers {
		if dwf.Needs.ToRead < LIBRARY_READ_CUTOFF {
			continue
		}
		if readingAlreadyExists(j, dwf) {
			continue
		}
		destination, ok := getBookshelfDestination(mp, *dwf)
		if !ok {
			continue
		}
		jb := job.NewLibraryRead([]int{destination}, int(dwf.Characteristics.DesireToRead*TIME_FACTOR))
		SetWorkerAndMove(jb, dwf, mp)
		j.Jobs = append(j.Jobs, jb)
		dwf.Needs.ToRead = 0
	}
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

func readingAlreadyExists(g *Service, dwf *dwarf.Dwarf) bool {
	for _, jb := range g.Jobs {
		rd, ok := jb.(*job.LibraryRead)
		if !ok {
			continue
		}
		if rd.GetWorker() == dwf {
			return true
		}
	}
	return false
}
