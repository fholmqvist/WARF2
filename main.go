package main

import (
	"log"

	"github.com/beefsack/go-astar"
	"github.com/hajimehoshi/ebiten"

	g "projects/games/warf2/game"
	"projects/games/warf2/worldmap"
	m "projects/games/warf2/worldmap"
)

var _ astar.Pather = &worldmap.Tile{}

func main() {
	g := g.NewGame()

	factor := 1
	ebiten.SetWindowSize(m.ScreenWidth*factor, m.ScreenHeight*factor)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(m.TPS)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
