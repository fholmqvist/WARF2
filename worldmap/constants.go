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
	Transparent = iota
	Ground
	BoundarySolid
	BoundaryExposed
	WallSolid
	WallExposed
	WallSelectedSolid
	WallSelectedExposed
)
