package worldmap

const actualTileSize = 16

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
	WallCrumbled1
	WallCrumbled2
	WallCrumbled3
	WallCrumbled4
)
const (
	FloorBricks1 = iota + actualTileSize
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
	WoodFloor1 = iota + actualTileSize*2
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
