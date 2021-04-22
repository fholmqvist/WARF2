package jobsystem

import (
	"fmt"
	"projects/games/warf2/item"
	m "projects/games/warf2/worldmap"
)

type LibraryRead struct {
	worker      *Worker
	destination int
	readingTime int
}

func NewLibraryRead(destination, readingTime int) *LibraryRead {
	return &LibraryRead{nil, destination, readingTime}
}

func (l *LibraryRead) NeedsToBeRemoved(*m.Map) bool {
	return l.readingTime <= 0
}

func (l *LibraryRead) PerformWork(m *m.Map) bool {
	worker := *l.worker
	if !item.IsChair(m.Items[worker.GetPosition()].Sprite) && worker.GetState() != WorkerMovingTowards {
		dst, ok := item.FindNearestChair(m, l.destination)
		if !ok {
			return false
		}
		worker.SetState(WorkerMovingTowards)
		fmt.Println("Moving!")
		worker.MoveTo(dst, m)
		l.SetWorker(&worker)
		return false
	}
	fmt.Println("We're here!")
	if worker.GetState() != WorkerArrived {
		fmt.Println("Hasn't arrived yet")
		return false
	}
	if l.readingTime > 0 {
		fmt.Println("Reading!")
		l.readingTime--
		return false
	}
	fmt.Println("Done reading!")
	return true
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

func (l *LibraryRead) GetDestination() int {
	return l.destination
}
