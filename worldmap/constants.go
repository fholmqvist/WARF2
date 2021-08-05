package worldmap

import "github.com/Holmqvist1990/WARF2/globals"

// World tile sprite constant.
const (
	None = iota
	WARNING

	Ground

	BoundarySolid
	BoundaryExposed

	WallSolid
	WallExposed
	WallSelectedSolid
	WallSelectedExposed
)
const (
	StorageFloor1 = iota + globals.TilesetW
	StorageFloor2
	StorageFloor3
	StorageFloor4
	StorageFloor5
	StorageFloor6
	StorageFloor7
	StorageFloor8
	StorageFloor9
	StorageFloor10
)
const (
	LibraryFloor1 = iota + globals.TilesetW*2
	LibraryFloor2
	LibraryFloor3
	LibraryFloor4

	SleepHallFloor

	BreweryFloor

	BarFloor
)
const (
	Straight = iota + 1
	Curve
	Stop
	Cross
	Cart
)
