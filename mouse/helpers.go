package mouse

import (
	m "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
)

// Most mouse modes share the same functionality:
// 	1. Handle the first click.
// 	2. Handle the dragging of mouse to select some range.

// These two functions below wrap this functionality
// and use lambdas to inject specific behaviour.

func clickFunctions(mp *m.Map, currentMousePos int, firstClick func(), dragClick func(*m.Map, int, int)) {
	if !hasClicked {
		firstClick()
		setHasClicked(currentMousePos)
	}
	if startPoint >= 0 {
		FuncOverRange(mp, currentMousePos, startPoint, dragClick)
	}
}

// FuncOverRange runs over a map
// between the start and end
// points against the supplied
// function.
func FuncOverRange(mp *m.Map, start, end int, f func(*m.Map, int, int)) {
	x1, y1, x2, y2 := TileRange(start, end)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			f(mp, x, y)
		}
	}
}

// TileRange returns a
// two-dimensional range
// between start and end,
// regardless of direction.
func TileRange(start, end int) (int, int, int, int) {
	x1, y1 := m.IdxToXY(start)
	x2, y2 := m.IdxToXY(end)

	if x1 > x2 {
		x1, x2 = x2, x1
	}

	if y1 > y2 {
		y1, y2 = y2, y1
	}

	return x1, y1, x2, y2
}

func mousePos() int {
	mx, my := ebiten.CursorPosition()
	mx, my = mx/m.TileSize, my/m.TileSize
	if mx >= m.TilesW {
		return -1
	}
	return mx + (my * m.TilesW)
}

func setHasClicked(currentMousePos int) {
	startPoint = currentMousePos
	hasClicked = true
}

func unsetHasClicked() {
	startPoint = -1
	hasClicked = false
}
