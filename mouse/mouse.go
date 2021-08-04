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

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// This cluster of variables
// help with (de)selecting walls.
var startPoint = -1
var endPoint = -1
var hasClicked = false
var firstClickedSprite = -1
var justPlacedRoom = false

// System for handling
// all functionality by mouse.
type System struct {
	Mode Mode
}

func NewSystem() *System {
	return &System{}
}

// Handle all the mouse interactivity.
func (s *System) Handle(mp *m.Map, rs *room.Service, dwarves *[]*dwarf.Dwarf) string {
	mousePos := MouseIdx()
	if mousePos < 0 || mousePos > globals.TilesT {
		mp.ClearSelectedTiles()
		unsetHasClicked()
		return ""
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.mouseClick(mp, rs, dwarves, mousePos)
		return ""
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		endPoint = mousePos
		s.mouseUp(mp, rs)
		return ""
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		s.Mode = Normal
		return ""
	}
	return s.mouseHover(mp, dwarves, mousePos)
}

func (s *System) mouseClick(mp *m.Map, rs *room.Service, dwarves *[]*dwarf.Dwarf, currentMousePos int) {
	var rm room.Room
	switch s.Mode {
	case Normal:
		noneMode(mp, dwarves, currentMousePos)
	case Storage:
		rm = &room.Storage{}
	case SleepHall:
		rm = &room.SleepHall{}
	case Farm:
		rm = &room.Farm{}
	case Brewery:
		rm = &room.Brewery{}
	case Bar:
		rm = &room.Bar{}
	case Library:
		rm = &room.Library{}
	case Delete:
		setHasClicked(currentMousePos)
	default:
		panic(fmt.Sprintf("mouseClick: unknown MouseMode: %v", s.Mode))
	}
	if rm == nil {
		return
	}
	if !justPlacedRoom {
		rs.AddRoom(mp, currentMousePos, rm)
	}
	justPlacedRoom = true
	globals.Delay(func() { justPlacedRoom = false })
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

func (s *System) mouseHover(mp *m.Map, dwarves *[]*dwarf.Dwarf, currentMousePos int) string {
	for _, dwarf := range *dwarves {
		if dwarf.Idx == currentMousePos {
			return dwarf.String()
		}
	}
	if mp.Items[currentMousePos].Sprite != entity.NoItem {
		itm := mp.Items[currentMousePos]
		if itm.ResourceAmount == 0 {
			return fmt.Sprintf("ITEM: %v.", entity.ItemToString(itm.Sprite))
		}
		return fmt.Sprintf("ITEM: %v.  AMOUNT: %v.", entity.ItemToString(itm.Sprite), itm.ResourceAmount)
	}
	return ""
}
