package game

import (
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/mouse"

	e "github.com/hajimehoshi/ebiten"
	i "github.com/hajimehoshi/ebiten/inpututil"
)

// SetMouseMode set the internal mousemode,
// and updates the equivalent UI text(s).
func (g *Game) SetMouseMode(mode mouse.Mode) {
	g.mouseSystem.Mode = mode
	state := mode.String()
	g.ui.MouseMode.Text = "GOWARF - " + state
}

func HandleKeyboard(g *Game) {
	if i.IsKeyJustPressed(e.KeySpace) {
		globals.PAUSE_GAME = true
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
		g.SetMouseMode(mouse.Farm)
	}
	if i.IsKeyJustPressed(e.Key4) {
		g.SetMouseMode(mouse.Library)
	}
	if i.IsKeyJustPressed(e.Key5) {
		g.SetMouseMode(mouse.Delete)
	}
}
