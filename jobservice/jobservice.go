// Package jobservice manages
// all the life cycles for jobs
// given to in-game characters.
package jobservice

import (
	"math/rand"
	"sort"
	"time"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/job"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Service manages all
// ingame jobs for dwarves.
type Service struct {
	Jobs             []job.Job      `json:"jobs"`
	Workers          []*dwarf.Dwarf `json:"workers"`
	AvailableWorkers []*dwarf.Dwarf `json:"-"`
	Map              *m.Map         `json:"-"`
}

func NewService(mp *m.Map) *Service {
	return &Service{Map: mp}
}

// Update runs every frame, handling
// the lifetime cycle of jobs.
func (j *Service) Update(rs *room.Service, mp *m.Map) {
	// CLEANUP.
	j.removeFinishedJobs(rs)
	j.updateAvailableWorkers()
	// JOB CHECKS.
	j.checkForDiggingJobs()
	j.checkForCarryingJobs(rs)
	j.checkForFarmingJobs(rs)
	// NEED CHECKS.
	j.checkForSleep(mp, rs)
	j.checkForReading(mp)
	// SORT.
	j.sortJobPriorities()
	// PERFORM.
	j.assignWorkers()
	j.performWork(rs)
}

func (j *Service) sortJobPriorities() {
	sort.Sort(j)
}

func (j *Service) removeFinishedJobs(rs *room.Service) {
	var jobs []job.Job
	for _, job := range j.Jobs {
		if job.NeedsToBeRemoved(j.Map, rs) {
			job.Finish(j.Map, rs)
			continue
		}
		jobs = append(jobs, job)
	}
	j.Jobs = jobs
}

func (j *Service) assignWorkers() {
	available := j.AvailableWorkers
	for _, job := range j.Jobs {
		if HasWorker(job) {
			continue
		}
		var foundWorker bool
	lookingForWorker:
		for _, worker := range available {
			if worker.HasJob() {
				continue
			}
			foundWorker = SetWorkerAndMove(job, worker, j.Map)
			if !foundWorker {
				continue
			}
			break lookingForWorker
		}
		if foundWorker {
			available = j.AvailableWorkers
		}
	}
}

func (j *Service) updateAvailableWorkers() {
	var available []*dwarf.Dwarf
	for _, dwarf := range j.Workers {
		if dwarf.Available() {
			available = append(available, dwarf)
		}
	}
	j.AvailableWorkers = available
}

func (j *Service) performWork(rs *room.Service) {
	for _, jb := range j.Jobs {
		dw := jb.GetWorker()
		if dw == nil {
			continue
		}
		if !dw.HasJob() {
			continue
		}
		for _, destination := range jb.GetDestinations() {
			if dw.Idx == destination {
				dw.State = dwarf.Arrived
				break
			}
		}
		if dw.State != dwarf.Arrived {
			if len(dw.Path) == 0 && jb.NeedsToBeRemoved(j.Map, rs) {
				dw.SetToAvailable()
				continue
			}
			if !jb.HasInternalMove() {
				continue
			}
		}
		finished := jb.PerformWork(j.Map, j.Workers, rs)
		if !finished {
			continue
		}
		dw.SetToAvailable()
	}
}
