package room

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

const MAX_STORAGE = 8

type StorageTile struct {
	*m.Tile
}

func (s *StorageTile) Available(res entity.Resource) bool {
	///////////////////
	// TODO
	// Switch amount on
	// resource type.
	///////////////////
	if res == s.Resource && s.ResourceAmount < MAX_STORAGE {
		return true
	}
	if s.ResourceAmount == 0 {
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
	for s.ResourceAmount < MAX_STORAGE && amount > 0 {
		s.ResourceAmount++
		amount--
	}
	return amount
}

// Returns up to the desiredAmount,
// reducing the stored tile amount.
func (s *StorageTile) Take(desiredAmount uint) uint {
	returnAmount := s.ResourceAmount
	if s.ResourceAmount < desiredAmount {
		return s.TakeAll()
	}
	s.ResourceAmount -= desiredAmount
	return returnAmount - desiredAmount
}

func (s *StorageTile) TakeAll() uint {
	all := s.ResourceAmount
	s.ResourceAmount = 0
	s.Resource = entity.ResourceNone
	return all
}

func (s *StorageTile) Remaining() uint {
	return MAX_STORAGE - s.ResourceAmount
}

func createStorageTiles(mp *m.Map, tiles []int) []StorageTile {
	var sts []StorageTile
	for _, idx := range tiles {
		var amount uint
		if mp.Items[idx].Resource != entity.ResourceNone {
			amount++
		}
		sts = append(sts, StorageTile{&mp.Items[idx]})
	}
	return sts
}
