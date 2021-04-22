package main

import (
	"fmt"
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

	logo()

	log.SetFlags(log.Lshortfile)

	g := g.NewGame(arg)

	factor := 1
	ebiten.SetWindowSize(m.ScreenWidth*factor, m.ScreenHeight*factor)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(m.TPS)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func logo() {
	lines := []string{
		"##########################",
		"########          ########",
		"########   WARF   ########",
		"########          ########",
		"##########################",
		"by Fredrik Holmqvist"
	}
	fmt.Println()
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
}
