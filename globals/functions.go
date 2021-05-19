package globals

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
)

func DrawTile(sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha float64, op *ebiten.DrawImageOptions) {
	sx := (sprite % TilesetW) * TileSize
	sy := (sprite / TilesetW) * TileSize
	si := image.Rect(sx, sy, sx+TileSize, sy+TileSize)
	_ = screen.DrawImage(tileset.SubImage(si).(*ebiten.Image), op)
}

func DrawOptions(idx int, alpha, rotation float64) *ebiten.DrawImageOptions {
	const xNum = ScreenWidth / TileSize
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	op.GeoM.Translate(-TileSize/2, -TileSize/2)
	op.GeoM.Rotate(rotation)
	op.GeoM.Translate(TileSize/2, TileSize/2)
	op.GeoM.Translate(float64((idx%xNum)*TileSize), float64((idx/xNum)*TileSize))
	return op
}

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
