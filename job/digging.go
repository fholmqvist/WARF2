package job

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/item"
	m "projects/games/warf2/worldmap"
)

type Digging struct {
	dwarf        *dwarf.Dwarf
	destinations []int
	wallIdx      int
}

func NewDigging(destinations []int, wallIdx int) *Digging {
	return &Digging{nil, destinations, wallIdx}
}

// Checks if the tile of to-be-dug wall is still selected.
func (d *Digging) NeedsToBeRemoved(mp *m.Map) bool {
	return !m.IsSelectedWall(mp.Tiles[d.wallIdx].Sprite) || d.dwarf == nil
}

func (d *Digging) Reset(*m.Map) {
	if d.dwarf == nil {
		return
	}
	d.dwarf.SetToAvailable()
	d.dwarf = nil
}

// Ran on arrival.
func (d *Digging) PerformWork(mp *m.Map) bool {
	t := &mp.Tiles[d.wallIdx]
	if !m.IsSelectedWall(t.Sprite) {
		// Job is, in a sense, done.
		return finished
	}
	t.Sprite = m.None
	mp.Items[t.Idx].Sprite = item.RandomCrumbledWall()
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

func (d *Digging) GetDestinations() []int {
	return d.destinations
}

func (d *Digging) GetWallIdx() int {
	return d.wallIdx
}
