package room

import (
	"fmt"
	"projects/games/warf2/resource"
	m "projects/games/warf2/worldmap"
)

type StorageTile struct {
	Idx    int
	Tpe    resource.Resource
	Amount uint
}

func (s *StorageTile) Available(tpe resource.Resource) bool {
	///////////////////
	// TODO
	// Switch amount on
	// resource type.
	///////////////////
	if tpe == s.Tpe && s.Amount <= 8 {
		return true
	}
	if s.Amount == 0 {
		s.Tpe = resource.None
	}
	return s.Tpe == resource.None && s.Amount == 0
}

func (s *StorageTile) Unavailable(tpe resource.Resource) bool {
	return !s.Available(tpe)
}

func (s *StorageTile) AddItem(r resource.Resource) {
	tpe := s.Tpe
	if tpe != r && tpe != 0 {
		panic(fmt.Sprintf("storage: AddItem: trying to add %v to a tile with type of %v", r, tpe))
	}
	s.Tpe = r
	s.Amount++
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
