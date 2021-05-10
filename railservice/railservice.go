package rail

// This currently assumes some responsibilities
// that should perhaps be delegated to Map?

import (
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
	/////////////////////////////////////////////////
	// TODO
	// Check which rail to place.
	/////////////////////////////////////////////////
	rt.Sprite = Straight
	r.FixRails()
}

func (r *RailService) PlaceRailXY(x, y int) {
	idx := mp.XYToIdx(x, y)
	r.PlaceRail(idx)
}

func (r *RailService) FixRails() {
	for idx, rt := range r.Map.Rails {
		if rt.Sprite == 0 {
			continue
		}
		const (
			up    = 0
			right = 1
			down  = 2
			left  = 3
		)
		rails := []mp.Tile{
			r.Map.OneRailUp(rt.Idx),
			r.Map.OneRailRight(rt.Idx),
			r.Map.OneRailDown(rt.Idx),
			r.Map.OneRailLeft(rt.Idx),
		}
		if IsRail(r.Map, rails[left].Idx) && IsRail(r.Map, rails[right].Idx) {
			r.Map.Rails[idx].Rotation = math.Pi * 1.5
		}
	}
}
