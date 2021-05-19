// Package jobservice manages
// all the life cycles for jobs
// given to in-game characters.
package jobservice

import (
	"math/rand"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/job"
	"projects/games/warf2/room"
	m "projects/games/warf2/worldmap"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// JobService manages all ingame jobs
// for dwarves.
type JobService struct {
	Jobs             []job.Job      `json:"jobs"`
	Workers          []*dwarf.Dwarf `json:"-"`
	AvailableWorkers []*dwarf.Dwarf `json:"-"`
	Map              *m.Map         `json:"-"`
}

// Update runs every frame, handling
// the lifetime cycle of jobs.
func (j *JobService) Update(rs *room.Service) {
	j.sortPriority()
	j.removeFinishedJobs(rs)
	j.updateAvailableWorkers()

	/* ---------------------------------- Check --------------------------------- */

	j.checkForDiggingJobs()
	j.checkForCarryingJobs(rs)

	/* ----------------------------- Assign and work ---------------------------- */

	j.assignWorkers()
	j.performWork()
}

func (j *JobService) sortPriority() {
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
	availableWorkers := j.AvailableWorkers
	for _, job := range j.Jobs {
		if HasWorker(job) {
			continue
		}
		var foundWorker bool
	lookingForWorker:
		for _, worker := range availableWorkers {
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
			availableWorkers = j.AvailableWorkers
		}
	}
}

func (j *JobService) updateAvailableWorkers() {
	var availableDwarves []*dwarf.Dwarf
	for _, dwarf := range j.Workers {
		if dwarf.Available() {
			availableDwarves = append(availableDwarves, dwarf)
		}
	}
	j.AvailableWorkers = availableDwarves
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
		finished := jb.PerformWork(j.Map)
		if !finished {
			continue
		}
		d.SetToAvailable()
	}
}

/* ----------------------------- sort.Interface ----------------------------- */

func (jb *JobService) Len() int {
	return len(jb.Jobs)
}

func (jb *JobService) Less(i, j int) bool {
	fst := jb.Jobs[i].Priority()
	snd := jb.Jobs[j].Priority()
	// Randomize equally prioritized.
	if fst == snd {
		return rand.Intn(2) == 1
	}
	// Highest priority first.
	return fst > snd
}

func (jb *JobService) Swap(i, j int) {
	jb.Jobs[i], jb.Jobs[j] = jb.Jobs[j], jb.Jobs[i]
}
