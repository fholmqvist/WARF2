package worldmap

// Tile data struct.
type Tile struct {
	Idx      int     `json:"i"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
	Sprite   int     `json:"s"`
	Island   int     `json:"-"`
	Map      *Map    `json:"-"`
	Blocked  bool    `json:"b"`
	Rotation float64 `json:"r"`
}

type Tiles []Tile

func (t Tiles) Len() int           { return len(t) }
func (t Tiles) Less(i, j int) bool { return t[i].Idx < t[j].Idx }
func (t Tiles) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

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
