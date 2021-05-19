package room

import (
	"fmt"
	"projects/games/warf2/resource"
	m "projects/games/warf2/worldmap"
)

const MAX_STORAGE = 8

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
	if tpe == s.Tpe && s.Amount < MAX_STORAGE {
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

// Adds the amount of resource to the tile,
// given that the tile is of that type
// (panics otherwise, enforcing hygenic caller).
// Remainder is returned to the caller.
func (s *StorageTile) Add(r resource.Resource, amount uint) (remaining uint) {
	tpe := s.Tpe
	if tpe != r && tpe != 0 {
		panic(fmt.Sprintf("storage: AddItem: trying to add %v to a tile with type of %v", r, tpe))
	}
	s.Tpe = r
	for s.Amount < MAX_STORAGE && amount > 0 {
		s.Amount++
		amount--
	}
	return amount
}

// Returns up to the desiredAmount,
// reducing the stored tile amount.
func (s *StorageTile) Take(desiredAmount uint) uint {
	returnAmount := s.Amount
	if s.Amount < desiredAmount {
		s.Amount = 0
		return returnAmount
	}
	s.Amount -= desiredAmount
	return returnAmount - desiredAmount
}

func (s *StorageTile) Remaining() uint {
	return MAX_STORAGE - s.Amount
}

func createStorageTiles(tt m.Tiles, itt m.Tiles) []StorageTile {
	var st []StorageTile
	for _, t := range tt {
		var amount uint
		if itt[t.Idx].Resource != resource.None {
			amount++
		}
		st = append(st, StorageTile{
			Idx:    t.Idx,
			Tpe:    itt[t.Idx].Resource,
			Amount: amount,
		})
	}
	return st
}
