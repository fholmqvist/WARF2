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

	bookShelves = []int{
		BookShelfOne, BookShelfTwo, BookShelfThree,
		BookShelfFour,
	}
)

// RandomFloorBrick returns
// the sprite for a random FloorBrick.
func RandomFloorBrick() int {
	return floorTiles[rand.Intn(len(floorTiles))]
}

// RandomBookshelf returns
// the sprite for a random BookShelf.
func RandomBookshelf() int {
	return bookShelves[rand.Intn(len(bookShelves))]
}
