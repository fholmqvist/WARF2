package characters

import (
	"projects/games/warf2/jobsystem"
	"projects/games/warf2/worldmap"
)

// HasJob returns whether
// characters job is nil.
func (ch *Character) HasJob() bool {
	return ch.job != nil
}

// SetJob sets job for
// given character.
func (ch *Character) SetJob(job jobsystem.Job) bool {
	if job == nil {
		return false
	}

	ch.job = &job
	return true
}

// Available checks whether worker is available.
func (ch *Character) Available() bool {
	return ch.state == jobsystem.WorkerIdle
}

// SetAvailable sets availability of worker.
func (ch *Character) SetAvailable(available bool) {
	if available {
		ch.state = jobsystem.WorkerIdle
		ch.job = nil
		return
	}

	ch.state = jobsystem.WorkerHasJob
}

// MoveTo calculates a new path
// and sends worker to it.
func (ch *Character) MoveTo(idx int, mp *worldmap.Map) bool {
	from, ok := mp.GetTileByIndex(ch.Entity.Idx)
	if !ok {
		return false
	}

	to, ok := mp.GetTileByIndex(idx)
	if !ok {
		return false
	}

	ok = ch.Walker.InitiateWalk(from, to)
	if !ok {
		return false
	}

	ch.state = jobsystem.WorkerMovingTowards
	return true
}

// PerformWork checks if the character is ready,
// and performs the given work.
func (ch *Character) PerformWork(mp *worldmap.Map) {
	job := *ch.job

	if ch.Entity.Idx != job.GetDestination() {
		return
	}

	work := job.PerformWork(mp)
	work()
	ch.SetAvailable(true)
}
