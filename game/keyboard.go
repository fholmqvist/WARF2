package game

import (
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/mouse"

	e "github.com/hajimehoshi/ebiten"
	i "github.com/hajimehoshi/ebiten/inpututil"
)

// SetMouseMode sets the internal mouse.Mode,
// and updates the equivalent UI text(s).
func (g *Game) SetMouseMode(mode mouse.Mode) {
	g.mouseSystem.Mode = mode
	g.ui.MouseMode.Text = "GOWARF - " + mode.String()
}

// Pausing, mouse.Mode.
func HandleKeyboard(g *Game) {
	if i.IsKeyJustPressed(e.KeySpace) {
		globals.GAME_PAUSED = !globals.GAME_PAUSED
		globals.ESC_MENU = false
	}
	if i.IsKeyJustPressed(e.KeyEscape) {
		globals.ESC_MENU = !globals.ESC_MENU
		globals.GAME_PAUSED = globals.ESC_MENU
	}
	if globals.GAME_PAUSED || globals.ESC_MENU {
		return
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
		g.SetMouseMode(mouse.SleepHall)
	}
	if i.IsKeyJustPressed(e.Key4) {
		g.SetMouseMode(mouse.Farm)
	}
	if i.IsKeyJustPressed(e.Key5) {
		g.SetMouseMode(mouse.Brewery)
	}
	if i.IsKeyJustPressed(e.Key6) {
		g.SetMouseMode(mouse.Library)
	}
	if i.IsKeyJustPressed(e.Key7) {
		g.SetMouseMode(mouse.Delete)
	}
}
