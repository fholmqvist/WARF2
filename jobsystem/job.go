package jobsystem

import (
	"projects/games/warf2/worldmap"
)

// Job declares the common interface
// for jobs, in order to be used within
// the job system.
type Job interface {
	WaitingForWorker() bool
	SetWorkerAndMove(Worker, *worldmap.Map) bool
	CheckState() JobState
	NeedsToBeRemoved(*worldmap.Map) bool
	PerformWork(*worldmap.Map) func() bool
	GetDestination() int
}
