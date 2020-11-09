// Package worldmap handles all
// the in-game rendering of in-game
// tiles and their functionality.
package worldmap

// Map holds all the tiles
// for a game.
type Map struct {
	Tiles         []Tile `json:"t"`
	SelectedTiles []Tile `json:"s"`
	Items         []Tile `json:"i"`
}

// GetTile returns a pointer to the tile
// from the XY-indexed tile on the map,
// and a bool to determine if the
// function was successful.
func (m Map) GetTile(x, y int) (*Tile, bool) {
	return m.GetTileByIndex(x + y*TilesW)
}

// GetSelectionTile returns a pointer to
// the tile from the XY-indexed tile from
// the selected layer on map, and a bool
// to determine if the function was successful.
func (m Map) GetSelectionTile(x, y int) (*Tile, bool) {
	return m.GetSelectionTileByIndex(x + y*TilesW)
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

// ResetIslands resets the islands
// after using flood fill.
func (m *Map) ResetIslands() {
	for i := range m.Tiles {
		m.Tiles[i].Island = 0
	}
}

func getTileFrom(idx int, tiles []Tile) (*Tile, bool) {
	if idx < 0 || idx >= TilesT {
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

// Tile data struct
type Tile struct {
	Idx              int  `json:"i"`
	X                int  `json:"x"`
	Y                int  `json:"y"`
	Sprite           int  `json:"s"`
	Island           int  `json:"-"`
	Map              *Map `json:"-"`
	NeedsInteraction bool `json:"n"`
	Blocked          bool `json:"b"`
}

// CreateTile returns a new tile
// at the given index with the given sprite.
func CreateTile(idx, spr int, m *Map) Tile {
	return Tile{
		Idx:    idx,
		X:      IdxToX(idx),
		Y:      IdxToY(idx),
		Sprite: spr,
		Map:    m,
	}
}

// TileDir contains an index
// and the direction it is in
// relation to the index it was
// requested from.
type TileDir struct {
	Idx int
	Dir Direction
}

// IdxToXY returns the corresponding
// X and Y values for a given index.
func IdxToXY(idx int) (int, int) {
	return IdxToX(idx), IdxToY(idx)
}

// IdxToX returns the corresponding
// X value to for given index.
func IdxToX(idx int) int {
	return idx % TilesW
}

// IdxToY returns the corresponding
// Y value to for given index.
func IdxToY(idx int) int {
	return idx / TilesW
}

// XYToIdx returns the corresponding
// index based on the given X and Y values.
func XYToIdx(x, y int) int {
	return x + y*TilesW
}
