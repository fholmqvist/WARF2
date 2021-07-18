package jobservice

import (
	"fmt"
	"math/rand"

	"github.com/Holmqvist1990/WARF2/job"
)

func (jb *Service) Len() int {
	return len(jb.Jobs)
}

func (jb *Service) Less(i, j int) bool {
	fst := jb.GetPriority(jb.Jobs[i])
	snd := jb.GetPriority(jb.Jobs[j])
	// Randomize equally prioritized.
	if fst == snd {
		return rand.Intn(2) == 1
	}
	// Highest priority first.
	return fst > snd
}

func (jb *Service) Swap(i, j int) {
	jb.Jobs[i], jb.Jobs[j] = jb.Jobs[j], jb.Jobs[i]
}

// Priority is in ascending order.
func (jb *Service) GetPriority(j job.Job) int {
	switch j.(type) {
	case *job.Digging:
		return 5
	case *job.Carrying:
		return 3
	case *job.LibraryRead:
		return 1
	case *job.Farming:
		return 3
	case *job.PlantFarm:
		return 2
	default:
		panic(fmt.Sprint("missing job type:", j))
	}
}
