package globals

import (
	"math"
)

func Dist(ax, ay, bx, by int) float64 {
	xDist := math.Abs(float64(bx - ax))
	yDist := math.Abs(float64(by - ay))
	return xDist + yDist
}

// IdxToXY returns the corresponding
// X and Y values for a given index.
func IdxToXY(idx int) (int, int) {
	return IdxToX(idx), IdxToY(idx)
}

// IdxToX returns the corresponding
// X value to for given index.
func IdxToX(idx int) int {
	return idx % TilesW
}

// IdxToY returns the corresponding
// Y value to for given index.
func IdxToY(idx int) int {
	return idx / TilesW
}

// XYToIdx returns the corresponding
// index based on the given X and Y values.
func XYToIdx(x, y int) int {
	return x + y*TilesW
}
