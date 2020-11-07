package game

import (
	c "projects/games/warf2/characters"
	h "projects/games/warf2/helpers"
	"projects/games/warf2/mouse"
	m "projects/games/warf2/worldmap"

	e "github.com/hajimehoshi/ebiten"
	i "github.com/hajimehoshi/ebiten/inpututil"
)

func handleKeyboard(g *Game) {
	handleTilesettingInput(g)
}

func handleTilesettingInput(g *Game) {
	mm := &g.mouseSystem.Mode
	mt := &g.ui.MouseMode.Text

	if i.IsKeyJustPressed(e.Key1) {
		*mm = mouse.Normal
		*mt = "GOWARF - WALL MODE"
	}

	if i.IsKeyJustPressed(e.Key2) {
		*mm = mouse.FloorTiles
		*mt = "GOWARF - FLOORTILE MODE"
	}

	if i.IsKeyJustPressed(e.Key3) {
		*mm = mouse.ResetFloor
		*mt = "GOWARF - RESET FLOORTILE MODE"
	}

	if i.IsKeyJustPressed(e.Key4) {
		*mm = mouse.PlaceItem
		*mt = "GOWARF - PLACE ITEM MODE"
	}

	if i.IsKeyJustPressed(e.Key5) {
		*mm = mouse.PlaceFurniture
		*mt = "GOWARF - PLACE FURNITURE MODE"
	}

	if i.IsKeyJustPressed(e.Key6) {
		*mm = mouse.RemoveItem
		*mt = "GOWARF - REMOVE ITEM MODE"
	}
}

// For debugging purposes using
// in-game moveable character.
func handleCharacterInput(chr *c.Character, mp *m.Map, t *h.Time) {
	if !t.TimeToMove() {
		return
	}

	w := &chr.Walker
	et := &chr.Entity

	if keyIsHeld(e.KeyUp) {
		w.Move(mp, et, m.Up)
		return
	}

	if keyIsHeld(e.KeyRight) {
		w.Move(mp, et, m.Right)
		return
	}

	if keyIsHeld(e.KeyDown) {
		w.Move(mp, et, m.Down)
		return
	}

	if keyIsHeld(e.KeyLeft) {
		w.Move(mp, et, m.Left)
		return
	}
}

func keyIsHeld(k e.Key) bool {
	return i.KeyPressDuration(k) > 0
}
