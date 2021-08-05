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

func NeighTileFour(idx int) []int {
	return []int{
		OneUp(idx), OneRight(idx),
		OneDown(idx), OneLeft(idx),
	}
}

func NeighTileDirFour(idx int) []TileDir {
	return []TileDir{
		{Idx: OneUp(idx), Dir: Up},
		{Idx: OneRight(idx), Dir: Right},
		{Idx: OneDown(idx), Dir: Down},
		{Idx: OneLeft(idx), Dir: Left}}
}

func NeighWallTileDirFour(m *Map, idx int) []TileDir {
	tiles := NeighTileDirFour(idx)
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
		{Idx: OneUpLeft(idx), Dir: UpLeft},
		{Idx: OneUpRight(idx), Dir: UpRight},
		{Idx: OneDownLeft(idx), Dir: DownLeft},
		{Idx: OneDownRight(idx), Dir: DownRight}}

	return append(NeighTileDirFour(idx), corners...)
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

func IsStorageFloorBrick(sprite int) bool {
	return sprite >= StorageFloor1 && sprite <= StorageFloor10
}

func IsLibraryWoodFloor(sprite int) bool {
	return sprite >= LibraryFloor1 && sprite < LibraryFloor4
}

func IsSleepHallWoodFloor(sprite int) bool {
	return sprite == SleepHallFloor
}

// If we're only surrounded by two walls,
// and those two walls are aligned,
// we are in a door opening.
func IsDoorOpening(mp *Map, idx int) bool {
	if IsAnyWall(mp.Tiles[idx].Sprite) {
		return false
	}
	tiles := NeighWallTileDirFour(mp, idx)
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

func IsNextToDoorOpening(mp *Map, idx int) bool {
	if IsDoorOpening(mp, OneUp(idx)) ||
		IsDoorOpening(mp, OneDown(idx)) ||
		IsDoorOpening(mp, OneLeft(idx)) ||
		IsDoorOpening(mp, OneRight(idx)) ||
		IsAnyWall(mp.OneTileDown(idx).Sprite) ||
		IsAnyWall(mp.OneTileDownLeft(idx).Sprite) ||
		IsAnyWall(mp.OneTileDownRight(idx).Sprite) {
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
