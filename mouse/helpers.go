package mouse

import (
	m "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
)

func firstClick(mp *m.Map, currentMousePos int, click func(), drags []func(*m.Map, int, int)) {
	if !hasClicked {
		click()
		setHasClicked(currentMousePos)
	}

	if startPoint >= 0 {
		mouseRange(mp, currentMousePos, startPoint, drags)
	}
}

func mouseRange(mp *m.Map, start, end int, fs []func(*m.Map, int, int)) {
	x1, y1, x2, y2 := tileRange(start, end)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			for _, f := range fs {
				f(mp, x, y)
			}
		}
	}

	previousStartPoint = startPoint
	previousEndPoint = end
}

func mousePos() int {
	mx, my := ebiten.CursorPosition()
	mx, my = mx/m.TileSize, my/m.TileSize

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
