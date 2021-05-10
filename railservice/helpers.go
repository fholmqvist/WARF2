package rail

import mp "projects/games/warf2/worldmap"

func IsRail(m *mp.Map, idx int) bool {
	return m.Rails[idx].Sprite >= Straight &&
		m.Rails[idx].Sprite <= Cross
}
