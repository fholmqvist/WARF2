// Package room defines all
// the different in-game rooms.
package room

import (
	"projects/games/warf2/dwarf"
	m "projects/games/warf2/worldmap"
)

var globalID uint16 = 0

// System for gathering data
// and functionality related to rooms.
type System struct {
	Libraries []Library
}

func (s *System) AddLibrary(mp *m.Map, currentMousePos int) {
	x, y := m.IdxToXY(currentMousePos)
	l := NewLibrary(mp, x, y)
	if l != nil {
		s.Libraries = append(s.Libraries, *l)
	}
}

// Room wraps all the functionality
// and data related to a room.
type Room struct {
	Floors []m.Tile
	Items  []m.Tile

	id uint16
}

// NewRoom returns a new room.
func NewRoom() Room {
	globalID++
	return Room{
		id: globalID,
	}
}

// Use room with given dwarf.
func (r *Room) Use(d dwarf.Dwarf) {

}
