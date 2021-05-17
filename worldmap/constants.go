package worldmap

import "projects/games/warf2/globals"

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
	FloorBricks1 = iota + globals.ActualTileSize
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
	WoodFloor1 = iota + globals.ActualTileSize*2
	WoodFloor2
	WoodFloor3
	WoodFloor4
)
const (
	Straight = iota + 1
	Curve
	Stop
	Cross
	Cart
)
