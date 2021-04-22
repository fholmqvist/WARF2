package jobsystem

import (
	m "projects/games/warf2/worldmap"
)

func WaitingForWorker(j Job) bool {
	return j.GetWorker() == nil
}

// Sets worker for digging, returns whether setting was successful.
// On success, worker proceeds to move to destination.
func SetWorkerAndMove(j Job, w Worker, mp *m.Map) bool {
	if !WaitingForWorker(j) {
		return false
	}
	ok := w.MoveTo(j.GetDestination(), mp)
	if !ok {
		return false
	}
	j.SetWorker(&w)
	w.SetJob(j)
	return true
}
