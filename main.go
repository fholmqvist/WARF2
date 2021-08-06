package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten"

	g "github.com/Holmqvist1990/WARF2/game"
	"github.com/Holmqvist1990/WARF2/globals"
)

var zoom = 1

func main() {
	log.SetFlags(log.Lshortfile)
	args := handleArgs()
	printLogo(args)
	var game *g.Game
	if len(args) > 0 {
		game = g.NewGame(args)

	} else {
		game = g.NewGame([]string{""})
	}
	ebiten.SetWindowSize(globals.ScreenWidth*zoom, globals.ScreenHeight*zoom)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(globals.TPS)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func handleArgs() []string {
	g.GAME_SPEED = g.SUPER
	for _, arg := range os.Args[1:] {
		switch strings.ToLower(arg) {
		case "normal":
			g.GAME_SPEED = g.NORMAL
		case "fast":
			g.GAME_SPEED = g.FAST
		case "super":
			g.GAME_SPEED = g.SUPER
		case "pause":
			globals.GAME_PAUSED = true
		}
	}
	return os.Args[1:]
}

func printLogo(args []string) {
	lines := []string{
		"##########################",
		"########          ########",
		"########   WARF   ########",
		"########          ########",
		"##########################",
		"by Fredrik Holmqvist",
		"",
	}
	if len(args) > 0 {
		lines = append(lines, fmt.Sprintf("Running with args: %v.", args))
	}
	fmt.Println()
	for _, line := range lines {
		fmt.Println(line)
	}
}
