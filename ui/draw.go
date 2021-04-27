package ui

import (
	"github.com/hajimehoshi/ebiten"
)

func DrawSquare(screen *ebiten.Image, e Element) {
	square, _ := ebiten.NewImage(e.Width, e.Height, ebiten.FilterDefault)
	_ = square.Fill(e.Color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.X), float64(e.Y))
	_ = screen.DrawImage(square, op)
}
