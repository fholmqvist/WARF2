package jobsystem

import (
	m "projects/games/warf2/worldmap"
)

type LibraryRead struct {
	worker      *Worker
	destination int
	readingTime int
}

func NewLibraryRead(w Worker, destination, readingTime int) *LibraryRead {
	return &LibraryRead{&w, destination, readingTime}
}

func (l *LibraryRead) NeedsToBeRemoved(*m.Map) bool {
	return l.readingTime <= 0
}

func (l *LibraryRead) PerformWork(*m.Map) bool {
	/////////////////////////////////////////////////
	// TODO
	//
	// 1. Check that we are next to a bookshelf.
	// 2. Find nearest available seat.
	// 3. If no seat, start reading at slower speed.
	// 4. If seat, go to seat and start reading.
	/////////////////////////////////////////////////
	if l.readingTime > 0 {
		l.readingTime--
		return false
	}
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
