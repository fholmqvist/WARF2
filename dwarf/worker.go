package dwarf

import (
	m "projects/games/warf2/worldmap"
)

// HasJob returns whether
// characters job is nil.
func (d *Dwarf) HasJob() bool {
	return d.state != WorkerIdle
}

// SetJob sets job for
// given character.
func (d *Dwarf) SetJob() {
	d.state = WorkerHasJob
}

// Available checks whether worker is available.
func (d *Dwarf) Available() bool {
	return d.state == WorkerIdle
}

// SetToAvailable sets availability of worker.
func (d *Dwarf) SetToAvailable() {
	d.state = WorkerIdle
}

// MoveTo calculates a new path
// and sends worker to it.
func (d *Dwarf) MoveTo(idx int, mp *m.Map) bool {
	from, ok := mp.GetTileByIndex(d.Idx)
	if !ok {
		d.SetToAvailable()
		return false
	}

	to, ok := mp.GetTileByIndex(idx)
	if !ok {
		d.SetToAvailable()
		return false
	}

	ok = d.InitiateWalk(from, to)
	if !ok {
		d.SetToAvailable()
		return false
	}

	d.state = WorkerMovingTowards
	return true
}

// PerformWork checks if the character is ready,
// and performs the given work.
// func (d *Dwarf) PerformWork(mp *m.Map) {
// 	jb := *d.job
// 	if d.Idx != jb.GetDestination() {
// 		if len(d.path) == 0 {
// 			d.SetToAvailable()
// 		}
// 		return
// 	}
// 	d.SetState(WorkerArrived)
// 	finished := jb.PerformWork(mp)
// 	if !finished {
// 		return
// 	}
// 	d.SetToAvailable()
// }

func (d *Dwarf) GetPosition() int {
	return d.Idx
}

func (d *Dwarf) GetState() WorkerState {
	return d.state
}

func (d *Dwarf) SetState(st WorkerState) {
	d.state = st
}
