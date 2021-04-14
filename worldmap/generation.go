package worldmap

import (
	"math/rand"
)

var (
	floorTiles = []int{
		FloorBricksOne, FloorBricksTwo, FloorBricksThree,
		FloorBricksFour, FloorBricksFive, FloorBricksSix,
		FloorBricksSeven, FloorBricksEight, FloorBricksNine,
		FloorBricksTen,
	}
)

func (mp *Map) SetFloorTile(x, y int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}
	if IsAnyWall(tile.Sprite) {
		return
	}
	if IsFloorBrick(tile.Sprite) {
		return
	}
	tile.Sprite = RandomFloorBrick()
}

func (mp *Map) SetFloorTileIdx(idx int) {
	x, y := IdxToXY(idx)
	mp.SetFloorTile(x, y)
}

func RandomFloorBrick() int {
	return floorTiles[rand.Intn(len(floorTiles))]
}

// FloodFill finds an "island"
// of tiles based on the predicate
// function and sets the tiles island
// number to the given island number.
func FloodFill(x, y int, m *Map, island int, predicate func(int) bool) {
	idx := XYToIdx(x, y)
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
	if y < TilesH-1 {
		FloodFill(x, y+1, m, island, predicate)
	}
	if x < TilesW-1 {
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

// Inverse flips between filling walls (false) and filling ground (true).
func (m *Map) FillIslands(inverse bool) {
	island := 1
	for i, t := range m.Tiles {
		if t.Island != 0 {
			continue
		}
		x, y := IdxToXY(i)
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
