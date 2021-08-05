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
}

func NewFillBrewer(st *room.StorageTile, barrelIdx int, destinations []int) *FillBarrel {
	return &FillBarrel{
		StorageTile:  st,
		WheatIndex:   st.Idx,
		BarrelIndex:  barrelIdx,
		destinations: destinations,
	}
}

func (f *FillBarrel) NeedsToBeRemoved(mp *m.Map, rs *room.Service) bool {
	return f.path == nil && len(f.destinations) == 0
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
		return finished
	}
	f.dwarf.Idx = f.path[0]
	f.destinations[0] = f.path[0]
	f.path = f.path[1:]
	return unfinished
}

func (f *FillBarrel) Finish(mp *m.Map, rs *room.Service) {
	if f.dwarf == nil {
		return
	}
	mp.Items[f.BarrelIndex].Sprite = entity.FilledBarrel
	mp.Items[f.BarrelIndex].ResourceAmount = f.amount
	f.dwarf.SetToAvailable()
	f.dwarf = nil
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
	return false
}

func (f *FillBarrel) String() string {
	return "FillBarrel"
}

func (f *FillBarrel) setupPath(mp *m.Map) {
	path := []int{}
	for _, dst := range f.destinations {
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
