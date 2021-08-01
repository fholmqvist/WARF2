package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var sleepHallAutoID = 0

type SleepHall struct {
	ID    int
	tiles []int
}

func NewSleepHall(mp *m.Map, x, y int) *SleepHall {
	s := &SleepHall{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.SleepHallFloor })
	if len(tiles) == 0 {
		return nil
	}
	sort.Ints(tiles)
	s.tiles = tiles
	for _, idx := range s.tiles {
		mp.Tiles[idx].Room = s
		if m.IsAnyWall(mp.Tiles[idx].Sprite) ||
			m.IsAnyWall(mp.Tiles[m.OneTileDown(idx)].Sprite) {
			continue
		}
		if entity.IsBed(mp.Items[idx].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileDown(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileLeft(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileDownLeft(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileRight(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileDownRight(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileUp(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileUpLeft(idx)].Sprite) ||
			entity.IsBed(mp.Items[m.OneTileUpRight(idx)].Sprite) {
			continue
		}
		if m.IsNextToDoorOpening(mp, idx) || m.IsNextToDoorOpening(mp, m.OneTileDown(idx)) {
			continue
		}
		mp.Items[idx].Sprite = entity.BedRed1
		mp.Items[m.OneTileDown(idx)].Sprite = entity.BedRed2
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
	return s.tiles
}
