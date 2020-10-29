package gmap

// Map holds all the tiles
// for a game.
type Map struct {
	Tiles []Tile
}

// GetTile returns a pointer to a tile
// from the XY-indexed tile on the map,
// and a bool to determine whether the
// function was successful.
func (m Map) GetTile(x, y int) (*Tile, bool) {
	return m.GetTileByIndex(x + y*TilesW)
}

// GetTileByIndex returns a pointer
// to a tile from the map and a bool
// to determine whether the function
// was successful.
func (m Map) GetTileByIndex(idx int) (*Tile, bool) {
	if idx < 0 || idx >= TilesT {
		return nil, false
	}
	return &m.Tiles[idx], true
}

// Tile data struct
type Tile struct {
	Idx    int
	X, Y   int
	Sprite int
	Island int
	Map    *Map
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
