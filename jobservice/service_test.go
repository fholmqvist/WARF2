package jobservice

import (
	"testing"

	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/job"
)

func TestWorkQueue(t *testing.T) {
	js := &Service{
		Jobs: []job.Job{
			job.NewRead([]int{10}, 1),
			job.NewDigging([]int{20}, 0),
			job.NewCarrying([]int{30}, entity.ResourceRock, 0, 0, 0),
			job.NewRead([]int{11}, 1),
			job.NewDigging([]int{21}, 0),
			job.NewRead([]int{12}, 1),
			job.NewCarrying([]int{31}, entity.ResourceRock, 0, 0, 0),
			job.NewDigging([]int{22}, 0),
			job.NewCarrying([]int{32}, entity.ResourceRock, 0, 0, 0),
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
