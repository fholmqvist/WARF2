// Package mouse handles all
// mouse-oriented interactions.
package mouse

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"

	m "projects/games/warf2/worldmap"
)

// System for handling
// all functionality by mouse.
type System struct {
	Mode Mode
}

// Mode enum for managing mouse action state.
type Mode int

// Mode enum.
const (
	Normal Mode = iota

	FloorTiles
	ResetFloor

	PlaceItem
	PlaceFurniture
	RemoveItem
)

// Handle all the mouse interactivity.
func (s *System) Handle(mp *m.Map) {
	s.mouseHover(mp)

	idx := mousePos()

	if idx < 0 || idx > m.TilesT {
		return
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.mouseClick(mp, idx)
		return
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		endPoint = idx
		s.mouseUp(mp)
		return
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		s.Mode = Normal
		return
	}
}

func (s *System) mouseClick(mp *m.Map, currentMousePos int) {
	switch s.Mode {

	case Normal:
		noneMode(mp, currentMousePos)

	case FloorTiles:
		floorTileMode(mp, currentMousePos)

	case ResetFloor:
		resetFloorMode(mp, currentMousePos)

	case PlaceItem:
		placeItemMode(mp, currentMousePos)

	case PlaceFurniture:
		placeFurnitureMode(mp, currentMousePos)

	case RemoveItem:
		removeItemMode(mp, currentMousePos)

	default:
		fmt.Println("mouseClick: unknown MouseMode:", s.Mode)
	}
}

func (s *System) mouseUp(mp *m.Map) {
	if startPoint >= 0 {
		mouseRange(mp, startPoint, endPoint, []func(*m.Map, int, int){mouseUpSetWalls})
	}

	unsetHasClicked()
}

// TODO: Overlays, placeholders, highlights...
func (s *System) mouseHover(mp *m.Map) {
	switch s.Mode {
	default:
	}
}
