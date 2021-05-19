package jobservice

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/job"
	m "projects/games/warf2/worldmap"
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
