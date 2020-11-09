// Package room defines all
// the different in-game rooms.
package room

import "projects/games/warf2/character"

// Room interface for
// all room related functionality.
type Room interface {
	Use(dwarf *character.Dwarf)
}
