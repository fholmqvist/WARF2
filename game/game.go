// Package game contains
// the root struct for all
// the data needed to run the game.
package game

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"projects/games/warf2/dwarf"
	j "projects/games/warf2/jobsystem"
	"projects/games/warf2/mouse"
	"projects/games/warf2/room"
	u "projects/games/warf2/ui"
	m "projects/games/warf2/worldmap"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

// Game wraps all data the needs to be accessible across domains.
type Game struct {
	/* ----------------------------- In-game objects ---------------------------- */

	WorldMap m.Map
	Dwarves  []dwarf.Dwarf
	Rooms    room.System

	/* ------------------------------ Loaded assets ----------------------------- */

	worldTiles *ebiten.Image
	dwarfTiles *ebiten.Image
	itemTiles  *ebiten.Image
	gameFont   font.Face

	/* ------------------------------- Interaction ------------------------------ */

	mouseSystem mouse.System

	/* -------------------------------- Services -------------------------------- */

	JobService   j.JobService
	DwarfService dwarf.DwarfService

	/* ------------------------------ Private state ----------------------------- */

	time  Time
	debug bool
	ui    u.UI
}

// NewGame returns a pointer to an instantiated and initiated game.
func NewGame(arg string) *Game {
	var game Game

	switch arg {

	case "library":
		// Debugging and testing library generation.
		game = GenerateGame(0, emptyMap())
		game.WorldMap.DrawOutline(6, 5, 38, 14, m.WallSolid)
		game.WorldMap.DrawOutline(24, 13, 38, 22, m.WallSolid)
		game.WorldMap.Tiles[252].Sprite = m.Ground
		game.WorldMap.Tiles[620].Sprite = m.Ground
		for idx := 623; idx <= 634; idx++ {
			game.WorldMap.Tiles[idx].Sprite = m.Ground
		}
		game.Rooms.AddLibrary(&game.WorldMap, 7, 7)
		game.WorldMap.FixWalls()
		addDwarfToGame(&game, "Test 1")
		addDwarfToGame(&game, "Test 2")
		d1 := game.Dwarves[0]
		d1.Characteristics.DesireToRead = 20
		d2 := game.Dwarves[1]
		d2.Characteristics.DesireToRead = 30

	case "walls":
		// Debugging and testing wall and floor fills.
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

	case "fill":
		// Debugging and testing wall selection.
		game = GenerateGame(0, boundariesMap())
		mp := &game.WorldMap
		mp.DrawSquare(1, 1, m.TilesW-1, m.TilesH-1, m.WallSolid)
		mp.FixWalls()

	case "clean":
		fmt.Println("Cleaning names...")
		ds := dwarf.NewService()
		ds.CleanNames()
		return nil

	case "load":
		game = loadGame()

	default:
		game = GenerateGame(4, normalMap())

	}
	game.SetMouseMode(mouse.Normal)

	loadAssets(&game)

	game.saveGame()

	return &game
}

// Layout implementation for Ebiten interface.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return m.ScreenWidth, m.ScreenHeight
}

func loadAssets(g *Game) {
	// Setting worldTiles.
	worldTiles, _, err := ebitenutil.NewImageFromFile("art/world.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	g.worldTiles = worldTiles

	// Setting dwarfTiles.
	dwarfTiles, _, err := ebitenutil.NewImageFromFile("art/dwarf.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	g.dwarfTiles = dwarfTiles

	// Setting itemTiles.
	itemTiles, _, err := ebitenutil.NewImageFromFile("art/item.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	g.itemTiles = itemTiles

	setFont(g)
}

func setFont(g *Game) {
	f, err := os.Open("art/barcade_brawl.ttf")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	tt, err := truetype.Parse(b)
	if err != nil {
		log.Fatalf("could not parse truetype: %v", err)
	}

	g.gameFont = truetype.NewFace(tt, &truetype.Options{
		Size:    8,
		Hinting: font.HintingFull,
	})
}
