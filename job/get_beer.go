package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type GetBeer struct {
	JobBase
	Bar             *room.Bar
	StorageTile     *room.StorageTile
	BeerRefillIndex int
	path            []int
	amount          uint
}

func NewGetBeer(bar *room.Bar, st *room.StorageTile, refillIdx int, destinations []int) *GetBeer {
	return &GetBeer{
		JobBase:         NewJobBase(destinations),
		Bar:             bar,
		StorageTile:     st,
		BeerRefillIndex: refillIdx,
	}
}

func (g *GetBeer) PerformWork(mp *m.Map, dwf []*dwarf.Dwarf, rs *room.Service) bool {
	// Just arrived, pick up beer.
	if g.path == nil {
		g.amount = g.StorageTile.TakeAll()
		g.setupPath(mp)
		return unfinished
	}
	// Finished.
	if len(g.path) == 0 {
		g.Bar.AddBeer(g.amount)
		g.remove = true
		return finished
	}
	// Move.
	g.dwarf.Idx = g.path[0]
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

func (g *GetBeer) HasInternalMove() bool {
	return false
}

func (g *GetBeer) String() string {
	return "GetBeer"
}
