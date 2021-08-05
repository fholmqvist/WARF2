package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type GetBeer struct {
	StorageTile     *room.StorageTile
	BeerRefillIndex int
	destinations    []int
	dwarf           *dwarf.Dwarf
	remove          bool
}

func NewGetBeer(st *room.StorageTile, refillIdx int, destinations []int) *GetBeer {
	return &GetBeer{
		StorageTile:     st,
		BeerRefillIndex: refillIdx,
		destinations:    destinations,
	}
}

func (g *GetBeer) PerformWork(*m.Map, []*dwarf.Dwarf, *room.Service) bool {
	return unfinished
}

func (g *GetBeer) Finish(*m.Map, *room.Service) {

}

func (g *GetBeer) Remove() bool {
	return g.remove
}

func (g *GetBeer) GetWorker() *dwarf.Dwarf {
	return g.dwarf
}

func (g *GetBeer) SetWorker(d *dwarf.Dwarf) {
	g.dwarf = d
}

func (g *GetBeer) GetDestinations() []int {
	return g.destinations
}

func (g *GetBeer) HasInternalMove() bool {
	return false
}

func (g *GetBeer) String() string {
	return "GetBeer"
}
