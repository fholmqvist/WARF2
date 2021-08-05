package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type PlantFarm struct {
	Farm         *room.Farm
	destinations []int
	dwarf        *dwarf.Dwarf
	path         []int
	remove       bool
}

func NewPlantFarm(farm *room.Farm, farmableDestinations []int) *PlantFarm {
	return &PlantFarm{farm, farmableDestinations, nil, nil, false}
}

func (p *PlantFarm) Remove() bool {
	return p.remove
}

func (p *PlantFarm) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	if p.remove {
		return finished
	}
	if _, ok := rs.GetFarm(p.Farm.ID); !ok {
		p.Farm = nil
		p.remove = true
		return finished
	}
	if len(p.destinations) == 0 {
		p.remove = true
		p.dwarf.SetToAvailable()
		p.dwarf = nil
		return finished
	}
	if p.dwarf == nil {
		return unfinished
	}
	return p.moveDwarf(mp)
}

func (p *PlantFarm) GetWorker() *dwarf.Dwarf {
	return p.dwarf
}

func (p *PlantFarm) SetWorker(d *dwarf.Dwarf) {
	p.dwarf = d
}

func (p *PlantFarm) GetDestinations() []int {
	return p.destinations
}

func (p *PlantFarm) HasInternalMove() bool {
	return true
}

func (p *PlantFarm) String() string {
	return "PlantFarm"
}

func (p *PlantFarm) moveDwarf(mp *m.Map) bool {
	currentIdx := getNextIdx(p.destinations)
	if p.dwarf.Idx == currentIdx {
		p.Farm.PlantFarm(mp, mp.Items[currentIdx])
		p.destinations = p.destinations[:len(p.destinations)-1]
		p.path = nil
		if len(p.destinations) == 0 {
			p.remove = true
			return finished
		}
	}
	if p.path != nil {
		p.moveAlongPath()
		return unfinished
	}
	nextIdx := getNextIdx(p.destinations)
	if nextIdx-p.dwarf.Idx == 1 {
		p.dwarf.Idx = nextIdx // Adjacent
	} else {
		path, ok := getPath(mp, nextIdx, p.dwarf) // Elsewhere
		if !ok {
			return unfinished
		}
		p.path = path
	}
	return unfinished
}

func (p *PlantFarm) moveAlongPath() {
	if len(p.path) == 0 {
		p.path = nil
		return
	}
	// Move indexes to current path index.
	p.dwarf.Idx = p.path[0]
	// Iterate path.
	p.path = p.path[1:]
}
