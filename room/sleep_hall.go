package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type SleepHall struct {
	tiles m.Tiles
}

func NewSleepHall(mp *m.Map, x, y int) *SleepHall {
	s := &SleepHall{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.WoodFloorVertical })
	if len(tiles) == 0 {
		return nil
	}
	sort.Sort(tiles)
	s.tiles = tiles
	for _, t := range s.tiles {
		if m.IsAnyWall(mp.Tiles[t.Idx].Sprite) ||
			m.IsAnyWall(mp.Tiles[m.OneTileDown(t.Idx)].Sprite) {
			continue
		}
		if globals.IsBed(mp.Items[t.Idx].Sprite) ||
			globals.IsBed(mp.Items[m.OneTileDown(t.Idx)].Sprite) ||
			globals.IsBed(mp.Items[m.OneTileLeft(t.Idx)].Sprite) ||
			globals.IsBed(mp.Items[m.OneTileDownLeft(t.Idx)].Sprite) ||
			globals.IsBed(mp.Items[m.OneTileRight(t.Idx)].Sprite) ||
			globals.IsBed(mp.Items[m.OneTileDownRight(t.Idx)].Sprite) ||
			globals.IsBed(mp.Items[m.OneTileUp(t.Idx)].Sprite) ||
			globals.IsBed(mp.Items[m.OneTileUpLeft(t.Idx)].Sprite) ||
			globals.IsBed(mp.Items[m.OneTileUpRight(t.Idx)].Sprite) {
			continue
		}
		if m.IsNextToDoorOpening(mp, t.Idx) {
			continue
		}
		mp.Items[t.Idx].Sprite = globals.BedRed1
		mp.Items[m.OneTileDown(t.Idx)].Sprite = globals.BedRed2
	}
	return s
}
