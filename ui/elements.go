package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
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

func (b *Button) Select() {
	b.Color = color.Gray{50}
}

func (b *Button) Deselect() {
	b.Color = color.Gray{20}
}
