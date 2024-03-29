package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/item"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var farmAutoID = 0

type Farm struct {
	ID           int     // First tile.
	AllTileIdxs  []int   // To be indexed against Worldmap.
	FarmableIdxs []int   // Only tiles with farms on them.
	farmTile     *m.Tile // Knows when farm has reached maturity.
}

func NewFarm(mp *m.Map, x, y int) (*Farm, bool) {
	f := &Farm{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.Ground })
	if len(tiles) == 0 {
		return nil, false
	}
	sort.Ints(tiles)
	for _, idx := range tiles {
		mp.Tiles[idx].Room = f
		f.PlantFarm(mp, mp.Tiles[idx])
		if f.farmTile != nil {
			continue
		}
		if !entity.IsFarm(mp.Items[idx].Sprite) {
			continue
		}
		f.farmTile = &mp.Items[idx]
	}
	f.AllTileIdxs = tiles
	f.FarmableIdxs = f.farmableIndexes(mp)
	f.ID = farmAutoID
	farmAutoID++
	return f, true
}

func (f *Farm) GetID() int {
	if f == nil {
		return -1
	}
	return f.ID
}

func (f *Farm) String() string {
	return "Farm"
}

func (f *Farm) Tiles() []int {
	return f.AllTileIdxs
}

func (f *Farm) Update(mp *m.Map) {
	for _, tIdx := range f.FarmableIdxs {
		tile := &mp.Items[tIdx]
		if tile.Sprite == 0 {
			continue
		}
		switch tile.Sprite {
		case entity.FarmSingleEmpty:
			tile.Sprite = entity.FarmSingleWheat1
		case entity.FarmLeftEmpty:
			tile.Sprite = entity.FarmLeftWheat1
		case entity.FarmMiddleEmpty:
			tile.Sprite = entity.FarmMiddleWheat1
		case entity.FarmRightEmpty:
			tile.Sprite = entity.FarmRightWheat1

		case entity.FarmSingleWheat1:
			tile.Sprite = entity.FarmSingleWheat2
		case entity.FarmLeftWheat1:
			tile.Sprite = entity.FarmLeftWheat2
		case entity.FarmMiddleWheat1:
			tile.Sprite = entity.FarmMiddleWheat2
		case entity.FarmRightWheat1:
			tile.Sprite = entity.FarmRightWheat2

		case entity.FarmSingleWheat2:
			tile.Sprite = entity.FarmSingleWheat3
		case entity.FarmLeftWheat2:
			tile.Sprite = entity.FarmLeftWheat3
		case entity.FarmMiddleWheat2:
			tile.Sprite = entity.FarmMiddleWheat3
		case entity.FarmRightWheat2:
			tile.Sprite = entity.FarmRightWheat3

		case entity.FarmSingleWheat3:
			tile.Sprite = entity.FarmSingleWheat4
		case entity.FarmLeftWheat3:
			tile.Sprite = entity.FarmLeftWheat4
		case entity.FarmMiddleWheat3:
			tile.Sprite = entity.FarmMiddleWheat4
		case entity.FarmRightWheat3:
			tile.Sprite = entity.FarmRightWheat4
		}
	}
}

func (f *Farm) FullyHarvestedAndCleaned(mp *m.Map) bool {
	for _, idx := range f.FarmableIdxs {
		if mp.Items[idx].Sprite != entity.NoItem {
			return false
		}
	}
	return true
}

func (f *Farm) FullyPlanted(mp *m.Map) bool {
	for _, idx := range f.FarmableIdxs {
		if !entity.IsFarm(mp.Items[idx].Sprite) {
			return false
		}
	}
	return true
}

func (f *Farm) ShouldHarvest(mp *m.Map) ([]int, bool) {
	if !entity.IsFarmHarvestable(f.farmTile.Sprite) {
		return nil, false
	}
	for _, idx := range f.FarmableIdxs {
		if !entity.IsFarmHarvestable(mp.Items[idx].Sprite) {
			return nil, false
		}
	}
	return f.FarmableIdxs, true
}

func (f *Farm) PlantFarm(mp *m.Map, t m.Tile) {
	if m.IsAnyWall(mp.OneTileLeft(t.Idx).Sprite) ||
		m.IsAnyWall(mp.OneTileRight(t.Idx).Sprite) ||
		m.IsAnyWall(mp.OneTileUp(t.Idx).Sprite) ||
		m.IsAnyWall(mp.OneTileUpLeft(t.Idx).Sprite) ||
		m.IsAnyWall(mp.OneTileUpRight(t.Idx).Sprite) ||
		m.IsAnyWall(mp.OneTileDown(t.Idx).Sprite) ||
		m.IsAnyWall(mp.OneTileDownLeft(t.Idx).Sprite) ||
		m.IsAnyWall(mp.OneTileDownRight(t.Idx).Sprite) {
		return
	}
	defer func() {
		mp.Items[t.Idx].Resource = entity.ResourceNone
	}()
	if entity.IsFarmSingle(mp.Items[m.OneLeft(t.Idx)].Sprite) {
		mp.Items[t.Idx-1].Sprite = entity.FarmLeftEmpty
		item.Place(mp, t.X, t.Y, entity.FarmRightEmpty)
		return
	}
	if entity.IsFarmRight(mp.Items[m.OneLeft(t.Idx)].Sprite) {
		mp.Items[t.Idx-1].Sprite = entity.FarmMiddleEmpty
		item.Place(mp, t.X, t.Y, entity.FarmRightEmpty)
		return
	}
	item.Place(mp, t.X, t.Y, entity.FarmSingleEmpty)
}

func (f *Farm) farmableIndexes(mp *m.Map) []int {
	idxs := []int{}
	for _, tIdx := range f.AllTileIdxs {
		if !entity.IsFarm(mp.Items[tIdx].Sprite) {
			continue
		}
		idxs = append(idxs, tIdx)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(idxs)))
	return idxs
}
