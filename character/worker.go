package character

import (
	"projects/games/warf2/jobsystem"
	"projects/games/warf2/worldmap"
)

// HasJob returns whether
// characters job is nil.
func (ch *Dwarf) HasJob() bool {
	return ch.job != nil
}

// SetJob sets job for
// given character.
func (ch *Dwarf) SetJob(job jobsystem.Job) bool {
	if job == nil {
		ch.SetToAvailable()
		return false
	}

	ch.job = &job
	ch.state = jobsystem.WorkerHasJob
	return true
}

// Available checks whether worker is available.
func (ch *Dwarf) Available() bool {
	return ch.state == jobsystem.WorkerIdle
}

// SetToAvailable sets availability of worker.
func (ch *Dwarf) SetToAvailable() {
	ch.state = jobsystem.WorkerIdle
	ch.job = nil
}

// MoveTo calculates a new path
// and sends worker to it.
func (ch *Dwarf) MoveTo(idx int, mp *worldmap.Map) bool {
	from, ok := mp.GetTileByIndex(ch.Idx)
	if !ok {
		ch.SetToAvailable()
		return false
	}

	to, ok := mp.GetTileByIndex(idx)
	if !ok {
		ch.SetToAvailable()
		return false
	}

	ok = ch.InitiateWalk(from, to)
	if !ok {
		ch.SetToAvailable()
		return false
	}

	ch.state = jobsystem.WorkerMovingTowards
	return true
}

// PerformWork checks if the character is ready,
// and performs the given work.
func (ch *Dwarf) PerformWork(mp *worldmap.Map) {
	job := *ch.job

	if ch.Idx != job.GetDestination() {
		if len(ch.path) == 0 {
			ch.SetToAvailable()
		}

		return
	}

	finished := job.PerformWork(mp)()
	if !finished {
		return
	}

	ch.SetToAvailable()
}
