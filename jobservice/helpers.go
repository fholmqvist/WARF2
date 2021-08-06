package jobservice

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/job"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func HasWorker(j job.Job) bool {
	return j.GetWorker() != nil
}

func WaitingForWorker(j job.Job) bool {
	return !HasWorker(j)
}

// Sets worker for digging, returns whether setting was successful.
// On success, worker proceeds to move to destination.
func SetWorkerAndMove(j job.Job, w *dwarf.Dwarf, mp *m.Map) bool {
	if HasWorker(j) {
		return false
	}
	var foundDestination bool
	for _, destination := range j.GetDestinations() {
		if foundDestination {
			break
		}
		ok := w.MoveTo(destination, mp)
		if !ok {
			continue
		}
		foundDestination = true
	}
	if !foundDestination {
		return false
	}
	j.SetWorker(w)
	w.SetJob()
	return true
}
