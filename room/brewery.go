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
	sort.Ints(tiles)
	for i, idx := range tiles {
		if i%2 == 0 {
			continue
		}
		if m.IsAnyWall(mp.OneTileLeft(idx).Sprite) ||
			m.IsAnyWall(mp.OneTileRight(idx).Sprite) ||
			m.IsAnyWall(mp.OneTileUp(idx).Sprite) ||
			m.IsAnyWall(mp.OneTileUpLeft(idx).Sprite) ||
			m.IsAnyWall(mp.OneTileUpRight(idx).Sprite) ||
			m.IsAnyWall(mp.OneTileDown(idx).Sprite) ||
			m.IsAnyWall(mp.OneTileDownLeft(idx).Sprite) ||
			m.IsAnyWall(mp.OneTileDownRight(idx).Sprite) {
			continue
		}
		mp.Items[idx].Resource = entity.ResourceNone
		mp.Items[idx].Sprite = entity.EmptyBarrel
		mp.Tiles[idx].Room = b
	}
	b.tiles = tiles
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
