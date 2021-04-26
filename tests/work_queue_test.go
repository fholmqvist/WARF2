package tests

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/job"
	j "projects/games/warf2/jobsystem"
	"projects/games/warf2/worldmap"
	"testing"
)

func TestWorkQueue(t *testing.T) {
	js := jobSystemWithJobs()
	originalOrder := []int{}
	for _, v := range js.Jobs {
		originalOrder = append(originalOrder, v.GetDestination())
	}
	js.Update()
	sameAsBefore := 0
	for i, v := range js.Jobs {
		if v.GetDestination() == originalOrder[i] {
			sameAsBefore++
		}
	}
	allOfThem := 6
	if allOfThem == sameAsBefore {
		t.Fatalf("should be random, wasn't")
	}
}

func jobSystemWithJobs() *j.JobService {
	js := &j.JobService{
		Jobs: []job.Job{
			job.NewLibraryRead(10, 1),
			job.NewLibraryRead(11, 1),
			job.NewLibraryRead(12, 1),
			job.NewDigging(20, 0),
			job.NewDigging(21, 0),
			job.NewDigging(22, 0),
		},
		Map:     worldmap.New(),
		Workers: []*dwarf.Dwarf{},
	}
	js.Map.Tiles[0].Sprite = worldmap.WallSelectedSolid
	return js
}
