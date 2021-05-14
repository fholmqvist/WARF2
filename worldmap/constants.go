package worldmap

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

	FloorBricks1 = iota + 8
	FloorBricks2
	FloorBricks3
	FloorBricks4
	FloorBricks5
	FloorBricks6
	FloorBricks7
	FloorBricks8
	FloorBricks9
	FloorBricks10

	WoodFloor1 = iota + 14
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
