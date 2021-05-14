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

	FloorBricksOne = iota + 8
	FloorBricksTwo
	FloorBricksThree
	FloorBricksFour
	FloorBricksFive
	FloorBricksSix
	FloorBricksSeven
	FloorBricksEight
	FloorBricksNine
	FloorBricksTen

	WoodFloorOne = iota + 14
	WoodFloorTwo
	WoodFloorThree
	WoodFloorFour
)

const (
	Straight = iota + 1
	Curve
	Stop
	Cross
	Cart
)
