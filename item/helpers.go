package item

var blockingItems = append([]int{Table}, bookShelves...)

// IsBlocking returns true
// if the sprite is of a
// type that blocks.
func IsBlocking(sprite int) bool {
	for _, i := range blockingItems {
		if i == sprite {
			return true
		}
	}
	return false
}

// IsBookshelf returns
// if tile is of any
// type of BookShelf.
func IsBookshelf(sprite int) bool {
	return sprite >= BookShelfOne && sprite <= BookShelfTen
}

// IsChair returns
// if tile is of any
// type of Chair.
func IsChair(sprite int) bool {
	return sprite == ChairLeft || sprite == ChairRight
}

// IsLibraryItem returns
// if tile is of any
// type item that are
// included in libraries.
func IsLibraryItem(sprite int) bool {
	return sprite >= ChairLeft && sprite <= ChairRight
}
