package item

import "math/rand"

var (
	bookShelves = []int{
		BookShelfOne, BookShelfTwo, BookShelfThree,
		BookShelfFour, BookShelfFive, BookShelfSix,
		BookShelfSeven, BookShelfEight, BookShelfNine,
		BookShelfTen,
	}

	furniture = []int{
		ChairLeft, Table, ChairRight,
	}
)

// RandomBookshelf returns
// the sprite for a random BookShelf.
func RandomBookshelf() int {
	return bookShelves[rand.Intn(len(bookShelves))]
}

// RandomFurniture returns
// the sprite for a random Furniture.
func RandomFurniture() int {
	return furniture[rand.Intn(len(furniture))]
}
