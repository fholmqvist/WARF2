package jobservice

import (
	"fmt"

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

func (s *Service) checkForSleep(mp *worldmap.Map, rs *room.Service) {
	for _, dwf := range s.AvailableWorkers {
		if dwf.Needs.Sleep < dwarf.MAX_NEED {
			continue
		}
		if sleepAlreadyExists(s, dwf) {
			continue
		}
		//////////////////////////////////////
		// TODO
		// Crashes here.
		// Needs to pick position next to bed.
		//////////////////////////////////////
		bedIndex, ok := item.FindNearestBed(mp, dwf.Idx)
		if !ok {
			continue
		}
		fmt.Println("NEW SLEEP")
		jb := job.NewSleep(bedIndex)
		SetWorkerAndMove(jb, dwf, mp)
		s.Jobs = append(s.Jobs, jb)
		dwf.Needs.Sleep = 0
	}
}

func (s *Service) checkForReading(mp *worldmap.Map) {
	for _, dwf := range s.AvailableWorkers {
		if dwf.Needs.ToRead < LIBRARY_READ_CUTOFF {
			continue
		}
		if readingAlreadyExists(s, dwf) {
			continue
		}
		destination, ok := getBookshelfDestination(mp, *dwf)
		if !ok {
			continue
		}
		jb := job.NewLibraryRead([]int{destination}, int(dwf.Characteristics.DesireToRead*TIME_FACTOR))
		SetWorkerAndMove(jb, dwf, mp)
		s.Jobs = append(s.Jobs, jb)
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
