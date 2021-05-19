package job

import (
	"fmt"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/resource"
	"projects/games/warf2/room"
	m "projects/games/warf2/worldmap"
)

type Carrying struct {
	resource        resource.Resource
	dwarf           *dwarf.Dwarf
	destinations    []int
	goalDestination int
	storageIdx      int
	sprite          int
	path            []int
	prev            int
}

func NewCarrying(destinations []int, r resource.Resource, storageIdx int, goalDestination, sprite int) *Carrying {
	return &Carrying{
		resource:        r,
		dwarf:           nil,
		destinations:    destinations,
		goalDestination: goalDestination,
		storageIdx:      storageIdx,
		sprite:          sprite,
		path:            nil,
		prev:            0,
	}
}

func (c *Carrying) NeedsToBeRemoved(mp *m.Map) bool {
	return c.path != nil && len(c.path) == 0
}

func (c *Carrying) Finish(mp *m.Map, s *room.Service) {
	if c.dwarf == nil {
		return
	}
	dropIdx, ok := s.Storages[c.storageIdx].AddItem(c.dwarf.Idx, 1, c.resource)
	c.dwarf.SetToAvailable()
	c.dwarf = nil
	if !ok {
		fmt.Println("Carrying: Finish: Couldn't find storage tile.",
			"Ignoring item (forever lost!).")
		return
	}
	mp.Items[dropIdx].Sprite = c.sprite
}

func (c *Carrying) PerformWork(mp *m.Map) bool {
	if setupPath(c, mp) {
		return unfinished
	}
	if len(c.path) == 0 {
		return finished
	}
	moveDwarf(c, mp)
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

func (c *Carrying) String() string {
	return "Carrying"
}

func setupPath(c *Carrying, mp *m.Map) bool {
	if c.path != nil {
		return false
	}
	mp.Items[c.dwarf.Idx].Sprite = 0
	mp.Items[c.dwarf.Idx].Resource = 0
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

func moveDwarf(c *Carrying, mp *m.Map) {
	// Move indexes to current path index.
	c.dwarf.Idx = c.path[0]
	c.destinations[0] = c.path[0]
	c.prev = c.path[0]
	// Iterate path.
	c.path = c.path[1:]
}
