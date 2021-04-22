package game

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/item"
	"projects/games/warf2/jobsystem"
	"projects/games/warf2/worldmap"
)

const (
	TIME_FACTOR         = 20
	LIBRARY_READ_CUTOFF = 0
)

func (g *Game) checkForLibraryReading() {
	for _, worker := range g.JobSystem.AvailableWorkers {
		dwf := worker.(*dwarf.Dwarf)
		if dwf.Needs.ToRead < LIBRARY_READ_CUTOFF {
			continue
		}
		destination, ok := getBookshelfDestination(&g.WorldMap, *dwf)
		if !ok {
			continue
		}
		j := jobsystem.NewLibraryRead(destination, int(dwf.Characteristics.DesireToRead*TIME_FACTOR))
		jobsystem.SetWorkerAndMove(j, worker, &g.WorldMap)
		g.JobSystem.Jobs = append(g.JobSystem.Jobs, j)
		/////////////////////////////////////////////////
		// TODO
		//
		// This is not great.
		/////////////////////////////////////////////////
		dwf.Needs.ToRead = 0
	}
}

func getBookshelfDestination(m *worldmap.Map, dwf dwarf.Dwarf) (int, bool) {
	bookshelf, ok := item.FindNearestBookshelf(m, dwf.Idx)
	if !ok {
		return -1, false
	}
	destination := m.OneTileDown(bookshelf)
	if !worldmap.IsExposed(destination.Sprite) {
		return -1, false
	}
	return destination.Idx, true
}
