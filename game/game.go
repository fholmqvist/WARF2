// Package game contains
// the root struct for all
// the data needed to run the game.
package game

import (
	"io/ioutil"
	"log"
	"os"

	"projects/games/warf2/dwarf"
	"projects/games/warf2/globals"
	j "projects/games/warf2/jobservice"
	"projects/games/warf2/mouse"
	rail "projects/games/warf2/railservice"
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
	Rooms    room.Service

	/* ------------------------------ Loaded assets ----------------------------- */

	worldTiles *ebiten.Image
	dwarfTiles *ebiten.Image
	railTiles  *ebiten.Image
	itemTiles  *ebiten.Image
	font       font.Face

	/* ------------------------------- Interaction ------------------------------ */

	mouseSystem mouse.System

	/* -------------------------------- Services -------------------------------- */

	JobService   j.JobService
	DwarfService dwarf.DwarfService
	RailService  rail.RailService

	/* ------------------------------ Private state ----------------------------- */

	state GameState
	time  Time
	ui    u.UI

	// For injecting debugging
	// routines into the runtime.
	debugFunc *func(*Game)
}

// NewGame returns a pointer to an instantiated and initiated game.
func NewGame(arg string) *Game {
	game := gameFromArg(arg)
	game.LoadAssets()
	game.SaveGame()
	return game
}

// Layout implementation for Ebiten interface.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return globals.ScreenWidth, globals.ScreenHeight
}

func (g *Game) LoadAssets() {
	worldTiles, _, err := ebitenutil.NewImageFromFile("art/world.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	g.worldTiles = worldTiles
	dwarfTiles, _, err := ebitenutil.NewImageFromFile("art/dwarf.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	g.dwarfTiles = dwarfTiles
	railTiles, _, err := ebitenutil.NewImageFromFile("art/rail.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	g.railTiles = railTiles
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
	g.font = truetype.NewFace(tt, &truetype.Options{
		Size:    8,
		Hinting: font.HintingFull,
	})
}
