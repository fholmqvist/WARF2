// Package worldmap handles all
// the in-game rendering of in-game
// tiles and their functionality.
package worldmap

import (
	"projects/games/warf2/globals"
)

// Map holds all the tiles
// for a game.
type Map struct {
	Tiles         []Tile `json:"t"`
	SelectedTiles []Tile `json:"s"`
	Items         []Tile `json:"i"`
	Rails         []Tile `json:"r"`
}

func New() *Map {
	mp := &Map{}
	mp.Tiles = newTiles(mp, Ground)
	mp.SelectedTiles = newTiles(mp, None)
	mp.Items = newTiles(mp, None)
	mp.Rails = newRailTiles(mp, None)
	return mp
}

func NormalMap() *Map {
	mp := New()
	mp.Automata()
	mp.FillIslands(true)
	mp.FillIslands(false)
	mp.CreateBoundaryWalls()
	mp.FixWalls()
	return mp
}

func BoundariesMap() *Map {
	mp := New()
	mp.CreateBoundaryWalls()
	mp.FixWalls()
	return mp
}

func FilledMap() *Map {
	mp := New()
	mp.CreateBoundaryWalls()
	for _, t := range mp.Tiles {
		if IsAnyWall(t.Sprite) {
			continue
		}
		mp.Tiles[t.Idx].Sprite = WallSolid
	}
	mp.FixWalls()
	return mp
}

func newTiles(mp *Map, sprite int) []Tile {
	t := make([]Tile, globals.TilesW*globals.TilesH)
	for i := range t {
		t[i] = CreateTile(i, sprite, mp)
	}
	return t
}

func newRailTiles(mp *Map, sprite int) []Tile {
	t := make([]Tile, globals.TilesW*globals.TilesH)
	for i := range t {
		t[i] = CreateRailTile(i, mp)
	}
	return t
}

func (m *Map) Clear() {
	m.Tiles = newTiles(m, Ground)
	m.SelectedTiles = newTiles(m, None)
	m.Items = newTiles(m, None)
}

func (m *Map) ClearSelectedTiles() {
	m.SelectedTiles = newTiles(m, None)
}

// GetTile returns a pointer to the tile
// from the XY-indexed tile on the map,
// and a bool to determine if the
// function was successful.
func (m Map) GetTile(x, y int) (*Tile, bool) {
	return m.GetTileByIndex(x + y*globals.TilesW)
}

// GetSelectionTile returns a pointer to
// the tile from the XY-indexed tile from
// the selected layer on map, and a bool
// to determine if the function was successful.
func (m Map) GetSelectionTile(x, y int) (*Tile, bool) {
	return m.GetSelectionTileByIndex(x + y*globals.TilesW)
}

// GetTileByIndex returns a pointer
// to the tile from the map and a
// bool to determine if the
// function was successful.
func (m Map) GetTileByIndex(idx int) (*Tile, bool) {
	return getTileFrom(idx, m.Tiles)
}

// GetSelectionTileByIndex returns a
// pointer to the tile from the selected
// tiles layer and a bool to determine if
// the function was successful.
func (m Map) GetSelectionTileByIndex(idx int) (*Tile, bool) {
	return getTileFrom(idx, m.SelectedTiles)
}

// GetItemTileByIndex returns a pointer
// to the tile from the item layer and
// a bool to determine if the function
// was successful.
func (m Map) GetItemTileByIndex(idx int) (*Tile, bool) {
	return getTileFrom(idx, m.Items)
}

// GetItemTile returns a pointer
// to the tile from the item layer and
// a bool to determine if the function
// was successful.
func (m Map) GetItemTile(x, y int) (*Tile, bool) {
	idx := globals.XYToIdx(x, y)
	return getTileFrom(idx, m.Items)
}

// GetRailTileByIndex returns a pointer
// to the tile from the rail layer and
// a bool to determine if the function
// was successful.
func (m Map) GetRailTileByIndex(idx int) (*Tile, bool) {
	return getTileFrom(idx, m.Rails)
}

// GetRailTile returns a pointer
// to the tile from the rail layer and
// a bool to determine if the function
// was successful.
func (m Map) GetRailTile(x, y int) (*Tile, bool) {
	idx := globals.XYToIdx(x, y)
	return getTileFrom(idx, m.Rails)
}

// ResetIslands resets the islands
// after using flood fill.
func (m *Map) ResetIslands() {
	for i := range m.Tiles {
		m.Tiles[i].Island = 0
	}
}

// TilesForIsland returns the island
// tiles for a given island number.
func (m *Map) TilesForIsland(island int) []Tile {
	var islands []Tile
	for i := range m.Tiles {
		if m.Tiles[i].Island == island {
			islands = append(islands, m.Tiles[i])
		}
	}
	return islands
}

// ItemsFor returns all the
// items for a given set of tiles.
// func (m *Map) ItemsFor(tiles []Tile) []Tile {
// 	var items []Tile
// 	for _, tile := range tiles {
// 		if item.IsLibraryItem(m.Items[tile.Idx].Sprite) {
// 			items = append(items, m.Items[tile.Idx])
// 		}
// 	}
// 	return items
// }

func getTileFrom(idx int, tiles []Tile) (*Tile, bool) {
	if idx < 0 || idx >= globals.TilesT {
		return nil, false
	}
	return &tiles[idx], true
}

// getTileByIndexAndDirection returns a pointer
// to the tile from the map given the current index
// and a new direction. It also returns a boolean value
// to determine if the function was successful.
func (m Map) getTileByIndexAndDirection(idx int, dir Direction) (*Tile, bool) {
	t, ok := m.GetTileByIndex(IndexAtDirection(idx, dir))
	if !ok {
		return nil, false
	}
	return t, true
}
