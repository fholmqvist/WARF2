package game

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/jobsystem"
)

const (
	TIME_FACTOR         = 20
	LIBRARY_READ_CUTOFF = 50
)

func (g *Game) checkForLibraryReading() {
	for _, worker := range g.JobSystem.AvailableWorkers {
		d := worker.(*dwarf.Dwarf)
		if d.Needs.ToRead < LIBRARY_READ_CUTOFF {
			continue
		}
		j := jobsystem.NewLibraryRead(d, 0, int(d.Characteristics.DesireToRead*TIME_FACTOR))
		d.SetJob(j)
		g.JobSystem.Jobs = append(g.JobSystem.Jobs, j)
		/////////////////////////////////////////////////
		// TODO
		//
		// This is not great.
		/////////////////////////////////////////////////
		d.Needs.ToRead = 0
	}
}
