package dwarf

// WorkState defines the
// potential states for
// a given worker.
type WorkState int

// Enum
const (
	Idle WorkState = iota
	HasJob
	Moving
	Arrived
)

func (w WorkState) String() string {
	return []string{
		"Idle",
		"HasJob",
		"Moving",
		"Arrived",
	}[w]
}
