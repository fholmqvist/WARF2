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

func SurroundingWallTilesFour(m *Map, idx int) []TileDir {
	tiles := SurroundingTilesFour(idx)
	wallTiles := []TileDir{}
	for _, t := range tiles {
		if IsAnyWall(m.Tiles[t.Idx].Sprite) {
			wallTiles = append(wallTiles, t)
		}
	}
	return wallTiles
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

func IsNone(sprite int) bool {
	return sprite == None
}

func IsGround(sprite int) bool {
	return sprite == Ground
}

func IsExposed(sprite int) bool {
	return !IsAnyWall(sprite)
}

func IsAnyWall(sprite int) bool {
	return IsBoundary(sprite) || IsWall(sprite) || IsSelectedWall(sprite)
}

func IsBoundary(sprite int) bool {
	return sprite >= BoundarySolid && sprite <= BoundaryExposed
}

func IsWall(sprite int) bool {
	return sprite >= WallSolid && sprite <= WallExposed
}

func IsSelectedWall(sprite int) bool {
	return sprite >= WallSelectedSolid && sprite <= WallSelectedExposed
}

func IsWallOrSelected(sprite int) bool {
	return IsWall(sprite) || IsSelectedWall(sprite)
}

func IsRail(sprite int) bool {
	return sprite >= Straight &&
		sprite <= Cross
}

func IsFloorBrick(sprite int) bool {
	return sprite >= FloorBricksOne && sprite <= FloorBricksTen
}

// If we're only surrounded by two walls,
// and those two walls are aligned,
// we are in a door opening.
func IsDoorOpening(m *Map, idx int) bool {
	if IsAnyWall(m.Tiles[idx].Sprite) {
		return false
	}
	tiles := SurroundingWallTilesFour(m, idx)
	if len(tiles) != 2 {
		return false
	}
	// Vertical opening.
	if tiles[0].Dir == Up && tiles[1].Dir == Down {
		return true
	}
	if tiles[0].Dir == Down && tiles[1].Dir == Up {
		return true
	}
	// Horizontal opening.
	if tiles[0].Dir == Left && tiles[1].Dir == Right {
		return true
	}
	if tiles[0].Dir == Right && tiles[1].Dir == Left {
		return true
	}
	return false
}

func (mp *Map) FindNearestDoorOpenings(x, y, island int) []Tile {
	tiles := []Tile{}
	FloodFill(x, y, mp, island, func(idx int) bool {
		if !IsGround(mp.Tiles[idx].Sprite) {
			return false
		}

		if mp.Tiles[idx].Island == island {
			return false
		}

		if IsDoorOpening(mp, idx) {
			tiles = append(tiles, mp.Tiles[idx])
			return false
		}

		mp.Tiles[idx].Island = island
		return true
	})
	return tiles
}
