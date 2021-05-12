package worldmap

import "projects/games/warf2/globals"

// IdxToXY returns the corresponding
// X and Y values for a given index.
func IdxToXY(idx int) (int, int) {
	return IdxToX(idx), IdxToY(idx)
}

// IdxToX returns the corresponding
// X value to for given index.
func IdxToX(idx int) int {
	return idx % globals.TilesW
}

// IdxToY returns the corresponding
// Y value to for given index.
func IdxToY(idx int) int {
	return idx / globals.TilesW
}

// XYToIdx returns the corresponding
// index based on the given X and Y values.
func XYToIdx(x, y int) int {
	return x + y*globals.TilesW
}
