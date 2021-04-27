package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

// Element wraps data for UI elements
type Element struct {
	Text   string
	X      int
	Y      int
	Width  int
	Height int
	Color  color.Color
}

func (e Element) Draw(screen *ebiten.Image) {
	DrawSquare(screen, e)
}

type Button struct {
	Element
}

func (b Button) Draw(screen *ebiten.Image, font font.Face) {
	b.Element.Draw(screen)
	text.Draw(screen, b.Text, font, (b.X+b.Width/2)-len(b.Text)*4, (b.Y+b.Height/2)+4, color.White)
}

func (b *Button) Select() {
	b.Color = color.Gray{50}
}

func (b *Button) Deselect() {
	b.Color = color.Gray{20}
}
