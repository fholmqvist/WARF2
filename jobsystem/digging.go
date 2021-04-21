package jobsystem

import (
	m "projects/games/warf2/worldmap"
)

type Digging struct {
	worker      *Worker
	state       JobState
	destination int
	wallIdx     int
}

// Recalculates (!) and return state.
func (d *Digging) CheckState() JobState {
	return d.state
}

// Checks if the tile of to-be-dug wall is still selected.
func (d *Digging) NeedsToBeRemoved(mp *m.Map) bool {
	return !m.IsSelectedWall(mp.Tiles[d.wallIdx].Sprite)
}

// Ran on arrival.
func (d *Digging) PerformWork(mp *m.Map) func() bool {
	return func() bool {
		t := &mp.Tiles[d.wallIdx]
		if !m.IsSelectedWall(t.Sprite) {
			// Job is, in a sense, done.
			return true
		}
		t.Sprite = m.Ground
		for _, nb := range m.SurroundingTilesFour(t.Idx) {
			mp.FixWall(&mp.Tiles[nb.Idx])
		}
		return true
	}
}

func (d *Digging) GetDestination() int {
	return d.destination
}

func (d *Digging) Priority() int {
	return 0
}

func (d *Digging) GetWorker() *Worker {
	return d.worker
}

func (d *Digging) SetWorker(w *Worker) {
	d.worker = w
}

func (d *Digging) GetState() JobState {
	return d.state
}

func (d *Digging) SetState(j JobState) {
	d.state = j
}
