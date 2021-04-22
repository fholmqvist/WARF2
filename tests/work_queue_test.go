package tests

import (
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

func jobSystemWithJobs() *j.JobSystem {
	js := &j.JobSystem{
		Jobs: []j.Job{
			j.NewLibraryRead(10, 1),
			j.NewLibraryRead(11, 1),
			j.NewLibraryRead(12, 1),
			j.NewDigging(20, 0),
			j.NewDigging(21, 0),
			j.NewDigging(22, 0),
		},
		Map:     worldmap.New(),
		Workers: []j.Worker{},
	}
	js.Map.Tiles[0].Sprite = worldmap.WallSelectedSolid
	return js
}
