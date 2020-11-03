package characters

import (
	"log"
	"math/rand"
	e "projects/games/warf2/entity"
	"projects/games/warf2/jobsystem"
	m "projects/games/warf2/worldmap"
)

// Character is the foundational struct
// for in game characters.
type Character struct {
	e.Entity
	Walker

	state jobsystem.WorkerState
	job   *jobsystem.Job
}

// Walk placeholder, called every frame.
func (ch *Character) Walk(mp *m.Map) {
	if len(ch.path) == 0 {
		ch.randomWalk(mp)
		return
	}

	ch.traversePath(mp)
}

func (ch *Character) randomWalk(mp *m.Map) {
	// Pause most of the time
	if rand.Intn(100) > 90 {
		dir, err := m.GetDirection(rand.Intn(4))
		if err != nil {
			log.Fatal(err)
		}

		ch.Move(mp, &ch.Entity, dir)
	}
}

func (ch *Character) traversePath(mp *m.Map) {
	if len(ch.path) == 0 {
		return
	}

	next := ch.path[0]

	if ch.Idx == next {
		ch.path = ch.path[1:]
		return
	}

	dir, err := m.NextIdxToDir(ch.Idx, next)
	if err != nil {
		log.Fatal(err)
	}

	if ch.Move(mp, &ch.Entity, dir) {
		ch.path = ch.path[1:]
	}
}
