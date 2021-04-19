package jobsystem

import (
	m "projects/games/warf2/worldmap"
)

// Job declares the common interface
// for jobs, in order to be used within
// the job system.
type Job interface {
	WaitingForWorker() bool
	SetWorkerAndMove(Worker, *m.Map) bool
	CheckState() JobState
	NeedsToBeRemoved(*m.Map) bool
	PerformWork(*m.Map) func() bool
	GetDestination() int
}
