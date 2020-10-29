package game

import (
	"image/color"
	"io/ioutil"
	"log"
	"os"

	ch "projects/games/warf2/characters"
	e "projects/games/warf2/entity"
	h "projects/games/warf2/helpers"
	j "projects/games/warf2/jobsystem"
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

	WorldMap  m.Map
	JobSystem *j.JobSystem

	/* ------------------------------ Loaded assets ----------------------------- */

	tilesWorld   *ebiten.Image
	tilesDwarves *ebiten.Image
	gameFont     font.Face

	/* ----------------------------- External state ----------------------------- */

	Data *e.Data

	/* ----------------------------- Internal state ----------------------------- */

	time      h.Time
	debug     bool
	mouseMode MouseMode
	mousePos  int
	ui        u.UI

	/* ------------------------------ Experimental ------------------------------ */

	testChar *ch.Character
}

// NewGame returns a pointer to an instantiated and initiated game.
func NewGame(debug bool) *Game {
	worldmap := makeMap()
	generateTempMap(&worldmap)

	game := Game{
		WorldMap:  worldmap,
		JobSystem: &j.JobSystem{},

		Data: &e.Data{},

		time:      h.Time{Frame: 1},
		debug:     debug,
		mouseMode: None,
		ui: u.UI{
			MouseMode: u.Element{
				Text:  "GOWARF",
				X:     m.TileSize,
				Y:     m.TileSize*m.TilesH - m.TileSize,
				Color: color.White,
			},
		},

		testChar: testChar(worldmap),
	}

	game.JobSystem.Map = &game.WorldMap
	game.JobSystem.Workers = append(game.JobSystem.Workers, game.testChar)

	loadAssets(&game)

	return &game
}

// Layout implementation for Ebiten interface.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return m.ScreenWidth, m.ScreenHeight
}

func loadAssets(g *Game) {
	/* --------------------------------- Tileset -------------------------------- */

	worldTiles, _, err := ebitenutil.NewImageFromFile("art/world.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	g.tilesWorld = worldTiles

	dwarfTiles, _, err := ebitenutil.NewImageFromFile("art/dwarf.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	g.tilesDwarves = dwarfTiles

	/* ---------------------------------- Font ---------------------------------- */

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
