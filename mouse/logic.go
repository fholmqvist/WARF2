package mouse

import (
	"fmt"

	m "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// This cluster of variables
// help with (de)selecting walls.
var startPoint = -1
var endPoint = -1
var hasClicked = false
var firstClickedSprite = -1

// Remembering last frame
// in order to reset selected
// tiles without having to
// redraw the entire screen.
var previousStartPoint = -1
var previousEndPoint = -1

// Handle all the mouse interactivity.
func (s *System) Handle(mp *m.Map) {
	s.mouseHover(mp)

	idx := mousePos()

	if idx < 0 || idx > m.TilesT {
		return
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.mouseClick(mp, idx)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		endPoint = idx
		s.mouseUp(mp)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		s.Mode = Normal
	}
}

func (s *System) mouseClick(mp *m.Map, currentMousePos int) {
	switch s.Mode {

	case Normal:
		noneMode(mp, currentMousePos)

	case FloorTiles:
		floorTileMode(mp, currentMousePos)

	default:
		fmt.Println("mouseClick: unknown MouseMode:", s.Mode)
	}
}

func (s *System) mouseUp(mp *m.Map) {
	if startPoint >= 0 {
		mouseUpSetWalls(mp, startPoint, endPoint)
	}

	unsetHasClicked()
}

func (s *System) mouseHover(mp *m.Map) {
	switch s.Mode {
	default:
	}
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
