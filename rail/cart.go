package rail

import (
	"log"

	e "github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"

	"github.com/beefsack/go-astar"
)

type Cart struct {
	e.Entity
	e.Walker
}

func NewCart(idx int) *Cart {
	return &Cart{
		Entity: e.Entity{
			Idx:    idx,
			Sprite: m.Cart,
		},
	}
}

func (c *Cart) InitiateRide(mp *m.Map, to *m.Tile) bool {
	from := &mp.Rails[c.Idx]
	path, _, ok := astar.Path(from, to)
	if !ok {
		return false
	}
	var pathIdxs []int
	for _, t := range m.Reverse(path) {
		tile := t.(*m.Tile)
		pathIdxs = append(pathIdxs, tile.Idx)
	}
	c.Path = pathIdxs
	return true
}

func (c *Cart) traversePath(mp *m.Map) {
	if len(c.Path) == 0 {
		return
	}
	next := c.Path[0]
	if c.Idx == next {
		c.Path = c.Path[1:]
		return
	}
	dir, err := m.NextIdxToDir(c.Idx, next)
	if err != nil {
		log.Fatal(err, c.Idx, next, c.Path)
	}
	if c.Move(mp, &c.Entity, dir) {
		c.Path = c.Path[1:]
	}
}
