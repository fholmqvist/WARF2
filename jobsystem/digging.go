package jobsystem

import (
	m "projects/games/warf2/worldmap"
)

type Digging struct {
	worker      *Worker
	destination int
	wallIdx     int
}

func NewDigging(w *Worker, destination, wallIdx int) *Digging {
	return &Digging{w, destination, wallIdx}
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

func (d *Digging) Priority() int {
	return 1
}

func (d *Digging) GetWorker() *Worker {
	return d.worker
}

func (d *Digging) SetWorker(w *Worker) {
	d.worker = w
}

func (d *Digging) GetDestination() int {
	return d.destination
}
