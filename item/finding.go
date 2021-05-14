package item

import (
	"math"
	"projects/games/warf2/globals"
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
		d := globals.Dist(current.X, current.Y, itemTile.X, itemTile.Y)
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
