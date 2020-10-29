package helpers

import (
	"image"

	e "projects/games/warf2/entity"
	m "projects/games/warf2/worldmap"

	"github.com/hajimehoshi/ebiten"
)

// DrawGraphic is a wrapper for drawing a tile to the screen
func DrawGraphic(idx, sprite int, screen *ebiten.Image, tileset *ebiten.Image, alpha float64) {
	const xNum = m.ScreenWidth / m.TileSize

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	op.GeoM.Translate(float64((idx%xNum)*m.TileSize), float64((idx/xNum)*m.TileSize))

	sx := (sprite % m.TilesetW) * m.TileSize
	sy := (sprite / m.TilesetW) * m.TileSize
	si := image.Rect(sx, sy, sx+m.TileSize, sy+m.TileSize)
	_ = screen.DrawImage(tileset.SubImage(si).(*ebiten.Image), op)
}

// DrawGraphics is a wrapper for drawing many tiles to the screen
func DrawGraphics(entities []e.Entity, screen *ebiten.Image, tileset *ebiten.Image) {
	for i := range entities {
		DrawGraphic(entities[i].Idx, entities[i].Sprite, screen, tileset, 1)
	}
}
