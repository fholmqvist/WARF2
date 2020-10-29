package jobsystem

// JobState defines the
// potential states of
// a given job.
type JobState int

// JobState enum
const (
	New JobState = iota
	Ongoing
	Broken
	Done
)
