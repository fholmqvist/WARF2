package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten"

	g "projects/games/warf2/game"
	"projects/games/warf2/globals"
)

func main() {
	logo()
	log.SetFlags(log.Lshortfile)
	var arg string
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	game := g.NewGame(arg)
	if game == nil {
		return
	}
	if len(os.Args) > 2 {
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("speed variable: %v", err)
		}
		g.FramesToMove = i
	}
	factor := 1
	ebiten.SetWindowSize(globals.ScreenWidth*factor, globals.ScreenHeight*factor)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(globals.TPS)
	if err := ebiten.RunGame(game); err != nil {
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
		"by Fredrik Holmqvist",
	}
	fmt.Println()
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
}
