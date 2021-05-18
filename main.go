package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten"

	g "projects/games/warf2/game"
	"projects/games/warf2/globals"
)

var zoom = 1

func main() {
	logo()
	log.SetFlags(log.Lshortfile)
	arg := handleArgs()
	game := g.NewGame(arg)
	ebiten.SetWindowSize(globals.ScreenWidth*zoom, globals.ScreenHeight*zoom)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(globals.TPS)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func handleArgs() string {
	var arg string
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	if len(os.Args) > 2 {
		var speed int
		switch strings.ToLower(os.Args[2]) {
		case "normal":
			speed = g.NORMAL
		case "fast":
			speed = g.FAST
		case "super":
			speed = g.SUPER
		}
		g.FramesToMove = speed
	}
	return arg
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
