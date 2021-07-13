// Package game contains
// the root struct for all
// the data needed to run the game.
package game

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/globals"
	j "github.com/Holmqvist1990/WARF2/jobservice"
	"github.com/Holmqvist1990/WARF2/mouse"
	rail "github.com/Holmqvist1990/WARF2/railservice"
	"github.com/Holmqvist1990/WARF2/room"
	u "github.com/Holmqvist1990/WARF2/ui"
	m "github.com/Holmqvist1990/WARF2/worldmap"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

// Game wraps all data the needs to be accessible across domains.
type Game struct {
	/* ----------------------------- In-game objects ---------------------------- */

	WorldMap *m.Map
	Rooms    *room.Service

	/* ------------------------------ Loaded assets ----------------------------- */

	worldTiles *ebiten.Image
	dwarfTiles *ebiten.Image
	railTiles  *ebiten.Image
	itemTiles  *ebiten.Image
	uiTiles    *ebiten.Image
	font       font.Face

	/* ------------------------------- Interaction ------------------------------ */

	mouseSystem *mouse.System

	/* -------------------------------- Services -------------------------------- */

	JobService   *j.JobService
	DwarfService *dwarf.DwarfService
	RailService  *rail.RailService

	/* ------------------------------ Private state ----------------------------- */

	state GameState
	time  Time
	ui    u.UI

	// For injecting debugging
	// routines into the runtime.
	//
	// This function can access
	// anything the game can, so
	// it should be able to modify
	// any functionality that way.
	debugFunc *func(*Game)
}

// NewGame returns a pointer to an instantiated and initiated game.
func NewGame(args []string) *Game {
	game := initWithArgs(args)
	game.LoadAssets()
	game.SaveGame()
	return game
}

// Layout implementation for Ebiten interface.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return globals.ScreenWidth, globals.ScreenHeight
}

func (g *Game) LoadAssets() {
	var (
		tilesets = []*ebiten.Image{}
		paths    = []string{
			"art/world.png",
			"art/dwarf.png",
			"art/rail.png",
			"art/item.png",
			"art/ui.png",
		}
	)
	for _, path := range paths {
		tiles, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
		}
		tilesets = append(tilesets, tiles)
	}
	g.worldTiles = tilesets[0]
	g.dwarfTiles = tilesets[1]
	g.railTiles = tilesets[2]
	g.itemTiles = tilesets[3]
	g.uiTiles = tilesets[4]
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
