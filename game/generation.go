package game

import (
	"image/color"
	"math/rand"
	ch "projects/games/warf2/character"
	"projects/games/warf2/entity"
	e "projects/games/warf2/entity"
	j "projects/games/warf2/jobsystem"
	"projects/games/warf2/mouse"
	u "projects/games/warf2/ui"
	m "projects/games/warf2/worldmap"
)

/* --------------------------------------------------------------------------- */
/*                                     TODO                                    */
/* --------------------------------------------------------------------------- */
/* This is just a placeholder for map generation and/or loading at the moment. */
/* --------------------------------------------------------------------------- */

func tempGame() Game {
	worldmap := makeMap()
	generateTempMap(&worldmap)

	game := Game{
		WorldMap:  worldmap,
		JobSystem: &j.JobSystem{},
		Data:      &entity.Data{},

		time:        Time{Frame: 1},
		mouseSystem: mouse.System{},
		ui: u.UI{
			MouseMode: u.Element{
				X:     m.TileSize,
				Y:     m.TileSize*m.TilesH - m.TileSize,
				Color: color.White,
			},
		},
	}

	game.JobSystem.Map = &game.WorldMap
	for i := 0; i < 4; i++ {
		dwarf := placeNewDwarf(game.WorldMap)
		game.Dwarves = append(game.Dwarves, dwarf)
		game.JobSystem.Workers = append(game.JobSystem.Workers, &dwarf)
	}

	return game
}

func makeMap() m.Map {
	mp := &m.Map{}

	mp.Tiles = newTiles(mp, m.Ground)
	mp.SelectedTiles = newTiles(mp, m.None)
	mp.Items = newTiles(mp, m.None)

	return *mp
}

func newTiles(mp *m.Map, sprite int) []m.Tile {
	t := make([]m.Tile, m.TilesW*m.TilesH)

	for i := range t {
		t[i] = m.CreateTile(i, sprite, mp)
	}

	return t
}

func generateTempMap(mp *m.Map) {
	automata(mp)
	fillIslands(mp, true)
	fillIslands(mp, false)

	m.DrawHLine(mp, 0, m.TilesW, m.BoundarySolid)
	m.DrawHLine(mp, m.TilesT-m.TilesW-1, m.TilesW, m.BoundarySolid)

	m.DrawVLine(mp, 0, m.TilesH, m.BoundarySolid)
	m.DrawVLine(mp, m.TilesW-1, m.TilesH, m.BoundarySolid)

	mp.FixWalls()
}

func automata(mp *m.Map) {
	randomizeWalls(mp, 40)
	setInnerWalls(mp)

	for i := range mp.Tiles {
		neighbors := 0

		for _, st := range m.SurroundingTilesEight(i) {
			if m.IndexOutOfBounds(st.Idx, st.Dir) {
				continue
			}

			if m.IsAnyWall(mp.Tiles[st.Idx].Sprite) {
				neighbors++
			}
		}

		if neighbors > 3 {
			if rand.Intn(100) < 80 {
				mp.Tiles[i].Sprite = m.WallSolid
			}
		} else {
			if rand.Intn(100) < 80 {
				mp.Tiles[i].Sprite = m.Ground
			}
		}
	}

	setInnerWalls(mp)
}

func randomizeWalls(mp *m.Map, chance int) {
	for i := range mp.Tiles {
		risk := rand.Intn(100)
		if risk < chance {
			mp.Tiles[i].Sprite = m.WallSolid
		}
	}
}

// Inverse flips between filling walls (false) and filling ground (true).
func fillIslands(mp *m.Map, inverse bool) {
	island := 1
	for i, t := range mp.Tiles {
		if t.Island != 0 {
			continue
		}

		x, y := m.IdxToXY(i)

		if m.IsWall(t.Sprite) {
			floodFill(x, y, mp, island, func(idx int) bool {
				if !m.IsAnyWall(mp.Tiles[idx].Sprite) {
					return false
				}

				if mp.Tiles[idx].Island == island {
					return false
				}

				mp.Tiles[idx].Island = island
				return true
			})
			island++
		} else if m.IsGround(t.Sprite) {
			floodFill(x, y, mp, island, func(idx int) bool {
				if !m.IsGround(mp.Tiles[idx].Sprite) {
					return false
				}

				if mp.Tiles[idx].Island == island {
					return false
				}

				mp.Tiles[idx].Island = island
				return true
			})
			island++
		}
	}

	for currentIsland := 1; currentIsland <= island; currentIsland++ {
		var islandTiles []int
		for _, t := range mp.Tiles {
			if t.Island == currentIsland {
				islandTiles = append(islandTiles, t.Idx)
			}
		}

		if len(islandTiles) == 0 {
			continue
		}

		if len(islandTiles) <= 5 {
			for _, currentIslandIdx := range islandTiles {
				mp.Tiles[currentIslandIdx].Island = 0
				if inverse {
					mp.Tiles[currentIslandIdx].Sprite = m.Ground
				} else {
					mp.Tiles[currentIslandIdx].Sprite = m.WallSolid
				}
			}
		}
	}

	mp.ResetIslands()
}

func setInnerWalls(mp *m.Map) {
	m.DrawHLine(mp, m.TilesW+1, (m.TilesW-1)*2, m.WallSolid)
	m.DrawHLine(mp, m.TilesT-(m.TilesW*2)-1, m.TilesW, m.WallSolid)

	m.DrawVLine(mp, m.TilesW+1, m.TilesH-1, m.WallSolid)
	m.DrawVLine(mp, m.TilesW-2, m.TilesH-1, m.WallSolid)
}

// Inverse flips between filling walls (false) and filling ground (true).
func floodFill(x, y int, mp *m.Map, island int, predicate func(int) bool) {
	idx := m.XYToIdx(x, y)

	ok := predicate(idx)
	if !ok {
		return
	}

	if y > 0 {
		floodFill(x, y-1, mp, island, predicate)
	}
	if x > 0 {
		floodFill(x-1, y, mp, island, predicate)
	}
	if y < m.TilesH-1 {
		floodFill(x, y+1, mp, island, predicate)
	}
	if x < m.TilesW-1 {
		floodFill(x+1, y, mp, island, predicate)
	}
}

func placeNewDwarf(mp m.Map) ch.Dwarf {
	var availableSpots []int
	for i := range mp.Tiles {
		if m.IsGround(mp.Tiles[i].Sprite) {
			availableSpots = append(availableSpots, mp.Tiles[i].Idx)
		}
	}

	return ch.Dwarf{
		Entity: e.Entity{
			Sprite: rand.Intn(ch.DwarfTeal),
			Idx:    availableSpots[rand.Intn(len(availableSpots))],
		},
	}
}
