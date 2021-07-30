package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Farming struct {
	FarmID       int
	dwarf        *dwarf.Dwarf
	destinations []int
	path         []int
}

func NewFarming(farmID int, destinations []int) *Farming {
	return &Farming{farmID, nil, destinations, nil}
}

func (f *Farming) NeedsToBeRemoved(mp *m.Map, r *room.Service) bool {
	return len(f.destinations) == 0 && f.path == nil
}

func (d *Farming) Finish(*m.Map, *room.Service) {
	if d.dwarf == nil {
		return
	}
	d.dwarf.SetToAvailable()
	d.dwarf = nil
}

func (f *Farming) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	if _, ok := rs.GetFarm(f.FarmID); !ok {
		f.destinations = []int{}
		f.path = nil
		return finished
	}
	if len(f.destinations) == 0 {
		return finished
	}
	if f.dwarf == nil {
		return unfinished
	}
	return f.moveDwarf(mp)
}

func (f *Farming) GetWorker() *dwarf.Dwarf {
	return f.dwarf
}

func (f *Farming) SetWorker(dw *dwarf.Dwarf) {
	f.dwarf = dw
}

func (f *Farming) GetDestinations() []int {
	return f.destinations
}

func (f *Farming) HasInternalMove() bool {
	return true
}

func (f *Farming) String() string {
	return "Farming"
}

func (f *Farming) moveDwarf(mp *m.Map) bool {
	currentIdx := getNextIdx(f.destinations)
	if f.dwarf.Idx == currentIdx {
		mp.Items[currentIdx].Sprite = entity.Wheat
		mp.Items[currentIdx].Resource = entity.ResourceWheat
		f.destinations = f.destinations[:len(f.destinations)-1]
	}
	if f.NeedsToBeRemoved(mp, nil) {
		return finished
	}
	if f.path != nil {
		f.moveAlongPath()
		return unfinished
	}
	nextIdx := getNextIdx(f.destinations)
	if nextIdx-f.dwarf.Idx == 1 {
		f.dwarf.Idx = nextIdx // Adjacent
	} else {
		path, ok := getPath(mp, nextIdx, f.dwarf) // Elsewhere
		if !ok {
			return unfinished
		}
		f.path = path
	}
	return unfinished
}

func (f *Farming) moveAlongPath() {
	if len(f.path) == 0 {
		f.path = nil
		return
	}
	// Move indexes to current path index.
	f.dwarf.Idx = f.path[0]
	// Iterate path.
	f.path = f.path[1:]
}
