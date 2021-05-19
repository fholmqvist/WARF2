package game

import (
	"fmt"
	"projects/games/warf2/mouse"

	e "github.com/hajimehoshi/ebiten"
	i "github.com/hajimehoshi/ebiten/inpututil"
)

// SetMouseMode set the internal mousemode,
// and updates the equivalent UI text(s).
func (g *Game) SetMouseMode(mode mouse.Mode) {
	g.mouseSystem.Mode = mode
	mt := &g.ui.MouseMode.Text
	state := ""

	switch mode {

	case mouse.Normal:
		state = "WALL MODE"

	case mouse.Storage:
		state = "STORAGE"

	case mouse.Library:
		state = "LIBRARY"

	case mouse.Delete:
		state = "DELETE"

	default:
		fmt.Println("no such mouse mode:", mode)
	}

	*mt = "GOWARF - " + state
}

func HandleKeyboard(g *Game) {
	if i.IsKeyJustPressed(e.KeySpace) {
		g.time.Stop()
	}
	handleTileSettingInput(g)
}

func handleTileSettingInput(g *Game) {
	if i.IsKeyJustPressed(e.Key1) {
		g.SetMouseMode(mouse.Normal)
	}
	if i.IsKeyJustPressed(e.Key2) {
		g.SetMouseMode(mouse.Storage)
	}
	if i.IsKeyJustPressed(e.Key3) {
		g.SetMouseMode(mouse.Library)
	}
	if i.IsKeyJustPressed(e.Key4) {
		g.SetMouseMode(mouse.Delete)
	}
}
