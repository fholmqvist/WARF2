package jobsystem

import (
	"projects/games/warf2/job"
	m "projects/games/warf2/worldmap"
)

func (j *JobService) checkForDiggingJobs() {
	for _, wall := range j.Map.Tiles {
		if !m.IsSelectedWall(wall.Sprite) {
			continue
		}
		var destinations []int
		for _, destination := range m.SurroundingTilesFour(wall.Idx) {
			if m.IsColliding(j.Map, wall.Idx, destination.Dir) {
				continue
			}
			if j.diggingJobAlreadyExists(destination.Idx, wall.Idx) {
				continue
			}
			destinations = append(destinations, destination.Idx)

		}
		// We have satisfied the need
		// as a worker is on the way.
		diggingJob := job.NewDigging(destinations, wall.Idx)

		j.Jobs = append(j.Jobs, diggingJob)
	}
}

func (j *JobService) diggingJobAlreadyExists(dIdx, wIdx int) bool {
	for _, jb := range j.Jobs {
		d, ok := jb.(*job.Digging)
		if !ok {
			continue
		}
		for _, destination := range d.GetDestinations() {
			if destination == dIdx && d.GetWallIdx() == wIdx {
				return true
			}
		}

	}
	return false
}
