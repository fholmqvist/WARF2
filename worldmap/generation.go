package worldmap

import (
	"math/rand"
)

var (
	floorTiles = []int{
		FloorBricksOne, FloorBricksTwo, FloorBricksThree,
		FloorBricksFour, FloorBricksFive, FloorBricksSix,
		FloorBricksSeven, FloorBricksEight, FloorBricksNine,
		FloorBricksTen,
	}
)

// RandomFloorBrick returns
// the sprite for a random FloorBrick.
func RandomFloorBrick() int {
	return floorTiles[rand.Intn(len(floorTiles))]
}

// FloodFill finds an "island"
// of tiles based on the predicate
// function and sets the tiles island
// number to the given island number.
func FloodFill(x, y int, mp *Map, island int, predicate func(int) bool) {
	idx := XYToIdx(x, y)

	ok := predicate(idx)
	if !ok {
		return
	}

	if y > 0 {
		FloodFill(x, y-1, mp, island, predicate)
	}
	if x > 0 {
		FloodFill(x-1, y, mp, island, predicate)
	}
	if y < TilesH-1 {
		FloodFill(x, y+1, mp, island, predicate)
	}
	if x < TilesW-1 {
		FloodFill(x+1, y, mp, island, predicate)
	}
}
