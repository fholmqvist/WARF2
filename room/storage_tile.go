package room

import (
	"projects/games/warf2/resource"
	m "projects/games/warf2/worldmap"
)

type StorageTile struct {
	Idx    int
	Tpe    resource.Resource
	Amount uint
}

func (s *StorageTile) IsAvailable() bool {
	if s.Amount == 0 {
		s.Tpe = resource.None
	}
	return s.Tpe == resource.None && s.Amount == 0
}

func (s *StorageTile) IsEmpty() bool {
	return !s.IsAvailable()
}

func createStorageTiles(tt m.Tiles) []StorageTile {
	var st []StorageTile
	////////////////////////////////////
	// TODO
	// Use worldmap to determine if
	// we might be building on a tile
	// that already contains resources.
	// If so, that storage tile should
	// reflect that.
	////////////////////////////////////
	for _, t := range tt {
		st = append(st, StorageTile{
			Idx:    t.Idx,
			Tpe:    resource.None,
			Amount: 0,
		})
	}
	return st
}
