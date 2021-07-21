package worldmap

import "github.com/Holmqvist1990/WARF2/globals"

// World tile sprite constant.
const (
	None = iota

	Ground

	BoundarySolid
	BoundaryExposed

	WallSolid
	WallExposed
	WallSelectedSolid
	WallSelectedExposed
)
const (
	FloorBricks1 = iota + globals.TilesetW
	FloorBricks2
	FloorBricks3
	FloorBricks4
	FloorBricks5
	FloorBricks6
	FloorBricks7
	FloorBricks8
	FloorBricks9
	FloorBricks10
)
const (
	WoodFloor1 = iota + globals.TilesetW*2
	WoodFloor2
	WoodFloor3
	WoodFloor4

	WoodFloorVertical
)
const (
	Straight = iota + 1
	Curve
	Stop
	Cross
	Cart
)
