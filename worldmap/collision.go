package worldmap

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

	switch dir {

	case Up:
		t, ok := mp.GetTileByIndex(OneTileUp(idx))
		if !ok || Blocking(t.Sprite) {
			return true
		}
	case Right:
		t, ok := mp.GetTileByIndex(OneTileRight(idx))
		if !ok || Blocking(t.Sprite) {
			return true
		}
	case Down:
		t, ok := mp.GetTileByIndex(OneTileDown(idx))
		if !ok || Blocking(t.Sprite) {
			return true
		}
	case Left:
		t, ok := mp.GetTileByIndex(OneTileLeft(idx))
		if !ok || Blocking(t.Sprite) {
			return true
		}
	case UpLeft:
		t, ok := mp.GetTileByIndex(OneTileUpLeft(idx))
		if !ok || Blocking(t.Sprite) {
			return true
		}
	case UpRight:
		t, ok := mp.GetTileByIndex(OneTileUpRight(idx))
		if !ok || Blocking(t.Sprite) {
			return true
		}
	case DownLeft:
		t, ok := mp.GetTileByIndex(OneTileDownLeft(idx))
		if !ok || Blocking(t.Sprite) {
			return true
		}
	case DownRight:
		t, ok := mp.GetTileByIndex(OneTileDownRight(idx))
		if !ok || Blocking(t.Sprite) {
			return true
		}
	default:
		fmt.Println("unknown direction:", DirectionToText(dir))
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
	return idx <= 0 || idx >= TilesT-1
}

func overflowUp(idx int) bool {
	return idx < TilesW
}

func overflowDown(idx int) bool {
	return idx > (TilesT - TilesW)
}

func overflowUpLeft(idx int) bool {
	return idx < TilesW+1
}

func overflowDownRight(idx int) bool {
	return idx > (TilesT-TilesW)-1
}
