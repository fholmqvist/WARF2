package worldmap

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
	if t.Idx >= TilesBottom {
		return
	}

	if IsBoundary(t.Sprite) {
		if IsExposed(m.Tiles[OneTileDown(t.Idx)].Sprite) {
			m.Tiles[t.Idx].Sprite = BoundaryExposed
		} else {
			m.Tiles[t.Idx].Sprite = BoundarySolid
		}
	} else if IsWall(t.Sprite) {
		if IsExposed(m.Tiles[OneTileDown(t.Idx)].Sprite) {
			m.Tiles[t.Idx].Sprite = WallExposed
		} else {
			m.Tiles[t.Idx].Sprite = WallSolid
		}
	} else if IsSelectedWall(t.Sprite) {
		if IsExposed(m.Tiles[OneTileDown(t.Idx)].Sprite) {
			m.Tiles[t.Idx].Sprite = WallSelectedExposed
		} else {
			m.Tiles[t.Idx].Sprite = WallSelectedSolid
		}
	}
}
