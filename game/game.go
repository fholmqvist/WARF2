// Package game contains
// the root struct for all
// the data needed to run the game.
package game

import (
	"io/ioutil"
	"log"
	"os"

	"projects/games/warf2/character"
	"projects/games/warf2/entity"
	j "projects/games/warf2/jobsystem"
	"projects/games/warf2/mouse"
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
	Dwarves  []character.Dwarf

	/* ------------------------------ Loaded assets ----------------------------- */

	worldTiles *ebiten.Image
	dwarfTiles *ebiten.Image
	itemTiles  *ebiten.Image
	gameFont   font.Face

	/* ------------------------------ Public state ------------------------------ */

	JobSystem *j.JobSystem
	Data      *entity.Data

	/* ------------------------------- Interaction ------------------------------ */

	mouseSystem mouse.System

	/* ------------------------------ Private state ----------------------------- */

	time  Time
	debug bool
	ui    u.UI
}

// NewGame returns a pointer to an instantiated and initiated game.
func NewGame() *Game {
	game := tempGame()
	//game := loadGame()

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
