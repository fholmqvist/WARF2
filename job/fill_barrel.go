package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type FillBarrel struct {
	StorageTile  *room.StorageTile
	WheatIndex   int
	BarrelIndex  int
	destinations []int
	dwarf        *dwarf.Dwarf
	path         []int
	amount       uint
	remove       bool
}

func NewFillBrewer(st *room.StorageTile, barrelIdx int, destinations []int) *FillBarrel {
	return &FillBarrel{
		StorageTile:  st,
		WheatIndex:   st.Idx,
		BarrelIndex:  barrelIdx,
		destinations: destinations,
	}
}

func (f *FillBarrel) Remove() bool {
	return f.remove
}

func (f *FillBarrel) PerformWork(mp *m.Map, d []*dwarf.Dwarf, rs *room.Service) bool {
	if f.path == nil {
		f.amount = f.StorageTile.TakeAll()
		f.setupPath(mp)
		return unfinished
	}
	if len(f.path) == 0 {
		f.dwarf.Idx = f.BarrelIndex
		f.path = nil
		f.destinations = []int{}
		f.remove = true
		mp.Items[f.BarrelIndex].Sprite = entity.FilledBarrel
		mp.Items[f.BarrelIndex].ResourceAmount = f.amount
		f.dwarf.SetToAvailable()
		f.dwarf = nil
		return finished
	}
	f.dwarf.Idx = f.path[0]
	f.destinations[0] = f.path[0]
	f.path = f.path[1:]
	return unfinished
}

func (f *FillBarrel) GetWorker() *dwarf.Dwarf {
	return f.dwarf
}

func (f *FillBarrel) SetWorker(d *dwarf.Dwarf) {
	f.dwarf = d
}

func (f *FillBarrel) GetDestinations() []int {
	return f.destinations
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
	mp.Items[f.dwarf.Idx].Sprite = entity.NoItem
	f.path = path
}
