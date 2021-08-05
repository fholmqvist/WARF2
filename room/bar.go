package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var barAutoID = 0

type Bar struct {
	ID    int
	tiles []int
}

func NewBar(mp *m.Map, x, y int) (*Bar, bool) {
	b := &Bar{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.BarFloor })
	if len(tiles) == 0 {
		return nil, false
	}
	sort.Ints(tiles)
	hasPlacedBar := false
	for _, idx := range tiles {
		tile := &mp.Tiles[idx]
		tile.Room = b
		if !hasPlacedBar {
			hasPlacedBar = placeBar(mp, tiles, idx)
		}
	}
	if !hasPlacedBar {
		for _, idx := range tiles {
			mp.Tiles[idx].Sprite = m.Ground
			mp.Tiles[idx].Room = nil
		}
		return nil, false
	}
	b.ID = barAutoID
	barAutoID++
	return b, true
}

func (b *Bar) GetID() int {
	return b.ID
}

func (b *Bar) String() string {
	return "Bar"
}

func (b *Bar) Update(mp *m.Map) {

}

func (b *Bar) Tiles() []int {
	return b.tiles
}

func placeBar(mp *m.Map, tiles []int, idx int) bool {
	placements := []int{}
	idxX, idxY := globals.IdxToXY(idx)
	width, height := 5, 5
	for y := idxY; y < idxY+height; y++ {
		for x := idxX; x < idxX+width; x++ {
			curr := globals.XYToIdx(x, y)
			if mp.Items[curr].Sprite != entity.NoItem {
				return false
			}
			if m.IsAnyWall(mp.Tiles[curr].Sprite) {
				return false
			}
			placements = append(placements, curr)
		}
	}
	barItems := []int{
		entity.BarDrinksLeft, entity.BarDrinksRight, entity.NoItem, entity.NoItem, entity.NoItem,
		entity.NoItem, entity.NoItem, entity.NoItem, entity.BarV, entity.BarStool,
		entity.BarLeft, entity.BarH, entity.BarH, entity.BarRight, entity.BarStool,
		entity.BarStool, entity.BarStool, entity.BarStool, entity.BarStool, entity.BarStool,
		entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem,
	}
	for i := 0; i < len(placements); i++ {
		mp.Items[placements[i]].Sprite = barItems[i]
	}
	return true
}
