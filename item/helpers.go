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

func IsBookshelf(sprite int) bool {
	return sprite >= BookShelfOne && sprite <= BookShelfTen
}

func IsChair(sprite int) bool {
	return sprite == ChairLeft || sprite == ChairRight
}

func IsLibraryItem(sprite int) bool {
	return sprite >= ChairLeft && sprite <= ChairRight
}

func IsCrumbledWall(sprite int) bool {
	return sprite >= WallCrumbled1 && sprite <= WallCrumbled4
}

func IsFarmSingle(sprite int) bool {
	return sprite == FarmSingleEmpty
}

func IsFarmRight(sprite int) bool {
	return sprite == FarmRightEmpty
}
