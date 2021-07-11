package ui

import (
	"image/color"
	gl "projects/games/warf2/globals"

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

func (e Element) MouseIsOver(x, y int) bool {
	return (x >= e.X && x <= e.X+e.Width) &&
		(y >= e.Y && y <= e.Y+e.Height)
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

type ButtonTiled struct {
	///////////////////////
	// TODO
	// Height does nothing.
	///////////////////////
	Element
}

func (b ButtonTiled) Draw(screen *ebiten.Image, uiTiles *ebiten.Image, font font.Face) {
	// July 11 2021: Current golf record.
	drawTiledButton(
		screen, font, uiTiles,
		b.X+(gl.TilesW*b.Y), b.X+(gl.TilesW*b.Y)+gl.TilesW,
		b.X+(gl.TilesW*b.Y)+b.Width, b.X+(gl.TilesW*b.Y)+b.Width+gl.TilesW,
	)
	x := (b.X * gl.TileSize) + (b.Width*gl.TileSize)/2 - len(b.Text)*3
	y := (b.Y * gl.TileSize) + (b.Y / 2) + 4
	text.Draw(screen, b.Text, font, x, y, color.White)
}

type Dropdown struct {
	Main     ButtonTiled
	Buttons  []ButtonTiled
	hovering bool
}

func NewDropdown(text string, x, y, width int, buttons []ButtonTiled) Dropdown {
	return Dropdown{
		Main: ButtonTiled{
			Element{text, x, y, width, 1, color.White},
		},
		Buttons: buttons,
	}
}

func (d Dropdown) Draw(screen *ebiten.Image, uiTiles *ebiten.Image, font font.Face) {
	d.Main.Draw(screen, uiTiles, font)
	if !d.hovering {
		return
	}
	for _, button := range d.Buttons {
		button.Draw(screen, uiTiles, font)
	}
}

func (d *Dropdown) Hover(x, y int) {
	if d.IsNavigatingMenu(x, y) {
		return
	}
	if !d.Main.MouseIsOver(x, y) {
		d.hovering = false
		return
	}
	d.hovering = true
}

func (d Dropdown) IsNavigatingMenu(x, y int) bool {
	for _, button := range d.Buttons {
		if button.MouseIsOver(x, y) && d.hovering {
			return true
		}
	}
	return false
}
