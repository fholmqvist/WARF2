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

// Job declares the common interface
// for jobs, in order to be used within
// the job service.
type Job interface {
	PerformWork(*m.Map, []*dwarf.Dwarf, *room.Service) bool
	Remove() bool
	GetWorker() *dwarf.Dwarf
	SetWorker(*dwarf.Dwarf)
	GetDestinations() []int
	HasInternalMove() bool
	String() string
}

type JobBase struct {
	dwarf        *dwarf.Dwarf
	destinations []int
	remove       bool
}

func NewJobBase(destinations []int) JobBase {
	return JobBase{
		destinations: destinations,
	}
}

func (j *JobBase) GetDestinations() []int {
	return j.destinations
}

func (j *JobBase) Remove() bool {
	return j.remove
}

func (j *JobBase) GetWorker() *dwarf.Dwarf {
	return j.dwarf
}

func (j *JobBase) SetWorker(d *dwarf.Dwarf) {
	j.dwarf = d
}
