package job

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Carrying struct {
	resource        entity.Resource
	amount          uint
	dwarf           *dwarf.Dwarf
	destinations    []int
	goalDestination int
	storageIdx      int
	sprite          int
	path            []int
	prev            int
}

func NewCarrying(destinations []int, r entity.Resource, storageIdx int, goalDestination, sprite int) *Carrying {
	if r == entity.ResourceBeer {
		fmt.Println("NEW CARRYING BEER")
	}
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

func (c *Carrying) NeedsToBeRemoved(mp *m.Map, r *room.Service) bool {
	return c.path != nil && len(c.path) == 0
}

func (c *Carrying) Finish(mp *m.Map, s *room.Service) {
	if c.resource == entity.ResourceBeer {
		fmt.Println("BEER FINISH")
	}
	if c.dwarf == nil {
		return
	}
	defer func() {
		c.dwarf.SetToAvailable()
		c.dwarf = nil
	}()
	if c.sprite == m.None {
		return
	}
	if c.storageIdx > len(s.Rooms)-1 {
		return
	}
	storage, ok := s.Rooms[c.storageIdx].(*room.Storage)
	if !ok {
		return
	}
	dropIdx, ok := storage.AddItem(c.dwarf.Idx, c.amount, c.resource)
	if !ok {
		////////
		// TODO
		// Yeah.
		////////
		fmt.Println("Carrying: Finish: Couldn't find storage tile.",
			"Ignoring item (forever lost!).")
		return
	}
	mp.Items[dropIdx].Sprite = c.sprite
}

func (c *Carrying) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	if c.resource == entity.ResourceBeer {
		fmt.Println(c.resource.String())
	}
	if storageMissingOrFull(c, rs) {
		// Try again with
		// new storage.
		return finished
	}
	if c.path == nil {
		///////////////////////////////////
		// TODO
		// Item is no longer there, abort.
		// What should we actually do here?
		///////////////////////////////////
		if !entity.IsCarriable(mp.Items[c.dwarf.Idx].Sprite) {
			c.path = []int{}
			return finished
		}
		c.setupPath(mp)
		return unfinished
	}
	if len(c.path) == 0 {
		return finished
	}
	moveAlongPath(c, mp)
	return unfinished
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

func (c *Carrying) HasInternalMove() bool {
	return false
}

func (c *Carrying) String() string {
	return "Carrying"
}

func (c *Carrying) setupPath(mp *m.Map) {
	c.amount = mp.Items[c.dwarf.Idx].ResourceAmount
	mp.Items[c.dwarf.Idx].Sprite = 0
	mp.Items[c.dwarf.Idx].Resource = 0
	mp.Items[c.dwarf.Idx].ResourceAmount = 0
	c.prev = c.dwarf.Idx
	c.destinations[0] = c.dwarf.Idx
	path, ok := c.dwarf.CreatePath(
		&mp.Tiles[c.dwarf.Idx],
		&mp.Tiles[c.goalDestination],
	)
	if !ok {
		return
	}
	c.path = path
}

func moveAlongPath(c *Carrying, mp *m.Map) {
	// Move indexes to current path index.
	c.dwarf.Idx = c.path[0]
	c.destinations[0] = c.path[0]
	c.prev = c.path[0]
	// Iterate path.
	c.path = c.path[1:]
}

func storageMissingOrFull(c *Carrying, rs *room.Service) bool {
	if len(rs.Rooms)-1 < c.storageIdx {
		c.path = []int{}
		return true
	}
	storage, ok := rs.Rooms[c.storageIdx].(*room.Storage)
	if !ok {
		return true
	}
	if !storage.HasSpace(c.resource) {
		c.path = []int{}
		return true
	}
	return false
}
