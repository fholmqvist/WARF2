package tests

import (
	j "projects/games/warf2/jobsystem"
	"projects/games/warf2/worldmap"
	"testing"
)

func TestWorkQueue(t *testing.T) {
	js := &j.JobSystem{
		Jobs: []j.Job{
			j.NewLibraryRead(nil, 0, 1),
			j.NewLibraryRead(nil, 1, 1),
			j.NewLibraryRead(nil, 2, 1),
		},
		Map:     &worldmap.Map{},
		Workers: []j.Worker{},
	}
	firstOrder := []int{}
	for _, v := range js.Jobs {
		firstOrder = append(firstOrder, v.GetDestination())
	}
	js.Update()
	identical := 0
	for i, v := range js.Jobs {
		if v.GetDestination() == firstOrder[i] {
			identical++
		}
	}
	if identical == 3 {
		t.Fatalf("should be random, wasn't")
	}
}
