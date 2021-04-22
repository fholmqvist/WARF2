package item

import (
	"math"
	"projects/games/warf2/worldmap"
)

func FindNearest(m *worldmap.Map, idx int, f func(int) bool) (int, bool) {
	nearestIdx := -1
	nearest := math.MaxFloat64
	current, ok := m.GetTileByIndex(idx)
	if !ok {
		return -1, false
	}
	for _, itemTile := range m.Items {
		if !f(itemTile.Sprite) {
			continue
		}
		d := dist(*current, itemTile)
		if d < nearest {
			nearest = d
			nearestIdx = itemTile.Idx
		}
	}
	return nearestIdx, true
}

func FindNearestBookshelf(m *worldmap.Map, idx int) (int, bool) {
	return FindNearest(m, idx, IsBookshelf)
}

func FindNearestChair(m *worldmap.Map, idx int) (int, bool) {
	return FindNearest(m, idx, IsChair)
}

func dist(a, b worldmap.Tile) float64 {
	xDist := math.Abs(float64(b.X - a.X))
	yDist := math.Abs(float64(b.Y - a.Y))
	return xDist + yDist
}
