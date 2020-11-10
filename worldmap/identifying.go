package worldmap

import "fmt"

// GraphicName returns the
// name of the corresponding
// graphics constant based
// on its index.
func GraphicName(sprite int) string {
	switch sprite {
	case None:
		return "Transparent"
	case Ground:
		return "Ground"
	case BoundarySolid:
		return "BoundarySolid"
	case BoundaryExposed:
		return "BoundaryExposed"
	case WallSolid:
		return "WallSolid"
	case WallExposed:
		return "WallExposed"
	case WallSelectedSolid:
		return "WallSelectedSolid"
	case WallSelectedExposed:
		return "WallSelectedExposed"
	}
	return fmt.Sprintf("unknown graphic #%d", sprite)
}

// SurroundingTilesFour returns four
// adjacent tiles of a given index.
func SurroundingTilesFour(idx int) []TileDir {
	return []TileDir{
		{Idx: OneTileUp(idx), Dir: Up},
		{Idx: OneTileRight(idx), Dir: Right},
		{Idx: OneTileDown(idx), Dir: Down},
		{Idx: OneTileLeft(idx), Dir: Left}}
}

// SurroundingTilesEight returns eight
// adjacent tiles of a given index.
func SurroundingTilesEight(idx int) []TileDir {
	corners := []TileDir{
		{Idx: OneTileUpLeft(idx), Dir: UpLeft},
		{Idx: OneTileUpRight(idx), Dir: UpRight},
		{Idx: OneTileDownLeft(idx), Dir: DownLeft},
		{Idx: OneTileDownRight(idx), Dir: DownRight}}

	return append(SurroundingTilesFour(idx), corners...)
}

// IsNone returns if
// tile is of type None.
func IsNone(sprite int) bool {
	return sprite == None
}

// IsGround returns if
// tile is of type Ground.
func IsGround(sprite int) bool {
	return sprite == Ground
}

// IsExposed returns if
// tile is in the open.
func IsExposed(sprite int) bool {
	return !IsAnyWall(sprite)
}

// IsAnyWall returns if
// tile is any type of wall.
func IsAnyWall(sprite int) bool {
	return IsBoundary(sprite) || IsWall(sprite) || IsSelectedWall(sprite)
}

// IsBoundary returns if
// tile is of type Boundary.
func IsBoundary(sprite int) bool {
	return sprite >= BoundarySolid && sprite <= BoundaryExposed
}

// IsWall returns if
// tile is of type Wall.
func IsWall(sprite int) bool {
	return sprite >= WallSolid && sprite <= WallExposed
}

// IsSelectedWall returns if
// tile is of type SelectedWall.
func IsSelectedWall(sprite int) bool {
	return sprite >= WallSelectedSolid && sprite <= WallSelectedExposed
}

// IsWallOrSelected returns if
// tile is of type Wall or SelectedWall.
func IsWallOrSelected(sprite int) bool {
	return IsWall(sprite) || IsSelectedWall(sprite)
}

// IsFloorBrick returns if
// tile is of type FloorBrick.
func IsFloorBrick(sprite int) bool {
	return sprite >= FloorBricksOne && sprite <= FloorBricksTen
}
