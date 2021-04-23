package jobsystem

import (
	"projects/games/warf2/job"
	m "projects/games/warf2/worldmap"
)

func (j *JobSystem) checkForDiggingJobs() {
	for _, wall := range j.Map.Tiles {
		if !m.IsSelectedWall(wall.Sprite) || !wall.NeedsInteraction {
			continue
		}

		hasFoundJob := false
		for _, destination := range m.SurroundingTilesFour(wall.Idx) {
			if hasFoundJob {
				break
			}

			if m.IsColliding(j.Map, wall.Idx, destination.Dir) {
				continue
			}

			if j.diggingJobAlreadyExists(destination.Idx, wall.Idx) {
				continue
			}

			diggingJob := job.NewDigging(destination.Idx, wall.Idx)

			// We have satisfied the need
			// as a worker is on the way.
			wall.NeedsInteraction = false

			j.Jobs = append(j.Jobs, diggingJob)
			hasFoundJob = true
		}
	}
}

func (j *JobSystem) diggingJobAlreadyExists(dIdx, wIdx int) bool {
	for _, jb := range j.Jobs {
		d, ok := jb.(*job.Digging)
		if !ok {
			continue
		}
		if d.GetDestination() == dIdx && d.GetWallIdx() == wIdx {
			return true
		}
	}
	return false
}
