package jobservice

import (
	"projects/games/warf2/globals"
	"projects/games/warf2/item"
	"projects/games/warf2/job"
	"projects/games/warf2/room"
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

func (j *JobService) diggingJobAlreadyExists(dIdx, jIdx int) bool {
	for _, jb := range j.Jobs {
		d, ok := jb.(*job.Digging)
		if !ok {
			continue
		}
		for _, destination := range d.GetDestinations() {
			if destination == dIdx && d.GetWallIdx() == jIdx {
				return true
			}
		}

	}
	return false
}

func (j *JobService) checkForCarryingJobs(rs *room.Service) {
	for _, it := range j.Map.Items {
		///////////////////////////
		// TODO
		// Handle more than rocks.
		///////////////////////////
		if !item.IsCrumbledWall(it.Sprite) {
			continue
		}
		////////////////////////////////////
		// TODO
		// FloorBrick is _not_ an
		// adequate definition of storage.
		////////////////////////////////////
		if m.IsFloorBrick(j.Map.Tiles[it.Idx].Sprite) {
			continue
		}
		if j.carryingJobAlreadyExists(it.Idx, j.Map) {
			continue
		}
		x, y := globals.IdxToXY(it.Idx)
		nearest, ok := rs.FindNearestStorage(j.Map, x, y)
		if !ok {
			continue
		}
		dst, ok := nearest.GetAvailableTile(it.Resource)
		if !ok {
			continue
		}
		if it.Idx == dst {
			// This shouldn't happen.
			panic("job_service: check_for_job: it.Idx == dst")
		}
		j.Jobs = append(j.Jobs, job.NewCarrying(
			[]int{it.Idx},
			dst,
			it.Sprite,
		))
	}
}

func (j *JobService) carryingJobAlreadyExists(idx int, mp *m.Map) bool {
	for _, jb := range j.Jobs {
		c, ok := jb.(*job.Carrying)
		if !ok {
			continue
		}
		if c.GetDestinations()[0] != idx {
			continue
		}
		////////////////////////////////////
		// TODO
		// FloorBrick is _not_ an
		// adequate definition of storage.
		////////////////////////////////////
		if m.IsFloorBrick(mp.Tiles[idx].Sprite) {
			return true
		}
		for _, jb2 := range j.Jobs {
			c2, ok := jb2.(*job.Carrying)
			if !ok {
				continue
			}
			if c.GetDestinations()[0] == c2.GetDestinations()[0] {
				return true
			}
		}
	}
	return false
}
