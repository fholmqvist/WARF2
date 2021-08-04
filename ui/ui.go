// Package ui handles all the
// in-game graphical overlays.
package ui

import (
	"image/color"

	gl "github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/mouse"
)

var (
	transparentColor = color.Transparent
	backgroundColor  = color.RGBA{0, 0, 0, 100}
	textColor        = color.White
)

// UI wraps all the UI elements for Game
type UI struct {
	*MainMenu
	*HelpMenu
	*GameplayUI
}

func GenerateUI() UI {
	return UI{
		MainMenu:   NewMainMenu(),
		HelpMenu:   NewHelpMenu(),
		GameplayUI: NewGameplayUI(),
	}
}

func (ui *UI) UpdateGameplayMenus() (mouse.Mode, bool) {
	mousePos := mouse.MouseIdx()
	x, y := gl.IdxToXY(mousePos)
	mode, clicked := ui.BuildMenu.Update(x, y)
	if !clicked {
		return 0, false
	}
	return mode, true
}
