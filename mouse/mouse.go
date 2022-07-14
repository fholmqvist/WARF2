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
	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// System for handling
// all functionality by mouse.
type System struct {
	Mode               Mode
	startPoint         int  // First click.
	endPoint           int  // Mouse up.
	hasClicked         bool // Prevent resetting startPoint.
	firstClickedSprite int  // Remember what our first "hit" was.
	justPlacedRoom     bool // To prevent double-placing rooms.
}

func NewSystem() *System {
	return &System{
		startPoint:         -1,
		endPoint:           -1,
		firstClickedSprite: -1,
	}
}

// Handle all the mouse interactivity.
func (s *System) Handle(mp *m.Map, rs *room.Service, dwarves *[]*dwarf.Dwarf) string {
	mousePos := MouseIdx()
	if mousePos < 0 || mousePos > gl.TilesT {
		mp.ClearSelectedTiles()
		s.unsetHasClicked()
		return ""
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.mouseClick(mp, rs, dwarves, mousePos)
		return ""
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		s.endPoint = mousePos
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
		s.noneMode(mp, dwarves, currentMousePos)
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
		s.setHasClicked(currentMousePos)
	default:
		panic(fmt.Sprintf("mouseClick: unknown MouseMode: %v", s.Mode))
	}
	// To prevent double placement of rooms
	// due to the refiring nature of this function.
	if !s.justPlacedRoom && rm != nil {
		rs.AddRoomByType(mp, currentMousePos, rm)
	}
	s.justPlacedRoom = true
	gl.Delay(func() { s.justPlacedRoom = false })
}

func (s *System) mouseUp(mp *m.Map, rs *room.Service) {
	if s.startPoint == -1 {
		return
	}
	switch s.Mode {
	case Normal:
		FuncOverRange(mp, s.startPoint, s.endPoint, s.mouseUpSetWalls)

	case Delete:
		rs.DeleteRoomAtMousePos(mp, s.startPoint)
		s.unsetHasClicked()
	}
	mp.ClearSelectedTiles()
	s.unsetHasClicked()
}

func (s *System) mouseHover(mp *m.Map, dwarves *[]*dwarf.Dwarf, currentMousePos int) string {
	if currentMousePos < 0 || currentMousePos > gl.TilesT {
		return ""
	}
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
