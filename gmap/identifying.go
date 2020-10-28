package gmap

import "fmt"

// GraphicName returns the name of the corresponding
// graphics constant based on its index.
func GraphicName(sprite int) string {
	switch sprite {
	case Transparent:
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
		{
			Idx: OneTileUp(idx),
			Dir: Up,
		},
		{
			Idx: OneTileRight(idx),
			Dir: Right,
		},
		{
			Idx: OneTileDown(idx),
			Dir: Down,
		},
		{
			Idx: OneTileLeft(idx),
			Dir: Left,
		},
	}
}

// SurroundingTilesEight returns eight
// adjacent tiles of a given index.
func SurroundingTilesEight(idx int) []TileDir {
	firstFour := SurroundingTilesFour(idx)
	corners := []TileDir{
		{
			Idx: OneTileUp(OneTileLeft(idx)),
			Dir: UpLeft,
		},
		{
			Idx: OneTileUp(OneTileRight(idx)),
			Dir: UpRight,
		},
		{
			Idx: OneTileDown(OneTileLeft(idx)),
			Dir: DownLeft,
		},
		{
			Idx: OneTileDown(OneTileRight(idx)),
			Dir: DownRight,
		},
	}

	return append(firstFour, corners...)
}

// IsNone returns whether tile is of type None.
func IsNone(tile int) bool {
	return tile == Transparent
}

// IsGround returns whether tile is of type Ground.
func IsGround(tile int) bool {
	return tile == Ground
}

// IsExposed returns whether tile is in the open.
func IsExposed(tile int) bool {
	return !(IsBoundary(tile) || IsWall(tile) || IsSelectedWall(tile))
}

// IsAnyWall returns whether tile is any type of wall.
func IsAnyWall(tile int) bool {
	return IsBoundary(tile) || IsWall(tile) || IsSelectedWall(tile)
}

// IsBoundary returns whether tile is of type Boundary.
func IsBoundary(tile int) bool {
	return tile >= BoundarySolid && tile <= BoundaryExposed
}

// IsWall returns whether tile is of type Wall.
func IsWall(tile int) bool {
	return tile >= WallSolid && tile <= WallExposed
}

// IsSelectedWall returns whether tile is of type SelectedWall.
func IsSelectedWall(tile int) bool {
	return tile >= WallSelectedSolid && tile <= WallSelectedExposed
}
