package room

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

const MAX_STORAGE = 8

/////////////////////////////////////
// TODO
// Remove and just use regular tiles.
// Code duplication, accidental
// complexity and broken hovering.
/////////////////////////////////////
type StorageTile struct {
	Idx int
	entity.Resource
	Amount uint
}

func (s *StorageTile) Available(res entity.Resource) bool {
	///////////////////
	// TODO
	// Switch amount on
	// resource type.
	///////////////////
	if res == s.Resource && s.Amount < MAX_STORAGE {
		return true
	}
	if s.Amount == 0 {
		s.Resource = entity.ResourceNone
		return true
	}
	return false
}

func (s *StorageTile) Unavailable(tpe entity.Resource) bool {
	return !s.Available(tpe)
}

// Adds the amount of resource to the tile,
// given that the tile is of that type
// (panics otherwise, enforcing hygenic caller).
// Remainder is returned to the caller.
func (s *StorageTile) Add(res entity.Resource, amount uint) (remaining uint) {
	if s.Resource != res && s.Resource != 0 {
		panic(fmt.Sprintf("storage: AddItem: trying to add %v to a tile with type of %v", res, s.Resource))
	}
	s.Resource = res
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

func createStorageTiles(mp *m.Map, tiles []int) []StorageTile {
	var st []StorageTile
	for _, idx := range tiles {
		var amount uint
		if mp.Items[idx].Resource != entity.ResourceNone {
			amount++
		}
		st = append(st, StorageTile{
			Idx:      idx,
			Resource: mp.Items[idx].Resource,
			Amount:   amount,
		})
	}
	return st
}
