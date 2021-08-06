package job

import (
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func getNextIdx(destinations []int) int {
	if len(destinations) == 1 {
		return destinations[0]
	}
	return destinations[len(destinations)-1]
}

func getPath(mp *m.Map, next int, dwarfIdx int) ([]int, bool) {
	return m.CreatePath(
		&mp.Tiles[dwarfIdx],
		&mp.Tiles[next],
	)
}
