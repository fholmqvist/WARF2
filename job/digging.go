package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/item"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Digging struct {
	JobBase
	wallIdx int
}

func NewDigging(destinations []int, wallIdx int) *Digging {
	return &Digging{
		JobBase: NewJobBase(destinations),
		wallIdx: wallIdx,
	}
}

func (d *Digging) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	t := &mp.Tiles[d.wallIdx]
	// Job is, in a sense, done.
	if !m.IsSelectedWall(t.Sprite) {
		d.remove = true
		return finished
	}
	// Finished.
	t.Sprite = m.Ground
	mp.Items[t.Idx].Sprite = item.RandomCrumbledWall()
	mp.Items[t.Idx].Resource = entity.ResourceRock
	mp.Items[t.Idx].ResourceAmount = 1
	for _, nb := range m.NeighTileDirFour(t.Idx) {
		mp.FixWall(&mp.Tiles[nb.Idx])
	}
	d.remove = true
	return finished
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
