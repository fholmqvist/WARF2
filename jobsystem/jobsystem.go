// Package jobsystem manages
// all the life cycles for jobs
// given to in-game characters.
package jobsystem

import (
	"math/rand"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/job"
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
func (j *JobService) Update() {
	j.sortPriority()
	j.removeFinishedJobs()
	j.AvailableWorkers = j.availableWorkers()

	/* ---------------------------------- Check --------------------------------- */

	j.checkForDiggingJobs()

	/* ----------------------------- Assign and work ---------------------------- */

	j.assignWorkers(j.AvailableWorkers)
	j.performWork()
}

func (j *JobService) sortPriority() {
	sort.Sort(j)
}

func (j *JobService) removeFinishedJobs() {
	var jobs []job.Job
	for _, job := range j.Jobs {
		if !job.NeedsToBeRemoved(j.Map) {
			jobs = append(jobs, job)
		}
	}
	j.Jobs = jobs
}

func (j *JobService) assignWorkers(availableWorkers []*dwarf.Dwarf) {
	for _, job := range j.Jobs {
		if !WaitingForWorker(job) {
			continue
		}
		foundWorker := false
	found:
		for _, worker := range availableWorkers {
			if worker.Available() {
				foundWorker = SetWorkerAndMove(job, worker, j.Map)
				if !foundWorker {
					continue
				}
				break found
			}
		}
		if foundWorker {
			availableWorkers = j.availableWorkers()
			continue
		}
	}
	// for _, worker := range availableWorkers {
	// 	if !worker.Available() {
	// 		continue
	// 	}
	// 	for _, job := range j.Jobs {
	// 		if HasWorker(job) {
	// 			continue
	// 		}
	// 		if !SetWorkerAndMove(job, worker, j.Map) {
	// 			job.Refresh()
	// 		}
	// 	}
	// }
}

func (j *JobService) availableWorkers() []*dwarf.Dwarf {
	var dwarves []*dwarf.Dwarf
	for _, dwarf := range j.Workers {
		if dwarf.Available() {
			dwarves = append(dwarves, dwarf)
		}
	}
	return dwarves
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
		if d.Idx != jb.GetDestination() {
			if len(d.Path) == 0 {
				d.SetToAvailable()
			}
			return
		}
		finished := jb.PerformWork(j.Map)
		if !finished {
			return
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
