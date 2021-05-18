package job

import (
	"projects/games/warf2/dwarf"
	m "projects/games/warf2/worldmap"
)

type Carrying struct {
	dwarf           *dwarf.Dwarf
	destinations    []int
	goalDestination int
	sprite          int
	path            []int
	prev            int
}

func NewCarrying(destinations []int, goalDestination, sprite int) *Carrying {
	return &Carrying{
		dwarf:           nil,
		destinations:    destinations,
		goalDestination: goalDestination,
		sprite:          sprite,
		path:            nil,
		prev:            0,
	}
}

func (c *Carrying) NeedsToBeRemoved(mp *m.Map) bool {
	if c.dwarf == nil && c.path != nil {
		return true
	}
	if c.path != nil && len(c.path) == 0 {
		return true
	}
	return false
}

func (c *Carrying) Reset(mp *m.Map) {
	if c.dwarf == nil {
		return
	}
	c.dwarf.SetToAvailable()
	c.dwarf = nil
}

func (c *Carrying) PerformWork(mp *m.Map) bool {
	if setupPath(c, mp) {
		return unfinished
	}
	if len(c.path) == 0 {
		return finished
	}
	moveDwarfAndSprite(c, mp)
	return unfinished
}

func (c *Carrying) Priority() int {
	return 1
}

func (c *Carrying) GetWorker() *dwarf.Dwarf {
	return c.dwarf
}

func (c *Carrying) SetWorker(dw *dwarf.Dwarf) {
	c.dwarf = dw
}

func (c *Carrying) GetDestinations() []int {
	return c.destinations
}

func setupPath(c *Carrying, mp *m.Map) bool {
	if c.path != nil {
		return false
	}
	c.prev = c.dwarf.Idx
	c.destinations[0] = c.dwarf.Idx
	path, ok := c.dwarf.CreatePath(
		&mp.Tiles[c.dwarf.Idx],
		&mp.Tiles[c.goalDestination],
	)
	if !ok {
		return false
	}
	c.path = path
	return true
}

func moveDwarfAndSprite(c *Carrying, mp *m.Map) {
	///////////////////////////////////////
	// Remove previous item sprite.
	// TODO
	// This will overwrite any sprite if
	// the tile already has one! Very dumb.
	///////////////////////////////////////
	mp.Items[c.prev].Sprite = 0
	// Move indexes to current path index.
	c.dwarf.Idx = c.path[0]
	c.destinations[0] = c.path[0]
	c.prev = c.path[0]
	// Set current item sprite.
	mp.Items[c.path[0]].Sprite = c.sprite
	// Iterate path.
	c.path = c.path[1:]
}
