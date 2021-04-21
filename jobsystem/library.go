package jobsystem

import m "projects/games/warf2/worldmap"

type LibraryRead struct {
	worker      *Worker
	state       JobState
	destination int
}

func (l *LibraryRead) NeedsToBeRemoved(*m.Map) bool {
	return l.state == Done
}

func (l *LibraryRead) PerformWork(*m.Map) func() bool {
	return func() bool { return false }
}

func (l *LibraryRead) GetDestination() int {
	return l.destination
}

func (l *LibraryRead) Priority() int {
	return 0
}

func (l *LibraryRead) GetWorker() *Worker {
	return l.worker
}

func (l *LibraryRead) SetWorker(w *Worker) {
	l.worker = w
}

func (l *LibraryRead) GetState() JobState {
	return l.state
}

func (l *LibraryRead) SetState(j JobState) {
	l.state = j
}
