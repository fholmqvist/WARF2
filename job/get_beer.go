package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type GetBeer struct {
	Bar             *room.Bar
	StorageTile     *room.StorageTile
	BeerRefillIndex int
	destinations    []int
	dwarf           *dwarf.Dwarf
	path            []int
	amount          uint
	remove          bool
}

func NewGetBeer(bar *room.Bar, st *room.StorageTile, refillIdx int, destinations []int) *GetBeer {
	return &GetBeer{
		Bar:             bar,
		StorageTile:     st,
		BeerRefillIndex: refillIdx,
		destinations:    destinations,
	}
}

func (g *GetBeer) PerformWork(mp *m.Map, dwf []*dwarf.Dwarf, rs *room.Service) bool {
	if g.path == nil {
		g.amount = g.StorageTile.TakeAll()
		g.setupPath(mp)
		return unfinished
	}
	if len(g.path) == 0 {
		g.Bar.AddBeer(g.amount)
		g.path = nil
		g.destinations = []int{}
		g.remove = true
		return finished
	}
	g.dwarf.Idx = g.path[0]
	g.destinations[0] = g.path[0]
	g.path = g.path[1:]
	return unfinished
}

func (g *GetBeer) setupPath(mp *m.Map) {
	path, ok := m.CreatePath(
		&mp.Tiles[g.dwarf.Idx],
		&mp.Tiles[g.BeerRefillIndex],
	)
	if !ok {
		return
	}
	if len(path) == 0 {
		return
	}
	mp.Items[g.dwarf.Idx].Sprite = entity.NoItem
	g.path = path
}

func (g *GetBeer) Finish(mp *m.Map, rs *room.Service) {
	if g.dwarf == nil {
		return
	}
	g.dwarf.SetToAvailable()
	g.dwarf = nil
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
	return true
}

func (g *GetBeer) String() string {
	return "GetBeer"
}
