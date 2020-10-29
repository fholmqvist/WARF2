package jobsystem

import (
	"projects/games/warf2/worldmap"
)

// Digging defines the job
// for digging walls.
type Digging struct {
	worker      *Worker
	state       JobState
	destination int
	wallIdx     int
}

// WaitingForWorker returns
// whether worker is missing.
func (d *Digging) WaitingForWorker() bool {
	return d.worker == nil && d.state == New
}

// SetWorkerAndMove sets worker for digging,
// and returns a bool whether setting was successful.
// On success, worker proceeds to move to destination.
func (d *Digging) SetWorkerAndMove(worker Worker, mp *worldmap.Map) bool {
	if !d.WaitingForWorker() {
		return false
	}

	ok := worker.MoveTo(d.destination, mp)
	if !ok {
		return false
	}

	worker.SetAvailable(false)
	d.worker = &worker

	worker.SetJob(d)

	return true
}

// CheckState sets and returns
// current jobstate.
func (d *Digging) CheckState() JobState {
	return d.state
}

// NeedsToBeRemoved checks if the
// tile of to-be-dug wall is still selected.
func (d *Digging) NeedsToBeRemoved(mp *worldmap.Map) bool {
	return !worldmap.IsSelectedWall(mp.Tiles[d.wallIdx].Sprite)
}

// PerformWork is the function to get
// called when worker arrives at destination.
func (d *Digging) PerformWork(mp *worldmap.Map) func() {
	return func() {
		mp.Tiles[d.wallIdx].Sprite = worldmap.Ground
	}
}

// GetDestination returns the
// destination for the worker.
func (d *Digging) GetDestination() int {
	return d.destination
}
