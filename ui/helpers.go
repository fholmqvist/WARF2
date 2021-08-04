package ui

import (
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/hajimehoshi/ebiten"
)

func CenterTextX(text string) int {
	return globals.ScreenWidth/2 - (len(text)*globals.TileSize/2)/2
}

func DrawSquare(screen *ebiten.Image, e Element) {
	square, _ := ebiten.NewImage(e.Width, e.Height, ebiten.FilterDefault)
	if e.BackgroundColor != nil {
		_ = square.Fill(e.BackgroundColor)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.X), float64(e.Y))
	_ = screen.DrawImage(square, op)
}
