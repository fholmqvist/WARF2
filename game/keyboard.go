package game

import (
	"projects/games/warf2/mouse"

	e "github.com/hajimehoshi/ebiten"
	i "github.com/hajimehoshi/ebiten/inpututil"
)

// SetMouseMode set the internal mousemode,
// and updates the equivalent UI text(s).
func (g *Game) SetMouseMode(mode mouse.Mode) {
	g.mouseSystem.Mode = mode
	mt := &g.ui.MouseMode.Text

	switch mode {

	case mouse.Normal:
		*mt = "GOWARF - WALL MODE"

	case mouse.FloorTiles:
		*mt = "GOWARF - FLOORTILE MODE"

	case mouse.ResetFloor:
		*mt = "GOWARF - RESET FLOORTILE MODE"

	case mouse.PlaceItem:
		*mt = "GOWARF - PLACE ITEM MODE"

	case mouse.PlaceFurniture:
		*mt = "GOWARF - PLACE FURNITURE MODE"

	case mouse.RemoveItem:
		*mt = "GOWARF - REMOVE ITEM MODE"
	}
}

func handleKeyboard(g *Game) {
	handleTilesettingInput(g)
}

func handleTilesettingInput(g *Game) {
	if i.IsKeyJustPressed(e.Key1) {
		g.SetMouseMode(mouse.Normal)
	}
	if i.IsKeyJustPressed(e.Key2) {
		g.SetMouseMode(mouse.FloorTiles)
	}
	if i.IsKeyJustPressed(e.Key3) {
		g.SetMouseMode(mouse.ResetFloor)
	}
	if i.IsKeyJustPressed(e.Key4) {
		g.SetMouseMode(mouse.PlaceItem)
	}
	if i.IsKeyJustPressed(e.Key5) {
		g.SetMouseMode(mouse.PlaceFurniture)
	}
	if i.IsKeyJustPressed(e.Key6) {
		g.SetMouseMode(mouse.RemoveItem)
	}
	if i.IsKeyJustPressed(e.Key7) {
		g.SetMouseMode(mouse.Library)
	}
}
