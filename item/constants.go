// Package item describes the
// in-game items that are placed
// on top of game tiles.
package item

import "github.com/Holmqvist1990/WARF2/globals"

// Item tile sprite constant.
const (
	NoItem = iota
	WallCrumbled1
	WallCrumbled2
	WallCrumbled3
	WallCrumbled4
)
const (
	BookShelfOne = iota + globals.TilesetW
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
	FarmSingleEmpty = iota + globals.TilesetW*2
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
