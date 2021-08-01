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

func TileDirsToIdxs(t []TileDir) []int {
	idxs := make([]int, len(t))
	for i := range t {
		idxs = append(idxs, t[i].Idx)
	}
	return idxs
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
		return OneUp(idx)
	case Right:
		return OneRight(idx)
	case Down:
		return OneDown(idx)
	case Left:
		return OneLeft(idx)
	case UpLeft:
		return OneUpLeft(idx)
	case UpRight:
		return OneUpRight(idx)
	case DownLeft:
		return OneDownLeft(idx)
	case DownRight:
		return OneDownRight(idx)
	default:
		fmt.Println("unknown direction:", DirectionToText(dir))
		return -1
	}
}

func OneUp(idx int) int {
	return idx - globals.TilesW
}

func (m *Map) OneTileUp(idx int) Tile {
	return m.Tiles[OneUp(idx)]
}

func (m *Map) OneRailUp(idx int) Tile {
	return m.Rails[OneUp(idx)]
}

func OneDown(idx int) int {
	return idx + globals.TilesW
}

func (m *Map) OneTileDown(idx int) Tile {
	return m.Tiles[OneDown(idx)]
}

func (m *Map) OneRailDown(idx int) Tile {
	return m.Rails[OneDown(idx)]
}

func OneLeft(idx int) int {
	return idx - 1
}

func (m *Map) OneTileLeft(idx int) Tile {
	return m.Tiles[OneLeft(idx)]
}

func (m *Map) OneRailLeft(idx int) Tile {
	return m.Rails[OneLeft(idx)]
}

func OneRight(idx int) int {
	return idx + 1
}

func (m *Map) OneTileRight(idx int) Tile {
	return m.Tiles[OneRight(idx)]
}

func (m *Map) OneRailRight(idx int) Tile {
	return m.Rails[OneRight(idx)]
}

func OneUpLeft(idx int) int {
	return OneUp(OneLeft(idx))
}

func (m *Map) OneTileUpLeft(idx int) Tile {
	return m.Tiles[OneUp(OneLeft(idx))]
}

func OneUpRight(idx int) int {
	return OneUp(OneRight(idx))
}

func (m *Map) OneTileUpRight(idx int) Tile {
	return m.Tiles[OneUp(OneRight(idx))]
}

func OneDownLeft(idx int) int {
	return OneDown(OneLeft(idx))
}

func (m *Map) OneTileDownLeft(idx int) Tile {
	return m.Tiles[OneDown(OneLeft(idx))]
}

func OneDownRight(idx int) int {
	return OneDown(OneRight(idx))
}

func (m *Map) OneTileDownRight(idx int) Tile {
	return m.Tiles[OneDown(OneRight(idx))]
}

// NextIdxToDir returns the direction needed
// to traverse to the next position
func NextIdxToDir(idx, next int) (Direction, error) {
	if next == OneUp(idx) {
		return Up, nil
	}
	if next == OneDown(idx) {
		return Down, nil
	}
	if next == OneLeft(idx) {
		return Left, nil
	}
	if next == OneRight(idx) {
		return Right, nil
	}
	return Up, nextIdxToDirError(idx, next)
}

func nextIdxToDirError(idx, next int) error {
	x1, y1 := globals.IdxToXY(idx)
	x2, y2 := globals.IdxToXY(next)
	return fmt.Errorf("NextIdxToDir error:\n\tidx %v and next %v not adjacent, vdiff %v, hdiff %v", idx, next, y1-y2, x1-x2)
}
