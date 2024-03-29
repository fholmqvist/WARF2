package job

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Carrying struct {
	JobBase
	resource        entity.Resource
	amount          uint
	goalDestination int
	storageIdx      int
	sprite          int
	path            []int
	prev            int
}

func NewCarrying(destinations []int, r entity.Resource, storageIdx int, goalDestination, sprite int) *Carrying {
	return &Carrying{
		JobBase:         NewJobBase(destinations),
		resource:        r,
		goalDestination: goalDestination,
		storageIdx:      storageIdx,
		sprite:          sprite,
		path:            nil,
		prev:            0,
	}
}

func (c *Carrying) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf, rs *room.Service) bool {
	// Try again with
	// new storage.
	if storageMissingOrFull(c, rs) {
		return finished
	}
	// Just arrived.
	if c.path == nil {
		///////////////////////////////////
		// TODO
		// Item is no longer there, abort.
		// What should we actually do here?
		///////////////////////////////////
		if !entity.IsCarriable(mp.Items[c.dwarf.Idx].Sprite) {
			c.path = []int{}
			c.remove = true
			return finished
		}
		c.setupPath(mp)
		return unfinished
	}
	// Finished.
	if len(c.path) == 0 {
		c.finish(mp, rs)
		return finished
	}
	// Move.
	moveAlongPath(c, mp)
	return unfinished
}

func (c *Carrying) finish(mp *m.Map, rs *room.Service) {
	c.remove = true
	if c.sprite == m.None {
		return
	}
	if c.storageIdx > len(rs.Rooms)-1 {
		return
	}
	storage, ok := rs.Rooms[c.storageIdx].(*room.Storage)
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

func (c *Carrying) HasInternalMove() bool {
	return false
}

func (c *Carrying) String() string {
	return "Carrying"
}

func (c *Carrying) setupPath(mp *m.Map) {
	switch mp.Items[c.dwarf.Idx].Sprite {
	case entity.FilledBarrel:
		mp.Items[c.dwarf.Idx].Sprite = entity.EmptyBarrel
	default:
		mp.Items[c.dwarf.Idx].Sprite = entity.NoItem
	}
	c.amount = mp.Items[c.dwarf.Idx].ResourceAmount
	mp.Items[c.dwarf.Idx].Resource = 0
	mp.Items[c.dwarf.Idx].ResourceAmount = 0
	c.prev = c.dwarf.Idx
	c.destinations[0] = c.dwarf.Idx
	path, ok := m.CreatePath(
		&mp.Tiles[c.dwarf.Idx],
		&mp.Tiles[c.goalDestination],
	)
	if !ok {
		fmt.Printf("Carrying: No path for %v from %v to %v.\n", c.dwarf.Name, c.dwarf.Idx, c.goalDestination)
		return
	}
	c.path = path
}

func moveAlongPath(c *Carrying, mp *m.Map) {
	// Move indexes to current path index.
	c.dwarf.Idx = c.path[0]
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
