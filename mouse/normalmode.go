package mouse

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/dwarf"
	gl "github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func (s *System) noneMode(mp *m.Map, dwarves *[]*dwarf.Dwarf, currentMousePos int) {
	mp.ClearSelectedTiles()
	s.clickFunctions(mp, currentMousePos,
		func() {
			printMousePos(currentMousePos)
			printDwarf(dwarves, currentMousePos)
			// Get tile from real tiles.
			tile, ok := mp.GetTileByIndex(currentMousePos)
			if !ok {
				return
			}
			s.firstClickedSprite = tile.Sprite
			// Replace that tile with one from SelectedTiles.
			tile, ok = mp.GetSelectionTileByIndex(currentMousePos)
			if !ok {
				return
			}
			// Selecting a non-wall defaults to
			// wall in order to enable wall selection
			// without having first clicked on a wall.
			if !m.IsSelectedWall(s.firstClickedSprite) {
				s.firstClickedSprite = m.WallSolid
			}
			if m.IsWallOrSelected(tile.Sprite) {
				tile.Sprite = invertSelected(s.firstClickedSprite)
			}
		},
		func(mp *m.Map, x int, y int) {
			selectionWalls(s, mp, x, y)
		})
}

// Attempting to collapse these three similar
// functions into one just made the interface
// that much more complicated. Sometimes, not
// having DRY everywhere ain't that bad.
func (s *System) mouseUpSetWalls(mp *m.Map, x, y int) {
	selectionTile, ok := mp.GetSelectionTile(x, y)
	if !ok {
		return
	}
	// No change.
	if m.IsNone(selectionTile.Sprite) {
		return
	}
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}
	setWalls(s, tile)
	selectionTile.Sprite = m.None
}

func selectionWalls(s *System, mp *m.Map, x, y int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}
	if !m.IsWallOrSelected(tile.Sprite) {
		return
	}
	selectionTile, ok := mp.GetSelectionTile(x, y)
	if !ok {
		return
	}
	// In order to invert between (un)selected.
	selectionTile.Sprite = tile.Sprite
	setWalls(s, selectionTile)
}

func setWalls(s *System, tile *m.Tile) {
	if !m.IsWallOrSelected(tile.Sprite) {
		return
	}
	if m.IsWall(s.firstClickedSprite) {
		setToSelected(tile)
		return
	}
	setToNormalInteractFalse(tile)
}

func invertSelected(sprite int) int {
	if m.IsWall(sprite) {
		if sprite == m.WallSolid {
			return m.WallSelectedSolid
		}
		return m.WallSelectedExposed
	}
	if sprite == m.WallSelectedSolid {
		return m.WallSolid
	}
	return m.WallExposed
}

func setToSelected(tile *m.Tile) {
	if m.IsSelectedWall(tile.Sprite) {
		return
	}
	if tile.Sprite == m.WallSolid {
		tile.Sprite = m.WallSelectedSolid
		return
	}
	tile.Sprite = m.WallSelectedExposed
}

func setToNormalInteractFalse(tile *m.Tile) {
	if m.IsWall(tile.Sprite) {
		return
	}
	if tile.Sprite == m.WallSelectedSolid {
		tile.Sprite = m.WallSolid
		return
	}
	tile.Sprite = m.WallExposed
}

func printMousePos(idx int) {
	x, y := gl.IdxToXY(idx)
	fmt.Printf("IDX: %d. XY: {%d, %d}.\n", idx, x, y)
}

func printDwarf(dwarves *[]*dwarf.Dwarf, currentMousePos int) {
	for _, dwarf := range *dwarves {
		if dwarf.Idx != currentMousePos {
			continue
		}
		fmt.Println(dwarf.String())
		return
	}
}
