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

func (l *LibraryRead) PerformWork(m *m.Map) bool {
	worker := *l.dwarf
	if !item.IsChair(m.Items[worker.GetPosition()].Sprite) && worker.GetState() != dwarf.WorkerMovingTowards {
		dst, ok := item.FindNearestChair(m, l.destination)
		if !ok {
			return false
		}
		worker.SetState(dwarf.WorkerMovingTowards)
		fmt.Println("Moving!")
		worker.MoveTo(dst, m)
		l.SetWorker(&worker)
		return false
	}
	fmt.Println("We're here!")
	if worker.GetState() != dwarf.WorkerArrived {
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

func (l *LibraryRead) GetWorker() *dwarf.Dwarf {
	return l.dwarf
}

func (l *LibraryRead) SetWorker(dw *dwarf.Dwarf) {
	l.dwarf = dw
}

func (l *LibraryRead) GetDestination() int {
	return l.destination
}
