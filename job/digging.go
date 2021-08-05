package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/item"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
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
func (d *Digging) NeedsToBeRemoved(mp *m.Map, r *room.Service) bool {
	return !m.IsSelectedWall(mp.Tiles[d.wallIdx].Sprite) || d.dwarf == nil
}

func (d *Digging) Finish(*m.Map, *room.Service) {
	if d.dwarf == nil {
		return
	}
	d.dwarf.SetToAvailable()
	d.dwarf = nil
}

func (d *Digging) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	t := &mp.Tiles[d.wallIdx]
	if !m.IsSelectedWall(t.Sprite) {
		// Job is, in a sense, done.
		return finished
	}
	t.Sprite = m.Ground
	mp.Items[t.Idx].Sprite = item.RandomCrumbledWall()
	mp.Items[t.Idx].Resource = entity.ResourceRock
	mp.Items[t.Idx].ResourceAmount = 1
	for _, nb := range m.NeighTileDirFour(t.Idx) {
		mp.FixWall(&mp.Tiles[nb.Idx])
	}
	return finished
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

func (d *Digging) HasInternalMove() bool {
	return false
}

func (d *Digging) GetWallIdx() int {
	return d.wallIdx
}

func (d *Digging) String() string {
	return "Digging"
}
