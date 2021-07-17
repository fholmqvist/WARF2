// Package job provides the interface
// implementations of Job to be used
// in JobService.
package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

const (
	finished   = true
	unfinished = false
)

///////////////////////////////////
// TODO
// This interface is crap.
// In fact, jobs in general
// need to be revamped.
// Expand WorkStates to include
// more types of walking so that
// we can simplify mid-work walking
// logic and generally just remove
// this crap interface.
//
// DONE: Step 1: Remove priority,
// just inline it with type switch so
// that we can see it in one place
// instead of having to dig around
// every single file by hand.
//
// Step 2: JobType enum and just
// switch on everything? Yeah yeah
// #SOLID etc, but being able to see
// everything is increasingly
// important. Clarity over readability,
// I'm starting to get it.
///////////////////////////////////

// Job declares the common interface
// for jobs, in order to be used within
// the job service.
type Job interface {
	NeedsToBeRemoved(*m.Map) bool
	PerformWork(*m.Map, []*dwarf.Dwarf) bool
	Finish(*m.Map, *room.Service)
	GetWorker() *dwarf.Dwarf
	SetWorker(*dwarf.Dwarf)
	GetDestinations() []int
	String() string
}
