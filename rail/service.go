package rail

// This currently assumes some responsibilities
// that should perhaps be delegated to Map?
//
// Or perhaps, is this more correct, and we need
// an ItemService as well? I think this sounds even
// better to be honest.

import (
	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Service struct {
	Carts []*Cart
	Map   *m.Map
}

func NewService(mp *m.Map) *Service {
	return &Service{Map: mp}
}

func (r *Service) Update(mp *m.Map) {
	for _, cart := range r.Carts {
		cart.traversePath(mp)
	}
}

func (r *Service) PlaceRail(idx int) {
	r.PlaceRails([]int{idx})
}

func (r *Service) PlaceRails(idxs []int) {
	min := globals.TilesT + 1
	max := -1
	for _, idx := range idxs {
		t, ok := r.Map.GetTileByIndex(idx)
		if !ok {
			continue
		}
		itemT, ok := r.Map.GetItemTileByIndex(idx)
		if !ok {
			continue
		}
		if m.Blocking(t, itemT) {
			continue
		}
		rt, ok := r.Map.GetRailTileByIndex(idx)
		if !ok {
			continue
		}
		if rt.Sprite != m.None {
			continue
		}
		rt.Sprite = m.Straight
		if idx < min {
			min = idx
		}
		if max < idx {
			max = idx
		}
	}
	r.FixRails(min, max)
}

func (r *Service) PlaceRailXY(x, y int) {
	idx := globals.XYToIdx(x, y)
	r.PlaceRails([]int{idx})
}

func (r *Service) PlaceRailsXY(xys [][2]int) {
	var idxs []int
	for i := range xys {
		idxs = append(idxs, globals.XYToIdx(xys[i][0], xys[i][1]))
	}
	r.PlaceRails(idxs)
}
