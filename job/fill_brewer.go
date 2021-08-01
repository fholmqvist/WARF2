package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type FillBrewer struct {
	StorageTile  *room.StorageTile
	WheatIndex   int
	BarrelIndex  int
	destinations []int
	dwarf        *dwarf.Dwarf
	path         []int
	amount       uint
}

func NewFillBrewer(st *room.StorageTile, barrelIdx int, destinations []int) *FillBrewer {
	return &FillBrewer{
		StorageTile:  st,
		WheatIndex:   st.Idx,
		BarrelIndex:  barrelIdx,
		destinations: destinations,
	}
}

func (f *FillBrewer) NeedsToBeRemoved(mp *m.Map, rs *room.Service) bool {
	return f.path == nil && len(f.destinations) == 0
}

func (f *FillBrewer) PerformWork(mp *m.Map, d []*dwarf.Dwarf, rs *room.Service) bool {
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

func (f *FillBrewer) Finish(mp *m.Map, rs *room.Service) {
	if f.dwarf == nil {
		return
	}
	mp.Items[f.BarrelIndex].Sprite = entity.FilledBarrel
	mp.Items[f.BarrelIndex].ResourceAmount = f.amount
	f.dwarf.SetToAvailable()
	f.dwarf = nil
}

func (f *FillBrewer) GetWorker() *dwarf.Dwarf {
	return f.dwarf
}

func (f *FillBrewer) SetWorker(d *dwarf.Dwarf) {
	f.dwarf = d
}

func (f *FillBrewer) GetDestinations() []int {
	return f.destinations
}

func (f *FillBrewer) HasInternalMove() bool {
	return false
}

func (f *FillBrewer) String() string {
	return "FillBrewer"
}

func (f *FillBrewer) setupPath(mp *m.Map) {
	path := []int{}
	for _, dst := range f.destinations {
		p, ok := f.dwarf.CreatePath(
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
