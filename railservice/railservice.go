package rail

// This currently assumes some responsibilities
// that should perhaps be delegated to Map?
//
// Or perhaps, is this more correct, and we need
// an ItemService as well? I think this sounds even
// better to be honest.

import (
	"projects/games/warf2/globals"
	mp "projects/games/warf2/worldmap"
)

const (
	None = iota
	Straight
	Curve
	Stop
	Cross
)

type RailService struct {
	Map *mp.Map
}

func (r *RailService) PlaceRail(idx int) {
	r.PlaceRails([]int{idx})
}

func (r *RailService) PlaceRails(idxs []int) {
	min := globals.TilesT + 1
	max := -1
	for _, idx := range idxs {
		t, ok := r.Map.GetTileByIndex(idx)
		if !ok {
			return
		}
		if mp.Blocking(t) {
			return
		}
		rt, ok := r.Map.GetRailTileByIndex(idx)
		if !ok {
			return
		}
		if rt.Sprite != None {
			return
		}
		rt.Sprite = Straight
		if idx < min {
			min = idx
		}
		if max < idx {
			max = idx
		}
	}
	r.FixRails(min, max)
}

func (r *RailService) PlaceRailXY(x, y int) {
	idx := mp.XYToIdx(x, y)
	r.PlaceRails([]int{idx})
}

func (r *RailService) PlaceRailsXY(xys [][2]int) {
	var idxs []int
	for i := range xys {
		idxs = append(idxs, mp.XYToIdx(xys[i][0], xys[i][1]))
	}
	r.PlaceRails(idxs)
}
