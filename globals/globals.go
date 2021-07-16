package globals

// Game globals.
const (
	TileSize = 16
	TilesetW = 16

	ScreenWidth  = (36 + 10) * TileSize
	ScreenHeight = (24 + 10) * TileSize

	TilesW      = ScreenWidth / TileSize
	TilesH      = (ScreenHeight / TileSize) - 2
	TilesT      = TilesW * TilesH
	TilesBottom = TilesT - TilesW

	TPS         = 30
	CycleLength = TPS * 8
)

var (
	DEBUG      = false
	PAUSE_GAME = false
)
