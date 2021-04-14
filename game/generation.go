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

func standardMap() *m.Map {
	mp := makeMap()
	mp.Automata()
	mp.FillIslands(true)
	mp.FillIslands(false)
	mp.CreateBoundaryWalls()
	mp.FixWalls()
	return mp
}

func emptyMap() *m.Map {
	mp := makeMap()
	mp.CreateBoundaryWalls()
	mp.FixWalls()
	return mp
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

func makeMap() *m.Map {
	mp := &m.Map{}
	mp.Tiles = newTiles(mp, m.Ground)
	mp.SelectedTiles = newTiles(mp, m.None)
	mp.Items = newTiles(mp, m.None)
	return mp
}

func newTiles(mp *m.Map, sprite int) []m.Tile {
	t := make([]m.Tile, m.TilesW*m.TilesH)
	for i := range t {
		t[i] = m.CreateTile(i, sprite, mp)
	}
	return t
}
