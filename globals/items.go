package globals

// Due to "import cycle not allowed" :(
// Generally a great rule, but really sucks
// when you have to do things like this.
//
// item.Table makes more sense than globals.Table.

const (
	NoItem = iota
	WallCrumbled1
	WallCrumbled2
	WallCrumbled3
	WallCrumbled4
)
const (
	BookShelfOne = iota + TilesetW
	BookShelfTwo
	BookShelfThree
	BookShelfFour
	BookShelfFive
	BookShelfSix
	BookShelfSeven
	BookShelfEight
	BookShelfNine
	BookShelfTen

	ChairLeft
	Table
	ChairRight
)
const (
	FarmSingleEmpty = iota + TilesetW*2
	FarmLeftEmpty
	FarmMiddleEmpty
	FarmRightEmpty

	FarmSingleWheat1
	FarmLeftWheat1
	FarmMiddleWheat1
	FarmRightWheat1

	FarmSingleWheat2
	FarmLeftWheat2
	FarmMiddleWheat2
	FarmRightWheat2

	FarmSingleWheat3
	FarmLeftWheat3
	FarmMiddleWheat3
	FarmRightWheat3

	FarmSingleWheat4
	FarmLeftWheat4
	FarmMiddleWheat4
	FarmRightWheat4

	Wheat
)

var (
	BookShelves = []int{
		BookShelfOne, BookShelfTwo, BookShelfThree,
		BookShelfFour, BookShelfFive, BookShelfSix,
		BookShelfSeven, BookShelfEight, BookShelfNine,
		BookShelfTen,
	}

	Furniture = []int{
		ChairLeft, Table, ChairRight,
	}
)

var blockingItems = append([]int{Table}, BookShelves...)

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

func IsCrumbledWall(sprite int) bool {
	return sprite >= WallCrumbled1 && sprite <= WallCrumbled4
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

func IsFarm(sprite int) bool {
	// HELL YEAH!
	return sprite == FarmLeftEmpty ||
		sprite == FarmMiddleEmpty ||
		sprite == FarmRightEmpty ||
		sprite == FarmSingleEmpty ||
		sprite == FarmLeftWheat1 ||
		sprite == FarmMiddleWheat1 ||
		sprite == FarmRightWheat1 ||
		sprite == FarmSingleWheat1 ||
		sprite == FarmLeftWheat2 ||
		sprite == FarmMiddleWheat2 ||
		sprite == FarmRightWheat2 ||
		sprite == FarmSingleWheat2 ||
		sprite == FarmLeftWheat3 ||
		sprite == FarmMiddleWheat3 ||
		sprite == FarmRightWheat3 ||
		sprite == FarmSingleWheat3 ||
		sprite == FarmLeftWheat4 ||
		sprite == FarmMiddleWheat4 ||
		sprite == FarmRightWheat4 ||
		sprite == FarmSingleWheat4
}

func IsFarmSingle(sprite int) bool {
	return sprite == FarmSingleEmpty
}

func IsFarmRight(sprite int) bool {
	return sprite == FarmRightEmpty
}

func IsFarmHarvestable(sprite int) bool {
	return sprite == FarmLeftWheat4 ||
		sprite == FarmMiddleWheat4 ||
		sprite == FarmRightWheat4 ||
		sprite == FarmSingleWheat4
}

func IsFarmTileHarvested(sprite int) bool {
	///////////////////////////
	// TODO
	// Support more than Wheat.
	///////////////////////////
	return sprite == FarmLeftEmpty ||
		sprite == FarmMiddleEmpty ||
		sprite == FarmRightEmpty ||
		sprite == FarmSingleEmpty ||
		sprite == Wheat
}
