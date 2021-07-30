package dwarf

import (
	"fmt"
	"math/rand"

	e "github.com/Holmqvist1990/WARF2/entity"
	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

// Dwarf is the foundational struct
// for in game characters.
type Dwarf struct {
	e.Entity
	Walker
	Characteristics
	Needs
	State WorkState
}

func New(startingIdx int, name string) *Dwarf {
	return &Dwarf{
		Entity: e.Entity{
			Sprite: rand.Intn(DwarfTeal),
			Idx:    startingIdx,
		},
		Characteristics: GenerateCharacteristics(name),
	}
}

func (d Dwarf) String() string {
	return fmt.Sprintf("NAME: %v.\n\tIDX: %v.\n\tSTATE: %v.\n\tPATH-LEN: %v.",
		d.Name, d.Idx, d.State, len(d.Path))
}

// Walk placeholder, called every frame.
func (d *Dwarf) Walk(mp *m.Map) {
	if len(d.Path) == 0 {
		if d.HasJob() {
			return
		}
		d.randomWalk(mp)
		return
	}
	d.traversePath(mp)
}

func (d *Dwarf) randomWalk(mp *m.Map) {
	// Pause most of the time
	if rand.Intn(100) > 90 {
		dir, err := m.GetDirection(rand.Intn(4))
		if err != nil {
			globals.PAUSE_GAME = true
			fmt.Println(err)
		}
		d.Move(mp, &d.Entity, dir)
	}
}

func (d *Dwarf) traversePath(mp *m.Map) {
	if len(d.Path) == 0 {
		return
	}
	if d.Idx == d.Path[0] {
		d.Path = d.Path[1:]
		return
	}
tryagain:
	var dir m.Direction
	var err error
	next := d.Path[0]
	dir, err = m.NextIdxToDir(d.Idx, next)
	if err != nil {
		fmt.Println(err)
		d.Path = d.Path[1:]
		goto tryagain
	}
	if d.Move(mp, &d.Entity, dir) {
		d.Path = d.Path[1:]
	}
}
