package game

import (
	"fmt"
	"math/rand"

	d "github.com/Holmqvist1990/WARF2/dwarf"
	j "github.com/Holmqvist1990/WARF2/jobservice"
	"github.com/Holmqvist1990/WARF2/mouse"
	rail "github.com/Holmqvist1990/WARF2/rail"
	r "github.com/Holmqvist1990/WARF2/room"
	"github.com/Holmqvist1990/WARF2/ui"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func GenerateGame(dwarves int, worldmap *m.Map) Game {
	game := Game{
		WorldMap:     worldmap,
		Rooms:        r.NewService(),
		JobService:   j.NewService(worldmap),
		DwarfService: d.NewService(),
		RailService:  rail.NewService(worldmap),

		time:        Time{Frame: 1},
		mouseSystem: mouse.NewSystem(),
		ui:          ui.GenerateUI(),
	}
	for i := 0; i < dwarves; i++ {
		addDwarfToGame(&game, game.DwarfService.RandomName())
	}
	return game
}

func emptyMap() *m.Map {
	return m.New()
}

func placeNewDwarf(mp *m.Map, name string) (*d.Dwarf, bool) {
	var availableSpots []int
	for i := range mp.Tiles {
		if m.IsGround(mp.Tiles[i].Sprite) {
			availableSpots = append(availableSpots, mp.Tiles[i].Idx)
		}
	}
	if len(availableSpots) == 0 {
		fmt.Println("generation.go:placeNewDwarf: no available spaces")
		return nil, false
	}
	startingPosition := availableSpots[rand.Intn(len(availableSpots))]
	return d.New(startingPosition, name), true
}

func addDwarfToGame(g *Game, name string) {
	dwarf, ok := placeNewDwarf(g.WorldMap, name)
	if !ok {
		fmt.Println("generation.go:addDwarfToGame: dwarf was nil")
		return
	}
	g.JobService.Workers = append(g.JobService.Workers, dwarf)
	g.JobService.AvailableWorkers = append(g.JobService.AvailableWorkers, dwarf)
}
