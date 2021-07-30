package item

import (
	"math"
	"sort"

	"github.com/Holmqvist1990/WARF2/entity"
	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/worldmap"
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
		d := gl.Dist(current.X, current.Y, itemTile.X, itemTile.Y)
		if d < nearest {
			nearest = d
			nearestIdx = itemTile.Idx
		}
	}
	return nearestIdx, true
}

func FindNearestMany(m *worldmap.Map, idx int, f func(int) bool) ([]int, bool) {
	current, ok := m.GetTileByIndex(idx)
	if !ok {
		return nil, false
	}
	tiles := []worldmap.Tile{}
	for _, itemTile := range m.Items {
		if !f(itemTile.Sprite) {
			continue
		}
		tiles = append(tiles, itemTile)
	}
	sort.Slice(tiles, func(i, j int) bool {
		first := gl.Dist(current.X, current.Y, tiles[i].X, tiles[i].Y)
		second := gl.Dist(current.X, current.Y, tiles[j].X, tiles[j].Y)
		return first < second
	})
	idxs := make([]int, len(tiles))
	for i := range tiles {
		idxs[i] = tiles[i].Idx
	}
	return idxs, true
}

func FindNearestBookshelf(m *worldmap.Map, idx int) (int, bool) {
	return FindNearest(m, idx, entity.IsBookshelf)
}

func FindNearestChair(m *worldmap.Map, idx int) (int, bool) {
	return FindNearest(m, idx, entity.IsChair)
}

func FindNearestChairs(m *worldmap.Map, idx int) ([]int, bool) {
	return FindNearestMany(m, idx, entity.IsChair)
}

func FindNearestBeds(m *worldmap.Map, idx int) ([]int, bool) {
	return FindNearestMany(m, idx, entity.IsBedTop)
}
