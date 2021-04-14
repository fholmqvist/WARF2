package worldmap

// DrawHLine draws a horizontal line of specified
// sprite from index, to the right, for a number of
// n tiles.
func DrawHLine(mp *Map, idx, n, spr int) {
	for i := idx; i < idx+n; i++ {
		mp.Tiles[i].Sprite = spr
	}
}

// DrawVLine draws a vertical line of specified
// sprite from index, to the right, for a number of
// n tiles.
func DrawVLine(mp *Map, idx, n, spr int) {
	for i := idx; i < idx+TilesW*n; i += TilesW {
		mp.Tiles[i].Sprite = spr
	}
}

// DrawHRandomLine draws a horizontal line of a
// randomized sprite from index, to the right,
// for a number of n tiles.
func DrawHRandomLine(mp *Map, idx, n int, f func() int) {
	for i := idx; i < idx+n; i++ {
		mp.Tiles[i].Sprite = f()
	}
}

// DrawVRandomLine draws a vertical line of specified
// sprite from index, to the right, for a number of
// n tiles.
func DrawVRandomLine(mp *Map, idx, n int, f func() int) {
	for i := idx; i < idx+TilesW*n; i += TilesW {
		mp.Tiles[i].Sprite = f()
	}
}

func (m *Map) DrawRandomSquare(x1, y1, x2, y2 int, f func() int) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			idx := XYToIdx(x, y)
			m.Tiles[idx].Sprite = f()
		}
	}
}

func (m *Map) DrawSquareMutate(x1, y1, x2, y2 int, f func(int, int)) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			f(x, y)
		}
	}
}
