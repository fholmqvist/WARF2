package dwarf

import (
	"fmt"
	m "projects/games/warf2/worldmap"
)

// SetJob sets job for
// given character.
func (d *Dwarf) SetJob() {
	d.State = WorkerHasJob
}

// HasJob returns whether
// characters job is nil.
func (d *Dwarf) HasJob() bool {
	return d.State != WorkerIdle
}

// Available checks whether worker is available.
func (d *Dwarf) Available() bool {
	return !d.HasJob()
}

// SetToAvailable sets availability of worker.
func (d *Dwarf) SetToAvailable() {
	d.State = WorkerIdle
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
	d.State = WorkerMoving
	fmt.Println(d.State)
	return true
}
