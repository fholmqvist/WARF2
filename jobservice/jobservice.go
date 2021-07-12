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

// JobService manages all
// ingame jobs for dwarves.
type JobService struct {
	Jobs             []job.Job      `json:"jobs"`
	Workers          []*dwarf.Dwarf `json:"workers"`
	AvailableWorkers []*dwarf.Dwarf `json:"-"`
	Map              *m.Map         `json:"-"`
}

// Update runs every frame, handling
// the lifetime cycle of jobs.
func (j *JobService) Update(rs *room.Service, mp *m.Map) {
	j.sortJobPriorities()
	j.removeFinishedJobs(rs)
	j.updateAvailableWorkers()

	/* ---------------------------------- Check --------------------------------- */
	j.checkForDiggingJobs()
	j.checkForCarryingJobs(rs)
	j.checkForReading(mp)

	/* ----------------------------- Assign and work ---------------------------- */
	j.assignWorkers()
	j.performWork()
}

func (j *JobService) sortJobPriorities() {
	sort.Sort(j)
}

func (j *JobService) removeFinishedJobs(rs *room.Service) {
	var jobs []job.Job
	for _, job := range j.Jobs {
		if job.NeedsToBeRemoved(j.Map) {
			job.Finish(j.Map, rs)
			continue
		}
		jobs = append(jobs, job)
	}
	j.Jobs = jobs
}

func (j *JobService) assignWorkers() {
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

func (j *JobService) updateAvailableWorkers() {
	var available []*dwarf.Dwarf
	for _, dwarf := range j.Workers {
		if dwarf.Available() {
			available = append(available, dwarf)
		}
	}
	j.AvailableWorkers = available
}

func (j *JobService) performWork() {
	for _, jb := range j.Jobs {
		d := jb.GetWorker()
		if d == nil {
			continue
		}
		if !d.HasJob() {
			continue
		}
		var hasArrived bool
		for _, destination := range jb.GetDestinations() {
			if d.Idx == destination {
				hasArrived = true
				break
			}
		}
		if !hasArrived {
			if len(d.Path) == 0 {
				d.SetToAvailable()
			}
			continue
		}
		finished := false
		switch jb.(type) {
		case *job.LibraryRead:
			finished = jb.PerformWork(j.Map, j.Workers)
		default:
			finished = jb.PerformWork(j.Map, nil)
		}
		if !finished {
			continue
		}
		d.SetToAvailable()
	}
}
