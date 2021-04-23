package dwarf

// WorkerState defines the
// potential states for
// a given worker.
type WorkerState int

// Enum
const (
	WorkerIdle WorkerState = iota
	WorkerHasJob
	WorkerMoving
	WorkerArrived
)

func (w WorkerState) String() string {
	return []string{
		"WorkerIdle",
		"WorkerHasJob",
		"WorkerMoving",
		"WorkerArrived",
	}[w]
}
