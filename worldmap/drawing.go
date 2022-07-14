package worldmap

import gl "github.com/Holmqvist1990/WARF2/globals"

// Draws a horizontal line of specified
// sprite from index, to the right,
// for a number of n tiles.
func DrawHLineIdx(mp *Map, idx, n, spr int) {
	for i := idx; i < idx+n; i++ {
		mp.Tiles[i].Sprite = spr
	}
}

// Draws a vertical line of specified
// sprite from index, to the right,
// for a number of n tiles.
func DrawVLineIdx(mp *Map, idx, n, spr int) {
	for i := idx; i < idx+gl.TilesW*n; i += gl.TilesW {
		mp.Tiles[i].Sprite = spr
	}
}

// Draws a horizontal line of a
// randomized sprite from index,
// to the right, for a number of n tiles.
func DrawHRandomLineIdx(mp *Map, idx, n int, f func() int) {
	for i := idx; i < idx+n; i++ {
		mp.Tiles[i].Sprite = f()
	}
}

// Draws a vertical line of specified
// sprite from index, to the right,
// for a number of n tiles.
func DrawVRandomLineIdx(mp *Map, idx, n int, f func() int) {
	for i := idx; i < idx+gl.TilesW*n; i += gl.TilesW {
		mp.Tiles[i].Sprite = f()
	}
}

// Draws a square of sprite.
func (m *Map) DrawSquare(x1, y1, x2, y2, sprite int) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			m.Tiles[gl.XYToIdx(x, y)].Sprite = sprite
		}
	}
}

// Draws a square of random sprites.
func (m *Map) DrawRandomSquare(x1, y1, x2, y2 int, f func() int) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			m.Tiles[gl.XYToIdx(x, y)].Sprite = f()
		}
	}
}

// Draws a square of specified sprite.
func (m *Map) DrawSquareSprite(x1, y1, x2, y2, sprite int) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			m.Tiles[gl.XYToIdx(x, y)].Sprite = sprite
		}
	}
}

// Draws a square outline of sprite.
func (m *Map) DrawOutline(x1, y1, x2, y2, sprite int) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			if x == x1 || x == x2-1 || y == y1 || y == y2-1 {
				m.Tiles[gl.XYToIdx(x, y)].Sprite = sprite
			}
		}
	}
}
