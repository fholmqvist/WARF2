package game

import (
	"fmt"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/globals"
	m "projects/games/warf2/worldmap"
	"time"
)

func gameFromArg(arg string) *Game {
	var game Game
	state := Gameplay

	globals.DEBUG = true

	switch arg {

	case "library":
		///////////////////////////////////////////////////////
		// Debugging and testing library generation.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, emptyMap())
		game.WorldMap.DrawOutline(6, 5, 38, 14, m.WallSolid)
		game.WorldMap.DrawOutline(24, 13, 38, 22, m.WallSolid)
		game.WorldMap.Tiles[252].Sprite = m.Ground
		game.WorldMap.Tiles[620].Sprite = m.Ground
		for idx := 623; idx <= 634; idx++ {
			game.WorldMap.Tiles[idx].Sprite = m.Ground
		}
		game.Rooms.AddLibrary(&game.WorldMap, m.XYToIdx(7, 7))
		game.WorldMap.FixWalls()
		addDwarfToGame(&game, "Test 1")
		addDwarfToGame(&game, "Test 2")
		d1 := game.Dwarves[0]
		d1.Characteristics.DesireToRead = 20
		d2 := game.Dwarves[1]
		d2.Characteristics.DesireToRead = 30

	case "walls":
		///////////////////////////////////////////////////////
		// Debugging and testing wall and floor fills.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, boundariesMap())
		mp := &game.WorldMap

		// Room 1.
		mp.DrawOutline(5, 5, 10, 10, m.WallSolid)
		mp.Tiles[m.XYToIdx(5, 7)].Sprite = m.Ground
		mp.Tiles[m.XYToIdx(7, 5)].Sprite = m.Ground

		// Room 2.
		mp.DrawOutline(12, 5, 24, 12, m.WallSolid)
		mp.Tiles[m.XYToIdx(23, 8)].Sprite = m.Ground
		mp.Tiles[m.XYToIdx(16, 11)].Sprite = m.Ground

		// Room 3.
		mp.DrawOutline(26, 5, 38, 12, m.WallSolid)
		mp.DrawOutline(32, 11, 38, 18, m.WallSolid)
		mp.Tiles[536].Sprite = m.Ground
		for idx := 539; idx <= 542; idx++ {
			game.WorldMap.Tiles[idx].Sprite = m.Ground
		}

		go func() {
			time.Sleep(time.Millisecond * 500)
			_ = mp.FloodFillRoom(6, 6, m.RandomFloorBrick)
			_ = mp.FloodFillRoom(13, 6, m.RandomFloorBrick)
			_ = mp.FloodFillRoom(27, 6, m.RandomFloorBrick)
			mp.FixWalls()
		}()

	case "wall-debug":
		///////////////////////////////////////////////////////
		// Debugging pathfinding to wall digging jobs.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, boundariesMap())
		mp := &game.WorldMap
		mp.DrawOutline(5, 5, 10, 10, m.WallSolid)
		mp.DrawOutline(10, 5, 15, 10, m.WallSolid)
		game.Dwarves = append(game.Dwarves, dwarf.New(328, "test"))
		game.JobService.Workers = append(game.JobService.Workers, &game.Dwarves[0])
		mp.Tiles[331].Sprite = m.WallSelectedSolid
		mp.Tiles[332].Sprite = m.WallSelectedSolid

	case "fill":
		///////////////////////////////////////////////////////
		// Debugging and testing wall selection.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, boundariesMap())
		mp := &game.WorldMap
		mp.DrawSquare(1, 1, globals.TilesW-1, globals.TilesH-1, m.WallSolid)
		mp.FixWalls()

	case "rails":
		///////////////////////////////////////////////////////
		// Debugging rails.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, boundariesMap())

		game.RailService.PlaceRailsXY([][2]int{
			{8, 10},
			{8, 11},
			{9, 11},
			{10, 11},
			{11, 11},
			{11, 12},
			{11, 13},
			{11, 14},
			{10, 14},
			{9, 14},
			{9, 13},
			{8, 13},
			{7, 13},
			{7, 14},
			{7, 15},
			{6, 15},
		})

		game.RailService.PlaceRailsXY([][2]int{
			{8, 20},
			{8, 21},
			{7, 21},
			{9, 21},
			{8, 22},
		})

	case "clean":
		///////////////////////////////////////////////////////
		// Various cleaning and correcting files.
		///////////////////////////////////////////////////////
		fmt.Println("Cleaning names...")
		ds := dwarf.NewService()
		ds.CleanNames()
		return nil

	case "load":
		///////////////////////////////////////////////////////
		///////////////////////////////////////////////////////
		game = loadGame()

	case "menu":
		///////////////////////////////////////////////////////
		// Main Menu.
		///////////////////////////////////////////////////////
		game = GenerateGame(4, normalMap())
		state = MainMenu

	default:
		///////////////////////////////////////////////////////
		// TODO:
		// Standard game, skipping menu.
		// On release this should default to menu.
		///////////////////////////////////////////////////////
		game = GenerateGame(4, normalMap())
		globals.DEBUG = false

	}
	game.state = state
	return &game
}
