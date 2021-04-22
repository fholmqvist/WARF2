package dwarf

import (
	"log"
	"math/rand"
	e "projects/games/warf2/entity"
	"projects/games/warf2/jobsystem"
	m "projects/games/warf2/worldmap"
)

// Dwarf is the foundational struct
// for in game characters.
type Dwarf struct {
	e.Entity
	Walker
	Characteristics
	Needs

	state jobsystem.WorkerState
	job   *jobsystem.Job
}

func New(startingIdx int) Dwarf {
	return Dwarf{
		Entity: e.Entity{
			Sprite: rand.Intn(DwarfTeal),
			Idx:    startingIdx,
		},
		Characteristics: GenerateCharacteristics(),
	}
}

// Walk placeholder, called every frame.
func (d *Dwarf) Walk(mp *m.Map) {
	if len(d.path) == 0 {
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
			log.Fatal(err)
		}

		d.Move(mp, &d.Entity, dir)
	}
}

func (d *Dwarf) traversePath(mp *m.Map) {
	if len(d.path) == 0 {
		return
	}

	next := d.path[0]

	if d.Idx == next {
		d.path = d.path[1:]
		return
	}

	dir, err := m.NextIdxToDir(d.Idx, next)
	if err != nil {
		log.Fatal(err)
	}

	if d.Move(mp, &d.Entity, dir) {
		d.path = d.path[1:]
	}
}
