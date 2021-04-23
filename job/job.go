package job

import (
	"projects/games/warf2/dwarf"
	m "projects/games/warf2/worldmap"
)

const (
	finished   = true
	unfinished = false
)

// Job declares the common interface
// for jobs, in order to be used within
// the job system.
type Job interface {
	NeedsToBeRemoved(*m.Map) bool
	PerformWork(*m.Map) bool
	Priority() int

	GetWorker() *dwarf.Dwarf
	SetWorker(*dwarf.Dwarf)
	GetDestination() int
}
