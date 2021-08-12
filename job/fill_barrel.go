package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type FillBarrel struct {
	JobBase
	StorageTile *room.StorageTile
	WheatIndex  int
	BarrelIndex int
	path        []int
	amount      uint
}

func NewFillBrewer(st *room.StorageTile, barrelIdx int, destinations []int) *FillBarrel {
	return &FillBarrel{
		JobBase:     NewJobBase(destinations),
		StorageTile: st,
		WheatIndex:  st.Idx,
		BarrelIndex: barrelIdx,
	}
}

func (f *FillBarrel) PerformWork(mp *m.Map, d []*dwarf.Dwarf, rs *room.Service) bool {
	// Just arrived.
	if f.path == nil {
		f.amount = f.StorageTile.TakeAll()
		f.setupPath(mp)
		return unfinished
	}
	// Finished.
	if len(f.path) == 0 {
		mp.Items[f.BarrelIndex].Sprite = entity.FilledBarrel
		mp.Items[f.BarrelIndex].ResourceAmount = f.amount
		f.dwarf.Idx = f.BarrelIndex
		f.remove = true
		return finished
	}
	// Move.
	f.dwarf.Idx = f.path[0]
	f.path = f.path[1:]
	return unfinished
}

func (f *FillBarrel) HasInternalMove() bool {
	return true
}

func (f *FillBarrel) String() string {
	return "FillBarrel"
}

func (f *FillBarrel) setupPath(mp *m.Map) {
	path := []int{}
	for _, dst := range m.NeighTileFour(f.BarrelIndex) {
		p, ok := m.CreatePath(
			&mp.Tiles[f.dwarf.Idx],
			&mp.Tiles[dst],
		)
		if !ok {
			continue
		}
		path = p
		break
	}
	if len(path) == 0 {
		return
	}
	f.path = path
}
