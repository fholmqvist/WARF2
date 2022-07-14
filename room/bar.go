package room

import (
	"math/rand"
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	gl "github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var barAutoID = 0

const NEEDS_MORE_BEER_RATE = 5

type Bar struct {
	ID              int
	tiles           []int
	Beers           uint
	BeerRefillIndex int
}

func NewBar(mp *m.Map, x, y int) (*Bar, bool) {
	b := &Bar{}
	tiles := mp.FloodFillRoom(x, y, func() int { return m.BarFloor })
	if len(tiles) == 0 {
		return nil, false
	}
	b.tiles = tiles
	sort.Ints(tiles)
	hasPlacedBar := false
	for _, idx := range tiles {
		mp.Tiles[idx].Room = b
		if !hasPlacedBar {
			hasPlacedBar = b.placeBar(mp, tiles, idx)
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

func (b *Bar) AddBeer(amount uint) {
	b.Beers += amount
}

func (b *Bar) ConsumeBeer() bool {
	if b.Beers <= 0 {
		return false
	}
	b.Beers--
	return true
}

func (b *Bar) NeedsMoreBeer() bool {
	return b.Beers <= NEEDS_MORE_BEER_RATE
}

var bars = [][]int{
	{
		entity.BarDrinksLeft, entity.BarDrinksRight, entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem,
		entity.NoItem, entity.NoItem, entity.NoItem, entity.BarV, entity.BarStool, entity.NoItem,
		entity.BarLeft, entity.BarH, entity.BarH, entity.BarRight, entity.BarStool, entity.NoItem,
		entity.BarStool, entity.BarStool, entity.BarStool, entity.BarStool, entity.BarStool, entity.NoItem,
		entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem,
	},
	{
		entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem, entity.BarDrinksLeft, entity.BarDrinksRight,
		entity.NoItem, entity.BarStool, entity.BarV, entity.NoItem, entity.NoItem, entity.NoItem,
		entity.NoItem, entity.BarStool, entity.BarLeft, entity.BarH, entity.BarH, entity.BarRight,
		entity.NoItem, entity.BarStool, entity.BarStool, entity.BarStool, entity.BarStool, entity.BarStool,
		entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem, entity.NoItem,
	},
}

func (b *Bar) placeBar(mp *m.Map, tiles []int, idx int) bool {
	////////////////
	// TODO
	// This is crap.
	////////////////
	placements := []int{}
	idxX, idxY := gl.IdxToXY(idx)
	width, height := 6, 5
	for y := idxY; y < idxY+height; y++ {
		for x := idxX; x < idxX+width; x++ {
			curr := gl.XYToIdx(x, y)
			if mp.Items[curr].Sprite != entity.NoItem {
				return false
			}
			if m.IsAnyWall(mp.Tiles[curr].Sprite) {
				return false
			}
			placements = append(placements, curr)
		}
	}
	randomBar := bars[rand.Intn(len(bars))]
	for i := 0; i < len(placements); i++ {
		mp.Items[placements[i]].Sprite = randomBar[i]
		if randomBar[i] == entity.BarDrinksLeft {
			b.BeerRefillIndex = m.OneDown(placements[i])
		}
	}
	return true
}
