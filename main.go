package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	g "projects/games/warf2/game"
	m "projects/games/warf2/worldmap"
)

func main() {
	g := g.NewGame()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	factor := 1
	ebiten.SetWindowSize(m.ScreenWidth*factor, m.ScreenHeight*factor)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(m.TPS)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
