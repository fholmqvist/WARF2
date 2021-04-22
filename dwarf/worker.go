package dwarf

import (
	"projects/games/warf2/jobsystem"
	m "projects/games/warf2/worldmap"
)

// HasJob returns whether
// characters job is nil.
func (d *Dwarf) HasJob() bool {
	return d.job != nil
}

// SetJob sets job for
// given character.
func (d *Dwarf) SetJob(job jobsystem.Job) bool {
	if job == nil {
		d.SetToAvailable()
		return false
	}

	d.job = &job
	d.state = jobsystem.WorkerHasJob
	return true
}

// Available checks whether worker is available.
func (d *Dwarf) Available() bool {
	return d.state == jobsystem.WorkerIdle
}

// SetToAvailable sets availability of worker.
func (d *Dwarf) SetToAvailable() {
	d.state = jobsystem.WorkerIdle
	d.job = nil
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

	d.state = jobsystem.WorkerMovingTowards
	return true
}

// PerformWork checks if the character is ready,
// and performs the given work.
func (d *Dwarf) PerformWork(mp *m.Map) {
	job := *d.job

	if d.Idx != job.GetDestination() {
		if len(d.path) == 0 {
			d.SetToAvailable()
		}
		return
	}

	finished := job.PerformWork(mp)
	if !finished {
		return
	}

	d.SetToAvailable()
}
