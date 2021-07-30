package worldmap

import (
	"math/rand"

	"github.com/Holmqvist1990/WARF2/globals"
)

func RandomFloorBrick() int {
	return rand.Intn(FloorBricks10-FloorBricks1+1) + FloorBricks1
}

func RandomWoodFloor() int {
	if rand.Intn(3) < 2 {
		return WoodFloor1
	}
	return rand.Intn(WoodFloor4-WoodFloor1+1) + WoodFloor1
}

// FloodFill finds an "island"
// of tiles based on the predicate
// function and sets the tiles island
// number to the given island number.
func FloodFill(x, y int, m *Map, island int, predicate func(int) bool) {
	idx := globals.XYToIdx(x, y)
	ok := predicate(idx)
	if !ok {
		return
	}
	if y > 0 {
		FloodFill(x, y-1, m, island, predicate)
	}
	if x > 0 {
		FloodFill(x-1, y, m, island, predicate)
	}
	if y < globals.TilesH-1 {
		FloodFill(x, y+1, m, island, predicate)
	}
	if x < globals.TilesW-1 {
		FloodFill(x+1, y, m, island, predicate)
	}
}

func FloodFillWalls(x, y int, m *Map, island int) {
	FloodFill(x, y, m, island, func(idx int) bool {
		if !IsAnyWall(m.Tiles[idx].Sprite) {
			return false
		}
		if m.Tiles[idx].Island == island {
			return false
		}
		m.Tiles[idx].Island = island
		return true
	})
}

func FloodFillGround(x, y int, m *Map, island int) {
	FloodFill(x, y, m, island, func(idx int) bool {
		if !IsGround(m.Tiles[idx].Sprite) {
			return false
		}

		if m.Tiles[idx].Island == island {
			return false
		}

		m.Tiles[idx].Island = island
		return true
	})
}

// Resets islands!
func (mp *Map) FloodFillRoom(x, y int, spriteGenerator func() int) Tiles {
	///////////////////////////////////
	// TODO
	// New feature:
	// Filling a room of the same type
	// should merge the two rooms,
	// extending the first.
	///////////////////////////////////
	island := 99
	tiles := []Tile{}
	FloodFill(x, y, mp, island, func(idx int) bool {
		if !IsGround(mp.Tiles[idx].Sprite) {
			return false
		}
		if mp.Tiles[idx].Island == island {
			return false
		}
		if IsDoorOpening(mp, idx) {
			return false
		}
		mp.Tiles[idx].Sprite = spriteGenerator()
		mp.Tiles[idx].Island = island
		tiles = append(tiles, mp.Tiles[idx])
		return true
	})
	mp.ResetIslands()
	return tiles
}

// Inverse flips between filling walls (false) and filling ground (true).
func (m *Map) FillIslands(inverse bool) {
	island := 1
	for i, t := range m.Tiles {
		if t.Island != 0 {
			continue
		}
		x, y := globals.IdxToXY(i)
		if IsWall(t.Sprite) {
			FloodFillWalls(x, y, m, island)
			island++
		} else if IsGround(t.Sprite) {
			FloodFillGround(x, y, m, island)
			island++
		}
	}
	for currentIsland := 1; currentIsland <= island; currentIsland++ {
		var islandTiles []int
		for _, t := range m.Tiles {
			if t.Island == currentIsland {
				islandTiles = append(islandTiles, t.Idx)
			}
		}
		if len(islandTiles) == 0 {
			continue
		}
		if len(islandTiles) <= 5 {
			for _, currentIslandIdx := range islandTiles {
				m.Tiles[currentIslandIdx].Island = 0
				if inverse {
					m.Tiles[currentIslandIdx].Sprite = Ground
				} else {
					m.Tiles[currentIslandIdx].Sprite = WallSolid
				}
			}
		}
	}
	m.ResetIslands()
}

func (m *Map) Automata() {
	m.RandomizeWalls(40)
	m.CreateOutmostWalls()
	for i := range m.Tiles {
		neighbors := 0
		for _, st := range SurroundingTilesEight(i) {
			if IndexOutOfBounds(st.Idx, st.Dir) {
				continue
			}
			if IsAnyWall(m.Tiles[st.Idx].Sprite) {
				neighbors++
			}
		}
		if neighbors > 3 {
			if rand.Intn(100) < 80 {
				m.Tiles[i].Sprite = WallSolid
			}
		} else {
			if rand.Intn(100) < 80 {
				m.Tiles[i].Sprite = Ground
			}
		}
	}
	m.CreateOutmostWalls()
}
