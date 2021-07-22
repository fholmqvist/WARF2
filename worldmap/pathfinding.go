package worldmap

import (
	"github.com/Holmqvist1990/WARF2/globals"

	"github.com/beefsack/go-astar"
)

var offsets = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

var (
	_ astar.Pather = &Tile{}
)

// PathNeighbors is the implementation for
// the interface required by go-astar for
// determining surrounding (walkable) neighbors.
func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, offset := range offsets {
		offsetX, offsetY := offset[0], offset[1]
		var neighbor *Tile
		switch t.TileType {
		case NormalTile:
			normalTile, ok := t.Map.GetTile(t.X+offsetX, t.Y+offsetY)
			if !ok {
				continue
			}
			itemTile, ok := t.Map.GetItemTile(t.X+offsetX, t.Y+offsetY)
			if !ok {
				continue
			}
			if IsAnyWall(normalTile.Sprite) || Blocking(normalTile, itemTile) {
				continue
			}
			neighbor = normalTile
		case RailTile:
			n, ok := t.Map.GetRailTile(t.X+offsetX, t.Y+offsetY)
			if !ok {
				continue
			}
			if !IsRail(n.Sprite) {
				continue
			}
			neighbor = n
		}
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

// PathNeighborCost is the implementation for
// the interface required by go-astar for
// determining cost-to-walk for a given tile.
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

// PathEstimatedCost is the implementation for
// the interface required by go-astar for
// determining the cost of the entire path.
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	return globals.Dist(t.X, t.Y, toT.X, toT.Y)
}

// Reverse an astar path
func Reverse(path []astar.Pather) []astar.Pather {
	var reverse []astar.Pather
	for i := len(path) - 1; i >= 0; i-- {
		reverse = append(reverse, path[i])
	}
	return reverse
}
