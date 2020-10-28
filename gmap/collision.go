package gmap

import "fmt"

// NotColliding is the inverse of IsColliding below.
func NotColliding(mp *Map, idx int, dir Direction) bool {
	return !IsColliding(mp, idx, dir)
}

// IsColliding returns whether character will collide on next step.
func IsColliding(mp *Map, idx int, dir Direction) bool {
	if IndexOutOfBounds(idx, dir) {
		return true
	}

	t := mp.Tiles

	switch dir {

	case Up:
		if Blocking(t[OneTileUp(idx)].Sprite) {
			return true
		}
	case Right:
		if Blocking(t[OneTileRight(idx)].Sprite) {
			return true
		}
	case Down:
		if Blocking(t[OneTileDown(idx)].Sprite) {
			return true
		}
	case Left:
		if Blocking(t[OneTileLeft(idx)].Sprite) {
			return true
		}
	default:
		fmt.Println("unknown direction")
	}

	return false
}

// Blocking returns a boolean value
// answering whether the tile at the
// current index is Blocking movement.
func Blocking(tile int) bool {
	return !(IsNone(tile) || IsGround(tile))
}

// IndexOutOfBounds checks whether the given index
// plus a direction will produce an out of bounds exception.
func IndexOutOfBounds(idx int, dir Direction) bool {
	return overflowX(idx) || overflowUp(dir, idx) || overflowDown(dir, idx)
}

func overflowX(idx int) bool {
	return idx <= 0 || idx >= TilesT-1
}

func overflowUp(dir Direction, idx int) bool {
	return (dir == Up || dir == UpLeft || dir == UpRight) && idx < TilesW
}

func overflowDown(dir Direction, idx int) bool {
	return (dir == Down || dir == DownLeft || dir == DownRight) && idx > (TilesT-TilesW)
}
