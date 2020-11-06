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
