package tests

import (
	"projects/games/warf2/mouse"
	"projects/games/warf2/worldmap"
	"testing"
)

var m = worldmap.Map{
	Tiles:         make([]worldmap.Tile, 9),
	SelectedTiles: make([]worldmap.Tile, 9),
}

func TestTileRange(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		affected []int
	}{
		{"downright", 0, 4, []int{0, 1, 3, 4}},
		{"downleft", 1, 3, []int{0, 1, 3, 4}},
		{"upright", 3, 1, []int{0, 1, 3, 4}},
		{"upleft", 4, 0, []int{0, 1, 3, 4}},
	}

	for _, tt := range tests {
		copy := m

		mouse.FuncOverRange(&copy, tt.start, tt.end, func(mp *worldmap.Map, x, y int) {
			idx := worldmap.XYToIdx(x, y)
			mp.Tiles[idx].Sprite = -1
		})

		for _, i := range tt.affected {
			if copy.Tiles[i].Sprite != -1 {
				t.Fatalf("%s, %d", tt.name, tt.affected)
			}
		}
	}

}
