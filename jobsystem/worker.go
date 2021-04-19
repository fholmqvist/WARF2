package jobsystem

import m "projects/games/warf2/worldmap"

// Worker defines the interface
// for all characters who are
// eligible workers.
type Worker interface {
	HasJob() bool
	SetJob(Job) bool
	Available() bool
	SetToAvailable()
	MoveTo(int, *m.Map) bool
	PerformWork(*m.Map)
}
