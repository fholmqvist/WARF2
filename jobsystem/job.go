package jobsystem

import (
	m "projects/games/warf2/worldmap"
)

// Job declares the common interface
// for jobs, in order to be used within
// the job system.
type Job interface {
	NeedsToBeRemoved(*m.Map) bool
	PerformWork(*m.Map) func() bool
	Priority() int

	GetWorker() *Worker
	SetWorker(*Worker)
	GetState() JobState
	SetState(JobState)
	GetDestination() int
}
