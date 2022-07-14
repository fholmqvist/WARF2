package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten"

	g "github.com/Holmqvist1990/WARF2/game"
	gl "github.com/Holmqvist1990/WARF2/globals"
)

var zoom = 1

func main() {
	log.SetFlags(log.Lshortfile)
	args := handleArgs()
	printLogo(args)
	game := g.NewGame(args)
	ebiten.SetWindowSize(gl.ScreenWidth*zoom, gl.ScreenHeight*zoom)
	ebiten.SetWindowTitle("GOWARF")
	ebiten.SetMaxTPS(gl.TPS)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func handleArgs() []string {
	g.GAME_SPEED = g.SUPER
	if len(os.Args) < 2 {
		return []string{"default"}
	}
	for _, arg := range os.Args[1:] {
		switch strings.ToLower(arg) {
		case "normal":
			g.GAME_SPEED = g.NORMAL
		case "fast":
			g.GAME_SPEED = g.FAST
		case "super":
			g.GAME_SPEED = g.SUPER
		case "pause":
			gl.GAME_PAUSED = true
		}
	}
	return os.Args[1:]
}

func printLogo(args []string) {
	lines := `##########################
########          ########
########   WARF   ########
########          ########
##########################
by Fredrik Holmqvist`
	if len(args) > 0 {
		lines += fmt.Sprintf("\nRunning with args: %v.\n", args)
	}
	fmt.Println(lines)
}
