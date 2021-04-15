package game

import (
	"image/color"
	"math/rand"
	d "projects/games/warf2/dwarf"
	"projects/games/warf2/entity"
	e "projects/games/warf2/entity"
	j "projects/games/warf2/jobsystem"
	"projects/games/warf2/mouse"
	u "projects/games/warf2/ui"
	m "projects/games/warf2/worldmap"
)

func GenerateGame(dwarves int, worldmap *m.Map) Game {
	game := Game{
		WorldMap:  *worldmap,
		JobSystem: j.JobSystem{Map: worldmap},
		Data:      entity.Data{},

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
	for i := 0; i < dwarves; i++ {
		dwarf := placeNewDwarf(game.WorldMap)
		game.Dwarves = append(game.Dwarves, dwarf)
		game.JobSystem.Workers = append(game.JobSystem.Workers, &dwarf)
	}
	return game
}

func normalMap() *m.Map {
	mp := m.New()
	mp.Automata()
	mp.FillIslands(true)
	mp.FillIslands(false)
	mp.CreateBoundaryWalls()
	mp.FixWalls()
	return mp
}

func boundariesMap() *m.Map {
	mp := m.New()
	mp.CreateBoundaryWalls()
	mp.FixWalls()
	return mp
}

func emptyMap() *m.Map {
	return m.New()
}

func placeNewDwarf(mp m.Map) d.Dwarf {
	var availableSpots []int
	for i := range mp.Tiles {
		if m.IsGround(mp.Tiles[i].Sprite) {
			availableSpots = append(availableSpots, mp.Tiles[i].Idx)
		}
	}
	return d.Dwarf{
		Entity: e.Entity{
			Sprite: rand.Intn(d.DwarfTeal),
			Idx:    availableSpots[rand.Intn(len(availableSpots))],
		},
	}
}

func debugLibrary(game *Game, offset int) {
	game.WorldMap.DrawOutline(6, 5+offset, 38, 14+offset, m.WallSolid)
	game.WorldMap.DrawOutline(24, 13+offset, 38, 22+offset, m.WallSolid)
	game.WorldMap.Tiles[620+m.TilesW*offset].Sprite = m.Ground
	for idx := 623 + m.TilesW*offset; idx <= 634+m.TilesW*offset; idx++ {
		game.WorldMap.Tiles[idx].Sprite = m.Ground
	}
	game.Rooms.AddLibrary(&game.WorldMap, 7, 7+offset)
}
