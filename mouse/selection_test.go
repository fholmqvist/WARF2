package mouse

import (
	"testing"

	gl "github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var mp = m.Map{
	Tiles:         make([]m.Tile, 9),
	SelectedTiles: make([]m.Tile, 9),
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
		copy := mp

		FuncOverRange(&copy, tt.start, tt.end, func(mp *m.Map, x, y int) {
			idx := gl.XYToIdx(x, y)
			mp.Tiles[idx].Sprite = -1
		})

		for _, i := range tt.affected {
			if copy.Tiles[i].Sprite != -1 {
				t.Fatalf("%s, %d", tt.name, tt.affected)
			}
		}
	}
}
