package worldmap

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/entity"
	gl "github.com/Holmqvist1990/WARF2/globals"
)

// Tile data struct.
type Tile struct {
	TileType
	Idx            int             `json:"i"`
	X              int             `json:"x"`
	Y              int             `json:"y"`
	Sprite         int             `json:"s"`
	Island         int             `json:"-"`
	Map            *Map            `json:"-"`
	Rotation       float64         `json:"rt"`
	Resource       entity.Resource `json:"rs"`
	ResourceAmount uint            `json:"rsa"`
	Room           interface{}     // *room.Room, import cycle
}

func (t *Tile) String() string {
	return fmt.Sprintf("IDX: %v. SPRITE: %v. RESOURCE: %v. AMOUNT: %v.",
		t.Idx, entity.ItemToString(t.Sprite), t.Resource, t.ResourceAmount)
}

type Tiles []Tile

func (t Tiles) Len() int           { return len(t) }
func (t Tiles) Less(i, j int) bool { return t[i].Idx < t[j].Idx }
func (t Tiles) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func (t Tiles) ToIdxs() []int {
	idxs := make([]int, len(t))
	for i, tile := range t {
		idxs[i] = tile.Idx
	}
	return idxs
}

// Returns a new tile at the
// given index with the given sprite.
func CreateTile(idx, spr int, m *Map) Tile {
	return Tile{
		Idx:    idx,
		X:      gl.IdxToX(idx),
		Y:      gl.IdxToY(idx),
		Sprite: spr,
		Map:    m,
	}
}

// Returns a new RailTile at the given index.
func CreateRailTile(idx int, m *Map) Tile {
	t := CreateTile(idx, None, m)
	t.TileType = RailTile
	return t
}

type TileType int

const (
	NormalTile TileType = iota
	RailTile
)
