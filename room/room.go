// Package room defines all
// the different in-game rooms.
package room

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/worldmap"
)

var globalID uint16 = 0

// System for gathering data
// and functionality related to rooms.
type System struct {
	Libraries []Library
}

// Room wraps all the functionality
// and data related to a room.
type Room struct {
	Floors []worldmap.Tile
	Items  []worldmap.Tile

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
