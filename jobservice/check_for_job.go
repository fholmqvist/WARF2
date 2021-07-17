package jobservice

import (
	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/job"
	"github.com/Holmqvist1990/WARF2/resource"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func (j *Service) checkForDiggingJobs() {
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

func (j *Service) diggingJobAlreadyExists(dIdx, jIdx int) bool {
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

func (j *Service) checkForCarryingJobs(rs *room.Service) {
	for _, it := range j.Map.Items {
		if it.Resource == resource.None {
			continue
		}
		if !typeOfCarryingJob(it.Sprite) {
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
		x, y := gl.IdxToXY(it.Idx)
		nearest, storageIdx, ok := rs.FindNearestStorage(j.Map, x, y)
		if !ok {
			continue
		}
		dst, ok := nearest.GetAvailableTile(it.Resource)
		if !ok {
			continue
		}
		if it.Idx == dst {
			// If the item is to be carried
			// to its own position, something
			// has gone terribly wrong.
			panic("job_service: check_for_job: it.Idx == dst")
		}
		j.Jobs = append(j.Jobs, job.NewCarrying(
			[]int{it.Idx},
			resource.SpriteToResource(it.Sprite),
			storageIdx,
			dst,
			it.Sprite,
		))
	}
}

func (j *Service) carryingJobAlreadyExists(idx int, mp *m.Map) bool {
	for _, jb1 := range j.Jobs {
		c1, ok := jb1.(*job.Carrying)
		if !ok {
			continue
		}
		if c1.GetDestinations()[0] != idx {
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
			if c1.GetDestinations()[0] == c2.GetDestinations()[0] {
				return true
			}
		}
	}
	return false
}

func (j *Service) checkForFarmingJobs(rs *room.Service) {
	for _, farm := range rs.Farms {
		if farm.FullyHarvestedAndCleaned(j.Map) {
			if j.plantFarmJobAlreadyExists(farm) {
				continue
			}
			j.Jobs = append(j.Jobs, job.NewPlantFarm(farm.ID))
			continue
		}
		idxs, should := farm.ShouldHarvest(j.Map)
		if !should {
			continue
		}
		if j.farmJobAlreadyExists(farm) {
			continue
		}
		j.Jobs = append(j.Jobs, job.NewFarming(farm.ID, idxs))
	}
}

func (j *Service) plantFarmJobAlreadyExists(farm room.Farm) bool {
	for _, j := range j.Jobs {
		p, ok := j.(*job.PlantFarm)
		if !ok {
			return false
		}
		if p.FarmID == farm.ID {
			return true
		}
	}
	return false
}

func (j *Service) farmJobAlreadyExists(farm room.Farm) bool {
	for _, j := range j.Jobs {
		f, ok := j.(*job.Farming)
		if !ok {
			return false
		}
		if f.FarmID == farm.ID {
			return true
		}
	}
	return false
}

func typeOfCarryingJob(sprite int) bool {
	return !gl.IsCrumbledWall(sprite) || !gl.IsFarmTileHarvested(sprite)
}
