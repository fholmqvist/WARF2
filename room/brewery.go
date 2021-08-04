package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var breweryAutoID = 0

type Brewery struct {
	ID      int
	barrels []int
	tiles   []int
}

func NewBrewery(mp *m.Map, x, y int) *Brewery {
	b := &Brewery{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.BreweryFloor })
	if len(tiles) == 0 {
		return nil
	}
	sort.Ints(tiles)
	for _, idx := range tiles {
		tile := &mp.Tiles[idx]
		tile.Room = b
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
		if entity.IsBarrel(mp.Items[m.OneUp(idx)].Sprite) &&
			entity.IsBarrel(mp.Items[m.OneUp(m.OneUp(idx))].Sprite) {
			continue
		}
		mp.Items[idx].ResourceAmount = 0
		mp.Items[idx].Resource = entity.ResourceNone
		mp.Items[idx].Sprite = entity.EmptyBarrel
		b.barrels = append(b.barrels, idx)
	}
	b.tiles = tiles
	b.ID = breweryAutoID
	breweryAutoID++
	return b
}

func (b *Brewery) Update(*m.Map) {}

func (b *Brewery) GetID() int {
	if b == nil {
		return -1
	}
	return b.ID
}

func (b *Brewery) String() string {
	return "Brewery"
}

func (b *Brewery) Tiles() []int {
	return b.tiles
}

func (b *Brewery) GetEmptyBarrel(mp *m.Map) (int, bool) {
	for _, idx := range b.barrels {
		if !entity.IsEmptyBarrel(mp.Items[idx].Sprite) {
			continue
		}
		return idx, true
	}
	return -1, false
}
