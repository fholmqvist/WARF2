package globals

// Game globals.
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

	ActualTileSize = 16
)

var (
	DEBUG = false
)
