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

func IsFarm(sprite int) bool {
	for _, tile := range []int{
		FarmLeftEmpty,
		FarmMiddleEmpty,
		FarmRightEmpty,
		FarmSingleEmpty,
		FarmLeftWheat1,
		FarmMiddleWheat1,
		FarmRightWheat1,
		FarmSingleWheat1,
		FarmLeftWheat2,
		FarmMiddleWheat2,
		FarmRightWheat2,
		FarmSingleWheat2,
		FarmLeftWheat3,
		FarmMiddleWheat3,
		FarmRightWheat3,
		FarmSingleWheat3,
		FarmLeftWheat4,
		FarmMiddleWheat4,
		FarmRightWheat4,
		FarmSingleWheat4,
	} {
		if sprite == tile {
			return true
		}
	}
	return false
}

func IsFarmSingle(sprite int) bool {
	return sprite == FarmSingleEmpty
}

func IsFarmRight(sprite int) bool {
	return sprite == FarmRightEmpty
}

func IsFarmHarvestable(sprite int) bool {
	for _, tile := range []int{FarmLeftWheat4, FarmMiddleWheat4, FarmRightWheat4, FarmSingleWheat4} {
		if sprite == tile {
			return true
		}
	}
	return false
}
