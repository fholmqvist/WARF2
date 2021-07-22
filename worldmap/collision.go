package worldmap

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/globals"
)

// NotColliding is the inverse of IsColliding below.
func NotColliding(mp *Map, idx int, dir Direction) bool {
	return !IsColliding(mp, idx, dir)
}

// IsColliding returns whether the given
// position is colliding with the entity
// that desires to move to it.
func IsColliding(mp *Map, current int, next Direction) bool {
	if IndexOutOfBounds(current, next) {
		return true
	}
	t, ok := mp.getTileByIndexAndDirection(current, next)
	if !ok {
		return true
	}
	itemT, ok := mp.getItemTileByIndexAndDirection(current, next)
	if !ok {
		return true
	}
	return Blocking(t, itemT)
}

// Blocking returns a boolean value
// answering whether the tile at the
// current index is Blocking movement.
func Blocking(t *Tile, itemT *Tile) bool {
	return IsAnyWall(t.Sprite) || globals.IsItemBlocking(itemT.Sprite)
}

// IndexOutOfBounds checks whether the given index
// plus a direction will produce an out of bounds exception.
func IndexOutOfBounds(idx int, dir Direction) bool {
	if outOfBounds(idx) {
		return true
	}
	switch dir {
	case Up:
		return overflowUp(idx)
	case UpLeft:
		return overflowUpLeft(idx)
	case Down:
		return overflowDown(idx)
	case DownLeft:
		return overflowDown(idx)
	case UpRight:
		return overflowUp(idx)
	case DownRight:
		return overflowDownRight(idx)
	case Left:
	case Right:
	default:
		fmt.Println("unknown direction:", DirectionToText(dir))
	}
	return false
}

func outOfBounds(idx int) bool {
	return idx <= 0 || idx >= globals.TilesT-1
}

func overflowUp(idx int) bool {
	return idx < globals.TilesW
}

func overflowDown(idx int) bool {
	return idx > (globals.TilesT - globals.TilesW)
}

func overflowUpLeft(idx int) bool {
	return idx < globals.TilesW+1
}

func overflowDownRight(idx int) bool {
	return idx > (globals.TilesT-globals.TilesW)-1
}
