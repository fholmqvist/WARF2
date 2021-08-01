package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type FillBrewer struct {
	WheatIndex  int
	BarrelIndex int
	dwarf       *dwarf.Dwarf
}

func NewFillBrewer(wheatIdx, barrelIdx int) *FillBrewer {
	return &FillBrewer{
		WheatIndex:  wheatIdx,
		BarrelIndex: barrelIdx,
	}
}

func (f *FillBrewer) NeedsToBeRemoved(mp *m.Map, rs *room.Service) bool {
	return false
}

func (f *FillBrewer) PerformWork(mp *m.Map, d []*dwarf.Dwarf, rs *room.Service) bool {
	return false
}

func (f *FillBrewer) Finish(mp *m.Map, rs *room.Service) {
}

func (f *FillBrewer) GetWorker() *dwarf.Dwarf {
	return f.dwarf
}

func (f *FillBrewer) SetWorker(d *dwarf.Dwarf) {
	f.dwarf = d
}

func (f *FillBrewer) GetDestinations() []int {
	return []int{f.WheatIndex}
}

func (f *FillBrewer) HasInternalMove() bool {
	return false
}

func (f *FillBrewer) String() string {
	return "FillBrewer"
}
