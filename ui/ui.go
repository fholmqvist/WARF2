// Package ui handles all the
// in-game graphical overlays.
package ui

import (
	"image/color"
	"projects/games/warf2/dwarf"
	gl "projects/games/warf2/globals"
	"projects/games/warf2/mouse"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

var textColor = color.White

// UI wraps all the UI elements for Game
type UI struct {
	*MainMenu
	MouseMode Element
	BuildMenu Dropdown
}

func GenerateUI() UI {
	buildMenuButtons := []ButtonTiled{
		{Element: Element{Text: "Delete", X: 34, Y: 32 - 2, Width: 11, Height: 1, Color: textColor}},
		{Element: Element{Text: "Library", X: 34, Y: 32 - 4, Width: 11, Height: 1, Color: textColor}},
		{Element: Element{Text: "Storage", X: 34, Y: 32 - 6, Width: 11, Height: 1, Color: textColor}},
		{Element: Element{Text: "Wall", X: 34, Y: 32 - 8, Width: 11, Height: 1, Color: textColor}},
	}
	return UI{
		MouseMode: NewMouseOverlay(),
		MainMenu:  NewMainMenu(),
		BuildMenu: NewDropdown("Build", 34, 32, 11, buildMenuButtons),
	}
}

func NewMouseOverlay() Element {
	return Element{
		X:     gl.TileSize,
		Y:     (gl.TileSize * gl.TilesH) + gl.TileSize + 4,
		Color: textColor,
	}
}

func (ui *UI) DrawGameplay(screen *ebiten.Image, gameFont font.Face, dw []*dwarf.Dwarf, uiTiles *ebiten.Image) {
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		drawOverview(screen, gameFont, dw, ui.MouseMode)
	}
	drawMouseMode(screen, uiTiles, gameFont, ui.MouseMode)
	ui.BuildMenu.Draw(screen, uiTiles, gameFont)
}

func (ui *UI) UpdateGameplayMenu() {
	mousePos := mouse.MouseIdx()
	x, y := gl.IdxToXY(mousePos)
	ui.BuildMenu.Update(x, y)
}

func drawMouseMode(screen *ebiten.Image, uiTiles *ebiten.Image, gameFont font.Face, mouseMode Element) {
	drawTiledButton(
		screen, gameFont, uiTiles,
		gl.TilesT, gl.TilesT+gl.TilesW,
		gl.TilesT+gl.TilesW-1-12, gl.TilesT+(gl.TilesW*2)-1-12,
		false, // Unselectable
	)
	text.Draw(screen, mouseMode.Text, gameFont, mouseMode.X, mouseMode.Y, mouseMode.Color)
}

func drawOverview(screen *ebiten.Image, gameFont font.Face, dw []*dwarf.Dwarf, mm Element) {
	x, y, xo, yo := 20, 20, 10, 20
	drawBackground(screen, x, y, (len(dw)*y)+y+yo)
	text.Draw(screen, "Dwarves:", gameFont, x+xo, y*2, mm.Color)
	for i, d := range dw {
		text.Draw(screen, d.Characteristics.Name, gameFont, (x*2)+xo, yo+(y*(i+2)), mm.Color)
	}
}

func drawBackground(screen *ebiten.Image, x, y, height int) {
	sw, _ := screen.Size()
	e := Element{
		"",
		x,
		y,
		sw - 40,
		height,
		color.Gray{50},
	}
	DrawSquare(screen, e)
}

func drawTiledButton(screen *ebiten.Image, gameFont font.Face, uiTiles *ebiten.Image, tl, bl, tr, br int, highlighted bool) {
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
