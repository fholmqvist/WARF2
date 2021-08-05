package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var (
	breweryAutoID = 0
	brewDone      = uint(5)
)

type Brewery struct {
	ID      int
	barrels []*brewingBarrel
	tiles   []int
}

type brewingBarrel struct {
	idx int
	val uint
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
		b.barrels = append(b.barrels, &brewingBarrel{idx, 0})
	}
	b.tiles = tiles
	b.ID = breweryAutoID
	breweryAutoID++
	return b
}

func (b *Brewery) Update(mp *m.Map) {
	for _, barrel := range b.barrels {
		if !entity.IsFilledBarrel(mp.Items[barrel.idx].Sprite) {
			barrel.val = 0
			continue
		}
		barrel.val++
		if barrel.val < brewDone {
			continue
		}
		// This opens up for
		// checkForCarryingJob
		// on this tile.
		mp.Items[barrel.idx].Resource = entity.ResourceBeer
	}
}

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
	for _, barrel := range b.barrels {
		if !entity.IsEmptyBarrel(mp.Items[barrel.idx].Sprite) {
			continue
		}
		return barrel.idx, true
	}
	return -1, false
}
