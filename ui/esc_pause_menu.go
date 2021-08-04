package ui

import (
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type EscPauseMenu struct {
	Background Element
}

func NewEscPauseMenu() EscPauseMenu {
	return EscPauseMenu{
		Background: Element{
			X:               0,
			Y:               0,
			Width:           globals.ScreenWidth,
			Height:          globals.ScreenHeight,
			BackgroundColor: backgroundColor,
		},
	}
}

func (e *EscPauseMenu) DrawPause(screen *ebiten.Image, font font.Face) {
	e.Background.Draw(screen)
	text.Draw(screen, "PAUSED", font, CenterTextX("PAUSED"), globals.ScreenHeight/2, textColor)
}

func (e *EscPauseMenu) DrawESC(screen *ebiten.Image, font font.Face) {
	e.Background.Draw(screen)
	text.Draw(screen, "WARF", font, CenterTextX("WARF"), globals.ScreenHeight/2, textColor)
}
