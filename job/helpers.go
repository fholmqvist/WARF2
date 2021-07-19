package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func getNextIdx(destinations []int) int {
	if len(destinations) == 1 {
		return destinations[0]
	}
	return destinations[len(destinations)-1]
}

func getPath(mp *m.Map, next int, dwarf *dwarf.Dwarf) ([]int, bool) {
	return dwarf.CreatePath(
		&mp.Tiles[dwarf.Idx],
		&mp.Tiles[next],
	)
}
