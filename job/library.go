package job

import (
	"fmt"
	"projects/games/warf2/dwarf"
	"projects/games/warf2/item"
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

func (l *LibraryRead) PerformWork(m *m.Map) (finished bool) {
	if !item.IsChair(m.Items[l.dwarf.GetPosition()].Sprite) && l.dwarf.GetState() != dwarf.WorkerMoving {
		dst, ok := item.FindNearestChair(m, l.destination)
		if !ok {
			return false
		}
		l.destination = dst
		if l.dwarf.MoveTo(dst, m) {
			return false
		}
		fmt.Println("MoveTo went wrong")
		return false
	}
	if !item.IsChair(m.Items[l.dwarf.GetPosition()].Sprite) && l.dwarf.GetState() == dwarf.WorkerMoving {
		fmt.Println("Haven't arrived yet")
		return false
	}
	if l.readingTime > 0 {
		l.readingTime--
		return false
	}
	return true
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
