package worldmap

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/globals"
)

// TileDir contains an index
// and the direction it is in
// relation to the index it was
// requested from.
type TileDir struct {
	Idx int
	Dir Direction
}

// Direction type for collision checking.
type Direction int

// Enum for directions.
const (
	Up Direction = iota
	Down
	Left
	Right
	UpLeft
	UpRight
	DownLeft
	DownRight
)

// GetDirection converts an integer into a Direction.
func GetDirection(i int) (Direction, error) {
	switch i {
	case 0:
		return Up, nil
	case 1:
		return Down, nil
	case 2:
		return Left, nil
	case 3:
		return Right, nil
	}

	return Up, fmt.Errorf("no such direction: %d", i)
}

// DirectionToText takes the integer
// representation for a direction and
// returns the name of the direction.
func DirectionToText(dir Direction) string {
	switch dir {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}

	return "Unknown direction"
}

// IndexAtDirection returns the index
// at the direction from the current index.
func IndexAtDirection(idx int, dir Direction) int {
	switch dir {

	case Up:
		return OneTileUp(idx)
	case Right:
		return OneTileRight(idx)
	case Down:
		return OneTileDown(idx)
	case Left:
		return OneTileLeft(idx)
	case UpLeft:
		return OneTileUpLeft(idx)
	case UpRight:
		return OneTileUpRight(idx)
	case DownLeft:
		return OneTileDownLeft(idx)
	case DownRight:
		return OneTileDownRight(idx)

	default:
		fmt.Println("unknown direction:", DirectionToText(dir))
		return -1
	}
}

func OneTileUp(idx int) int {
	return idx - globals.TilesW
}

func (m *Map) OneTileUp(idx int) Tile {
	return m.Tiles[OneTileUp(idx)]
}

func (m *Map) OneRailUp(idx int) Tile {
	return m.Rails[OneTileUp(idx)]
}

func OneTileDown(idx int) int {
	return idx + globals.TilesW
}

func (m *Map) OneTileDown(idx int) Tile {
	return m.Tiles[OneTileDown(idx)]
}

func (m *Map) OneRailDown(idx int) Tile {
	return m.Rails[OneTileDown(idx)]
}

func OneTileLeft(idx int) int {
	return idx - 1
}

func (m *Map) OneTileLeft(idx int) Tile {
	return m.Tiles[OneTileLeft(idx)]
}

func (m *Map) OneRailLeft(idx int) Tile {
	return m.Rails[OneTileLeft(idx)]
}

func OneTileRight(idx int) int {
	return idx + 1
}

func (m *Map) OneTileRight(idx int) Tile {
	return m.Tiles[OneTileRight(idx)]
}

func (m *Map) OneRailRight(idx int) Tile {
	return m.Rails[OneTileRight(idx)]
}

func OneTileUpLeft(idx int) int {
	return OneTileUp(OneTileLeft(idx))
}

func (m *Map) OneTileUpLeft(idx int) Tile {
	return m.Tiles[OneTileUp(OneTileLeft(idx))]
}

func OneTileUpRight(idx int) int {
	return OneTileUp(OneTileRight(idx))
}

func (m *Map) OneTileUpRight(idx int) Tile {
	return m.Tiles[OneTileUp(OneTileRight(idx))]
}

func OneTileDownLeft(idx int) int {
	return OneTileDown(OneTileLeft(idx))
}

func (m *Map) OneTileDownLeft(idx int) Tile {
	return m.Tiles[OneTileDown(OneTileLeft(idx))]
}

func OneTileDownRight(idx int) int {
	return OneTileDown(OneTileRight(idx))
}

func (m *Map) OneTileDownRight(idx int) Tile {
	return m.Tiles[OneTileDown(OneTileRight(idx))]
}

// NextIdxToDir returns the direction needed
// to traverse to the next position
func NextIdxToDir(idx, next int) (Direction, error) {
	if next == OneTileUp(idx) {
		return Up, nil
	}
	if next == OneTileDown(idx) {
		return Down, nil
	}
	if next == OneTileLeft(idx) {
		return Left, nil
	}
	if next == OneTileRight(idx) {
		return Right, nil
	}
	return Up, fmt.Errorf("idx %v and next %v not adjacent", idx, next)
}
