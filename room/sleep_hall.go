package room

import (
	"fmt"
	"sort"

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
		fmt.Println(t)
	}
	return s
}
