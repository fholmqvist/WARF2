package jobsystem

// Worker defines the interface
// for all characters who are
// eligible workers.
type Worker interface {
	Available() bool
}
