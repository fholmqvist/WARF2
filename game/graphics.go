package game

import (
	"image"
	"projects/games/warf2/globals"

	"github.com/hajimehoshi/ebiten"
)

// DrawGraphic is a wrapper for drawing a tile to the screen
func DrawGraphic(idx, sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha float64) {
	op := newOption(idx, alpha, 0)
	draw(idx, sprite, screen, tileset, alpha, op)
}

func DrawRailGraphic(idx, sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha, rotation float64) {
	if sprite == 0 {
		return
	}
	op := newOption(idx, alpha, rotation)
	draw(idx, sprite, screen, tileset, alpha, op)
}

func draw(idx, sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha float64, op *ebiten.DrawImageOptions) {
	sx := (sprite % globals.TilesetW) * globals.TileSize
	sy := (sprite / globals.TilesetW) * globals.TileSize
	si := image.Rect(sx, sy, sx+globals.TileSize, sy+globals.TileSize)
	_ = screen.DrawImage(tileset.SubImage(si).(*ebiten.Image), op)
}

func newOption(idx int, alpha, rotation float64) *ebiten.DrawImageOptions {
	const xNum = globals.ScreenWidth / globals.TileSize
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	op.GeoM.Translate(-globals.TileSize/2, -globals.TileSize/2)
	op.GeoM.Rotate(rotation)
	op.GeoM.Translate(globals.TileSize/2, globals.TileSize/2)
	op.GeoM.Translate(float64((idx%xNum)*globals.TileSize), float64((idx/xNum)*globals.TileSize))
	return op
}
