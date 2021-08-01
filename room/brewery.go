package room

import (
	"sort"

	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var breweryAutoID = 0

type Brewery struct {
	ID    int
	tiles []int
}

func NewBrewery(mp *m.Map, x, y int) *Brewery {
	b := &Brewery{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.BreweryFloor })
	if len(tiles) == 0 {
		return nil
	}
	sort.Sort(tiles)
	for _, t := range tiles {
		mp.Tiles[t.Idx].Room = b
	}
	b.tiles = tiles.ToIdxs()
	b.ID = breweryAutoID
	breweryAutoID++
	return b
}

func (b *Brewery) Update(*m.Map) {

}

func (b *Brewery) GetID() int {
	return b.ID
}

func (b *Brewery) String() string {
	return "Brewery"
}

func (b *Brewery) Tiles() []int {
	return b.tiles
}
