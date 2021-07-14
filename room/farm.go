package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/item"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Farm struct {
	tiles m.Tiles
}

func NewFarm(mp *m.Map, x, y int) *Farm {
	f := &Farm{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.None })
	if len(tiles) == 0 {
		return nil
	}
	sort.Sort(tiles)
	f.tiles = tiles
	for _, t := range f.tiles {
		f.placeFarm(mp, t)
	}
	return f
}

func (f *Farm) placeFarm(mp *m.Map, t m.Tile) {
	earlyExists := []bool{
		m.IsAnyWall(mp.OneTileLeft(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileRight(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileUp(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileUpLeft(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileUpRight(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileDown(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileDownLeft(t.Idx).Sprite),
		m.IsAnyWall(mp.OneTileDownRight(t.Idx).Sprite),
	}
	for _, ee := range earlyExists {
		if ee {
			return
		}
	}
	if item.IsFarmSingle(mp.Items[m.OneTileLeft(t.Idx)].Sprite) {
		mp.Items[t.Idx-1].Sprite = item.FarmLeftEmpty
		item.Place(mp, t.X, t.Y, item.FarmRightEmpty)
		return
	}
	if item.IsFarmRight(mp.Items[m.OneTileLeft(t.Idx)].Sprite) {
		mp.Items[t.Idx-1].Sprite = item.FarmMiddleEmpty
		item.Place(mp, t.X, t.Y, item.FarmRightEmpty)
		return
	}
	item.Place(mp, t.X, t.Y, item.FarmSingleEmpty)
	return
}
