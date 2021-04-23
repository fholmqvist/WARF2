package job

import (
	"projects/games/warf2/dwarf"
	"projects/games/warf2/item"
	"projects/games/warf2/worldmap"
	m "projects/games/warf2/worldmap"
)

type LibraryRead struct {
	dwarf       *dwarf.Dwarf
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
	if shouldGetChair(m, l) {
		dst, ok := item.FindNearestChair(m, l.destination)
		if !ok {
			return unfinished
		}
		l.destination = dst
		if l.dwarf.MoveTo(dst, m) {
			return unfinished
		}
		return unfinished
	}
	if l.readingTime > 1 {
		l.readingTime--
		return unfinished
	}
	l.readingTime = 0
	return finished
}

func (l *LibraryRead) Priority() int {
	return 0
}

func (l *LibraryRead) GetWorker() *dwarf.Dwarf {
	return l.dwarf
}

func (l *LibraryRead) SetWorker(dw *dwarf.Dwarf) {
	l.dwarf = dw
}

func (l *LibraryRead) GetDestination() int {
	return l.destination
}

func shouldGetChair(m *worldmap.Map, l *LibraryRead) bool {
	return !item.IsChair(m.Items[l.dwarf.GetPosition()].Sprite) &&
		l.dwarf.GetState() != dwarf.WorkerMoving
}
