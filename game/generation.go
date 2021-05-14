package game

import (
	"image/color"
	"math/rand"
	d "projects/games/warf2/dwarf"
	"projects/games/warf2/globals"
	j "projects/games/warf2/jobservice"
	"projects/games/warf2/mouse"
	rail "projects/games/warf2/railservice"
	"projects/games/warf2/ui"
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
				X:     globals.TileSize,
				Y:     globals.TileSize*globals.TilesH - globals.TileSize,
				Color: color.White,
			},
			MainMenu: ui.NewMainMenu(),
		},
	}
	for i := 0; i < dwarves; i++ {
		addDwarfToGame(&game, game.DwarfService.RandomName())
	}
	return game
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
