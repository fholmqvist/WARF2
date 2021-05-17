// Package job provides the interface
// implementations of Job to be used
// in JobService.
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
	Reset(*m.Map)
	Priority() int

	GetWorker() *dwarf.Dwarf
	SetWorker(*dwarf.Dwarf)
	GetDestinations() []int
}
