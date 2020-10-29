package jobsystem

import (
	m "projects/games/warf2/worldmap"
)

// JobSystem manages all ingame jobs
// for dwarves.
type JobSystem struct {
	Jobs    []Job
	Workers []Worker
	Map     *m.Map
}

// Update runs every frame, handling
// the lifetime cycle of jobs.
func (j *JobSystem) Update() {
	j.removeFinishedJobs()
	j.assignWorkers()
	j.checkForDiggingJobs()
	j.performWork()
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
		if !job.WaitingForWorker() {
			continue
		}

		foundWorker := false
		for _, worker := range availableWorkers {
			if worker.Available() {
				foundWorker = job.SetWorkerAndMove(worker, j.Map)
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
		if !m.IsSelectedWall(wall.Sprite) {
			continue
		}

		hasFoundJob := false
		neighbors := m.SurroundingTilesFour(wall.Idx)
		for _, destination := range neighbors {
			if hasFoundJob {
				break
			}

			if m.IsColliding(j.Map, destination.Idx, destination.Dir) {
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

func (j *JobSystem) performWork() {
	for _, worker := range j.Workers {
		if !worker.HasJob() {
			continue
		}
		worker.PerformWork(j.Map)
	}
}
