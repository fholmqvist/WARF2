package job

import (
	"fmt"
	"math/rand"
	"projects/games/warf2/dwarf"
	m "projects/games/warf2/worldmap"
)

type Carrying struct {
	id              int
	dwarf           *dwarf.Dwarf
	destinations    []int
	goalDestination int
	sprite          int
}

func NewCarrying(destinations []int, goalDestination, sprite int) *Carrying {
	return &Carrying{rand.Intn(20), nil, destinations, goalDestination, sprite}
}

func (c *Carrying) NeedsToBeRemoved(mp *m.Map) bool {
	return c.dwarf == nil || c.dwarf.Idx == c.goalDestination
}

func (c *Carrying) Reset(mp *m.Map) {
	if c.dwarf == nil {
		return
	}
	c.dwarf.SetToAvailable()
	c.dwarf = nil
	mp.Items[c.goalDestination].Sprite = c.sprite
	fmt.Println("DONE", c.id)
}

func (c *Carrying) PerformWork(mp *m.Map) bool {
	mp.Items[c.destinations[0]].Sprite = 0
	c.dwarf.MoveTo(c.goalDestination, mp)
	return finished
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
