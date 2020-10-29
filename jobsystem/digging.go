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

// SetWorker sets worker for digging,
// and returns a bool whether setting
// was successful.
func (d *Digging) SetWorker(worker Worker) bool {
	if !d.WaitingForWorker() {
		return false
	}

	ok := worker.SetAvailable(false)
	if !ok {
		return false
	}

	d.worker = &worker
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
