package ui

import (
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type HelpMenu struct {
	Description string
	Back        Button
	Clickable   bool
}

func NewHelpMenu() *HelpMenu {
	return &HelpMenu{
		Description: "Currently no help. WIP.",
		Back: Button{Element{
			Text:      "Back",
			X:         xOffset,
			Y:         globals.TilesH / 2,
			Width:     width,
			Height:    1,
			TextColor: textColor,
		}, false},
	}
}

func (h *HelpMenu) Draw(screen *ebiten.Image, uiTiles *ebiten.Image, font font.Face) {
	text.Draw(screen, h.Description, font, xOffset*globals.TileSize, globals.TileSize*2, textColor)
	h.Back.Draw(screen, uiTiles, font)
}

func (h *HelpMenu) Update() bool {
	x, y := ebiten.CursorPosition()
	x /= globals.TileSize
	y /= globals.TileSize
	if h.Back.MouseIsOver(x, y) {
		h.Back.hovering = true
		if !h.Clickable {
			return false
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			h.Clickable = false
			return true
		}
		return false
	}
	h.Back.hovering = false
	return false
}
