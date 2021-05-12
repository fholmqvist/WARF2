package rail

// This currently assumes some responsibilities
// that should perhaps be delegated to Map?
//
// Or perhaps, is this more correct, and we need
// an ItemService as well? I think this sounds even
// better to be honest.

import (
	"fmt"
	"math"
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
	min := mp.TilesT + 1
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
	fmt.Println(min, max)
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

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

func (r *RailService) FixRails(min, max int) {
	for idx := min; idx <= max; idx++ {
		if r.Map.Rails[idx].Sprite == 0 {
			continue
		}
		rails := []mp.Tile{
			r.Map.OneRailUp(idx),
			r.Map.OneRailRight(idx),
			r.Map.OneRailDown(idx),
			r.Map.OneRailLeft(idx),
		}
		// CROSS
		if IsRail(r.Map, rails[up].Idx) && IsRail(r.Map, rails[right].Idx) &&
			IsRail(r.Map, rails[down].Idx) && IsRail(r.Map, rails[left].Idx) {
			r.Map.Rails[idx].Sprite = Cross
			continue
		}
		// HORIZONTAL
		if IsRail(r.Map, rails[left].Idx) && IsRail(r.Map, rails[right].Idx) {
			r.Map.Rails[idx].Rotation = math.Pi * 1.5
			continue
		}
		// UP RIGHT
		if IsRail(r.Map, rails[up].Idx) && IsRail(r.Map, rails[right].Idx) {
			r.Map.Rails[idx].Sprite = Curve
			continue
		}
		// UP LEFT
		if IsRail(r.Map, rails[up].Idx) && IsRail(r.Map, rails[left].Idx) {
			r.Map.Rails[idx].Sprite = Curve
			r.Map.Rails[idx].Rotation = math.Pi * 1.5
			continue
		}
		// DOWN RIGHT
		if IsRail(r.Map, rails[down].Idx) && IsRail(r.Map, rails[right].Idx) {
			r.Map.Rails[idx].Sprite = Curve
			r.Map.Rails[idx].Rotation = -math.Pi * 1.5
			continue
		}
		// DOWN LEFT
		if IsRail(r.Map, rails[down].Idx) && IsRail(r.Map, rails[left].Idx) {
			r.Map.Rails[idx].Sprite = Curve
			r.Map.Rails[idx].Rotation = math.Pi * 3.0
			continue
		}

		// LONE UP
		if !IsRail(r.Map, rails[up].Idx) && IsRail(r.Map, rails[down].Idx) {
			r.Map.Rails[idx].Sprite = Stop
			continue
		}
		// LONE RIGHT
		if !IsRail(r.Map, rails[right].Idx) && IsRail(r.Map, rails[left].Idx) {
			r.Map.Rails[idx].Sprite = Stop
			r.Map.Rails[idx].Rotation = -math.Pi * 1.5
			continue
		}
		// LONE DOWN
		if !IsRail(r.Map, rails[down].Idx) && IsRail(r.Map, rails[up].Idx) {
			r.Map.Rails[idx].Sprite = Stop
			r.Map.Rails[idx].Rotation = math.Pi * 3.0
			continue
		}
		// LONE LEFT
		if !IsRail(r.Map, rails[left].Idx) && IsRail(r.Map, rails[right].Idx) {
			r.Map.Rails[idx].Sprite = Stop
			r.Map.Rails[idx].Rotation = math.Pi * 1.5
			continue
		}
	}
}
