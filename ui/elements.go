package ui

import (
	"image/color"

	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/mouse"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

// Element wraps data for UI elements.
type Element struct {
	Text            string
	X               int
	Y               int
	Width           int
	Height          int
	TextColor       color.Color
	BackgroundColor color.Color
}

func (e Element) Draw(screen *ebiten.Image) {
	DrawSquare(screen, e)
}

func (e Element) DrawWithText(screen *ebiten.Image, font font.Face) {
	e.Draw(screen)
	text.Draw(screen, e.Text, font, (e.X+e.Width/2)-len(e.Text)*4, (e.Y+e.Height/2)+4, e.TextColor)
}

func (e Element) MouseIsOver(x, y int) bool {
	return (x >= e.X && x <= e.X+e.Width) &&
		(y >= e.Y && y <= e.Y+e.Height)
}

type Button struct {
	///////////////////////
	// TODO
	// Height does nothing.
	///////////////////////
	Element
	hovering bool
}

func (b Button) Draw(screen *ebiten.Image, uiTiles *ebiten.Image, font font.Face) {
	// July 11 2021: Current golf record.
	drawButtonTiles(
		screen, font, uiTiles,
		b.X+(gl.TilesW*b.Y), b.X+(gl.TilesW*b.Y)+gl.TilesW,
		b.X+(gl.TilesW*b.Y)+b.Width, b.X+(gl.TilesW*b.Y)+b.Width+gl.TilesW,
		b.hovering,
	)
	x := (b.X * gl.TileSize) + (b.Width*gl.TileSize)/2 - len(b.Text)*3
	y := (b.Y * gl.TileSize) + gl.TileSize + gl.TileSize/4
	text.Draw(screen, b.Text, font, x, y, textColor)
}

type Dropdown struct {
	Main    Button
	Buttons []Button
}

func NewDropdown(text string, x, y, width int, buttons []Button) Dropdown {
	return Dropdown{
		Main: Button{
			Element:  Element{text, x, y, width, 1, textColor, backgroundColor},
			hovering: false,
		},
		Buttons: buttons,
	}
}

func (d Dropdown) Draw(screen *ebiten.Image, uiTiles *ebiten.Image, font font.Face) {
	d.Main.Draw(screen, uiTiles, font)
	if !d.Main.hovering {
		return
	}
	for _, button := range d.Buttons {
		button.Draw(screen, uiTiles, font)
	}
}

func (d *Dropdown) Update(x, y int) (mode mouse.Mode, clicked bool) {
	if d.isNavigatingMenu(x, y) {
		return d.handleMenuNavigation(x, y)
	}
	if !d.Main.MouseIsOver(x, y) {
		d.Main.hovering = false
		return
	}
	d.Main.hovering = true
	return
}

func (d Dropdown) isNavigatingMenu(x, y int) bool {
	if !d.Main.hovering {
		return false
	}
	for _, b := range d.Buttons {
		if b.MouseIsOver(x, y) {
			return true
		}
	}
	return false
}

func (d *Dropdown) handleMenuNavigation(x, y int) (mode mouse.Mode, clicked bool) {
	for idx, b := range d.Buttons {
		if !b.MouseIsOver(x, y) {
			d.Buttons[idx].hovering = false
			continue
		}
		d.Buttons[idx].hovering = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			return mouse.ModeFromString[b.Text], true
		}
	}
	return
}

func drawButtonTiles(screen *ebiten.Image, gameFont font.Face, uiTiles *ebiten.Image, tl, bl, tr, br int, highlighted bool) {
	// Wrapper to reduce noice.
	draw := func(sprite, idx int) {
		gl.DrawTile(sprite, screen, uiTiles, 1, gl.DrawOptions(idx, 1, 0))
	}
	var (
		tlc = TopLeft
		blc = BottomLeft
		tc  = Top
		bc  = Bottom
		trc = TopRight
		brc = BottomRight
	)
	if highlighted {
		tlc = Highlighted_TopLeft
		blc = Highlighted_BottomLeft
		tc = Highlighted_Top
		bc = Highlighted_Bottom
		trc = Highlighted_TopRight
		brc = Highlighted_BottomRight
	}
	// Left.
	draw(tlc, tl)
	draw(blc, bl)
	// Middle.
	for i := 1; i < tr-tl; i++ {
		draw(tc, tl+i)
		draw(bc, tl+gl.TilesW+i)
	}
	// Right.
	draw(trc, tr)
	draw(brc, br)
}
