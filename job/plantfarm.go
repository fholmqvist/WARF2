package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

////////////////
// TODO
// Does nothing.
////////////////

type PlantFarm struct {
	FarmID int
	dwarf  *dwarf.Dwarf
}

func NewPlantFarm(farmId int) *PlantFarm {
	return &PlantFarm{farmId, nil}
}

func (p *PlantFarm) NeedsToBeRemoved(*m.Map) bool {
	return true
}

func (p *PlantFarm) PerformWork(*m.Map, []*dwarf.Dwarf) bool {
	return true
}

func (p *PlantFarm) Finish(*m.Map, *room.Service) {
	return
}

func (p *PlantFarm) GetWorker() *dwarf.Dwarf {
	return p.dwarf
}

func (p *PlantFarm) SetWorker(d *dwarf.Dwarf) {
	p.dwarf = d
}

func (p *PlantFarm) GetDestinations() []int {
	return []int{}
}

func (p *PlantFarm) String() string {
	return "PlantFarm"
}
