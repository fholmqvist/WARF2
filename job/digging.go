package job

import (
	"projects/games/warf2/dwarf"
	m "projects/games/warf2/worldmap"
)

type Digging struct {
	dwarf       *dwarf.Dwarf
	destination int
	wallIdx     int
}

func NewDigging(destination, wallIdx int) *Digging {
	return &Digging{nil, destination, wallIdx}
}

// Checks if the tile of to-be-dug wall is still selected.
func (d *Digging) NeedsToBeRemoved(mp *m.Map) bool {
	return !m.IsSelectedWall(mp.Tiles[d.wallIdx].Sprite)
}

// Ran on arrival.
func (d *Digging) PerformWork(mp *m.Map) bool {
	t := &mp.Tiles[d.wallIdx]
	if !m.IsSelectedWall(t.Sprite) {
		// Job is, in a sense, done.
		return finished
	}
	t.Sprite = m.Ground
	for _, nb := range m.SurroundingTilesFour(t.Idx) {
		mp.FixWall(&mp.Tiles[nb.Idx])
	}
	return finished
}

func (d *Digging) Priority() int {
	return 1
}

func (d *Digging) GetWorker() *dwarf.Dwarf {
	return d.dwarf
}

func (d *Digging) SetWorker(dw *dwarf.Dwarf) {
	d.dwarf = dw
}

func (d *Digging) GetDestination() int {
	return d.destination
}

func (d *Digging) GetWallIdx() int {
	return d.wallIdx
}
