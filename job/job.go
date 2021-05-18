// Package job provides the interface
// implementations of Job to be used
// in JobService.
package job

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/room"
	m "projects/games/warf2/worldmap"
)

const (
	finished   = true
	unfinished = false
)

// Job declares the common interface
// for jobs, in order to be used within
// the job service.
type Job interface {
	NeedsToBeRemoved(*m.Map) bool
	PerformWork(*m.Map) bool
	Finish(*m.Map, *room.Service)
	Priority() int

	GetWorker() *dwarf.Dwarf
	SetWorker(*dwarf.Dwarf)
	GetDestinations() []int
	String() string
}
