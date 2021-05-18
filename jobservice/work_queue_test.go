package jobservice

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/job"
	"projects/games/warf2/worldmap"
	"testing"
)

func TestWorkQueue(t *testing.T) {
	js := jobServicemWithJobs()
	originalOrder := []int{}
	for _, v := range js.Jobs {
		originalOrder = append(originalOrder, v.GetDestinations()...)
	}
	js.Update(nil) // TODO
	sameAsBefore := 0
	for i, v := range js.Jobs {
		for _, destination := range v.GetDestinations() {
			if destination == originalOrder[i] {
				sameAsBefore++
			}
		}
	}
	allOfThem := 6
	if allOfThem == sameAsBefore {
		t.Fatalf("should be random, wasn't")
	}
}

func jobServicemWithJobs() *JobService {
	js := &JobService{
		Jobs: []job.Job{
			job.NewLibraryRead([]int{10}, 1),
			job.NewLibraryRead([]int{11}, 1),
			job.NewLibraryRead([]int{12}, 1),
			job.NewDigging([]int{20}, 0),
			job.NewDigging([]int{21}, 0),
			job.NewDigging([]int{22}, 0),
		},
		Map:     worldmap.New(),
		Workers: []*dwarf.Dwarf{},
	}
	js.Map.Tiles[0].Sprite = worldmap.WallSelectedSolid
	return js
}
