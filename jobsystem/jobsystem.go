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
	Jobs    []Job    `json:"jobs"`
	Workers []Worker `json:"-"`
	Map     *m.Map   `json:"-"`
}

// Update runs every frame, handling
// the lifetime cycle of jobs.
func (j *JobSystem) Update() {
	j.sortJobs()
	j.removeFinishedJobs()
	j.assignWorkers()
	j.checkForDiggingJobs()
	j.performWork()
}

func (j *JobSystem) sortJobs() {
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

func (j *JobSystem) assignWorkers() {
	availableWorkers := j.availableWorkers()
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

func (j *JobSystem) checkForDiggingJobs() {
	for _, wall := range j.Map.Tiles {
		if !m.IsSelectedWall(wall.Sprite) || !wall.NeedsInteraction {
			continue
		}

		hasFoundJob := false
		for _, destination := range m.SurroundingTilesFour(wall.Idx) {
			if hasFoundJob {
				break
			}

			if m.IsColliding(j.Map, wall.Idx, destination.Dir) {
				continue
			}

			if j.diggingJobAlreadyExists(destination.Idx, wall.Idx) {
				continue
			}

			diggingJob := Digging{
				worker:      nil,
				destination: destination.Idx,
				wallIdx:     wall.Idx,
			}

			// We have satisfied the need
			// as a worker is on the way.
			wall.NeedsInteraction = false

			j.Jobs = append(j.Jobs, &diggingJob)
			hasFoundJob = true
		}
	}
}

func (j *JobSystem) diggingJobAlreadyExists(dIdx, wIdx int) bool {
	for _, job := range j.Jobs {
		d, ok := job.(*Digging)
		if !ok {
			continue
		}
		if d.destination == dIdx && d.wallIdx == wIdx {
			return true
		}
	}
	return false
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
