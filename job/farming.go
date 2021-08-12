package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Farming struct {
	JobBase
	FarmID int
	path   []int
}

func NewFarming(farmID int, destinations []int) *Farming {
	return &Farming{
		NewJobBase(destinations),
		farmID,
		nil,
	}
}

func (f *Farming) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	// No farm, abort.
	if _, ok := rs.GetFarm(f.FarmID); !ok {
		f.remove = true
		return finished
	}
	// Finished.
	if len(f.destinations) == 0 {
		f.remove = true
		return finished
	}
	// Move.
	return f.moveDwarf(mp)
}

func (f *Farming) HasInternalMove() bool {
	return true
}

func (f *Farming) String() string {
	return "Farming"
}

func (f *Farming) moveDwarf(mp *m.Map) bool {
	// Harvest.
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
	// Move to next tile.
	if f.path != nil {
		f.moveAlongPath()
		return unfinished
	}
	// Find next tile.
	nextIdx := getNextIdx(f.destinations)
	if nextIdx-f.dwarf.Idx == 1 {
		f.dwarf.Idx = nextIdx // Adjacent
	} else {
		path, ok := getPath(mp, nextIdx, f.dwarf.Idx) // Elsewhere
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
