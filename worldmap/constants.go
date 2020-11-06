package worldmap

// Game constants.
const (
	TileSize = 16
	TilesetW = 16

	ScreenWidth  = (36 + 10) * TileSize
	ScreenHeight = (24 + 8) * TileSize

	TilesW      = ScreenWidth / TileSize
	TilesH      = ScreenHeight / TileSize
	TilesT      = TilesW * TilesH
	TilesBottom = TilesT - TilesW

	TPS         = 30
	CycleLength = TPS * 8
)

// World tiles.
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
)
