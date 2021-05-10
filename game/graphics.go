package game

import (
	"image"

	e "projects/games/warf2/entity"
	m "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
)

// DrawGraphic is a wrapper for drawing a tile to the screen
func DrawGraphic(idx, sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha float64) {
	if sprite == 0 {
		return
	}
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

// DrawGraphics is a wrapper for drawing many tiles to the screen
func DrawGraphics(entities []e.Entity, screen *ebiten.Image, tileset *ebiten.Image) {
	for i := range entities {
		DrawGraphic(entities[i].Idx, entities[i].Sprite, screen, tileset, 1)
	}
}

func draw(idx, sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha float64, op *ebiten.DrawImageOptions) {
	sx := (sprite % m.TilesetW) * m.TileSize
	sy := (sprite / m.TilesetW) * m.TileSize
	si := image.Rect(sx, sy, sx+m.TileSize, sy+m.TileSize)
	_ = screen.DrawImage(tileset.SubImage(si).(*ebiten.Image), op)
}

func newOption(idx int, alpha, rotation float64) *ebiten.DrawImageOptions {
	const xNum = m.ScreenWidth / m.TileSize
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	op.GeoM.Translate(-m.TileSize/2, -m.TileSize/2)
	op.GeoM.Rotate(rotation)
	op.GeoM.Translate(m.TileSize/2, m.TileSize/2)
	op.GeoM.Translate(float64((idx%xNum)*m.TileSize), float64((idx/xNum)*m.TileSize))
	return op
}
