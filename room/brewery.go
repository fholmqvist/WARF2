package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
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
	for i, t := range tiles {
		if i%2 == 0 {
			continue
		}
		if m.IsAnyWall(mp.OneTileLeft(t.Idx).Sprite) ||
			m.IsAnyWall(mp.OneTileRight(t.Idx).Sprite) ||
			m.IsAnyWall(mp.OneTileUp(t.Idx).Sprite) ||
			m.IsAnyWall(mp.OneTileUpLeft(t.Idx).Sprite) ||
			m.IsAnyWall(mp.OneTileUpRight(t.Idx).Sprite) ||
			m.IsAnyWall(mp.OneTileDown(t.Idx).Sprite) ||
			m.IsAnyWall(mp.OneTileDownLeft(t.Idx).Sprite) ||
			m.IsAnyWall(mp.OneTileDownRight(t.Idx).Sprite) {
			continue
		}
		mp.Items[t.Idx].Resource = entity.ResourceNone
		mp.Items[t.Idx].Sprite = entity.EmptyBarrel
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
