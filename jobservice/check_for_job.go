package jobservice

import (
	"github.com/Holmqvist1990/WARF2/entity"
	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/job"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func (s *Service) checkForJobs(rs *room.Service) {
	for _, tile := range s.Map.Tiles {
		if checkForDiggingJob(s, tile) {
			continue
		}
	}
	for _, itm := range s.Map.Items {
		if checkForCarryingJob(s, itm, rs) {
			continue
		}
	}
	for _, rm := range rs.Rooms {
		if f, ok := rm.(*room.Farm); ok {
			if f == nil {
				continue
			}
			checkForFarmingJobs(s, *f, rs)
		}
		if b, ok := rm.(*room.Brewery); ok {
			if b == nil {
				continue
			}
			checkBreweryJobs(s, *b, rs)
		}
	}
}

func checkForDiggingJob(s *Service, wall m.Tile) (added bool) {
	if !m.IsSelectedWall(wall.Sprite) {
		return false
	}
	var destinations []int
	for _, destination := range m.SurroundingTilesFour(wall.Idx) {
		if m.IsColliding(s.Map, wall.Idx, destination.Dir) {
			continue
		}
		if diggingJobExists(s, destination.Idx, wall.Idx) {
			continue
		}
		destinations = append(destinations, destination.Idx)
	}
	// We have satisfied the need
	// as a worker is on the way.
	diggingJob := job.NewDigging(destinations, wall.Idx)
	s.Jobs = append(s.Jobs, diggingJob)
	return true
}

func checkForCarryingJob(s *Service, itm m.Tile, rs *room.Service) (added bool) {
	if itm.Resource == entity.ResourceNone {
		return false
	}
	if !entity.IsCarriable(itm.Sprite) {
		return false
	}
	////////////////////////////////////
	// TODO
	// FloorBrick is _not_ an
	// adequate definition of storage.
	////////////////////////////////////
	if m.IsStorageFloorBrick(s.Map.Tiles[itm.Idx].Sprite) {
		return false
	}
	if carryingJobExists(s, itm.Idx, s.Map) {
		return false
	}
	x, y := gl.IdxToXY(itm.Idx)
	nearest, storageIdx, ok := rs.FindNearestStorage(s.Map, x, y, itm.Resource)
	if !ok {
		return false
	}
	dst, ok := nearest.GetAvailableTile(itm.Resource)
	if !ok {
		return false
	}
	if itm.Idx == dst {
		// If the item is to be carried
		// to its own position, something
		// has gone terribly wrong.
		panic("job_service: check_for_job: it.Idx == dst")
	}
	if _, ok := m.CreatePath(&s.Map.Tiles[itm.Idx], &s.Map.Tiles[dst]); !ok {
		return false
	}
	s.Jobs = append(s.Jobs, job.NewCarrying(
		[]int{itm.Idx},
		entity.SpriteToResource(itm.Sprite),
		storageIdx,
		dst,
		itm.Sprite,
	))
	return true
}

func checkForFarmingJobs(s *Service, farm room.Farm, rs *room.Service) {
	if farm.FullyHarvestedAndCleaned(s.Map) {
		if plantFarmJobExists(s, farm) {
			return
		}
		s.Jobs = append(s.Jobs, job.NewPlantFarm(&farm, farm.FarmableIdxs))
		return
	}
	idxs, should := farm.ShouldHarvest(s.Map)
	if !should {
		return
	}
	if farmJobExists(s, farm) {
		return
	}
	s.Jobs = append(s.Jobs, job.NewFarming(farm.ID, idxs))
	return
}

func checkBreweryJobs(s *Service, brewery room.Brewery, rs *room.Service) {
	for _, rm := range rs.Rooms {
		storage, ok := rm.(*room.Storage)
		if !ok || storage == nil {
			continue
		}
		storageTile, has := storage.HasWheat()
		if !has {
			continue
		}
		barrelIdx, ok := brewery.GetEmptyBarrel(s.Map)
		if !ok {
			continue
		}
		if fillBarrelJobExists(s, storageTile.Idx, barrelIdx) {
			continue
		}
		s.Jobs = append(s.Jobs, job.NewFillBrewer(
			storageTile,
			barrelIdx,
			m.TileDirsToIdxs(m.SurroundingTilesFour(barrelIdx))),
		)
	}
}

func diggingJobExists(s *Service, dIdx, jIdx int) bool {
	for _, jb := range s.Jobs {
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

func carryingJobExists(s *Service, idx int, mp *m.Map) bool {
	for _, jb1 := range s.Jobs {
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
		if m.IsStorageFloorBrick(mp.Tiles[idx].Sprite) {
			return true
		}
		for _, jb2 := range s.Jobs {
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

func plantFarmJobExists(s *Service, farm room.Farm) bool {
	for _, j := range s.Jobs {
		p, ok := j.(*job.PlantFarm)
		if !ok {
			continue
		}
		if p.Farm.ID == farm.ID {
			return true
		}
	}
	return false
}

func farmJobExists(s *Service, farm room.Farm) bool {
	for _, j := range s.Jobs {
		f, ok := j.(*job.Farming)
		if !ok {
			continue
		}
		if f.FarmID == farm.ID {
			return true
		}
	}
	return false
}

func fillBarrelJobExists(s *Service, wheatIdx int, barrelIndex int) bool {
	for _, j := range s.Jobs {
		f, ok := j.(*job.FillBarrel)
		if !ok {
			continue
		}
		if f.WheatIndex == wheatIdx || f.BarrelIndex == barrelIndex {
			return true
		}
	}
	return false
}
