package ui

import (
	"image/color"

	"github.com/Holmqvist1990/WARF2/dwarf"
	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/mouse"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type GameplayUI struct {
	MouseMode   Element
	BuildMenu   Dropdown
	OverviewTab Element
}

func NewGameplayUI() *GameplayUI {
	buildMenuButtons := buildMenuButtons()
	overviewTab := Element{
		Text:            "Hover",
		X:               0,
		Y:               0,
		Width:           gl.TilesW * gl.TileSize,
		Height:          gl.TileSize,
		TextColor:       textColor,
		BackgroundColor: backgroundColor,
	}
	return &GameplayUI{
		MouseMode:   NewMouseOverlay(),
		BuildMenu:   NewDropdown("Build", 34, 32, 11, buildMenuButtons),
		OverviewTab: overviewTab,
	}
}

func (g *GameplayUI) Draw(screen *ebiten.Image, uiTiles *ebiten.Image, font font.Face, ui *UI, dw []*dwarf.Dwarf) {
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		drawOverview(screen, font, dw, ui.MouseMode)
	}
	drawMouseMode(screen, uiTiles, font, ui.MouseMode)
	ui.BuildMenu.Draw(screen, uiTiles, font)
	if ui.OverviewTab.Text == "" {
		ui.OverviewTab.BackgroundColor = transparentColor
	} else {
		ui.OverviewTab.BackgroundColor = backgroundColor
	}
	ui.OverviewTab.DrawWithText(screen, font)
}

func NewMouseOverlay() Element {
	return Element{
		X:               gl.TileSize,
		Y:               (gl.TileSize * gl.TilesH) + gl.TileSize + 4,
		TextColor:       textColor,
		BackgroundColor: backgroundColor,
	}
}

func buildMenuButtons() []Button {
	texts := []string{
		mouse.Normal.String(),
		mouse.Storage.String(),
		mouse.SleepHall.String(),
		mouse.Farm.String(),
		mouse.Brewery.String(),
		mouse.Library.String(),
		mouse.Delete.String(),
	}
	offset := len(texts) * 2
	buildMenuButtons := make([]Button, len(texts))
	for i, text := range texts {
		b := Button{
			Element: Element{
				Text:            text,
				X:               34,
				Y:               32 - offset,
				Width:           11,
				Height:          1,
				TextColor:       textColor,
				BackgroundColor: backgroundColor,
			},
		}
		buildMenuButtons[i] = b
		offset -= 2
	}
	return buildMenuButtons
}

func drawMouseMode(screen *ebiten.Image, uiTiles *ebiten.Image, gameFont font.Face, mouseMode Element) {
	drawButtonTiles(
		screen, gameFont, uiTiles,
		gl.TilesT, gl.TilesT+gl.TilesW,
		gl.TilesT+gl.TilesW-1-12, gl.TilesT+(gl.TilesW*2)-1-12,
		false, // Unselectable
	)
	text.Draw(screen, mouseMode.Text, gameFont, mouseMode.X, mouseMode.Y, mouseMode.TextColor)
}

func drawOverview(screen *ebiten.Image, gameFont font.Face, dw []*dwarf.Dwarf, mouseMode Element) {
	x, y, xo, yo := 20, 20, 10, 20
	drawBackground(screen, x, y, (len(dw)*y)+y+yo)
	text.Draw(screen, "Dwarves:", gameFont, x+xo, y*2, mouseMode.BackgroundColor)
	for i, d := range dw {
		text.Draw(screen, d.Characteristics.Name, gameFont, (x*2)+xo, yo+(y*(i+2)), mouseMode.TextColor)
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
		color.Opaque,
		color.Gray{50},
	}
	DrawSquare(screen, e)
}
