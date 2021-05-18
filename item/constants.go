// Package item describes the
// in-game items that are placed
// on top of game tiles.
package item

import "projects/games/warf2/globals"

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
