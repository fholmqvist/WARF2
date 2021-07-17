// Package mouse handles all
// mouse-oriented interactions.
package mouse

///////////////////////////////////////
// TODO
// Perhaps this entire system
// should be baked into UI?
// There are cross concerns here
// where the UI now has mouse-over
// and click functionality of its own,
// splitting and duplicating efforts.
///////////////////////////////////////

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

// This cluster of variables
// help with (de)selecting walls.
var startPoint = -1
var endPoint = -1
var hasClicked = false
var firstClickedSprite = -1

// System for handling
// all functionality by mouse.
type System struct {
	Mode Mode
}

func NewSystem() *System {
	return &System{}
}

// Handle all the mouse interactivity.
func (s *System) Handle(mp *m.Map, rs *room.Service, dwarves *[]*dwarf.Dwarf) {
	idx := MouseIdx()
	if idx < 0 || idx > globals.TilesT {
		mp.ClearSelectedTiles()
		unsetHasClicked()
		return
	}
	s.mouseHover(mp)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.mouseClick(mp, rs, dwarves, idx)
		return
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		endPoint = idx
		s.mouseUp(mp, rs)
		return
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		s.Mode = Normal
		return
	}
}

func (s *System) mouseClick(mp *m.Map, rs *room.Service, dwarves *[]*dwarf.Dwarf, currentMousePos int) {
	/////////////////////////////////
	// TODO
	// Setting and deleting rooms
	// removes items that are
	// already there.
	/////////////////////////////////
	switch s.Mode {

	case Normal:
		noneMode(mp, dwarves, currentMousePos)

	case Storage:
		rs.AddStorage(mp, currentMousePos)

	case Farm:
		rs.AddFarm(mp, currentMousePos)

	case Library:
		rs.AddLibrary(mp, currentMousePos)

	case Delete:
		setHasClicked(currentMousePos)

	default:
		fmt.Println("mouseClick: unknown MouseMode:", s.Mode)
	}
}

func (s *System) mouseUp(mp *m.Map, rs *room.Service) {
	if startPoint == -1 {
		return
	}
	switch s.Mode {

	case Normal:
		FuncOverRange(mp, startPoint, endPoint, mouseUpSetWalls)

	case Delete:
		rs.DeleteRoomAtMousePos(mp, startPoint)
		unsetHasClicked()
	}
	mp.ClearSelectedTiles()
	unsetHasClicked()
}

/////////////////////////////////////////////////
// TODO
// Overlays, placeholders, highlights...
/////////////////////////////////////////////////
func (s *System) mouseHover(mp *m.Map) {
	switch s.Mode {
	default:
	}
}
