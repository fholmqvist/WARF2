package room

import (
	"sort"

	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/item"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Farm struct {
	ID       int     // First tile.
	tileIdxs []int   // To be indexed against Worldmap.
	farmTile *m.Tile // Knows when farm has reached maturity.
}

func NewFarm(mp *m.Map, x, y int) *Farm {
	f := &Farm{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.None })
	if len(tiles) == 0 {
		return nil
	}
	sort.Sort(tiles)
	f.ID = tiles[0].Idx
	f.tileIdxs = tiles.ToIdxs()
	for _, t := range tiles {
		f.PlantFarm(mp, t)
		if f.farmTile != nil {
			continue
		}
		if !gl.IsFarm(mp.Items[t.Idx].Sprite) {
			continue
		}
		f.farmTile = &mp.Items[t.Idx]
	}
	return f
}

func (f *Farm) Update(mp *m.Map) {
	for _, tIdx := range f.tileIdxs {
		tile := &mp.Items[tIdx]
		if tile.Sprite == 0 {
			continue
		}
		switch tile.Sprite {
		case gl.FarmSingleEmpty:
			tile.Sprite = gl.FarmSingleWheat1
		case gl.FarmLeftEmpty:
			tile.Sprite = gl.FarmLeftWheat1
		case gl.FarmMiddleEmpty:
			tile.Sprite = gl.FarmMiddleWheat1
		case gl.FarmRightEmpty:
			tile.Sprite = gl.FarmRightWheat1

		case gl.FarmSingleWheat1:
			tile.Sprite = gl.FarmSingleWheat2
		case gl.FarmLeftWheat1:
			tile.Sprite = gl.FarmLeftWheat2
		case gl.FarmMiddleWheat1:
			tile.Sprite = gl.FarmMiddleWheat2
		case gl.FarmRightWheat1:
			tile.Sprite = gl.FarmRightWheat2

		case gl.FarmSingleWheat2:
			tile.Sprite = gl.FarmSingleWheat3
		case gl.FarmLeftWheat2:
			tile.Sprite = gl.FarmLeftWheat3
		case gl.FarmMiddleWheat2:
			tile.Sprite = gl.FarmMiddleWheat3
		case gl.FarmRightWheat2:
			tile.Sprite = gl.FarmRightWheat3

		case gl.FarmSingleWheat3:
			tile.Sprite = gl.FarmSingleWheat4
		case gl.FarmLeftWheat3:
			tile.Sprite = gl.FarmLeftWheat4
		case gl.FarmMiddleWheat3:
			tile.Sprite = gl.FarmMiddleWheat4
		case gl.FarmRightWheat3:
			tile.Sprite = gl.FarmRightWheat4
		}
	}
}

func (f *Farm) FullyHarvestedAndCleaned(mp *m.Map) bool {
	for _, idx := range f.tileIdxs {
		if !gl.IsFarmTileHarvested(mp.Items[idx].Sprite) {
			return false
		}
	}
	return true
}

func (f *Farm) ShouldHarvest(mp *m.Map) ([]int, bool) {
	if !gl.IsFarmHarvestable(f.farmTile.Sprite) {
		return nil, false
	}
	return f.GetHarvestIdxs(mp), true
}

func (f *Farm) GetHarvestIdxs(mp *m.Map) []int {
	idxs := []int{}
	for _, tIdx := range f.tileIdxs {
		if !gl.IsFarm(mp.Items[tIdx].Sprite) {
			continue
		}
		idxs = append(idxs, tIdx)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(idxs)))
	return idxs
}

func (f *Farm) PlantFarm(mp *m.Map, t m.Tile) {
	skips := []bool{
		m.IsAnyWall(mp.OneTileLeft(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileRight(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileUp(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileUpLeft(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileUpRight(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileDown(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileDownLeft(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileDownRight(t.Idx).Sprite),
	}
	for _, skip := range skips {
		if skip {
			return
		}
	}
	if gl.IsFarmSingle(mp.Items[m.OneTileLeft(t.Idx)].Sprite) {
		mp.Items[t.Idx-1].Sprite = gl.FarmLeftEmpty
		item.Place(mp, t.X, t.Y, gl.FarmRightEmpty)
		return
	}
	if gl.IsFarmRight(mp.Items[m.OneTileLeft(t.Idx)].Sprite) {
		mp.Items[t.Idx-1].Sprite = gl.FarmMiddleEmpty
		item.Place(mp, t.X, t.Y, gl.FarmRightEmpty)
		return
	}
	item.Place(mp, t.X, t.Y, gl.FarmSingleEmpty)
	return
}
