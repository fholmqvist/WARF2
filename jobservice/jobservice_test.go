package jobservice

import (
	"testing"

	"github.com/Holmqvist1990/WARF2/job"
	"github.com/Holmqvist1990/WARF2/resource"
)

func TestWorkQueue(t *testing.T) {
	js := &JobService{
		Jobs: []job.Job{
			job.NewLibraryRead([]int{10}, 1),
			job.NewDigging([]int{20}, 0),
			job.NewCarrying([]int{30}, resource.Rock, 0, 0, 0),
			job.NewLibraryRead([]int{11}, 1),
			job.NewDigging([]int{21}, 0),
			job.NewLibraryRead([]int{12}, 1),
			job.NewCarrying([]int{31}, resource.Rock, 0, 0, 0),
			job.NewDigging([]int{22}, 0),
			job.NewCarrying([]int{32}, resource.Rock, 0, 0, 0),
		},
	}
	js.sortJobPriorities()
	order := []string{
		"Digging",
		"Digging",
		"Digging",
		"Carrying",
		"Carrying",
		"Carrying",
		"Library",
		"Library",
		"Library",
	}
	for i, ord := range order {
		if js.Jobs[i].String() != ord {
			t.Fatalf("wanted %v got %v", ord, js.Jobs[i].String())
		}
	}
}
