package jobsystem

import (
	"fmt"
	m "projects/games/warf2/worldmap"
)

// Job declares the common interface
// for jobs, in order to be used within
// the job system.
type Job interface {
	WaitingForWorker() bool
	SetWorker(Worker) bool
	CheckState() JobState
}

// JobSystem manages all ingame jobs
// for dwarves.
type JobSystem struct {
	Jobs    []Job
	Workers []Worker
}

// Update runs every frame, handling
// the lifetime cycle of jobs.
func (j *JobSystem) Update() {
	j.removeFinishedJobs()
	j.assignWorkers()

	fmt.Println("JOBS:", len(j.Jobs))
}

// CheckForDiggingJobs is called by
// the consumer of the JobSystem in
// the equivalent update loop.
func (j *JobSystem) CheckForDiggingJobs(mp *m.Map) {
	for _, wall := range mp.Tiles {
		if !m.IsSelectedWall(wall.Sprite) {
			continue
		}

		hasFoundJob := false
		neighbors := m.SurroundingTilesEight(wall.Idx)
		for _, destination := range neighbors {
			if hasFoundJob {
				break
			}

			if m.IsColliding(mp, destination.Idx, destination.Dir) {
				continue
			}

			if j.diggingJobAlreadyExists(destination.Idx, wall.Idx) {
				continue
			}

			diggingJob := Digging{
				worker:      nil,
				state:       New,
				destination: destination.Idx,
				wallIdx:     wall.Idx,
			}

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

func (j *JobSystem) removeFinishedJobs() {
	var jobs []Job
	for _, job := range j.Jobs {
		if job.CheckState() != Done {
			jobs = append(jobs, job)
		}
	}
	j.Jobs = jobs
}

func (j *JobSystem) assignWorkers() {
	availableWorkers := j.availableWorkers()
	for _, job := range j.Jobs {
		if !job.WaitingForWorker() {
			continue
		}

		foundWorker := false
		for _, worker := range availableWorkers {
			if worker.Available() {
				foundWorker = job.SetWorker(worker)
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
