package game

import (
	"image/color"
	"math/rand"
	d "projects/games/warf2/dwarf"
	j "projects/games/warf2/jobsystem"
	"projects/games/warf2/mouse"
	rail "projects/games/warf2/railservice"
	u "projects/games/warf2/ui"
	m "projects/games/warf2/worldmap"
)

func GenerateGame(dwarves int, worldmap *m.Map) Game {
	game := Game{
		WorldMap:     *worldmap,
		JobService:   j.JobService{Map: worldmap},
		DwarfService: d.NewService(),
		RailService:  rail.RailService{Map: worldmap},

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
		addDwarfToGame(&game, game.DwarfService.RandomName())
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

func placeNewDwarf(mp m.Map, name string) d.Dwarf {
	var availableSpots []int
	for i := range mp.Tiles {
		if m.IsGround(mp.Tiles[i].Sprite) {
			availableSpots = append(availableSpots, mp.Tiles[i].Idx)
		}
	}
	startingPosition := availableSpots[rand.Intn(len(availableSpots))]
	return d.New(startingPosition, name)
}

func addDwarfToGame(g *Game, name string) {
	dwarf := placeNewDwarf(g.WorldMap, name)
	g.Dwarves = append(g.Dwarves, dwarf)
	g.JobService.Workers = append(g.JobService.Workers, &dwarf)
}
