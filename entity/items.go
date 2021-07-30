package entity

import gl "github.com/Holmqvist1990/WARF2/globals"

func ItemToString(itm int) string {
	switch itm {
	case NoItem:
		return "No item"
	case WallCrumbled1:
		return "Crumbled wall"
	case WallCrumbled2:
		return "Crumbled wall"
	case WallCrumbled3:
		return "Crumbled wall"
	case WallCrumbled4:
		return "Crumbled wall"
	case BookShelfOne:
		return "Bookshelf"
	case BookShelfTwo:
		return "Bookshelf"
	case BookShelfThree:
		return "Bookshelf"
	case BookShelfFour:
		return "Bookshelf"
	case BookShelfFive:
		return "Bookshelf"
	case BookShelfSix:
		return "Bookshelf"
	case BookShelfSeven:
		return "Bookshelf"
	case BookShelfEight:
		return "Bookshelf"
	case BookShelfNine:
		return "Bookshelf"
	case BookShelfTen:
		return "Bookshelf"
	case ChairLeft:
		return "Chair"
	case Table:
		return "Table"
	case ChairRight:
		return "Chair"
	case FarmSingleEmpty:
		return "Farm"
	case FarmLeftEmpty:
		return "Farm"
	case FarmMiddleEmpty:
		return "Farm"
	case FarmRightEmpty:
		return "Farm"
	case FarmSingleWheat1:
		return "Farm"
	case FarmLeftWheat1:
		return "Farm"
	case FarmMiddleWheat1:
		return "Farm"
	case FarmRightWheat1:
		return "Farm"
	case FarmSingleWheat2:
		return "Farm"
	case FarmLeftWheat2:
		return "Farm"
	case FarmMiddleWheat2:
		return "Farm"
	case FarmRightWheat2:
		return "Farm"
	case FarmSingleWheat3:
		return "Farm"
	case FarmLeftWheat3:
		return "Farm"
	case FarmMiddleWheat3:
		return "Farm"
	case FarmRightWheat3:
		return "Farm"
	case FarmSingleWheat4:
		return "Farm"
	case FarmLeftWheat4:
		return "Farm"
	case FarmMiddleWheat4:
		return "Farm"
	case FarmRightWheat4:
		return "Farm"
	case Wheat:
		return "Wheat"
	case BedRed1:
		return "Bed"
	case BedRed2:
		return "Bed"
	default:
		return "unknown"
	}
}

const (
	NoItem = iota
	WallCrumbled1
	WallCrumbled2
	WallCrumbled3
	WallCrumbled4
)
const (
	BookShelfOne = iota + gl.TilesetW
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
	FarmSingleEmpty = iota + gl.TilesetW*2
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
const (
	BedRed1 = iota + gl.TilesetW*4
	BedRed2 = gl.TilesetW * 5
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
	Beds = []int{
		BedRed1, BedRed2,
	}
)

var blockingItems = [][]int{BookShelves, Furniture, Beds}

func IsItemBlocking(sprite int) bool {
	for _, xs := range blockingItems {
		for _, x := range xs {
			if x == sprite {
				return true
			}
		}
	}
	return false
}

func IsCarriable(sprite int) bool {
	return IsCrumbledWall(sprite) || IsFarmTileHarvested(sprite)
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
	return IsChair(sprite) || sprite == Table || IsBookshelf(sprite)
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
	return sprite == Wheat
}

func IsBed(sprite int) bool {
	return sprite >= BedRed1 && sprite <= BedRed2
}

func IsBedTop(sprite int) bool {
	return sprite == BedRed1
}
