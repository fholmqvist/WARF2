package game

import (
	c "projects/games/warf2/characters"
	h "projects/games/warf2/helpers"
	m "projects/games/warf2/worldmap"

	e "github.com/hajimehoshi/ebiten"
	i "github.com/hajimehoshi/ebiten/inpututil"
)

func handleKeyboard(g *Game) {
	handleTilesettingInput(g)
	handleCharacterInput(g.testChar, &g.WorldMap, &g.time)
}

func handleTilesettingInput(g *Game) {
	mm := &g.mouseMode
	mt := &g.ui.MouseMode.Text

	if i.IsKeyJustPressed(e.KeyEscape) {
		*mm = None
		*mt = "GOWARF"
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
