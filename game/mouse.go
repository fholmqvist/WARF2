package game

import (
	"fmt"

	m "projects/games/warf2/gmap"

	"github.com/hajimehoshi/ebiten"
)

// MouseMode enum for managing mouse action state.
type MouseMode int

// MouseMode enum.
const (
	None MouseMode = iota
)

func handleMouse(g *Game) {
	mouseHover(g)

	idx := mousePos()

	if idx < 0 || idx > m.TilesT {
		return
	}

	g.mousePos = idx

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if g.debug {
			fmt.Println("tile:", idx, m.GraphicName(g.Gmap.Tiles[idx].Idx))
		}

		mouseClick(g, idx)

	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.mouseMode = None
	}
}

func mouseClick(g *Game, idx int) {
	// s := &g.Gmap.Tiles[idx].Sprite

	switch g.mouseMode {

	case None:
		// Identity

	default:
		fmt.Println("mouseClick: unknown MouseMode:", g.mouseMode)
	}
}

func mouseHover(g *Game) {
	// idx := mousePos()

	switch g.mouseMode {

	default:
	}
}

func mousePos() int {
	mx, my := ebiten.CursorPosition()
	mx, my = mx/m.TileSize, my/m.TileSize

	return mx + (my * m.TilesW)
}
