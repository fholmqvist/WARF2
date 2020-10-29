package jobsystem

// WorkerState defines the
// potential states for
// a given worker.
type WorkerState int

// Enum
const (
	WorkerIdle WorkerState = iota
	WorkerHasJob
	WorkerMovingTowards
	WorkerArrived
)
