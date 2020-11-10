// Package room defines all
// the different in-game rooms.
package room

import (
	"projects/games/warf2/character"
	"projects/games/warf2/worldmap"
)

// System for gathering data
// and functionality related to rooms.
type System struct {
	Libraries []Room
}

// Room wraps all the functionality
// and data related to a room.
type Room struct {
	Floors []worldmap.Tile
	Items  []worldmap.Tile
}

// Use room with given dwarf.
func (r *Room) Use(d character.Dwarf) {

}
