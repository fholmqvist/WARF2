package room

import (
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var barAutoID = 0

type Bar struct {
	ID    int
	tiles []int
}

func NewBar(mp *m.Map, x, y int) *Bar {
	b := &Bar{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.BarFloor })
	if len(tiles) == 0 {
		return nil
	}
	sort.Ints(tiles)
	for _, idx := range tiles {
		tile := &mp.Tiles[idx]
		tile.Room = b
	}
	///////////////////////////
	// TODO
	// Actually generate items.
	///////////////////////////
	mp.Items[tiles[0]].Sprite = entity.BarDrinksLeft
	mp.Items[tiles[1]].Sprite = entity.BarDrinksRight
	mp.Items[tiles[12]].Sprite = entity.BarLeft
	mp.Items[tiles[13]].Sprite = entity.BarH
	mp.Items[tiles[14]].Sprite = entity.BarH
	mp.Items[tiles[15]].Sprite = entity.BarRight
	mp.Items[tiles[9]].Sprite = entity.BarV
	mp.Items[tiles[18]].Sprite = entity.BarStool
	mp.Items[tiles[19]].Sprite = entity.BarStool
	mp.Items[tiles[20]].Sprite = entity.BarStool
	mp.Items[tiles[21]].Sprite = entity.BarStool
	mp.Items[tiles[16]].Sprite = entity.BarStool
	mp.Items[tiles[10]].Sprite = entity.BarStool
	b.ID = barAutoID
	barAutoID++
	return b
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
