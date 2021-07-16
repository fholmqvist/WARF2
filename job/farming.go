package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/item"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Farming struct {
	FarmID       int
	dwarf        *dwarf.Dwarf
	destinations []int
}

func NewFarming(farmID int, destinations []int) *Farming {
	return &Farming{farmID, nil, destinations}
}

func (d *Farming) NeedsToBeRemoved(mp *m.Map) bool {
	return len(d.destinations) == 0
}

func (d *Farming) Finish(*m.Map, *room.Service) {
	if d.dwarf == nil {
		return
	}
	d.dwarf.SetToAvailable()
	d.dwarf = nil
}

// Ran on arrival.
func (f *Farming) PerformWork(mp *m.Map, dwarves []*dwarf.Dwarf) bool {
	if len(f.destinations) == 0 {
		return true
	}
	if f.dwarf == nil {
		return false
	}
	lastIdx := f.destinations[len(f.destinations)-1]
	if f.dwarf.Idx != lastIdx {
		return false
	}
	mp.Items[lastIdx].Sprite = item.Wheat
	f.destinations = f.destinations[:len(f.destinations)-1]
	////////////
	// TODO
	// Teleport!
	////////////
	if len(f.destinations) == 0 {
		return true
	}
	if len(f.destinations) == 1 {
		f.dwarf.Idx = f.destinations[0]
	} else {
		f.dwarf.Idx = f.destinations[len(f.destinations)-1]
	}
	return false
}

func (f *Farming) Priority() int {
	return 1
}

func (f *Farming) GetWorker() *dwarf.Dwarf {
	return f.dwarf
}

func (f *Farming) SetWorker(dw *dwarf.Dwarf) {
	f.dwarf = dw
}

func (f *Farming) GetDestinations() []int {
	return f.destinations
}

func (f *Farming) GetWallIdx() int {
	return -1
}

func (f *Farming) String() string {
	return "Farming"
}
