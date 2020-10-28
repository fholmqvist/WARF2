package gmap

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
