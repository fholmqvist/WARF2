package game

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/mouse"
	rail "github.com/Holmqvist1990/WARF2/railservice"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func gameFromArg(args []string) *Game {
	var game Game
	state := Gameplay

	globals.DEBUG = true

	switch args[0] {

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
		game.Rooms.AddLibrary(game.WorldMap, globals.XYToIdx(7, 7))
		game.WorldMap.FixWalls()
		addDwarfToGame(&game, "Test 1")
		addDwarfToGame(&game, "Test 2")
		d1 := game.JobService.Workers[0]
		d1.Characteristics.DesireToRead = 20
		d2 := game.JobService.Workers[1]
		d2.Characteristics.DesireToRead = 30

	case "storage":
		game = GenerateGame(0, m.BoundariesMap())
		mp := game.WorldMap
		mp.DrawOutline(5, 5, 10, 10, m.WallSolid)
		mp.DrawOutline(20, 5, 25, 10, m.WallSolid)
		s1 := room.NewStorage(mp, 6, 6)
		s2 := room.NewStorage(mp, 21, 6)
		game.Rooms.Storages = append(game.Rooms.Storages, *s1)
		game.Rooms.Storages = append(game.Rooms.Storages, *s2)
		ns, _, ok := game.Rooms.FindNearestStorage(mp, 1, 1)
		if !ok {
			panic(ok)
		}
		fmt.Println(ns.Center)

	case "walls":
		///////////////////////////////////////////////////////
		// Debugging and testing wall and floor fills.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, m.BoundariesMap())
		mp := game.WorldMap

		// Room 1.
		mp.DrawOutline(5, 5, 10, 10, m.WallSolid)
		mp.Tiles[globals.XYToIdx(5, 7)].Sprite = m.Ground
		mp.Tiles[globals.XYToIdx(7, 5)].Sprite = m.Ground

		// Room 2.
		mp.DrawOutline(12, 5, 24, 12, m.WallSolid)
		mp.Tiles[globals.XYToIdx(23, 8)].Sprite = m.Ground
		mp.Tiles[globals.XYToIdx(16, 11)].Sprite = m.Ground

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
		// Debugging pathfinding to wall digging
		// and carrying jobs.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, m.BoundariesMap())
		mp := game.WorldMap
		mp.DrawOutline(4, 4, 16, 11, m.WallSelectedSolid)
		mp.DrawOutline(5, 5, 10, 10, m.WallSelectedSolid)
		mp.DrawOutline(10, 5, 15, 10, m.WallSelectedSolid)
		for i := 0; i < 2; i++ {
			d := dwarf.New(282+i, fmt.Sprintf("test%v", i+1))
			game.JobService.Workers = append(game.JobService.Workers, d)
		}
		game.Rooms.AddStorage(mp, globals.XYToIdx(6, 6))

	case "wall-stress":
		///////////////////////////////////////////////////////
		// Stress test for digging jobs.
		///////////////////////////////////////////////////////
		mp := m.FilledMap()
		for offset := 2; offset < globals.TilesW-2; offset += 4 {
			mp.DrawSquareSprite(offset, 2, offset+2, globals.TilesH-4, m.WallSelectedExposed)
		}
		for offset := 2; offset < globals.TilesH-2; offset += 4 {
			mp.DrawSquareSprite(2, offset, globals.TilesW-4, offset+2, m.WallSelectedExposed)
		}
		mp.DrawSquareSprite(42, 26, 44, 28, m.Ground)
		mp.DrawSquareSprite(2, 2, 4, 4, m.Ground)
		mp.FixWalls()
		game = GenerateGame(128, mp)

	case "fill":
		///////////////////////////////////////////////////////
		// Debugging and testing wall selection.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, m.BoundariesMap())
		mp := game.WorldMap
		mp.DrawSquare(1, 1, globals.TilesW-1, globals.TilesH-1, m.WallSolid)
		mp.FixWalls()

	case "rails":
		///////////////////////////////////////////////////////
		// Debugging rails.
		///////////////////////////////////////////////////////
		game = GenerateGame(0, m.BoundariesMap())
		game.RailService.Carts = append(game.RailService.Carts, rail.NewCart(globals.XYToIdx(2, 2)))
		var halfCircle [][2]int
		for line := 2; line < globals.TilesW-2; line++ {
			halfCircle = append(halfCircle, [2]int{line, 2})
		}
		for line := 2; line < globals.TilesH-2; line++ {
			halfCircle = append(halfCircle, [2]int{globals.TilesW - 3, line})
		}
		game.RailService.PlaceRailsXY(halfCircle)
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
		f := func(g *Game) {
			mp := g.WorldMap
			cart := g.RailService.Carts[0]
			if len(cart.Path) > 0 {
				return
			}
			if cart.Idx == globals.XYToIdx(2, 2) {
				g.RailService.Carts[0].InitiateRide(mp, &mp.Rails[globals.XYToIdx(43, 29)])
			}
			if cart.Idx == globals.XYToIdx(43, 29) {
				cart.InitiateRide(mp, &mp.Rails[globals.XYToIdx(2, 2)])
			}
		}
		game.debugFunc = &f

	case "maintain":
		///////////////////////////////////////////////////////
		// Runs procedures that clean and maintain
		// generated files.
		///////////////////////////////////////////////////////
		maintenance()
		os.Exit(3)

	case "git":
		///////////////////////////////////////////////////////
		// Runs maintenance.
		// Adds changes to GIT with message.
		///////////////////////////////////////////////////////

		fmt.Println("Adding to GIT with comment:", args[1:])
		file := "./push_to_git.sh"
		f, _ := os.ReadFile(file)
		lines := strings.Split(string(f), "\n")
		lines[1] = "message=\"" + args[1]
		for _, arg := range args[2:] {
			lines[1] += " " + arg
		}
		lines[1] += "\""
		os.WriteFile(file, []byte(strings.Join(lines, "\n")), fs.FileMode(os.O_TRUNC))
		out, err := exec.Command("C:/Program Files/Git/usr/bin/sh.exe", file).Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(out, err)
		os.Exit(3)

	case "load":
		///////////////////////////////////////////////////////
		// Load saved game.
		///////////////////////////////////////////////////////
		game = loadGame()

	case "menu":
		///////////////////////////////////////////////////////
		// Main Menu.
		///////////////////////////////////////////////////////
		game = GenerateGame(4, m.NormalMap())
		state = MainMenu

	case "game":
		game = GenerateGame(4, m.NormalMap())

	default:
		game = GenerateGame(4, m.NormalMap())
		state = MainMenu
		globals.DEBUG = false

	}
	game.state = state
	game.SetMouseMode(mouse.Normal)
	return &game
}

func maintenance() {
	fmt.Println("Cleaning names...")
	ds := dwarf.NewService()
	ds.CleanNames()
	fmt.Println("Generating Todo file...")
	globals.GenerateTodos()
}
