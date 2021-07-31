package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var sleepHallAutoID = 0

type SleepHall struct {
	ID    int
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
		if entity.IsBed(mp.Items[t.Idx].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileDown(t.Idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileLeft(t.Idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileDownLeft(t.Idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileRight(t.Idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileDownRight(t.Idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileUp(t.Idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileUpLeft(t.Idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileUpRight(t.Idx)].Sprite) {
			continue
		}
		if m.IsNextToDoorOpening(mp, t.Idx) || m.IsNextToDoorOpening(mp, m.OneTileDown(t.Idx)) {
			continue
		}
		mp.Items[t.Idx].Sprite = entity.BedRed1
		mp.Items[m.OneTileDown(t.Idx)].Sprite = entity.BedRed2
	}
	for _, t := range tiles {
		mp.Tiles[t.Idx].Room = s
	}
	s.ID = sleepHallAutoID
	sleepHallAutoID++
	return s
}

func (s *SleepHall) GetID() int {
	return s.ID
}

func (s *SleepHall) String() string {
	return "SleepHall"
}

func (s *SleepHall) Update(mp *m.Map) {}

func (s *SleepHall) Tiles() []int {
	return s.tiles.ToIdxs()
}
