// Package jobsystem manages
// all the life cycles for jobs
// given to in-game characters.
package jobsystem

import (
	"math/rand"
	m "projects/games/warf2/worldmap"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// JobSystem manages all ingame jobs
// for dwarves.
type JobSystem struct {
	Jobs             []Job    `json:"jobs"`
	Workers          []Worker `json:"-"`
	AvailableWorkers []Worker `json:"-"`
	Map              *m.Map   `json:"-"`
}

// Update runs every frame, handling
// the lifetime cycle of jobs.
func (j *JobSystem) Update() {
	j.sortPriority()
	j.removeFinishedJobs()
	j.AvailableWorkers = j.availableWorkers()

	/* ---------------------------------- Check --------------------------------- */

	j.checkForDiggingJobs()

	/* ----------------------------- Assign and work ---------------------------- */

	j.assignWorkers(j.AvailableWorkers)
	j.performWork()
}

func (j *JobSystem) sortPriority() {
	sort.Sort(j)
}

func (j *JobSystem) removeFinishedJobs() {
	var jobs []Job
	for _, job := range j.Jobs {
		if !job.NeedsToBeRemoved(j.Map) {
			jobs = append(jobs, job)
		}
	}
	j.Jobs = jobs
}

func (j *JobSystem) assignWorkers(availableWorkers []Worker) {
	for _, job := range j.Jobs {
		if !WaitingForWorker(job) {
			continue
		}

		foundWorker := false
		for _, worker := range availableWorkers {
			if worker.Available() {
				foundWorker = SetWorkerAndMove(job, worker, j.Map)
			}
		}

		if foundWorker {
			availableWorkers = j.availableWorkers()
			continue
		}
	}
}

func (j *JobSystem) availableWorkers() []Worker {
	var workers []Worker
	for _, worker := range j.Workers {
		if worker.Available() {
			workers = append(workers, worker)
		}
	}
	return workers
}

func (j *JobSystem) performWork() {
	for _, worker := range j.Workers {
		if !worker.HasJob() {
			continue
		}
		worker.PerformWork(j.Map)
	}
}

/* ----------------------------- sort.Interface ----------------------------- */

func (jb *JobSystem) Len() int {
	return len(jb.Jobs)
}

func (jb *JobSystem) Less(i, j int) bool {
	fst := jb.Jobs[i].Priority()
	snd := jb.Jobs[j].Priority()
	// Randomize equally prioritized.
	if fst == snd {
		return rand.Intn(2) == 1
	}
	// Highest priority first.
	return fst > snd
}

func (jb *JobSystem) Swap(i, j int) {
	jb.Jobs[i], jb.Jobs[j] = jb.Jobs[j], jb.Jobs[i]
}
