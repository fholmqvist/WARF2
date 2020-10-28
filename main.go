package main

import (
	"log"

	"github.com/beefsack/go-astar"
	"github.com/hajimehoshi/ebiten"

	g "projects/games/warf2/game"
	"projects/games/warf2/gmap"
	m "projects/games/warf2/gmap"
)

var _ astar.Pather = &gmap.Tile{}

func main() {
	g := g.NewGame(false)

	factor := 1
	ebiten.SetWindowSize(m.ScreenWidth*factor, m.ScreenHeight*factor)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(m.TPS)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
