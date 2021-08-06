package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/item"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var sleepHallAutoID = 0

type SleepHall struct {
	ID    int
	tiles []int
}

func NewSleepHall(mp *m.Map, x, y int) (*SleepHall, bool) {
	s := &SleepHall{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.SleepHallFloor })
	if len(tiles) == 0 {
		return nil, false
	}
	sort.Ints(tiles)
	s.tiles = tiles
	for _, idx := range s.tiles {
		mp.Tiles[idx].Room = s
		if m.IsAnyWall(mp.Tiles[idx].Sprite) ||
			m.IsAnyWall(mp.Tiles[m.OneDown(idx)].Sprite) {
			continue
		}
		if entity.IsBed(mp.Items[idx].Sprite) ||
			entity.IsBed(mp.Items[m.OneDown(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneLeft(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneDownLeft(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneRight(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneDownRight(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneUp(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneUpLeft(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneUpRight(idx)].Sprite) {
			continue
		}
		if m.IsNextToDoorOpening(mp, idx) || m.IsNextToDoorOpening(mp, m.OneDown(idx)) {
			continue
		}
		mp.Items[idx].Sprite, mp.Items[m.OneDown(idx)].Sprite = item.RandomBed()
	}
	s.ID = sleepHallAutoID
	sleepHallAutoID++
	return s, true
}

func (s *SleepHall) GetID() int {
	if s == nil {
		return -1
	}
	return s.ID
}

func (s *SleepHall) String() string {
	return "SleepHall"
}

func (s *SleepHall) Update(mp *m.Map) {}

func (s *SleepHall) Tiles() []int {
	return s.tiles
}
