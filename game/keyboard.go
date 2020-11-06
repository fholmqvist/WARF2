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

	if i.IsKeyJustPressed(e.KeyEscape) {
		*mm = mouse.Normal
		*mt = "GOWARF"
	}

	if i.IsKeyJustPressed(e.Key1) {
		*mm = mouse.FloorTiles
		*mt = "GOWARF - FLOOR TILES"
	}

	if i.IsKeyJustPressed(e.Key2) {
		*mm = mouse.ResetFloor
		*mt = "GOWARF - RESET FLOOR TILES"
	}
}

func handleCharacterInput(chr *c.Character, mp *m.Map, t *h.Time) {
	if !t.TimeToMove() {
		return
	}

	w := &chr.Walker
	et := &chr.Entity

	if key(e.KeyUp) {
		w.Move(mp, et, m.Up)
	} else if key(e.KeyRight) {
		w.Move(mp, et, m.Right)
	} else if key(e.KeyDown) {
		w.Move(mp, et, m.Down)
	} else if key(e.KeyLeft) {
		w.Move(mp, et, m.Left)
	}
}

func key(k e.Key) bool {
	return i.KeyPressDuration(k) > 0
}
