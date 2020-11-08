package worldmap

import (
	"math"

	"github.com/beefsack/go-astar"
)

var offsets = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

var _ astar.Pather = &Tile{}

// PathNeighbors is the implementation for
// the interface required by go-astar for
// determining surrounding (walkable) neighbors.
func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}

	for _, offset := range offsets {
		offsetX, offsetY := offset[0], offset[1]
		neighbor, ok := t.Map.GetTile(t.X+offsetX, t.Y+offsetY)

		if !ok || !IsExposed(neighbor.Sprite) || neighbor.Blocked {
			continue
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

	xDist := math.Abs(float64(toT.X - t.X))
	yDist := math.Abs(float64(toT.Y - t.Y))

	return xDist + yDist
}

// Reverse an astar path
func Reverse(path []astar.Pather) []astar.Pather {
	var reverse []astar.Pather
	for i := len(path) - 1; i >= 0; i-- {
		reverse = append(reverse, path[i])
	}
	return reverse
}
