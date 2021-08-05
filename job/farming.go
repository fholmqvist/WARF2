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
	remove       bool
}

func NewFarming(farmID int, destinations []int) *Farming {
	return &Farming{farmID, nil, destinations, nil, false}
}

func (f *Farming) Remove() bool {
	return f.remove
}

func (f *Farming) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	if f.remove {
		return finished
	}
	if _, ok := rs.GetFarm(f.FarmID); !ok {
		f.destinations = []int{}
		f.path = nil
		f.remove = true
		f.dwarf.SetToAvailable()
		f.dwarf = nil
		return finished
	}
	if len(f.destinations) == 0 {
		f.remove = true
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
		mp.Items[currentIdx].ResourceAmount = 1
		f.destinations = f.destinations[:len(f.destinations)-1]
		if len(f.destinations) == 0 {
			f.remove = true
			return finished
		}
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
