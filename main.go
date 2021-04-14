package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"

	g "projects/games/warf2/game"
	m "projects/games/warf2/worldmap"
)

func main() {
	var arg string
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	g := g.NewGame(arg)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	factor := 1
	ebiten.SetWindowSize(m.ScreenWidth*factor, m.ScreenHeight*factor)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(m.TPS)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
