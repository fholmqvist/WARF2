package worldmap

import (
	"math/rand"

	gl "github.com/Holmqvist1990/WARF2/globals"
)

func (m *Map) CreateBoundaryWalls() {
	m.DrawOutline(0, 0, gl.TilesW, gl.TilesH, BoundarySolid)
}

func (m *Map) CreateOutmostWalls() {
	m.DrawOutline(1, 1, gl.TilesW-1, gl.TilesH-1, WallSolid)
}

func (m *Map) RandomizeWalls(chance int) {
	for i := range m.Tiles {
		risk := rand.Intn(100)
		if risk < chance {
			m.Tiles[i].Sprite = WallSolid
		}
	}
}

// FixWalls sets the graphic for all
// wall types so that solid and exposed
// variants match with the surrounding
// environment.
func (m *Map) FixWalls() {
	for _, t := range m.Tiles {
		m.FixWall(&t)
	}
}

// FixWall sets the graphic for any
// wall type at the current index so
// that solid and exposed variants match
// with the surrounding environment.
func (m *Map) FixWall(t *Tile) {
	if t.Idx >= gl.TilesBottom {
		return
	}
	if IsBoundary(t.Sprite) {
		if IsExposed(m.Tiles[OneDown(t.Idx)].Sprite) {
			m.Tiles[t.Idx].Sprite = BoundaryExposed
		} else {
			m.Tiles[t.Idx].Sprite = BoundarySolid
		}
		return
	}
	if IsWall(t.Sprite) {
		if IsExposed(m.Tiles[OneDown(t.Idx)].Sprite) {
			m.Tiles[t.Idx].Sprite = WallExposed
		} else {
			m.Tiles[t.Idx].Sprite = WallSolid
		}
		return
	}
	if IsSelectedWall(t.Sprite) {
		if IsExposed(m.Tiles[OneDown(t.Idx)].Sprite) {
			m.Tiles[t.Idx].Sprite = WallSelectedExposed
		} else {
			m.Tiles[t.Idx].Sprite = WallSelectedSolid
		}
		return
	}
}
