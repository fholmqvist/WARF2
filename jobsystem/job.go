package jobsystem

// Job declares the common interface
// for jobs, in order to be used within
// the job system.
type Job interface {
	WaitingForWorker() bool
	SetWorker(*Worker) bool
	OnGoing() bool
	Done() bool
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
}

func (j *JobSystem) removeFinishedJobs() {
	var jobs []Job
	for _, job := range j.Jobs {
		if !job.Done() {
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
				foundWorker = job.SetWorker(&worker)
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
